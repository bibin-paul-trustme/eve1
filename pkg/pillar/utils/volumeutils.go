// Copyright (c) 2020 Zededa, Inc.
// SPDX-License-Identifier: Apache-2.0

package utils

import (
	"errors"
	"fmt"
	"github.com/lf-edge/eve/pkg/pillar/base"
	"os"
	"path/filepath"

	"github.com/lf-edge/eve/pkg/pillar/diskmetrics"
)

// dirSize returns the size of the directory
func dirSize(path string) (uint64, error) {
	var size int64
	err := filepath.Walk(path, func(_ string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() {
			size += info.Size()
		}
		return err
	})
	return uint64(size), err
}

// GetVolumeSize returns the actual and maximum size of the volume
// plus a DiskType and a DirtyFlag
func GetVolumeSize(log *base.LogObject, name string) (uint64, uint64, string, bool, error) {
	info, err := os.Stat(name)
	if err != nil {
		errStr := fmt.Sprintf("GetVolumeMaxSize failed for %s: %v",
			name, err)
		return 0, 0, "", false, errors.New(errStr)
	}
	if info.IsDir() {
		// Assume this is a container
		size, err := dirSize(name)
		if err != nil {
			errStr := fmt.Sprintf("GetVolumeMaxSize failed for %s: %v",
				name, err)
			return 0, 0, "", false, errors.New(errStr)
		}
		return size, size, "CONTAINER", false, nil
	}
	imgInfo, err := diskmetrics.GetImgInfo(log, name)
	if err != nil {
		errStr := fmt.Sprintf("GetVolumeMaxSize/GetImgInfo failed for %s: %v",
			name, err)
		return 0, 0, "", false, errors.New(errStr)
	}
	return imgInfo.ActualSize, imgInfo.VirtualSize, imgInfo.Format,
		imgInfo.DirtyFlag, nil
}
