// Copyright(c) 2017-2018 Zededa, Inc.
// All rights reserved.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0-devel
// 	protoc        v3.6.1
// source: config/baseosconfig.proto

package config

import (
	proto "github.com/golang/protobuf/proto"
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

// OS version key and value pair
type OSKeyTags struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	OSVerKey   string `protobuf:"bytes,1,opt,name=OSVerKey,proto3" json:"OSVerKey,omitempty"`
	OSVerValue string `protobuf:"bytes,2,opt,name=OSVerValue,proto3" json:"OSVerValue,omitempty"`
}

func (x *OSKeyTags) Reset() {
	*x = OSKeyTags{}
	if protoimpl.UnsafeEnabled {
		mi := &file_config_baseosconfig_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OSKeyTags) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OSKeyTags) ProtoMessage() {}

func (x *OSKeyTags) ProtoReflect() protoreflect.Message {
	mi := &file_config_baseosconfig_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OSKeyTags.ProtoReflect.Descriptor instead.
func (*OSKeyTags) Descriptor() ([]byte, []int) {
	return file_config_baseosconfig_proto_rawDescGZIP(), []int{0}
}

func (x *OSKeyTags) GetOSVerKey() string {
	if x != nil {
		return x.OSVerKey
	}
	return ""
}

func (x *OSKeyTags) GetOSVerValue() string {
	if x != nil {
		return x.OSVerValue
	}
	return ""
}

// repeated key value tags compromising
type OSVerDetails struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BaseOSParams []*OSKeyTags `protobuf:"bytes,12,rep,name=baseOSParams,proto3" json:"baseOSParams,omitempty"`
}

func (x *OSVerDetails) Reset() {
	*x = OSVerDetails{}
	if protoimpl.UnsafeEnabled {
		mi := &file_config_baseosconfig_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OSVerDetails) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OSVerDetails) ProtoMessage() {}

func (x *OSVerDetails) ProtoReflect() protoreflect.Message {
	mi := &file_config_baseosconfig_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OSVerDetails.ProtoReflect.Descriptor instead.
func (*OSVerDetails) Descriptor() ([]byte, []int) {
	return file_config_baseosconfig_proto_rawDescGZIP(), []int{1}
}

func (x *OSVerDetails) GetBaseOSParams() []*OSKeyTags {
	if x != nil {
		return x.BaseOSParams
	}
	return nil
}

type BaseOSConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Uuidandversion *UUIDandVersion `protobuf:"bytes,1,opt,name=uuidandversion,proto3" json:"uuidandversion,omitempty"`
	Drives         []*Drive        `protobuf:"bytes,3,rep,name=drives,proto3" json:"drives,omitempty"`
	Activate       bool            `protobuf:"varint,4,opt,name=activate,proto3" json:"activate,omitempty"`
	BaseOSVersion  string          `protobuf:"bytes,10,opt,name=baseOSVersion,proto3" json:"baseOSVersion,omitempty"`
	BaseOSDetails  *OSVerDetails   `protobuf:"bytes,11,opt,name=baseOSDetails,proto3" json:"baseOSDetails,omitempty"`
}

func (x *BaseOSConfig) Reset() {
	*x = BaseOSConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_config_baseosconfig_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BaseOSConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BaseOSConfig) ProtoMessage() {}

func (x *BaseOSConfig) ProtoReflect() protoreflect.Message {
	mi := &file_config_baseosconfig_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BaseOSConfig.ProtoReflect.Descriptor instead.
func (*BaseOSConfig) Descriptor() ([]byte, []int) {
	return file_config_baseosconfig_proto_rawDescGZIP(), []int{2}
}

func (x *BaseOSConfig) GetUuidandversion() *UUIDandVersion {
	if x != nil {
		return x.Uuidandversion
	}
	return nil
}

func (x *BaseOSConfig) GetDrives() []*Drive {
	if x != nil {
		return x.Drives
	}
	return nil
}

func (x *BaseOSConfig) GetActivate() bool {
	if x != nil {
		return x.Activate
	}
	return false
}

func (x *BaseOSConfig) GetBaseOSVersion() string {
	if x != nil {
		return x.BaseOSVersion
	}
	return ""
}

func (x *BaseOSConfig) GetBaseOSDetails() *OSVerDetails {
	if x != nil {
		return x.BaseOSDetails
	}
	return nil
}

var File_config_baseosconfig_proto protoreflect.FileDescriptor

var file_config_baseosconfig_proto_rawDesc = []byte{
	0x0a, 0x19, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2f, 0x62, 0x61, 0x73, 0x65, 0x6f, 0x73, 0x63,
	0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x15, 0x6f, 0x72, 0x67,
	0x2e, 0x6c, 0x66, 0x65, 0x64, 0x67, 0x65, 0x2e, 0x65, 0x76, 0x65, 0x2e, 0x63, 0x6f, 0x6e, 0x66,
	0x69, 0x67, 0x1a, 0x16, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2f, 0x64, 0x65, 0x76, 0x63, 0x6f,
	0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x14, 0x63, 0x6f, 0x6e, 0x66,
	0x69, 0x67, 0x2f, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0x47, 0x0a, 0x09, 0x4f, 0x53, 0x4b, 0x65, 0x79, 0x54, 0x61, 0x67, 0x73, 0x12, 0x1a, 0x0a,
	0x08, 0x4f, 0x53, 0x56, 0x65, 0x72, 0x4b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x4f, 0x53, 0x56, 0x65, 0x72, 0x4b, 0x65, 0x79, 0x12, 0x1e, 0x0a, 0x0a, 0x4f, 0x53, 0x56,
	0x65, 0x72, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x4f,
	0x53, 0x56, 0x65, 0x72, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x22, 0x54, 0x0a, 0x0c, 0x4f, 0x53, 0x56,
	0x65, 0x72, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x12, 0x44, 0x0a, 0x0c, 0x62, 0x61, 0x73,
	0x65, 0x4f, 0x53, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x18, 0x0c, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x20, 0x2e, 0x6f, 0x72, 0x67, 0x2e, 0x6c, 0x66, 0x65, 0x64, 0x67, 0x65, 0x2e, 0x65, 0x76, 0x65,
	0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x4f, 0x53, 0x4b, 0x65, 0x79, 0x54, 0x61, 0x67,
	0x73, 0x52, 0x0c, 0x62, 0x61, 0x73, 0x65, 0x4f, 0x53, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x22,
	0xa0, 0x02, 0x0a, 0x0c, 0x42, 0x61, 0x73, 0x65, 0x4f, 0x53, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67,
	0x12, 0x4d, 0x0a, 0x0e, 0x75, 0x75, 0x69, 0x64, 0x61, 0x6e, 0x64, 0x76, 0x65, 0x72, 0x73, 0x69,
	0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x25, 0x2e, 0x6f, 0x72, 0x67, 0x2e, 0x6c,
	0x66, 0x65, 0x64, 0x67, 0x65, 0x2e, 0x65, 0x76, 0x65, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67,
	0x2e, 0x55, 0x55, 0x49, 0x44, 0x61, 0x6e, 0x64, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x52,
	0x0e, 0x75, 0x75, 0x69, 0x64, 0x61, 0x6e, 0x64, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12,
	0x34, 0x0a, 0x06, 0x64, 0x72, 0x69, 0x76, 0x65, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x1c, 0x2e, 0x6f, 0x72, 0x67, 0x2e, 0x6c, 0x66, 0x65, 0x64, 0x67, 0x65, 0x2e, 0x65, 0x76, 0x65,
	0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x44, 0x72, 0x69, 0x76, 0x65, 0x52, 0x06, 0x64,
	0x72, 0x69, 0x76, 0x65, 0x73, 0x12, 0x1a, 0x0a, 0x08, 0x61, 0x63, 0x74, 0x69, 0x76, 0x61, 0x74,
	0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x08, 0x52, 0x08, 0x61, 0x63, 0x74, 0x69, 0x76, 0x61, 0x74,
	0x65, 0x12, 0x24, 0x0a, 0x0d, 0x62, 0x61, 0x73, 0x65, 0x4f, 0x53, 0x56, 0x65, 0x72, 0x73, 0x69,
	0x6f, 0x6e, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x62, 0x61, 0x73, 0x65, 0x4f, 0x53,
	0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x49, 0x0a, 0x0d, 0x62, 0x61, 0x73, 0x65, 0x4f,
	0x53, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x23,
	0x2e, 0x6f, 0x72, 0x67, 0x2e, 0x6c, 0x66, 0x65, 0x64, 0x67, 0x65, 0x2e, 0x65, 0x76, 0x65, 0x2e,
	0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x4f, 0x53, 0x56, 0x65, 0x72, 0x44, 0x65, 0x74, 0x61,
	0x69, 0x6c, 0x73, 0x52, 0x0d, 0x62, 0x61, 0x73, 0x65, 0x4f, 0x53, 0x44, 0x65, 0x74, 0x61, 0x69,
	0x6c, 0x73, 0x42, 0x3d, 0x0a, 0x15, 0x6f, 0x72, 0x67, 0x2e, 0x6c, 0x66, 0x65, 0x64, 0x67, 0x65,
	0x2e, 0x65, 0x76, 0x65, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x5a, 0x24, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6c, 0x66, 0x2d, 0x65, 0x64, 0x67, 0x65, 0x2f,
	0x65, 0x76, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x67, 0x6f, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x69,
	0x67, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_config_baseosconfig_proto_rawDescOnce sync.Once
	file_config_baseosconfig_proto_rawDescData = file_config_baseosconfig_proto_rawDesc
)

func file_config_baseosconfig_proto_rawDescGZIP() []byte {
	file_config_baseosconfig_proto_rawDescOnce.Do(func() {
		file_config_baseosconfig_proto_rawDescData = protoimpl.X.CompressGZIP(file_config_baseosconfig_proto_rawDescData)
	})
	return file_config_baseosconfig_proto_rawDescData
}

var file_config_baseosconfig_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_config_baseosconfig_proto_goTypes = []interface{}{
	(*OSKeyTags)(nil),      // 0: org.lfedge.eve.config.OSKeyTags
	(*OSVerDetails)(nil),   // 1: org.lfedge.eve.config.OSVerDetails
	(*BaseOSConfig)(nil),   // 2: org.lfedge.eve.config.BaseOSConfig
	(*UUIDandVersion)(nil), // 3: org.lfedge.eve.config.UUIDandVersion
	(*Drive)(nil),          // 4: org.lfedge.eve.config.Drive
}
var file_config_baseosconfig_proto_depIdxs = []int32{
	0, // 0: org.lfedge.eve.config.OSVerDetails.baseOSParams:type_name -> org.lfedge.eve.config.OSKeyTags
	3, // 1: org.lfedge.eve.config.BaseOSConfig.uuidandversion:type_name -> org.lfedge.eve.config.UUIDandVersion
	4, // 2: org.lfedge.eve.config.BaseOSConfig.drives:type_name -> org.lfedge.eve.config.Drive
	1, // 3: org.lfedge.eve.config.BaseOSConfig.baseOSDetails:type_name -> org.lfedge.eve.config.OSVerDetails
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_config_baseosconfig_proto_init() }
func file_config_baseosconfig_proto_init() {
	if File_config_baseosconfig_proto != nil {
		return
	}
	file_config_devcommon_proto_init()
	file_config_storage_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_config_baseosconfig_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OSKeyTags); i {
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
		file_config_baseosconfig_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OSVerDetails); i {
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
		file_config_baseosconfig_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BaseOSConfig); i {
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
			RawDescriptor: file_config_baseosconfig_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_config_baseosconfig_proto_goTypes,
		DependencyIndexes: file_config_baseosconfig_proto_depIdxs,
		MessageInfos:      file_config_baseosconfig_proto_msgTypes,
	}.Build()
	File_config_baseosconfig_proto = out.File
	file_config_baseosconfig_proto_rawDesc = nil
	file_config_baseosconfig_proto_goTypes = nil
	file_config_baseosconfig_proto_depIdxs = nil
}
