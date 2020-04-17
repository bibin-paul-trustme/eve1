// Code generated by protoc-gen-go. DO NOT EDIT.
// source: config/devcommon.proto

package config

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	common "github.com/lf-edge/eve/api/go/common"
	math "math"
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

type UUIDandVersion struct {
	Uuid                 string   `protobuf:"bytes,1,opt,name=uuid,proto3" json:"uuid,omitempty"`
	Version              string   `protobuf:"bytes,2,opt,name=version,proto3" json:"version,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UUIDandVersion) Reset()         { *m = UUIDandVersion{} }
func (m *UUIDandVersion) String() string { return proto.CompactTextString(m) }
func (*UUIDandVersion) ProtoMessage()    {}
func (*UUIDandVersion) Descriptor() ([]byte, []int) {
	return fileDescriptor_929a6154e0c58b98, []int{0}
}

func (m *UUIDandVersion) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UUIDandVersion.Unmarshal(m, b)
}
func (m *UUIDandVersion) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UUIDandVersion.Marshal(b, m, deterministic)
}
func (m *UUIDandVersion) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UUIDandVersion.Merge(m, src)
}
func (m *UUIDandVersion) XXX_Size() int {
	return xxx_messageInfo_UUIDandVersion.Size(m)
}
func (m *UUIDandVersion) XXX_DiscardUnknown() {
	xxx_messageInfo_UUIDandVersion.DiscardUnknown(m)
}

var xxx_messageInfo_UUIDandVersion proto.InternalMessageInfo

func (m *UUIDandVersion) GetUuid() string {
	if m != nil {
		return m.Uuid
	}
	return ""
}

func (m *UUIDandVersion) GetVersion() string {
	if m != nil {
		return m.Version
	}
	return ""
}

// Device Operational Commands Semantic
// For rebooting device,  command=Reset, counter = counter+delta, desiredState = on
// For poweroff device,  command=Reset, counter = counter+delta, desiredState = off
// For backup at midnight, command=Backup, counter = counter+delta, desiredState=n/a, opsTime = mm/dd/yy:hh:ss
// Current implementation does support only single command outstanding for each type
// In future can be extended to have more scheduled events
//
type DeviceOpsCmd struct {
	Counter      uint32 `protobuf:"varint,2,opt,name=counter,proto3" json:"counter,omitempty"`
	DesiredState bool   `protobuf:"varint,3,opt,name=desiredState,proto3" json:"desiredState,omitempty"`
	// FIXME: change to timestamp, once we move to gogo proto
	OpsTime              string   `protobuf:"bytes,4,opt,name=opsTime,proto3" json:"opsTime,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeviceOpsCmd) Reset()         { *m = DeviceOpsCmd{} }
func (m *DeviceOpsCmd) String() string { return proto.CompactTextString(m) }
func (*DeviceOpsCmd) ProtoMessage()    {}
func (*DeviceOpsCmd) Descriptor() ([]byte, []int) {
	return fileDescriptor_929a6154e0c58b98, []int{1}
}

func (m *DeviceOpsCmd) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeviceOpsCmd.Unmarshal(m, b)
}
func (m *DeviceOpsCmd) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeviceOpsCmd.Marshal(b, m, deterministic)
}
func (m *DeviceOpsCmd) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeviceOpsCmd.Merge(m, src)
}
func (m *DeviceOpsCmd) XXX_Size() int {
	return xxx_messageInfo_DeviceOpsCmd.Size(m)
}
func (m *DeviceOpsCmd) XXX_DiscardUnknown() {
	xxx_messageInfo_DeviceOpsCmd.DiscardUnknown(m)
}

var xxx_messageInfo_DeviceOpsCmd proto.InternalMessageInfo

func (m *DeviceOpsCmd) GetCounter() uint32 {
	if m != nil {
		return m.Counter
	}
	return 0
}

func (m *DeviceOpsCmd) GetDesiredState() bool {
	if m != nil {
		return m.DesiredState
	}
	return false
}

func (m *DeviceOpsCmd) GetOpsTime() string {
	if m != nil {
		return m.OpsTime
	}
	return ""
}

// Timers and other per-device policy which relates to the interaction
// with zedcloud. Note that the timers are randomized on the device
// to avoid synchronization with other devices. Random range is between
// between .5 and 1.5 of these nominal values. If not set (i.e. zero),
// it means the default value of 60 seconds.
type ConfigItem struct {
	Key                  string   `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	Value                string   `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ConfigItem) Reset()         { *m = ConfigItem{} }
func (m *ConfigItem) String() string { return proto.CompactTextString(m) }
func (*ConfigItem) ProtoMessage()    {}
func (*ConfigItem) Descriptor() ([]byte, []int) {
	return fileDescriptor_929a6154e0c58b98, []int{2}
}

func (m *ConfigItem) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ConfigItem.Unmarshal(m, b)
}
func (m *ConfigItem) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ConfigItem.Marshal(b, m, deterministic)
}
func (m *ConfigItem) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ConfigItem.Merge(m, src)
}
func (m *ConfigItem) XXX_Size() int {
	return xxx_messageInfo_ConfigItem.Size(m)
}
func (m *ConfigItem) XXX_DiscardUnknown() {
	xxx_messageInfo_ConfigItem.DiscardUnknown(m)
}

var xxx_messageInfo_ConfigItem proto.InternalMessageInfo

func (m *ConfigItem) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func (m *ConfigItem) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

// Adapter bundles corresponding to a subset of what is in ZioBundle
type Adapter struct {
	Type                 common.PhyIoType `protobuf:"varint,1,opt,name=type,proto3,enum=PhyIoType" json:"type,omitempty"`
	Name                 string           `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *Adapter) Reset()         { *m = Adapter{} }
func (m *Adapter) String() string { return proto.CompactTextString(m) }
func (*Adapter) ProtoMessage()    {}
func (*Adapter) Descriptor() ([]byte, []int) {
	return fileDescriptor_929a6154e0c58b98, []int{3}
}

func (m *Adapter) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Adapter.Unmarshal(m, b)
}
func (m *Adapter) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Adapter.Marshal(b, m, deterministic)
}
func (m *Adapter) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Adapter.Merge(m, src)
}
func (m *Adapter) XXX_Size() int {
	return xxx_messageInfo_Adapter.Size(m)
}
func (m *Adapter) XXX_DiscardUnknown() {
	xxx_messageInfo_Adapter.DiscardUnknown(m)
}

var xxx_messageInfo_Adapter proto.InternalMessageInfo

func (m *Adapter) GetType() common.PhyIoType {
	if m != nil {
		return m.Type
	}
	return common.PhyIoType_PhyIoNoop
}

func (m *Adapter) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func init() {
	proto.RegisterType((*UUIDandVersion)(nil), "UUIDandVersion")
	proto.RegisterType((*DeviceOpsCmd)(nil), "DeviceOpsCmd")
	proto.RegisterType((*ConfigItem)(nil), "ConfigItem")
	proto.RegisterType((*Adapter)(nil), "Adapter")
}

func init() {
	proto.RegisterFile("config/devcommon.proto", fileDescriptor_929a6154e0c58b98)
}

var fileDescriptor_929a6154e0c58b98 = []byte{
	// 302 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x54, 0x51, 0x51, 0x4b, 0xf3, 0x30,
	0x14, 0x65, 0xdf, 0xf6, 0x39, 0xbd, 0xcc, 0x21, 0x45, 0xa4, 0x28, 0xe8, 0x28, 0x3e, 0xec, 0xc5,
	0x16, 0xd4, 0x57, 0x05, 0xdd, 0x40, 0xf6, 0xa4, 0xd4, 0xcd, 0x07, 0xdf, 0xb2, 0xe4, 0xae, 0x0b,
	0x36, 0xbd, 0xa1, 0x4d, 0x0a, 0xf5, 0xd7, 0x4b, 0x92, 0x0d, 0xf4, 0xed, 0x9c, 0x73, 0xcf, 0xe1,
	0xde, 0x93, 0xc0, 0x19, 0xa7, 0x6a, 0x23, 0x8b, 0x4c, 0x60, 0xcb, 0x49, 0x29, 0xaa, 0x52, 0x5d,
	0x93, 0xa1, 0xf3, 0x8b, 0xc0, 0x9c, 0xae, 0x48, 0x60, 0xf9, 0x7b, 0x98, 0x3c, 0xc2, 0x78, 0xb5,
	0x5a, 0xcc, 0x59, 0x25, 0x3e, 0xb0, 0x6e, 0x24, 0x55, 0x51, 0x04, 0x03, 0x6b, 0xa5, 0x88, 0x7b,
	0x93, 0xde, 0xf4, 0x28, 0xf7, 0x38, 0x8a, 0x61, 0xd8, 0x86, 0x71, 0xfc, 0xcf, 0xcb, 0x7b, 0x9a,
	0x6c, 0x60, 0x34, 0xc7, 0x56, 0x72, 0x7c, 0xd5, 0xcd, 0x4c, 0x79, 0x27, 0x27, 0x5b, 0x19, 0xac,
	0xbd, 0xf3, 0x38, 0xdf, 0xd3, 0x28, 0x81, 0x91, 0xc0, 0x46, 0xd6, 0x28, 0xde, 0x0d, 0x33, 0x18,
	0xf7, 0x27, 0xbd, 0xe9, 0x61, 0xfe, 0x47, 0x73, 0x69, 0xd2, 0xcd, 0x52, 0x2a, 0x8c, 0x07, 0x61,
	0xcf, 0x8e, 0x26, 0xf7, 0x00, 0x33, 0x5f, 0x6f, 0x61, 0x50, 0x45, 0x27, 0xd0, 0xff, 0xc2, 0x6e,
	0x77, 0xa2, 0x83, 0xd1, 0x29, 0xfc, 0x6f, 0x59, 0x69, 0x71, 0x77, 0x5f, 0x20, 0xc9, 0x03, 0x0c,
	0x9f, 0x04, 0xd3, 0x6e, 0xfd, 0x25, 0x0c, 0x4c, 0xa7, 0xd1, 0x67, 0xc6, 0xb7, 0x90, 0xbe, 0x6d,
	0xbb, 0x05, 0x2d, 0x3b, 0x8d, 0xb9, 0xd7, 0x5d, 0xed, 0x8a, 0xa9, 0x7d, 0xde, 0xe3, 0xe7, 0x17,
	0xb8, 0xe2, 0xa4, 0xd2, 0x6f, 0x14, 0x28, 0x58, 0xca, 0x4b, 0xb2, 0x22, 0xb5, 0x0d, 0xd6, 0xae,
	0x6f, 0x78, 0xbf, 0xcf, 0xeb, 0x42, 0x9a, 0xad, 0x5d, 0xa7, 0x9c, 0x54, 0x56, 0x6e, 0x6e, 0x50,
	0x14, 0x98, 0x61, 0x8b, 0x19, 0xd3, 0x32, 0x2b, 0x28, 0x0b, 0x5f, 0xb2, 0x3e, 0xf0, 0xe6, 0xbb,
	0x9f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x3c, 0x9e, 0x13, 0xd1, 0xa3, 0x01, 0x00, 0x00,
}
