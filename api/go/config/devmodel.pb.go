// Code generated by protoc-gen-go. DO NOT EDIT.
// source: devmodel.proto

package config

import (
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

// Deprecate; replace by level 2 specification
type SWAdapterType int32

const (
	SWAdapterType_IGNORE SWAdapterType = 0
	SWAdapterType_VLAN   SWAdapterType = 1
	SWAdapterType_BOND   SWAdapterType = 2
)

var SWAdapterType_name = map[int32]string{
	0: "IGNORE",
	1: "VLAN",
	2: "BOND",
}

var SWAdapterType_value = map[string]int32{
	"IGNORE": 0,
	"VLAN":   1,
	"BOND":   2,
}

func (x SWAdapterType) String() string {
	return proto.EnumName(SWAdapterType_name, int32(x))
}

func (SWAdapterType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_9fb58492383773ea, []int{0}
}

type PhyIoType int32

const (
	PhyIoType_PhyIoNoop    PhyIoType = 0
	PhyIoType_PhyIoNetEth  PhyIoType = 1
	PhyIoType_PhyIoUSB     PhyIoType = 2
	PhyIoType_PhyIoCOM     PhyIoType = 3
	PhyIoType_PhyIoAudio   PhyIoType = 4
	PhyIoType_PhyIoNetWLAN PhyIoType = 5
	PhyIoType_PhyIoNetWWAN PhyIoType = 6
	PhyIoType_PhyIoHDMI    PhyIoType = 7
	PhyIoType_PhyIoOther   PhyIoType = 255
)

var PhyIoType_name = map[int32]string{
	0:   "PhyIoNoop",
	1:   "PhyIoNetEth",
	2:   "PhyIoUSB",
	3:   "PhyIoCOM",
	4:   "PhyIoAudio",
	5:   "PhyIoNetWLAN",
	6:   "PhyIoNetWWAN",
	7:   "PhyIoHDMI",
	255: "PhyIoOther",
}

var PhyIoType_value = map[string]int32{
	"PhyIoNoop":    0,
	"PhyIoNetEth":  1,
	"PhyIoUSB":     2,
	"PhyIoCOM":     3,
	"PhyIoAudio":   4,
	"PhyIoNetWLAN": 5,
	"PhyIoNetWWAN": 6,
	"PhyIoHDMI":    7,
	"PhyIoOther":   255,
}

func (x PhyIoType) String() string {
	return proto.EnumName(PhyIoType_name, int32(x))
}

func (PhyIoType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_9fb58492383773ea, []int{1}
}

// How does EVE should use them, for what purpose it is for
type PhyIoMemberUsage int32

const (
	PhyIoMemberUsage_PhyIoUsageNone      PhyIoMemberUsage = 0
	PhyIoMemberUsage_PhyIoUsageMgmt      PhyIoMemberUsage = 1
	PhyIoMemberUsage_PhyIoUsageShared    PhyIoMemberUsage = 2
	PhyIoMemberUsage_PhyIoUsageDedicated PhyIoMemberUsage = 3
)

var PhyIoMemberUsage_name = map[int32]string{
	0: "PhyIoUsageNone",
	1: "PhyIoUsageMgmt",
	2: "PhyIoUsageShared",
	3: "PhyIoUsageDedicated",
}

var PhyIoMemberUsage_value = map[string]int32{
	"PhyIoUsageNone":      0,
	"PhyIoUsageMgmt":      1,
	"PhyIoUsageShared":    2,
	"PhyIoUsageDedicated": 3,
}

func (x PhyIoMemberUsage) String() string {
	return proto.EnumName(PhyIoMemberUsage_name, int32(x))
}

func (PhyIoMemberUsage) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_9fb58492383773ea, []int{2}
}

// Deprecate; replace by level 2 specification
type SWAdapterParams struct {
	AType SWAdapterType `protobuf:"varint,1,opt,name=aType,proto3,enum=SWAdapterType" json:"aType,omitempty"`
	// vlan
	UnderlayInterface string `protobuf:"bytes,8,opt,name=underlayInterface,proto3" json:"underlayInterface,omitempty"`
	VlanId            uint32 `protobuf:"varint,9,opt,name=vlanId,proto3" json:"vlanId,omitempty"`
	// OR : repeated physical interfaces for bond0
	Bondgroup            []string `protobuf:"bytes,10,rep,name=bondgroup,proto3" json:"bondgroup,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SWAdapterParams) Reset()         { *m = SWAdapterParams{} }
func (m *SWAdapterParams) String() string { return proto.CompactTextString(m) }
func (*SWAdapterParams) ProtoMessage()    {}
func (*SWAdapterParams) Descriptor() ([]byte, []int) {
	return fileDescriptor_9fb58492383773ea, []int{0}
}

func (m *SWAdapterParams) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SWAdapterParams.Unmarshal(m, b)
}
func (m *SWAdapterParams) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SWAdapterParams.Marshal(b, m, deterministic)
}
func (m *SWAdapterParams) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SWAdapterParams.Merge(m, src)
}
func (m *SWAdapterParams) XXX_Size() int {
	return xxx_messageInfo_SWAdapterParams.Size(m)
}
func (m *SWAdapterParams) XXX_DiscardUnknown() {
	xxx_messageInfo_SWAdapterParams.DiscardUnknown(m)
}

var xxx_messageInfo_SWAdapterParams proto.InternalMessageInfo

func (m *SWAdapterParams) GetAType() SWAdapterType {
	if m != nil {
		return m.AType
	}
	return SWAdapterType_IGNORE
}

func (m *SWAdapterParams) GetUnderlayInterface() string {
	if m != nil {
		return m.UnderlayInterface
	}
	return ""
}

func (m *SWAdapterParams) GetVlanId() uint32 {
	if m != nil {
		return m.VlanId
	}
	return 0
}

func (m *SWAdapterParams) GetBondgroup() []string {
	if m != nil {
		return m.Bondgroup
	}
	return nil
}

// systemAdapters, are the higher l2 concept built on physicalIOs.
// systemAdapters, gives all the required bits to turn the physical IOs
// into useful IP endpoints
// These endpoints can be further used to connect to controller or
// can be shared between workload/services running on the node.
type SystemAdapter struct {
	// name of the adapter; hardware-specific e.g., eth0
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// this is part of the freelink group
	// deprecate: look at PhysicalIO->UsagePolicy instead
	FreeUplink bool `protobuf:"varint,2,opt,name=freeUplink,proto3" json:"freeUplink,omitempty"`
	// this is part of the uplink group
	// deprecate: look at PhysicalIO->Usage instead
	Uplink bool `protobuf:"varint,3,opt,name=uplink,proto3" json:"uplink,omitempty"`
	// attach this network config for this adapter
	NetworkUUID string `protobuf:"bytes,4,opt,name=networkUUID,proto3" json:"networkUUID,omitempty"`
	// if its static network we need ip address
	Addr string `protobuf:"bytes,5,opt,name=addr,proto3" json:"addr,omitempty"`
	// alias/logical name which will be reported to zedcloud
	// and used for app instances
	LogicalName          string   `protobuf:"bytes,6,opt,name=logicalName,proto3" json:"logicalName,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SystemAdapter) Reset()         { *m = SystemAdapter{} }
func (m *SystemAdapter) String() string { return proto.CompactTextString(m) }
func (*SystemAdapter) ProtoMessage()    {}
func (*SystemAdapter) Descriptor() ([]byte, []int) {
	return fileDescriptor_9fb58492383773ea, []int{1}
}

func (m *SystemAdapter) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SystemAdapter.Unmarshal(m, b)
}
func (m *SystemAdapter) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SystemAdapter.Marshal(b, m, deterministic)
}
func (m *SystemAdapter) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SystemAdapter.Merge(m, src)
}
func (m *SystemAdapter) XXX_Size() int {
	return xxx_messageInfo_SystemAdapter.Size(m)
}
func (m *SystemAdapter) XXX_DiscardUnknown() {
	xxx_messageInfo_SystemAdapter.DiscardUnknown(m)
}

var xxx_messageInfo_SystemAdapter proto.InternalMessageInfo

func (m *SystemAdapter) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *SystemAdapter) GetFreeUplink() bool {
	if m != nil {
		return m.FreeUplink
	}
	return false
}

func (m *SystemAdapter) GetUplink() bool {
	if m != nil {
		return m.Uplink
	}
	return false
}

func (m *SystemAdapter) GetNetworkUUID() string {
	if m != nil {
		return m.NetworkUUID
	}
	return ""
}

func (m *SystemAdapter) GetAddr() string {
	if m != nil {
		return m.Addr
	}
	return ""
}

func (m *SystemAdapter) GetLogicalName() string {
	if m != nil {
		return m.LogicalName
	}
	return ""
}

// Given additional details for EVE softwar to how to treat this
// interface. Example policies could be limit use of LTE interface
// or only use Eth1 only if Eth0 is not available etc
type PhyIOUsagePolicy struct {
	// Used only when one other normal uplinks are down
	FreeUplink           bool     `protobuf:"varint,1,opt,name=freeUplink,proto3" json:"freeUplink,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PhyIOUsagePolicy) Reset()         { *m = PhyIOUsagePolicy{} }
func (m *PhyIOUsagePolicy) String() string { return proto.CompactTextString(m) }
func (*PhyIOUsagePolicy) ProtoMessage()    {}
func (*PhyIOUsagePolicy) Descriptor() ([]byte, []int) {
	return fileDescriptor_9fb58492383773ea, []int{2}
}

func (m *PhyIOUsagePolicy) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PhyIOUsagePolicy.Unmarshal(m, b)
}
func (m *PhyIOUsagePolicy) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PhyIOUsagePolicy.Marshal(b, m, deterministic)
}
func (m *PhyIOUsagePolicy) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PhyIOUsagePolicy.Merge(m, src)
}
func (m *PhyIOUsagePolicy) XXX_Size() int {
	return xxx_messageInfo_PhyIOUsagePolicy.Size(m)
}
func (m *PhyIOUsagePolicy) XXX_DiscardUnknown() {
	xxx_messageInfo_PhyIOUsagePolicy.DiscardUnknown(m)
}

var xxx_messageInfo_PhyIOUsagePolicy proto.InternalMessageInfo

func (m *PhyIOUsagePolicy) GetFreeUplink() bool {
	if m != nil {
		return m.FreeUplink
	}
	return false
}

type PhyAddress struct {
	Pcilong string `protobuf:"bytes,2,opt,name=pcilong,proto3" json:"pcilong,omitempty"`
	Ifname  string `protobuf:"bytes,2,opt,name=ifname,proto3" json:"ifname,omitempty"`
	Serial  string `protobuf:"bytes,2,opt,name=serial,proto3" json:"serial,omitempty"`
	Irq     string `protobuf:"bytes,2,opt,name=irq,proto3" json:"irq,omitempty"`
	Ioports string `protobuf:"bytes,2,opt,name=ioports,proto3" json:"ioports,omitempty"`
}

// PhysicalIO:
//    Absolute low level description of physical buses and ports that are
//    available on given platfrom.
//    Collection of these IOs, connstitue what we would call as hardware
//    model. Each physical IO is manageable and visible to EVE software, and
//    it can be further configured to either provide IP connectivity or
//    directly be given to workloads
type PhysicalIO struct {
	Ptype PhyIoType `protobuf:"varint,1,opt,name=ptype,proto3,enum=PhyIoType" json:"ptype,omitempty"`
	// physical label typically printed on box.
	// Example Eth0, Eth1, Wifi0, ComA, ComB
	Phylabel string `protobuf:"bytes,2,opt,name=phylabel,proto3" json:"phylabel,omitempty"`
	// The hardware bus address. The key to this map can be of the following
	// (case-insensitive) values:
	// "pcilong": the address is a PCI id of the form 0000:02:00.0
	// "ifname": the addresss is a string for a network interface like "eth1"
	// "serial": the address is a Linux serial port alias such as "/dev/ttyS2"
	// "irq": the address is a number such as "5"
	// "ioports": the address is a string such as "2f8-2ff"
	Phyaddrs map[string]string `protobuf:"bytes,3,rep,name=phyaddrs,proto3" json:"phyaddrs,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`

	// The hardware bus address. This Obsoletes Phyaddrs map above.
	PhyAddr *PhyAddress

	// provides the ability to model designer to rename the physicalIO
	// port to more understandable
	// For example Eth0->Mgmt0
	//  or USBA->ConfigDiskA etc
	Logicallabel string `protobuf:"bytes,4,opt,name=logicallabel,proto3" json:"logicallabel,omitempty"`
	// Assignment Group, is unique label that is applied across PhysicalIOs
	// EntireGroup can be assigned to application or nothing at all
	Assigngrp   string            `protobuf:"bytes,5,opt,name=assigngrp,proto3" json:"assigngrp,omitempty"`
	Usage       PhyIoMemberUsage  `protobuf:"varint,6,opt,name=usage,proto3,enum=PhyIoMemberUsage" json:"usage,omitempty"`
	UsagePolicy *PhyIOUsagePolicy `protobuf:"bytes,7,opt,name=usagePolicy,proto3" json:"usagePolicy,omitempty"`
	// physical and logical attributes
	//    For example in WWAN to which firmware version to laod etc
	// OBSOLETE. TO BE REMOVED>
	Cbattr               map[string]string `protobuf:"bytes,8,rep,name=cbattr,proto3" json:"cbattr,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *PhysicalIO) Reset()         { *m = PhysicalIO{} }
func (m *PhysicalIO) String() string { return proto.CompactTextString(m) }
func (*PhysicalIO) ProtoMessage()    {}
func (*PhysicalIO) Descriptor() ([]byte, []int) {
	return fileDescriptor_9fb58492383773ea, []int{3}
}

func (m *PhysicalIO) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PhysicalIO.Unmarshal(m, b)
}
func (m *PhysicalIO) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PhysicalIO.Marshal(b, m, deterministic)
}
func (m *PhysicalIO) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PhysicalIO.Merge(m, src)
}
func (m *PhysicalIO) XXX_Size() int {
	return xxx_messageInfo_PhysicalIO.Size(m)
}
func (m *PhysicalIO) XXX_DiscardUnknown() {
	xxx_messageInfo_PhysicalIO.DiscardUnknown(m)
}

var xxx_messageInfo_PhysicalIO proto.InternalMessageInfo

func (m *PhysicalIO) GetPtype() PhyIoType {
	if m != nil {
		return m.Ptype
	}
	return PhyIoType_PhyIoNoop
}

func (m *PhysicalIO) GetPhylabel() string {
	if m != nil {
		return m.Phylabel
	}
	return ""
}

func (m *PhysicalIO) GetPhyaddrs() map[string]string {
	if m != nil {
		return m.Phyaddrs
	}
	return nil
}

func (m *PhysicalIO) GetLogicallabel() string {
	if m != nil {
		return m.Logicallabel
	}
	return ""
}

func (m *PhysicalIO) GetAssigngrp() string {
	if m != nil {
		return m.Assigngrp
	}
	return ""
}

func (m *PhysicalIO) GetUsage() PhyIoMemberUsage {
	if m != nil {
		return m.Usage
	}
	return PhyIoMemberUsage_PhyIoUsageNone
}

func (m *PhysicalIO) GetUsagePolicy() *PhyIOUsagePolicy {
	if m != nil {
		return m.UsagePolicy
	}
	return nil
}

func (m *PhysicalIO) GetCbattr() map[string]string {
	if m != nil {
		return m.Cbattr
	}
	return nil
}

func init() {
	proto.RegisterEnum("SWAdapterType", SWAdapterType_name, SWAdapterType_value)
	proto.RegisterEnum("PhyIoType", PhyIoType_name, PhyIoType_value)
	proto.RegisterEnum("PhyIoMemberUsage", PhyIoMemberUsage_name, PhyIoMemberUsage_value)
	proto.RegisterType((*SWAdapterParams)(nil), "sWAdapterParams")
	proto.RegisterType((*SystemAdapter)(nil), "SystemAdapter")
	proto.RegisterType((*PhyIOUsagePolicy)(nil), "PhyIOUsagePolicy")
	proto.RegisterType((*PhysicalIO)(nil), "PhysicalIO")
	proto.RegisterMapType((map[string]string)(nil), "PhysicalIO.CbattrEntry")
	proto.RegisterMapType((map[string]string)(nil), "PhysicalIO.PhyaddrsEntry")
}

func init() { proto.RegisterFile("devmodel.proto", fileDescriptor_9fb58492383773ea) }

var fileDescriptor_9fb58492383773ea = []byte{
	// 698 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x54, 0xdd, 0x6e, 0xe2, 0x46,
	0x14, 0x8e, 0x31, 0x10, 0x38, 0x04, 0x32, 0x99, 0x46, 0x8d, 0x1b, 0x55, 0xad, 0x85, 0x22, 0x15,
	0x45, 0xad, 0x91, 0x88, 0x2a, 0xf5, 0xe7, 0x8a, 0x84, 0x28, 0xb5, 0x54, 0x7e, 0xe4, 0x2c, 0x1b,
	0x69, 0xef, 0x06, 0xcf, 0x60, 0xac, 0xd8, 0x1e, 0x6b, 0x6c, 0xb3, 0xf2, 0xbe, 0x4a, 0x1e, 0x63,
	0xdf, 0x62, 0x5f, 0x6a, 0x57, 0xe3, 0xf1, 0x82, 0xc9, 0x5e, 0xed, 0xdd, 0x39, 0xdf, 0x77, 0x7e,
	0xbf, 0x33, 0x36, 0xf4, 0x28, 0xdb, 0x86, 0x9c, 0xb2, 0xc0, 0x8a, 0x05, 0x4f, 0x79, 0xff, 0x45,
	0x83, 0xd3, 0xe4, 0x69, 0x4c, 0x49, 0x9c, 0x32, 0xb1, 0x20, 0x82, 0x84, 0x09, 0xbe, 0x82, 0x06,
	0x79, 0x93, 0xc7, 0xcc, 0xd0, 0x4c, 0x6d, 0xd0, 0x1b, 0xf5, 0xac, 0x5d, 0x80, 0x44, 0x1d, 0x45,
	0xe2, 0xdf, 0xe1, 0x2c, 0x8b, 0x28, 0x13, 0x01, 0xc9, 0xed, 0x28, 0x65, 0x62, 0x4d, 0x5c, 0x66,
	0xb4, 0x4c, 0x6d, 0xd0, 0x76, 0xbe, 0x25, 0xf0, 0x8f, 0xd0, 0xdc, 0x06, 0x24, 0xb2, 0xa9, 0xd1,
	0x36, 0xb5, 0x41, 0xd7, 0x29, 0x3d, 0xfc, 0x33, 0xb4, 0x57, 0x3c, 0xa2, 0x9e, 0xe0, 0x59, 0x6c,
	0x80, 0xa9, 0x0f, 0xda, 0xce, 0x1e, 0xe8, 0x7f, 0xd4, 0xa0, 0xfb, 0x98, 0x27, 0x29, 0x0b, 0xcb,
	0x01, 0x30, 0x86, 0x7a, 0x44, 0x42, 0x35, 0x5a, 0xdb, 0x29, 0x6c, 0xfc, 0x0b, 0xc0, 0x5a, 0x30,
	0xb6, 0x8c, 0x03, 0x3f, 0x7a, 0x36, 0x6a, 0xa6, 0x36, 0x68, 0x39, 0x15, 0x44, 0xf6, 0xce, 0x14,
	0xa7, 0x17, 0x5c, 0xe9, 0x61, 0x13, 0x3a, 0x11, 0x4b, 0xdf, 0x73, 0xf1, 0xbc, 0x5c, 0xda, 0x13,
	0xa3, 0x5e, 0x94, 0xac, 0x42, 0xb2, 0x1b, 0xa1, 0x54, 0x18, 0x0d, 0xd5, 0x4d, 0xda, 0x32, 0x2b,
	0xe0, 0x9e, 0xef, 0x92, 0x60, 0x26, 0x07, 0x69, 0xaa, 0xac, 0x0a, 0xd4, 0x1f, 0x01, 0x5a, 0x6c,
	0x72, 0x7b, 0xbe, 0x4c, 0x88, 0xc7, 0x16, 0x3c, 0xf0, 0xdd, 0xfc, 0xd5, 0x8c, 0xda, 0xeb, 0x19,
	0xfb, 0x9f, 0x74, 0x80, 0xc5, 0x26, 0x4f, 0x64, 0x11, 0x7b, 0x8e, 0x4d, 0x68, 0xc4, 0xe9, 0xfe,
	0x04, 0x60, 0xc9, 0x82, 0x5c, 0xc9, 0x5f, 0x10, 0xf8, 0x12, 0x5a, 0xf1, 0x26, 0x0f, 0xc8, 0x8a,
	0x05, 0xc5, 0xca, 0x6d, 0x67, 0xe7, 0xe3, 0x3f, 0x0b, 0x4e, 0x4e, 0x9b, 0x18, 0xba, 0xa9, 0x0f,
	0x3a, 0xa3, 0x9f, 0xac, 0x7d, 0x71, 0x69, 0x16, 0xdc, 0x7d, 0x94, 0x8a, 0xdc, 0xd9, 0x85, 0xe2,
	0x3e, 0x9c, 0x94, 0x6b, 0xa8, 0xb2, 0x4a, 0x90, 0x03, 0x4c, 0xde, 0x8b, 0x24, 0x89, 0xef, 0x45,
	0x9e, 0x88, 0x4b, 0x59, 0xf6, 0x00, 0xfe, 0x0d, 0x1a, 0x99, 0x5c, 0xba, 0x50, 0xa5, 0x37, 0x3a,
	0x53, 0x63, 0x4f, 0x59, 0xb8, 0x62, 0xa2, 0x50, 0xc3, 0x51, 0x3c, 0xbe, 0x81, 0x4e, 0xb6, 0x57,
	0xc7, 0x38, 0x36, 0xb5, 0x41, 0xa7, 0x0c, 0xaf, 0xca, 0xe6, 0x54, 0xa3, 0xf0, 0x10, 0x9a, 0xee,
	0x8a, 0xa4, 0xa9, 0x30, 0x5a, 0xc5, 0x52, 0x17, 0xd5, 0xa5, 0xee, 0x0a, 0x46, 0xad, 0x54, 0x86,
	0x5d, 0xfe, 0x0b, 0xdd, 0x83, 0x5d, 0x31, 0x02, 0xfd, 0x99, 0xe5, 0xe5, 0xe3, 0x91, 0x26, 0x3e,
	0x87, 0xc6, 0x96, 0x04, 0x19, 0x2b, 0x35, 0x54, 0xce, 0x3f, 0xb5, 0xbf, 0xb4, 0xcb, 0xbf, 0xa1,
	0x53, 0xa9, 0xf9, 0x3d, 0xa9, 0xd7, 0x43, 0xe8, 0x1e, 0x7c, 0x32, 0x18, 0xa0, 0x69, 0x3f, 0xcc,
	0xe6, 0xce, 0x3d, 0x3a, 0xc2, 0x2d, 0xa8, 0xbf, 0xfd, 0x7f, 0x3c, 0x43, 0x9a, 0xb4, 0x6e, 0xe7,
	0xb3, 0x09, 0xaa, 0x5d, 0xbf, 0x68, 0xd0, 0xde, 0x5d, 0x18, 0x77, 0x4b, 0x67, 0xc6, 0x79, 0x8c,
	0x8e, 0xf0, 0x29, 0x74, 0x94, 0xcb, 0xd2, 0xfb, 0x74, 0x83, 0x34, 0x7c, 0x02, 0xad, 0x02, 0x58,
	0x3e, 0xde, 0xa2, 0xda, 0xce, 0xbb, 0x9b, 0x4f, 0x91, 0x8e, 0x7b, 0xc5, 0x33, 0xb2, 0xf9, 0x38,
	0xa3, 0x3e, 0x47, 0x75, 0x8c, 0xe0, 0xe4, 0x6b, 0xf2, 0x93, 0xec, 0xda, 0x38, 0x40, 0x9e, 0xc6,
	0x33, 0xd4, 0xdc, 0xf5, 0xfb, 0x6f, 0x32, 0xb5, 0xd1, 0x31, 0x3e, 0x2d, 0x4b, 0xcc, 0xd3, 0x0d,
	0x13, 0xe8, 0xb3, 0x76, 0xed, 0xab, 0xf7, 0x5c, 0xbd, 0x23, 0xc6, 0xd0, 0x53, 0x33, 0x48, 0x6f,
	0xc6, 0x23, 0x86, 0x8e, 0x0e, 0xb1, 0xa9, 0x17, 0xa6, 0x48, 0xc3, 0xe7, 0x65, 0x6e, 0x81, 0x3d,
	0x6e, 0x88, 0x60, 0x14, 0xd5, 0xf0, 0x05, 0xfc, 0xb0, 0x47, 0x27, 0x8c, 0xfa, 0x2e, 0x49, 0x19,
	0x45, 0xfa, 0xed, 0x03, 0xfc, 0xea, 0xf2, 0xd0, 0xfa, 0xc0, 0x28, 0xa3, 0xc4, 0x72, 0x03, 0x9e,
	0x51, 0x2b, 0x4b, 0x98, 0xd8, 0xfa, 0x2e, 0x53, 0x7f, 0xac, 0x77, 0x57, 0x9e, 0x9f, 0x6e, 0xb2,
	0x95, 0xe5, 0xf2, 0x70, 0x18, 0xac, 0xff, 0x60, 0xd4, 0x63, 0x43, 0xb6, 0x65, 0x43, 0x12, 0xfb,
	0x43, 0x8f, 0x0f, 0x5d, 0x1e, 0xad, 0x7d, 0x6f, 0xd5, 0x2c, 0x82, 0x6f, 0xbe, 0x04, 0x00, 0x00,
	0xff, 0xff, 0xd3, 0xb6, 0xd8, 0x5a, 0xf0, 0x04, 0x00, 0x00,
}
