// Code generated by protoc-gen-go. DO NOT EDIT.
// source: cipherinfo.proto

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
	return fileDescriptor_d32c1c7154980027, []int{0}
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
	return fileDescriptor_d32c1c7154980027, []int{1}
}

// Cipher information to decrypt Sensitive Data
type CipherContext struct {
	Uuidandversion    *UUIDandVersion   `protobuf:"bytes,100,opt,name=uuidandversion,proto3" json:"uuidandversion,omitempty"`
	KeyExchangeScheme KeyExchangeScheme `protobuf:"varint,1,opt,name=keyExchangeScheme,proto3,enum=KeyExchangeScheme" json:"keyExchangeScheme,omitempty"`
	EncryptionScheme  EncryptionScheme  `protobuf:"varint,2,opt,name=encryptionScheme,proto3,enum=EncryptionScheme" json:"encryptionScheme,omitempty"`
	// Initial Value for Symmetric Key derivation
	InitialValue []byte `protobuf:"bytes,3,opt,name=initialValue,proto3" json:"initialValue,omitempty"`
	// controller certificate
	ControllerCert []byte `protobuf:"bytes,4,opt,name=controllerCert,proto3" json:"controllerCert,omitempty"`
	// device certificate identifier/sha
	DeviceCertId         string   `protobuf:"bytes,5,opt,name=deviceCertId,proto3" json:"deviceCertId,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CipherContext) Reset()         { *m = CipherContext{} }
func (m *CipherContext) String() string { return proto.CompactTextString(m) }
func (*CipherContext) ProtoMessage()    {}
func (*CipherContext) Descriptor() ([]byte, []int) {
	return fileDescriptor_d32c1c7154980027, []int{0}
}

func (m *CipherContext) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CipherContext.Unmarshal(m, b)
}
func (m *CipherContext) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CipherContext.Marshal(b, m, deterministic)
}
func (m *CipherContext) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CipherContext.Merge(m, src)
}
func (m *CipherContext) XXX_Size() int {
	return xxx_messageInfo_CipherContext.Size(m)
}
func (m *CipherContext) XXX_DiscardUnknown() {
	xxx_messageInfo_CipherContext.DiscardUnknown(m)
}

var xxx_messageInfo_CipherContext proto.InternalMessageInfo

func (m *CipherContext) GetUuidandversion() *UUIDandVersion {
	if m != nil {
		return m.Uuidandversion
	}
	return nil
}

func (m *CipherContext) GetKeyExchangeScheme() KeyExchangeScheme {
	if m != nil {
		return m.KeyExchangeScheme
	}
	return KeyExchangeScheme_KEA_NONE
}

func (m *CipherContext) GetEncryptionScheme() EncryptionScheme {
	if m != nil {
		return m.EncryptionScheme
	}
	return EncryptionScheme_SA_NONE
}

func (m *CipherContext) GetInitialValue() []byte {
	if m != nil {
		return m.InitialValue
	}
	return nil
}

func (m *CipherContext) GetControllerCert() []byte {
	if m != nil {
		return m.ControllerCert
	}
	return nil
}

func (m *CipherContext) GetDeviceCertId() string {
	if m != nil {
		return m.DeviceCertId
	}
	return ""
}

// Encrypted sensitive data information
type CipherBlock struct {
	// cipher context id
	CipherContextId *UUIDandVersion `protobuf:"bytes,1,opt,name=cipherContextId,proto3" json:"cipherContextId,omitempty"`
	// encrypted sensitive data
	CipherData []byte `protobuf:"bytes,2,opt,name=cipherData,proto3" json:"cipherData,omitempty"`
	// sha256 of the plaintext sensitive data
	Sha256               []byte   `protobuf:"bytes,3,opt,name=sha256,proto3" json:"sha256,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CipherBlock) Reset()         { *m = CipherBlock{} }
func (m *CipherBlock) String() string { return proto.CompactTextString(m) }
func (*CipherBlock) ProtoMessage()    {}
func (*CipherBlock) Descriptor() ([]byte, []int) {
	return fileDescriptor_d32c1c7154980027, []int{1}
}

func (m *CipherBlock) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CipherBlock.Unmarshal(m, b)
}
func (m *CipherBlock) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CipherBlock.Marshal(b, m, deterministic)
}
func (m *CipherBlock) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CipherBlock.Merge(m, src)
}
func (m *CipherBlock) XXX_Size() int {
	return xxx_messageInfo_CipherBlock.Size(m)
}
func (m *CipherBlock) XXX_DiscardUnknown() {
	xxx_messageInfo_CipherBlock.DiscardUnknown(m)
}

var xxx_messageInfo_CipherBlock proto.InternalMessageInfo

func (m *CipherBlock) GetCipherContextId() *UUIDandVersion {
	if m != nil {
		return m.CipherContextId
	}
	return nil
}

func (m *CipherBlock) GetCipherData() []byte {
	if m != nil {
		return m.CipherData
	}
	return nil
}

func (m *CipherBlock) GetSha256() []byte {
	if m != nil {
		return m.Sha256
	}
	return nil
}

// This message will be filled with the
// credential details and will be encrypted
// across wire for data in transit
type CredentialBlock struct {
	Identity             string   `protobuf:"bytes,1,opt,name=identity,proto3" json:"identity,omitempty"`
	Password             string   `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CredentialBlock) Reset()         { *m = CredentialBlock{} }
func (m *CredentialBlock) String() string { return proto.CompactTextString(m) }
func (*CredentialBlock) ProtoMessage()    {}
func (*CredentialBlock) Descriptor() ([]byte, []int) {
	return fileDescriptor_d32c1c7154980027, []int{2}
}

func (m *CredentialBlock) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CredentialBlock.Unmarshal(m, b)
}
func (m *CredentialBlock) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CredentialBlock.Marshal(b, m, deterministic)
}
func (m *CredentialBlock) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CredentialBlock.Merge(m, src)
}
func (m *CredentialBlock) XXX_Size() int {
	return xxx_messageInfo_CredentialBlock.Size(m)
}
func (m *CredentialBlock) XXX_DiscardUnknown() {
	xxx_messageInfo_CredentialBlock.DiscardUnknown(m)
}

var xxx_messageInfo_CredentialBlock proto.InternalMessageInfo

func (m *CredentialBlock) GetIdentity() string {
	if m != nil {
		return m.Identity
	}
	return ""
}

func (m *CredentialBlock) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

func init() {
	proto.RegisterEnum("KeyExchangeScheme", KeyExchangeScheme_name, KeyExchangeScheme_value)
	proto.RegisterEnum("EncryptionScheme", EncryptionScheme_name, EncryptionScheme_value)
	proto.RegisterType((*CipherContext)(nil), "CipherContext")
	proto.RegisterType((*CipherBlock)(nil), "CipherBlock")
	proto.RegisterType((*CredentialBlock)(nil), "CredentialBlock")
}

func init() { proto.RegisterFile("cipherinfo.proto", fileDescriptor_d32c1c7154980027) }

var fileDescriptor_d32c1c7154980027 = []byte{
	// 444 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x92, 0xc1, 0x6e, 0xd3, 0x40,
	0x10, 0x86, 0x71, 0x81, 0xd2, 0x4c, 0x42, 0xe2, 0xec, 0x01, 0x59, 0x3d, 0x40, 0x14, 0x21, 0x14,
	0x55, 0xc2, 0x96, 0x52, 0xb5, 0x88, 0x03, 0x12, 0x89, 0x63, 0x20, 0xaa, 0x54, 0x24, 0x47, 0xed,
	0x81, 0x4b, 0xb4, 0xdd, 0x9d, 0xd8, 0xab, 0xda, 0xbb, 0xd6, 0x7a, 0x1d, 0x1a, 0x4e, 0xbc, 0x16,
	0x6f, 0x87, 0x6c, 0x87, 0xa8, 0x71, 0x38, 0xfe, 0xdf, 0x7c, 0x9a, 0xf5, 0xcc, 0x18, 0x6c, 0x26,
	0xb2, 0x18, 0xb5, 0x90, 0x2b, 0xe5, 0x66, 0x5a, 0x19, 0x75, 0xda, 0xe3, 0xb8, 0x66, 0x2a, 0x4d,
	0x95, 0xac, 0xc1, 0xf0, 0xcf, 0x11, 0xbc, 0xf4, 0x2b, 0xcb, 0x57, 0xd2, 0xe0, 0x83, 0x21, 0x1f,
	0xa0, 0x5b, 0x14, 0x82, 0x53, 0xc9, 0xd7, 0xa8, 0x73, 0xa1, 0xa4, 0xc3, 0x07, 0xd6, 0xa8, 0x3d,
	0xee, 0xb9, 0x37, 0x37, 0xf3, 0x19, 0x95, 0xfc, 0xb6, 0xc6, 0x61, 0x43, 0x23, 0x9f, 0xa1, 0x7f,
	0x8f, 0x9b, 0xe0, 0x81, 0xc5, 0x54, 0x46, 0xb8, 0x60, 0x31, 0xa6, 0xe8, 0x58, 0x03, 0x6b, 0xd4,
	0x1d, 0x13, 0xf7, 0xaa, 0x59, 0x09, 0x0f, 0x65, 0xf2, 0x09, 0x6c, 0x94, 0x4c, 0x6f, 0x32, 0x23,
	0x94, 0xdc, 0x36, 0x38, 0xaa, 0x1a, 0xf4, 0xdd, 0xa0, 0x51, 0x08, 0x0f, 0x54, 0x32, 0x84, 0x8e,
	0x90, 0xc2, 0x08, 0x9a, 0xdc, 0xd2, 0xa4, 0x40, 0xe7, 0xe9, 0xc0, 0x1a, 0x75, 0xc2, 0x3d, 0x46,
	0xde, 0x41, 0x97, 0x29, 0x69, 0xb4, 0x4a, 0x12, 0xd4, 0x3e, 0x6a, 0xe3, 0x3c, 0xab, 0xac, 0x06,
	0x2d, 0x7b, 0x71, 0x5c, 0x0b, 0x86, 0x65, 0x9a, 0x73, 0xe7, 0xf9, 0xc0, 0x1a, 0xb5, 0xc2, 0x3d,
	0x36, 0xfc, 0x6d, 0x41, 0xbb, 0xde, 0xdd, 0x34, 0x51, 0xec, 0x9e, 0x7c, 0x84, 0x1e, 0x7b, 0xbc,
	0xca, 0x39, 0xaf, 0xc6, 0xff, 0xcf, 0xea, 0x9a, 0x1e, 0x79, 0x0d, 0x50, 0xa3, 0x19, 0x35, 0xb4,
	0x9a, 0xb9, 0x13, 0x3e, 0x22, 0xe4, 0x15, 0x1c, 0xe7, 0x31, 0x1d, 0x5f, 0x5c, 0x6e, 0x87, 0xda,
	0xa6, 0xe1, 0x1c, 0x7a, 0xbe, 0x46, 0x8e, 0xb2, 0x9c, 0xb0, 0xfe, 0x8a, 0x53, 0x38, 0x11, 0x15,
	0x30, 0x9b, 0xea, 0xf9, 0x56, 0xb8, 0xcb, 0x65, 0x2d, 0xa3, 0x79, 0xfe, 0x53, 0x69, 0x5e, 0x3d,
	0xd2, 0x0a, 0x77, 0xf9, 0xcc, 0x83, 0xfe, 0xc1, 0x91, 0x48, 0x07, 0x4e, 0xae, 0x82, 0xc9, 0xf2,
	0xfa, 0xfb, 0x75, 0x60, 0x3f, 0xf9, 0x97, 0x02, 0x7f, 0xf6, 0xcd, 0xb6, 0xce, 0xce, 0xc1, 0x6e,
	0x1e, 0x85, 0xb4, 0xe1, 0xc5, 0x62, 0xa7, 0x13, 0xe8, 0x2e, 0x26, 0xcb, 0x49, 0xb0, 0x58, 0x8e,
	0x2f, 0x2e, 0x97, 0xfe, 0x97, 0xa9, 0x6d, 0x4d, 0xbf, 0xc2, 0x1b, 0xa6, 0x52, 0xf7, 0x17, 0x72,
	0xe4, 0xd4, 0x65, 0x89, 0x2a, 0xb8, 0x5b, 0xe4, 0xa8, 0xcb, 0xb5, 0xd6, 0xbf, 0xe4, 0x8f, 0xb7,
	0x91, 0x30, 0x71, 0x71, 0xe7, 0x32, 0x95, 0x7a, 0xc9, 0xea, 0x3d, 0xf2, 0x08, 0x3d, 0x5c, 0xa3,
	0x47, 0x33, 0xe1, 0x45, 0xca, 0x63, 0x4a, 0xae, 0x44, 0x74, 0x77, 0x5c, 0xc9, 0xe7, 0x7f, 0x03,
	0x00, 0x00, 0xff, 0xff, 0x19, 0xea, 0xb7, 0x47, 0xe4, 0x02, 0x00, 0x00,
}
