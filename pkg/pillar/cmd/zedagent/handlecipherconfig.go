// Copyright (c) 2017-2018 Zededa, Inc.
// SPDX-License-Identifier: Apache-2.0

package zedagent

// cipher specific parser/utility routines

import (
	"bytes"
	"crypto/sha256"
	"fmt"

	zconfig "github.com/lf-edge/eve/api/go/config"
	"github.com/lf-edge/eve/pkg/pillar/types"
	log "github.com/sirupsen/logrus"
)

var cipherCtxHash []byte

// cipher context parsing routine
func parseCipherContext(ctx *cipherContext, config *zconfig.EdgeDevConfig) {

	log.Infof("Started parsing cipher context")
	cfgCipherContextList := config.GetCipherContexts()
	h := sha256.New()
	for _, cfgContextConfig := range cfgCipherContextList {
		computeConfigElementSha(h, cfgContextConfig)
	}
	newHash := h.Sum(nil)
	if bytes.Equal(newHash, cipherCtxHash) {
		return
	}
	log.Infof("parseCipherContext: Applying updated config "+
		"Last Sha: % x, "+
		"New  Sha: % x, "+
		"Num of cfgContextConfig: %d",
		cipherCtxHash, newHash, len(cfgCipherContextList))

	cipherCtxHash = newHash

	// First look for deleted ones
	items := ctx.pubCipherContextConfig.GetAll()
	for idStr := range items {
		found := false
		for _, cfgContextConfig := range cfgCipherContextList {
			if cfgContextConfig.GetContextId() == idStr {
				found = true
				break
			}
		}
		// cipherContext not found, delete
		if !found {
			log.Infof("parseCipherContext: deleting %s", idStr)
			unpublishCipherContextConfig(ctx, idStr)
		}
	}

	for _, cfgContextConfig := range cfgCipherContextList {
		if cfgContextConfig.GetContextId() == "" {
			log.Debugf("parseCipherContext ignoring empty")
			continue
		}
		context := types.CipherContextConfig{
			ContextID:          cfgContextConfig.GetContextId(),
			HashScheme:         cfgContextConfig.GetHashScheme(),
			KeyExchangeScheme:  cfgContextConfig.GetKeyExchangeScheme(),
			EncryptionScheme:   cfgContextConfig.GetEncryptionScheme(),
			DeviceCertHash:     cfgContextConfig.GetDeviceCertHash(),
			ControllerCertHash: cfgContextConfig.GetControllerCertHash(),
		}
		publishCipherContextConfig(ctx, context)
	}
	log.Infof("parsing cipher context done")
}

// parseCipherBlock : will collate all the relevant information
// ciphercontext will be used to get the certs and encryption schemes
func parseCipherBlock(ctx *getconfigContext, key string,
	cfgCipherBlock *zconfig.CipherBlock) types.CipherBlockStatus {

	log.Infof("parseCipherBlock(%s) started", key)
	if cfgCipherBlock == nil {
		log.Infof("parseCipherBlock(%s) nil cipher block", key)
		return types.CipherBlockStatus{CipherBlockID: key}
	}
	cipherBlock := types.CipherBlockStatus{
		CipherBlockID:   key,
		CipherContextID: cfgCipherBlock.GetCipherContextId(),
		InitialValue:    cfgCipherBlock.GetInitialValue(),
		CipherData:      cfgCipherBlock.GetCipherData(),
		ClearTextHash:   cfgCipherBlock.GetClearTextSha256(),
	}

	// should contain valid cipher data
	if len(cipherBlock.CipherData) == 0 ||
		len(cipherBlock.CipherContextID) == 0 {
		errStr := fmt.Sprintf("%s, block contains incomplete data", key)
		log.Errorf(errStr)
		cipherBlock.SetErrorNow(errStr)
		return cipherBlock
	}
	log.Infof("%s, marking cipher as true", key)
	cipherBlock.IsCipher = true

	log.Infof("parseCipherBlock(%s) done", key)
	return cipherBlock
}

// orchestration routines
// on cipher context, controller certificate, eve node certificate change

func handleControllerCertConfigModify(ctx *zedagentContext,
	config types.ZCertConfig) {

	// generate controller certificate status
	status := types.ZCertStatus{
		ZCertConfig: config,
	}
	// TBD:XXX generate signing verification status
	publishControllerCertStatus(ctx.cipherCtx, status)

	pub := ctx.cipherCtx.pubCipherContextConfig
	items := pub.GetAll()
	for _, item := range items {
		context := item.(types.CipherContextConfig)
		if bytes.Equal(config.Hash, context.ControllerCertHash) {
			deviceKey := context.EveNodeCertKey()
			dcertstatus := lookupEveNodeCertStatus(ctx.cipherCtx, deviceKey)
			updateCipherContextStatus(ctx, context, &status, dcertstatus)
		}
	}
}

func handleControllerCertConfigDelete(ctx *zedagentContext, key string) {
	config := lookupControllerCertStatus(ctx.cipherCtx, key)
	// update related cipher contexts
	pub := ctx.cipherCtx.pubCipherContextConfig
	items := pub.GetAll()
	for _, item := range items {
		context := item.(types.CipherContextConfig)
		if bytes.Equal(config.Hash, context.ControllerCertHash) {
			deviceKey := context.EveNodeCertKey()
			dcertstatus := lookupEveNodeCertStatus(ctx.cipherCtx, deviceKey)
			updateCipherContextStatus(ctx, context, nil, dcertstatus)
		}
	}
	unpublishControllerCertStatus(ctx.cipherCtx, key)
}

func handleCipherContextConfigModify(ctx *zedagentContext,
	config types.CipherContextConfig) {
	deviceKey := config.EveNodeCertKey()
	controllerKey := config.ControllerCertKey()
	ccertstatus := lookupControllerCertStatus(ctx.cipherCtx, controllerKey)
	dcertstatus := lookupEveNodeCertStatus(ctx.cipherCtx, deviceKey)
	updateCipherContextStatus(ctx, config, ccertstatus, dcertstatus)
}

func updateCipherContextStatus(ctx *zedagentContext, config types.CipherContextConfig,
	certStatus *types.ZCertStatus, dcert *types.ZCertStatus) {
	var errStr, errStr0, errStr1 string

	// fill up the structure, with config values
	status := types.CipherContextStatus{
		CipherContextConfig: config,
	}

	// update the error string, if any
	if certStatus != nil {
		status.ControllerCert = certStatus.Cert
		if len(certStatus.Error) != 0 {
			errStr0 = "controllerCert: " + certStatus.Error
		}
	}

	// ciper context do not need device cert
	if dcert != nil {
		if len(dcert.Error) != 0 {
			errStr1 = "deviceCert: " + dcert.Error
		}
	}

	if len(errStr0) != 0 || len(errStr1) != 0 {
		if len(errStr0) != 0 && len(errStr1) != 0 {
			errStr = errStr0 + "," + errStr1
		} else {
			if len(errStr0) != 0 {
				errStr = errStr0
			}
			if len(errStr1) != 0 {
				errStr = errStr1
			}
		}
		status.ErrorAndTime.SetErrorNow(errStr)
	}
	publishCipherContextStatus(ctx.cipherCtx, status)
}

// trigger handler routines
// eve node cert config change will trigger status and
// related cipher context status update
func handleEveNodeCertConfigModify(ctxArg interface{}, key string,
	configArg interface{}) {
	ctx := ctxArg.(*zedagentContext)
	config := configArg.(types.ZCertConfig)
	// generate eve node certificate status
	nodeCertStatus := types.ZCertStatus{
		ZCertConfig: config,
	}

	// TBD:XXX generate signing verification status
	publishEveNodeCertStatus(ctx.cipherCtx, nodeCertStatus)

	// update related cipher contexts
	pub := ctx.cipherCtx.pubCipherContextConfig
	items := pub.GetAll()
	for _, item := range items {
		context := item.(types.CipherContextConfig)
		if bytes.Equal(config.Hash, context.DeviceCertHash) {
			cKey := context.ControllerCertKey()
			certStatus := lookupControllerCertStatus(ctx.cipherCtx, cKey)
			updateCipherContextStatus(ctx, context, certStatus, &nodeCertStatus)
		}
	}

	// trigger eve node certificate push to controller
	triggerEveNodeCertsEvent(ctx)
}

func handleEveNodeCertConfigDelete(ctxArg interface{}, key string,
	configArg interface{}) {
	ctx := ctxArg.(*zedagentContext)
	config := configArg.(types.ZCertConfig)
	// update related cipher contexts
	pub := ctx.cipherCtx.pubCipherContextConfig
	items := pub.GetAll()
	for _, item := range items {
		context := item.(types.CipherContextConfig)
		if bytes.Equal(config.Hash, context.DeviceCertHash) {
			cKey := context.ControllerCertKey()
			ccertstatus := lookupControllerCertStatus(ctx.cipherCtx, cKey)
			updateCipherContextStatus(ctx, context, ccertstatus, nil)
		}
	}
	unpublishEveNodeCertStatus(ctx.cipherCtx, key)

	// trigger eve node certificate push to controller
	triggerEveNodeCertsEvent(ctx)
}
