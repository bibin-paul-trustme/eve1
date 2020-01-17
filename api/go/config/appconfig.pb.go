// Code generated by protoc-gen-go. DO NOT EDIT.
// source: appconfig.proto

package config

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
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

type InstanceOpsCmd struct {
	Counter              uint32   `protobuf:"varint,2,opt,name=counter,proto3" json:"counter,omitempty"`
	OpsTime              string   `protobuf:"bytes,4,opt,name=opsTime,proto3" json:"opsTime,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *InstanceOpsCmd) Reset()         { *m = InstanceOpsCmd{} }
func (m *InstanceOpsCmd) String() string { return proto.CompactTextString(m) }
func (*InstanceOpsCmd) ProtoMessage()    {}
func (*InstanceOpsCmd) Descriptor() ([]byte, []int) {
	return fileDescriptor_6183fdf07ef5608d, []int{0}
}

func (m *InstanceOpsCmd) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_InstanceOpsCmd.Unmarshal(m, b)
}
func (m *InstanceOpsCmd) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_InstanceOpsCmd.Marshal(b, m, deterministic)
}
func (m *InstanceOpsCmd) XXX_Merge(src proto.Message) {
	xxx_messageInfo_InstanceOpsCmd.Merge(m, src)
}
func (m *InstanceOpsCmd) XXX_Size() int {
	return xxx_messageInfo_InstanceOpsCmd.Size(m)
}
func (m *InstanceOpsCmd) XXX_DiscardUnknown() {
	xxx_messageInfo_InstanceOpsCmd.DiscardUnknown(m)
}

var xxx_messageInfo_InstanceOpsCmd proto.InternalMessageInfo

func (m *InstanceOpsCmd) GetCounter() uint32 {
	if m != nil {
		return m.Counter
	}
	return 0
}

func (m *InstanceOpsCmd) GetOpsTime() string {
	if m != nil {
		return m.OpsTime
	}
	return ""
}

type AppInstanceConfig struct {
	Uuidandversion *UUIDandVersion `protobuf:"bytes,1,opt,name=uuidandversion,proto3" json:"uuidandversion,omitempty"`
	Displayname    string          `protobuf:"bytes,2,opt,name=displayname,proto3" json:"displayname,omitempty"`
	Fixedresources *VmConfig       `protobuf:"bytes,3,opt,name=fixedresources,proto3" json:"fixedresources,omitempty"`
	Drives         []*Drive        `protobuf:"bytes,4,rep,name=drives,proto3" json:"drives,omitempty"`
	Activate       bool            `protobuf:"varint,5,opt,name=activate,proto3" json:"activate,omitempty"`
	// NetworkAdapter are virtual adapters assigned to the application
	// Physical adapters such as eth1 are part of Adapter
	Interfaces []*NetworkAdapter `protobuf:"bytes,6,rep,name=interfaces,proto3" json:"interfaces,omitempty"`
	// adapters - Name in Adapter should be set to PhysicalIO.assigngrp
	Adapters []*Adapter `protobuf:"bytes,7,rep,name=adapters,proto3" json:"adapters,omitempty"`
	// The device behavior for a restart command (if counter increased)
	// is to restart the application instance honoring the persist setting
	// for the disks/drives.
	// The device can assume that the adapters did not change.
	Restart *InstanceOpsCmd `protobuf:"bytes,9,opt,name=restart,proto3" json:"restart,omitempty"`
	// The device behavior for a purge command is to restart the domU.
	// with the disks/drives recreated from the downloaded images
	// (whether preserve is set or not).
	//    if the manifest is changed with purge option, new manifest will
	//    be used. Device doesn't know what has changed, it will get the
	//    changed config.
	//
	//    if disks section have changed will be purged automatically.
	//    phase 1: we would purge all disks irrespective preserve flag
	Purge *InstanceOpsCmd `protobuf:"bytes,10,opt,name=purge,proto3" json:"purge,omitempty"`
	// App Instance initialization configuration data provided by user
	// This will be used as "user-data" in cloud-init
	// Empty string will indicate that cloud-init is not required
	UserData string `protobuf:"bytes,11,opt,name=userData,proto3" json:"userData,omitempty"`
	// Config flag if the app-instance should be made accessible
	// through a remote console session established by the device.
	RemoteConsole bool `protobuf:"varint,12,opt,name=remoteConsole,proto3" json:"remoteConsole,omitempty"`
	// Trigger refresh of Images used by the App Instance.
	// This will trigger re-download of images. If we have a new version of the
	// image, this will result in a purge of the image.
	// All the images used by the App Instance would be refreshed, except
	// the ones of type HDD_EMPTY
	RefreshImage         *InstanceOpsCmd `protobuf:"bytes,13,opt,name=refreshImage,proto3" json:"refreshImage,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *AppInstanceConfig) Reset()         { *m = AppInstanceConfig{} }
func (m *AppInstanceConfig) String() string { return proto.CompactTextString(m) }
func (*AppInstanceConfig) ProtoMessage()    {}
func (*AppInstanceConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_6183fdf07ef5608d, []int{1}
}

func (m *AppInstanceConfig) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AppInstanceConfig.Unmarshal(m, b)
}
func (m *AppInstanceConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AppInstanceConfig.Marshal(b, m, deterministic)
}
func (m *AppInstanceConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AppInstanceConfig.Merge(m, src)
}
func (m *AppInstanceConfig) XXX_Size() int {
	return xxx_messageInfo_AppInstanceConfig.Size(m)
}
func (m *AppInstanceConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_AppInstanceConfig.DiscardUnknown(m)
}

var xxx_messageInfo_AppInstanceConfig proto.InternalMessageInfo

func (m *AppInstanceConfig) GetUuidandversion() *UUIDandVersion {
	if m != nil {
		return m.Uuidandversion
	}
	return nil
}

func (m *AppInstanceConfig) GetDisplayname() string {
	if m != nil {
		return m.Displayname
	}
	return ""
}

func (m *AppInstanceConfig) GetFixedresources() *VmConfig {
	if m != nil {
		return m.Fixedresources
	}
	return nil
}

func (m *AppInstanceConfig) GetDrives() []*Drive {
	if m != nil {
		return m.Drives
	}
	return nil
}

func (m *AppInstanceConfig) GetActivate() bool {
	if m != nil {
		return m.Activate
	}
	return false
}

func (m *AppInstanceConfig) GetInterfaces() []*NetworkAdapter {
	if m != nil {
		return m.Interfaces
	}
	return nil
}

func (m *AppInstanceConfig) GetAdapters() []*Adapter {
	if m != nil {
		return m.Adapters
	}
	return nil
}

func (m *AppInstanceConfig) GetRestart() *InstanceOpsCmd {
	if m != nil {
		return m.Restart
	}
	return nil
}

func (m *AppInstanceConfig) GetPurge() *InstanceOpsCmd {
	if m != nil {
		return m.Purge
	}
	return nil
}

func (m *AppInstanceConfig) GetUserData() string {
	if m != nil {
		return m.UserData
	}
	return ""
}

func (m *AppInstanceConfig) GetRemoteConsole() bool {
	if m != nil {
		return m.RemoteConsole
	}
	return false
}

func (m *AppInstanceConfig) GetRefreshImage() *InstanceOpsCmd {
	if m != nil {
		return m.RefreshImage
	}
	return nil
}

func init() {
	proto.RegisterType((*InstanceOpsCmd)(nil), "InstanceOpsCmd")
	proto.RegisterType((*AppInstanceConfig)(nil), "AppInstanceConfig")
}

func init() { proto.RegisterFile("appconfig.proto", fileDescriptor_6183fdf07ef5608d) }

var fileDescriptor_6183fdf07ef5608d = []byte{
	// 436 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x92, 0xcb, 0x6b, 0x1b, 0x31,
	0x18, 0xc4, 0xd9, 0x3a, 0xb1, 0xd7, 0x72, 0x6c, 0x53, 0x9d, 0x44, 0x0e, 0xed, 0x12, 0x5c, 0xd8,
	0x1e, 0xaa, 0xa5, 0xc9, 0xa1, 0xe7, 0x34, 0x86, 0xe2, 0x4b, 0x0b, 0xa2, 0xc9, 0xa1, 0x37, 0x45,
	0xfa, 0x76, 0x23, 0x6a, 0x3d, 0x90, 0xb4, 0xdb, 0xc7, 0xff, 0x5e, 0x28, 0xfb, 0x32, 0x76, 0xc8,
	0x71, 0x66, 0x7e, 0xd2, 0x48, 0x9f, 0x84, 0xd6, 0xdc, 0x39, 0x61, 0x4d, 0xa9, 0x2a, 0xea, 0xbc,
	0x8d, 0xf6, 0x72, 0x2d, 0xa1, 0x11, 0x56, 0x6b, 0x6b, 0x06, 0x63, 0x19, 0xa2, 0xf5, 0xbc, 0x82,
	0x41, 0xa6, 0x8d, 0x1e, 0x49, 0x03, 0xf1, 0x78, 0xe9, 0xd5, 0x16, 0xad, 0x76, 0x26, 0x44, 0x6e,
	0x04, 0x7c, 0x73, 0xe1, 0x4e, 0x4b, 0x4c, 0xd0, 0x4c, 0xd8, 0xda, 0x44, 0xf0, 0xe4, 0x55, 0x96,
	0xe4, 0x4b, 0x36, 0xca, 0x36, 0xb1, 0x2e, 0x7c, 0x57, 0x1a, 0xc8, 0x59, 0x96, 0xe4, 0x73, 0x36,
	0xca, 0xab, 0x7f, 0x13, 0xf4, 0xfa, 0xd6, 0xb9, 0x71, 0xa7, 0xbb, 0xae, 0x01, 0x7f, 0x42, 0xab,
	0xba, 0x56, 0x92, 0x1b, 0xd9, 0x80, 0x0f, 0xca, 0x1a, 0x92, 0x64, 0x49, 0xbe, 0xb8, 0x5e, 0xd3,
	0xfb, 0xfb, 0xdd, 0x96, 0x1b, 0xf9, 0xd0, 0xdb, 0xec, 0x19, 0x86, 0x33, 0xb4, 0x90, 0x2a, 0xb8,
	0x3d, 0xff, 0x63, 0xb8, 0x86, 0xee, 0x18, 0x73, 0x76, 0x6c, 0xe1, 0x8f, 0x68, 0x55, 0xaa, 0xdf,
	0x20, 0x3d, 0x04, 0x5b, 0x7b, 0x01, 0x81, 0x4c, 0xba, 0xad, 0xe7, 0xf4, 0x41, 0xf7, 0xed, 0xec,
	0x19, 0x80, 0xdf, 0xa0, 0xa9, 0xf4, 0xaa, 0x81, 0x40, 0xce, 0xb2, 0x49, 0xbe, 0xb8, 0x9e, 0xd2,
	0x6d, 0x2b, 0xd9, 0xe0, 0xe2, 0x4b, 0x94, 0x72, 0x11, 0x55, 0xc3, 0x23, 0x90, 0xf3, 0x2c, 0xc9,
	0x53, 0x76, 0xd0, 0xb8, 0x40, 0x48, 0xb5, 0x23, 0x28, 0x79, 0x5b, 0x35, 0xed, 0xd6, 0xaf, 0xe9,
	0x57, 0x88, 0xbf, 0xac, 0xff, 0x79, 0x2b, 0xb9, 0x8b, 0xe0, 0xd9, 0x11, 0x82, 0x37, 0x28, 0xe5,
	0xbd, 0x1d, 0xc8, 0xac, 0xc3, 0x53, 0x3a, 0x72, 0x87, 0x04, 0xbf, 0x47, 0x33, 0x0f, 0x21, 0x72,
	0x1f, 0xc9, 0x7c, 0x98, 0xcc, 0xe9, 0x63, 0xb0, 0x31, 0xc7, 0xef, 0xd0, 0xb9, 0xab, 0x7d, 0x05,
	0x04, 0xbd, 0x0c, 0xf6, 0x69, 0x7b, 0x89, 0x3a, 0x80, 0xdf, 0xf2, 0xc8, 0xc9, 0xa2, 0x1b, 0xdb,
	0x41, 0xe3, 0x0d, 0x5a, 0x7a, 0xd0, 0x36, 0xb6, 0xcf, 0x13, 0xec, 0x1e, 0xc8, 0x45, 0x77, 0xcb,
	0x53, 0x13, 0xdf, 0xa0, 0x0b, 0x0f, 0xa5, 0x87, 0xf0, 0xb4, 0xd3, 0xbc, 0x02, 0xb2, 0x7c, 0xb9,
	0xef, 0x04, 0xfa, 0xfc, 0x05, 0xbd, 0x15, 0x56, 0xd3, 0xbf, 0x20, 0x41, 0x72, 0x2a, 0xf6, 0xb6,
	0x96, 0xb4, 0xed, 0x6d, 0x94, 0x18, 0xfe, 0xe0, 0x8f, 0x4d, 0xa5, 0xe2, 0x53, 0xfd, 0x48, 0x85,
	0xd5, 0xc5, 0xbe, 0xfc, 0x00, 0xb2, 0x82, 0x02, 0x1a, 0x28, 0xb8, 0x53, 0x45, 0x65, 0x8b, 0xfe,
	0x53, 0x3e, 0x4e, 0x3b, 0xf8, 0xe6, 0x7f, 0x00, 0x00, 0x00, 0xff, 0xff, 0xff, 0x12, 0x54, 0xb4,
	0xe3, 0x02, 0x00, 0x00,
}
