// Copyright (c) 2017-2018 Zededa, Inc.
// SPDX-License-Identifier: Apache-2.0

package client

import (
	"bytes"
	"crypto/tls"
	"encoding/base64"
	"flag"
	"fmt"
	"io/ioutil"
	"mime"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/golang/protobuf/proto"
	"github.com/google/go-cmp/cmp"
	"github.com/lf-edge/eve/api/go/register"
	"github.com/lf-edge/eve/pkg/pillar/agentlog"
	"github.com/lf-edge/eve/pkg/pillar/flextimer"
	"github.com/lf-edge/eve/pkg/pillar/hardware"
	"github.com/lf-edge/eve/pkg/pillar/pidfile"
	"github.com/lf-edge/eve/pkg/pillar/pubsub"
	"github.com/lf-edge/eve/pkg/pillar/types"
	"github.com/lf-edge/eve/pkg/pillar/utils"
	fileutils "github.com/lf-edge/eve/pkg/pillar/utils/file"
	"github.com/lf-edge/eve/pkg/pillar/zedcloud"
	"github.com/satori/go.uuid"
	log "github.com/sirupsen/logrus"
)

const (
	agentName   = "zedclient"
	maxDelay    = time.Second * 600 // 10 minutes
	uuidMaxWait = time.Second * 60  // 1 minute
	// Time limits for event loop handlers
	errorTime   = 3 * time.Minute
	warningTime = 40 * time.Second
	return400   = false
)

// Really a constant
var nilUUID uuid.UUID

// Set from Makefile
var Version = "No version specified"

// Assumes the config files are in IdentityDirname, which is /config
// by default. The files are
//  root-certificate.pem	Root CA cert(s) for object signing
//  server			Fixed? Written if redirected. factory-root-cert?
//  onboard.cert.pem, onboard.key.pem	Per device onboarding certificate/key
//  		   		for selfRegister operation
//  device.cert.pem,
//  device.key.pem		Device certificate/key created before this
//  		     		client is started.
//  uuid			Written by getUuid operation
//  hardwaremodel		Written by getUuid if server returns a hardwaremodel
//  enterprise			Written by getUuid if server returns an enterprise
//  name			Written by getUuid if server returns a name
//
//

type clientContext struct {
	subDeviceNetworkStatus pubsub.Subscription
	deviceNetworkStatus    *types.DeviceNetworkStatus
	usableAddressCount     int
	subGlobalConfig        pubsub.Subscription
	globalConfig           *types.GlobalConfig
	zedcloudCtx            *zedcloud.ZedCloudContext
	getCertsTimer          *time.Timer
}

var (
	debug             = false
	debugOverride     bool // From command line arg
	serverNameAndPort string
	serverName        string
	onboardTLSConfig  *tls.Config
	devtlsConfig      *tls.Config
)

func Run(ps *pubsub.PubSub) { //nolint:gocyclo
	versionPtr := flag.Bool("v", false, "Version")
	debugPtr := flag.Bool("d", false, "Debug flag")
	curpartPtr := flag.String("c", "", "Current partition")
	noPidPtr := flag.Bool("p", false, "Do not check for running client")
	maxRetriesPtr := flag.Int("r", 0, "Max retries")
	flag.Parse()

	versionFlag := *versionPtr
	debug = *debugPtr
	debugOverride = debug
	if debugOverride {
		log.SetLevel(log.DebugLevel)
	} else {
		log.SetLevel(log.InfoLevel)
	}
	curpart := *curpartPtr
	noPidFlag := *noPidPtr
	maxRetries := *maxRetriesPtr
	args := flag.Args()
	if versionFlag {
		fmt.Printf("%s: %s\n", os.Args[0], Version)
		return
	}
	// Sending json log format to stdout
	err := agentlog.Init("client", curpart)
	if err != nil {
		log.Fatal(err)
	}
	if !noPidFlag {
		if err := pidfile.CheckAndCreatePidfile(agentName); err != nil {
			log.Fatal(err)
		}
	}
	log.Infof("Starting %s\n", agentName)
	operations := map[string]bool{
		"selfRegister": false,
		"getUuid":      false,
	}
	for _, op := range args {
		if _, ok := operations[op]; ok {
			operations[op] = true
		} else {
			log.Errorf("Unknown arg %s\n", op)
			log.Fatal("Usage: " + os.Args[0] +
				"[-o] [<operations>...]")
		}
	}

	hardwaremodelFileName := types.IdentityDirname + "/hardwaremodel"
	enterpriseFileName := types.IdentityDirname + "/enterprise"
	nameFileName := types.IdentityDirname + "/name"

	cms := zedcloud.GetCloudMetrics() // Need type of data
	pub, err := ps.NewPublication(pubsub.PublicationOptions{
		AgentName: agentName,
		TopicType: cms,
	})
	if err != nil {
		log.Fatal(err)
	}

	var oldUUID uuid.UUID
	b, err := ioutil.ReadFile(types.UUIDFileName)
	if err == nil {
		uuidStr := strings.TrimSpace(string(b))
		oldUUID, err = uuid.FromString(uuidStr)
		if err != nil {
			log.Warningf("Malformed UUID file ignored: %s\n", err)
		}
	}
	// Check if we have a /config/hardwaremodel file
	oldHardwaremodel := hardware.GetHardwareModelOverride()

	clientCtx := clientContext{
		deviceNetworkStatus: &types.DeviceNetworkStatus{},
		globalConfig:        &types.GlobalConfigDefaults,
	}

	// Look for global config such as log levels
	subGlobalConfig, err := ps.NewSubscription(pubsub.SubscriptionOptions{
		CreateHandler: handleGlobalConfigModify,
		ModifyHandler: handleGlobalConfigModify,
		DeleteHandler: handleGlobalConfigDelete,
		WarningTime:   warningTime,
		ErrorTime:     errorTime,
		TopicImpl:     types.GlobalConfig{},
		Ctx:           &clientCtx,
	})

	if err != nil {
		log.Fatal(err)
	}
	clientCtx.subGlobalConfig = subGlobalConfig
	subGlobalConfig.Activate()

	subDeviceNetworkStatus, err := ps.NewSubscription(pubsub.SubscriptionOptions{
		CreateHandler: handleDNSModify,
		ModifyHandler: handleDNSModify,
		DeleteHandler: handleDNSDelete,
		WarningTime:   warningTime,
		ErrorTime:     errorTime,
		AgentName:     "nim",
		TopicImpl:     types.DeviceNetworkStatus{},
		Ctx:           &clientCtx,
	})
	if err != nil {
		log.Fatal(err)
	}
	clientCtx.subDeviceNetworkStatus = subDeviceNetworkStatus
	subDeviceNetworkStatus.Activate()

	zedcloudCtx := zedcloud.ZedCloudContext{
		DeviceNetworkStatus: clientCtx.deviceNetworkStatus,
		FailureFunc:         zedcloud.ZedCloudFailure,
		SuccessFunc:         zedcloud.ZedCloudSuccess,
		NetworkSendTimeout:  clientCtx.globalConfig.NetworkSendTimeout,
		V2API:               zedcloud.UseV2API(),
	}

	// Get device serial number
	zedcloudCtx.DevSerial = hardware.GetProductSerial()
	zedcloudCtx.DevSoftSerial = hardware.GetSoftSerial()
	clientCtx.zedcloudCtx = &zedcloudCtx
	log.Infof("Client Get Device Serial %s, Soft Serial %s\n", zedcloudCtx.DevSerial,
		zedcloudCtx.DevSoftSerial)

	// Run a periodic timer so we always update StillRunning
	stillRunning := time.NewTicker(25 * time.Second)
	agentlog.StillRunning(agentName, warningTime, errorTime)

	// Wait for a usable IP address.
	// After 5 seconds we check; if we already have a UUID we proceed.
	// That ensures that we will start zedagent and it will check
	// the cloudGoneTime if we are doing an imake update.
	t1 := time.NewTimer(5 * time.Second)

	ticker := flextimer.NewExpTicker(time.Second, maxDelay, 0.0)

	// XXX redo in ticker case to handle change to servername?
	server, err := ioutil.ReadFile(types.ServerFileName)
	if err != nil {
		log.Fatal(err)
	}
	serverNameAndPort = strings.TrimSpace(string(server))
	serverName := strings.Split(serverNameAndPort, ":")[0]

	var onboardCert tls.Certificate
	var deviceCertPem []byte
	var gotServerCerts bool

	if operations["selfRegister"] {
		var err error
		onboardCert, err = tls.LoadX509KeyPair(types.OnboardCertName,
			types.OnboardKeyName)
		if err != nil {
			log.Fatal(err)
		}
		onboardTLSConfig, err = zedcloud.GetTlsConfig(zedcloudCtx.DeviceNetworkStatus,
			serverName, &onboardCert, &zedcloudCtx)
		if err != nil {
			log.Fatal(err)
		}
		// Load device text cert for upload
		deviceCertPem, err = ioutil.ReadFile(types.DeviceCertName)
		if err != nil {
			log.Fatal(err)
		}
	}

	// Load device cert
	deviceCert, err := zedcloud.GetClientCert()
	if err != nil {
		log.Fatal(err)
	}
	devtlsConfig, err = zedcloud.GetTlsConfig(zedcloudCtx.DeviceNetworkStatus,
		serverName, &deviceCert, &zedcloudCtx)
	if err != nil {
		log.Fatal(err)
	}

	done := false
	var devUUID uuid.UUID
	var hardwaremodel string
	var enterprise string
	var name string
	gotUUID := false
	gotRegister := false
	retryCount := 0
	clientCtx.getCertsTimer = time.NewTimer(1 * time.Second)
	clientCtx.getCertsTimer.Stop()

	for !done {
		log.Infof("Waiting for usableAddressCount %d and done %v\n",
			clientCtx.usableAddressCount, done)
		select {
		case change := <-subGlobalConfig.MsgChan():
			subGlobalConfig.ProcessChange(change)

		case change := <-subDeviceNetworkStatus.MsgChan():
			subDeviceNetworkStatus.ProcessChange(change)

		case <-ticker.C:
			if clientCtx.usableAddressCount == 0 {
				log.Infof("ticker and no usableAddressCount")
				// XXX keep exponential unchanged?
				break
			}

			// try to fetch the server certs chain first, if it's V2
			if !gotServerCerts && zedcloudCtx.V2API {
				gotServerCerts = fetchCertChain(zedcloudCtx, devtlsConfig, retryCount, true) // XXX always get certs from cloud for now
				log.Infof("client fetchCertChain, gotServerCerts %v\n", gotServerCerts)
				if !gotServerCerts {
					break
				}
			}

			if !gotRegister && operations["selfRegister"] {
				done = selfRegister(zedcloudCtx, onboardTLSConfig, deviceCertPem, retryCount)
				if done {
					gotRegister = true
				}
				if !done && operations["getUuid"] {
					// Check if getUUid succeeds
					done, devUUID, hardwaremodel, enterprise, name = doGetUUID(&clientCtx, devtlsConfig, retryCount)
					if done {
						log.Infof("getUUID succeeded; selfRegister no longer needed")
						gotUUID = true
					}
				}
			}
			if !gotUUID && operations["getUuid"] {
				done, devUUID, hardwaremodel, enterprise, name = doGetUUID(&clientCtx, devtlsConfig, retryCount)
				if done {
					log.Infof("getUUID succeeded; selfRegister no longer needed")
					gotUUID = true
				}
				if oldUUID != nilUUID && retryCount > 2 {
					log.Infof("Sticking with old UUID\n")
					devUUID = oldUUID
					done = true
					break
				}
			}
			retryCount++
			if maxRetries != 0 && retryCount > maxRetries {
				log.Errorf("Exceeded %d retries",
					maxRetries)
				os.Exit(1)
			}

		case <-t1.C:
			// If we already know a uuid we can skip
			// This might not set hardwaremodel when upgrading
			// an onboarded system without /config/hardwaremodel.
			// Unlikely to have a network outage during that
			// upgrade *and* require an override.
			if clientCtx.usableAddressCount == 0 &&
				operations["getUuid"] && oldUUID != nilUUID {

				log.Infof("Already have a UUID %s; declaring success\n",
					oldUUID.String())
				done = true
			}

		case <-clientCtx.getCertsTimer.C:
			// triggered by cert miss error in doGetUUID, so the TLS is device TLSConfig
			ok := fetchCertChain(zedcloudCtx, devtlsConfig, retryCount, true)
			log.Infof("client timer get cert chain %v\n", ok)

		case <-stillRunning.C:
		}
		agentlog.StillRunning(agentName, warningTime, errorTime)
	}

	// Post loop code
	if devUUID != nilUUID {
		doWrite := true
		if oldUUID != nilUUID {
			if oldUUID != devUUID {
				log.Infof("Replacing existing UUID %s\n",
					oldUUID.String())
			} else {
				log.Infof("No change to UUID %s\n",
					devUUID)
				doWrite = false
			}
		} else {
			log.Infof("Got config with UUID %s\n", devUUID)
		}

		if doWrite {
			b := []byte(fmt.Sprintf("%s\n", devUUID))
			err = ioutil.WriteFile(types.UUIDFileName, b, 0644)
			if err != nil {
				log.Fatal("WriteFile", err, types.UUIDFileName)
			}
			log.Debugf("Wrote UUID %s\n", devUUID)
		}
		doWrite = true
		if hardwaremodel != "" {
			if oldHardwaremodel != hardwaremodel {
				log.Infof("Replacing existing hardwaremodel %s with %s\n",
					oldHardwaremodel, hardwaremodel)
			} else {
				log.Infof("No change to hardwaremodel %s\n",
					hardwaremodel)
				doWrite = false
			}
		} else {
			log.Infof("Got config with no hardwaremodel\n")
			doWrite = false
		}

		if doWrite {
			// Note that no CRLF
			b := []byte(hardwaremodel)
			err = ioutil.WriteFile(hardwaremodelFileName, b, 0644)
			if err != nil {
				log.Fatal("WriteFile", err,
					hardwaremodelFileName)
			}
			log.Debugf("Wrote hardwaremodel %s\n", hardwaremodel)
		}
		// We write the strings even if empty to make sure we have the most
		// recents. Since this is for debug use we are less careful
		// than for the hardwaremodel.
		b = []byte(enterprise) // Note that no CRLF
		err = ioutil.WriteFile(enterpriseFileName, b, 0644)
		if err != nil {
			log.Fatal("WriteFile", err, enterpriseFileName)
		}
		log.Debugf("Wrote enterprise %s\n", enterprise)
		b = []byte(name) // Note that no CRLF
		err = ioutil.WriteFile(nameFileName, b, 0644)
		if err != nil {
			log.Fatal("WriteFile", err, nameFileName)
		}
		log.Debugf("Wrote name %s\n", name)
	}

	err = pub.Publish("global", zedcloud.GetCloudMetrics())
	if err != nil {
		log.Errorln(err)
	}
}

// Post something without a return type.
// Returns true when done; false when retry
// the third return value is the extra send status, for Cert Miss status for example
func myPost(zedcloudCtx zedcloud.ZedCloudContext, tlsConfig *tls.Config,
	requrl string, retryCount int, reqlen int64, b *bytes.Buffer) (bool, *http.Response, types.SenderResult, []byte) {

	senderStatus := types.SenderStatusNone
	zedcloudCtx.TlsConfig = tlsConfig
	resp, contents, rtf, err := zedcloud.SendOnAllIntf(zedcloudCtx,
		requrl, reqlen, b, retryCount, return400)
	if err != nil {
		if rtf == types.SenderStatusRemTempFail {
			log.Errorf("remoteTemporaryFailure %s", err)
		} else if rtf == types.SenderStatusCertMiss {
			log.Infof("client myPost: Cert Miss\n")
		} else {
			log.Errorln(err)
		}
		return false, resp, rtf, contents
	}

	if !zedcloudCtx.NoLedManager {
		// Inform ledmanager about cloud connectivity
		utils.UpdateLedManagerConfig(3)
	}
	switch resp.StatusCode {
	case http.StatusOK:
		if !zedcloudCtx.NoLedManager {
			// Inform ledmanager about existence in cloud
			utils.UpdateLedManagerConfig(4)
		}
		log.Infof("%s StatusOK\n", requrl)
	case http.StatusCreated:
		if !zedcloudCtx.NoLedManager {
			// Inform ledmanager about existence in cloud
			utils.UpdateLedManagerConfig(4)
		}
		log.Infof("%s StatusCreated\n", requrl)
	case http.StatusConflict:
		if !zedcloudCtx.NoLedManager {
			// Inform ledmanager about brokenness
			utils.UpdateLedManagerConfig(10)
		}
		log.Errorf("%s StatusConflict\n", requrl)
		// Retry until fixed
		log.Errorf("%s\n", string(contents))
		return false, resp, senderStatus, contents
	case http.StatusNotFound, http.StatusUnauthorized, http.StatusNotModified:
		// Caller needs to handle
		return false, resp, senderStatus, contents
	default:
		log.Errorf("%s statuscode %d %s\n",
			requrl, resp.StatusCode,
			http.StatusText(resp.StatusCode))
		log.Errorf("%s\n", string(contents))
		return false, resp, senderStatus, contents
	}

	contentType := resp.Header.Get("Content-Type")
	if contentType == "" {
		log.Errorf("%s no content-type\n", requrl)
		return false, resp, senderStatus, contents
	}
	mimeType, _, err := mime.ParseMediaType(contentType)
	if err != nil {
		log.Errorf("%s ParseMediaType failed %v\n", requrl, err)
		return false, resp, senderStatus, contents
	}
	switch mimeType {
	case "application/x-proto-binary", "application/json", "text/plain":
		log.Debugf("Received reply %s\n", string(contents))
	default:
		log.Errorln("Incorrect Content-Type " + mimeType)
		return false, resp, senderStatus, contents
	}
	return true, resp, senderStatus, contents
}

// Returns true when done; false when retry
func selfRegister(zedcloudCtx zedcloud.ZedCloudContext, tlsConfig *tls.Config, deviceCertPem []byte, retryCount int) bool {
	// XXX add option to get this from a file in /config + override
	// logic
	productSerial := hardware.GetProductSerial()
	productSerial = strings.TrimSpace(productSerial)
	softSerial := hardware.GetSoftSerial()
	softSerial = strings.TrimSpace(softSerial)
	log.Infof("ProductSerial %s, SoftwareSerial %s\n", productSerial, softSerial)

	registerCreate := &register.ZRegisterMsg{
		PemCert:    []byte(base64.StdEncoding.EncodeToString(deviceCertPem)),
		Serial:     productSerial,
		SoftSerial: softSerial,
	}
	b, err := proto.Marshal(registerCreate)
	if err != nil {
		log.Errorln(err)
		return false
	}
	// in V2 API, register does not send UUID string
	requrl := zedcloud.URLPathString(serverNameAndPort, zedcloudCtx.V2API, false, nilUUID, "register")
	done, resp, _, contents := myPost(zedcloudCtx, tlsConfig,
		requrl, retryCount,
		int64(len(b)), bytes.NewBuffer(b))
	if resp != nil && resp.StatusCode == http.StatusNotModified {
		if !zedcloudCtx.NoLedManager {
			// Inform ledmanager about brokenness
			utils.UpdateLedManagerConfig(10)
		}
		log.Errorf("%s StatusNotModified\n", requrl)
		// Retry until fixed
		log.Errorf("%s\n", string(contents))
		done = false
	}

	return done
}

// fetch V2 certs from cloud, return GotCloudCerts and ServerIsV1 boolean
// if got certs, the leaf is saved to types.ServerCertFileName file
func fetchCertChain(zedcloudCtx zedcloud.ZedCloudContext, tlsConfig *tls.Config, retryCount int, force bool) bool {
	var resp *http.Response
	var b, contents []byte
	var done bool

	if !force {
		_, err := os.Stat(types.ServerCertFileName)
		if err == nil {
			return true
		}
	}

	// certs API is always V2, and without UUID, use http for now
	requrl := zedcloud.URLPathString(serverNameAndPort, true, true, nilUUID, "certs")
	// currently there is no data included for the request, same as myGet()
	done, resp, _, contents = myPost(zedcloudCtx, tlsConfig, requrl, retryCount, int64(len(b)), bytes.NewBuffer(b))
	if resp != nil {
		log.Infof("client fetchCertChain done %v, resp-code %d, content len %d", done, resp.StatusCode, len(contents))
		if resp.StatusCode == http.StatusNotFound || resp.StatusCode == http.StatusUnauthorized ||
			resp.StatusCode == http.StatusNotImplemented || resp.StatusCode == http.StatusBadRequest {
			// cloud server does not support V2 API
			log.Infof("client fetchCertChain: server %s does not support V2 API\n", serverName)
			return false
		}
		// catch default return status, if not done, will return false later
		log.Infof("client fetchCertChain: server %s return status %s, done %v\n", serverName, resp.Status, done)
	} else {
		log.Infof("client fetchCertChain done %v, resp null, content len %d", done, len(contents))
	}
	if !done {
		return false
	}

	zedcloudCtx.TlsConfig = tlsConfig
	// verify the certificate chain and write the siging cert to file
	certBytes, err := zedcloud.VerifyCloudCertChain(zedcloudCtx, serverName, contents)
	if err != nil {
		log.Errorf("client fetchCertChain: verify err %v", err)
		return false
	}

	err = fileutils.WriteRename(types.ServerCertFileName, certBytes)
	if err != nil {
		log.Errorf("client fetchCertChain: file save err %v", err)
		return false
	}

	log.Infof("client fetchCertChain: ok\n")
	return true
}

func doGetUUID(ctx *clientContext, tlsConfig *tls.Config,
	retryCount int) (bool, uuid.UUID, string, string, string) {

	var resp *http.Response
	var contents []byte
	var rtf types.SenderResult
	zedcloudCtx := *ctx.zedcloudCtx

	// get UUID does not have UUID string in V2 API
	requrl := zedcloud.URLPathString(serverNameAndPort, zedcloudCtx.V2API, false, nilUUID, "config")
	b, err := generateConfigRequest()
	if err != nil {
		log.Errorln(err)
		return false, nilUUID, "", "", ""
	}
	var done bool
	done, resp, rtf, contents = myPost(zedcloudCtx, tlsConfig, requrl, retryCount,
		int64(len(b)), bytes.NewBuffer(b))
	if resp != nil && resp.StatusCode == http.StatusNotModified {
		// Acceptable response for a ConfigRequest POST
		done = true
	}
	if !done {
		// This may be due to the cloud cert file is stale, since the hash does not match.
		// acquire new cert chain.
		if rtf == types.SenderStatusCertMiss {
			interval := time.Duration(1)
			ctx.getCertsTimer = time.NewTimer(interval * time.Second)
			log.Infof("doGetUUID: Cert miss. Setup timer to acquire\n")
		}
		return false, nilUUID, "", "", ""
	}
	log.Infof("doGetUUID: client getUUID ok\n")
	devUUID, hardwaremodel, enterprise, name, err := parseConfig(requrl, resp, contents)
	if err == nil {
		// Inform ledmanager about config received from cloud
		if !zedcloudCtx.NoLedManager {
			utils.UpdateLedManagerConfig(4)
		}
		return true, devUUID, hardwaremodel, enterprise, name
	}
	// Keep on trying until it parses
	log.Errorf("Failed parsing uuid: %s\n", err)
	return false, nilUUID, "", "", ""
}

// Handles both create and modify events
func handleGlobalConfigModify(ctxArg interface{}, key string,
	statusArg interface{}) {

	ctx := ctxArg.(*clientContext)
	if key != "global" {
		log.Debugf("handleGlobalConfigModify: ignoring %s\n", key)
		return
	}
	log.Infof("handleGlobalConfigModify for %s\n", key)
	var gcp *types.GlobalConfig
	debug, gcp = agentlog.HandleGlobalConfig(ctx.subGlobalConfig, agentName,
		debugOverride)
	if gcp != nil {
		ctx.globalConfig = gcp
	}
	log.Infof("handleGlobalConfigModify done for %s\n", key)
}

func handleGlobalConfigDelete(ctxArg interface{}, key string,
	statusArg interface{}) {

	ctx := ctxArg.(*clientContext)
	if key != "global" {
		log.Debugf("handleGlobalConfigDelete: ignoring %s\n", key)
		return
	}
	log.Infof("handleGlobalConfigDelete for %s\n", key)
	debug, _ = agentlog.HandleGlobalConfig(ctx.subGlobalConfig, agentName,
		debugOverride)
	*ctx.globalConfig = types.GlobalConfigDefaults
	log.Infof("handleGlobalConfigDelete done for %s\n", key)
}

// Handles both create and modify events
func handleDNSModify(ctxArg interface{}, key string, statusArg interface{}) {

	status := statusArg.(types.DeviceNetworkStatus)
	ctx := ctxArg.(*clientContext)
	if key != "global" {
		log.Infof("handleDNSModify: ignoring %s\n", key)
		return
	}
	log.Infof("handleDNSModify for %s\n", key)
	if cmp.Equal(ctx.deviceNetworkStatus, status) {
		log.Infof("handleDNSModify no change\n")
		return
	}

	log.Infof("handleDNSModify: changed %v",
		cmp.Diff(ctx.deviceNetworkStatus, status))
	*ctx.deviceNetworkStatus = status
	newAddrCount := types.CountLocalAddrAnyNoLinkLocal(*ctx.deviceNetworkStatus)
	if newAddrCount != ctx.usableAddressCount {
		log.Infof("DeviceNetworkStatus from %d to %d addresses\n",
			ctx.usableAddressCount, newAddrCount)
		// ledmanager subscribes to DeviceNetworkStatus to see changes
		ctx.usableAddressCount = newAddrCount
	}

	// update proxy certs if configured
	ctx.zedcloudCtx.DeviceNetworkStatus = &status
	// if there is proxy certs change, needs to update both
	// onboard and device tlsconfig
	cloudCtx := zedcloud.ZedCloudContext{}
	cloudCtx.TlsConfig = devtlsConfig
	cloudCtx.PrevCertPEM = ctx.zedcloudCtx.PrevCertPEM
	cloudCtx.DeviceNetworkStatus = ctx.zedcloudCtx.DeviceNetworkStatus
	if cloudCtx.TlsConfig == nil {
		log.Infof("handleDNSModify: client onboard, tlsconfig nil\n")
	}
	updated := zedcloud.UpdateTLSProxyCerts(&cloudCtx)
	if updated {
		if onboardTLSConfig != nil {
			onboardTLSConfig.RootCAs = cloudCtx.TlsConfig.RootCAs
		}
		devtlsConfig.RootCAs = cloudCtx.TlsConfig.RootCAs
		log.Infof("handleDNSModify: client rootCAs updated\n")
	}

	log.Infof("handleDNSModify done for %s\n", key)
}

func handleDNSDelete(ctxArg interface{}, key string,
	statusArg interface{}) {

	log.Infof("handleDNSDelete for %s\n", key)
	ctx := ctxArg.(*clientContext)

	if key != "global" {
		log.Infof("handleDNSDelete: ignoring %s\n", key)
		return
	}
	*ctx.deviceNetworkStatus = types.DeviceNetworkStatus{}
	newAddrCount := types.CountLocalAddrAnyNoLinkLocal(*ctx.deviceNetworkStatus)
	ctx.usableAddressCount = newAddrCount
	log.Infof("handleDNSDelete done for %s\n", key)
}
