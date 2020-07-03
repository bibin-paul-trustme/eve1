// Copyright (c) 2020 Zededa, Inc.
// SPDX-License-Identifier: Apache-2.0

package volumemgr

import (
	"fmt"
	"strings"
	"time"

	zconfig "github.com/lf-edge/eve/api/go/config"
	"github.com/lf-edge/eve/pkg/pillar/containerd"
	"github.com/lf-edge/eve/pkg/pillar/types"
	"github.com/lf-edge/eve/pkg/pillar/utils"
	uuid "github.com/satori/go.uuid"
	log "github.com/sirupsen/logrus"
)

// doUpdate handles any updates to a VolumeStatus, called by any event handlers
// that either capture a VolumeStatus change or create one and want it processed.
// Returns changed
// XXX remove "done" boolean return?
func doUpdateContentTree(ctx *volumemgrContext, status *types.ContentTreeStatus) (bool, bool) {

	log.Infof("doUpdateContentTree(%s) name %s state %s", status.Key(), status.DisplayName, status.State)
	status.WaitingForCerts = false

	changed := false
	addedBlobs := []string{}
	sv := SignatureVerifier{
		Signature:        status.ImageSignature,
		PublicKey:        status.SignatureKey,
		CertificateChain: status.CertificateChain,
	}

	if status.State < types.VERIFIED {

		if status.IsContainer() {
			maybeLatchContentTreeHash(ctx, status)
			if status.ContentSha256 == "" {
				rs := lookupResolveStatus(ctx, status.ResolveKey())
				if rs == nil {
					log.Infof("Resolve status not found for %s",
						status.ContentID)
					status.HasResolverRef = true
					MaybeAddResolveConfig(ctx, *status)
					status.State = types.RESOLVING_TAG
					changed = true
					return changed, false
				}
				log.Infof("Processing ResolveStatus for content tree (%v)", status.ContentID)
				status.State = types.RESOLVED_TAG
				changed = true
				if rs.HasError() {
					errStr := fmt.Sprintf("Received error from resolver for %s, SHA (%s): %s",
						status.ResolveKey(), rs.ImageSha256, rs.Error)
					log.Error(errStr)
					status.SetErrorWithSource(errStr, types.ResolveStatus{},
						rs.ErrorTime)
					changed = true
					return changed, false
				} else if rs.ImageSha256 == "" {
					errStr := fmt.Sprintf("Received empty SHA from resolver for %s, SHA (%s): %s",
						status.ResolveKey(), rs.ImageSha256, rs.Error)
					log.Error(errStr)
					status.SetErrorWithSource(errStr, types.ResolveStatus{},
						rs.ErrorTime)
					changed = true
					return changed, false
				} else if status.IsErrorSource(types.ResolveStatus{}) {
					log.Infof("Clearing resolver error %s", status.Error)
					status.ClearErrorWithSource()
					changed = true
				}
				foundSha := strings.ToLower(rs.ImageSha256)
				log.Infof("Added Image SHA (%s) for content tree (%s)",
					foundSha, status.ContentID)
				status.ContentSha256 = foundSha
				status.HasResolverRef = false
				status.RelativeURL = maybeInsertSha(status.RelativeURL, status.ContentSha256)
				latchContentTreeHash(ctx, status.ContentID,
					status.ContentSha256, uint32(status.GenerationCounter))
				maybeLatchContentTreeHash(ctx, status)
				deleteResolveConfig(ctx, rs.Key())
				changed = true

			}

			// at this point, we have a resolved tag
			if status.State < types.RESOLVED_TAG {
				status.State = types.RESOLVED_TAG
				changed = true
			}

			// at this point, we will have the hash of the root blob as status.ContentSha256,
			// so we need to create the BlobStatus, if it does not exist
			// if it does not already exist
			if lookupOrCreateBlobStatus(ctx, sv, status.ObjType, status.ContentSha256) == nil {
				rootBlob := &types.BlobStatus{
					DatastoreID: status.DatastoreID,
					RelativeURL: status.RelativeURL,
					Sha256:      status.ContentSha256,
					Size:        status.MaxDownloadSize,
					State:       types.INITIAL,
					BlobType:    types.BlobUnknown,
					ObjType:     status.ObjType,
				}
				log.Infof("doUpdateContentTree: publishing root BlobStatus (%s) for content tree (%s)",
					status.ContentSha256, status.ContentID)
				publishBlobStatus(ctx, rootBlob)
			}
			if len(status.Blobs) == 0 {
				status.Blobs = append(status.Blobs, status.ContentSha256)
			}
		}

		// at this point, we at least are downloading
		if status.State < types.DOWNLOADING {
			status.State = types.DOWNLOADING
			changed = true
		}

		// loop through each blob, see if it is downloaded and verified.
		// we set the contenttree to verified when all of the blobs are verified
		if status.State < types.DOWNLOADING {
			status.State = types.DOWNLOADING
			changed = true
		}
		leftToProcess := false

		var (
			currentSize, totalSize int64
			blobErrors             = []string{}
			blobErrorTime          time.Time
		)
		for _, blobSha := range status.Blobs {
			// get the actual blobStatus
			blob := lookupOrCreateBlobStatus(ctx, sv, status.ObjType, blobSha)
			if blob == nil {
				log.Errorf("doUpdateContentTree: could not find BlobStatus(%s)", blobSha)
				leftToProcess = true
				continue
			}
			totalSize += blob.TotalSize
			currentSize += blob.CurrentSize

			// now the type should not be unknown (unless it is in error state)
			// these calls might update Blob.State hence we check
			// sequentially
			if blob.State <= types.DOWNLOADING {
				// any state less than downloaded, we ask for download, so that we have the refcount;
				// downloadBlob() is smart enough to look for existing references
				log.Infof("doUpdateContentTree: blob sha %s download state %v less than DOWNLOADED", blob.Sha256, blob.State)
				if downloadBlob(ctx, status.ObjType, sv, blob) {
					publishBlobStatus(ctx, blob)
					changed = true
				}
			}
			if blob.State == types.DOWNLOADED || blob.State == types.VERIFYING {
				// downloaded: kick off verifier for this blob
				log.Infof("doUpdateContentTree: blob sha %s download state %v less than VERIFIED", blob.Sha256, blob.State)
				if verifyBlob(ctx, status.ObjType, sv, blob) {
					publishBlobStatus(ctx, blob)
					changed = true
				}
			}
			if blob.State != types.VERIFIED {
				leftToProcess = true
				log.Errorf("doUpdateContentTree: left to process due to state '%s' for content blob %s",
					blob.State, blob.Sha256)
			} else {
				log.Infof("doUpdateContentTree: blob sha %s download state VERIFIED", blob.Sha256)
				if !blob.HasPersistRef {
					log.Infof("doUpdateContentTree: Adding PersistImageStatus reference for blob: %s",
						blob.Sha256)
					AddOrRefCountPersistImageStatus(ctx, "",
						blob.ObjType, blob.Path,
						blob.Sha256, int64(blob.Size))
					blob.HasPersistRef = true
				}
				// if verified, check for any children and start them off
				// resolve any unknown types and get manifests of index, or children of manifest
				blobType, err := resolveBlobType(blob)
				if blobType != blob.BlobType || err != nil {
					blob.BlobType = blobType
					publishBlobStatus(ctx, blob)
					changed = true
				}
				if err != nil {
					log.Infof("doUpdateContentTree(%s): error resolving blob type: %v", blob.Sha256, err)
					blob.SetError(err.Error(), time.Now())
					publishBlobStatus(ctx, blob)
					changed = true
				}
				blobChildren := blobsNotInList(getBlobChildren(blob), status.Blobs)
				if len(blobChildren) > 0 {
					log.Infof("doUpdateContentTree: adding %d children", len(blobChildren))
					// add all of the children
					for _, blob := range blobChildren {
						addedBlobs = append(addedBlobs, blob.Sha256)
					}
					status.Blobs = append(status.Blobs, addedBlobs...)
					// only publish those that do not already exist
					publishBlobStatus(ctx, blobsNotInStatusOrCreate(ctx, sv, status.ObjType, blobChildren)...)
				}
				if blob.BlobType == types.BlobManifest {
					size := resolveManifestSize(*blob)
					if size != status.TotalSize {
						status.TotalSize = size
						changed = true
					}
				}
			}
			// if any errors, catch them
			// Note that the downloadBlob above could have cleared
			// previous errors due to a retry hence we check for
			// errors here at the end
			if blob.HasError() {
				log.Errorf("doUpdateContentTree: BlobStatus(%s) has error: %s", blobSha, blob.Error)
				blobErrors = append(blobErrors, blob.Error)
				if blob.ErrorTime.After(blobErrorTime) {
					blobErrorTime = blob.ErrorTime
				}
				leftToProcess = true
			}
		}

		// Check if sizes changed before setting changed
		if status.CurrentSize != currentSize || status.TotalSize != totalSize {
			changed = true
			status.CurrentSize = currentSize
			status.TotalSize = totalSize
			if status.TotalSize > 0 {
				status.Progress = uint(status.CurrentSize / status.TotalSize * 100)
			}
			log.Infof("doUpdateContentTree: updating CurrentSize/TotalSize/Progress %d/%d/%d",
				currentSize, totalSize, status.Progress)
		}

		// update errors from blobs to status
		if len(blobErrors) != 0 {
			status.SetError(strings.Join(blobErrors, " / "), blobErrorTime)
			log.Infof("doUpdateContentTree(%s) had errors: %v", status.Key(), status.Error)
			changed = true
		} else if status.HasError() {
			log.Infof("doUpdateContentTree(%s) clearing errors", status.Key())
			status.ClearErrorWithSource()
			changed = true
		}

		// if we added any blobs, we need to reprocess this
		if len(addedBlobs) > 0 {
			log.Infof("doUpdateContentTree(%s) rerunning with added blobs: %v", status.Key(), addedBlobs)
			return doUpdateContentTree(ctx, status)
		}

		// if there are any left to process, do not do anything else
		// the rest of this flow should happen only when every part of the content tree
		// is downloaded and verified
		if leftToProcess {
			log.Infof("doUpdateContentTree(%s) leftToProcess=true, so returning `true,false`", status.Key())
			return true, false
		}

		// if we made it this far, the entire tree has been verified
		// before we mark it as verified, load it into the CAS store
		// for now, we only load containers into containerd
		// TODO: next stage, load disk images as well
		if status.Format == zconfig.Format_CONTAINER {
			// get all of the blobs
			if err := containerd.LoadBlobs(lookupBlobStatuses(ctx, status.Blobs...), status.RelativeURL); err != nil {
				log.Errorf("doUpdateContentTree(%s) unable to load blobs into containerd: %v", status.Key(), err)
				status.SetErrorWithSource(fmt.Sprintf("unable to load blobs into containerd: %v", err),
					types.OldVolumeStatus{}, time.Now())
				return changed, false
			}
			log.Infof("doUpdateContentTree(%s) successfully loaded all blobs into containerd", status.Key())
		}
		status.State = types.VERIFIED
		changed = true
	}

	return changed, status.State == types.VERIFIED
}

// Returns changed
// XXX remove "done" boolean return?
func doUpdateVol(ctx *volumemgrContext, status *types.VolumeStatus) (bool, bool) {

	log.Infof("doUpdateVol(%s) name %s", status.Key(), status.DisplayName)

	// Anything to do?
	if status.State == types.CREATED_VOLUME {
		log.Infof("doUpdateVol(%s) name %s nothing to do",
			status.Key(), status.DisplayName)
		return false, true
	}
	changed := false
	switch status.VolumeContentOriginType {
	case zconfig.VolumeContentOriginType_VCOT_BLANK:
		// XXX TBD
		errStr := fmt.Sprintf("doUpdateVol(%s) name %s: Volume content origin type %v is not implemeted yet.",
			status.Key(), status.DisplayName, status.VolumeContentOriginType)
		status.SetErrorWithSource(errStr,
			types.VolumeStatus{}, time.Now())
		changed = true
		return changed, false
	case zconfig.VolumeContentOriginType_VCOT_DOWNLOAD:
		ctStatus := lookupContentTreeStatus(ctx, status.ContentID.String())
		if ctStatus == nil {
			// Content tree not available
			log.Errorf("doUpdateVol(%s) name %s: waiting for content tree status %v",
				status.Key(), status.DisplayName, status.ContentID)
			errStr := fmt.Sprintf("ContentTreeStatus(%s) attached to volume %s is nil",
				status.ContentID.String(), status.DisplayName)
			status.SetErrorWithSource(errStr, types.ContentTreeStatus{}, time.Now())
			changed = true
			return changed, false
		} else if status.IsErrorSource(types.ContentTreeStatus{}) {
			log.Infof("doUpdate: Clearing volume error %s", status.Error)
			status.ClearErrorWithSource()
			changed = true
		}
		if status.Progress != ctStatus.Progress ||
			status.TotalSize != ctStatus.TotalSize ||
			status.CurrentSize != ctStatus.CurrentSize {
			status.Progress = ctStatus.Progress
			status.TotalSize = ctStatus.TotalSize
			status.CurrentSize = ctStatus.CurrentSize
			changed = true
		}
		if status.State != ctStatus.State {
			status.State = ctStatus.State
			changed = true
		}
		if ctStatus.State < types.VERIFIED {
			// Waiting for content tree to be processed
			if ctStatus.HasError() {
				log.Errorf("doUpdateVol(%s) name %s: content tree status has following error %v",
					status.Key(), status.DisplayName, ctStatus.Error)
				errStr := fmt.Sprintf("Found error in content tree %s attached to volume %s: %v",
					ctStatus.DisplayName, status.DisplayName, ctStatus.Error)
				status.SetErrorWithSource(errStr, types.ContentTreeStatus{}, time.Now())
				changed = true
				return changed, false
			}
			log.Infof("doUpdateVol(%s) name %s: waiting for content tree status %v to be verified",
				status.Key(), status.DisplayName, ctStatus.DisplayName)
			return changed, false
		}
		if ctStatus.State == types.VERIFIED &&
			status.State != types.CREATING_VOLUME &&
			!status.VolumeCreated {

			status.State = types.CREATING_VOLUME
			// first blob is always the root
			if len(ctStatus.Blobs) < 1 {
				log.Errorf("doUpdateVol(%s) name %s: content tree status has no blobs",
					status.Key(), status.DisplayName)
				return changed, false
			}
			sv := SignatureVerifier{
				Signature:        ctStatus.ImageSignature,
				PublicKey:        ctStatus.SignatureKey,
				CertificateChain: ctStatus.CertificateChain,
			}
			rootBlob := lookupOrCreateBlobStatus(ctx, sv, ctStatus.ObjType, ctStatus.Blobs[0])
			if rootBlob == nil {
				log.Errorf("doUpdateVol(%s) name %s: could not find BlobStatus(%s)",
					status.Key(), status.DisplayName, ctStatus.Blobs[0])
				return changed, false
			}
			status.FileLocation = rootBlob.Path
			status.ReferenceName = ctStatus.RelativeURL
			status.ContentFormat = ctStatus.Format
			changed = true
			// Asynch creation; ensure we have requested it
			MaybeAddWorkCreate(ctx, status)
		}
		if status.IsErrorSource(types.ContentTreeStatus{}) {
			log.Infof("doUpdate: Clearing volume error %s", status.Error)
			status.ClearErrorWithSource()
			changed = true
		}
		if status.State == types.CREATING_VOLUME && !status.VolumeCreated {
			vr := lookupVolumeWorkResult(ctx, status.Key())
			if vr != nil {
				log.Infof("doUpdateVol: VolumeWorkResult(%s) location %s, created %t",
					status.Key(), vr.FileLocation, vr.VolumeCreated)
				deleteVolumeWorkResult(ctx, status.Key())
				if status.VolumeCreated != vr.VolumeCreated {
					log.Infof("From vr set VolumeCreated to %s for %s",
						vr.FileLocation, status.VolumeID)
					status.VolumeCreated = vr.VolumeCreated
					changed = true
				}
				if status.FileLocation != vr.FileLocation {
					log.Infof("doUpdateContentTree: From vr set FileLocation to %s for %s",
						vr.FileLocation, status.VolumeID)
					status.FileLocation = vr.FileLocation
					changed = true
				}
				if vr.Error != nil {
					log.Errorf("doUpdate: Error recieved from the volume worker %v",
						vr.Error)
					status.SetErrorWithSource(vr.Error.Error(),
						types.VolumeStatus{}, vr.ErrorTime)
					changed = true
					return changed, false
				} else if status.IsErrorSource(types.VolumeStatus{}) {
					log.Infof("doUpdateContentTree: Clearing volume error %s", status.Error)
					status.ClearErrorWithSource()
					changed = true
				}
			} else {
				log.Infof("doUpdateVol: VolumeWorkResult(%s) not found", status.Key())
			}
		}
		if status.State == types.CREATING_VOLUME && status.VolumeCreated {
			if !status.HasError() {
				status.State = types.CREATED_VOLUME
			}
			changed = true
			// Work is done
			DeleteWorkCreate(ctx, status)
			if status.MaxVolSize == 0 {
				var err error
				log.Infof("doUpdateVol: MaxVolSize is 0 for %s. Filling it up.",
					status.FileLocation)
				_, status.MaxVolSize, err = utils.GetVolumeSize(status.FileLocation)
				if err != nil {
					log.Error(err)
				}
			}
			return changed, true
		}
	default:
		// Unsupported volume content origin type
		errStr := fmt.Sprintf("doUpdateVol(%s) name %s: Volume content origin type %v not supported",
			status.Key(), status.DisplayName, status.VolumeContentOriginType)
		status.SetErrorWithSource(errStr,
			types.VolumeStatus{}, time.Now())
		changed = true
		return changed, false
	}
	return changed, false
}

// updateStatus update all VolumeStatus/ContentTreeStatus which include a blob
// that has this Sha256
func updateStatus(ctx *volumemgrContext, objType, sha string) {

	log.Infof("updateStatus(%s) objType %s", sha, objType)
	found := false
	volPub := ctx.publication(types.OldVolumeStatus{}, objType)
	volItems := volPub.GetAll()
	for _, st := range volItems {
		status := st.(types.OldVolumeStatus)
		var hasSha bool
		for _, blobSha := range status.Blobs {
			if blobSha == sha {
				log.Infof("Found blob %s on VolumeStatus %s", sha, status.Key())
				hasSha = true
			}
		}
		if hasSha {
			found = true
			if changed, _ := doUpdateOld(ctx, &status); changed {
				publishOldVolumeStatus(ctx, &status)
			}
		}
	}
	ctPub := ctx.pubContentTreeStatus
	ctItems := ctPub.GetAll()
	for _, st := range ctItems {
		status := st.(types.ContentTreeStatus)
		var hasSha bool
		for _, blobSha := range status.Blobs {
			if blobSha == sha {
				log.Infof("Found blob %s on ContentTreeStatus %s", sha, status.Key())
				hasSha = true
			}
		}
		if hasSha {
			found = true
			if changed, _ := doUpdateContentTree(ctx, &status); changed {
				log.Infof("updateStatus(%s) publishing ContentTreeStatus", status.Key())
				publishContentTreeStatus(ctx, &status)
			}
			// Volume status referring to this content UUID needs to get updated
			log.Infof("updateStatus(%s) updating volume status from content ID %v", status.Key(), status.ContentID)
			updateVolumeStatusFromContentID(ctx, status.ContentID)
		}
	}
	if !found {
		log.Warnf("XXX updateStatus(%s) objType %s NOT FOUND", sha, objType)
	}
}

func updateContentTreeStatus(ctx *volumemgrContext, contentSha256 string, contentID uuid.UUID) {

	log.Infof("updateContentTreeStatus(%s)", contentID)
	found := false
	pub := ctx.pubContentTreeStatus
	items := pub.GetAll()
	for _, st := range items {
		status := st.(types.ContentTreeStatus)
		if status.ContentSha256 == contentSha256 {
			log.Infof("Found ContentTreeStatus %s", status.Key())
			found = true
			changed, _ := doUpdateContentTree(ctx, &status)
			if changed {
				publishContentTreeStatus(ctx, &status)
			}
			// Volume status referring to this content UUID needs to get updated
			log.Infof("updateStatus(%s) updating volume status from content ID %v", status.Key(), status.ContentID)
			updateVolumeStatusFromContentID(ctx, status.ContentID)
		}
	}
	if !found {
		log.Warnf("XXX updateContentTreeStatus(%s) NOT FOUND", contentID)
	}
}

// Find all the VolumeStatus which refer to this volume uuid
func updateVolumeStatus(ctx *volumemgrContext, volumeID uuid.UUID) {

	log.Infof("updateVolumeStatus for %s", volumeID)
	found := false
	pub := ctx.pubVolumeStatus
	items := pub.GetAll()
	for _, st := range items {
		status := st.(types.VolumeStatus)
		if status.VolumeID == volumeID {
			log.Infof("Found VolumeStatus %s: name %s",
				status.Key(), status.DisplayName)
			found = true
			changed, _ := doUpdateVol(ctx, &status)
			if changed {
				publishVolumeStatus(ctx, &status)
				updateVolumeRefStatus(ctx, &status)
			}
		}
	}
	if !found {
		log.Warnf("XXX updateVolumeStatus(%s) NOT FOUND", volumeID)
	}
}

// Find all the VolumeStatus which refer to this content uuid
func updateVolumeStatusFromContentID(ctx *volumemgrContext, contentID uuid.UUID) {

	log.Infof("updateVolumeStatusFromContentID for %s", contentID)
	found := false
	pub := ctx.pubVolumeStatus
	items := pub.GetAll()
	for _, st := range items {
		status := st.(types.VolumeStatus)
		if status.ContentID == contentID {
			log.Infof("Found VolumeStatus %s: name %s",
				status.Key(), status.DisplayName)
			found = true
			changed, _ := doUpdateVol(ctx, &status)
			if changed {
				publishVolumeStatus(ctx, &status)
				updateVolumeRefStatus(ctx, &status)
			}
		}
	}
	if !found {
		log.Warnf("XXX updateVolumeStatusFromContentID(%s) NOT FOUND", contentID)
	}
}
