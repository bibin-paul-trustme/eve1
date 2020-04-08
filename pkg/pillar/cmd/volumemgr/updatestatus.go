// Copyright (c) 2020 Zededa, Inc.
// SPDX-License-Identifier: Apache-2.0

package volumemgr

import (
	"time"

	"github.com/lf-edge/eve/pkg/pillar/pubsub"
	"github.com/lf-edge/eve/pkg/pillar/types"
	"github.com/satori/go.uuid"
	log "github.com/sirupsen/logrus"
)

// Returns changed
// XXX remove "done" boolean return?
func doUpdate(ctx *volumemgrContext, status *types.VolumeStatus) (bool, bool) {

	log.Infof("doUpdate(%s) name %s", status.Key(), status.DisplayName)
	status.WaitingForCerts = false

	// Anything to do?
	if status.State == types.DELIVERED {
		log.Infof("doUpdate(%s) name %s nothing to do",
			status.Key(), status.DisplayName)
		return false, true
	}

	// Check if Verified Status already exists.
	vs, changed := lookForVerified(ctx, status)
	if vs != nil {
		log.Infof("Found %s based on VolumeID %s sha %s",
			status.DisplayName, status.VolumeID, status.BlobSha256)

		// XXX should not happen once ResolveConfig in place
		// XXX messes up container since end up with different object
		// XXX temporary container field HACK
		// XXX create logic...
		if status.ContainerSha256 != status.BlobSha256 {
			status.ContainerSha256 = status.BlobSha256
		}
		// XXX compare logic ...
		if status.ContainerSha256 != vs.ImageSha256 {
			log.Infof("updating image sha from %s to %s",
				status.ContainerSha256, vs.ImageSha256)
			status.ContainerSha256 = vs.ImageSha256
			changed = true
		}
		if status.State != vs.State {
			status.State = vs.State
			changed = true
		}
		if vs.Pending() {
			log.Infof("lookupVerifyImageStatus %s Pending\n", status.VolumeID)
			return changed, false
		}
		if vs.LastErr != "" {
			log.Errorf("Received error from verifier for %s: %s\n",
				status.VolumeID, vs.LastErr)
			status.ErrorSource = pubsub.TypeToName(types.VerifyImageStatus{})
			status.LastErr = vs.LastErr
			status.LastErrTime = vs.LastErrTime
			changed = true
			return changed, false
		} else if status.ErrorSource == pubsub.TypeToName(types.VerifyImageStatus{}) {
			log.Infof("Clearing verifier error %s\n", status.LastErr)
			status.LastErr = ""
			status.ErrorSource = ""
			status.LastErrTime = time.Time{}
			changed = true
		}
		switch vs.State {
		case types.INITIAL:
			// Nothing to do
		case types.DELIVERED:
			if status.FileLocation != vs.FileLocation {
				status.FileLocation = vs.FileLocation
				log.Infof("Update SSL FileLocation for %s: %s\n",
					status.Key(), status.FileLocation)
				changed = true
			}
			c, err := createVolume(ctx, status, vs.FileLocation)
			if err != nil {
				status.ErrorSource = pubsub.TypeToName(types.VolumeStatus{})
				status.LastErr = err.Error()
				status.LastErrTime = time.Now()
				changed = true
				return changed, false
			}
			if c {
				changed = true
			}
		default:
			if status.FileLocation != vs.FileLocation {
				status.FileLocation = vs.FileLocation
				log.Infof("Update SSL FileLocation for %s: %s\n",
					status.Key(), status.FileLocation)
				changed = true
			}
		}
		return changed, false
	}

	log.Infof("VerifyImageStatus %s for %s sha %s not found",
		status.DisplayName, status.VolumeID,
		status.BlobSha256)

	// Make sure we kick the downloader and have a refcount
	if !status.DownloadOrigin.HasDownloaderRef {
		AddOrRefcountDownloaderConfig(ctx, *status)
		status.DownloadOrigin.HasDownloaderRef = true
		changed = true
	}
	// Check if we have a DownloadStatus if not put a DownloadConfig
	// in place
	ds := lookupDownloaderStatus(ctx, status.ObjType, status.VolumeID)
	if ds == nil || ds.Expired || ds.RefCount == 0 {
		if ds == nil {
			log.Infof("downloadStatus not found. name: %s", status.VolumeID)
		} else if ds.Expired {
			log.Infof("downloadStatus Expired set. name: %s", status.VolumeID)
		} else {
			log.Infof("downloadStatus RefCount=0. name: %s", status.VolumeID)
		}
		status.State = types.DOWNLOAD_STARTED
		changed = true
		return changed, false
	}
	if ds.FileLocation != "" && status.FileLocation == "" {
		status.FileLocation = ds.FileLocation
		changed = true
		log.Infof("From ds set FileLocation to %s for %s",
			ds.FileLocation, status.VolumeID)
	}
	if status.State != ds.State {
		status.State = ds.State
		changed = true
	}
	if ds.Progress != status.Progress {
		status.Progress = ds.Progress
		changed = true
	}
	if ds.Pending() {
		log.Infof("lookupDownloaderStatus %s Pending\n",
			status.VolumeID)
		return changed, false
	}
	if ds.LastErr != "" {
		log.Errorf("Received error from downloader for %s: %s\n",
			status.VolumeID, ds.LastErr)
		status.ErrorSource = pubsub.TypeToName(types.DownloaderStatus{})
		status.LastErr = ds.LastErr
		status.LastErrTime = ds.LastErrTime
		changed = true
		return changed, false
	} else if status.ErrorSource == pubsub.TypeToName(types.DownloaderStatus{}) {
		log.Infof("Clearing downloader error %s\n", status.LastErr)
		status.LastErr = ""
		status.ErrorSource = ""
		status.LastErrTime = time.Time{}
		changed = true
	}
	switch ds.State {
	case types.INITIAL:
		// Nothing to do
	case types.DOWNLOAD_STARTED:
		// Nothing to do
	case types.DOWNLOADED:
		// Kick verifier to start if it hasn't already; add RefCount
		c := kickVerifier(ctx, status, true)
		if c {
			changed = true
		}
	}
	if status.WaitingForCerts {
		log.Infof("Waiting for certs for %s\n", status.Key())
		return changed, false
	}
	log.Infof("Waiting for download for %s\n", status.Key())
	return changed, false
}

// Returns changed
// Updates status with WaitingForCerts if checkCerts is set
func kickVerifier(ctx *volumemgrContext, status *types.VolumeStatus, checkCerts bool) bool {
	changed := false
	if !status.DownloadOrigin.HasVerifierRef {
		ret, errInfo := MaybeAddVerifyImageConfig(ctx, *status, checkCerts)
		if ret {
			status.DownloadOrigin.HasVerifierRef = true
			changed = true
		} else {
			// if errors, set the certError flag
			// otherwise, mark as waiting for certs
			if errInfo.Error != "" {
				status.ErrorSource = errInfo.ErrorSource
				status.LastErr = errInfo.Error
				status.LastErrTime = errInfo.ErrorTime
				changed = true
			} else {
				if !status.WaitingForCerts {
					changed = true
					status.WaitingForCerts = true
				}
			}
		}
	}
	return changed
}

// lookForVerified handles the split between PersistImageStatus and
// VerifyImageStatus. If it only finds the Persist it returns nil but
// sets up a VerifyImageConfig.
// Also returns changed=true if the VolumeStatus is changed
func lookForVerified(ctx *volumemgrContext, status *types.VolumeStatus) (*types.VerifyImageStatus, bool) {
	changed := false
	vs := lookupVerifyImageStatus(ctx, status.ObjType, status.VolumeID)
	if vs == nil {
		ps := lookupPersistImageStatus(ctx, status.ObjType, status.BlobSha256)
		if ps == nil || ps.Expired {
			log.Infof("Verify/PersistImageStatus for %s sha %s not found",
				status.VolumeID, status.BlobSha256)
		} else {
			log.Infof("Found %s based on ImageSha256 %s VolumeID %s",
				status.DisplayName, status.BlobSha256, status.VolumeID)
			if status.State != types.DOWNLOADED {
				status.State = types.DOWNLOADED
				status.Progress = 100
				changed = true
			}
			// If we don't already have a RefCount add one
			if !status.DownloadOrigin.HasVerifierRef {
				log.Infof("!HasVerifierRef")
				// We don't need certs since Status already exists
				MaybeAddVerifyImageConfig(ctx, *status, false)
				status.DownloadOrigin.HasVerifierRef = true
				changed = true
			}
			// Wait for VerifyImageStatus to appear
			return nil, changed
		}
	} else {
		log.Infof("Found %s based on VolumeID %s sha %s",
			status.DisplayName, status.VolumeID, status.BlobSha256)
		// If we don't already have a RefCount add one
		// No need to checkCerts since we have a VerifyImageStatus
		c := kickVerifier(ctx, status, false)
		if c {
			changed = true
		}
		return vs, changed
	}
	return vs, changed
}

// Find all the VolumeStatus which refer to this VolumeID
func updateVolumeStatus(ctx *volumemgrContext, objType string, volumeID uuid.UUID) {

	log.Infof("updateVolumeStatus(%s) objType %s", volumeID, objType)
	found := false
	pub := ctx.publication(types.VolumeStatus{}, objType)
	items := pub.GetAll()
	for _, st := range items {
		status := st.(types.VolumeStatus)
		if status.VolumeID == volumeID {
			log.Infof("Found VolumeStatus %s", status.Key())
			found = true
			changed, _ := doUpdate(ctx, &status)
			if changed {
				publishVolumeStatus(ctx, &status)
			}
		}
	}
	if !found {
		log.Warnf("XXX updateVolumeStatus(%s) objType %s NOT FOUND",
			volumeID, objType)
	}
}

// doDelete returns changed boolean
// XXX need return value?
func doDelete(ctx *volumemgrContext, status *types.VolumeStatus) bool {
	changed := false

	// XXX support other types
	if status.Origin == types.OriginTypeDownload {
		if status.DownloadOrigin.HasDownloaderRef {
			MaybeRemoveDownloaderConfig(ctx, status.ObjType,
				status.VolumeID) // XXX volumeID?
			status.DownloadOrigin.HasDownloaderRef = false
			changed = true
		}
		if status.DownloadOrigin.HasVerifierRef {
			MaybeRemoveVerifyImageConfig(ctx, status.ObjType,
				status.VolumeID) // XXX volumeID?
			status.DownloadOrigin.HasVerifierRef = false
			changed = true
		}
	}

	if status.VolumeCreated {
		_, err := destroyVolume(ctx, status)
		if err != nil {
			status.ErrorSource = pubsub.TypeToName(types.VolumeStatus{})
			status.LastErr = err.Error()
			status.LastErrTime = time.Now()
			changed = true
			return changed
		}
		status.VolumeCreated = false
		changed = true
	}
	return changed
}
