// Copyright(c) 2017-2018 Zededa, Inc.
// All rights reserved.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0-devel
// 	protoc        v3.6.1
// source: config/vm.proto

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

// For now we need to tell the device which virtualization mode
// to use. Later we might use a single one for all VMs (on any particular
// ISA). If we end up keeping this we should make the names be less
// tied to a particular hypervisor.
type VmMode int32

const (
	VmMode_PV      VmMode = 0
	VmMode_HVM     VmMode = 1
	VmMode_Filler  VmMode = 2 // PVH = 2;
	VmMode_FML     VmMode = 3 // Experimental machine learning mode
	VmMode_NOHYPER VmMode = 4 // Do not use a hypervisor
)

// Enum value maps for VmMode.
var (
	VmMode_name = map[int32]string{
		0: "PV",
		1: "HVM",
		2: "Filler",
		3: "FML",
		4: "NOHYPER",
	}
	VmMode_value = map[string]int32{
		"PV":      0,
		"HVM":     1,
		"Filler":  2,
		"FML":     3,
		"NOHYPER": 4,
	}
)

func (x VmMode) Enum() *VmMode {
	p := new(VmMode)
	*p = x
	return p
}

func (x VmMode) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (VmMode) Descriptor() protoreflect.EnumDescriptor {
	return file_config_vm_proto_enumTypes[0].Descriptor()
}

func (VmMode) Type() protoreflect.EnumType {
	return &file_config_vm_proto_enumTypes[0]
}

func (x VmMode) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use VmMode.Descriptor instead.
func (VmMode) EnumDescriptor() ([]byte, []int) {
	return file_config_vm_proto_rawDescGZIP(), []int{0}
}

type VmConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Kernel             string   `protobuf:"bytes,1,opt,name=kernel,proto3" json:"kernel,omitempty"`
	Ramdisk            string   `protobuf:"bytes,2,opt,name=ramdisk,proto3" json:"ramdisk,omitempty"`
	Memory             uint32   `protobuf:"varint,3,opt,name=memory,proto3" json:"memory,omitempty"`
	Maxmem             uint32   `protobuf:"varint,4,opt,name=maxmem,proto3" json:"maxmem,omitempty"`
	Vcpus              uint32   `protobuf:"varint,5,opt,name=vcpus,proto3" json:"vcpus,omitempty"`
	Maxcpus            uint32   `protobuf:"varint,6,opt,name=maxcpus,proto3" json:"maxcpus,omitempty"`
	Rootdev            string   `protobuf:"bytes,7,opt,name=rootdev,proto3" json:"rootdev,omitempty"`
	Extraargs          string   `protobuf:"bytes,8,opt,name=extraargs,proto3" json:"extraargs,omitempty"`
	Bootloader         string   `protobuf:"bytes,9,opt,name=bootloader,proto3" json:"bootloader,omitempty"`
	Cpus               string   `protobuf:"bytes,10,opt,name=cpus,proto3" json:"cpus,omitempty"`
	Devicetree         string   `protobuf:"bytes,11,opt,name=devicetree,proto3" json:"devicetree,omitempty"`
	Dtdev              []string `protobuf:"bytes,12,rep,name=dtdev,proto3" json:"dtdev,omitempty"`
	Irqs               []uint32 `protobuf:"varint,13,rep,packed,name=irqs,proto3" json:"irqs,omitempty"`
	Iomem              []string `protobuf:"bytes,14,rep,name=iomem,proto3" json:"iomem,omitempty"`
	VirtualizationMode VmMode   `protobuf:"varint,15,opt,name=virtualizationMode,proto3,enum=org.lfedge.eve.config.VmMode" json:"virtualizationMode,omitempty"`
	EnableVnc          bool     `protobuf:"varint,16,opt,name=enableVnc,proto3" json:"enableVnc,omitempty"`
	VncDisplay         uint32   `protobuf:"varint,17,opt,name=vncDisplay,proto3" json:"vncDisplay,omitempty"`
	VncPasswd          string   `protobuf:"bytes,18,opt,name=vncPasswd,proto3" json:"vncPasswd,omitempty"`
	OomScore           int32    `protobuf:"varint,19,opt,name=oomScore,proto3" json:"oomScore,omitempty"`
}

func (x *VmConfig) Reset() {
	*x = VmConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_config_vm_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *VmConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VmConfig) ProtoMessage() {}

func (x *VmConfig) ProtoReflect() protoreflect.Message {
	mi := &file_config_vm_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VmConfig.ProtoReflect.Descriptor instead.
func (*VmConfig) Descriptor() ([]byte, []int) {
	return file_config_vm_proto_rawDescGZIP(), []int{0}
}

func (x *VmConfig) GetKernel() string {
	if x != nil {
		return x.Kernel
	}
	return ""
}

func (x *VmConfig) GetRamdisk() string {
	if x != nil {
		return x.Ramdisk
	}
	return ""
}

func (x *VmConfig) GetMemory() uint32 {
	if x != nil {
		return x.Memory
	}
	return 0
}

func (x *VmConfig) GetMaxmem() uint32 {
	if x != nil {
		return x.Maxmem
	}
	return 0
}

func (x *VmConfig) GetVcpus() uint32 {
	if x != nil {
		return x.Vcpus
	}
	return 0
}

func (x *VmConfig) GetMaxcpus() uint32 {
	if x != nil {
		return x.Maxcpus
	}
	return 0
}

func (x *VmConfig) GetRootdev() string {
	if x != nil {
		return x.Rootdev
	}
	return ""
}

func (x *VmConfig) GetExtraargs() string {
	if x != nil {
		return x.Extraargs
	}
	return ""
}

func (x *VmConfig) GetBootloader() string {
	if x != nil {
		return x.Bootloader
	}
	return ""
}

func (x *VmConfig) GetCpus() string {
	if x != nil {
		return x.Cpus
	}
	return ""
}

func (x *VmConfig) GetDevicetree() string {
	if x != nil {
		return x.Devicetree
	}
	return ""
}

func (x *VmConfig) GetDtdev() []string {
	if x != nil {
		return x.Dtdev
	}
	return nil
}

func (x *VmConfig) GetIrqs() []uint32 {
	if x != nil {
		return x.Irqs
	}
	return nil
}

func (x *VmConfig) GetIomem() []string {
	if x != nil {
		return x.Iomem
	}
	return nil
}

func (x *VmConfig) GetVirtualizationMode() VmMode {
	if x != nil {
		return x.VirtualizationMode
	}
	return VmMode_PV
}

func (x *VmConfig) GetEnableVnc() bool {
	if x != nil {
		return x.EnableVnc
	}
	return false
}

func (x *VmConfig) GetVncDisplay() uint32 {
	if x != nil {
		return x.VncDisplay
	}
	return 0
}

func (x *VmConfig) GetVncPasswd() string {
	if x != nil {
		return x.VncPasswd
	}
	return ""
}

func (x *VmConfig) GetOomScore() int32 {
	if x != nil {
		return x.OomScore
	}
	return 0
}

var File_config_vm_proto protoreflect.FileDescriptor

var file_config_vm_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2f, 0x76, 0x6d, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x15, 0x6f, 0x72, 0x67, 0x2e, 0x6c, 0x66, 0x65, 0x64, 0x67, 0x65, 0x2e, 0x65, 0x76,
	0x65, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x22, 0xaf, 0x04, 0x0a, 0x08, 0x56, 0x6d, 0x43,
	0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x16, 0x0a, 0x06, 0x6b, 0x65, 0x72, 0x6e, 0x65, 0x6c, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6b, 0x65, 0x72, 0x6e, 0x65, 0x6c, 0x12, 0x18, 0x0a,
	0x07, 0x72, 0x61, 0x6d, 0x64, 0x69, 0x73, 0x6b, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07,
	0x72, 0x61, 0x6d, 0x64, 0x69, 0x73, 0x6b, 0x12, 0x16, 0x0a, 0x06, 0x6d, 0x65, 0x6d, 0x6f, 0x72,
	0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x06, 0x6d, 0x65, 0x6d, 0x6f, 0x72, 0x79, 0x12,
	0x16, 0x0a, 0x06, 0x6d, 0x61, 0x78, 0x6d, 0x65, 0x6d, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0d, 0x52,
	0x06, 0x6d, 0x61, 0x78, 0x6d, 0x65, 0x6d, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x63, 0x70, 0x75, 0x73,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x05, 0x76, 0x63, 0x70, 0x75, 0x73, 0x12, 0x18, 0x0a,
	0x07, 0x6d, 0x61, 0x78, 0x63, 0x70, 0x75, 0x73, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x07,
	0x6d, 0x61, 0x78, 0x63, 0x70, 0x75, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x72, 0x6f, 0x6f, 0x74, 0x64,
	0x65, 0x76, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x72, 0x6f, 0x6f, 0x74, 0x64, 0x65,
	0x76, 0x12, 0x1c, 0x0a, 0x09, 0x65, 0x78, 0x74, 0x72, 0x61, 0x61, 0x72, 0x67, 0x73, 0x18, 0x08,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x65, 0x78, 0x74, 0x72, 0x61, 0x61, 0x72, 0x67, 0x73, 0x12,
	0x1e, 0x0a, 0x0a, 0x62, 0x6f, 0x6f, 0x74, 0x6c, 0x6f, 0x61, 0x64, 0x65, 0x72, 0x18, 0x09, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0a, 0x62, 0x6f, 0x6f, 0x74, 0x6c, 0x6f, 0x61, 0x64, 0x65, 0x72, 0x12,
	0x12, 0x0a, 0x04, 0x63, 0x70, 0x75, 0x73, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x63,
	0x70, 0x75, 0x73, 0x12, 0x1e, 0x0a, 0x0a, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x74, 0x72, 0x65,
	0x65, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x74,
	0x72, 0x65, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x64, 0x74, 0x64, 0x65, 0x76, 0x18, 0x0c, 0x20, 0x03,
	0x28, 0x09, 0x52, 0x05, 0x64, 0x74, 0x64, 0x65, 0x76, 0x12, 0x12, 0x0a, 0x04, 0x69, 0x72, 0x71,
	0x73, 0x18, 0x0d, 0x20, 0x03, 0x28, 0x0d, 0x52, 0x04, 0x69, 0x72, 0x71, 0x73, 0x12, 0x14, 0x0a,
	0x05, 0x69, 0x6f, 0x6d, 0x65, 0x6d, 0x18, 0x0e, 0x20, 0x03, 0x28, 0x09, 0x52, 0x05, 0x69, 0x6f,
	0x6d, 0x65, 0x6d, 0x12, 0x4d, 0x0a, 0x12, 0x76, 0x69, 0x72, 0x74, 0x75, 0x61, 0x6c, 0x69, 0x7a,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4d, 0x6f, 0x64, 0x65, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x0e, 0x32,
	0x1d, 0x2e, 0x6f, 0x72, 0x67, 0x2e, 0x6c, 0x66, 0x65, 0x64, 0x67, 0x65, 0x2e, 0x65, 0x76, 0x65,
	0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x56, 0x6d, 0x4d, 0x6f, 0x64, 0x65, 0x52, 0x12,
	0x76, 0x69, 0x72, 0x74, 0x75, 0x61, 0x6c, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4d, 0x6f,
	0x64, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x56, 0x6e, 0x63, 0x18,
	0x10, 0x20, 0x01, 0x28, 0x08, 0x52, 0x09, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x56, 0x6e, 0x63,
	0x12, 0x1e, 0x0a, 0x0a, 0x76, 0x6e, 0x63, 0x44, 0x69, 0x73, 0x70, 0x6c, 0x61, 0x79, 0x18, 0x11,
	0x20, 0x01, 0x28, 0x0d, 0x52, 0x0a, 0x76, 0x6e, 0x63, 0x44, 0x69, 0x73, 0x70, 0x6c, 0x61, 0x79,
	0x12, 0x1c, 0x0a, 0x09, 0x76, 0x6e, 0x63, 0x50, 0x61, 0x73, 0x73, 0x77, 0x64, 0x18, 0x12, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x09, 0x76, 0x6e, 0x63, 0x50, 0x61, 0x73, 0x73, 0x77, 0x64, 0x12, 0x1a,
	0x0a, 0x08, 0x6f, 0x6f, 0x6d, 0x53, 0x63, 0x6f, 0x72, 0x65, 0x18, 0x13, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x08, 0x6f, 0x6f, 0x6d, 0x53, 0x63, 0x6f, 0x72, 0x65, 0x2a, 0x3b, 0x0a, 0x06, 0x56, 0x6d,
	0x4d, 0x6f, 0x64, 0x65, 0x12, 0x06, 0x0a, 0x02, 0x50, 0x56, 0x10, 0x00, 0x12, 0x07, 0x0a, 0x03,
	0x48, 0x56, 0x4d, 0x10, 0x01, 0x12, 0x0a, 0x0a, 0x06, 0x46, 0x69, 0x6c, 0x6c, 0x65, 0x72, 0x10,
	0x02, 0x12, 0x07, 0x0a, 0x03, 0x46, 0x4d, 0x4c, 0x10, 0x03, 0x12, 0x0b, 0x0a, 0x07, 0x4e, 0x4f,
	0x48, 0x59, 0x50, 0x45, 0x52, 0x10, 0x04, 0x42, 0x3d, 0x0a, 0x15, 0x6f, 0x72, 0x67, 0x2e, 0x6c,
	0x66, 0x65, 0x64, 0x67, 0x65, 0x2e, 0x65, 0x76, 0x65, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67,
	0x5a, 0x24, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6c, 0x66, 0x2d,
	0x65, 0x64, 0x67, 0x65, 0x2f, 0x65, 0x76, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x67, 0x6f, 0x2f,
	0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_config_vm_proto_rawDescOnce sync.Once
	file_config_vm_proto_rawDescData = file_config_vm_proto_rawDesc
)

func file_config_vm_proto_rawDescGZIP() []byte {
	file_config_vm_proto_rawDescOnce.Do(func() {
		file_config_vm_proto_rawDescData = protoimpl.X.CompressGZIP(file_config_vm_proto_rawDescData)
	})
	return file_config_vm_proto_rawDescData
}

var file_config_vm_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_config_vm_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_config_vm_proto_goTypes = []interface{}{
	(VmMode)(0),      // 0: org.lfedge.eve.config.VmMode
	(*VmConfig)(nil), // 1: org.lfedge.eve.config.VmConfig
}
var file_config_vm_proto_depIdxs = []int32{
	0, // 0: org.lfedge.eve.config.VmConfig.virtualizationMode:type_name -> org.lfedge.eve.config.VmMode
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_config_vm_proto_init() }
func file_config_vm_proto_init() {
	if File_config_vm_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_config_vm_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*VmConfig); i {
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
			RawDescriptor: file_config_vm_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_config_vm_proto_goTypes,
		DependencyIndexes: file_config_vm_proto_depIdxs,
		EnumInfos:         file_config_vm_proto_enumTypes,
		MessageInfos:      file_config_vm_proto_msgTypes,
	}.Build()
	File_config_vm_proto = out.File
	file_config_vm_proto_rawDesc = nil
	file_config_vm_proto_goTypes = nil
	file_config_vm_proto_depIdxs = nil
}
