// Code generated by protoc-gen-go. DO NOT EDIT.
// source: register.proto

package register

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

// XXX is this used? Not in client.go; deprecate
type ZRegisterResult int32

const (
	ZRegisterResult_ZRegNone        ZRegisterResult = 0
	ZRegisterResult_ZRegSuccess     ZRegisterResult = 1
	ZRegisterResult_ZRegNotActive   ZRegisterResult = 2
	ZRegisterResult_ZRegAlreadyDone ZRegisterResult = 3
	ZRegisterResult_ZRegDeviceNA    ZRegisterResult = 4
	ZRegisterResult_ZRegFailed      ZRegisterResult = 5
)

var ZRegisterResult_name = map[int32]string{
	0: "ZRegNone",
	1: "ZRegSuccess",
	2: "ZRegNotActive",
	3: "ZRegAlreadyDone",
	4: "ZRegDeviceNA",
	5: "ZRegFailed",
}

var ZRegisterResult_value = map[string]int32{
	"ZRegNone":        0,
	"ZRegSuccess":     1,
	"ZRegNotActive":   2,
	"ZRegAlreadyDone": 3,
	"ZRegDeviceNA":    4,
	"ZRegFailed":      5,
}

func (x ZRegisterResult) String() string {
	return proto.EnumName(ZRegisterResult_name, int32(x))
}

func (ZRegisterResult) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_1303fe8288f4efb6, []int{0}
}

// This is the request payload for POST /api/v1/edgeDevice/register
// ZRegisterMsg carries the pem-encoded device certificate plus additional
// identifying information such as device serial number(s).
// The message is assumed to be protected by a TLS session bound to the
// onboarding certificate.
type ZRegisterMsg struct {
	OnBoardKey           string   `protobuf:"bytes,1,opt,name=onBoardKey,proto3" json:"onBoardKey,omitempty"`
	PemCert              []byte   `protobuf:"bytes,2,opt,name=pemCert,proto3" json:"pemCert,omitempty"`
	Serial               string   `protobuf:"bytes,3,opt,name=serial,proto3" json:"serial,omitempty"`
	SoftSerial           string   `protobuf:"bytes,4,opt,name=softSerial,proto3" json:"softSerial,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ZRegisterMsg) Reset()         { *m = ZRegisterMsg{} }
func (m *ZRegisterMsg) String() string { return proto.CompactTextString(m) }
func (*ZRegisterMsg) ProtoMessage()    {}
func (*ZRegisterMsg) Descriptor() ([]byte, []int) {
	return fileDescriptor_1303fe8288f4efb6, []int{0}
}

func (m *ZRegisterMsg) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ZRegisterMsg.Unmarshal(m, b)
}
func (m *ZRegisterMsg) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ZRegisterMsg.Marshal(b, m, deterministic)
}
func (m *ZRegisterMsg) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ZRegisterMsg.Merge(m, src)
}
func (m *ZRegisterMsg) XXX_Size() int {
	return xxx_messageInfo_ZRegisterMsg.Size(m)
}
func (m *ZRegisterMsg) XXX_DiscardUnknown() {
	xxx_messageInfo_ZRegisterMsg.DiscardUnknown(m)
}

var xxx_messageInfo_ZRegisterMsg proto.InternalMessageInfo

func (m *ZRegisterMsg) GetOnBoardKey() string {
	if m != nil {
		return m.OnBoardKey
	}
	return ""
}

func (m *ZRegisterMsg) GetPemCert() []byte {
	if m != nil {
		return m.PemCert
	}
	return nil
}

func (m *ZRegisterMsg) GetSerial() string {
	if m != nil {
		return m.Serial
	}
	return ""
}

func (m *ZRegisterMsg) GetSoftSerial() string {
	if m != nil {
		return m.SoftSerial
	}
	return ""
}

// XXX is this used? Not in client.go; deprecate
type ZRegisterResp struct {
	Result               ZRegisterResult `protobuf:"varint,2,opt,name=result,proto3,enum=ZRegisterResult" json:"result,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *ZRegisterResp) Reset()         { *m = ZRegisterResp{} }
func (m *ZRegisterResp) String() string { return proto.CompactTextString(m) }
func (*ZRegisterResp) ProtoMessage()    {}
func (*ZRegisterResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_1303fe8288f4efb6, []int{1}
}

func (m *ZRegisterResp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ZRegisterResp.Unmarshal(m, b)
}
func (m *ZRegisterResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ZRegisterResp.Marshal(b, m, deterministic)
}
func (m *ZRegisterResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ZRegisterResp.Merge(m, src)
}
func (m *ZRegisterResp) XXX_Size() int {
	return xxx_messageInfo_ZRegisterResp.Size(m)
}
func (m *ZRegisterResp) XXX_DiscardUnknown() {
	xxx_messageInfo_ZRegisterResp.DiscardUnknown(m)
}

var xxx_messageInfo_ZRegisterResp proto.InternalMessageInfo

func (m *ZRegisterResp) GetResult() ZRegisterResult {
	if m != nil {
		return m.Result
	}
	return ZRegisterResult_ZRegNone
}

func init() {
	proto.RegisterEnum("ZRegisterResult", ZRegisterResult_name, ZRegisterResult_value)
	proto.RegisterType((*ZRegisterMsg)(nil), "ZRegisterMsg")
	proto.RegisterType((*ZRegisterResp)(nil), "ZRegisterResp")
}

func init() { proto.RegisterFile("register.proto", fileDescriptor_1303fe8288f4efb6) }

var fileDescriptor_1303fe8288f4efb6 = []byte{
	// 275 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x5c, 0x90, 0x41, 0x4b, 0xc3, 0x40,
	0x10, 0x85, 0x4d, 0x5b, 0xa3, 0x8e, 0x31, 0x5d, 0x47, 0x90, 0x9c, 0xa4, 0xf4, 0x20, 0x41, 0x30,
	0x01, 0x3d, 0x79, 0x4c, 0x2d, 0x5e, 0xc4, 0x1c, 0xd2, 0x5b, 0x6f, 0x69, 0x32, 0x8d, 0x0b, 0xdb,
	0x6e, 0xd8, 0xdd, 0x04, 0xea, 0xc9, 0x9f, 0x2e, 0xe9, 0xa6, 0x5a, 0x3c, 0xbe, 0xef, 0x31, 0xef,
	0x0d, 0x0f, 0x7c, 0x45, 0x15, 0xd7, 0x86, 0x54, 0x54, 0x2b, 0x69, 0xe4, 0xf4, 0xdb, 0x01, 0x6f,
	0x99, 0xf5, 0xec, 0x43, 0x57, 0x78, 0x07, 0x20, 0xb7, 0x33, 0x99, 0xab, 0xf2, 0x9d, 0x76, 0x81,
	0x33, 0x71, 0xc2, 0x8b, 0xec, 0x88, 0x60, 0x00, 0x67, 0x35, 0x6d, 0x5e, 0x49, 0x99, 0x60, 0x30,
	0x71, 0x42, 0x2f, 0x3b, 0x48, 0xbc, 0x05, 0x57, 0x93, 0xe2, 0xb9, 0x08, 0x86, 0xfb, 0xab, 0x5e,
	0x75, 0x89, 0x5a, 0xae, 0xcd, 0xc2, 0x7a, 0x23, 0x9b, 0xf8, 0x47, 0xa6, 0x2f, 0x70, 0xf5, 0xfb,
	0x41, 0x46, 0xba, 0xc6, 0x10, 0x5c, 0x45, 0xba, 0x11, 0xb6, 0xc1, 0x7f, 0x62, 0xd1, 0xb1, 0xdf,
	0x08, 0x93, 0xf5, 0xfe, 0xc3, 0x17, 0x8c, 0xff, 0x59, 0xe8, 0xc1, 0x79, 0x87, 0x52, 0xb9, 0x25,
	0x76, 0x82, 0x63, 0xb8, 0xec, 0xd4, 0xa2, 0x29, 0x0a, 0xd2, 0x9a, 0x39, 0x78, 0x6d, 0xcb, 0x52,
	0x69, 0x92, 0xc2, 0xf0, 0x96, 0xd8, 0x00, 0x6f, 0x6c, 0x48, 0x22, 0x14, 0xe5, 0xe5, 0x6e, 0xde,
	0x1d, 0x0e, 0x91, 0xd9, 0x59, 0xe6, 0xd4, 0xf2, 0x82, 0xd2, 0x84, 0x8d, 0xd0, 0x07, 0xe8, 0xc8,
	0x5b, 0xce, 0x05, 0x95, 0xec, 0x74, 0x16, 0x2e, 0xef, 0x2b, 0x6e, 0x3e, 0x9b, 0x55, 0x54, 0xc8,
	0x4d, 0x2c, 0xd6, 0x8f, 0x54, 0x56, 0x14, 0x53, 0x4b, 0x71, 0x5e, 0xf3, 0xb8, 0x92, 0xf1, 0x61,
	0xe9, 0x95, 0xbb, 0x9f, 0xfa, 0xf9, 0x27, 0x00, 0x00, 0xff, 0xff, 0xbb, 0x0a, 0xfc, 0x48, 0x7c,
	0x01, 0x00, 0x00,
}
