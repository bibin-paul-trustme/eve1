// Code generated by protoc-gen-go. DO NOT EDIT.
// source: devcommon.proto

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

// Security Key Exchange Method
type KeyExchangeScheme int32

const (
	KeyExchangeScheme_KEA_NONE KeyExchangeScheme = 0
	KeyExchangeScheme_KEA_ECDH KeyExchangeScheme = 1
)

var KeyExchangeScheme_name = map[int32]string{
	0: "KEA_NONE",
	1: "KEA_ECDH",
}

var KeyExchangeScheme_value = map[string]int32{
	"KEA_NONE": 0,
	"KEA_ECDH": 1,
}

func (x KeyExchangeScheme) String() string {
	return proto.EnumName(KeyExchangeScheme_name, int32(x))
}

func (KeyExchangeScheme) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_c2bb9fc347232ae8, []int{0}
}

// Encryption Scheme for Cipher Payload
type EncryptionScheme int32

const (
	EncryptionScheme_SA_NONE        EncryptionScheme = 0
	EncryptionScheme_SA_AES_256_CFB EncryptionScheme = 1
)

var EncryptionScheme_name = map[int32]string{
	0: "SA_NONE",
	1: "SA_AES_256_CFB",
}

var EncryptionScheme_value = map[string]int32{
	"SA_NONE":        0,
	"SA_AES_256_CFB": 1,
}

func (x EncryptionScheme) String() string {
	return proto.EnumName(EncryptionScheme_name, int32(x))
}

func (EncryptionScheme) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_c2bb9fc347232ae8, []int{1}
}

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
	return fileDescriptor_c2bb9fc347232ae8, []int{0}
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
	return fileDescriptor_c2bb9fc347232ae8, []int{1}
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
	return fileDescriptor_c2bb9fc347232ae8, []int{2}
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
	Type                 PhyIoType `protobuf:"varint,1,opt,name=type,proto3,enum=PhyIoType" json:"type,omitempty"`
	Name                 string    `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *Adapter) Reset()         { *m = Adapter{} }
func (m *Adapter) String() string { return proto.CompactTextString(m) }
func (*Adapter) ProtoMessage()    {}
func (*Adapter) Descriptor() ([]byte, []int) {
	return fileDescriptor_c2bb9fc347232ae8, []int{3}
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

func (m *Adapter) GetType() PhyIoType {
	if m != nil {
		return m.Type
	}
	return PhyIoType_PhyIoNoop
}

func (m *Adapter) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type CipherInfo struct {
	Id                string            `protobuf:"bytes,100,opt,name=id,proto3" json:"id,omitempty"`
	KeyExchangeScheme KeyExchangeScheme `protobuf:"varint,1,opt,name=keyExchangeScheme,proto3,enum=KeyExchangeScheme" json:"keyExchangeScheme,omitempty"`
	EncryptionScheme  EncryptionScheme  `protobuf:"varint,2,opt,name=encryptionScheme,proto3,enum=EncryptionScheme" json:"encryptionScheme,omitempty"`
	// Initial Value for ECDH computation
	InitialValue []byte `protobuf:"bytes,3,opt,name=initialValue,proto3" json:"initialValue,omitempty"`
	// controller provided public certificate
	PublicCert []byte `protobuf:"bytes,4,opt,name=publicCert,proto3" json:"publicCert,omitempty"`
	// sha of the public certificate
	Sha256 string `protobuf:"bytes,5,opt,name=sha256,proto3" json:"sha256,omitempty"`
	// signature for the sha
	Signature            []byte   `protobuf:"bytes,6,opt,name=signature,proto3" json:"signature,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CipherInfo) Reset()         { *m = CipherInfo{} }
func (m *CipherInfo) String() string { return proto.CompactTextString(m) }
func (*CipherInfo) ProtoMessage()    {}
func (*CipherInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_c2bb9fc347232ae8, []int{4}
}

func (m *CipherInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CipherInfo.Unmarshal(m, b)
}
func (m *CipherInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CipherInfo.Marshal(b, m, deterministic)
}
func (m *CipherInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CipherInfo.Merge(m, src)
}
func (m *CipherInfo) XXX_Size() int {
	return xxx_messageInfo_CipherInfo.Size(m)
}
func (m *CipherInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_CipherInfo.DiscardUnknown(m)
}

var xxx_messageInfo_CipherInfo proto.InternalMessageInfo

func (m *CipherInfo) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *CipherInfo) GetKeyExchangeScheme() KeyExchangeScheme {
	if m != nil {
		return m.KeyExchangeScheme
	}
	return KeyExchangeScheme_KEA_NONE
}

func (m *CipherInfo) GetEncryptionScheme() EncryptionScheme {
	if m != nil {
		return m.EncryptionScheme
	}
	return EncryptionScheme_SA_NONE
}

func (m *CipherInfo) GetInitialValue() []byte {
	if m != nil {
		return m.InitialValue
	}
	return nil
}

func (m *CipherInfo) GetPublicCert() []byte {
	if m != nil {
		return m.PublicCert
	}
	return nil
}

func (m *CipherInfo) GetSha256() string {
	if m != nil {
		return m.Sha256
	}
	return ""
}

func (m *CipherInfo) GetSignature() []byte {
	if m != nil {
		return m.Signature
	}
	return nil
}

func init() {
	proto.RegisterEnum("KeyExchangeScheme", KeyExchangeScheme_name, KeyExchangeScheme_value)
	proto.RegisterEnum("EncryptionScheme", EncryptionScheme_name, EncryptionScheme_value)
	proto.RegisterType((*UUIDandVersion)(nil), "UUIDandVersion")
	proto.RegisterType((*DeviceOpsCmd)(nil), "DeviceOpsCmd")
	proto.RegisterType((*ConfigItem)(nil), "ConfigItem")
	proto.RegisterType((*Adapter)(nil), "Adapter")
	proto.RegisterType((*CipherInfo)(nil), "CipherInfo")
}

func init() { proto.RegisterFile("devcommon.proto", fileDescriptor_c2bb9fc347232ae8) }

var fileDescriptor_c2bb9fc347232ae8 = []byte{
	// 488 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x64, 0x52, 0xc1, 0x6e, 0xd3, 0x40,
	0x10, 0x25, 0x69, 0x9a, 0xb4, 0x43, 0x08, 0xce, 0x0a, 0x21, 0x0b, 0xa1, 0x52, 0x59, 0x1c, 0xaa,
	0x4a, 0xd8, 0x52, 0x4a, 0x7b, 0x2b, 0x22, 0x75, 0x0c, 0x44, 0x95, 0x5a, 0xe4, 0xb4, 0x3d, 0x70,
	0x89, 0x36, 0xbb, 0x13, 0x7b, 0x55, 0x7b, 0xd7, 0xb2, 0xd7, 0x16, 0xe6, 0x47, 0xf8, 0x5d, 0xe4,
	0x4d, 0x02, 0x34, 0xb9, 0xcd, 0x7b, 0xf3, 0xde, 0xec, 0xbe, 0xd1, 0xc0, 0x4b, 0x8e, 0x15, 0x53,
	0x69, 0xaa, 0xa4, 0x9b, 0xe5, 0x4a, 0xab, 0x37, 0x03, 0x8e, 0x55, 0xaa, 0x38, 0x26, 0x2b, 0xec,
	0x7c, 0x82, 0xc1, 0xfd, 0xfd, 0x74, 0x42, 0x25, 0x7f, 0xc0, 0xbc, 0x10, 0x4a, 0x12, 0x02, 0x9d,
	0xb2, 0x14, 0xdc, 0x6e, 0x1d, 0xb7, 0x4e, 0x0e, 0x43, 0x53, 0x13, 0x1b, 0x7a, 0xd5, 0xaa, 0x6d,
	0xb7, 0x0d, 0xbd, 0x81, 0xce, 0x12, 0xfa, 0x13, 0xac, 0x04, 0xc3, 0xdb, 0xac, 0xf0, 0x53, 0xa3,
	0x64, 0xaa, 0x94, 0x1a, 0x73, 0xa3, 0x7c, 0x11, 0x6e, 0x20, 0x71, 0xa0, 0xcf, 0xb1, 0x10, 0x39,
	0xf2, 0x99, 0xa6, 0x1a, 0xed, 0xbd, 0xe3, 0xd6, 0xc9, 0x41, 0xf8, 0x84, 0x6b, 0xdc, 0x2a, 0x2b,
	0xee, 0x44, 0x8a, 0x76, 0x67, 0xf5, 0xce, 0x1a, 0x3a, 0x1f, 0x01, 0x7c, 0x25, 0x97, 0x22, 0x9a,
	0x6a, 0x4c, 0x89, 0x05, 0x7b, 0x8f, 0x58, 0xaf, 0xbf, 0xd8, 0x94, 0xe4, 0x15, 0xec, 0x57, 0x34,
	0x29, 0x71, 0xfd, 0xbf, 0x15, 0x70, 0x2e, 0xa1, 0x37, 0xe6, 0x34, 0x6b, 0x9e, 0x3f, 0x82, 0x8e,
	0xae, 0x33, 0x34, 0x9e, 0xc1, 0x08, 0xdc, 0xef, 0x71, 0x3d, 0x55, 0x77, 0x75, 0x86, 0xa1, 0xe1,
	0x9b, 0xd8, 0x92, 0xa6, 0x1b, 0xbf, 0xa9, 0x9d, 0xdf, 0x6d, 0x00, 0x5f, 0x64, 0x31, 0xe6, 0x53,
	0xb9, 0x54, 0x64, 0x00, 0x6d, 0xc1, 0x6d, 0x6e, 0x04, 0x6d, 0xc1, 0xc9, 0x67, 0x18, 0x3e, 0x62,
	0x1d, 0xfc, 0x64, 0x31, 0x95, 0x11, 0xce, 0x58, 0x8c, 0xe9, 0x66, 0x3e, 0x71, 0xaf, 0xb7, 0x3b,
	0xe1, 0xae, 0x98, 0x5c, 0x82, 0x85, 0x92, 0xe5, 0x75, 0xa6, 0x85, 0x92, 0xeb, 0x01, 0x6d, 0x33,
	0x60, 0xe8, 0x06, 0x5b, 0x8d, 0x70, 0x47, 0xda, 0xac, 0x54, 0x48, 0xa1, 0x05, 0x4d, 0x1e, 0x4c,
	0xf6, 0x66, 0xa5, 0xfd, 0xf0, 0x09, 0x47, 0x8e, 0x00, 0xb2, 0x72, 0x91, 0x08, 0xe6, 0x63, 0xae,
	0xcd, 0x56, 0xfb, 0xe1, 0x7f, 0x0c, 0x79, 0x0d, 0xdd, 0x22, 0xa6, 0xa3, 0xf3, 0x0b, 0x7b, 0xdf,
	0x04, 0x5b, 0x23, 0xf2, 0x16, 0x0e, 0x0b, 0x11, 0x49, 0xaa, 0xcb, 0x1c, 0xed, 0xae, 0xb1, 0xfd,
	0x23, 0x4e, 0x3d, 0x18, 0xee, 0x04, 0x24, 0x7d, 0x38, 0xb8, 0x0e, 0xc6, 0xf3, 0x9b, 0xdb, 0x9b,
	0xc0, 0x7a, 0xb6, 0x41, 0x81, 0x3f, 0xf9, 0x66, 0xb5, 0x4e, 0xcf, 0xc0, 0xda, 0x0e, 0x44, 0x9e,
	0x43, 0x6f, 0xf6, 0x57, 0x4e, 0x60, 0x30, 0x1b, 0xcf, 0xc7, 0xc1, 0x6c, 0x3e, 0x3a, 0xbf, 0x98,
	0xfb, 0x5f, 0xae, 0xac, 0xd6, 0xd5, 0x57, 0x78, 0xc7, 0x54, 0xea, 0xfe, 0x42, 0x8e, 0x9c, 0xba,
	0x2c, 0x51, 0x25, 0x77, 0xcb, 0x02, 0xf3, 0xe6, 0xde, 0x56, 0xf7, 0xfb, 0xe3, 0x7d, 0x24, 0x74,
	0x5c, 0x2e, 0x5c, 0xa6, 0x52, 0x2f, 0x59, 0x7e, 0x40, 0x1e, 0xa1, 0x87, 0x15, 0x7a, 0x34, 0x13,
	0x5e, 0xa4, 0x3c, 0x66, 0x6e, 0x66, 0xd1, 0x35, 0xe2, 0xb3, 0x3f, 0x01, 0x00, 0x00, 0xff, 0xff,
	0xad, 0x41, 0x60, 0xeb, 0x0f, 0x03, 0x00, 0x00,
}
