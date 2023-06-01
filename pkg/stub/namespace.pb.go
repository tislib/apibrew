// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        (unknown)
// source: stub/namespace.proto

package stub

import (
	model "github.com/apibrew/apibrew/pkg/model"
	_ "github.com/google/gnostic/openapiv3"
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

type ListNamespaceRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Token string `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
}

func (x *ListNamespaceRequest) Reset() {
	*x = ListNamespaceRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_stub_namespace_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListNamespaceRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListNamespaceRequest) ProtoMessage() {}

func (x *ListNamespaceRequest) ProtoReflect() protoreflect.Message {
	mi := &file_stub_namespace_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListNamespaceRequest.ProtoReflect.Descriptor instead.
func (*ListNamespaceRequest) Descriptor() ([]byte, []int) {
	return file_stub_namespace_proto_rawDescGZIP(), []int{0}
}

func (x *ListNamespaceRequest) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

type ListNamespaceResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Content []*model.Namespace `protobuf:"bytes,2,rep,name=content,proto3" json:"content,omitempty"`
}

func (x *ListNamespaceResponse) Reset() {
	*x = ListNamespaceResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_stub_namespace_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListNamespaceResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListNamespaceResponse) ProtoMessage() {}

func (x *ListNamespaceResponse) ProtoReflect() protoreflect.Message {
	mi := &file_stub_namespace_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListNamespaceResponse.ProtoReflect.Descriptor instead.
func (*ListNamespaceResponse) Descriptor() ([]byte, []int) {
	return file_stub_namespace_proto_rawDescGZIP(), []int{1}
}

func (x *ListNamespaceResponse) GetContent() []*model.Namespace {
	if x != nil {
		return x.Content
	}
	return nil
}

type CreateNamespaceRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Token      string             `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	Namespaces []*model.Namespace `protobuf:"bytes,2,rep,name=Namespaces,proto3" json:"Namespaces,omitempty"`
}

func (x *CreateNamespaceRequest) Reset() {
	*x = CreateNamespaceRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_stub_namespace_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateNamespaceRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateNamespaceRequest) ProtoMessage() {}

func (x *CreateNamespaceRequest) ProtoReflect() protoreflect.Message {
	mi := &file_stub_namespace_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateNamespaceRequest.ProtoReflect.Descriptor instead.
func (*CreateNamespaceRequest) Descriptor() ([]byte, []int) {
	return file_stub_namespace_proto_rawDescGZIP(), []int{2}
}

func (x *CreateNamespaceRequest) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

func (x *CreateNamespaceRequest) GetNamespaces() []*model.Namespace {
	if x != nil {
		return x.Namespaces
	}
	return nil
}

type CreateNamespaceResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Namespaces []*model.Namespace `protobuf:"bytes,1,rep,name=Namespaces,proto3" json:"Namespaces,omitempty"`
}

func (x *CreateNamespaceResponse) Reset() {
	*x = CreateNamespaceResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_stub_namespace_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateNamespaceResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateNamespaceResponse) ProtoMessage() {}

func (x *CreateNamespaceResponse) ProtoReflect() protoreflect.Message {
	mi := &file_stub_namespace_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateNamespaceResponse.ProtoReflect.Descriptor instead.
func (*CreateNamespaceResponse) Descriptor() ([]byte, []int) {
	return file_stub_namespace_proto_rawDescGZIP(), []int{3}
}

func (x *CreateNamespaceResponse) GetNamespaces() []*model.Namespace {
	if x != nil {
		return x.Namespaces
	}
	return nil
}

type UpdateNamespaceRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Token      string             `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	Namespaces []*model.Namespace `protobuf:"bytes,2,rep,name=Namespaces,proto3" json:"Namespaces,omitempty"`
}

func (x *UpdateNamespaceRequest) Reset() {
	*x = UpdateNamespaceRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_stub_namespace_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateNamespaceRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateNamespaceRequest) ProtoMessage() {}

func (x *UpdateNamespaceRequest) ProtoReflect() protoreflect.Message {
	mi := &file_stub_namespace_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateNamespaceRequest.ProtoReflect.Descriptor instead.
func (*UpdateNamespaceRequest) Descriptor() ([]byte, []int) {
	return file_stub_namespace_proto_rawDescGZIP(), []int{4}
}

func (x *UpdateNamespaceRequest) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

func (x *UpdateNamespaceRequest) GetNamespaces() []*model.Namespace {
	if x != nil {
		return x.Namespaces
	}
	return nil
}

type UpdateNamespaceResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Namespaces []*model.Namespace `protobuf:"bytes,1,rep,name=Namespaces,proto3" json:"Namespaces,omitempty"`
}

func (x *UpdateNamespaceResponse) Reset() {
	*x = UpdateNamespaceResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_stub_namespace_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateNamespaceResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateNamespaceResponse) ProtoMessage() {}

func (x *UpdateNamespaceResponse) ProtoReflect() protoreflect.Message {
	mi := &file_stub_namespace_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateNamespaceResponse.ProtoReflect.Descriptor instead.
func (*UpdateNamespaceResponse) Descriptor() ([]byte, []int) {
	return file_stub_namespace_proto_rawDescGZIP(), []int{5}
}

func (x *UpdateNamespaceResponse) GetNamespaces() []*model.Namespace {
	if x != nil {
		return x.Namespaces
	}
	return nil
}

type DeleteNamespaceRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Token string   `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	Ids   []string `protobuf:"bytes,2,rep,name=ids,proto3" json:"ids,omitempty"`
}

func (x *DeleteNamespaceRequest) Reset() {
	*x = DeleteNamespaceRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_stub_namespace_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteNamespaceRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteNamespaceRequest) ProtoMessage() {}

func (x *DeleteNamespaceRequest) ProtoReflect() protoreflect.Message {
	mi := &file_stub_namespace_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteNamespaceRequest.ProtoReflect.Descriptor instead.
func (*DeleteNamespaceRequest) Descriptor() ([]byte, []int) {
	return file_stub_namespace_proto_rawDescGZIP(), []int{6}
}

func (x *DeleteNamespaceRequest) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

func (x *DeleteNamespaceRequest) GetIds() []string {
	if x != nil {
		return x.Ids
	}
	return nil
}

type DeleteNamespaceResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Namespaces []*model.Namespace `protobuf:"bytes,1,rep,name=Namespaces,proto3" json:"Namespaces,omitempty"`
}

func (x *DeleteNamespaceResponse) Reset() {
	*x = DeleteNamespaceResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_stub_namespace_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteNamespaceResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteNamespaceResponse) ProtoMessage() {}

func (x *DeleteNamespaceResponse) ProtoReflect() protoreflect.Message {
	mi := &file_stub_namespace_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteNamespaceResponse.ProtoReflect.Descriptor instead.
func (*DeleteNamespaceResponse) Descriptor() ([]byte, []int) {
	return file_stub_namespace_proto_rawDescGZIP(), []int{7}
}

func (x *DeleteNamespaceResponse) GetNamespaces() []*model.Namespace {
	if x != nil {
		return x.Namespaces
	}
	return nil
}

type GetNamespaceRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Token string `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	Id    string `protobuf:"bytes,2,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *GetNamespaceRequest) Reset() {
	*x = GetNamespaceRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_stub_namespace_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetNamespaceRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetNamespaceRequest) ProtoMessage() {}

func (x *GetNamespaceRequest) ProtoReflect() protoreflect.Message {
	mi := &file_stub_namespace_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetNamespaceRequest.ProtoReflect.Descriptor instead.
func (*GetNamespaceRequest) Descriptor() ([]byte, []int) {
	return file_stub_namespace_proto_rawDescGZIP(), []int{8}
}

func (x *GetNamespaceRequest) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

func (x *GetNamespaceRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type GetNamespaceResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Namespace *model.Namespace `protobuf:"bytes,1,opt,name=Namespace,proto3" json:"Namespace,omitempty"`
}

func (x *GetNamespaceResponse) Reset() {
	*x = GetNamespaceResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_stub_namespace_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetNamespaceResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetNamespaceResponse) ProtoMessage() {}

func (x *GetNamespaceResponse) ProtoReflect() protoreflect.Message {
	mi := &file_stub_namespace_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetNamespaceResponse.ProtoReflect.Descriptor instead.
func (*GetNamespaceResponse) Descriptor() ([]byte, []int) {
	return file_stub_namespace_proto_rawDescGZIP(), []int{9}
}

func (x *GetNamespaceResponse) GetNamespace() *model.Namespace {
	if x != nil {
		return x.Namespace
	}
	return nil
}

var File_stub_namespace_proto protoreflect.FileDescriptor

var file_stub_namespace_proto_rawDesc = []byte{
	0x0a, 0x14, 0x73, 0x74, 0x75, 0x62, 0x2f, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04, 0x73, 0x74, 0x75, 0x62, 0x1a, 0x15, 0x6d, 0x6f,
	0x64, 0x65, 0x6c, 0x2f, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x11, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2f, 0x65, 0x72, 0x72, 0x6f, 0x72,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61,
	0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x24, 0x67, 0x6e, 0x6f, 0x73, 0x74, 0x69, 0x63, 0x2f, 0x6f, 0x70,
	0x65, 0x6e, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x33, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x2c, 0x0a, 0x14, 0x4c, 0x69,
	0x73, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x22, 0x43, 0x0a, 0x15, 0x4c, 0x69, 0x73, 0x74,
	0x4e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x2a, 0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x10, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x4e, 0x61, 0x6d, 0x65, 0x73,
	0x70, 0x61, 0x63, 0x65, 0x52, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x22, 0x60, 0x0a,
	0x16, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x30, 0x0a,
	0x0a, 0x4e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x10, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x4e, 0x61, 0x6d, 0x65, 0x73, 0x70,
	0x61, 0x63, 0x65, 0x52, 0x0a, 0x4e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x73, 0x22,
	0x4b, 0x0a, 0x17, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61,
	0x63, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x30, 0x0a, 0x0a, 0x4e, 0x61,
	0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x10,
	0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x4e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65,
	0x52, 0x0a, 0x4e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x73, 0x22, 0x60, 0x0a, 0x16,
	0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x30, 0x0a, 0x0a,
	0x4e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x10, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x4e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61,
	0x63, 0x65, 0x52, 0x0a, 0x4e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x73, 0x22, 0x4b,
	0x0a, 0x17, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63,
	0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x30, 0x0a, 0x0a, 0x4e, 0x61, 0x6d,
	0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x10, 0x2e,
	0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x4e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x52,
	0x0a, 0x4e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x73, 0x22, 0x40, 0x0a, 0x16, 0x44,
	0x65, 0x6c, 0x65, 0x74, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x10, 0x0a, 0x03, 0x69,
	0x64, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x09, 0x52, 0x03, 0x69, 0x64, 0x73, 0x22, 0x4b, 0x0a,
	0x17, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x30, 0x0a, 0x0a, 0x4e, 0x61, 0x6d, 0x65,
	0x73, 0x70, 0x61, 0x63, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x6d,
	0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x4e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x52, 0x0a,
	0x4e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x73, 0x22, 0x3b, 0x0a, 0x13, 0x47, 0x65,
	0x74, 0x4e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x46, 0x0a, 0x14, 0x47, 0x65, 0x74, 0x4e, 0x61,
	0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x2e, 0x0a, 0x09, 0x4e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x10, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x4e, 0x61, 0x6d, 0x65, 0x73,
	0x70, 0x61, 0x63, 0x65, 0x52, 0x09, 0x4e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x32,
	0xf9, 0x03, 0x0a, 0x09, 0x4e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x12, 0x64, 0x0a,
	0x06, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x12, 0x1c, 0x2e, 0x73, 0x74, 0x75, 0x62, 0x2e, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1d, 0x2e, 0x73, 0x74, 0x75, 0x62, 0x2e, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x1d, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x17, 0x3a, 0x01, 0x2a, 0x22,
	0x12, 0x2f, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x2f, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61,
	0x63, 0x65, 0x73, 0x12, 0x5b, 0x0a, 0x04, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x1a, 0x2e, 0x73, 0x74,
	0x75, 0x62, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1b, 0x2e, 0x73, 0x74, 0x75, 0x62, 0x2e, 0x4c,
	0x69, 0x73, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x1a, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x14, 0x12, 0x12, 0x2f, 0x73,
	0x79, 0x73, 0x74, 0x65, 0x6d, 0x2f, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x73,
	0x12, 0x64, 0x0a, 0x06, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x12, 0x1c, 0x2e, 0x73, 0x74, 0x75,
	0x62, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63,
	0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1d, 0x2e, 0x73, 0x74, 0x75, 0x62, 0x2e,
	0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x1d, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x17, 0x3a,
	0x01, 0x2a, 0x1a, 0x12, 0x2f, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x2f, 0x6e, 0x61, 0x6d, 0x65,
	0x73, 0x70, 0x61, 0x63, 0x65, 0x73, 0x12, 0x64, 0x0a, 0x06, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65,
	0x12, 0x1c, 0x2e, 0x73, 0x74, 0x75, 0x62, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x4e, 0x61,
	0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1d,
	0x2e, 0x73, 0x74, 0x75, 0x62, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x4e, 0x61, 0x6d, 0x65,
	0x73, 0x70, 0x61, 0x63, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x1d, 0x82,
	0xd3, 0xe4, 0x93, 0x02, 0x17, 0x3a, 0x01, 0x2a, 0x2a, 0x12, 0x2f, 0x73, 0x79, 0x73, 0x74, 0x65,
	0x6d, 0x2f, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x73, 0x12, 0x5d, 0x0a, 0x03,
	0x47, 0x65, 0x74, 0x12, 0x19, 0x2e, 0x73, 0x74, 0x75, 0x62, 0x2e, 0x47, 0x65, 0x74, 0x4e, 0x61,
	0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1a,
	0x2e, 0x73, 0x74, 0x75, 0x62, 0x2e, 0x47, 0x65, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61,
	0x63, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x1f, 0x82, 0xd3, 0xe4, 0x93,
	0x02, 0x19, 0x12, 0x17, 0x2f, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x2f, 0x6e, 0x61, 0x6d, 0x65,
	0x73, 0x70, 0x61, 0x63, 0x65, 0x73, 0x2f, 0x7b, 0x69, 0x64, 0x7d, 0x42, 0x25, 0x5a, 0x23, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x61, 0x70, 0x69, 0x62, 0x72, 0x65,
	0x77, 0x2f, 0x61, 0x70, 0x69, 0x62, 0x72, 0x65, 0x77, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x73, 0x74,
	0x75, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_stub_namespace_proto_rawDescOnce sync.Once
	file_stub_namespace_proto_rawDescData = file_stub_namespace_proto_rawDesc
)

func file_stub_namespace_proto_rawDescGZIP() []byte {
	file_stub_namespace_proto_rawDescOnce.Do(func() {
		file_stub_namespace_proto_rawDescData = protoimpl.X.CompressGZIP(file_stub_namespace_proto_rawDescData)
	})
	return file_stub_namespace_proto_rawDescData
}

var file_stub_namespace_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_stub_namespace_proto_goTypes = []interface{}{
	(*ListNamespaceRequest)(nil),    // 0: stub.ListNamespaceRequest
	(*ListNamespaceResponse)(nil),   // 1: stub.ListNamespaceResponse
	(*CreateNamespaceRequest)(nil),  // 2: stub.CreateNamespaceRequest
	(*CreateNamespaceResponse)(nil), // 3: stub.CreateNamespaceResponse
	(*UpdateNamespaceRequest)(nil),  // 4: stub.UpdateNamespaceRequest
	(*UpdateNamespaceResponse)(nil), // 5: stub.UpdateNamespaceResponse
	(*DeleteNamespaceRequest)(nil),  // 6: stub.DeleteNamespaceRequest
	(*DeleteNamespaceResponse)(nil), // 7: stub.DeleteNamespaceResponse
	(*GetNamespaceRequest)(nil),     // 8: stub.GetNamespaceRequest
	(*GetNamespaceResponse)(nil),    // 9: stub.GetNamespaceResponse
	(*model.Namespace)(nil),         // 10: model.Namespace
}
var file_stub_namespace_proto_depIdxs = []int32{
	10, // 0: stub.ListNamespaceResponse.content:type_name -> model.Namespace
	10, // 1: stub.CreateNamespaceRequest.Namespaces:type_name -> model.Namespace
	10, // 2: stub.CreateNamespaceResponse.Namespaces:type_name -> model.Namespace
	10, // 3: stub.UpdateNamespaceRequest.Namespaces:type_name -> model.Namespace
	10, // 4: stub.UpdateNamespaceResponse.Namespaces:type_name -> model.Namespace
	10, // 5: stub.DeleteNamespaceResponse.Namespaces:type_name -> model.Namespace
	10, // 6: stub.GetNamespaceResponse.Namespace:type_name -> model.Namespace
	2,  // 7: stub.Namespace.Create:input_type -> stub.CreateNamespaceRequest
	0,  // 8: stub.Namespace.List:input_type -> stub.ListNamespaceRequest
	4,  // 9: stub.Namespace.Update:input_type -> stub.UpdateNamespaceRequest
	6,  // 10: stub.Namespace.Delete:input_type -> stub.DeleteNamespaceRequest
	8,  // 11: stub.Namespace.Get:input_type -> stub.GetNamespaceRequest
	3,  // 12: stub.Namespace.Create:output_type -> stub.CreateNamespaceResponse
	1,  // 13: stub.Namespace.List:output_type -> stub.ListNamespaceResponse
	5,  // 14: stub.Namespace.Update:output_type -> stub.UpdateNamespaceResponse
	7,  // 15: stub.Namespace.Delete:output_type -> stub.DeleteNamespaceResponse
	9,  // 16: stub.Namespace.Get:output_type -> stub.GetNamespaceResponse
	12, // [12:17] is the sub-list for method output_type
	7,  // [7:12] is the sub-list for method input_type
	7,  // [7:7] is the sub-list for extension type_name
	7,  // [7:7] is the sub-list for extension extendee
	0,  // [0:7] is the sub-list for field type_name
}

func init() { file_stub_namespace_proto_init() }
func file_stub_namespace_proto_init() {
	if File_stub_namespace_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_stub_namespace_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListNamespaceRequest); i {
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
		file_stub_namespace_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListNamespaceResponse); i {
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
		file_stub_namespace_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateNamespaceRequest); i {
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
		file_stub_namespace_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateNamespaceResponse); i {
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
		file_stub_namespace_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateNamespaceRequest); i {
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
		file_stub_namespace_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateNamespaceResponse); i {
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
		file_stub_namespace_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteNamespaceRequest); i {
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
		file_stub_namespace_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteNamespaceResponse); i {
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
		file_stub_namespace_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetNamespaceRequest); i {
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
		file_stub_namespace_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetNamespaceResponse); i {
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
			RawDescriptor: file_stub_namespace_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_stub_namespace_proto_goTypes,
		DependencyIndexes: file_stub_namespace_proto_depIdxs,
		MessageInfos:      file_stub_namespace_proto_msgTypes,
	}.Build()
	File_stub_namespace_proto = out.File
	file_stub_namespace_proto_rawDesc = nil
	file_stub_namespace_proto_goTypes = nil
	file_stub_namespace_proto_depIdxs = nil
}
