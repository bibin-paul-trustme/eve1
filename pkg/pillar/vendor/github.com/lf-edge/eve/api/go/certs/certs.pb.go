// Copyright(c) 2020 Zededa, Inc.
// All rights reserved.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.21.0-devel
// 	protoc        v3.6.1
// source: certs/certs.proto

package certs

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

type ZCertType int32

const (
	ZCertType_CERT_TYPE_CONTROLLER_NONE          ZCertType = 0
	ZCertType_CERT_TYPE_CONTROLLER_SIGNING       ZCertType = 1 //set for the leaf certificate used by controller to sign payload envelopes
	ZCertType_CERT_TYPE_CONTROLLER_INTERMEDIATE  ZCertType = 2 //set for intermediate certs used to validate the certificates
	ZCertType_CERT_TYPE_CONTROLLER_ECDH_EXCHANGE ZCertType = 3 //set for certificate used by controller to share any symmetric key using ECDH
)

// Enum value maps for ZCertType.
var (
	ZCertType_name = map[int32]string{
		0: "CERT_TYPE_CONTROLLER_NONE",
		1: "CERT_TYPE_CONTROLLER_SIGNING",
		2: "CERT_TYPE_CONTROLLER_INTERMEDIATE",
		3: "CERT_TYPE_CONTROLLER_ECDH_EXCHANGE",
	}
	ZCertType_value = map[string]int32{
		"CERT_TYPE_CONTROLLER_NONE":          0,
		"CERT_TYPE_CONTROLLER_SIGNING":       1,
		"CERT_TYPE_CONTROLLER_INTERMEDIATE":  2,
		"CERT_TYPE_CONTROLLER_ECDH_EXCHANGE": 3,
	}
)

func (x ZCertType) Enum() *ZCertType {
	p := new(ZCertType)
	*p = x
	return p
}

func (x ZCertType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ZCertType) Descriptor() protoreflect.EnumDescriptor {
	return file_certs_certs_proto_enumTypes[0].Descriptor()
}

func (ZCertType) Type() protoreflect.EnumType {
	return &file_certs_certs_proto_enumTypes[0]
}

func (x ZCertType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ZCertType.Descriptor instead.
func (ZCertType) EnumDescriptor() ([]byte, []int) {
	return file_certs_certs_proto_rawDescGZIP(), []int{0}
}

// This is the response payload for GET /api/v1/edgeDevice/certs
// or /api/v2/edgeDevice/certs
// ZControllerCert carries a set of X.509 certificate and their properties
// from Controller to EVE.
type ZControllerCert struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Certs []*ZCert `protobuf:"bytes,1,rep,name=certs,proto3" json:"certs,omitempty"` //list of certificates sent by controller
}

func (x *ZControllerCert) Reset() {
	*x = ZControllerCert{}
	if protoimpl.UnsafeEnabled {
		mi := &file_certs_certs_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ZControllerCert) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ZControllerCert) ProtoMessage() {}

func (x *ZControllerCert) ProtoReflect() protoreflect.Message {
	mi := &file_certs_certs_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ZControllerCert.ProtoReflect.Descriptor instead.
func (*ZControllerCert) Descriptor() ([]byte, []int) {
	return file_certs_certs_proto_rawDescGZIP(), []int{0}
}

func (x *ZControllerCert) GetCerts() []*ZCert {
	if x != nil {
		return x.Certs
	}
	return nil
}

type ZCert struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	HashAlgo common.HashAlgorithm `protobuf:"varint,1,opt,name=hashAlgo,proto3,enum=common.HashAlgorithm" json:"hashAlgo,omitempty"` //hash method used to arrive at certHash
	CertHash []byte               `protobuf:"bytes,2,opt,name=certHash,proto3" json:"certHash,omitempty"`                            //truncated hash of the cert, according to hashing scheme in hashAlgo
	Type     ZCertType            `protobuf:"varint,3,opt,name=type,proto3,enum=ZCertType" json:"type,omitempty"`                    //what kind of certificate(to identify the target use case)
	Cert     []byte               `protobuf:"bytes,4,opt,name=cert,proto3" json:"cert,omitempty"`                                    //X509 cert in .PEM format
}

func (x *ZCert) Reset() {
	*x = ZCert{}
	if protoimpl.UnsafeEnabled {
		mi := &file_certs_certs_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ZCert) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ZCert) ProtoMessage() {}

func (x *ZCert) ProtoReflect() protoreflect.Message {
	mi := &file_certs_certs_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ZCert.ProtoReflect.Descriptor instead.
func (*ZCert) Descriptor() ([]byte, []int) {
	return file_certs_certs_proto_rawDescGZIP(), []int{1}
}

func (x *ZCert) GetHashAlgo() common.HashAlgorithm {
	if x != nil {
		return x.HashAlgo
	}
	return common.HashAlgorithm_HASH_ALGORITHM_INVALID
}

func (x *ZCert) GetCertHash() []byte {
	if x != nil {
		return x.CertHash
	}
	return nil
}

func (x *ZCert) GetType() ZCertType {
	if x != nil {
		return x.Type
	}
	return ZCertType_CERT_TYPE_CONTROLLER_NONE
}

func (x *ZCert) GetCert() []byte {
	if x != nil {
		return x.Cert
	}
	return nil
}

var File_certs_certs_proto protoreflect.FileDescriptor

var file_certs_certs_proto_rawDesc = []byte{
	0x0a, 0x11, 0x63, 0x65, 0x72, 0x74, 0x73, 0x2f, 0x63, 0x65, 0x72, 0x74, 0x73, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x13, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x63, 0x6f, 0x6d, 0x6d,
	0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x2f, 0x0a, 0x0f, 0x5a, 0x43, 0x6f, 0x6e,
	0x74, 0x72, 0x6f, 0x6c, 0x6c, 0x65, 0x72, 0x43, 0x65, 0x72, 0x74, 0x12, 0x1c, 0x0a, 0x05, 0x63,
	0x65, 0x72, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x06, 0x2e, 0x5a, 0x43, 0x65,
	0x72, 0x74, 0x52, 0x05, 0x63, 0x65, 0x72, 0x74, 0x73, 0x22, 0x8a, 0x01, 0x0a, 0x05, 0x5a, 0x43,
	0x65, 0x72, 0x74, 0x12, 0x31, 0x0a, 0x08, 0x68, 0x61, 0x73, 0x68, 0x41, 0x6c, 0x67, 0x6f, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x15, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x48,
	0x61, 0x73, 0x68, 0x41, 0x6c, 0x67, 0x6f, 0x72, 0x69, 0x74, 0x68, 0x6d, 0x52, 0x08, 0x68, 0x61,
	0x73, 0x68, 0x41, 0x6c, 0x67, 0x6f, 0x12, 0x1a, 0x0a, 0x08, 0x63, 0x65, 0x72, 0x74, 0x48, 0x61,
	0x73, 0x68, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x08, 0x63, 0x65, 0x72, 0x74, 0x48, 0x61,
	0x73, 0x68, 0x12, 0x1e, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0e,
	0x32, 0x0a, 0x2e, 0x5a, 0x43, 0x65, 0x72, 0x74, 0x54, 0x79, 0x70, 0x65, 0x52, 0x04, 0x74, 0x79,
	0x70, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x65, 0x72, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0c,
	0x52, 0x04, 0x63, 0x65, 0x72, 0x74, 0x2a, 0x9b, 0x01, 0x0a, 0x09, 0x5a, 0x43, 0x65, 0x72, 0x74,
	0x54, 0x79, 0x70, 0x65, 0x12, 0x1d, 0x0a, 0x19, 0x43, 0x45, 0x52, 0x54, 0x5f, 0x54, 0x59, 0x50,
	0x45, 0x5f, 0x43, 0x4f, 0x4e, 0x54, 0x52, 0x4f, 0x4c, 0x4c, 0x45, 0x52, 0x5f, 0x4e, 0x4f, 0x4e,
	0x45, 0x10, 0x00, 0x12, 0x20, 0x0a, 0x1c, 0x43, 0x45, 0x52, 0x54, 0x5f, 0x54, 0x59, 0x50, 0x45,
	0x5f, 0x43, 0x4f, 0x4e, 0x54, 0x52, 0x4f, 0x4c, 0x4c, 0x45, 0x52, 0x5f, 0x53, 0x49, 0x47, 0x4e,
	0x49, 0x4e, 0x47, 0x10, 0x01, 0x12, 0x25, 0x0a, 0x21, 0x43, 0x45, 0x52, 0x54, 0x5f, 0x54, 0x59,
	0x50, 0x45, 0x5f, 0x43, 0x4f, 0x4e, 0x54, 0x52, 0x4f, 0x4c, 0x4c, 0x45, 0x52, 0x5f, 0x49, 0x4e,
	0x54, 0x45, 0x52, 0x4d, 0x45, 0x44, 0x49, 0x41, 0x54, 0x45, 0x10, 0x02, 0x12, 0x26, 0x0a, 0x22,
	0x43, 0x45, 0x52, 0x54, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x43, 0x4f, 0x4e, 0x54, 0x52, 0x4f,
	0x4c, 0x4c, 0x45, 0x52, 0x5f, 0x45, 0x43, 0x44, 0x48, 0x5f, 0x45, 0x58, 0x43, 0x48, 0x41, 0x4e,
	0x47, 0x45, 0x10, 0x03, 0x42, 0x46, 0x0a, 0x1f, 0x63, 0x6f, 0x6d, 0x2e, 0x7a, 0x65, 0x64, 0x65,
	0x64, 0x61, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x5a, 0x23, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x6c, 0x66, 0x2d, 0x65, 0x64, 0x67, 0x65, 0x2f, 0x65, 0x76, 0x65, 0x2f,
	0x61, 0x70, 0x69, 0x2f, 0x67, 0x6f, 0x2f, 0x63, 0x65, 0x72, 0x74, 0x73, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_certs_certs_proto_rawDescOnce sync.Once
	file_certs_certs_proto_rawDescData = file_certs_certs_proto_rawDesc
)

func file_certs_certs_proto_rawDescGZIP() []byte {
	file_certs_certs_proto_rawDescOnce.Do(func() {
		file_certs_certs_proto_rawDescData = protoimpl.X.CompressGZIP(file_certs_certs_proto_rawDescData)
	})
	return file_certs_certs_proto_rawDescData
}

var file_certs_certs_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_certs_certs_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_certs_certs_proto_goTypes = []interface{}{
	(ZCertType)(0),            // 0: ZCertType
	(*ZControllerCert)(nil),   // 1: ZControllerCert
	(*ZCert)(nil),             // 2: ZCert
	(common.HashAlgorithm)(0), // 3: common.HashAlgorithm
}
var file_certs_certs_proto_depIdxs = []int32{
	2, // 0: ZControllerCert.certs:type_name -> ZCert
	3, // 1: ZCert.hashAlgo:type_name -> common.HashAlgorithm
	0, // 2: ZCert.type:type_name -> ZCertType
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_certs_certs_proto_init() }
func file_certs_certs_proto_init() {
	if File_certs_certs_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_certs_certs_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ZControllerCert); i {
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
		file_certs_certs_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ZCert); i {
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
			RawDescriptor: file_certs_certs_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_certs_certs_proto_goTypes,
		DependencyIndexes: file_certs_certs_proto_depIdxs,
		EnumInfos:         file_certs_certs_proto_enumTypes,
		MessageInfos:      file_certs_certs_proto_msgTypes,
	}.Build()
	File_certs_certs_proto = out.File
	file_certs_certs_proto_rawDesc = nil
	file_certs_certs_proto_goTypes = nil
	file_certs_certs_proto_depIdxs = nil
}
