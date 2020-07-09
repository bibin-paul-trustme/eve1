// Copyright (c) 2020 Zededa, Inc.
// SPDX-License-Identifier: Apache-2.0

package volumemgr

// Interface to worker to run the create and destroy in separate goroutines

import (
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	zconfig "github.com/lf-edge/eve/api/go/config"
	"github.com/lf-edge/eve/pkg/pillar/pubsub"
	"github.com/lf-edge/eve/pkg/pillar/types"
	uuid "github.com/satori/go.uuid"
)

func TestHandleWorkCreate(t *testing.T) {
	status := types.OldVolumeStatus{
		VolumeID:     uuid.NewV4(), // XXX future
		AppInstID:    uuid.NewV4(),
		PurgeCounter: 0,
		Origin:       types.OriginTypeDownload,
		Format:       zconfig.Format_QCOW2,
		ObjType:      types.AppImgObj,
	}
	testMatrix := map[string]struct {
		ReadOnly         bool
		SrcLocation      string
		BlobSha256       string
		ExpectFail       bool
		ExpectedLocation string
		ExpectCreated    bool
	}{
		"read-only": {
			ReadOnly:         true,
			SrcLocation:      "/dev/null",
			BlobSha256:       "somesha",
			ExpectFail:       false,
			ExpectedLocation: "/dev/null",
			ExpectCreated:    true,
		},
		"read-write fail": {
			ReadOnly:    false,
			SrcLocation: "/dev/null",
			BlobSha256:  "somesha",
			ExpectFail:  true,
			ExpectedLocation: appRwOldVolumeName("somesha", status.AppInstID.String(),
				status.PurgeCounter, status.Format, status.Origin, false),

			ExpectCreated: true,
		},
	}
	ctx := initCtx(t)
	ctx.workerOld = InitHandleWorkOld(&ctx)
	for testname, test := range testMatrix {
		t.Logf("Running test case %s", testname)
		t.Run(testname, func(t *testing.T) {

			status.ReadOnly = test.ReadOnly
			status.FileLocation = test.SrcLocation
			status.BlobSha256 = test.BlobSha256
			status.Blobs = []string{
				test.BlobSha256,
			}
			ctx.pubBlobStatus.Publish(test.BlobSha256, types.BlobStatus{
				State:  types.VERIFIED,
				Path:   test.SrcLocation,
				Sha256: strings.ToLower(test.BlobSha256),
			})

			MaybeAddWorkCreateOld(&ctx, &status)
			assert.Equal(t, ctx.workerOld.NumPending(), 1)
			MaybeAddWorkCreateOld(&ctx, &status)
			assert.Equal(t, ctx.workerOld.NumPending(), 1)

			vr := lookupVolumeWorkResult(&ctx, status.Key())
			assert.Nil(t, vr)

			res := <-ctx.workerOld.MsgChan()
			HandleWorkResultOld(&ctx, ctx.workerOld.Process(res))
			assert.Equal(t, ctx.workerOld.NumPending(), 0)
			vr = lookupVolumeWorkResult(&ctx, status.Key())
			assert.NotNil(t, vr)
			deleteVolumeWorkResult(&ctx, status.Key())
			DeleteWorkCreateOld(&ctx, &status)

			if test.ExpectFail {
				assert.NotNil(t, vr.Error, "Error")
				assert.NotEqual(t, vr.ErrorTime, time.Time{},
					"ErrorTime")
			} else {
				assert.Nil(t, vr.Error, "Error")
				assert.Equal(t, vr.ErrorTime, time.Time{},
					"ErrorTime")
			}
			assert.Equal(t, vr.VolumeCreated, test.ExpectCreated,
				"VolumeCreated")
			assert.Equal(t, vr.FileLocation, test.ExpectedLocation,
				"FileLocation")
		})
	}
}

func TestHandleWorkDestroy(t *testing.T) {
	testMatrix := map[string]struct {
		ReadOnly         bool
		SrcLocation      string
		BlobSha256       string
		VolumeCreated    bool
		ExpectFail       bool
		ExpectedLocation string
		ExpectCreated    bool
	}{
		"read-only created": {
			ReadOnly:         true,
			SrcLocation:      "/tmp/xyzzy",
			BlobSha256:       "somesha",
			VolumeCreated:    true,
			ExpectFail:       false,
			ExpectedLocation: "",
			ExpectCreated:    false,
		},
		"read-write created": {
			ReadOnly:         false,
			SrcLocation:      "/tmp/xyzzy",
			BlobSha256:       "somesha",
			VolumeCreated:    true,
			ExpectFail:       true,
			ExpectedLocation: "",
			ExpectCreated:    true,
		},
		"read-write not created": {
			ReadOnly:         false,
			SrcLocation:      "/dev/null",
			BlobSha256:       "somesha",
			VolumeCreated:    false,
			ExpectFail:       false,
			ExpectedLocation: "/dev/null",
			ExpectCreated:    false,
		},
	}
	ctx := initCtx(t)
	ctx.workerOld = InitHandleWorkOld(&ctx)
	status := types.OldVolumeStatus{
		VolumeID:     uuid.NewV4(), // XXX future
		AppInstID:    uuid.NewV4(),
		PurgeCounter: 0,
		Origin:       types.OriginTypeDownload,
		Format:       zconfig.Format_QCOW2,
		ObjType:      types.AppImgObj,
	}
	for testname, test := range testMatrix {
		t.Logf("Running test case %s", testname)
		t.Run(testname, func(t *testing.T) {

			status.ReadOnly = test.ReadOnly
			status.FileLocation = test.SrcLocation
			status.BlobSha256 = test.BlobSha256
			status.VolumeCreated = test.VolumeCreated
			MaybeAddWorkDestroyOld(&ctx, &status)
			assert.Equal(t, ctx.workerOld.NumPending(), 1)
			MaybeAddWorkDestroyOld(&ctx, &status)
			assert.Equal(t, ctx.workerOld.NumPending(), 1)

			vr := lookupVolumeWorkResult(&ctx, status.Key())
			assert.Nil(t, vr)

			res := <-ctx.workerOld.MsgChan()
			HandleWorkResultOld(&ctx, ctx.workerOld.Process(res))
			assert.Equal(t, ctx.workerOld.NumPending(), 0)
			vr = lookupVolumeWorkResult(&ctx, status.Key())
			assert.NotNil(t, vr)
			deleteVolumeWorkResult(&ctx, status.Key())
			DeleteWorkDestroyOld(&ctx, &status)

			if test.ExpectFail {
				assert.NotNil(t, vr.Error, "Error")
				assert.NotEqual(t, vr.ErrorTime, time.Time{},
					"ErrorTime")
			} else {
				assert.Nil(t, vr.Error, "Error")
				assert.Equal(t, vr.ErrorTime, time.Time{},
					"ErrorTime")
			}
			assert.Equal(t, vr.VolumeCreated, test.ExpectCreated,
				"VolumeCreated")
			assert.Equal(t, vr.FileLocation, test.ExpectedLocation,
				"FileLocation")
		})
	}
}

func initCtx(t *testing.T) volumemgrContext {
	ctx := volumemgrContext{}
	ps := pubsub.New(&pubsub.EmptyDriver{})

	pubAppVolumeStatus, err := ps.NewPublication(pubsub.PublicationOptions{
		AgentName:  agentName,
		AgentScope: types.AppImgObj,
		TopicType:  types.OldVolumeStatus{},
	})
	assert.Nil(t, err)
	ctx.pubAppVolumeStatus = pubAppVolumeStatus

	pubContentTreeStatus, err := ps.NewPublication(pubsub.PublicationOptions{
		AgentName:  agentName,
		AgentScope: types.AppImgObj,
		TopicType:  types.ContentTreeStatus{},
	})
	assert.Nil(t, err)
	ctx.pubContentTreeStatus = pubContentTreeStatus

	pubBaseOsContentTreeStatus, err := ps.NewPublication(pubsub.PublicationOptions{
		AgentName:  agentName,
		AgentScope: types.BaseOsObj,
		TopicType:  types.ContentTreeStatus{},
	})
	assert.Nil(t, err)
	ctx.pubBaseOsContentTreeStatus = pubBaseOsContentTreeStatus

	pubBlobStatus, err := ps.NewPublication(pubsub.PublicationOptions{
		AgentName:  agentName,
		AgentScope: types.AppImgObj,
		TopicType:  types.BlobStatus{},
	})
	assert.Nil(t, err)
	ctx.pubBlobStatus = pubBlobStatus
	return ctx
}
