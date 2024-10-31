// Copyright (c) 2024 Zededa, Inc.
// SPDX-License-Identifier: Apache-2.0

package monitor

import (
	"github.com/lf-edge/eve/pkg/pillar/types"
	uuid "github.com/satori/go.uuid"
)

type nodeStatus struct {
	Server         string                   `json:"server,omitempty"`
	NodeUUID       uuid.UUID                `json:"node_uuid,omitempty"`
	Onboarded      bool                     `json:"onboarded"`
	AppSummary     types.AppInstanceSummary `json:"app_summary,omitempty"`
	ZedAgentStatus types.ZedAgentStatus     `json:"zedagent_status,omitempty"`
}

type appInstancesStatus struct {
	Apps []types.AppInstanceStatus `json:"apps"`
}

func (ctx *monitorContext) isOnboarded() (bool, uuid.UUID) {
	sub := ctx.subscriptions["OnboardingStatus"]
	if item, err := sub.Get("global"); err == nil {
		onboardingStatus := item.(types.OnboardingStatus)
		return true, onboardingStatus.DeviceUUID
	}
	return false, uuid.UUID{}
}

func (ctx *monitorContext) getAppSummary() types.AppInstanceSummary {
	// send the network status to the server
	sub := ctx.subscriptions["AppSummary"]
	if item, err := sub.Get("global"); err == nil {
		appSummary := item.(types.AppInstanceSummary)
		return appSummary
	}
	return types.AppInstanceSummary{}
}

func (ctx *monitorContext) sendNodeStatus() {
	// send the node status to the server
	nodeStatus := nodeStatus{
		Server: ctx.serverNameAndPort,
	}
	if onboarded, nodeUUID := ctx.isOnboarded(); onboarded {
		nodeStatus.NodeUUID = nodeUUID
		nodeStatus.Onboarded = true
	}

	nodeStatus.ZedAgentStatus = ctx.getZedAgentStatus()
	nodeStatus.AppSummary = ctx.getAppSummary()

	if ctx.lastNodeStatus != nil && *ctx.lastNodeStatus == nodeStatus {
		return
	}
	ctx.lastNodeStatus = &nodeStatus

	ctx.IPCServer.sendIpcMessage("NodeStatus", nodeStatus)
}

func (ctx *monitorContext) getAppInstancesStatus() []types.AppInstanceStatus {
	// send the network status to the server
	sub := ctx.subscriptions["AppStatus"]
	items := sub.GetAll()
	apps := make([]types.AppInstanceStatus, 0)
	for _, item := range items {
		appSummary := item.(types.AppInstanceStatus)
		apps = append(apps, appSummary)
	}
	return apps
}

func (ctx *monitorContext) getZedAgentStatus() types.ZedAgentStatus {
	sub := ctx.subscriptions["ZedAgentStatus"]
	if item, err := sub.Get("global"); err == nil {
		zedAgentStatus := item.(types.ZedAgentStatus)
		return zedAgentStatus
	}
	return types.ZedAgentStatus{}
}

func (ctx *monitorContext) sendAppsList() {
	// send the node status to the server
	appStatus := ctx.getAppInstancesStatus()
	if len(appStatus) > 0 {
		apps := appInstancesStatus{
			Apps: appStatus,
		}
		ctx.IPCServer.sendIpcMessage("AppsList", apps)
	}
}
