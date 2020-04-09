// Copyright (c) 2017-2020 Zededa, Inc.
// SPDX-License-Identifier: Apache-2.0

// Handle publishing the existing-at-boot VolumeStatus
// Published under "unknown" objType with refcount=0. Moved to
// other objType when there is a reference.

package volumemgr

import (
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"

	zconfig "github.com/lf-edge/eve/api/go/config"
	"github.com/lf-edge/eve/pkg/pillar/types"
	uuid "github.com/satori/go.uuid"
	"github.com/shirou/gopsutil/host"
	log "github.com/sirupsen/logrus"
)

// Really a constant
var nilUUID = uuid.UUID{}

// appRwVolumeName - Returns name of the image ( including parent dir )
// Note that we still use the sha in the filename to not impact running images. Otherwise
// we could switch this to imageID
// XXX other types of volumes might want a different name.
func appRwVolumeName(sha256, uuidStr string, purgeCounter uint32, format zconfig.Format,
	origin types.OriginType) string {

	if origin != types.OriginTypeDownload {
		log.Fatalf("XXX unsupported origin %v", origin)
	}
	purgeString := ""
	if purgeCounter != 0 {
		purgeString = fmt.Sprintf("#%d", purgeCounter)
	}
	formatStr := strings.ToLower(format.String())
	return fmt.Sprintf("%s/%s-%s%s.%s", rwImgDirname, sha256,
		uuidStr, purgeString, formatStr)
}

// parseAppRwVolumeName - Returns rwImgDirname, sha256, uuidStr, purgeCounter
func parseAppRwVolumeName(image string) (string, string, string, uint32) {
	// VolumeSha is provided by the controller - it can be uppercase
	// or lowercase.
	re1 := regexp.MustCompile(`(.+)/([0-9A-Fa-f]+)-([0-9a-fA-F\-]+)#([0-9]+)\.(.+)`)
	if re1.MatchString(image) {
		// With purgeCounter
		parsedStrings := re1.FindStringSubmatch(image)
		count, err := strconv.ParseUint(parsedStrings[4], 10, 32)
		if err != nil {
			log.Error(err)
			count = 0
		}
		return parsedStrings[1], parsedStrings[2], parsedStrings[3],
			uint32(count)
	}
	// Without purgeCounter
	re2 := regexp.MustCompile(`(.+)/([0-9A-Fa-f]+)-([0-9a-fA-F\-]+)\.([^\.]+)`)
	if !re2.MatchString(image) {
		log.Errorf("AppRwVolumeName %s doesn't match pattern", image)
		return "", "", "", 0
	}
	parsedStrings := re2.FindStringSubmatch(image)
	return parsedStrings[1], parsedStrings[2], parsedStrings[3], 0
}

// recursive scanning for verified objects,
// to recreate the status files
func populateInitialVolumeStatus(ctx *volumemgrContext, dirName string) {

	log.Infof("populateInitialVolumeStatus(%s)\n", dirName)

	// Record host boot time for comparisons
	hinfo, err := host.Info()
	if err != nil {
		log.Fatalf("host.Info(): %s", err)
	}
	deviceBootTime := time.Unix(int64(hinfo.BootTime), 0).UTC()

	locations, err := ioutil.ReadDir(dirName)
	if err != nil {
		log.Fatal(err)
	}

	for _, location := range locations {
		filelocation := dirName + "/" + location.Name()
		if location.IsDir() {
			log.Debugf("populateInitialVolumeStatus: directory %s ignored\n", filelocation)
			continue
		}

		info, err := os.Stat(filelocation)
		if err != nil {
			log.Errorf("Error in getting file information. Err: %s. "+
				"Deleting file %s", err, filelocation)
			deleteFile(filelocation)
			continue
		}

		size := info.Size()
		_, sha256, appUUIDStr, purgeCounter := parseAppRwVolumeName(filelocation)
		log.Infof("populateInitialVolumeStatus: Processing sha256: %s, AppUuid: %s, "+
			"%d Mbytes, fileLocation:%s",
			sha256, appUUIDStr, size/(1024*1024), filelocation)

		appUUID, err := uuid.FromString(appUUIDStr)
		if err != nil {
			log.Errorf("populateInitialVolumeStatus: Invalid UUIDStr(%s) in "+
				"filename (%s). err: %s. Deleting the File",
				appUUIDStr, filelocation, err)
			deleteFile(filelocation)
			continue
		}

		status := types.VolumeStatus{
			BlobSha256:   sha256,
			AppInstID:    appUUID,
			VolumeID:     nilUUID, // XXX known for other origins?
			PurgeCounter: purgeCounter,
			DisplayName:  "Found in /persist/img",
			FileLocation: filelocation,
			// XXX Is this the correct size? vs. qcow2 size?
			TargetSizeBytes: uint64(size),
			ObjType:         types.UnknownObj,
			VolumeCreated:   true,
			RefCount:        0,
			LastUse:         info.ModTime(),
			PreReboot:       info.ModTime().Before(deviceBootTime),
		}

		publishVolumeStatus(ctx, &status)
	}
}

// Remove from VolumeStatus since fileLocation has been deleted
// XXX implement and call.
func unpublishInitialVolumeStatus(ctx *volumemgrContext, volumeKey string) {

	pub := ctx.publication(types.VolumeStatus{}, types.UnknownObj)
	st, _ := pub.Get(volumeKey)
	if st == nil {
		log.Errorf("unpublishInitialVolumeStatus(%s) key not found",
			volumeKey)
		return
	}
	pub.Unpublish(volumeKey)
}

// XXX for now only handle those with a sha and appInstID
// XXX format arg is not used
func lookupInitVolumeStatus(ctx *volumemgrContext, volumeKey string, originType types.OriginType, format zconfig.Format) *types.VolumeStatus {

	log.Infof("lookupInitVolumeStatus(%s) type %d format %d", volumeKey,
		originType, format)
	// XXX for now
	if originType != types.OriginTypeDownload {
		return nil
	}
	// XXX do we need these check or just look up on volumeKey
	blobSha256, appInstID, volumeID, purgeCounter, err := types.VolumeKeyToParts(volumeKey)
	if err != nil {
		log.Errorf("lookupInitVolumeStatus failed: err %s", err)
		return nil
	}
	// XXX debug
	log.Infof("lookupInitVolumeStatus sha %s appinst %s volume %s purgeCounter %d",
		blobSha256, appInstID, volumeID, purgeCounter)
	if blobSha256 == "" {
		log.Infof("lookupInitVolumeStatus(%s) no sha; not found", volumeKey)
		return nil
	}
	if appInstID == nilUUID {
		log.Infof("lookupInitVolumeStatus(%s) no appInstID; not found", volumeKey)
		return nil
	}
	pub := ctx.publication(types.VolumeStatus{}, types.UnknownObj)
	st, _ := pub.Get(volumeKey)
	if st == nil {
		log.Infof("lookupInitVolumeStatus(%s) key not found", volumeKey)
		return nil
	}
	status := st.(types.VolumeStatus)
	return &status
}

// Periodic garbage collection looking at RefCount=0 files in the unknown
// Others have their delete handler.
func gcObjects(ctx *volumemgrContext, dirName string) {

	log.Debugf("gcObjects()\n")

	pub := ctx.publication(types.VolumeStatus{}, types.UnknownObj)
	items := pub.GetAll()
	for _, st := range items {
		status := st.(types.VolumeStatus)
		if status.RefCount != 0 {
			log.Debugf("gcObjects: skipping RefCount %d: %s\n",
				status.RefCount, status.Key())
			continue
		}
		timePassed := time.Since(status.LastUse)
		timeLimit := time.Duration(ctx.vdiskGCTime) * time.Second
		if timePassed < timeLimit {
			log.Debugf("gcObjects: skipping recently used %s remains %d seconds\n",
				status.Key(), (timePassed-timeLimit)/time.Second)
			continue
		}
		filelocation := status.FileLocation
		if filelocation == "" {
			log.Errorf("No filelocation to remove for %s", status.Key())
		} else {
			log.Infof("gcObjects: removing %s LastUse %v now %v: %s\n",
				filelocation, status.LastUse, time.Now(), status.Key())
			if err := os.Remove(filelocation); err != nil {
				log.Errorln(err)
			}
		}
		unpublishVolumeStatus(ctx, &status)
	}
}

// gc timer just started, reset the LastUse timestamp
func gcResetObjectsLastUse(ctx *volumemgrContext, dirName string) {

	log.Debugf("gcResetObjectsLastUse()\n")

	pub := ctx.publication(types.VolumeStatus{}, types.UnknownObj)
	items := pub.GetAll()
	for _, st := range items {
		status := st.(types.VolumeStatus)
		if status.RefCount == 0 {
			log.Infof("gcResetObjectsLastUse: reset %v LastUse to now\n", status.Key())
			status.LastUse = time.Now()
			publishVolumeStatus(ctx, &status)
		}
	}
}

func deleteFile(filelocation string) {
	if err := os.Remove(filelocation); err != nil {
		log.Errorf("Failed to delete file %s. Error: %s",
			filelocation, err.Error())
	}
}
