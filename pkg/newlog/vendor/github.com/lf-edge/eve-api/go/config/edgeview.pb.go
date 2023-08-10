// Copyright(c) 2021-2022 Zededa, Inc.
// All rights reserved.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v4.23.4
// source: config/edgeview.proto

package config

import (
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

type EdgeViewConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// JWT token for signed info, it contains the dispatcher
	// endpoint IP:Port, device UUID, nonce and expiration time
	Token string `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	// dispatcher certificate(s) if it's not well-known CA signed
	DispCertPem [][]byte `protobuf:"bytes,2,rep,name=disp_cert_pem,json=dispCertPem,proto3" json:"disp_cert_pem,omitempty"`
	// policy for device access through edge-view
	DevPolicy *DevDebugAccessPolicy `protobuf:"bytes,3,opt,name=dev_policy,json=devPolicy,proto3" json:"dev_policy,omitempty"`
	// policy access for apps through edge-view
	AppPolicy *AppDebugAccessPolicy `protobuf:"bytes,4,opt,name=app_policy,json=appPolicy,proto3" json:"app_policy,omitempty"`
	// policy access for external endpoint through edge-view
	ExtPolicy *ExternalEndPointPolicy `protobuf:"bytes,5,opt,name=ext_policy,json=extPolicy,proto3" json:"ext_policy,omitempty"`
	// Generation ID for re-start edgeview without parameter changes
	GenerationId uint32 `protobuf:"varint,6,opt,name=generation_id,json=generationId,proto3" json:"generation_id,omitempty"`
}

func (x *EdgeViewConfig) Reset() {
	*x = EdgeViewConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_config_edgeview_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EdgeViewConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EdgeViewConfig) ProtoMessage() {}

func (x *EdgeViewConfig) ProtoReflect() protoreflect.Message {
	mi := &file_config_edgeview_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EdgeViewConfig.ProtoReflect.Descriptor instead.
func (*EdgeViewConfig) Descriptor() ([]byte, []int) {
	return file_config_edgeview_proto_rawDescGZIP(), []int{0}
}

func (x *EdgeViewConfig) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

func (x *EdgeViewConfig) GetDispCertPem() [][]byte {
	if x != nil {
		return x.DispCertPem
	}
	return nil
}

func (x *EdgeViewConfig) GetDevPolicy() *DevDebugAccessPolicy {
	if x != nil {
		return x.DevPolicy
	}
	return nil
}

func (x *EdgeViewConfig) GetAppPolicy() *AppDebugAccessPolicy {
	if x != nil {
		return x.AppPolicy
	}
	return nil
}

func (x *EdgeViewConfig) GetExtPolicy() *ExternalEndPointPolicy {
	if x != nil {
		return x.ExtPolicy
	}
	return nil
}

func (x *EdgeViewConfig) GetGenerationId() uint32 {
	if x != nil {
		return x.GenerationId
	}
	return 0
}

// Dev debug policy applicable to edge-view
type DevDebugAccessPolicy struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// device side of edge-view access is allowed or not
	AllowDev bool `protobuf:"varint,1,opt,name=allow_dev,json=allowDev,proto3" json:"allow_dev,omitempty"`
}

func (x *DevDebugAccessPolicy) Reset() {
	*x = DevDebugAccessPolicy{}
	if protoimpl.UnsafeEnabled {
		mi := &file_config_edgeview_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DevDebugAccessPolicy) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DevDebugAccessPolicy) ProtoMessage() {}

func (x *DevDebugAccessPolicy) ProtoReflect() protoreflect.Message {
	mi := &file_config_edgeview_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DevDebugAccessPolicy.ProtoReflect.Descriptor instead.
func (*DevDebugAccessPolicy) Descriptor() ([]byte, []int) {
	return file_config_edgeview_proto_rawDescGZIP(), []int{1}
}

func (x *DevDebugAccessPolicy) GetAllowDev() bool {
	if x != nil {
		return x.AllowDev
	}
	return false
}

// App debug policy applicable to edge-view
type AppDebugAccessPolicy struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// app side of edge-view access is allowed or not
	AllowApp bool `protobuf:"varint,1,opt,name=allow_app,json=allowApp,proto3" json:"allow_app,omitempty"`
}

func (x *AppDebugAccessPolicy) Reset() {
	*x = AppDebugAccessPolicy{}
	if protoimpl.UnsafeEnabled {
		mi := &file_config_edgeview_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AppDebugAccessPolicy) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AppDebugAccessPolicy) ProtoMessage() {}

func (x *AppDebugAccessPolicy) ProtoReflect() protoreflect.Message {
	mi := &file_config_edgeview_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AppDebugAccessPolicy.ProtoReflect.Descriptor instead.
func (*AppDebugAccessPolicy) Descriptor() ([]byte, []int) {
	return file_config_edgeview_proto_rawDescGZIP(), []int{2}
}

func (x *AppDebugAccessPolicy) GetAllowApp() bool {
	if x != nil {
		return x.AllowApp
	}
	return false
}

// External Endpoint applicable to edge-view
// To mean the entity external to the device, e.g. a local-profile server on the LAN outside of mgmt
// or app-shared ports. since it's not part of EVE, and not part of EVE applications. In the EdgeView code,
// if tcp session setup is to an address we don't have, it identifies the request as 'external'
type ExternalEndPointPolicy struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// external of device side of edge-view access is allowed or not
	AllowExt bool `protobuf:"varint,1,opt,name=allow_ext,json=allowExt,proto3" json:"allow_ext,omitempty"`
}

func (x *ExternalEndPointPolicy) Reset() {
	*x = ExternalEndPointPolicy{}
	if protoimpl.UnsafeEnabled {
		mi := &file_config_edgeview_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ExternalEndPointPolicy) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ExternalEndPointPolicy) ProtoMessage() {}

func (x *ExternalEndPointPolicy) ProtoReflect() protoreflect.Message {
	mi := &file_config_edgeview_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ExternalEndPointPolicy.ProtoReflect.Descriptor instead.
func (*ExternalEndPointPolicy) Descriptor() ([]byte, []int) {
	return file_config_edgeview_proto_rawDescGZIP(), []int{3}
}

func (x *ExternalEndPointPolicy) GetAllowExt() bool {
	if x != nil {
		return x.AllowExt
	}
	return false
}

var File_config_edgeview_proto protoreflect.FileDescriptor

var file_config_edgeview_proto_rawDesc = []byte{
	0x0a, 0x15, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2f, 0x65, 0x64, 0x67, 0x65, 0x76, 0x69, 0x65,
	0x77, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x15, 0x6f, 0x72, 0x67, 0x2e, 0x6c, 0x66, 0x65,
	0x64, 0x67, 0x65, 0x2e, 0x65, 0x76, 0x65, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x22, 0xd5,
	0x02, 0x0a, 0x0e, 0x45, 0x64, 0x67, 0x65, 0x56, 0x69, 0x65, 0x77, 0x43, 0x6f, 0x6e, 0x66, 0x69,
	0x67, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x22, 0x0a, 0x0d, 0x64, 0x69, 0x73, 0x70, 0x5f,
	0x63, 0x65, 0x72, 0x74, 0x5f, 0x70, 0x65, 0x6d, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0c, 0x52, 0x0b,
	0x64, 0x69, 0x73, 0x70, 0x43, 0x65, 0x72, 0x74, 0x50, 0x65, 0x6d, 0x12, 0x4a, 0x0a, 0x0a, 0x64,
	0x65, 0x76, 0x5f, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x2b, 0x2e, 0x6f, 0x72, 0x67, 0x2e, 0x6c, 0x66, 0x65, 0x64, 0x67, 0x65, 0x2e, 0x65, 0x76, 0x65,
	0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x44, 0x65, 0x76, 0x44, 0x65, 0x62, 0x75, 0x67,
	0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x52, 0x09, 0x64, 0x65,
	0x76, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x12, 0x4a, 0x0a, 0x0a, 0x61, 0x70, 0x70, 0x5f, 0x70,
	0x6f, 0x6c, 0x69, 0x63, 0x79, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2b, 0x2e, 0x6f, 0x72,
	0x67, 0x2e, 0x6c, 0x66, 0x65, 0x64, 0x67, 0x65, 0x2e, 0x65, 0x76, 0x65, 0x2e, 0x63, 0x6f, 0x6e,
	0x66, 0x69, 0x67, 0x2e, 0x41, 0x70, 0x70, 0x44, 0x65, 0x62, 0x75, 0x67, 0x41, 0x63, 0x63, 0x65,
	0x73, 0x73, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x52, 0x09, 0x61, 0x70, 0x70, 0x50, 0x6f, 0x6c,
	0x69, 0x63, 0x79, 0x12, 0x4c, 0x0a, 0x0a, 0x65, 0x78, 0x74, 0x5f, 0x70, 0x6f, 0x6c, 0x69, 0x63,
	0x79, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2d, 0x2e, 0x6f, 0x72, 0x67, 0x2e, 0x6c, 0x66,
	0x65, 0x64, 0x67, 0x65, 0x2e, 0x65, 0x76, 0x65, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e,
	0x45, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x45, 0x6e, 0x64, 0x50, 0x6f, 0x69, 0x6e, 0x74,
	0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x52, 0x09, 0x65, 0x78, 0x74, 0x50, 0x6f, 0x6c, 0x69, 0x63,
	0x79, 0x12, 0x23, 0x0a, 0x0d, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f,
	0x69, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0c, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x22, 0x33, 0x0a, 0x14, 0x44, 0x65, 0x76, 0x44, 0x65, 0x62,
	0x75, 0x67, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x12, 0x1b,
	0x0a, 0x09, 0x61, 0x6c, 0x6c, 0x6f, 0x77, 0x5f, 0x64, 0x65, 0x76, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x08, 0x52, 0x08, 0x61, 0x6c, 0x6c, 0x6f, 0x77, 0x44, 0x65, 0x76, 0x22, 0x33, 0x0a, 0x14, 0x41,
	0x70, 0x70, 0x44, 0x65, 0x62, 0x75, 0x67, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x50, 0x6f, 0x6c,
	0x69, 0x63, 0x79, 0x12, 0x1b, 0x0a, 0x09, 0x61, 0x6c, 0x6c, 0x6f, 0x77, 0x5f, 0x61, 0x70, 0x70,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x08, 0x61, 0x6c, 0x6c, 0x6f, 0x77, 0x41, 0x70, 0x70,
	0x22, 0x35, 0x0a, 0x16, 0x45, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x45, 0x6e, 0x64, 0x50,
	0x6f, 0x69, 0x6e, 0x74, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x12, 0x1b, 0x0a, 0x09, 0x61, 0x6c,
	0x6c, 0x6f, 0x77, 0x5f, 0x65, 0x78, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x08, 0x61,
	0x6c, 0x6c, 0x6f, 0x77, 0x45, 0x78, 0x74, 0x42, 0x3d, 0x0a, 0x15, 0x6f, 0x72, 0x67, 0x2e, 0x6c,
	0x66, 0x65, 0x64, 0x67, 0x65, 0x2e, 0x65, 0x76, 0x65, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67,
	0x5a, 0x24, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6c, 0x66, 0x2d,
	0x65, 0x64, 0x67, 0x65, 0x2f, 0x65, 0x76, 0x65, 0x2d, 0x61, 0x70, 0x69, 0x2f, 0x67, 0x6f, 0x2f,
	0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_config_edgeview_proto_rawDescOnce sync.Once
	file_config_edgeview_proto_rawDescData = file_config_edgeview_proto_rawDesc
)

func file_config_edgeview_proto_rawDescGZIP() []byte {
	file_config_edgeview_proto_rawDescOnce.Do(func() {
		file_config_edgeview_proto_rawDescData = protoimpl.X.CompressGZIP(file_config_edgeview_proto_rawDescData)
	})
	return file_config_edgeview_proto_rawDescData
}

var file_config_edgeview_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_config_edgeview_proto_goTypes = []interface{}{
	(*EdgeViewConfig)(nil),         // 0: org.lfedge.eve.config.EdgeViewConfig
	(*DevDebugAccessPolicy)(nil),   // 1: org.lfedge.eve.config.DevDebugAccessPolicy
	(*AppDebugAccessPolicy)(nil),   // 2: org.lfedge.eve.config.AppDebugAccessPolicy
	(*ExternalEndPointPolicy)(nil), // 3: org.lfedge.eve.config.ExternalEndPointPolicy
}
var file_config_edgeview_proto_depIdxs = []int32{
	1, // 0: org.lfedge.eve.config.EdgeViewConfig.dev_policy:type_name -> org.lfedge.eve.config.DevDebugAccessPolicy
	2, // 1: org.lfedge.eve.config.EdgeViewConfig.app_policy:type_name -> org.lfedge.eve.config.AppDebugAccessPolicy
	3, // 2: org.lfedge.eve.config.EdgeViewConfig.ext_policy:type_name -> org.lfedge.eve.config.ExternalEndPointPolicy
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_config_edgeview_proto_init() }
func file_config_edgeview_proto_init() {
	if File_config_edgeview_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_config_edgeview_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EdgeViewConfig); i {
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
		file_config_edgeview_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DevDebugAccessPolicy); i {
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
		file_config_edgeview_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AppDebugAccessPolicy); i {
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
		file_config_edgeview_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ExternalEndPointPolicy); i {
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
			RawDescriptor: file_config_edgeview_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_config_edgeview_proto_goTypes,
		DependencyIndexes: file_config_edgeview_proto_depIdxs,
		MessageInfos:      file_config_edgeview_proto_msgTypes,
	}.Build()
	File_config_edgeview_proto = out.File
	file_config_edgeview_proto_rawDesc = nil
	file_config_edgeview_proto_goTypes = nil
	file_config_edgeview_proto_depIdxs = nil
}
