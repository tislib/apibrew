// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        (unknown)
// source: stub/node.proto

package stub

import (
	model "github.com/tislib/apibrew/pkg/model"
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

type NodeStatusRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Token string `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	Id    string `protobuf:"bytes,2,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *NodeStatusRequest) Reset() {
	*x = NodeStatusRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_stub_node_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NodeStatusRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NodeStatusRequest) ProtoMessage() {}

func (x *NodeStatusRequest) ProtoReflect() protoreflect.Message {
	mi := &file_stub_node_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NodeStatusRequest.ProtoReflect.Descriptor instead.
func (*NodeStatusRequest) Descriptor() ([]byte, []int) {
	return file_stub_node_proto_rawDescGZIP(), []int{0}
}

func (x *NodeStatusRequest) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

func (x *NodeStatusRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type NodeStatusResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ConnectionAlreadyInitiated bool `protobuf:"varint,1,opt,name=connectionAlreadyInitiated,proto3" json:"connectionAlreadyInitiated,omitempty"`
	TestConnection             bool `protobuf:"varint,2,opt,name=testConnection,proto3" json:"testConnection,omitempty"`
}

func (x *NodeStatusResponse) Reset() {
	*x = NodeStatusResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_stub_node_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NodeStatusResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NodeStatusResponse) ProtoMessage() {}

func (x *NodeStatusResponse) ProtoReflect() protoreflect.Message {
	mi := &file_stub_node_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NodeStatusResponse.ProtoReflect.Descriptor instead.
func (*NodeStatusResponse) Descriptor() ([]byte, []int) {
	return file_stub_node_proto_rawDescGZIP(), []int{1}
}

func (x *NodeStatusResponse) GetConnectionAlreadyInitiated() bool {
	if x != nil {
		return x.ConnectionAlreadyInitiated
	}
	return false
}

func (x *NodeStatusResponse) GetTestConnection() bool {
	if x != nil {
		return x.TestConnection
	}
	return false
}

type ListNodeRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Token string `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
}

func (x *ListNodeRequest) Reset() {
	*x = ListNodeRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_stub_node_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListNodeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListNodeRequest) ProtoMessage() {}

func (x *ListNodeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_stub_node_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListNodeRequest.ProtoReflect.Descriptor instead.
func (*ListNodeRequest) Descriptor() ([]byte, []int) {
	return file_stub_node_proto_rawDescGZIP(), []int{2}
}

func (x *ListNodeRequest) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

type ListNodeResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Content []*model.Node `protobuf:"bytes,2,rep,name=content,proto3" json:"content,omitempty"`
}

func (x *ListNodeResponse) Reset() {
	*x = ListNodeResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_stub_node_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListNodeResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListNodeResponse) ProtoMessage() {}

func (x *ListNodeResponse) ProtoReflect() protoreflect.Message {
	mi := &file_stub_node_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListNodeResponse.ProtoReflect.Descriptor instead.
func (*ListNodeResponse) Descriptor() ([]byte, []int) {
	return file_stub_node_proto_rawDescGZIP(), []int{3}
}

func (x *ListNodeResponse) GetContent() []*model.Node {
	if x != nil {
		return x.Content
	}
	return nil
}

type CreateNodeRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Token string        `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	Nodes []*model.Node `protobuf:"bytes,2,rep,name=Nodes,proto3" json:"Nodes,omitempty"`
}

func (x *CreateNodeRequest) Reset() {
	*x = CreateNodeRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_stub_node_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateNodeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateNodeRequest) ProtoMessage() {}

func (x *CreateNodeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_stub_node_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateNodeRequest.ProtoReflect.Descriptor instead.
func (*CreateNodeRequest) Descriptor() ([]byte, []int) {
	return file_stub_node_proto_rawDescGZIP(), []int{4}
}

func (x *CreateNodeRequest) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

func (x *CreateNodeRequest) GetNodes() []*model.Node {
	if x != nil {
		return x.Nodes
	}
	return nil
}

type CreateNodeResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Nodes []*model.Node `protobuf:"bytes,1,rep,name=Nodes,proto3" json:"Nodes,omitempty"`
}

func (x *CreateNodeResponse) Reset() {
	*x = CreateNodeResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_stub_node_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateNodeResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateNodeResponse) ProtoMessage() {}

func (x *CreateNodeResponse) ProtoReflect() protoreflect.Message {
	mi := &file_stub_node_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateNodeResponse.ProtoReflect.Descriptor instead.
func (*CreateNodeResponse) Descriptor() ([]byte, []int) {
	return file_stub_node_proto_rawDescGZIP(), []int{5}
}

func (x *CreateNodeResponse) GetNodes() []*model.Node {
	if x != nil {
		return x.Nodes
	}
	return nil
}

type UpdateNodeRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Token string        `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	Nodes []*model.Node `protobuf:"bytes,2,rep,name=Nodes,proto3" json:"Nodes,omitempty"`
}

func (x *UpdateNodeRequest) Reset() {
	*x = UpdateNodeRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_stub_node_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateNodeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateNodeRequest) ProtoMessage() {}

func (x *UpdateNodeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_stub_node_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateNodeRequest.ProtoReflect.Descriptor instead.
func (*UpdateNodeRequest) Descriptor() ([]byte, []int) {
	return file_stub_node_proto_rawDescGZIP(), []int{6}
}

func (x *UpdateNodeRequest) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

func (x *UpdateNodeRequest) GetNodes() []*model.Node {
	if x != nil {
		return x.Nodes
	}
	return nil
}

type UpdateNodeResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Nodes []*model.Node `protobuf:"bytes,1,rep,name=Nodes,proto3" json:"Nodes,omitempty"`
}

func (x *UpdateNodeResponse) Reset() {
	*x = UpdateNodeResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_stub_node_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateNodeResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateNodeResponse) ProtoMessage() {}

func (x *UpdateNodeResponse) ProtoReflect() protoreflect.Message {
	mi := &file_stub_node_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateNodeResponse.ProtoReflect.Descriptor instead.
func (*UpdateNodeResponse) Descriptor() ([]byte, []int) {
	return file_stub_node_proto_rawDescGZIP(), []int{7}
}

func (x *UpdateNodeResponse) GetNodes() []*model.Node {
	if x != nil {
		return x.Nodes
	}
	return nil
}

type DeleteNodeRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Token string   `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	Ids   []string `protobuf:"bytes,2,rep,name=ids,proto3" json:"ids,omitempty"`
}

func (x *DeleteNodeRequest) Reset() {
	*x = DeleteNodeRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_stub_node_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteNodeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteNodeRequest) ProtoMessage() {}

func (x *DeleteNodeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_stub_node_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteNodeRequest.ProtoReflect.Descriptor instead.
func (*DeleteNodeRequest) Descriptor() ([]byte, []int) {
	return file_stub_node_proto_rawDescGZIP(), []int{8}
}

func (x *DeleteNodeRequest) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

func (x *DeleteNodeRequest) GetIds() []string {
	if x != nil {
		return x.Ids
	}
	return nil
}

type DeleteNodeResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *DeleteNodeResponse) Reset() {
	*x = DeleteNodeResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_stub_node_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteNodeResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteNodeResponse) ProtoMessage() {}

func (x *DeleteNodeResponse) ProtoReflect() protoreflect.Message {
	mi := &file_stub_node_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteNodeResponse.ProtoReflect.Descriptor instead.
func (*DeleteNodeResponse) Descriptor() ([]byte, []int) {
	return file_stub_node_proto_rawDescGZIP(), []int{9}
}

type GetNodeRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Token string `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	Id    string `protobuf:"bytes,2,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *GetNodeRequest) Reset() {
	*x = GetNodeRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_stub_node_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetNodeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetNodeRequest) ProtoMessage() {}

func (x *GetNodeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_stub_node_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetNodeRequest.ProtoReflect.Descriptor instead.
func (*GetNodeRequest) Descriptor() ([]byte, []int) {
	return file_stub_node_proto_rawDescGZIP(), []int{10}
}

func (x *GetNodeRequest) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

func (x *GetNodeRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type GetNodeResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Node *model.Node `protobuf:"bytes,1,opt,name=Node,proto3" json:"Node,omitempty"`
}

func (x *GetNodeResponse) Reset() {
	*x = GetNodeResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_stub_node_proto_msgTypes[11]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetNodeResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetNodeResponse) ProtoMessage() {}

func (x *GetNodeResponse) ProtoReflect() protoreflect.Message {
	mi := &file_stub_node_proto_msgTypes[11]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetNodeResponse.ProtoReflect.Descriptor instead.
func (*GetNodeResponse) Descriptor() ([]byte, []int) {
	return file_stub_node_proto_rawDescGZIP(), []int{11}
}

func (x *GetNodeResponse) GetNode() *model.Node {
	if x != nil {
		return x.Node
	}
	return nil
}

var File_stub_node_proto protoreflect.FileDescriptor

var file_stub_node_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x73, 0x74, 0x75, 0x62, 0x2f, 0x6e, 0x6f, 0x64, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x07, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x1a, 0x10, 0x6d, 0x6f, 0x64, 0x65,
	0x6c, 0x2f, 0x6e, 0x6f, 0x64, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x11, 0x6d, 0x6f,
	0x64, 0x65, 0x6c, 0x2f, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x14, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x39, 0x0a, 0x11, 0x4e, 0x6f, 0x64, 0x65, 0x53, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f,
	0x6b, 0x65, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e,
	0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64,
	0x22, 0x7c, 0x0a, 0x12, 0x4e, 0x6f, 0x64, 0x65, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3e, 0x0a, 0x1a, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x41, 0x6c, 0x72, 0x65, 0x61, 0x64, 0x79, 0x49, 0x6e, 0x69, 0x74, 0x69,
	0x61, 0x74, 0x65, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x1a, 0x63, 0x6f, 0x6e, 0x6e,
	0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x41, 0x6c, 0x72, 0x65, 0x61, 0x64, 0x79, 0x49, 0x6e, 0x69,
	0x74, 0x69, 0x61, 0x74, 0x65, 0x64, 0x12, 0x26, 0x0a, 0x0e, 0x74, 0x65, 0x73, 0x74, 0x43, 0x6f,
	0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0e,
	0x74, 0x65, 0x73, 0x74, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x27,
	0x0a, 0x0f, 0x4c, 0x69, 0x73, 0x74, 0x4e, 0x6f, 0x64, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x22, 0x39, 0x0a, 0x10, 0x4c, 0x69, 0x73, 0x74, 0x4e,
	0x6f, 0x64, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x25, 0x0a, 0x07, 0x63,
	0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x6d,
	0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x4e, 0x6f, 0x64, 0x65, 0x52, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65,
	0x6e, 0x74, 0x22, 0x4c, 0x0a, 0x11, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4e, 0x6f, 0x64, 0x65,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x21, 0x0a,
	0x05, 0x4e, 0x6f, 0x64, 0x65, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x6d,
	0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x4e, 0x6f, 0x64, 0x65, 0x52, 0x05, 0x4e, 0x6f, 0x64, 0x65, 0x73,
	0x22, 0x37, 0x0a, 0x12, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4e, 0x6f, 0x64, 0x65, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x21, 0x0a, 0x05, 0x4e, 0x6f, 0x64, 0x65, 0x73, 0x18,
	0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x4e, 0x6f,
	0x64, 0x65, 0x52, 0x05, 0x4e, 0x6f, 0x64, 0x65, 0x73, 0x22, 0x4c, 0x0a, 0x11, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x4e, 0x6f, 0x64, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14,
	0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74,
	0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x21, 0x0a, 0x05, 0x4e, 0x6f, 0x64, 0x65, 0x73, 0x18, 0x02, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x4e, 0x6f, 0x64, 0x65,
	0x52, 0x05, 0x4e, 0x6f, 0x64, 0x65, 0x73, 0x22, 0x37, 0x0a, 0x12, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x4e, 0x6f, 0x64, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x21, 0x0a,
	0x05, 0x4e, 0x6f, 0x64, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x6d,
	0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x4e, 0x6f, 0x64, 0x65, 0x52, 0x05, 0x4e, 0x6f, 0x64, 0x65, 0x73,
	0x22, 0x3b, 0x0a, 0x11, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x4e, 0x6f, 0x64, 0x65, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x10, 0x0a, 0x03, 0x69,
	0x64, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x09, 0x52, 0x03, 0x69, 0x64, 0x73, 0x22, 0x14, 0x0a,
	0x12, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x4e, 0x6f, 0x64, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x36, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x4e, 0x6f, 0x64, 0x65, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x0e, 0x0a, 0x02, 0x69,
	0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x32, 0x0a, 0x0f, 0x47,
	0x65, 0x74, 0x4e, 0x6f, 0x64, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1f,
	0x0a, 0x04, 0x4e, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x6d,
	0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x4e, 0x6f, 0x64, 0x65, 0x52, 0x04, 0x4e, 0x6f, 0x64, 0x65, 0x32,
	0x99, 0x03, 0x0a, 0x04, 0x4e, 0x6f, 0x64, 0x65, 0x12, 0x43, 0x0a, 0x06, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x12, 0x1a, 0x2e, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x2e, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x4e, 0x6f, 0x64, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1b,
	0x2e, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4e,
	0x6f, 0x64, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x3d, 0x0a,
	0x04, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x18, 0x2e, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x2e,
	0x4c, 0x69, 0x73, 0x74, 0x4e, 0x6f, 0x64, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x19, 0x2e, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x4e, 0x6f,
	0x64, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x43, 0x0a, 0x06,
	0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x12, 0x1a, 0x2e, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72,
	0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4e, 0x6f, 0x64, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x1b, 0x2e, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x2e, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x4e, 0x6f, 0x64, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x00, 0x12, 0x43, 0x0a, 0x06, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x12, 0x1a, 0x2e, 0x63, 0x6c,
	0x75, 0x73, 0x74, 0x65, 0x72, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x4e, 0x6f, 0x64, 0x65,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1b, 0x2e, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65,
	0x72, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x4e, 0x6f, 0x64, 0x65, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x3a, 0x0a, 0x03, 0x47, 0x65, 0x74, 0x12, 0x17, 0x2e,
	0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x2e, 0x47, 0x65, 0x74, 0x4e, 0x6f, 0x64, 0x65, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x18, 0x2e, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72,
	0x2e, 0x47, 0x65, 0x74, 0x4e, 0x6f, 0x64, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x00, 0x12, 0x47, 0x0a, 0x0a, 0x4e, 0x6f, 0x64, 0x65, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x12, 0x1a, 0x2e, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x2e, 0x4e, 0x6f, 0x64, 0x65, 0x53,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1b, 0x2e, 0x63,
	0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x2e, 0x4e, 0x6f, 0x64, 0x65, 0x53, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x24, 0x5a, 0x22, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x74, 0x69, 0x73, 0x6c, 0x69, 0x62,
	0x2f, 0x61, 0x70, 0x69, 0x62, 0x72, 0x65, 0x77, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x73, 0x74, 0x75,
	0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_stub_node_proto_rawDescOnce sync.Once
	file_stub_node_proto_rawDescData = file_stub_node_proto_rawDesc
)

func file_stub_node_proto_rawDescGZIP() []byte {
	file_stub_node_proto_rawDescOnce.Do(func() {
		file_stub_node_proto_rawDescData = protoimpl.X.CompressGZIP(file_stub_node_proto_rawDescData)
	})
	return file_stub_node_proto_rawDescData
}

var file_stub_node_proto_msgTypes = make([]protoimpl.MessageInfo, 12)
var file_stub_node_proto_goTypes = []interface{}{
	(*NodeStatusRequest)(nil),  // 0: cluster.NodeStatusRequest
	(*NodeStatusResponse)(nil), // 1: cluster.NodeStatusResponse
	(*ListNodeRequest)(nil),    // 2: cluster.ListNodeRequest
	(*ListNodeResponse)(nil),   // 3: cluster.ListNodeResponse
	(*CreateNodeRequest)(nil),  // 4: cluster.CreateNodeRequest
	(*CreateNodeResponse)(nil), // 5: cluster.CreateNodeResponse
	(*UpdateNodeRequest)(nil),  // 6: cluster.UpdateNodeRequest
	(*UpdateNodeResponse)(nil), // 7: cluster.UpdateNodeResponse
	(*DeleteNodeRequest)(nil),  // 8: cluster.DeleteNodeRequest
	(*DeleteNodeResponse)(nil), // 9: cluster.DeleteNodeResponse
	(*GetNodeRequest)(nil),     // 10: cluster.GetNodeRequest
	(*GetNodeResponse)(nil),    // 11: cluster.GetNodeResponse
	(*model.Node)(nil),         // 12: model.Node
}
var file_stub_node_proto_depIdxs = []int32{
	12, // 0: cluster.ListNodeResponse.content:type_name -> model.Node
	12, // 1: cluster.CreateNodeRequest.Nodes:type_name -> model.Node
	12, // 2: cluster.CreateNodeResponse.Nodes:type_name -> model.Node
	12, // 3: cluster.UpdateNodeRequest.Nodes:type_name -> model.Node
	12, // 4: cluster.UpdateNodeResponse.Nodes:type_name -> model.Node
	12, // 5: cluster.GetNodeResponse.Node:type_name -> model.Node
	4,  // 6: cluster.Node.Create:input_type -> cluster.CreateNodeRequest
	2,  // 7: cluster.Node.List:input_type -> cluster.ListNodeRequest
	6,  // 8: cluster.Node.Update:input_type -> cluster.UpdateNodeRequest
	8,  // 9: cluster.Node.Delete:input_type -> cluster.DeleteNodeRequest
	10, // 10: cluster.Node.Get:input_type -> cluster.GetNodeRequest
	0,  // 11: cluster.Node.NodeStatus:input_type -> cluster.NodeStatusRequest
	5,  // 12: cluster.Node.Create:output_type -> cluster.CreateNodeResponse
	3,  // 13: cluster.Node.List:output_type -> cluster.ListNodeResponse
	7,  // 14: cluster.Node.Update:output_type -> cluster.UpdateNodeResponse
	9,  // 15: cluster.Node.Delete:output_type -> cluster.DeleteNodeResponse
	11, // 16: cluster.Node.Get:output_type -> cluster.GetNodeResponse
	1,  // 17: cluster.Node.NodeStatus:output_type -> cluster.NodeStatusResponse
	12, // [12:18] is the sub-list for method output_type
	6,  // [6:12] is the sub-list for method input_type
	6,  // [6:6] is the sub-list for extension type_name
	6,  // [6:6] is the sub-list for extension extendee
	0,  // [0:6] is the sub-list for field type_name
}

func init() { file_stub_node_proto_init() }
func file_stub_node_proto_init() {
	if File_stub_node_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_stub_node_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NodeStatusRequest); i {
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
		file_stub_node_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NodeStatusResponse); i {
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
		file_stub_node_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListNodeRequest); i {
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
		file_stub_node_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListNodeResponse); i {
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
		file_stub_node_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateNodeRequest); i {
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
		file_stub_node_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateNodeResponse); i {
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
		file_stub_node_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateNodeRequest); i {
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
		file_stub_node_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateNodeResponse); i {
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
		file_stub_node_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteNodeRequest); i {
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
		file_stub_node_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteNodeResponse); i {
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
		file_stub_node_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetNodeRequest); i {
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
		file_stub_node_proto_msgTypes[11].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetNodeResponse); i {
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
			RawDescriptor: file_stub_node_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   12,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_stub_node_proto_goTypes,
		DependencyIndexes: file_stub_node_proto_depIdxs,
		MessageInfos:      file_stub_node_proto_msgTypes,
	}.Build()
	File_stub_node_proto = out.File
	file_stub_node_proto_rawDesc = nil
	file_stub_node_proto_goTypes = nil
	file_stub_node_proto_depIdxs = nil
}
