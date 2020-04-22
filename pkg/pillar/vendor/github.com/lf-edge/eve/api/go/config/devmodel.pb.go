// Copyright(c) 2017-2018 Zededa, Inc.
// All rights reserved.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.21.0-devel
// 	protoc        v3.6.1
// source: config/devmodel.proto

package config

import (
	proto "github.com/golang/protobuf/proto"
	common "github.com/lf-edge/eve/api/go/common"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

// Deprecate; replace by level 2 specification
type SWAdapterType int32

const (
	SWAdapterType_IGNORE SWAdapterType = 0
	SWAdapterType_VLAN   SWAdapterType = 1
	SWAdapterType_BOND   SWAdapterType = 2
)

// Enum value maps for SWAdapterType.
var (
	SWAdapterType_name = map[int32]string{
		0: "IGNORE",
		1: "VLAN",
		2: "BOND",
	}
	SWAdapterType_value = map[string]int32{
		"IGNORE": 0,
		"VLAN":   1,
		"BOND":   2,
	}
)

func (x SWAdapterType) Enum() *SWAdapterType {
	p := new(SWAdapterType)
	*p = x
	return p
}

func (x SWAdapterType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (SWAdapterType) Descriptor() protoreflect.EnumDescriptor {
	return file_config_devmodel_proto_enumTypes[0].Descriptor()
}

func (SWAdapterType) Type() protoreflect.EnumType {
	return &file_config_devmodel_proto_enumTypes[0]
}

func (x SWAdapterType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use SWAdapterType.Descriptor instead.
func (SWAdapterType) EnumDescriptor() ([]byte, []int) {
	return file_config_devmodel_proto_rawDescGZIP(), []int{0}
}

// Deprecate; replace by level 2 specification
type SWAdapterParams struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AType SWAdapterType `protobuf:"varint,1,opt,name=aType,proto3,enum=SWAdapterType" json:"aType,omitempty"`
	// vlan
	UnderlayInterface string `protobuf:"bytes,8,opt,name=underlayInterface,proto3" json:"underlayInterface,omitempty"`
	VlanId            uint32 `protobuf:"varint,9,opt,name=vlanId,proto3" json:"vlanId,omitempty"`
	// OR : repeated physical interfaces for bond0
	Bondgroup []string `protobuf:"bytes,10,rep,name=bondgroup,proto3" json:"bondgroup,omitempty"`
}

func (x *SWAdapterParams) Reset() {
	*x = SWAdapterParams{}
	if protoimpl.UnsafeEnabled {
		mi := &file_config_devmodel_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SWAdapterParams) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SWAdapterParams) ProtoMessage() {}

func (x *SWAdapterParams) ProtoReflect() protoreflect.Message {
	mi := &file_config_devmodel_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SWAdapterParams.ProtoReflect.Descriptor instead.
func (*SWAdapterParams) Descriptor() ([]byte, []int) {
	return file_config_devmodel_proto_rawDescGZIP(), []int{0}
}

func (x *SWAdapterParams) GetAType() SWAdapterType {
	if x != nil {
		return x.AType
	}
	return SWAdapterType_IGNORE
}

func (x *SWAdapterParams) GetUnderlayInterface() string {
	if x != nil {
		return x.UnderlayInterface
	}
	return ""
}

func (x *SWAdapterParams) GetVlanId() uint32 {
	if x != nil {
		return x.VlanId
	}
	return 0
}

func (x *SWAdapterParams) GetBondgroup() []string {
	if x != nil {
		return x.Bondgroup
	}
	return nil
}

// systemAdapters, are the higher l2 concept built on physicalIOs.
// systemAdapters, gives all the required bits to turn the physical IOs
// into useful IP endpoints.
// These endpoints can be further used to connect to controller or
// can be shared between workload/services running on the node.
type SystemAdapter struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// name - Name of the Network Interface. This is the Port Name
	//  used in Info / Metrics / flowlog etc. Name cannot be changed.
	// This will be the Network Port name.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// deprecated; need level 2 spec. sWAdapterParams allocDetails = 20;
	// this is part of the freelink group
	// DEPRECATED by PhyIoAdapter.usagePolicy.
	FreeUplink bool `protobuf:"varint,2,opt,name=freeUplink,proto3" json:"freeUplink,omitempty"`
	// uplink - DEPRECATED by PhysicalIO.Usage / PhysicalIO.UsagePolicy
	// this is part of the uplink group
	// deprecate: have a separate device policy object in the API
	Uplink bool `protobuf:"varint,3,opt,name=uplink,proto3" json:"uplink,omitempty"`
	// networkUUID - attach this network config for this adapter
	// if not set, depending on Usage of Adapter, would be treated as
	// an L2 port
	NetworkUUID string `protobuf:"bytes,4,opt,name=networkUUID,proto3" json:"networkUUID,omitempty"`
	// addr - if its static network we need ip address
	// If this is specified, networkUUID must also be specified. addr
	// is expected to be in sync with the network object (same subnet etc ).
	Addr string `protobuf:"bytes,5,opt,name=addr,proto3" json:"addr,omitempty"`
	// alias - Device just reflects it back in status / Metrics back to
	// cloud.
	Alias string `protobuf:"bytes,7,opt,name=alias,proto3" json:"alias,omitempty"`
	// lowerLayerName - For example, if lower layer is PhysicalAdapter
	// ( physical interface), this should point to PhyLabel of the
	// physicalIO.
	LowerLayerName string `protobuf:"bytes,8,opt,name=lowerLayerName,proto3" json:"lowerLayerName,omitempty"`
}

func (x *SystemAdapter) Reset() {
	*x = SystemAdapter{}
	if protoimpl.UnsafeEnabled {
		mi := &file_config_devmodel_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SystemAdapter) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SystemAdapter) ProtoMessage() {}

func (x *SystemAdapter) ProtoReflect() protoreflect.Message {
	mi := &file_config_devmodel_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SystemAdapter.ProtoReflect.Descriptor instead.
func (*SystemAdapter) Descriptor() ([]byte, []int) {
	return file_config_devmodel_proto_rawDescGZIP(), []int{1}
}

func (x *SystemAdapter) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *SystemAdapter) GetFreeUplink() bool {
	if x != nil {
		return x.FreeUplink
	}
	return false
}

func (x *SystemAdapter) GetUplink() bool {
	if x != nil {
		return x.Uplink
	}
	return false
}

func (x *SystemAdapter) GetNetworkUUID() string {
	if x != nil {
		return x.NetworkUUID
	}
	return ""
}

func (x *SystemAdapter) GetAddr() string {
	if x != nil {
		return x.Addr
	}
	return ""
}

func (x *SystemAdapter) GetAlias() string {
	if x != nil {
		return x.Alias
	}
	return ""
}

func (x *SystemAdapter) GetLowerLayerName() string {
	if x != nil {
		return x.LowerLayerName
	}
	return ""
}

// Given additional details for EVE softwar to how to treat this
// interface. Example policies could be limit use of LTE interface
// or only use Eth1 only if Eth0 is not available etc
// XXX Note that this is the static information from the model.
// Current configuration is in systemAdapter
type PhyIOUsagePolicy struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// DEPRECATED - Used only when one other normal uplinks are down
	FreeUplink bool `protobuf:"varint,1,opt,name=freeUplink,proto3" json:"freeUplink,omitempty"`
	// fallBackPriority
	//  0 is the highest priority.
	//  Lower priority interfaces are used only when NONE of the higher
	//  priority interfaces are up.
	//  For example:
	//      First use all interfaces with priority 0
	//      if no priority 0 interfaces, use interfaces with priority 1
	//      if no priority 1 interfaces, use interfaces with priority 2
	//      and so on..
	FallBackPriority uint32 `protobuf:"varint,2,opt,name=fallBackPriority,proto3" json:"fallBackPriority,omitempty"`
}

func (x *PhyIOUsagePolicy) Reset() {
	*x = PhyIOUsagePolicy{}
	if protoimpl.UnsafeEnabled {
		mi := &file_config_devmodel_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PhyIOUsagePolicy) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PhyIOUsagePolicy) ProtoMessage() {}

func (x *PhyIOUsagePolicy) ProtoReflect() protoreflect.Message {
	mi := &file_config_devmodel_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PhyIOUsagePolicy.ProtoReflect.Descriptor instead.
func (*PhyIOUsagePolicy) Descriptor() ([]byte, []int) {
	return file_config_devmodel_proto_rawDescGZIP(), []int{2}
}

func (x *PhyIOUsagePolicy) GetFreeUplink() bool {
	if x != nil {
		return x.FreeUplink
	}
	return false
}

func (x *PhyIOUsagePolicy) GetFallBackPriority() uint32 {
	if x != nil {
		return x.FallBackPriority
	}
	return 0
}

// PhysicalIO:
//    Absolute low level description of physical buses and ports that are
//    available on given platfrom.
//    Collection of these IOs, connstitue what we would call as hardware
//    model. Each physical IO is manageable and visible to EVE software, and
//    it can be further configured to either provide IP connectivity or
//    directly be given to workloads
type PhysicalIO struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ptype common.PhyIoType `protobuf:"varint,1,opt,name=ptype,proto3,enum=PhyIoType" json:"ptype,omitempty"`
	// physical label typically printed on box.
	// Example Eth0, Eth1, Wifi0, ComA, ComB
	Phylabel string `protobuf:"bytes,2,opt,name=phylabel,proto3" json:"phylabel,omitempty"`
	// The hardware bus address. The key to this map can be of the following
	// (case-insensitive) values:
	// "pcilong": the address is a PCI id of the form 0000:02:00.0
	// "ifname": the addresss is a string for a network interface like "eth1"
	// "serial": the address is a Linux serial port alias such as "/dev/ttyS2"
	// "irq": the address is a number such as "5". This can be a comma
	//    separated list of integers or even a range of integers. Hence using
	//    a string to address this.
	// "ioports": the address is a string such as "2f8-2ff"
	// If the type is PhyIoNet*, then there needs to be an "ifname" physaddr.
	Phyaddrs map[string]string `protobuf:"bytes,3,rep,name=phyaddrs,proto3" json:"phyaddrs,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// logicallabel - provides the ability to model designer to refer
	//    the physicalIO port to using more friendly name
	// For example Eth0->Mgmt0
	//  or USBA->ConfigDiskA etc
	Logicallabel string `protobuf:"bytes,4,opt,name=logicallabel,proto3" json:"logicallabel,omitempty"`
	// assigngrp
	// Assignment Group, is unique label that is applied across PhysicalIOs
	// EntireGroup can be assigned to application or nothing at all
	//
	// This is the name used in AppInstanceConfig.adapters to assign an
	// adapter to an application.
	//
	// If assigngrp is not set, the Adapter cannot be assigned to any
	// application. One example is, when the adapter is on the same Pci
	// bus as another device required by Dom0.
	//
	// Even if there is only one device on the its PCIBus, the assignGrp Must
	// be set.
	Assigngrp string `protobuf:"bytes,5,opt,name=assigngrp,proto3" json:"assigngrp,omitempty"`
	// usage - indicates the role of adapter ( mgmt / blocked / app-direct
	//    etc. )
	Usage common.PhyIoMemberUsage `protobuf:"varint,6,opt,name=usage,proto3,enum=PhyIoMemberUsage" json:"usage,omitempty"`
	// usagePolicy - Policy Object used to further refine the usage.
	// For example, specify if this should be only used as fallback?
	//    Or used as the primary uplink? Allow App traffic? restrict
	//    app traffic?? etc..
	UsagePolicy *PhyIOUsagePolicy `protobuf:"bytes,7,opt,name=usagePolicy,proto3" json:"usagePolicy,omitempty"`
	// physical and logical attributes
	//    For example in WWAN to which firmware version to laod etc
	Cbattr map[string]string `protobuf:"bytes,8,rep,name=cbattr,proto3" json:"cbattr,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *PhysicalIO) Reset() {
	*x = PhysicalIO{}
	if protoimpl.UnsafeEnabled {
		mi := &file_config_devmodel_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PhysicalIO) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PhysicalIO) ProtoMessage() {}

func (x *PhysicalIO) ProtoReflect() protoreflect.Message {
	mi := &file_config_devmodel_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PhysicalIO.ProtoReflect.Descriptor instead.
func (*PhysicalIO) Descriptor() ([]byte, []int) {
	return file_config_devmodel_proto_rawDescGZIP(), []int{3}
}

func (x *PhysicalIO) GetPtype() common.PhyIoType {
	if x != nil {
		return x.Ptype
	}
	return common.PhyIoType_PhyIoNoop
}

func (x *PhysicalIO) GetPhylabel() string {
	if x != nil {
		return x.Phylabel
	}
	return ""
}

func (x *PhysicalIO) GetPhyaddrs() map[string]string {
	if x != nil {
		return x.Phyaddrs
	}
	return nil
}

func (x *PhysicalIO) GetLogicallabel() string {
	if x != nil {
		return x.Logicallabel
	}
	return ""
}

func (x *PhysicalIO) GetAssigngrp() string {
	if x != nil {
		return x.Assigngrp
	}
	return ""
}

func (x *PhysicalIO) GetUsage() common.PhyIoMemberUsage {
	if x != nil {
		return x.Usage
	}
	return common.PhyIoMemberUsage_PhyIoUsageNone
}

func (x *PhysicalIO) GetUsagePolicy() *PhyIOUsagePolicy {
	if x != nil {
		return x.UsagePolicy
	}
	return nil
}

func (x *PhysicalIO) GetCbattr() map[string]string {
	if x != nil {
		return x.Cbattr
	}
	return nil
}

var File_config_devmodel_proto protoreflect.FileDescriptor

var file_config_devmodel_proto_rawDesc = []byte{
	0x0a, 0x15, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2f, 0x64, 0x65, 0x76, 0x6d, 0x6f, 0x64, 0x65,
	0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f,
	0x64, 0x65, 0x76, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0x9b, 0x01, 0x0a, 0x0f, 0x73, 0x57, 0x41, 0x64, 0x61, 0x70, 0x74,
	0x65, 0x72, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x12, 0x24, 0x0a, 0x05, 0x61, 0x54, 0x79, 0x70,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0e, 0x2e, 0x73, 0x57, 0x41, 0x64, 0x61, 0x70,
	0x74, 0x65, 0x72, 0x54, 0x79, 0x70, 0x65, 0x52, 0x05, 0x61, 0x54, 0x79, 0x70, 0x65, 0x12, 0x2c,
	0x0a, 0x11, 0x75, 0x6e, 0x64, 0x65, 0x72, 0x6c, 0x61, 0x79, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x66,
	0x61, 0x63, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x11, 0x75, 0x6e, 0x64, 0x65, 0x72,
	0x6c, 0x61, 0x79, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x66, 0x61, 0x63, 0x65, 0x12, 0x16, 0x0a, 0x06,
	0x76, 0x6c, 0x61, 0x6e, 0x49, 0x64, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x06, 0x76, 0x6c,
	0x61, 0x6e, 0x49, 0x64, 0x12, 0x1c, 0x0a, 0x09, 0x62, 0x6f, 0x6e, 0x64, 0x67, 0x72, 0x6f, 0x75,
	0x70, 0x18, 0x0a, 0x20, 0x03, 0x28, 0x09, 0x52, 0x09, 0x62, 0x6f, 0x6e, 0x64, 0x67, 0x72, 0x6f,
	0x75, 0x70, 0x22, 0xcf, 0x01, 0x0a, 0x0d, 0x53, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x41, 0x64, 0x61,
	0x70, 0x74, 0x65, 0x72, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x66, 0x72, 0x65, 0x65,
	0x55, 0x70, 0x6c, 0x69, 0x6e, 0x6b, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0a, 0x66, 0x72,
	0x65, 0x65, 0x55, 0x70, 0x6c, 0x69, 0x6e, 0x6b, 0x12, 0x16, 0x0a, 0x06, 0x75, 0x70, 0x6c, 0x69,
	0x6e, 0x6b, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x06, 0x75, 0x70, 0x6c, 0x69, 0x6e, 0x6b,
	0x12, 0x20, 0x0a, 0x0b, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x55, 0x55, 0x49, 0x44, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x55, 0x55,
	0x49, 0x44, 0x12, 0x12, 0x0a, 0x04, 0x61, 0x64, 0x64, 0x72, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x61, 0x64, 0x64, 0x72, 0x12, 0x14, 0x0a, 0x05, 0x61, 0x6c, 0x69, 0x61, 0x73, 0x18,
	0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x61, 0x6c, 0x69, 0x61, 0x73, 0x12, 0x26, 0x0a, 0x0e,
	0x6c, 0x6f, 0x77, 0x65, 0x72, 0x4c, 0x61, 0x79, 0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x08,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x6c, 0x6f, 0x77, 0x65, 0x72, 0x4c, 0x61, 0x79, 0x65, 0x72,
	0x4e, 0x61, 0x6d, 0x65, 0x22, 0x5e, 0x0a, 0x10, 0x50, 0x68, 0x79, 0x49, 0x4f, 0x55, 0x73, 0x61,
	0x67, 0x65, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x12, 0x1e, 0x0a, 0x0a, 0x66, 0x72, 0x65, 0x65,
	0x55, 0x70, 0x6c, 0x69, 0x6e, 0x6b, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0a, 0x66, 0x72,
	0x65, 0x65, 0x55, 0x70, 0x6c, 0x69, 0x6e, 0x6b, 0x12, 0x2a, 0x0a, 0x10, 0x66, 0x61, 0x6c, 0x6c,
	0x42, 0x61, 0x63, 0x6b, 0x50, 0x72, 0x69, 0x6f, 0x72, 0x69, 0x74, 0x79, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0d, 0x52, 0x10, 0x66, 0x61, 0x6c, 0x6c, 0x42, 0x61, 0x63, 0x6b, 0x50, 0x72, 0x69, 0x6f,
	0x72, 0x69, 0x74, 0x79, 0x22, 0xca, 0x03, 0x0a, 0x0a, 0x50, 0x68, 0x79, 0x73, 0x69, 0x63, 0x61,
	0x6c, 0x49, 0x4f, 0x12, 0x20, 0x0a, 0x05, 0x70, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0e, 0x32, 0x0a, 0x2e, 0x50, 0x68, 0x79, 0x49, 0x6f, 0x54, 0x79, 0x70, 0x65, 0x52, 0x05,
	0x70, 0x74, 0x79, 0x70, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x68, 0x79, 0x6c, 0x61, 0x62, 0x65,
	0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x68, 0x79, 0x6c, 0x61, 0x62, 0x65,
	0x6c, 0x12, 0x35, 0x0a, 0x08, 0x70, 0x68, 0x79, 0x61, 0x64, 0x64, 0x72, 0x73, 0x18, 0x03, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x50, 0x68, 0x79, 0x73, 0x69, 0x63, 0x61, 0x6c, 0x49, 0x4f,
	0x2e, 0x50, 0x68, 0x79, 0x61, 0x64, 0x64, 0x72, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x08,
	0x70, 0x68, 0x79, 0x61, 0x64, 0x64, 0x72, 0x73, 0x12, 0x22, 0x0a, 0x0c, 0x6c, 0x6f, 0x67, 0x69,
	0x63, 0x61, 0x6c, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c,
	0x6c, 0x6f, 0x67, 0x69, 0x63, 0x61, 0x6c, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x12, 0x1c, 0x0a, 0x09,
	0x61, 0x73, 0x73, 0x69, 0x67, 0x6e, 0x67, 0x72, 0x70, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x09, 0x61, 0x73, 0x73, 0x69, 0x67, 0x6e, 0x67, 0x72, 0x70, 0x12, 0x27, 0x0a, 0x05, 0x75, 0x73,
	0x61, 0x67, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x11, 0x2e, 0x50, 0x68, 0x79, 0x49,
	0x6f, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x55, 0x73, 0x61, 0x67, 0x65, 0x52, 0x05, 0x75, 0x73,
	0x61, 0x67, 0x65, 0x12, 0x33, 0x0a, 0x0b, 0x75, 0x73, 0x61, 0x67, 0x65, 0x50, 0x6f, 0x6c, 0x69,
	0x63, 0x79, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x50, 0x68, 0x79, 0x49, 0x4f,
	0x55, 0x73, 0x61, 0x67, 0x65, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x52, 0x0b, 0x75, 0x73, 0x61,
	0x67, 0x65, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x12, 0x2f, 0x0a, 0x06, 0x63, 0x62, 0x61, 0x74,
	0x74, 0x72, 0x18, 0x08, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x50, 0x68, 0x79, 0x73, 0x69,
	0x63, 0x61, 0x6c, 0x49, 0x4f, 0x2e, 0x43, 0x62, 0x61, 0x74, 0x74, 0x72, 0x45, 0x6e, 0x74, 0x72,
	0x79, 0x52, 0x06, 0x63, 0x62, 0x61, 0x74, 0x74, 0x72, 0x1a, 0x3b, 0x0a, 0x0d, 0x50, 0x68, 0x79,
	0x61, 0x64, 0x64, 0x72, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65,
	0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c,
	0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x1a, 0x39, 0x0a, 0x0b, 0x43, 0x62, 0x61, 0x74, 0x74, 0x72,
	0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38,
	0x01, 0x2a, 0x2f, 0x0a, 0x0d, 0x73, 0x57, 0x41, 0x64, 0x61, 0x70, 0x74, 0x65, 0x72, 0x54, 0x79,
	0x70, 0x65, 0x12, 0x0a, 0x0a, 0x06, 0x49, 0x47, 0x4e, 0x4f, 0x52, 0x45, 0x10, 0x00, 0x12, 0x08,
	0x0a, 0x04, 0x56, 0x4c, 0x41, 0x4e, 0x10, 0x01, 0x12, 0x08, 0x0a, 0x04, 0x42, 0x4f, 0x4e, 0x44,
	0x10, 0x02, 0x42, 0x47, 0x0a, 0x1f, 0x63, 0x6f, 0x6d, 0x2e, 0x7a, 0x65, 0x64, 0x65, 0x64, 0x61,
	0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x5a, 0x24, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x6c, 0x66, 0x2d, 0x65, 0x64, 0x67, 0x65, 0x2f, 0x65, 0x76, 0x65, 0x2f, 0x61, 0x70,
	0x69, 0x2f, 0x67, 0x6f, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_config_devmodel_proto_rawDescOnce sync.Once
	file_config_devmodel_proto_rawDescData = file_config_devmodel_proto_rawDesc
)

func file_config_devmodel_proto_rawDescGZIP() []byte {
	file_config_devmodel_proto_rawDescOnce.Do(func() {
		file_config_devmodel_proto_rawDescData = protoimpl.X.CompressGZIP(file_config_devmodel_proto_rawDescData)
	})
	return file_config_devmodel_proto_rawDescData
}

var file_config_devmodel_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_config_devmodel_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_config_devmodel_proto_goTypes = []interface{}{
	(SWAdapterType)(0),           // 0: sWAdapterType
	(*SWAdapterParams)(nil),      // 1: sWAdapterParams
	(*SystemAdapter)(nil),        // 2: SystemAdapter
	(*PhyIOUsagePolicy)(nil),     // 3: PhyIOUsagePolicy
	(*PhysicalIO)(nil),           // 4: PhysicalIO
	nil,                          // 5: PhysicalIO.PhyaddrsEntry
	nil,                          // 6: PhysicalIO.CbattrEntry
	(common.PhyIoType)(0),        // 7: PhyIoType
	(common.PhyIoMemberUsage)(0), // 8: PhyIoMemberUsage
}
var file_config_devmodel_proto_depIdxs = []int32{
	0, // 0: sWAdapterParams.aType:type_name -> sWAdapterType
	7, // 1: PhysicalIO.ptype:type_name -> PhyIoType
	5, // 2: PhysicalIO.phyaddrs:type_name -> PhysicalIO.PhyaddrsEntry
	8, // 3: PhysicalIO.usage:type_name -> PhyIoMemberUsage
	3, // 4: PhysicalIO.usagePolicy:type_name -> PhyIOUsagePolicy
	6, // 5: PhysicalIO.cbattr:type_name -> PhysicalIO.CbattrEntry
	6, // [6:6] is the sub-list for method output_type
	6, // [6:6] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_config_devmodel_proto_init() }
func file_config_devmodel_proto_init() {
	if File_config_devmodel_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_config_devmodel_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SWAdapterParams); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_config_devmodel_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SystemAdapter); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_config_devmodel_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PhyIOUsagePolicy); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_config_devmodel_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PhysicalIO); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_config_devmodel_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_config_devmodel_proto_goTypes,
		DependencyIndexes: file_config_devmodel_proto_depIdxs,
		EnumInfos:         file_config_devmodel_proto_enumTypes,
		MessageInfos:      file_config_devmodel_proto_msgTypes,
	}.Build()
	File_config_devmodel_proto = out.File
	file_config_devmodel_proto_rawDesc = nil
	file_config_devmodel_proto_goTypes = nil
	file_config_devmodel_proto_depIdxs = nil
}
