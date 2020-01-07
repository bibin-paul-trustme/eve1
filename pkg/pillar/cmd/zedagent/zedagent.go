// Copyright (c) 2017-2018 Zededa, Inc.
// SPDX-License-Identifier: Apache-2.0

// zedAgent interfaces with zedcloud for
//   * config sync
//   * metric/info publish
// app instance config is published to zedmanager for orchestration
// baseos/certs config is published to baseosmgr for orchestration
// datastore config is published for downloader consideration
// event based baseos/app instance/device info published to ZedCloud
// periodic status/metric published to zedCloud

// zedagent handles the following configuration
//   * app instance config/status  <zedagent>   / <appimg> / <config | status>
//   * base os config/status       <zedagent>   / <baseos> / <config | status>
//   * certs config/status         <zedagent>   / certs>   / <config | status>
// <base os>
//   <zedagent>  <baseos> <config> --> <baseosmgr>  <baseos> <status>
// <certs>
//   <zedagent>  <certs> <config> --> <baseosmgr>   <certs> <status>
// <app image>
//   <zedagent>  <appimage> <config> --> <zedmanager> <appimage> <status>
// <datastore>
//   <zedagent>  <datastore> <config> --> <downloader>

package zedagent

import (
	"container/list"
	"flag"
	"fmt"
	"os"
	"time"

	"github.com/google/go-cmp/cmp"
	"github.com/lf-edge/eve/pkg/pillar/agentlog"
	"github.com/lf-edge/eve/pkg/pillar/cast"
	"github.com/lf-edge/eve/pkg/pillar/pidfile"
	"github.com/lf-edge/eve/pkg/pillar/pubsub"
	"github.com/lf-edge/eve/pkg/pillar/types"
	"github.com/lf-edge/eve/pkg/pillar/utils"
	"github.com/lf-edge/eve/pkg/pillar/zedcloud"

	log "github.com/sirupsen/logrus"
)

const (
	agentName          = "zedagent"
	restartCounterFile = types.IdentityDirname + "/restartcounter"
	firstbootFile      = types.TmpDirname + "/first-boot"
	// checkpointDirname - location of config checkpoint
	checkpointDirname = types.PersistDir + "/checkpoint"
	// Time limits for event loop handlers
	errorTime   = 3 * time.Minute
	warningTime = 40 * time.Second
)

// Set from Makefile
var Version = "No version specified"

// XXX move to a context? Which? Used in handleconfig and handlemetrics!
var deviceNetworkStatus *types.DeviceNetworkStatus = &types.DeviceNetworkStatus{}

// XXX globals filled in by subscription handlers and read by handlemetrics
// XXX could alternatively access sub object when adding them.
var clientMetrics interface{}
var logmanagerMetrics interface{}
var downloaderMetrics interface{}
var networkMetrics types.NetworkMetrics

// Context for handleDNSModify
type DNSContext struct {
	usableAddressCount     int
	DNSinitialized         bool // Received DeviceNetworkStatus
	subDeviceNetworkStatus *pubsub.Subscription
	triggerGetConfig       bool
	triggerDeviceInfo      bool
}

type zedagentContext struct {
	verifierRestarted         bool              // Information from handleVerifierRestarted
	getconfigCtx              *getconfigContext // Cross link
	assignableAdapters        *types.AssignableAdapters
	subAssignableAdapters     *pubsub.Subscription
	iteration                 int
	subNetworkInstanceStatus  *pubsub.Subscription
	subCertObjConfig          *pubsub.Subscription
	TriggerDeviceInfo         chan<- struct{}
	zbootRestarted            bool // published by baseosmgr
	subBaseOsStatus           *pubsub.Subscription
	subBaseOsDownloadStatus   *pubsub.Subscription
	subCertObjDownloadStatus  *pubsub.Subscription
	subBaseOsVerifierStatus   *pubsub.Subscription
	subAppImgDownloadStatus   *pubsub.Subscription
	subAppImgVerifierStatus   *pubsub.Subscription
	subNetworkInstanceMetrics *pubsub.Subscription
	subAppFlowMonitor         *pubsub.Subscription
	subAppVifIPTrig           *pubsub.Subscription
	pubGlobalConfig           *pubsub.Publication
	subGlobalConfig           *pubsub.Subscription
	subVaultStatus            *pubsub.Subscription
	GCInitialized             bool // Received initial GlobalConfig
	subZbootStatus            *pubsub.Subscription
	rebootCmd                 bool
	rebootCmdDeferred         bool
	deviceReboot              bool
	rebootReason              string
	rebootStack               string
	rebootTime                time.Time
	restartCounter            uint32
	subDevicePortConfigList   *pubsub.Subscription
	devicePortConfigList      types.DevicePortConfigList
	remainingTestTime         time.Duration
	physicalIoAdapterMap      map[string]types.PhysicalIOAdapter
	globalConfig              types.GlobalConfig
	globalStatus              types.GlobalStatus
}

var debug = false
var debugOverride bool // From command line arg
var flowQ *list.List

func Run() {
	versionPtr := flag.Bool("v", false, "Version")
	debugPtr := flag.Bool("d", false, "Debug flag")
	curpartPtr := flag.String("c", "", "Current partition")
	parsePtr := flag.String("p", "", "parse checkpoint file")
	validatePtr := flag.Bool("V", false, "validate UTF-8 in checkpoint")
	flag.Parse()
	debug = *debugPtr
	debugOverride = debug
	if debugOverride {
		log.SetLevel(log.DebugLevel)
	} else {
		log.SetLevel(log.InfoLevel)
	}
	curpart := *curpartPtr
	parse := *parsePtr
	validate := *validatePtr
	if *versionPtr {
		fmt.Printf("%s: %s\n", os.Args[0], Version)
		return
	}
	if validate && parse == "" {
		fmt.Printf("Setting -V requires -p\n")
		os.Exit(1)
	}
	if parse != "" {
		res, config := readValidateConfig(
			types.GlobalConfigDefaults.StaleConfigTime, parse)
		if !res {
			fmt.Printf("Failed to parse %s\n", parse)
			os.Exit(1)
		}
		fmt.Printf("parsed proto <%v>\n", config)
		if validate {
			valid := validateConfigUTF8(config)
			if !valid {
				fmt.Printf("Found some invalid UTF-8\n")
				os.Exit(1)
			}
		}
		return
	}
	logf, err := agentlog.Init(agentName, curpart)
	if err != nil {
		log.Fatal(err)
	}
	defer logf.Close()
	if err := pidfile.CheckAndCreatePidfile(agentName); err != nil {
		log.Fatal(err)
	}

	log.Infof("Starting %s\n", agentName)

	triggerDeviceInfo := make(chan struct{}, 1)
	zedagentCtx := zedagentContext{TriggerDeviceInfo: triggerDeviceInfo}
	zedagentCtx.globalConfig = types.GlobalConfigDefaults
	zedagentCtx.globalStatus.ConfigItems = make(
		map[string]types.ConfigItemStatus)
	zedagentCtx.globalStatus.UpdateItemValuesFromGlobalConfig(
		zedagentCtx.globalConfig)
	zedagentCtx.globalStatus.UnknownConfigItems = make(
		map[string]types.ConfigItemStatus)

	zedagentCtx.physicalIoAdapterMap = make(map[string]types.PhysicalIOAdapter)

	zedagentCtx.pubGlobalConfig, err = pubsub.PublishPersistent("",
		types.GlobalConfig{})
	if err != nil {
		log.Fatal(err)
	}

	// Run a periodic timer so we always update StillRunning
	stillRunning := time.NewTicker(25 * time.Second)
	agentlog.StillRunning(agentName, warningTime, errorTime)
	agentlog.StillRunning(agentName+"config", warningTime, errorTime)
	agentlog.StillRunning(agentName+"metrics", warningTime, errorTime)
	agentlog.StillRunning(agentName+"devinfo", warningTime, errorTime)

	// Tell ourselves to go ahead
	// initialize the module specifig stuff
	handleInit(zedagentCtx.globalConfig.NetworkSendTimeout)

	// Context to pass around
	getconfigCtx := getconfigContext{}

	// Pick up (mostly static) AssignableAdapters before we report
	// any device info
	aa := types.AssignableAdapters{}
	zedagentCtx.assignableAdapters = &aa

	// Cross link
	getconfigCtx.zedagentCtx = &zedagentCtx
	zedagentCtx.getconfigCtx = &getconfigCtx

	// Timer for deferred sends of info messages
	deferredChan := zedcloud.InitDeferred()

	// Make sure we have a GlobalConfig file with defaults
	utils.ReadAndUpdateGCFile(zedagentCtx.pubGlobalConfig)

	subAssignableAdapters, err := pubsub.Subscribe("domainmgr",
		types.AssignableAdapters{}, false, &zedagentCtx)
	if err != nil {
		log.Fatal(err)
	}
	subAssignableAdapters.MaxProcessTimeWarn = warningTime
	subAssignableAdapters.MaxProcessTimeError = errorTime
	subAssignableAdapters.ModifyHandler = handleAAModify
	subAssignableAdapters.CreateHandler = handleAAModify
	subAssignableAdapters.DeleteHandler = handleAADelete
	zedagentCtx.subAssignableAdapters = subAssignableAdapters
	subAssignableAdapters.Activate()

	pubPhysicalIOAdapters, err := pubsub.Publish(agentName,
		types.PhysicalIOAdapterList{})
	if err != nil {
		log.Fatal(err)
	}
	pubPhysicalIOAdapters.ClearRestarted()
	getconfigCtx.pubPhysicalIOAdapters = pubPhysicalIOAdapters

	pubDevicePortConfig, err := pubsub.Publish(agentName,
		types.DevicePortConfig{})
	if err != nil {
		log.Fatal(err)
	}
	getconfigCtx.pubDevicePortConfig = pubDevicePortConfig

	// Publish NetworkXObjectConfig and for outselves. XXX remove
	pubNetworkXObjectConfig, err := pubsub.Publish(agentName,
		types.NetworkXObjectConfig{})
	if err != nil {
		log.Fatal(err)
	}
	getconfigCtx.pubNetworkXObjectConfig = pubNetworkXObjectConfig

	pubNetworkInstanceConfig, err := pubsub.Publish(agentName,
		types.NetworkInstanceConfig{})
	if err != nil {
		log.Fatal(err)
	}
	getconfigCtx.pubNetworkInstanceConfig = pubNetworkInstanceConfig

	pubAppInstanceConfig, err := pubsub.Publish(agentName,
		types.AppInstanceConfig{})
	if err != nil {
		log.Fatal(err)
	}
	getconfigCtx.pubAppInstanceConfig = pubAppInstanceConfig
	pubAppInstanceConfig.ClearRestarted()

	pubAppNetworkConfig, err := pubsub.Publish(agentName,
		types.AppNetworkConfig{})
	if err != nil {
		log.Fatal(err)
	}
	pubAppNetworkConfig.ClearRestarted()
	getconfigCtx.pubAppNetworkConfig = pubAppNetworkConfig

	// XXX defer this until we have some config from cloud or saved copy
	pubAppInstanceConfig.SignalRestarted()

	pubCertObjConfig, err := pubsub.Publish(agentName,
		types.CertObjConfig{})
	if err != nil {
		log.Fatal(err)
	}
	pubCertObjConfig.ClearRestarted()
	getconfigCtx.pubCertObjConfig = pubCertObjConfig

	pubBaseOsConfig, err := pubsub.Publish(agentName,
		types.BaseOsConfig{})
	if err != nil {
		log.Fatal(err)
	}
	pubBaseOsConfig.ClearRestarted()
	getconfigCtx.pubBaseOsConfig = pubBaseOsConfig

	pubZedAgentStatus, err := pubsub.Publish(agentName,
		types.ZedAgentStatus{})
	if err != nil {
		log.Fatal(err)
	}
	pubZedAgentStatus.ClearRestarted()
	getconfigCtx.pubZedAgentStatus = pubZedAgentStatus
	pubDatastoreConfig, err := pubsub.Publish(agentName,
		types.DatastoreConfig{})
	if err != nil {
		log.Fatal(err)
	}
	getconfigCtx.pubDatastoreConfig = pubDatastoreConfig
	pubDatastoreConfig.ClearRestarted()

	// Look for global config such as log levels
	subGlobalConfig, err := pubsub.Subscribe("", types.GlobalConfig{},
		false, &zedagentCtx)
	if err != nil {
		log.Fatal(err)
	}
	subGlobalConfig.MaxProcessTimeWarn = warningTime
	subGlobalConfig.MaxProcessTimeError = errorTime
	subGlobalConfig.ModifyHandler = handleGlobalConfigModify
	subGlobalConfig.CreateHandler = handleGlobalConfigModify
	subGlobalConfig.DeleteHandler = handleGlobalConfigDelete
	zedagentCtx.subGlobalConfig = subGlobalConfig
	subGlobalConfig.Activate()

	subNetworkInstanceStatus, err := pubsub.Subscribe("zedrouter",
		types.NetworkInstanceStatus{}, false, &zedagentCtx)
	if err != nil {
		log.Fatal(err)
	}
	subNetworkInstanceStatus.MaxProcessTimeWarn = warningTime
	subNetworkInstanceStatus.MaxProcessTimeError = errorTime
	subNetworkInstanceStatus.ModifyHandler = handleNetworkInstanceModify
	subNetworkInstanceStatus.CreateHandler = handleNetworkInstanceModify
	subNetworkInstanceStatus.DeleteHandler = handleNetworkInstanceDelete
	zedagentCtx.subNetworkInstanceStatus = subNetworkInstanceStatus
	subNetworkInstanceStatus.Activate()

	subNetworkInstanceMetrics, err := pubsub.Subscribe("zedrouter",
		types.NetworkInstanceMetrics{}, false, &zedagentCtx)
	if err != nil {
		log.Fatal(err)
	}
	subNetworkInstanceMetrics.MaxProcessTimeWarn = warningTime
	subNetworkInstanceMetrics.MaxProcessTimeError = errorTime
	subNetworkInstanceMetrics.ModifyHandler = handleNetworkInstanceMetricsModify
	subNetworkInstanceMetrics.CreateHandler = handleNetworkInstanceMetricsModify
	subNetworkInstanceMetrics.DeleteHandler = handleNetworkInstanceMetricsDelete
	zedagentCtx.subNetworkInstanceMetrics = subNetworkInstanceMetrics
	subNetworkInstanceMetrics.Activate()

	subAppFlowMonitor, err := pubsub.Subscribe("zedrouter",
		types.IPFlow{}, false, &zedagentCtx)
	if err != nil {
		log.Fatal(err)
	}
	subAppFlowMonitor.MaxProcessTimeWarn = warningTime
	subAppFlowMonitor.MaxProcessTimeError = errorTime
	subAppFlowMonitor.ModifyHandler = handleAppFlowMonitorModify
	subAppFlowMonitor.CreateHandler = handleAppFlowMonitorModify
	subAppFlowMonitor.DeleteHandler = handleAppFlowMonitorDelete
	subAppFlowMonitor.Activate()
	flowQ = list.New()
	log.Infof("FlowStats: create subFlowStatus")

	subAppVifIPTrig, err := pubsub.Subscribe("zedrouter",
		types.VifIPTrig{}, false, &zedagentCtx)
	if err != nil {
		log.Fatal(err)
	}
	subAppVifIPTrig.MaxProcessTimeWarn = warningTime
	subAppVifIPTrig.MaxProcessTimeError = errorTime
	subAppVifIPTrig.ModifyHandler = handleAppVifIPTrigModify
	subAppVifIPTrig.CreateHandler = handleAppVifIPTrigModify
	subAppVifIPTrig.Activate()

	// Look for AppInstanceStatus from zedmanager
	subAppInstanceStatus, err := pubsub.Subscribe("zedmanager",
		types.AppInstanceStatus{}, false, &zedagentCtx)
	if err != nil {
		log.Fatal(err)
	}
	subAppInstanceStatus.MaxProcessTimeWarn = warningTime
	subAppInstanceStatus.MaxProcessTimeError = errorTime
	subAppInstanceStatus.ModifyHandler = handleAppInstanceStatusModify
	subAppInstanceStatus.CreateHandler = handleAppInstanceStatusModify
	subAppInstanceStatus.DeleteHandler = handleAppInstanceStatusDelete
	getconfigCtx.subAppInstanceStatus = subAppInstanceStatus
	subAppInstanceStatus.Activate()

	// Look for zboot status
	subZbootStatus, err := pubsub.Subscribe("baseosmgr",
		types.ZbootStatus{}, false, &zedagentCtx)
	if err != nil {
		log.Fatal(err)
	}
	subZbootStatus.MaxProcessTimeWarn = warningTime
	subZbootStatus.MaxProcessTimeError = errorTime
	subZbootStatus.ModifyHandler = handleZbootStatusModify
	subZbootStatus.CreateHandler = handleZbootStatusModify
	subZbootStatus.DeleteHandler = handleZbootStatusDelete
	subZbootStatus.RestartHandler = handleZbootRestarted
	zedagentCtx.subZbootStatus = subZbootStatus
	subZbootStatus.Activate()

	subBaseOsStatus, err := pubsub.Subscribe("baseosmgr",
		types.BaseOsStatus{}, false, &zedagentCtx)
	if err != nil {
		log.Fatal(err)
	}
	subBaseOsStatus.MaxProcessTimeWarn = warningTime
	subBaseOsStatus.MaxProcessTimeError = errorTime
	subBaseOsStatus.ModifyHandler = handleBaseOsStatusModify
	subBaseOsStatus.CreateHandler = handleBaseOsStatusModify
	subBaseOsStatus.DeleteHandler = handleBaseOsStatusDelete
	zedagentCtx.subBaseOsStatus = subBaseOsStatus
	subBaseOsStatus.Activate()

	subVaultStatus, err := pubsub.Subscribe("vaultmgr",
		types.VaultStatus{}, false, &zedagentCtx)
	if err != nil {
		log.Fatal(err)
	}
	subVaultStatus.MaxProcessTimeWarn = warningTime
	subVaultStatus.MaxProcessTimeError = errorTime
	subVaultStatus.ModifyHandler = handleVaultStatusModify
	subVaultStatus.DeleteHandler = handleVaultStatusDelete
	zedagentCtx.subVaultStatus = subVaultStatus
	subVaultStatus.Activate()

	// Look for DownloaderStatus from downloader
	// used only for downloader storage stats collection
	subBaseOsDownloadStatus, err := pubsub.SubscribeScope("downloader",
		types.BaseOsObj, types.DownloaderStatus{}, false, &zedagentCtx)
	if err != nil {
		log.Fatal(err)
	}
	subBaseOsDownloadStatus.MaxProcessTimeWarn = warningTime
	subBaseOsDownloadStatus.MaxProcessTimeError = errorTime
	zedagentCtx.subBaseOsDownloadStatus = subBaseOsDownloadStatus
	subBaseOsDownloadStatus.Activate()

	// Look for DownloaderStatus from downloader
	// used only for downloader storage stats collection
	subCertObjDownloadStatus, err := pubsub.SubscribeScope("downloader",
		types.CertObj, types.DownloaderStatus{}, false, &zedagentCtx)
	if err != nil {
		log.Fatal(err)
	}
	subCertObjDownloadStatus.MaxProcessTimeWarn = warningTime
	subCertObjDownloadStatus.MaxProcessTimeError = errorTime
	zedagentCtx.subCertObjDownloadStatus = subCertObjDownloadStatus
	subCertObjDownloadStatus.Activate()

	// Look for VerifyBaseOsImageStatus from verifier
	// used only for verifier storage stats collection
	subBaseOsVerifierStatus, err := pubsub.SubscribeScope("verifier",
		types.BaseOsObj, types.VerifyImageStatus{}, false, &zedagentCtx)
	if err != nil {
		log.Fatal(err)
	}
	subBaseOsVerifierStatus.MaxProcessTimeWarn = warningTime
	subBaseOsVerifierStatus.MaxProcessTimeError = errorTime
	subBaseOsVerifierStatus.ModifyHandler = handleVerifierStatusModify
	subBaseOsVerifierStatus.DeleteHandler = handleVerifierStatusDelete
	subBaseOsVerifierStatus.RestartHandler = handleVerifierRestarted
	zedagentCtx.subBaseOsVerifierStatus = subBaseOsVerifierStatus
	subBaseOsVerifierStatus.Activate()

	// Look for VerifyImageStatus from verifier
	// used only for verifier storage stats collection
	subAppImgVerifierStatus, err := pubsub.SubscribeScope("verifier",
		types.AppImgObj, types.VerifyImageStatus{}, false, &zedagentCtx)
	if err != nil {
		log.Fatal(err)
	}
	subAppImgVerifierStatus.MaxProcessTimeWarn = warningTime
	subAppImgVerifierStatus.MaxProcessTimeError = errorTime
	zedagentCtx.subAppImgVerifierStatus = subAppImgVerifierStatus
	subAppImgVerifierStatus.Activate()

	// Look for DownloaderStatus from downloader for metric reporting
	// used only for downloader storage stats collection
	subAppImgDownloadStatus, err := pubsub.SubscribeScope("downloader",
		types.AppImgObj, types.DownloaderStatus{}, false, &zedagentCtx)
	if err != nil {
		log.Fatal(err)
	}
	subAppImgDownloadStatus.MaxProcessTimeWarn = warningTime
	subAppImgDownloadStatus.MaxProcessTimeError = errorTime
	zedagentCtx.subAppImgDownloadStatus = subAppImgDownloadStatus
	subAppImgDownloadStatus.Activate()

	// Look for nodeagent status
	subNodeAgentStatus, err := pubsub.Subscribe("nodeagent",
		types.NodeAgentStatus{}, false, &getconfigCtx)
	if err != nil {
		log.Fatal(err)
	}
	subNodeAgentStatus.MaxProcessTimeWarn = warningTime
	subNodeAgentStatus.MaxProcessTimeError = errorTime
	subNodeAgentStatus.ModifyHandler = handleNodeAgentStatusModify
	subNodeAgentStatus.DeleteHandler = handleNodeAgentStatusDelete
	getconfigCtx.subNodeAgentStatus = subNodeAgentStatus
	subNodeAgentStatus.Activate()

	DNSctx := DNSContext{}
	DNSctx.usableAddressCount = types.CountLocalAddrAnyNoLinkLocal(*deviceNetworkStatus)

	subDeviceNetworkStatus, err := pubsub.Subscribe("nim",
		types.DeviceNetworkStatus{}, false, &DNSctx)
	if err != nil {
		log.Fatal(err)
	}
	subDeviceNetworkStatus.MaxProcessTimeWarn = warningTime
	subDeviceNetworkStatus.MaxProcessTimeError = errorTime
	subDeviceNetworkStatus.ModifyHandler = handleDNSModify
	subDeviceNetworkStatus.DeleteHandler = handleDNSDelete
	DNSctx.subDeviceNetworkStatus = subDeviceNetworkStatus
	subDeviceNetworkStatus.Activate()

	subDevicePortConfigList, err := pubsub.Subscribe("nim",
		types.DevicePortConfigList{}, false, &zedagentCtx)
	if err != nil {
		log.Fatal(err)
	}
	subDevicePortConfigList.MaxProcessTimeWarn = warningTime
	subDevicePortConfigList.MaxProcessTimeError = errorTime
	subDevicePortConfigList.ModifyHandler = handleDPCLModify
	subDevicePortConfigList.DeleteHandler = handleDPCLDelete
	zedagentCtx.subDevicePortConfigList = subDevicePortConfigList
	subDevicePortConfigList.Activate()

	// Read the GlobalConfig first
	// Wait for initial GlobalConfig
	for !zedagentCtx.GCInitialized {
		log.Infof("Waiting for GCInitialized\n")
		select {
		case change := <-subGlobalConfig.C:
			subGlobalConfig.ProcessChange(change)

		case change := <-getconfigCtx.subNodeAgentStatus.C:
			getconfigCtx.subNodeAgentStatus.ProcessChange(change)
		}
	}

	// wait till, zboot status is ready
	for !zedagentCtx.zbootRestarted {
		select {
		case change := <-subZbootStatus.C:
			subZbootStatus.ProcessChange(change)
			if zedagentCtx.zbootRestarted {
				log.Infof("Zboot reported restarted\n")
			}

		case change := <-getconfigCtx.subNodeAgentStatus.C:
			getconfigCtx.subNodeAgentStatus.ProcessChange(change)

		case <-stillRunning.C:
		}
		agentlog.StillRunning(agentName, warningTime, errorTime)
	}

	log.Infof("Waiting until we have some uplinks with usable addresses\n")
	for !DNSctx.DNSinitialized {
		log.Infof("Waiting for DeviceNetworkStatus %v\n",
			DNSctx.DNSinitialized)

		select {
		case change := <-subGlobalConfig.C:
			subGlobalConfig.ProcessChange(change)

		case change := <-zedagentCtx.subBaseOsVerifierStatus.C:
			zedagentCtx.subBaseOsVerifierStatus.ProcessChange(change)

		case change := <-zedagentCtx.subBaseOsDownloadStatus.C:
			zedagentCtx.subBaseOsDownloadStatus.ProcessChange(change)

		case change := <-zedagentCtx.subAppImgVerifierStatus.C:
			zedagentCtx.subAppImgVerifierStatus.ProcessChange(change)

		case change := <-zedagentCtx.subAppImgDownloadStatus.C:
			zedagentCtx.subAppImgDownloadStatus.ProcessChange(change)

		case change := <-zedagentCtx.subCertObjDownloadStatus.C:
			zedagentCtx.subCertObjDownloadStatus.ProcessChange(change)

		case change := <-subDeviceNetworkStatus.C:
			subDeviceNetworkStatus.ProcessChange(change)

		case change := <-subAssignableAdapters.C:
			subAssignableAdapters.ProcessChange(change)

		case change := <-subDevicePortConfigList.C:
			subDevicePortConfigList.ProcessChange(change)

		case change := <-getconfigCtx.subNodeAgentStatus.C:
			getconfigCtx.subNodeAgentStatus.ProcessChange(change)

		case change := <-subVaultStatus.C:
			subVaultStatus.ProcessChange(change)

		case change := <-deferredChan:
			start := time.Now()
			zedcloud.HandleDeferred(change, 100*time.Millisecond)
			pubsub.CheckMaxTimeTopic(agentName, "deferredChan", start,
				warningTime, errorTime)

		case <-stillRunning.C:
		}
		agentlog.StillRunning(agentName, warningTime, errorTime)
	}

	// Subscribe to network metrics from zedrouter
	subNetworkMetrics, err := pubsub.Subscribe("zedrouter",
		types.NetworkMetrics{}, true, &zedagentCtx)
	if err != nil {
		log.Fatal(err)
	}
	// Subscribe to cloud metrics from different agents
	cms := zedcloud.GetCloudMetrics()
	subClientMetrics, err := pubsub.Subscribe("zedclient", cms,
		true, &zedagentCtx)
	if err != nil {
		log.Fatal(err)
	}
	subLogmanagerMetrics, err := pubsub.Subscribe("logmanager",
		cms, true, &zedagentCtx)
	if err != nil {
		log.Fatal(err)
	}
	subDownloaderMetrics, err := pubsub.Subscribe("downloader",
		cms, true, &zedagentCtx)
	if err != nil {
		log.Fatal(err)
	}

	// Use a go routine to make sure we have wait/timeout without
	// blocking the main select loop
	go deviceInfoTask(&zedagentCtx, triggerDeviceInfo)

	// Publish initial device info.
	triggerPublishDevInfo(&zedagentCtx)

	// start the metrics reporting task
	handleChannel := make(chan interface{})
	go metricsTimerTask(&zedagentCtx, handleChannel)
	metricsTickerHandle := <-handleChannel
	getconfigCtx.metricsTickerHandle = metricsTickerHandle

	// Process the verifierStatus to avoid downloading an image we
	// already have in place
	log.Infof("Handling initial verifier Status\n")
	for !zedagentCtx.verifierRestarted {
		select {
		case change := <-subZbootStatus.C:
			subZbootStatus.ProcessChange(change)

		case change := <-subGlobalConfig.C:
			subGlobalConfig.ProcessChange(change)

		case change := <-subBaseOsVerifierStatus.C:
			subBaseOsVerifierStatus.ProcessChange(change)
			if zedagentCtx.verifierRestarted {
				log.Infof("Verifier reported restarted\n")
				break
			}

		case change := <-zedagentCtx.subBaseOsDownloadStatus.C:
			zedagentCtx.subBaseOsDownloadStatus.ProcessChange(change)

		case change := <-zedagentCtx.subAppImgVerifierStatus.C:
			zedagentCtx.subAppImgVerifierStatus.ProcessChange(change)

		case change := <-zedagentCtx.subAppImgDownloadStatus.C:
			zedagentCtx.subAppImgDownloadStatus.ProcessChange(change)

		case change := <-zedagentCtx.subCertObjDownloadStatus.C:
			zedagentCtx.subCertObjDownloadStatus.ProcessChange(change)

		case change := <-getconfigCtx.subNodeAgentStatus.C:
			getconfigCtx.subNodeAgentStatus.ProcessChange(change)

		case change := <-subDeviceNetworkStatus.C:
			subDeviceNetworkStatus.ProcessChange(change)
			if DNSctx.triggerDeviceInfo {
				// IP/DNS in device info could have changed
				log.Infof("NetworkStatus triggered PublishDeviceInfo\n")
				triggerPublishDevInfo(&zedagentCtx)
				DNSctx.triggerDeviceInfo = false
			}

		case change := <-subAssignableAdapters.C:
			subAssignableAdapters.ProcessChange(change)

		case change := <-subDevicePortConfigList.C:
			subDevicePortConfigList.ProcessChange(change)

		case change := <-subVaultStatus.C:
			subVaultStatus.ProcessChange(change)

		case change := <-deferredChan:
			zedcloud.HandleDeferred(change, 100*time.Millisecond)

		case <-stillRunning.C:
		}
		// XXX verifierRestarted can take 5 minutes??
		agentlog.StillRunning(agentName, warningTime, errorTime)
		// Need to tickle this since the configTimerTask is not yet started
		agentlog.StillRunning(agentName+"config", warningTime, errorTime)
	}

	// start the config fetch tasks, when zboot status is ready
	go configTimerTask(handleChannel, &getconfigCtx)
	configTickerHandle := <-handleChannel
	// XXX close handleChannels?
	getconfigCtx.configTickerHandle = configTickerHandle

	for {
		select {
		case change := <-subZbootStatus.C:
			subZbootStatus.ProcessChange(change)

		case change := <-subGlobalConfig.C:
			subGlobalConfig.ProcessChange(change)

		case change := <-subAppInstanceStatus.C:
			subAppInstanceStatus.ProcessChange(change)

		case change := <-subBaseOsStatus.C:
			subBaseOsStatus.ProcessChange(change)

		case change := <-zedagentCtx.subBaseOsVerifierStatus.C:
			zedagentCtx.subBaseOsVerifierStatus.ProcessChange(change)

		case change := <-zedagentCtx.subBaseOsDownloadStatus.C:
			zedagentCtx.subBaseOsDownloadStatus.ProcessChange(change)

		case change := <-zedagentCtx.subAppImgVerifierStatus.C:
			zedagentCtx.subAppImgVerifierStatus.ProcessChange(change)

		case change := <-zedagentCtx.subAppImgDownloadStatus.C:
			zedagentCtx.subAppImgDownloadStatus.ProcessChange(change)

		case change := <-zedagentCtx.subCertObjDownloadStatus.C:
			zedagentCtx.subCertObjDownloadStatus.ProcessChange(change)

		case change := <-getconfigCtx.subNodeAgentStatus.C:
			getconfigCtx.subNodeAgentStatus.ProcessChange(change)

		case change := <-subDeviceNetworkStatus.C:
			subDeviceNetworkStatus.ProcessChange(change)
			if DNSctx.triggerGetConfig {
				triggerGetConfig(configTickerHandle)
				DNSctx.triggerGetConfig = false
			}
			if DNSctx.triggerDeviceInfo {
				// IP/DNS in device info could have changed
				log.Infof("NetworkStatus triggered PublishDeviceInfo\n")
				triggerPublishDevInfo(&zedagentCtx)
				DNSctx.triggerDeviceInfo = false
			}

		case change := <-subAssignableAdapters.C:
			subAssignableAdapters.ProcessChange(change)

		case change := <-subNetworkMetrics.C:
			subNetworkMetrics.ProcessChange(change)
			m, err := subNetworkMetrics.Get("global")
			if err != nil {
				log.Errorf("subNetworkMetrics.Get failed: %s\n",
					err)
			} else {
				networkMetrics = types.CastNetworkMetrics(m)
			}

		case change := <-subClientMetrics.C:
			subClientMetrics.ProcessChange(change)
			m, err := subClientMetrics.Get("global")
			if err != nil {
				log.Errorf("subClientMetrics.Get failed: %s\n",
					err)
			} else {
				clientMetrics = m
			}

		case change := <-subLogmanagerMetrics.C:
			subLogmanagerMetrics.ProcessChange(change)
			m, err := subLogmanagerMetrics.Get("global")
			if err != nil {
				log.Errorf("subLogmanagerMetrics.Get failed: %s\n",
					err)
			} else {
				logmanagerMetrics = m
			}

		case change := <-subDownloaderMetrics.C:
			subDownloaderMetrics.ProcessChange(change)
			m, err := subDownloaderMetrics.Get("global")
			if err != nil {
				log.Errorf("subDownloaderMetrics.Get failed: %s\n",
					err)
			} else {
				downloaderMetrics = m
			}

		case change := <-deferredChan:
			start := time.Now()
			zedcloud.HandleDeferred(change, 100*time.Millisecond)
			pubsub.CheckMaxTimeTopic(agentName, "deferredChan", start,
				warningTime, errorTime)

		case change := <-subNetworkInstanceStatus.C:
			subNetworkInstanceStatus.ProcessChange(change)

		case change := <-subNetworkInstanceMetrics.C:
			subNetworkInstanceMetrics.ProcessChange(change)

		case change := <-subDevicePortConfigList.C:
			subDevicePortConfigList.ProcessChange(change)

		case change := <-subAppFlowMonitor.C:
			log.Debugf("FlowStats: change called")
			subAppFlowMonitor.ProcessChange(change)

		case change := <-subAppVifIPTrig.C:
			subAppVifIPTrig.ProcessChange(change)

		case change := <-subVaultStatus.C:
			subVaultStatus.ProcessChange(change)

		case <-stillRunning.C:
		}
		agentlog.StillRunning(agentName, warningTime, errorTime)
	}
}

func triggerPublishDevInfo(ctxPtr *zedagentContext) {

	log.Info("Triggered PublishDeviceInfo")
	select {
	case ctxPtr.TriggerDeviceInfo <- struct{}{}:
		// Do nothing more
	default:
		// This occurs if we are already trying to send a device info
		// and we get a second and third trigger before that is complete.
		log.Warnf("Failed to send on PublishDeviceInfo")
	}
}

func deviceInfoTask(ctxPtr *zedagentContext, triggerDeviceInfo <-chan struct{}) {

	// Run a periodic timer so we always update StillRunning
	stillRunning := time.NewTicker(25 * time.Second)

	for {
		select {
		case <-triggerDeviceInfo:
			start := time.Now()
			log.Info("deviceInfoTask got message")

			PublishDeviceInfoToZedCloud(ctxPtr)
			ctxPtr.iteration++
			log.Info("deviceInfoTask done with message")
			pubsub.CheckMaxTimeTopic(agentName+"devinfo", "PublishDeviceInfo", start,
				warningTime, errorTime)
		case <-stillRunning.C:
		}
		agentlog.StillRunning(agentName+"devinfo", warningTime, errorTime)
	}
}

func handleVerifierRestarted(ctxArg interface{}, done bool) {
	ctx := ctxArg.(*zedagentContext)
	log.Infof("handleVerifierRestarted(%v)\n", done)
	if done {
		ctx.verifierRestarted = true
	}
}

// base os verifier status modify event
func handleVerifierStatusModify(ctxArg interface{}, key string,
	statusArg interface{}) {

	status := cast.CastVerifyImageStatus(statusArg)
	if status.Key() != key {
		log.Errorf("handleVerifierStatusModify key/UUID mismatch %s vs %s; ignored %+v\n",
			key, status.Key(), status)
		return
	}
	log.Infof("handleVerifierStatusModify for %s\n", status.Safename)
	// Nothing to do
}

// base os verifier status delete event
func handleVerifierStatusDelete(ctxArg interface{}, key string,
	statusArg interface{}) {

	status := cast.CastVerifyImageStatus(statusArg)
	log.Infof("handleVeriferStatusDelete RefCount %d Expired %v for %s\n",
		status.RefCount, status.Expired, key)
	// Nothing to do
}

func handleZbootRestarted(ctxArg interface{}, done bool) {
	ctx := ctxArg.(*zedagentContext)
	log.Infof("handleZbootRestarted(%v)\n", done)
	if done {
		ctx.zbootRestarted = true
	}
}

func handleInit(networkSendTimeout uint32) {
	initializeDirs()
	handleConfigInit(networkSendTimeout)
}

func initializeDirs() {

	// create persistent holder directory
	if _, err := os.Stat(types.PersistDir); err != nil {
		log.Debugf("Create %s\n", types.PersistDir)
		if err := os.MkdirAll(types.PersistDir, 0700); err != nil {
			log.Fatal(err)
		}
	}
	if _, err := os.Stat(types.CertificateDirname); err != nil {
		log.Debugf("Create %s\n", types.CertificateDirname)
		if err := os.MkdirAll(types.CertificateDirname, 0700); err != nil {
			log.Fatal(err)
		}
	}
	if _, err := os.Stat(checkpointDirname); err != nil {
		log.Debugf("Create %s\n", checkpointDirname)
		if err := os.MkdirAll(checkpointDirname, 0700); err != nil {
			log.Fatal(err)
		}
	}
	if _, err := os.Stat(types.DownloadDirname); err != nil {
		log.Debugf("Create %s\n", types.DownloadDirname)
		if err := os.MkdirAll(types.DownloadDirname, 0700); err != nil {
			log.Fatal(err)
		}
	}
}

// app instance event watch to capture transitions
// and publish to zedCloud
// Handles both create and modify events
func handleAppInstanceStatusModify(ctxArg interface{}, key string,
	statusArg interface{}) {
	status := cast.CastAppInstanceStatus(statusArg)
	if status.Key() != key {
		log.Errorf("handleAppInstanceStatusModify key/UUID mismatch %s vs %s; ignored %+v\n",
			key, status.Key(), status)
		return
	}
	ctx := ctxArg.(*zedagentContext)
	uuidStr := status.Key()
	PublishAppInfoToZedCloud(ctx, uuidStr, &status, ctx.assignableAdapters,
		ctx.iteration)
	ctx.iteration++
}

func handleAppInstanceStatusDelete(ctxArg interface{}, key string,
	statusArg interface{}) {

	ctx := ctxArg.(*zedagentContext)
	uuidStr := key
	PublishAppInfoToZedCloud(ctx, uuidStr, nil, ctx.assignableAdapters,
		ctx.iteration)
	ctx.iteration++
}

func lookupAppInstanceStatus(ctx *zedagentContext, key string) *types.AppInstanceStatus {

	sub := ctx.getconfigCtx.subAppInstanceStatus
	st, _ := sub.Get(key)
	if st == nil {
		log.Infof("lookupAppInstanceStatus(%s) not found\n", key)
		return nil
	}
	status := cast.CastAppInstanceStatus(st)
	if status.Key() != key {
		log.Errorf("lookupAppInstanceStatus key/UUID mismatch %s vs %s; ignored %+v\n",
			key, status.Key(), status)
		return nil
	}
	return &status
}

func handleDNSModify(ctxArg interface{}, key string, statusArg interface{}) {

	status := cast.CastDeviceNetworkStatus(statusArg)
	ctx := ctxArg.(*DNSContext)
	if key != "global" {
		log.Infof("handleDNSModify: ignoring %s\n", key)
		return
	}
	log.Infof("handleDNSModify for %s\n", key)
	if cmp.Equal(*deviceNetworkStatus, status) {
		log.Infof("handleDNSModify no change\n")
		return
	}
	log.Infof("handleDNSModify: changed %v",
		cmp.Diff(*deviceNetworkStatus, status))
	*deviceNetworkStatus = status
	// Did we (re-)gain the first usable address?
	// XXX should we also trigger if the count increases?
	newAddrCount := types.CountLocalAddrAnyNoLinkLocal(*deviceNetworkStatus)
	if newAddrCount != 0 && ctx.usableAddressCount == 0 {
		log.Infof("DeviceNetworkStatus from %d to %d addresses\n",
			ctx.usableAddressCount, newAddrCount)
		ctx.triggerGetConfig = true
	}
	ctx.DNSinitialized = true
	ctx.usableAddressCount = newAddrCount
	ctx.triggerDeviceInfo = true
	log.Infof("handleDNSModify done for %s\n", key)
}

func handleDNSDelete(ctxArg interface{}, key string,
	statusArg interface{}) {

	log.Infof("handleDNSDelete for %s\n", key)
	ctx := ctxArg.(*DNSContext)

	if key != "global" {
		log.Infof("handleDNSDelete: ignoring %s\n", key)
		return
	}
	*deviceNetworkStatus = types.DeviceNetworkStatus{}
	newAddrCount := types.CountLocalAddrAnyNoLinkLocal(*deviceNetworkStatus)
	ctx.DNSinitialized = false
	ctx.usableAddressCount = newAddrCount
	log.Infof("handleDNSDelete done for %s\n", key)
}

func handleDPCLModify(ctxArg interface{}, key string, statusArg interface{}) {

	status := cast.CastDevicePortConfigList(statusArg)
	ctx := ctxArg.(*zedagentContext)
	if key != "global" {
		log.Infof("handleDPCLModify: ignoring %s\n", key)
		return
	}
	if cmp.Equal(ctx.devicePortConfigList, status) {
		log.Infof("handleDPCLModify no change\n")
		return
	}
	// Note that lastSucceeded will increment a lot; ignore it but compare
	// lastFailed/lastError?? XXX how?
	log.Infof("handleDPCLModify: changed %v",
		cmp.Diff(ctx.devicePortConfigList, status))
	ctx.devicePortConfigList = status
	triggerPublishDevInfo(ctx)
}

func handleDPCLDelete(ctxArg interface{}, key string, statusArg interface{}) {

	ctx := ctxArg.(*zedagentContext)
	if key != "global" {
		log.Infof("handleDPCLDelete: ignoring %s\n", key)
		return
	}
	log.Infof("handleDPCLDelete for %s\n", key)
	ctx.devicePortConfigList = types.DevicePortConfigList{}
	triggerPublishDevInfo(ctx)
}

// base os status event handlers
// Report BaseOsStatus to zedcloud
// Handles both create and modify events
func handleBaseOsStatusModify(ctxArg interface{}, key string, statusArg interface{}) {
	ctx := ctxArg.(*zedagentContext)
	status := cast.CastBaseOsStatus(statusArg)
	if status.Key() != key {
		log.Errorf("handleBaseOsStatusModify key/UUID mismatch %s vs %s; ignored %+v\n", key, status.Key(), status)
		return
	}
	triggerPublishDevInfo(ctx)
	log.Infof("handleBaseOsStatusModify(%s) done\n", key)
}

func handleBaseOsStatusDelete(ctxArg interface{}, key string,
	statusArg interface{}) {

	log.Infof("handleBaseOsStatusDelete(%s)\n", key)
	ctx := ctxArg.(*zedagentContext)
	triggerPublishDevInfo(ctx)
	log.Infof("handleBaseOsStatusDelete(%s) done\n", key)
}

// vault status event handlers
// Report VaultStatus to zedcloud
func handleVaultStatusModify(ctxArg interface{}, key string, statusArg interface{}) {
	ctx := ctxArg.(*zedagentContext)
	status := cast.VaultStatus(statusArg)
	if status.Key() != key {
		log.Errorf("handleVaultStatusModify key/UUID mismatch %s vs %s; ignored %+v\n", key, status.Key(), status)
		return
	}
	triggerPublishDevInfo(ctx)
	log.Infof("handleVaultStatusModify(%s) done\n", key)
}

func handleVaultStatusDelete(ctxArg interface{}, key string,
	statusArg interface{}) {

	log.Infof("handleVaultStatusDelete(%s)\n", key)
	ctx := ctxArg.(*zedagentContext)
	triggerPublishDevInfo(ctx)
	log.Infof("handleVaultStatusDelete(%s) done\n", key)
}

func appendError(allErrors string, prefix string, lasterr string) string {
	return fmt.Sprintf("%s%s: %s\n\n", allErrors, prefix, lasterr)
}

// Handles both create and modify events
func handleGlobalConfigModify(ctxArg interface{}, key string,
	statusArg interface{}) {

	ctx := ctxArg.(*zedagentContext)
	if key != "global" {
		log.Infof("handleGlobalConfigModify: ignoring %s\n", key)
		return
	}
	log.Infof("handleGlobalConfigModify for %s\n", key)
	var gcp *types.GlobalConfig
	debug, gcp = agentlog.HandleGlobalConfig(ctx.subGlobalConfig, agentName,
		debugOverride)
	if gcp != nil && !ctx.GCInitialized {
		updated := types.ApplyGlobalConfig(*gcp)
		log.Infof("handleGlobalConfigModify setting initials to %+v\n",
			updated)
		sane := types.EnforceGlobalConfigMinimums(updated)
		log.Infof("handleGlobalConfigModify: enforced minimums %v\n",
			cmp.Diff(updated, sane))
		ctx.globalConfig = sane
		ctx.GCInitialized = true
	}
	log.Infof("handleGlobalConfigModify done for %s\n", key)
}

func handleGlobalConfigDelete(ctxArg interface{}, key string,
	statusArg interface{}) {

	ctx := ctxArg.(*zedagentContext)
	if key != "global" {
		log.Infof("handleGlobalConfigDelete: ignoring %s\n", key)
		return
	}
	log.Infof("handleGlobalConfigDelete for %s\n", key)
	debug, _ = agentlog.HandleGlobalConfig(ctx.subGlobalConfig, agentName,
		debugOverride)
	ctx.globalConfig = types.GlobalConfigDefaults
	log.Infof("handleGlobalConfigDelete done for %s\n", key)
}

// Handles both create and modify events
func handleAAModify(ctxArg interface{}, key string,
	statusArg interface{}) {

	ctx := ctxArg.(*zedagentContext)
	status := cast.CastAssignableAdapters(statusArg)
	if key != "global" {
		log.Infof("handleAAModify: ignoring %s\n", key)
		return
	}
	log.Infof("handleAAModify() %+v\n", status)
	*ctx.assignableAdapters = status
	triggerPublishDevInfo(ctx)
	log.Infof("handleAAModify() done\n")
}

func handleAADelete(ctxArg interface{}, key string,
	statusArg interface{}) {

	ctx := ctxArg.(*zedagentContext)
	if key != "global" {
		log.Infof("handleAADelete: ignoring %s\n", key)
		return
	}
	log.Infof("handleAADelete()\n")
	ctx.assignableAdapters.Initialized = false
	triggerPublishDevInfo(ctx)
	log.Infof("handleAADelete() done\n")
}

// Handles both create and modify events
func handleZbootStatusModify(ctxArg interface{}, key string,
	statusArg interface{}) {

	ctx := ctxArg.(*zedagentContext)
	if !isZbootValidPartitionLabel(key) {
		log.Errorf("handleZbootStatusModify: invalid key %s\n", key)
		return
	}
	log.Infof("handleZbootStatusModify: for %s\n", key)
	// nothing to do
	triggerPublishDevInfo(ctx)
}

func handleZbootStatusDelete(ctxArg interface{}, key string,
	statusArg interface{}) {
	if !isZbootValidPartitionLabel(key) {
		log.Errorf("handleZbootStatusDelete: invalid key %s\n", key)
		return
	}
	log.Infof("handleZbootStatusDelete: for %s\n", key)
	// Nothing to do
}

func handleNodeAgentStatusModify(ctxArg interface{}, key string,
	statusArg interface{}) {

	getconfigCtx := ctxArg.(*getconfigContext)
	status := cast.NodeAgentStatus(statusArg)
	if status.Key() != key {
		log.Errorf("handleNodeAgentStatusModify: invalid key %s\n", key)
		return
	}
	updateInprogress := getconfigCtx.updateInprogress
	ctx := getconfigCtx.zedagentCtx
	ctx.remainingTestTime = status.RemainingTestTime
	getconfigCtx.updateInprogress = status.UpdateInprogress
	ctx.rebootTime = status.RebootTime
	ctx.rebootStack = status.RebootStack
	ctx.rebootReason = status.RebootReason
	ctx.restartCounter = status.RestartCounter
	// if config reboot command was initiated and
	// was deferred, and the device is not in inprogress
	// state, initiate the reboot process
	if ctx.rebootCmdDeferred &&
		updateInprogress && !status.UpdateInprogress {
		log.Infof("TestComplete and deferred reboot\n")
		ctx.rebootCmdDeferred = false
		infoStr := fmt.Sprintf("TestComplete and deferred Reboot Cmd\n")
		handleRebootCmd(ctx, infoStr)
	}
	if status.DeviceReboot {
		handleDeviceReboot(ctx)
	}
	triggerPublishDevInfo(ctx)
	log.Infof("handleNodeAgentStatusModify: done.\n")
}

func handleNodeAgentStatusDelete(ctxArg interface{}, key string,
	statusArg interface{}) {
	ctx := ctxArg.(*zedagentContext)
	status := cast.NodeAgentStatus(statusArg)
	if status.Key() != key {
		log.Errorf("handleNodeAgentStatusDelete: invalid key %s\n", key)
		return
	}
	log.Infof("handleNodeAgentStatusDelete: for %s\n", key)
	// Nothing to do
	triggerPublishDevInfo(ctx)
}
