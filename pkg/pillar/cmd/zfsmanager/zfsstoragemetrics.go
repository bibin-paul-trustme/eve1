// Copyright (c) 2022 Zededa, Inc.
// SPDX-License-Identifier: Apache-2.0

package zfsmanager

import (
	"time"

	libzfs "github.com/bicomsystems/go-libzfs"
	"github.com/lf-edge/eve/pkg/pillar/types"
	"github.com/lf-edge/eve/pkg/pillar/vault"
	"github.com/lf-edge/eve/pkg/pillar/zfs"
)

const (
	storageMetricsPublishInterval = 5 * time.Second // interval between publishing metrics from zfs
)

func storageMetricsPublisher(ctxPtr *zfsContext) {
	if vault.ReadPersistType() != types.PersistZFS {
		return
	}
	collectAndPublishStorageMetrics(ctxPtr)

	t := time.NewTicker(storageMetricsPublishInterval)
	for {
		select {
		case <-t.C:
			collectAndPublishStorageMetrics(ctxPtr)
		}
	}
}

func collectAndPublishStorageMetrics(ctxPtr *zfsContext) {
	zpoolList, err := libzfs.PoolOpenAll()
	if err != nil {
		log.Errorf("get zpool list for collect metrics failed %v", err)
	} else {
		for _, zpool := range zpoolList {
			defer zpool.Close()
			vdevs, err := zpool.VDevTree()
			if err != nil {
				log.Errorf("get vdev tree for collect metrics failed %v", err)
				continue
			}

			zfsPoolMetrics := zfs.GetAllDeviceMetricsFromZpool(vdevs)
			if err := ctxPtr.storageMetricsPub.Publish(zfsPoolMetrics.Key(), *zfsPoolMetrics); err != nil {
				log.Errorf("error in publishing of storageMetrics: %s", err)
			}
		}
	}
}
