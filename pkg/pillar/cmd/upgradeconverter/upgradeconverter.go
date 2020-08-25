// Copyright (c) 2017-2018 Zededa, Inc.
// SPDX-License-Identifier: Apache-2.0

package upgradeconverter

import (
	"flag"

	"github.com/lf-edge/eve/pkg/pillar/agentlog"
	"github.com/lf-edge/eve/pkg/pillar/base"
	"github.com/lf-edge/eve/pkg/pillar/pidfile"
	"github.com/lf-edge/eve/pkg/pillar/pubsub"
	"github.com/lf-edge/eve/pkg/pillar/types"
	"github.com/sirupsen/logrus"
)

var conversionHandlers = []ConversionHandler{
	{
		description: "Convert Global Settings to new format",
		handlerFunc: convertGlobalConfig,
	},
	{
		description: "Move volumes to /persist/vault",
		handlerFunc: convertPersistVolumes,
	},
	{
		description: "Move verified files to /persist/vault/verifier/verified",
		handlerFunc: renameVerifiedFiles,
	},
}

type ucContext struct {
	agentName     string
	debugOverride bool
	noFlag        bool

	// FilePaths. These are defined here instead of consts for easier unit tests
	persistDir       string
	persistConfigDir string
	varTmpDir        string
	ps               *pubsub.PubSub
}

func (ctx ucContext) configItemValueMapDir() string {
	return ctx.persistConfigDir + "/ConfigItemValueMap/"
}
func (ctx ucContext) configItemValueMapFile() string {
	return ctx.configItemValueMapDir() + "/global.json"
}
func (ctx ucContext) globalConfigDir() string {
	return ctx.persistConfigDir + "/GlobalConfig"
}
func (ctx ucContext) globalConfigFile() string {
	return ctx.globalConfigDir() + "/global.json"
}

// Old location for volumes
func (ctx ucContext) imgDir() string {
	return ctx.persistDir + "/img/"
}

// Old location for volumes
func (ctx ucContext) preparedDir() string {
	return ctx.persistDir + "/runx/pods/prepared/"
}

// New location for volumes
func (ctx ucContext) volumesDir() string {
	return ctx.persistDir + "/vault/volumes/"
}

// checkpoint file for EdgeDevConfig
func (ctx ucContext) configCheckpointFile() string {
	return ctx.persistDir + "/checkpoint/lastconfig"
}

func runHandlers(ctxPtr *ucContext) {
	for _, handler := range conversionHandlers {
		log.Infof("upgradeconverter.Run: Running Conversion handler: %s",
			handler.description)
		err := handler.handlerFunc(ctxPtr)
		if err != nil {
			log.Errorf("upgradeconverter.Run: Handler %s failed. err %s\n ctx:%+v",
				handler.description, err, *ctxPtr)
		}
	}
}

var log *base.LogObject

// Run - runs the main upgradeconverter process
func Run(ps *pubsub.PubSub) int {
	ctx := &ucContext{agentName: "upgradeconverter",
		persistDir:       types.PersistDir,
		persistConfigDir: types.PersistConfigDir,
		varTmpDir:        "/var/tmp",
		ps:               ps,
	}
	debugPtr := flag.Bool("d", false, "Debug flag")
	persistPtr := flag.String("p", "/persist", "persist directory")
	noFlagPtr := flag.Bool("n", false, "Don't do anything just log flag")
	flag.Parse()
	ctx.debugOverride = *debugPtr
	ctx.persistDir = *persistPtr // XXX remove? Or use for tests?
	ctx.noFlag = *noFlagPtr
	if ctx.debugOverride {
		logrus.SetLevel(logrus.DebugLevel)
	} else {
		logrus.SetLevel(logrus.InfoLevel)
	}

	// XXX Make logrus record a noticable global source
	agentlog.Init("xyzzy-" + ctx.agentName)

	log = agentlog.Init(ctx.agentName)
	if err := pidfile.CheckAndCreatePidfile(log, ctx.agentName); err != nil {
		log.Fatal(err)
	}
	log.Infof("Starting %s\n", ctx.agentName)
	runHandlers(ctx)
	return 0
}

// HandlerFunc - defines functions to handle each conversion
type HandlerFunc func(ctx *ucContext) error

// ConversionHandler - defines type for processing
type ConversionHandler struct {
	description string
	handlerFunc HandlerFunc
}
