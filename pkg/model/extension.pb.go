// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        (unknown)
// source: model/extension.proto

package model

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

type ExtensionOperationType int32

const (
	ExtensionOperationType_ExtensionOperationTypeCreate ExtensionOperationType = 0
	ExtensionOperationType_ExtensionOperationTypeUpdate ExtensionOperationType = 1
	ExtensionOperationType_ExtensionOperationTypeList   ExtensionOperationType = 2
	ExtensionOperationType_ExtensionOperationTypeGet    ExtensionOperationType = 3
	ExtensionOperationType_ExtensionOperationTypeDelete ExtensionOperationType = 4
)

// Enum value maps for ExtensionOperationType.
var (
	ExtensionOperationType_name = map[int32]string{
		0: "ExtensionOperationTypeCreate",
		1: "ExtensionOperationTypeUpdate",
		2: "ExtensionOperationTypeList",
		3: "ExtensionOperationTypeGet",
		4: "ExtensionOperationTypeDelete",
	}
	ExtensionOperationType_value = map[string]int32{
		"ExtensionOperationTypeCreate": 0,
		"ExtensionOperationTypeUpdate": 1,
		"ExtensionOperationTypeList":   2,
		"ExtensionOperationTypeGet":    3,
		"ExtensionOperationTypeDelete": 4,
	}
)

func (x ExtensionOperationType) Enum() *ExtensionOperationType {
	p := new(ExtensionOperationType)
	*p = x
	return p
}

func (x ExtensionOperationType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ExtensionOperationType) Descriptor() protoreflect.EnumDescriptor {
	return file_model_extension_proto_enumTypes[0].Descriptor()
}

func (ExtensionOperationType) Type() protoreflect.EnumType {
	return &file_model_extension_proto_enumTypes[0]
}

func (x ExtensionOperationType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ExtensionOperationType.Descriptor instead.
func (ExtensionOperationType) EnumDescriptor() ([]byte, []int) {
	return file_model_extension_proto_rawDescGZIP(), []int{0}
}

type ExtensionOperationStep int32

const (
	ExtensionOperationStep_ExtensionOperationStepBefore  ExtensionOperationStep = 0
	ExtensionOperationStep_ExtensionOperationStepInstead ExtensionOperationStep = 1
	ExtensionOperationStep_ExtensionOperationStepAfter   ExtensionOperationStep = 2
)

// Enum value maps for ExtensionOperationStep.
var (
	ExtensionOperationStep_name = map[int32]string{
		0: "ExtensionOperationStepBefore",
		1: "ExtensionOperationStepInstead",
		2: "ExtensionOperationStepAfter",
	}
	ExtensionOperationStep_value = map[string]int32{
		"ExtensionOperationStepBefore":  0,
		"ExtensionOperationStepInstead": 1,
		"ExtensionOperationStepAfter":   2,
	}
)

func (x ExtensionOperationStep) Enum() *ExtensionOperationStep {
	p := new(ExtensionOperationStep)
	*p = x
	return p
}

func (x ExtensionOperationStep) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ExtensionOperationStep) Descriptor() protoreflect.EnumDescriptor {
	return file_model_extension_proto_enumTypes[1].Descriptor()
}

func (ExtensionOperationStep) Type() protoreflect.EnumType {
	return &file_model_extension_proto_enumTypes[1]
}

func (x ExtensionOperationStep) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ExtensionOperationStep.Descriptor instead.
func (ExtensionOperationStep) EnumDescriptor() ([]byte, []int) {
	return file_model_extension_proto_rawDescGZIP(), []int{1}
}

type ExtensionServer struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Host string `protobuf:"bytes,1,opt,name=host,proto3" json:"host,omitempty"`
	Port int32  `protobuf:"varint,2,opt,name=port,proto3" json:"port,omitempty"`
}

func (x *ExtensionServer) Reset() {
	*x = ExtensionServer{}
	if protoimpl.UnsafeEnabled {
		mi := &file_model_extension_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ExtensionServer) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ExtensionServer) ProtoMessage() {}

func (x *ExtensionServer) ProtoReflect() protoreflect.Message {
	mi := &file_model_extension_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ExtensionServer.ProtoReflect.Descriptor instead.
func (*ExtensionServer) Descriptor() ([]byte, []int) {
	return file_model_extension_proto_rawDescGZIP(), []int{0}
}

func (x *ExtensionServer) GetHost() string {
	if x != nil {
		return x.Host
	}
	return ""
}

func (x *ExtensionServer) GetPort() int32 {
	if x != nil {
		return x.Port
	}
	return 0
}

type ExtensionOperation struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	OperationType ExtensionOperationType `protobuf:"varint,1,opt,name=operationType,proto3,enum=model.ExtensionOperationType" json:"operationType,omitempty"`
	Step          ExtensionOperationStep `protobuf:"varint,2,opt,name=step,proto3,enum=model.ExtensionOperationStep" json:"step,omitempty"`
	Sync          bool                   `protobuf:"varint,3,opt,name=sync,proto3" json:"sync,omitempty"` // sync operation is always true for instead step
}

func (x *ExtensionOperation) Reset() {
	*x = ExtensionOperation{}
	if protoimpl.UnsafeEnabled {
		mi := &file_model_extension_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ExtensionOperation) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ExtensionOperation) ProtoMessage() {}

func (x *ExtensionOperation) ProtoReflect() protoreflect.Message {
	mi := &file_model_extension_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ExtensionOperation.ProtoReflect.Descriptor instead.
func (*ExtensionOperation) Descriptor() ([]byte, []int) {
	return file_model_extension_proto_rawDescGZIP(), []int{1}
}

func (x *ExtensionOperation) GetOperationType() ExtensionOperationType {
	if x != nil {
		return x.OperationType
	}
	return ExtensionOperationType_ExtensionOperationTypeCreate
}

func (x *ExtensionOperation) GetStep() ExtensionOperationStep {
	if x != nil {
		return x.Step
	}
	return ExtensionOperationStep_ExtensionOperationStepBefore
}

func (x *ExtensionOperation) GetSync() bool {
	if x != nil {
		return x.Sync
	}
	return false
}

type ExtensionConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Namespace  string                `protobuf:"bytes,5,opt,name=namespace,proto3" json:"namespace,omitempty"`
	Resource   string                `protobuf:"bytes,6,opt,name=resource,proto3" json:"resource,omitempty"`
	Operations []*ExtensionOperation `protobuf:"bytes,8,rep,name=operations,proto3" json:"operations,omitempty"`
}

func (x *ExtensionConfig) Reset() {
	*x = ExtensionConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_model_extension_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ExtensionConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ExtensionConfig) ProtoMessage() {}

func (x *ExtensionConfig) ProtoReflect() protoreflect.Message {
	mi := &file_model_extension_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ExtensionConfig.ProtoReflect.Descriptor instead.
func (*ExtensionConfig) Descriptor() ([]byte, []int) {
	return file_model_extension_proto_rawDescGZIP(), []int{2}
}

func (x *ExtensionConfig) GetNamespace() string {
	if x != nil {
		return x.Namespace
	}
	return ""
}

func (x *ExtensionConfig) GetResource() string {
	if x != nil {
		return x.Resource
	}
	return ""
}

func (x *ExtensionConfig) GetOperations() []*ExtensionOperation {
	if x != nil {
		return x.Operations
	}
	return nil
}

type RemoteExtension struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          string           `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name        string           `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Description string           `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	Type        DataType         `protobuf:"varint,4,opt,name=type,proto3,enum=model.DataType" json:"type,omitempty"`
	Server      *ExtensionServer `protobuf:"bytes,7,opt,name=server,proto3" json:"server,omitempty"`
	Config      *ExtensionConfig `protobuf:"bytes,8,opt,name=config,proto3" json:"config,omitempty"`
	AuditData   *AuditData       `protobuf:"bytes,101,opt,name=auditData,proto3" json:"auditData,omitempty"`
	Version     uint32           `protobuf:"varint,102,opt,name=version,proto3" json:"version,omitempty"`
}

func (x *RemoteExtension) Reset() {
	*x = RemoteExtension{}
	if protoimpl.UnsafeEnabled {
		mi := &file_model_extension_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RemoteExtension) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RemoteExtension) ProtoMessage() {}

func (x *RemoteExtension) ProtoReflect() protoreflect.Message {
	mi := &file_model_extension_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RemoteExtension.ProtoReflect.Descriptor instead.
func (*RemoteExtension) Descriptor() ([]byte, []int) {
	return file_model_extension_proto_rawDescGZIP(), []int{3}
}

func (x *RemoteExtension) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *RemoteExtension) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *RemoteExtension) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *RemoteExtension) GetType() DataType {
	if x != nil {
		return x.Type
	}
	return DataType_USER
}

func (x *RemoteExtension) GetServer() *ExtensionServer {
	if x != nil {
		return x.Server
	}
	return nil
}

func (x *RemoteExtension) GetConfig() *ExtensionConfig {
	if x != nil {
		return x.Config
	}
	return nil
}

func (x *RemoteExtension) GetAuditData() *AuditData {
	if x != nil {
		return x.AuditData
	}
	return nil
}

func (x *RemoteExtension) GetVersion() uint32 {
	if x != nil {
		return x.Version
	}
	return 0
}

var File_model_extension_proto protoreflect.FileDescriptor

var file_model_extension_proto_rawDesc = []byte{
	0x0a, 0x15, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2f, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f,
	0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x1a, 0x11,
	0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2f, 0x61, 0x75, 0x64, 0x69, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x12, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x39, 0x0a, 0x0f, 0x45, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69,
	0x6f, 0x6e, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x12, 0x12, 0x0a, 0x04, 0x68, 0x6f, 0x73, 0x74,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x68, 0x6f, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04,
	0x70, 0x6f, 0x72, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x70, 0x6f, 0x72, 0x74,
	0x22, 0xa0, 0x01, 0x0a, 0x12, 0x45, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x4f, 0x70,
	0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x43, 0x0a, 0x0d, 0x6f, 0x70, 0x65, 0x72, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1d,
	0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x45, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e,
	0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x52, 0x0d, 0x6f,
	0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x12, 0x31, 0x0a, 0x04,
	0x73, 0x74, 0x65, 0x70, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1d, 0x2e, 0x6d, 0x6f, 0x64,
	0x65, 0x6c, 0x2e, 0x45, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x4f, 0x70, 0x65, 0x72,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x74, 0x65, 0x70, 0x52, 0x04, 0x73, 0x74, 0x65, 0x70, 0x12,
	0x12, 0x0a, 0x04, 0x73, 0x79, 0x6e, 0x63, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x04, 0x73,
	0x79, 0x6e, 0x63, 0x22, 0x86, 0x01, 0x0a, 0x0f, 0x45, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f,
	0x6e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x1c, 0x0a, 0x09, 0x6e, 0x61, 0x6d, 0x65, 0x73,
	0x70, 0x61, 0x63, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x6e, 0x61, 0x6d, 0x65,
	0x73, 0x70, 0x61, 0x63, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x12, 0x39, 0x0a, 0x0a, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18,
	0x08, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x45, 0x78,
	0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x52, 0x0a, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x22, 0xa6, 0x02, 0x0a,
	0x0f, 0x52, 0x65, 0x6d, 0x6f, 0x74, 0x65, 0x45, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e,
	0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64,
	0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72,
	0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x23, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x0e, 0x32, 0x0f, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x44, 0x61, 0x74,
	0x61, 0x54, 0x79, 0x70, 0x65, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x2e, 0x0a, 0x06, 0x73,
	0x65, 0x72, 0x76, 0x65, 0x72, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x6d, 0x6f,
	0x64, 0x65, 0x6c, 0x2e, 0x45, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x53, 0x65, 0x72,
	0x76, 0x65, 0x72, 0x52, 0x06, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x12, 0x2e, 0x0a, 0x06, 0x63,
	0x6f, 0x6e, 0x66, 0x69, 0x67, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x6d, 0x6f,
	0x64, 0x65, 0x6c, 0x2e, 0x45, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x43, 0x6f, 0x6e,
	0x66, 0x69, 0x67, 0x52, 0x06, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x2e, 0x0a, 0x09, 0x61,
	0x75, 0x64, 0x69, 0x74, 0x44, 0x61, 0x74, 0x61, 0x18, 0x65, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10,
	0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x41, 0x75, 0x64, 0x69, 0x74, 0x44, 0x61, 0x74, 0x61,
	0x52, 0x09, 0x61, 0x75, 0x64, 0x69, 0x74, 0x44, 0x61, 0x74, 0x61, 0x12, 0x18, 0x0a, 0x07, 0x76,
	0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x66, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x07, 0x76, 0x65,
	0x72, 0x73, 0x69, 0x6f, 0x6e, 0x2a, 0xbd, 0x01, 0x0a, 0x16, 0x45, 0x78, 0x74, 0x65, 0x6e, 0x73,
	0x69, 0x6f, 0x6e, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65,
	0x12, 0x20, 0x0a, 0x1c, 0x45, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x4f, 0x70, 0x65,
	0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x10, 0x00, 0x12, 0x20, 0x0a, 0x1c, 0x45, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x4f,
	0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x55, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x10, 0x01, 0x12, 0x1e, 0x0a, 0x1a, 0x45, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f,
	0x6e, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x4c, 0x69,
	0x73, 0x74, 0x10, 0x02, 0x12, 0x1d, 0x0a, 0x19, 0x45, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f,
	0x6e, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x47, 0x65,
	0x74, 0x10, 0x03, 0x12, 0x20, 0x0a, 0x1c, 0x45, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e,
	0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x44, 0x65, 0x6c,
	0x65, 0x74, 0x65, 0x10, 0x04, 0x2a, 0x7e, 0x0a, 0x16, 0x45, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69,
	0x6f, 0x6e, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x74, 0x65, 0x70, 0x12,
	0x20, 0x0a, 0x1c, 0x45, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x4f, 0x70, 0x65, 0x72,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x74, 0x65, 0x70, 0x42, 0x65, 0x66, 0x6f, 0x72, 0x65, 0x10,
	0x00, 0x12, 0x21, 0x0a, 0x1d, 0x45, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x4f, 0x70,
	0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x74, 0x65, 0x70, 0x49, 0x6e, 0x73, 0x74, 0x65,
	0x61, 0x64, 0x10, 0x01, 0x12, 0x1f, 0x0a, 0x1b, 0x45, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f,
	0x6e, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x74, 0x65, 0x70, 0x41, 0x66,
	0x74, 0x65, 0x72, 0x10, 0x02, 0x42, 0x2a, 0x5a, 0x28, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x74, 0x69, 0x73, 0x6c, 0x69, 0x62, 0x2f, 0x64, 0x61, 0x74, 0x61, 0x2d,
	0x68, 0x61, 0x6e, 0x64, 0x6c, 0x65, 0x72, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x6d, 0x6f, 0x64, 0x65,
	0x6c, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_model_extension_proto_rawDescOnce sync.Once
	file_model_extension_proto_rawDescData = file_model_extension_proto_rawDesc
)

func file_model_extension_proto_rawDescGZIP() []byte {
	file_model_extension_proto_rawDescOnce.Do(func() {
		file_model_extension_proto_rawDescData = protoimpl.X.CompressGZIP(file_model_extension_proto_rawDescData)
	})
	return file_model_extension_proto_rawDescData
}

var file_model_extension_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_model_extension_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_model_extension_proto_goTypes = []interface{}{
	(ExtensionOperationType)(0), // 0: model.ExtensionOperationType
	(ExtensionOperationStep)(0), // 1: model.ExtensionOperationStep
	(*ExtensionServer)(nil),     // 2: model.ExtensionServer
	(*ExtensionOperation)(nil),  // 3: model.ExtensionOperation
	(*ExtensionConfig)(nil),     // 4: model.ExtensionConfig
	(*RemoteExtension)(nil),     // 5: model.RemoteExtension
	(DataType)(0),               // 6: model.DataType
	(*AuditData)(nil),           // 7: model.AuditData
}
var file_model_extension_proto_depIdxs = []int32{
	0, // 0: model.ExtensionOperation.operationType:type_name -> model.ExtensionOperationType
	1, // 1: model.ExtensionOperation.step:type_name -> model.ExtensionOperationStep
	3, // 2: model.ExtensionConfig.operations:type_name -> model.ExtensionOperation
	6, // 3: model.RemoteExtension.type:type_name -> model.DataType
	2, // 4: model.RemoteExtension.server:type_name -> model.ExtensionServer
	4, // 5: model.RemoteExtension.config:type_name -> model.ExtensionConfig
	7, // 6: model.RemoteExtension.auditData:type_name -> model.AuditData
	7, // [7:7] is the sub-list for method output_type
	7, // [7:7] is the sub-list for method input_type
	7, // [7:7] is the sub-list for extension type_name
	7, // [7:7] is the sub-list for extension extendee
	0, // [0:7] is the sub-list for field type_name
}

func init() { file_model_extension_proto_init() }
func file_model_extension_proto_init() {
	if File_model_extension_proto != nil {
		return
	}
	file_model_audit_proto_init()
	file_model_common_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_model_extension_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ExtensionServer); i {
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
		file_model_extension_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ExtensionOperation); i {
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
		file_model_extension_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ExtensionConfig); i {
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
		file_model_extension_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RemoteExtension); i {
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
			RawDescriptor: file_model_extension_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_model_extension_proto_goTypes,
		DependencyIndexes: file_model_extension_proto_depIdxs,
		EnumInfos:         file_model_extension_proto_enumTypes,
		MessageInfos:      file_model_extension_proto_msgTypes,
	}.Build()
	File_model_extension_proto = out.File
	file_model_extension_proto_rawDesc = nil
	file_model_extension_proto_goTypes = nil
	file_model_extension_proto_depIdxs = nil
}
