// Copyright(c) 2020 Zededa, Inc.
// All rights reserved.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.24.0-devel
// 	protoc        v3.6.1
// source: attest/attest.proto

package attest

import (
	proto "github.com/golang/protobuf/proto"
	evecommon "github.com/lf-edge/eve/api/go/evecommon"
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

type ZAttestReqType int32

const (
	ZAttestReqType_ATTEST_REQ_NONE  ZAttestReqType = 0
	ZAttestReqType_ATTEST_REQ_CERT  ZAttestReqType = 1 //EVE X.509 certificates
	ZAttestReqType_ATTEST_REQ_NONCE ZAttestReqType = 2 //nonce request to Controller
	ZAttestReqType_ATTEST_REQ_QUOTE ZAttestReqType = 3 //quote msg
)

// Enum value maps for ZAttestReqType.
var (
	ZAttestReqType_name = map[int32]string{
		0: "ATTEST_REQ_NONE",
		1: "ATTEST_REQ_CERT",
		2: "ATTEST_REQ_NONCE",
		3: "ATTEST_REQ_QUOTE",
	}
	ZAttestReqType_value = map[string]int32{
		"ATTEST_REQ_NONE":  0,
		"ATTEST_REQ_CERT":  1,
		"ATTEST_REQ_NONCE": 2,
		"ATTEST_REQ_QUOTE": 3,
	}
)

func (x ZAttestReqType) Enum() *ZAttestReqType {
	p := new(ZAttestReqType)
	*p = x
	return p
}

func (x ZAttestReqType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ZAttestReqType) Descriptor() protoreflect.EnumDescriptor {
	return file_attest_attest_proto_enumTypes[0].Descriptor()
}

func (ZAttestReqType) Type() protoreflect.EnumType {
	return &file_attest_attest_proto_enumTypes[0]
}

func (x ZAttestReqType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ZAttestReqType.Descriptor instead.
func (ZAttestReqType) EnumDescriptor() ([]byte, []int) {
	return file_attest_attest_proto_rawDescGZIP(), []int{0}
}

type ZAttestRespType int32

const (
	ZAttestRespType_ATTEST_RESP_NONE       ZAttestRespType = 0
	ZAttestRespType_ATTEST_RESP_CERT       ZAttestRespType = 1 //response to cert msg
	ZAttestRespType_ATTEST_RESP_NONCE      ZAttestRespType = 2 //response to quote request
	ZAttestRespType_ATTEST_RESP_QUOTE_RESP ZAttestRespType = 3 //response to quote msg
)

// Enum value maps for ZAttestRespType.
var (
	ZAttestRespType_name = map[int32]string{
		0: "ATTEST_RESP_NONE",
		1: "ATTEST_RESP_CERT",
		2: "ATTEST_RESP_NONCE",
		3: "ATTEST_RESP_QUOTE_RESP",
	}
	ZAttestRespType_value = map[string]int32{
		"ATTEST_RESP_NONE":       0,
		"ATTEST_RESP_CERT":       1,
		"ATTEST_RESP_NONCE":      2,
		"ATTEST_RESP_QUOTE_RESP": 3,
	}
)

func (x ZAttestRespType) Enum() *ZAttestRespType {
	p := new(ZAttestRespType)
	*p = x
	return p
}

func (x ZAttestRespType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ZAttestRespType) Descriptor() protoreflect.EnumDescriptor {
	return file_attest_attest_proto_enumTypes[1].Descriptor()
}

func (ZAttestRespType) Type() protoreflect.EnumType {
	return &file_attest_attest_proto_enumTypes[1]
}

func (x ZAttestRespType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ZAttestRespType.Descriptor instead.
func (ZAttestRespType) EnumDescriptor() ([]byte, []int) {
	return file_attest_attest_proto_rawDescGZIP(), []int{1}
}

type ZAttestResponseCode int32

const (
	ZAttestResponseCode_ATTEST_RESPONSE_NONE    ZAttestResponseCode = 0
	ZAttestResponseCode_ATTEST_RESPONSE_SUCCESS ZAttestResponseCode = 1 //Attestation successful
	ZAttestResponseCode_ATTEST_RESPONSE_FAILURE ZAttestResponseCode = 2 //Attestation failed
)

// Enum value maps for ZAttestResponseCode.
var (
	ZAttestResponseCode_name = map[int32]string{
		0: "ATTEST_RESPONSE_NONE",
		1: "ATTEST_RESPONSE_SUCCESS",
		2: "ATTEST_RESPONSE_FAILURE",
	}
	ZAttestResponseCode_value = map[string]int32{
		"ATTEST_RESPONSE_NONE":    0,
		"ATTEST_RESPONSE_SUCCESS": 1,
		"ATTEST_RESPONSE_FAILURE": 2,
	}
)

func (x ZAttestResponseCode) Enum() *ZAttestResponseCode {
	p := new(ZAttestResponseCode)
	*p = x
	return p
}

func (x ZAttestResponseCode) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ZAttestResponseCode) Descriptor() protoreflect.EnumDescriptor {
	return file_attest_attest_proto_enumTypes[2].Descriptor()
}

func (ZAttestResponseCode) Type() protoreflect.EnumType {
	return &file_attest_attest_proto_enumTypes[2]
}

func (x ZAttestResponseCode) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ZAttestResponseCode.Descriptor instead.
func (ZAttestResponseCode) EnumDescriptor() ([]byte, []int) {
	return file_attest_attest_proto_rawDescGZIP(), []int{2}
}

type ZEveCertType int32

const (
	ZEveCertType_CERT_TYPE_DEVICE_NONE               ZEveCertType = 0
	ZEveCertType_CERT_TYPE_DEVICE_ONBOARDING         ZEveCertType = 1 //set for certificate used by edge node for identifying the device
	ZEveCertType_CERT_TYPE_DEVICE_RESTRICTED_SIGNING ZEveCertType = 2 //set for certificate used by edge node for attestation
	ZEveCertType_CERT_TYPE_DEVICE_ENDORSEMENT_RSA    ZEveCertType = 3 //set for endorsement key certificate with RSASSA signing algorithm
	ZEveCertType_CERT_TYPE_DEVICE_ECDH_EXCHANGE      ZEveCertType = 4 //set for certificate used by edge node to share symmetric key using ECDH
)

// Enum value maps for ZEveCertType.
var (
	ZEveCertType_name = map[int32]string{
		0: "CERT_TYPE_DEVICE_NONE",
		1: "CERT_TYPE_DEVICE_ONBOARDING",
		2: "CERT_TYPE_DEVICE_RESTRICTED_SIGNING",
		3: "CERT_TYPE_DEVICE_ENDORSEMENT_RSA",
		4: "CERT_TYPE_DEVICE_ECDH_EXCHANGE",
	}
	ZEveCertType_value = map[string]int32{
		"CERT_TYPE_DEVICE_NONE":               0,
		"CERT_TYPE_DEVICE_ONBOARDING":         1,
		"CERT_TYPE_DEVICE_RESTRICTED_SIGNING": 2,
		"CERT_TYPE_DEVICE_ENDORSEMENT_RSA":    3,
		"CERT_TYPE_DEVICE_ECDH_EXCHANGE":      4,
	}
)

func (x ZEveCertType) Enum() *ZEveCertType {
	p := new(ZEveCertType)
	*p = x
	return p
}

func (x ZEveCertType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ZEveCertType) Descriptor() protoreflect.EnumDescriptor {
	return file_attest_attest_proto_enumTypes[3].Descriptor()
}

func (ZEveCertType) Type() protoreflect.EnumType {
	return &file_attest_attest_proto_enumTypes[3]
}

func (x ZEveCertType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ZEveCertType.Descriptor instead.
func (ZEveCertType) EnumDescriptor() ([]byte, []int) {
	return file_attest_attest_proto_rawDescGZIP(), []int{3}
}

// This is the request payload for POST /api/v2/edgeDevice/id/<uuid>/attest
// The message is assumed to be protected by signing envelope
type ZAttestReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ReqType ZAttestReqType `protobuf:"varint,1,opt,name=reqType,proto3,enum=ZAttestReqType" json:"reqType,omitempty"` //type of the request
	Quote   *ZAttestQuote  `protobuf:"bytes,2,opt,name=quote,proto3" json:"quote,omitempty"`                          //attestation quote msg
	Certs   []*ZEveCert    `protobuf:"bytes,3,rep,name=certs,proto3" json:"certs,omitempty"`                          //X509 certs in .PEM format, signed by device certificate
}

func (x *ZAttestReq) Reset() {
	*x = ZAttestReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_attest_attest_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ZAttestReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ZAttestReq) ProtoMessage() {}

func (x *ZAttestReq) ProtoReflect() protoreflect.Message {
	mi := &file_attest_attest_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ZAttestReq.ProtoReflect.Descriptor instead.
func (*ZAttestReq) Descriptor() ([]byte, []int) {
	return file_attest_attest_proto_rawDescGZIP(), []int{0}
}

func (x *ZAttestReq) GetReqType() ZAttestReqType {
	if x != nil {
		return x.ReqType
	}
	return ZAttestReqType_ATTEST_REQ_NONE
}

func (x *ZAttestReq) GetQuote() *ZAttestQuote {
	if x != nil {
		return x.Quote
	}
	return nil
}

func (x *ZAttestReq) GetCerts() []*ZEveCert {
	if x != nil {
		return x.Certs
	}
	return nil
}

// This is the response payload for POST /api/v2/edgeDevice/id/<uuid>/attest
// The message is assumed to be protected by signing envelope
type ZAttestResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RespType  ZAttestRespType   `protobuf:"varint,1,opt,name=respType,proto3,enum=ZAttestRespType" json:"respType,omitempty"` //type of the response
	Nonce     *ZAttestNonceResp `protobuf:"bytes,2,opt,name=nonce,proto3" json:"nonce,omitempty"`                             //nonce from Controller
	QuoteResp *ZAttestQuoteResp `protobuf:"bytes,3,opt,name=quoteResp,proto3" json:"quoteResp,omitempty"`                     //attest quote response from Controller
}

func (x *ZAttestResponse) Reset() {
	*x = ZAttestResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_attest_attest_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ZAttestResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ZAttestResponse) ProtoMessage() {}

func (x *ZAttestResponse) ProtoReflect() protoreflect.Message {
	mi := &file_attest_attest_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ZAttestResponse.ProtoReflect.Descriptor instead.
func (*ZAttestResponse) Descriptor() ([]byte, []int) {
	return file_attest_attest_proto_rawDescGZIP(), []int{1}
}

func (x *ZAttestResponse) GetRespType() ZAttestRespType {
	if x != nil {
		return x.RespType
	}
	return ZAttestRespType_ATTEST_RESP_NONE
}

func (x *ZAttestResponse) GetNonce() *ZAttestNonceResp {
	if x != nil {
		return x.Nonce
	}
	return nil
}

func (x *ZAttestResponse) GetQuoteResp() *ZAttestQuoteResp {
	if x != nil {
		return x.QuoteResp
	}
	return nil
}

type ZAttestNonceResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Nonce []byte `protobuf:"bytes,1,opt,name=nonce,proto3" json:"nonce,omitempty"` //nonce to use in quote generation
}

func (x *ZAttestNonceResp) Reset() {
	*x = ZAttestNonceResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_attest_attest_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ZAttestNonceResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ZAttestNonceResp) ProtoMessage() {}

func (x *ZAttestNonceResp) ProtoReflect() protoreflect.Message {
	mi := &file_attest_attest_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ZAttestNonceResp.ProtoReflect.Descriptor instead.
func (*ZAttestNonceResp) Descriptor() ([]byte, []int) {
	return file_attest_attest_proto_rawDescGZIP(), []int{2}
}

func (x *ZAttestNonceResp) GetNonce() []byte {
	if x != nil {
		return x.Nonce
	}
	return nil
}

type ZAttestQuote struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AttestData []byte `protobuf:"bytes,1,opt,name=attestData,proto3" json:"attestData,omitempty"` //TPMS_ATTEST (Table 2:123) in https://trustedcomputinggroup.org/wp-content/uploads/TPM-Rev-2.0-Part-2-Structures-01.38.pdf
	//nonce is included in attestData
	Signature []byte `protobuf:"bytes,2,opt,name=signature,proto3" json:"signature,omitempty"` //signature to verify attestData
}

func (x *ZAttestQuote) Reset() {
	*x = ZAttestQuote{}
	if protoimpl.UnsafeEnabled {
		mi := &file_attest_attest_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ZAttestQuote) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ZAttestQuote) ProtoMessage() {}

func (x *ZAttestQuote) ProtoReflect() protoreflect.Message {
	mi := &file_attest_attest_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ZAttestQuote.ProtoReflect.Descriptor instead.
func (*ZAttestQuote) Descriptor() ([]byte, []int) {
	return file_attest_attest_proto_rawDescGZIP(), []int{3}
}

func (x *ZAttestQuote) GetAttestData() []byte {
	if x != nil {
		return x.AttestData
	}
	return nil
}

func (x *ZAttestQuote) GetSignature() []byte {
	if x != nil {
		return x.Signature
	}
	return nil
}

type ZAttestQuoteResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Response ZAttestResponseCode `protobuf:"varint,1,opt,name=response,proto3,enum=ZAttestResponseCode" json:"response,omitempty"` //result of quote validation
}

func (x *ZAttestQuoteResp) Reset() {
	*x = ZAttestQuoteResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_attest_attest_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ZAttestQuoteResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ZAttestQuoteResp) ProtoMessage() {}

func (x *ZAttestQuoteResp) ProtoReflect() protoreflect.Message {
	mi := &file_attest_attest_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ZAttestQuoteResp.ProtoReflect.Descriptor instead.
func (*ZAttestQuoteResp) Descriptor() ([]byte, []int) {
	return file_attest_attest_proto_rawDescGZIP(), []int{4}
}

func (x *ZAttestQuoteResp) GetResponse() ZAttestResponseCode {
	if x != nil {
		return x.Response
	}
	return ZAttestResponseCode_ATTEST_RESPONSE_NONE
}

type ZEveCert struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	HashAlgo   evecommon.HashAlgorithm `protobuf:"varint,1,opt,name=hashAlgo,proto3,enum=org.lfedge.eve.common.HashAlgorithm" json:"hashAlgo,omitempty"` //hash method used to arrive at certHash
	CertHash   []byte                  `protobuf:"bytes,2,opt,name=certHash,proto3" json:"certHash,omitempty"`                                           //Hash of the cert, computed using hashAlgo
	Type       ZEveCertType            `protobuf:"varint,3,opt,name=type,proto3,enum=ZEveCertType" json:"type,omitempty"`                                //what kind of certificate(to identify the target use case)
	Cert       []byte                  `protobuf:"bytes,4,opt,name=cert,proto3" json:"cert,omitempty"`                                                   //X509 cert in .PEM format
	Attributes *ZEveCertAttr           `protobuf:"bytes,5,opt,name=attributes,proto3" json:"attributes,omitempty"`                                       //properties of this certificate
}

func (x *ZEveCert) Reset() {
	*x = ZEveCert{}
	if protoimpl.UnsafeEnabled {
		mi := &file_attest_attest_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ZEveCert) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ZEveCert) ProtoMessage() {}

func (x *ZEveCert) ProtoReflect() protoreflect.Message {
	mi := &file_attest_attest_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ZEveCert.ProtoReflect.Descriptor instead.
func (*ZEveCert) Descriptor() ([]byte, []int) {
	return file_attest_attest_proto_rawDescGZIP(), []int{5}
}

func (x *ZEveCert) GetHashAlgo() evecommon.HashAlgorithm {
	if x != nil {
		return x.HashAlgo
	}
	return evecommon.HashAlgorithm_HASH_ALGORITHM_INVALID
}

func (x *ZEveCert) GetCertHash() []byte {
	if x != nil {
		return x.CertHash
	}
	return nil
}

func (x *ZEveCert) GetType() ZEveCertType {
	if x != nil {
		return x.Type
	}
	return ZEveCertType_CERT_TYPE_DEVICE_NONE
}

func (x *ZEveCert) GetCert() []byte {
	if x != nil {
		return x.Cert
	}
	return nil
}

func (x *ZEveCert) GetAttributes() *ZEveCertAttr {
	if x != nil {
		return x.Attributes
	}
	return nil
}

type ZEveCertAttr struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	IsMutable bool `protobuf:"varint,1,opt,name=isMutable,proto3" json:"isMutable,omitempty"` //set to false for immutable certificates
}

func (x *ZEveCertAttr) Reset() {
	*x = ZEveCertAttr{}
	if protoimpl.UnsafeEnabled {
		mi := &file_attest_attest_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ZEveCertAttr) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ZEveCertAttr) ProtoMessage() {}

func (x *ZEveCertAttr) ProtoReflect() protoreflect.Message {
	mi := &file_attest_attest_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ZEveCertAttr.ProtoReflect.Descriptor instead.
func (*ZEveCertAttr) Descriptor() ([]byte, []int) {
	return file_attest_attest_proto_rawDescGZIP(), []int{6}
}

func (x *ZEveCertAttr) GetIsMutable() bool {
	if x != nil {
		return x.IsMutable
	}
	return false
}

var File_attest_attest_proto protoreflect.FileDescriptor

var file_attest_attest_proto_rawDesc = []byte{
	0x0a, 0x13, 0x61, 0x74, 0x74, 0x65, 0x73, 0x74, 0x2f, 0x61, 0x74, 0x74, 0x65, 0x73, 0x74, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x19, 0x65, 0x76, 0x65, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e,
	0x2f, 0x65, 0x76, 0x65, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0x7d, 0x0a, 0x0a, 0x5a, 0x41, 0x74, 0x74, 0x65, 0x73, 0x74, 0x52, 0x65, 0x71, 0x12, 0x29,
	0x0a, 0x07, 0x72, 0x65, 0x71, 0x54, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32,
	0x0f, 0x2e, 0x5a, 0x41, 0x74, 0x74, 0x65, 0x73, 0x74, 0x52, 0x65, 0x71, 0x54, 0x79, 0x70, 0x65,
	0x52, 0x07, 0x72, 0x65, 0x71, 0x54, 0x79, 0x70, 0x65, 0x12, 0x23, 0x0a, 0x05, 0x71, 0x75, 0x6f,
	0x74, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x5a, 0x41, 0x74, 0x74, 0x65,
	0x73, 0x74, 0x51, 0x75, 0x6f, 0x74, 0x65, 0x52, 0x05, 0x71, 0x75, 0x6f, 0x74, 0x65, 0x12, 0x1f,
	0x0a, 0x05, 0x63, 0x65, 0x72, 0x74, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x09, 0x2e,
	0x5a, 0x45, 0x76, 0x65, 0x43, 0x65, 0x72, 0x74, 0x52, 0x05, 0x63, 0x65, 0x72, 0x74, 0x73, 0x22,
	0x99, 0x01, 0x0a, 0x0f, 0x5a, 0x41, 0x74, 0x74, 0x65, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x2c, 0x0a, 0x08, 0x72, 0x65, 0x73, 0x70, 0x54, 0x79, 0x70, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x10, 0x2e, 0x5a, 0x41, 0x74, 0x74, 0x65, 0x73, 0x74, 0x52,
	0x65, 0x73, 0x70, 0x54, 0x79, 0x70, 0x65, 0x52, 0x08, 0x72, 0x65, 0x73, 0x70, 0x54, 0x79, 0x70,
	0x65, 0x12, 0x27, 0x0a, 0x05, 0x6e, 0x6f, 0x6e, 0x63, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x11, 0x2e, 0x5a, 0x41, 0x74, 0x74, 0x65, 0x73, 0x74, 0x4e, 0x6f, 0x6e, 0x63, 0x65, 0x52,
	0x65, 0x73, 0x70, 0x52, 0x05, 0x6e, 0x6f, 0x6e, 0x63, 0x65, 0x12, 0x2f, 0x0a, 0x09, 0x71, 0x75,
	0x6f, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e,
	0x5a, 0x41, 0x74, 0x74, 0x65, 0x73, 0x74, 0x51, 0x75, 0x6f, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70,
	0x52, 0x09, 0x71, 0x75, 0x6f, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x22, 0x28, 0x0a, 0x10, 0x5a,
	0x41, 0x74, 0x74, 0x65, 0x73, 0x74, 0x4e, 0x6f, 0x6e, 0x63, 0x65, 0x52, 0x65, 0x73, 0x70, 0x12,
	0x14, 0x0a, 0x05, 0x6e, 0x6f, 0x6e, 0x63, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x05,
	0x6e, 0x6f, 0x6e, 0x63, 0x65, 0x22, 0x4c, 0x0a, 0x0c, 0x5a, 0x41, 0x74, 0x74, 0x65, 0x73, 0x74,
	0x51, 0x75, 0x6f, 0x74, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x61, 0x74, 0x74, 0x65, 0x73, 0x74, 0x44,
	0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x0a, 0x61, 0x74, 0x74, 0x65, 0x73,
	0x74, 0x44, 0x61, 0x74, 0x61, 0x12, 0x1c, 0x0a, 0x09, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75,
	0x72, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x09, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x74,
	0x75, 0x72, 0x65, 0x22, 0x44, 0x0a, 0x10, 0x5a, 0x41, 0x74, 0x74, 0x65, 0x73, 0x74, 0x51, 0x75,
	0x6f, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x12, 0x30, 0x0a, 0x08, 0x72, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x14, 0x2e, 0x5a, 0x41, 0x74, 0x74,
	0x65, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x43, 0x6f, 0x64, 0x65, 0x52,
	0x08, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0xce, 0x01, 0x0a, 0x08, 0x5a, 0x45,
	0x76, 0x65, 0x43, 0x65, 0x72, 0x74, 0x12, 0x40, 0x0a, 0x08, 0x68, 0x61, 0x73, 0x68, 0x41, 0x6c,
	0x67, 0x6f, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x24, 0x2e, 0x6f, 0x72, 0x67, 0x2e, 0x6c,
	0x66, 0x65, 0x64, 0x67, 0x65, 0x2e, 0x65, 0x76, 0x65, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e,
	0x2e, 0x48, 0x61, 0x73, 0x68, 0x41, 0x6c, 0x67, 0x6f, 0x72, 0x69, 0x74, 0x68, 0x6d, 0x52, 0x08,
	0x68, 0x61, 0x73, 0x68, 0x41, 0x6c, 0x67, 0x6f, 0x12, 0x1a, 0x0a, 0x08, 0x63, 0x65, 0x72, 0x74,
	0x48, 0x61, 0x73, 0x68, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x08, 0x63, 0x65, 0x72, 0x74,
	0x48, 0x61, 0x73, 0x68, 0x12, 0x21, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x0e, 0x32, 0x0d, 0x2e, 0x5a, 0x45, 0x76, 0x65, 0x43, 0x65, 0x72, 0x74, 0x54, 0x79, 0x70,
	0x65, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x65, 0x72, 0x74, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x04, 0x63, 0x65, 0x72, 0x74, 0x12, 0x2d, 0x0a, 0x0a, 0x61,
	0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x0d, 0x2e, 0x5a, 0x45, 0x76, 0x65, 0x43, 0x65, 0x72, 0x74, 0x41, 0x74, 0x74, 0x72, 0x52, 0x0a,
	0x61, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x73, 0x22, 0x2c, 0x0a, 0x0c, 0x5a, 0x45,
	0x76, 0x65, 0x43, 0x65, 0x72, 0x74, 0x41, 0x74, 0x74, 0x72, 0x12, 0x1c, 0x0a, 0x09, 0x69, 0x73,
	0x4d, 0x75, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x09, 0x69,
	0x73, 0x4d, 0x75, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x2a, 0x66, 0x0a, 0x0e, 0x5a, 0x41, 0x74, 0x74,
	0x65, 0x73, 0x74, 0x52, 0x65, 0x71, 0x54, 0x79, 0x70, 0x65, 0x12, 0x13, 0x0a, 0x0f, 0x41, 0x54,
	0x54, 0x45, 0x53, 0x54, 0x5f, 0x52, 0x45, 0x51, 0x5f, 0x4e, 0x4f, 0x4e, 0x45, 0x10, 0x00, 0x12,
	0x13, 0x0a, 0x0f, 0x41, 0x54, 0x54, 0x45, 0x53, 0x54, 0x5f, 0x52, 0x45, 0x51, 0x5f, 0x43, 0x45,
	0x52, 0x54, 0x10, 0x01, 0x12, 0x14, 0x0a, 0x10, 0x41, 0x54, 0x54, 0x45, 0x53, 0x54, 0x5f, 0x52,
	0x45, 0x51, 0x5f, 0x4e, 0x4f, 0x4e, 0x43, 0x45, 0x10, 0x02, 0x12, 0x14, 0x0a, 0x10, 0x41, 0x54,
	0x54, 0x45, 0x53, 0x54, 0x5f, 0x52, 0x45, 0x51, 0x5f, 0x51, 0x55, 0x4f, 0x54, 0x45, 0x10, 0x03,
	0x2a, 0x70, 0x0a, 0x0f, 0x5a, 0x41, 0x74, 0x74, 0x65, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x54,
	0x79, 0x70, 0x65, 0x12, 0x14, 0x0a, 0x10, 0x41, 0x54, 0x54, 0x45, 0x53, 0x54, 0x5f, 0x52, 0x45,
	0x53, 0x50, 0x5f, 0x4e, 0x4f, 0x4e, 0x45, 0x10, 0x00, 0x12, 0x14, 0x0a, 0x10, 0x41, 0x54, 0x54,
	0x45, 0x53, 0x54, 0x5f, 0x52, 0x45, 0x53, 0x50, 0x5f, 0x43, 0x45, 0x52, 0x54, 0x10, 0x01, 0x12,
	0x15, 0x0a, 0x11, 0x41, 0x54, 0x54, 0x45, 0x53, 0x54, 0x5f, 0x52, 0x45, 0x53, 0x50, 0x5f, 0x4e,
	0x4f, 0x4e, 0x43, 0x45, 0x10, 0x02, 0x12, 0x1a, 0x0a, 0x16, 0x41, 0x54, 0x54, 0x45, 0x53, 0x54,
	0x5f, 0x52, 0x45, 0x53, 0x50, 0x5f, 0x51, 0x55, 0x4f, 0x54, 0x45, 0x5f, 0x52, 0x45, 0x53, 0x50,
	0x10, 0x03, 0x2a, 0x69, 0x0a, 0x13, 0x5a, 0x41, 0x74, 0x74, 0x65, 0x73, 0x74, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x18, 0x0a, 0x14, 0x41, 0x54, 0x54,
	0x45, 0x53, 0x54, 0x5f, 0x52, 0x45, 0x53, 0x50, 0x4f, 0x4e, 0x53, 0x45, 0x5f, 0x4e, 0x4f, 0x4e,
	0x45, 0x10, 0x00, 0x12, 0x1b, 0x0a, 0x17, 0x41, 0x54, 0x54, 0x45, 0x53, 0x54, 0x5f, 0x52, 0x45,
	0x53, 0x50, 0x4f, 0x4e, 0x53, 0x45, 0x5f, 0x53, 0x55, 0x43, 0x43, 0x45, 0x53, 0x53, 0x10, 0x01,
	0x12, 0x1b, 0x0a, 0x17, 0x41, 0x54, 0x54, 0x45, 0x53, 0x54, 0x5f, 0x52, 0x45, 0x53, 0x50, 0x4f,
	0x4e, 0x53, 0x45, 0x5f, 0x46, 0x41, 0x49, 0x4c, 0x55, 0x52, 0x45, 0x10, 0x02, 0x2a, 0xbd, 0x01,
	0x0a, 0x0c, 0x5a, 0x45, 0x76, 0x65, 0x43, 0x65, 0x72, 0x74, 0x54, 0x79, 0x70, 0x65, 0x12, 0x19,
	0x0a, 0x15, 0x43, 0x45, 0x52, 0x54, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x44, 0x45, 0x56, 0x49,
	0x43, 0x45, 0x5f, 0x4e, 0x4f, 0x4e, 0x45, 0x10, 0x00, 0x12, 0x1f, 0x0a, 0x1b, 0x43, 0x45, 0x52,
	0x54, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x44, 0x45, 0x56, 0x49, 0x43, 0x45, 0x5f, 0x4f, 0x4e,
	0x42, 0x4f, 0x41, 0x52, 0x44, 0x49, 0x4e, 0x47, 0x10, 0x01, 0x12, 0x27, 0x0a, 0x23, 0x43, 0x45,
	0x52, 0x54, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x44, 0x45, 0x56, 0x49, 0x43, 0x45, 0x5f, 0x52,
	0x45, 0x53, 0x54, 0x52, 0x49, 0x43, 0x54, 0x45, 0x44, 0x5f, 0x53, 0x49, 0x47, 0x4e, 0x49, 0x4e,
	0x47, 0x10, 0x02, 0x12, 0x24, 0x0a, 0x20, 0x43, 0x45, 0x52, 0x54, 0x5f, 0x54, 0x59, 0x50, 0x45,
	0x5f, 0x44, 0x45, 0x56, 0x49, 0x43, 0x45, 0x5f, 0x45, 0x4e, 0x44, 0x4f, 0x52, 0x53, 0x45, 0x4d,
	0x45, 0x4e, 0x54, 0x5f, 0x52, 0x53, 0x41, 0x10, 0x03, 0x12, 0x22, 0x0a, 0x1e, 0x43, 0x45, 0x52,
	0x54, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x44, 0x45, 0x56, 0x49, 0x43, 0x45, 0x5f, 0x45, 0x43,
	0x44, 0x48, 0x5f, 0x45, 0x58, 0x43, 0x48, 0x41, 0x4e, 0x47, 0x45, 0x10, 0x04, 0x42, 0x3d, 0x0a,
	0x15, 0x6f, 0x72, 0x67, 0x2e, 0x6c, 0x66, 0x65, 0x64, 0x67, 0x65, 0x2e, 0x65, 0x76, 0x65, 0x2e,
	0x61, 0x74, 0x74, 0x65, 0x73, 0x74, 0x5a, 0x24, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x6c, 0x66, 0x2d, 0x65, 0x64, 0x67, 0x65, 0x2f, 0x65, 0x76, 0x65, 0x2f, 0x61,
	0x70, 0x69, 0x2f, 0x67, 0x6f, 0x2f, 0x61, 0x74, 0x74, 0x65, 0x73, 0x74, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_attest_attest_proto_rawDescOnce sync.Once
	file_attest_attest_proto_rawDescData = file_attest_attest_proto_rawDesc
)

func file_attest_attest_proto_rawDescGZIP() []byte {
	file_attest_attest_proto_rawDescOnce.Do(func() {
		file_attest_attest_proto_rawDescData = protoimpl.X.CompressGZIP(file_attest_attest_proto_rawDescData)
	})
	return file_attest_attest_proto_rawDescData
}

var file_attest_attest_proto_enumTypes = make([]protoimpl.EnumInfo, 4)
var file_attest_attest_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_attest_attest_proto_goTypes = []interface{}{
	(ZAttestReqType)(0),          // 0: ZAttestReqType
	(ZAttestRespType)(0),         // 1: ZAttestRespType
	(ZAttestResponseCode)(0),     // 2: ZAttestResponseCode
	(ZEveCertType)(0),            // 3: ZEveCertType
	(*ZAttestReq)(nil),           // 4: ZAttestReq
	(*ZAttestResponse)(nil),      // 5: ZAttestResponse
	(*ZAttestNonceResp)(nil),     // 6: ZAttestNonceResp
	(*ZAttestQuote)(nil),         // 7: ZAttestQuote
	(*ZAttestQuoteResp)(nil),     // 8: ZAttestQuoteResp
	(*ZEveCert)(nil),             // 9: ZEveCert
	(*ZEveCertAttr)(nil),         // 10: ZEveCertAttr
	(evecommon.HashAlgorithm)(0), // 11: org.lfedge.eve.common.HashAlgorithm
}
var file_attest_attest_proto_depIdxs = []int32{
	0,  // 0: ZAttestReq.reqType:type_name -> ZAttestReqType
	7,  // 1: ZAttestReq.quote:type_name -> ZAttestQuote
	9,  // 2: ZAttestReq.certs:type_name -> ZEveCert
	1,  // 3: ZAttestResponse.respType:type_name -> ZAttestRespType
	6,  // 4: ZAttestResponse.nonce:type_name -> ZAttestNonceResp
	8,  // 5: ZAttestResponse.quoteResp:type_name -> ZAttestQuoteResp
	2,  // 6: ZAttestQuoteResp.response:type_name -> ZAttestResponseCode
	11, // 7: ZEveCert.hashAlgo:type_name -> org.lfedge.eve.common.HashAlgorithm
	3,  // 8: ZEveCert.type:type_name -> ZEveCertType
	10, // 9: ZEveCert.attributes:type_name -> ZEveCertAttr
	10, // [10:10] is the sub-list for method output_type
	10, // [10:10] is the sub-list for method input_type
	10, // [10:10] is the sub-list for extension type_name
	10, // [10:10] is the sub-list for extension extendee
	0,  // [0:10] is the sub-list for field type_name
}

func init() { file_attest_attest_proto_init() }
func file_attest_attest_proto_init() {
	if File_attest_attest_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_attest_attest_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ZAttestReq); i {
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
		file_attest_attest_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ZAttestResponse); i {
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
		file_attest_attest_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ZAttestNonceResp); i {
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
		file_attest_attest_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ZAttestQuote); i {
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
		file_attest_attest_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ZAttestQuoteResp); i {
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
		file_attest_attest_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ZEveCert); i {
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
		file_attest_attest_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ZEveCertAttr); i {
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
			RawDescriptor: file_attest_attest_proto_rawDesc,
			NumEnums:      4,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_attest_attest_proto_goTypes,
		DependencyIndexes: file_attest_attest_proto_depIdxs,
		EnumInfos:         file_attest_attest_proto_enumTypes,
		MessageInfos:      file_attest_attest_proto_msgTypes,
	}.Build()
	File_attest_attest_proto = out.File
	file_attest_attest_proto_rawDesc = nil
	file_attest_attest_proto_goTypes = nil
	file_attest_attest_proto_depIdxs = nil
}
