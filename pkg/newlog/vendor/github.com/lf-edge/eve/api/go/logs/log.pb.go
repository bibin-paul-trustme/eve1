// Copyright(c) 2017-2018 Zededa, Inc.
// All rights reserved.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0-devel
// 	protoc        v3.12.2
// source: logs/log.proto

package logs

import (
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
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

type LogEntry struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Severity  string               `protobuf:"bytes,1,opt,name=severity,proto3" json:"severity,omitempty"`                                                                                 // e.g., INFO, DEBUG, ERROR etc.
	Source    string               `protobuf:"bytes,2,opt,name=source,proto3" json:"source,omitempty"`                                                                                     // Source of the msg, zedmanager etc.
	Iid       string               `protobuf:"bytes,3,opt,name=iid,proto3" json:"iid,omitempty"`                                                                                           // instance ID of the source (e.g., PID)
	Content   string               `protobuf:"bytes,4,opt,name=content,proto3" json:"content,omitempty"`                                                                                   // actual log message
	Msgid     uint64               `protobuf:"varint,5,opt,name=msgid,proto3" json:"msgid,omitempty"`                                                                                      // monotonically increasing number (detect drops)
	Tags      map[string]string    `protobuf:"bytes,6,rep,name=tags,proto3" json:"tags,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"` // additional meta info <key,value>
	Timestamp *timestamp.Timestamp `protobuf:"bytes,7,opt,name=timestamp,proto3" json:"timestamp,omitempty"`                                                                               // timestamp of the msg
	Filename  string               `protobuf:"bytes,8,opt,name=filename,proto3" json:"filename,omitempty"`
	Function  string               `protobuf:"bytes,9,opt,name=function,proto3" json:"function,omitempty"`
}

func (x *LogEntry) Reset() {
	*x = LogEntry{}
	if protoimpl.UnsafeEnabled {
		mi := &file_logs_log_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LogEntry) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LogEntry) ProtoMessage() {}

func (x *LogEntry) ProtoReflect() protoreflect.Message {
	mi := &file_logs_log_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LogEntry.ProtoReflect.Descriptor instead.
func (*LogEntry) Descriptor() ([]byte, []int) {
	return file_logs_log_proto_rawDescGZIP(), []int{0}
}

func (x *LogEntry) GetSeverity() string {
	if x != nil {
		return x.Severity
	}
	return ""
}

func (x *LogEntry) GetSource() string {
	if x != nil {
		return x.Source
	}
	return ""
}

func (x *LogEntry) GetIid() string {
	if x != nil {
		return x.Iid
	}
	return ""
}

func (x *LogEntry) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

func (x *LogEntry) GetMsgid() uint64 {
	if x != nil {
		return x.Msgid
	}
	return 0
}

func (x *LogEntry) GetTags() map[string]string {
	if x != nil {
		return x.Tags
	}
	return nil
}

func (x *LogEntry) GetTimestamp() *timestamp.Timestamp {
	if x != nil {
		return x.Timestamp
	}
	return nil
}

func (x *LogEntry) GetFilename() string {
	if x != nil {
		return x.Filename
	}
	return ""
}

func (x *LogEntry) GetFunction() string {
	if x != nil {
		return x.Function
	}
	return ""
}

// This is the request payload for POST /api/v1/edgeDevice/logs
// ZInfoMsg carries device logs to the controller.
// The messages need to be retransmitted until they make it to the controller.
// The message is assumed to be protected by a TLS session bound to the
// device certificate.
type LogBundle struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DevID      string               `protobuf:"bytes,1,opt,name=devID,proto3" json:"devID,omitempty"`           // Device UUID
	Image      string               `protobuf:"bytes,2,opt,name=image,proto3" json:"image,omitempty"`           // SW image the log got emitted from
	Log        []*LogEntry          `protobuf:"bytes,3,rep,name=log,proto3" json:"log,omitempty"`               // Log entries
	Timestamp  *timestamp.Timestamp `protobuf:"bytes,4,opt,name=timestamp,proto3" json:"timestamp,omitempty"`   // upload timestamp
	EveVersion string               `protobuf:"bytes,5,opt,name=eveVersion,proto3" json:"eveVersion,omitempty"` // EVE software version
}

func (x *LogBundle) Reset() {
	*x = LogBundle{}
	if protoimpl.UnsafeEnabled {
		mi := &file_logs_log_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LogBundle) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LogBundle) ProtoMessage() {}

func (x *LogBundle) ProtoReflect() protoreflect.Message {
	mi := &file_logs_log_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LogBundle.ProtoReflect.Descriptor instead.
func (*LogBundle) Descriptor() ([]byte, []int) {
	return file_logs_log_proto_rawDescGZIP(), []int{1}
}

func (x *LogBundle) GetDevID() string {
	if x != nil {
		return x.DevID
	}
	return ""
}

func (x *LogBundle) GetImage() string {
	if x != nil {
		return x.Image
	}
	return ""
}

func (x *LogBundle) GetLog() []*LogEntry {
	if x != nil {
		return x.Log
	}
	return nil
}

func (x *LogBundle) GetTimestamp() *timestamp.Timestamp {
	if x != nil {
		return x.Timestamp
	}
	return nil
}

func (x *LogBundle) GetEveVersion() string {
	if x != nil {
		return x.EveVersion
	}
	return ""
}

// This is the request payload for POST /api/v1/edgeDevice/apps/instances/id/<app-instance-uuid>/logs
type AppInstanceLogBundle struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Log       []*LogEntry          `protobuf:"bytes,1,rep,name=log,proto3" json:"log,omitempty"`             // Log entries
	Timestamp *timestamp.Timestamp `protobuf:"bytes,2,opt,name=timestamp,proto3" json:"timestamp,omitempty"` // upload timestamp
}

func (x *AppInstanceLogBundle) Reset() {
	*x = AppInstanceLogBundle{}
	if protoimpl.UnsafeEnabled {
		mi := &file_logs_log_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AppInstanceLogBundle) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AppInstanceLogBundle) ProtoMessage() {}

func (x *AppInstanceLogBundle) ProtoReflect() protoreflect.Message {
	mi := &file_logs_log_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AppInstanceLogBundle.ProtoReflect.Descriptor instead.
func (*AppInstanceLogBundle) Descriptor() ([]byte, []int) {
	return file_logs_log_proto_rawDescGZIP(), []int{2}
}

func (x *AppInstanceLogBundle) GetLog() []*LogEntry {
	if x != nil {
		return x.Log
	}
	return nil
}

func (x *AppInstanceLogBundle) GetTimestamp() *timestamp.Timestamp {
	if x != nil {
		return x.Timestamp
	}
	return nil
}

// This is the reply payload for device log processing from controller
type ServerMetrics struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CpuPercentage       float32 `protobuf:"fixed32,1,opt,name=cpu_percentage,json=cpuPercentage,proto3" json:"cpu_percentage,omitempty"`                      // controller CPU usage in percentage
	LogProcessDelayMsec uint32  `protobuf:"varint,2,opt,name=log_process_delay_msec,json=logProcessDelayMsec,proto3" json:"log_process_delay_msec,omitempty"` // log messages processing delay in msec
}

func (x *ServerMetrics) Reset() {
	*x = ServerMetrics{}
	if protoimpl.UnsafeEnabled {
		mi := &file_logs_log_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ServerMetrics) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ServerMetrics) ProtoMessage() {}

func (x *ServerMetrics) ProtoReflect() protoreflect.Message {
	mi := &file_logs_log_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ServerMetrics.ProtoReflect.Descriptor instead.
func (*ServerMetrics) Descriptor() ([]byte, []int) {
	return file_logs_log_proto_rawDescGZIP(), []int{3}
}

func (x *ServerMetrics) GetCpuPercentage() float32 {
	if x != nil {
		return x.CpuPercentage
	}
	return 0
}

func (x *ServerMetrics) GetLogProcessDelayMsec() uint32 {
	if x != nil {
		return x.LogProcessDelayMsec
	}
	return 0
}

var File_logs_log_proto protoreflect.FileDescriptor

var file_logs_log_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x6c, 0x6f, 0x67, 0x73, 0x2f, 0x6c, 0x6f, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x13, 0x6f, 0x72, 0x67, 0x2e, 0x6c, 0x66, 0x65, 0x64, 0x67, 0x65, 0x2e, 0x65, 0x76, 0x65,
	0x2e, 0x6c, 0x6f, 0x67, 0x73, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xe8, 0x02, 0x0a, 0x08, 0x4c, 0x6f, 0x67, 0x45, 0x6e,
	0x74, 0x72, 0x79, 0x12, 0x1a, 0x0a, 0x08, 0x73, 0x65, 0x76, 0x65, 0x72, 0x69, 0x74, 0x79, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x73, 0x65, 0x76, 0x65, 0x72, 0x69, 0x74, 0x79, 0x12,
	0x16, 0x0a, 0x06, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x06, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x69, 0x69, 0x64, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x69, 0x69, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6e,
	0x74, 0x65, 0x6e, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f, 0x6e, 0x74,
	0x65, 0x6e, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x6d, 0x73, 0x67, 0x69, 0x64, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x04, 0x52, 0x05, 0x6d, 0x73, 0x67, 0x69, 0x64, 0x12, 0x3b, 0x0a, 0x04, 0x74, 0x61, 0x67,
	0x73, 0x18, 0x06, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x27, 0x2e, 0x6f, 0x72, 0x67, 0x2e, 0x6c, 0x66,
	0x65, 0x64, 0x67, 0x65, 0x2e, 0x65, 0x76, 0x65, 0x2e, 0x6c, 0x6f, 0x67, 0x73, 0x2e, 0x4c, 0x6f,
	0x67, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x2e, 0x54, 0x61, 0x67, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79,
	0x52, 0x04, 0x74, 0x61, 0x67, 0x73, 0x12, 0x38, 0x0a, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74,
	0x61, 0x6d, 0x70, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65,
	0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70,
	0x12, 0x1a, 0x0a, 0x08, 0x66, 0x69, 0x6c, 0x65, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x08, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x08, 0x66, 0x69, 0x6c, 0x65, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08,
	0x66, 0x75, 0x6e, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x66, 0x75, 0x6e, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x1a, 0x37, 0x0a, 0x09, 0x54, 0x61, 0x67, 0x73,
	0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38,
	0x01, 0x22, 0xc2, 0x01, 0x0a, 0x09, 0x4c, 0x6f, 0x67, 0x42, 0x75, 0x6e, 0x64, 0x6c, 0x65, 0x12,
	0x14, 0x0a, 0x05, 0x64, 0x65, 0x76, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x64, 0x65, 0x76, 0x49, 0x44, 0x12, 0x14, 0x0a, 0x05, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x12, 0x2f, 0x0a, 0x03, 0x6c,
	0x6f, 0x67, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x6f, 0x72, 0x67, 0x2e, 0x6c,
	0x66, 0x65, 0x64, 0x67, 0x65, 0x2e, 0x65, 0x76, 0x65, 0x2e, 0x6c, 0x6f, 0x67, 0x73, 0x2e, 0x4c,
	0x6f, 0x67, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x03, 0x6c, 0x6f, 0x67, 0x12, 0x38, 0x0a, 0x09,
	0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x74, 0x69, 0x6d,
	0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x12, 0x1e, 0x0a, 0x0a, 0x65, 0x76, 0x65, 0x56, 0x65, 0x72,
	0x73, 0x69, 0x6f, 0x6e, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x65, 0x76, 0x65, 0x56,
	0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x22, 0x81, 0x01, 0x0a, 0x14, 0x41, 0x70, 0x70, 0x49, 0x6e,
	0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x4c, 0x6f, 0x67, 0x42, 0x75, 0x6e, 0x64, 0x6c, 0x65, 0x12,
	0x2f, 0x0a, 0x03, 0x6c, 0x6f, 0x67, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x6f,
	0x72, 0x67, 0x2e, 0x6c, 0x66, 0x65, 0x64, 0x67, 0x65, 0x2e, 0x65, 0x76, 0x65, 0x2e, 0x6c, 0x6f,
	0x67, 0x73, 0x2e, 0x4c, 0x6f, 0x67, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x03, 0x6c, 0x6f, 0x67,
	0x12, 0x38, 0x0a, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52,
	0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x22, 0x6b, 0x0a, 0x0d, 0x53, 0x65,
	0x72, 0x76, 0x65, 0x72, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73, 0x12, 0x25, 0x0a, 0x0e, 0x63,
	0x70, 0x75, 0x5f, 0x70, 0x65, 0x72, 0x63, 0x65, 0x6e, 0x74, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x02, 0x52, 0x0d, 0x63, 0x70, 0x75, 0x50, 0x65, 0x72, 0x63, 0x65, 0x6e, 0x74, 0x61,
	0x67, 0x65, 0x12, 0x33, 0x0a, 0x16, 0x6c, 0x6f, 0x67, 0x5f, 0x70, 0x72, 0x6f, 0x63, 0x65, 0x73,
	0x73, 0x5f, 0x64, 0x65, 0x6c, 0x61, 0x79, 0x5f, 0x6d, 0x73, 0x65, 0x63, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0d, 0x52, 0x13, 0x6c, 0x6f, 0x67, 0x50, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x44, 0x65,
	0x6c, 0x61, 0x79, 0x4d, 0x73, 0x65, 0x63, 0x42, 0x39, 0x0a, 0x13, 0x6f, 0x72, 0x67, 0x2e, 0x6c,
	0x66, 0x65, 0x64, 0x67, 0x65, 0x2e, 0x65, 0x76, 0x65, 0x2e, 0x6c, 0x6f, 0x67, 0x73, 0x5a, 0x22,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6c, 0x66, 0x2d, 0x65, 0x64,
	0x67, 0x65, 0x2f, 0x65, 0x76, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x67, 0x6f, 0x2f, 0x6c, 0x6f,
	0x67, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_logs_log_proto_rawDescOnce sync.Once
	file_logs_log_proto_rawDescData = file_logs_log_proto_rawDesc
)

func file_logs_log_proto_rawDescGZIP() []byte {
	file_logs_log_proto_rawDescOnce.Do(func() {
		file_logs_log_proto_rawDescData = protoimpl.X.CompressGZIP(file_logs_log_proto_rawDescData)
	})
	return file_logs_log_proto_rawDescData
}

var file_logs_log_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_logs_log_proto_goTypes = []interface{}{
	(*LogEntry)(nil),             // 0: org.lfedge.eve.logs.LogEntry
	(*LogBundle)(nil),            // 1: org.lfedge.eve.logs.LogBundle
	(*AppInstanceLogBundle)(nil), // 2: org.lfedge.eve.logs.AppInstanceLogBundle
	(*ServerMetrics)(nil),        // 3: org.lfedge.eve.logs.ServerMetrics
	nil,                          // 4: org.lfedge.eve.logs.LogEntry.TagsEntry
	(*timestamp.Timestamp)(nil),  // 5: google.protobuf.Timestamp
}
var file_logs_log_proto_depIdxs = []int32{
	4, // 0: org.lfedge.eve.logs.LogEntry.tags:type_name -> org.lfedge.eve.logs.LogEntry.TagsEntry
	5, // 1: org.lfedge.eve.logs.LogEntry.timestamp:type_name -> google.protobuf.Timestamp
	0, // 2: org.lfedge.eve.logs.LogBundle.log:type_name -> org.lfedge.eve.logs.LogEntry
	5, // 3: org.lfedge.eve.logs.LogBundle.timestamp:type_name -> google.protobuf.Timestamp
	0, // 4: org.lfedge.eve.logs.AppInstanceLogBundle.log:type_name -> org.lfedge.eve.logs.LogEntry
	5, // 5: org.lfedge.eve.logs.AppInstanceLogBundle.timestamp:type_name -> google.protobuf.Timestamp
	6, // [6:6] is the sub-list for method output_type
	6, // [6:6] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_logs_log_proto_init() }
func file_logs_log_proto_init() {
	if File_logs_log_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_logs_log_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LogEntry); i {
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
		file_logs_log_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LogBundle); i {
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
		file_logs_log_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AppInstanceLogBundle); i {
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
		file_logs_log_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ServerMetrics); i {
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
			RawDescriptor: file_logs_log_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_logs_log_proto_goTypes,
		DependencyIndexes: file_logs_log_proto_depIdxs,
		MessageInfos:      file_logs_log_proto_msgTypes,
	}.Build()
	File_logs_log_proto = out.File
	file_logs_log_proto_rawDesc = nil
	file_logs_log_proto_goTypes = nil
	file_logs_log_proto_depIdxs = nil
}
