// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        (unknown)
// source: model/query.proto

package model

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	structpb "google.golang.org/protobuf/types/known/structpb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type CompoundBooleanExpression struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Expressions []*BooleanExpression `protobuf:"bytes,1,rep,name=expressions,proto3" json:"expressions,omitempty"`
}

func (x *CompoundBooleanExpression) Reset() {
	*x = CompoundBooleanExpression{}
	if protoimpl.UnsafeEnabled {
		mi := &file_model_query_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CompoundBooleanExpression) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CompoundBooleanExpression) ProtoMessage() {}

func (x *CompoundBooleanExpression) ProtoReflect() protoreflect.Message {
	mi := &file_model_query_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CompoundBooleanExpression.ProtoReflect.Descriptor instead.
func (*CompoundBooleanExpression) Descriptor() ([]byte, []int) {
	return file_model_query_proto_rawDescGZIP(), []int{0}
}

func (x *CompoundBooleanExpression) GetExpressions() []*BooleanExpression {
	if x != nil {
		return x.Expressions
	}
	return nil
}

type Expression struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Expression:
	//
	//	*Expression_Property
	//	*Expression_Value
	Expression isExpression_Expression `protobuf_oneof:"expression"`
}

func (x *Expression) Reset() {
	*x = Expression{}
	if protoimpl.UnsafeEnabled {
		mi := &file_model_query_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Expression) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Expression) ProtoMessage() {}

func (x *Expression) ProtoReflect() protoreflect.Message {
	mi := &file_model_query_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Expression.ProtoReflect.Descriptor instead.
func (*Expression) Descriptor() ([]byte, []int) {
	return file_model_query_proto_rawDescGZIP(), []int{1}
}

func (m *Expression) GetExpression() isExpression_Expression {
	if m != nil {
		return m.Expression
	}
	return nil
}

func (x *Expression) GetProperty() string {
	if x, ok := x.GetExpression().(*Expression_Property); ok {
		return x.Property
	}
	return ""
}

func (x *Expression) GetValue() *structpb.Value {
	if x, ok := x.GetExpression().(*Expression_Value); ok {
		return x.Value
	}
	return nil
}

type isExpression_Expression interface {
	isExpression_Expression()
}

type Expression_Property struct {
	Property string `protobuf:"bytes,1,opt,name=property,proto3,oneof"`
}

type Expression_Value struct {
	Value *structpb.Value `protobuf:"bytes,3,opt,name=value,proto3,oneof"`
}

func (*Expression_Property) isExpression_Expression() {}

func (*Expression_Value) isExpression_Expression() {}

type PairExpression struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Left  *Expression `protobuf:"bytes,1,opt,name=left,proto3" json:"left,omitempty"`
	Right *Expression `protobuf:"bytes,2,opt,name=right,proto3" json:"right,omitempty"`
}

func (x *PairExpression) Reset() {
	*x = PairExpression{}
	if protoimpl.UnsafeEnabled {
		mi := &file_model_query_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PairExpression) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PairExpression) ProtoMessage() {}

func (x *PairExpression) ProtoReflect() protoreflect.Message {
	mi := &file_model_query_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PairExpression.ProtoReflect.Descriptor instead.
func (*PairExpression) Descriptor() ([]byte, []int) {
	return file_model_query_proto_rawDescGZIP(), []int{2}
}

func (x *PairExpression) GetLeft() *Expression {
	if x != nil {
		return x.Left
	}
	return nil
}

func (x *PairExpression) GetRight() *Expression {
	if x != nil {
		return x.Right
	}
	return nil
}

type Filters struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Filters map[string]*structpb.Value `protobuf:"bytes,1,rep,name=filters,proto3" json:"filters,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *Filters) Reset() {
	*x = Filters{}
	if protoimpl.UnsafeEnabled {
		mi := &file_model_query_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Filters) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Filters) ProtoMessage() {}

func (x *Filters) ProtoReflect() protoreflect.Message {
	mi := &file_model_query_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Filters.ProtoReflect.Descriptor instead.
func (*Filters) Descriptor() ([]byte, []int) {
	return file_model_query_proto_rawDescGZIP(), []int{3}
}

func (x *Filters) GetFilters() map[string]*structpb.Value {
	if x != nil {
		return x.Filters
	}
	return nil
}

type RegexMatchExpression struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Pattern    string      `protobuf:"bytes,1,opt,name=pattern,proto3" json:"pattern,omitempty"`
	Expression *Expression `protobuf:"bytes,2,opt,name=expression,proto3" json:"expression,omitempty"`
}

func (x *RegexMatchExpression) Reset() {
	*x = RegexMatchExpression{}
	if protoimpl.UnsafeEnabled {
		mi := &file_model_query_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RegexMatchExpression) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RegexMatchExpression) ProtoMessage() {}

func (x *RegexMatchExpression) ProtoReflect() protoreflect.Message {
	mi := &file_model_query_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RegexMatchExpression.ProtoReflect.Descriptor instead.
func (*RegexMatchExpression) Descriptor() ([]byte, []int) {
	return file_model_query_proto_rawDescGZIP(), []int{4}
}

func (x *RegexMatchExpression) GetPattern() string {
	if x != nil {
		return x.Pattern
	}
	return ""
}

func (x *RegexMatchExpression) GetExpression() *Expression {
	if x != nil {
		return x.Expression
	}
	return nil
}

type BooleanExpression struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Expression:
	//
	//	*BooleanExpression_And
	//	*BooleanExpression_Or
	//	*BooleanExpression_Not
	//	*BooleanExpression_Equal
	//	*BooleanExpression_LessThan
	//	*BooleanExpression_GreaterThan
	//	*BooleanExpression_LessThanOrEqual
	//	*BooleanExpression_GreaterThanOrEqual
	//	*BooleanExpression_In
	//	*BooleanExpression_IsNull
	//	*BooleanExpression_RegexMatch
	Expression isBooleanExpression_Expression `protobuf_oneof:"expression"`
	Filters    map[string]*structpb.Value     `protobuf:"bytes,12,rep,name=filters,proto3" json:"filters,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *BooleanExpression) Reset() {
	*x = BooleanExpression{}
	if protoimpl.UnsafeEnabled {
		mi := &file_model_query_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BooleanExpression) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BooleanExpression) ProtoMessage() {}

func (x *BooleanExpression) ProtoReflect() protoreflect.Message {
	mi := &file_model_query_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BooleanExpression.ProtoReflect.Descriptor instead.
func (*BooleanExpression) Descriptor() ([]byte, []int) {
	return file_model_query_proto_rawDescGZIP(), []int{5}
}

func (m *BooleanExpression) GetExpression() isBooleanExpression_Expression {
	if m != nil {
		return m.Expression
	}
	return nil
}

func (x *BooleanExpression) GetAnd() *CompoundBooleanExpression {
	if x, ok := x.GetExpression().(*BooleanExpression_And); ok {
		return x.And
	}
	return nil
}

func (x *BooleanExpression) GetOr() *CompoundBooleanExpression {
	if x, ok := x.GetExpression().(*BooleanExpression_Or); ok {
		return x.Or
	}
	return nil
}

func (x *BooleanExpression) GetNot() *BooleanExpression {
	if x, ok := x.GetExpression().(*BooleanExpression_Not); ok {
		return x.Not
	}
	return nil
}

func (x *BooleanExpression) GetEqual() *PairExpression {
	if x, ok := x.GetExpression().(*BooleanExpression_Equal); ok {
		return x.Equal
	}
	return nil
}

func (x *BooleanExpression) GetLessThan() *PairExpression {
	if x, ok := x.GetExpression().(*BooleanExpression_LessThan); ok {
		return x.LessThan
	}
	return nil
}

func (x *BooleanExpression) GetGreaterThan() *PairExpression {
	if x, ok := x.GetExpression().(*BooleanExpression_GreaterThan); ok {
		return x.GreaterThan
	}
	return nil
}

func (x *BooleanExpression) GetLessThanOrEqual() *PairExpression {
	if x, ok := x.GetExpression().(*BooleanExpression_LessThanOrEqual); ok {
		return x.LessThanOrEqual
	}
	return nil
}

func (x *BooleanExpression) GetGreaterThanOrEqual() *PairExpression {
	if x, ok := x.GetExpression().(*BooleanExpression_GreaterThanOrEqual); ok {
		return x.GreaterThanOrEqual
	}
	return nil
}

func (x *BooleanExpression) GetIn() *PairExpression {
	if x, ok := x.GetExpression().(*BooleanExpression_In); ok {
		return x.In
	}
	return nil
}

func (x *BooleanExpression) GetIsNull() *Expression {
	if x, ok := x.GetExpression().(*BooleanExpression_IsNull); ok {
		return x.IsNull
	}
	return nil
}

func (x *BooleanExpression) GetRegexMatch() *RegexMatchExpression {
	if x, ok := x.GetExpression().(*BooleanExpression_RegexMatch); ok {
		return x.RegexMatch
	}
	return nil
}

func (x *BooleanExpression) GetFilters() map[string]*structpb.Value {
	if x != nil {
		return x.Filters
	}
	return nil
}

type isBooleanExpression_Expression interface {
	isBooleanExpression_Expression()
}

type BooleanExpression_And struct {
	// logical expressions
	And *CompoundBooleanExpression `protobuf:"bytes,1,opt,name=and,proto3,oneof"`
}

type BooleanExpression_Or struct {
	Or *CompoundBooleanExpression `protobuf:"bytes,2,opt,name=or,proto3,oneof"`
}

type BooleanExpression_Not struct {
	Not *BooleanExpression `protobuf:"bytes,3,opt,name=not,proto3,oneof"`
}

type BooleanExpression_Equal struct {
	// basic comparison
	Equal *PairExpression `protobuf:"bytes,4,opt,name=equal,proto3,oneof"`
}

type BooleanExpression_LessThan struct {
	LessThan *PairExpression `protobuf:"bytes,5,opt,name=lessThan,proto3,oneof"`
}

type BooleanExpression_GreaterThan struct {
	GreaterThan *PairExpression `protobuf:"bytes,6,opt,name=greaterThan,proto3,oneof"`
}

type BooleanExpression_LessThanOrEqual struct {
	LessThanOrEqual *PairExpression `protobuf:"bytes,7,opt,name=lessThanOrEqual,proto3,oneof"`
}

type BooleanExpression_GreaterThanOrEqual struct {
	GreaterThanOrEqual *PairExpression `protobuf:"bytes,8,opt,name=greaterThanOrEqual,proto3,oneof"`
}

type BooleanExpression_In struct {
	In *PairExpression `protobuf:"bytes,9,opt,name=in,proto3,oneof"`
}

type BooleanExpression_IsNull struct {
	IsNull *Expression `protobuf:"bytes,10,opt,name=isNull,proto3,oneof"`
}

type BooleanExpression_RegexMatch struct {
	// other
	RegexMatch *RegexMatchExpression `protobuf:"bytes,11,opt,name=regexMatch,proto3,oneof"`
}

func (*BooleanExpression_And) isBooleanExpression_Expression() {}

func (*BooleanExpression_Or) isBooleanExpression_Expression() {}

func (*BooleanExpression_Not) isBooleanExpression_Expression() {}

func (*BooleanExpression_Equal) isBooleanExpression_Expression() {}

func (*BooleanExpression_LessThan) isBooleanExpression_Expression() {}

func (*BooleanExpression_GreaterThan) isBooleanExpression_Expression() {}

func (*BooleanExpression_LessThanOrEqual) isBooleanExpression_Expression() {}

func (*BooleanExpression_GreaterThanOrEqual) isBooleanExpression_Expression() {}

func (*BooleanExpression_In) isBooleanExpression_Expression() {}

func (*BooleanExpression_IsNull) isBooleanExpression_Expression() {}

func (*BooleanExpression_RegexMatch) isBooleanExpression_Expression() {}

var File_model_query_proto protoreflect.FileDescriptor

var file_model_query_proto_rawDesc = []byte{
	0x0a, 0x11, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2f, 0x71, 0x75, 0x65, 0x72, 0x79, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x05, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x73, 0x74, 0x72, 0x75,
	0x63, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x57, 0x0a, 0x19, 0x43, 0x6f, 0x6d, 0x70,
	0x6f, 0x75, 0x6e, 0x64, 0x42, 0x6f, 0x6f, 0x6c, 0x65, 0x61, 0x6e, 0x45, 0x78, 0x70, 0x72, 0x65,
	0x73, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x3a, 0x0a, 0x0b, 0x65, 0x78, 0x70, 0x72, 0x65, 0x73, 0x73,
	0x69, 0x6f, 0x6e, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x6d, 0x6f, 0x64,
	0x65, 0x6c, 0x2e, 0x42, 0x6f, 0x6f, 0x6c, 0x65, 0x61, 0x6e, 0x45, 0x78, 0x70, 0x72, 0x65, 0x73,
	0x73, 0x69, 0x6f, 0x6e, 0x52, 0x0b, 0x65, 0x78, 0x70, 0x72, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e,
	0x73, 0x22, 0x68, 0x0a, 0x0a, 0x45, 0x78, 0x70, 0x72, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x12,
	0x1c, 0x0a, 0x08, 0x70, 0x72, 0x6f, 0x70, 0x65, 0x72, 0x74, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x48, 0x00, 0x52, 0x08, 0x70, 0x72, 0x6f, 0x70, 0x65, 0x72, 0x74, 0x79, 0x12, 0x2e, 0x0a,
	0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x56,
	0x61, 0x6c, 0x75, 0x65, 0x48, 0x00, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x42, 0x0c, 0x0a,
	0x0a, 0x65, 0x78, 0x70, 0x72, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x22, 0x60, 0x0a, 0x0e, 0x50,
	0x61, 0x69, 0x72, 0x45, 0x78, 0x70, 0x72, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x25, 0x0a,
	0x04, 0x6c, 0x65, 0x66, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x6d, 0x6f,
	0x64, 0x65, 0x6c, 0x2e, 0x45, 0x78, 0x70, 0x72, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x04,
	0x6c, 0x65, 0x66, 0x74, 0x12, 0x27, 0x0a, 0x05, 0x72, 0x69, 0x67, 0x68, 0x74, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x45, 0x78, 0x70, 0x72,
	0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x05, 0x72, 0x69, 0x67, 0x68, 0x74, 0x22, 0x94, 0x01,
	0x0a, 0x07, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x73, 0x12, 0x35, 0x0a, 0x07, 0x66, 0x69, 0x6c,
	0x74, 0x65, 0x72, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x6d, 0x6f, 0x64,
	0x65, 0x6c, 0x2e, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x73, 0x2e, 0x46, 0x69, 0x6c, 0x74, 0x65,
	0x72, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x07, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x73,
	0x1a, 0x52, 0x0a, 0x0c, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79,
	0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b,
	0x65, 0x79, 0x12, 0x2c, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65,
	0x3a, 0x02, 0x38, 0x01, 0x22, 0x63, 0x0a, 0x14, 0x52, 0x65, 0x67, 0x65, 0x78, 0x4d, 0x61, 0x74,
	0x63, 0x68, 0x45, 0x78, 0x70, 0x72, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x18, 0x0a, 0x07,
	0x70, 0x61, 0x74, 0x74, 0x65, 0x72, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x70,
	0x61, 0x74, 0x74, 0x65, 0x72, 0x6e, 0x12, 0x31, 0x0a, 0x0a, 0x65, 0x78, 0x70, 0x72, 0x65, 0x73,
	0x73, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x6d, 0x6f, 0x64,
	0x65, 0x6c, 0x2e, 0x45, 0x78, 0x70, 0x72, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x0a, 0x65,
	0x78, 0x70, 0x72, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x22, 0x8e, 0x06, 0x0a, 0x11, 0x42, 0x6f,
	0x6f, 0x6c, 0x65, 0x61, 0x6e, 0x45, 0x78, 0x70, 0x72, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x12,
	0x34, 0x0a, 0x03, 0x61, 0x6e, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x20, 0x2e, 0x6d,
	0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x43, 0x6f, 0x6d, 0x70, 0x6f, 0x75, 0x6e, 0x64, 0x42, 0x6f, 0x6f,
	0x6c, 0x65, 0x61, 0x6e, 0x45, 0x78, 0x70, 0x72, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x48, 0x00,
	0x52, 0x03, 0x61, 0x6e, 0x64, 0x12, 0x32, 0x0a, 0x02, 0x6f, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x20, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x43, 0x6f, 0x6d, 0x70, 0x6f, 0x75,
	0x6e, 0x64, 0x42, 0x6f, 0x6f, 0x6c, 0x65, 0x61, 0x6e, 0x45, 0x78, 0x70, 0x72, 0x65, 0x73, 0x73,
	0x69, 0x6f, 0x6e, 0x48, 0x00, 0x52, 0x02, 0x6f, 0x72, 0x12, 0x2c, 0x0a, 0x03, 0x6e, 0x6f, 0x74,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x42,
	0x6f, 0x6f, 0x6c, 0x65, 0x61, 0x6e, 0x45, 0x78, 0x70, 0x72, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e,
	0x48, 0x00, 0x52, 0x03, 0x6e, 0x6f, 0x74, 0x12, 0x2d, 0x0a, 0x05, 0x65, 0x71, 0x75, 0x61, 0x6c,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x50,
	0x61, 0x69, 0x72, 0x45, 0x78, 0x70, 0x72, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x48, 0x00, 0x52,
	0x05, 0x65, 0x71, 0x75, 0x61, 0x6c, 0x12, 0x33, 0x0a, 0x08, 0x6c, 0x65, 0x73, 0x73, 0x54, 0x68,
	0x61, 0x6e, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c,
	0x2e, 0x50, 0x61, 0x69, 0x72, 0x45, 0x78, 0x70, 0x72, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x48,
	0x00, 0x52, 0x08, 0x6c, 0x65, 0x73, 0x73, 0x54, 0x68, 0x61, 0x6e, 0x12, 0x39, 0x0a, 0x0b, 0x67,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x72, 0x54, 0x68, 0x61, 0x6e, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x15, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x50, 0x61, 0x69, 0x72, 0x45, 0x78, 0x70,
	0x72, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x48, 0x00, 0x52, 0x0b, 0x67, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x72, 0x54, 0x68, 0x61, 0x6e, 0x12, 0x41, 0x0a, 0x0f, 0x6c, 0x65, 0x73, 0x73, 0x54, 0x68,
	0x61, 0x6e, 0x4f, 0x72, 0x45, 0x71, 0x75, 0x61, 0x6c, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x15, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x50, 0x61, 0x69, 0x72, 0x45, 0x78, 0x70, 0x72,
	0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x48, 0x00, 0x52, 0x0f, 0x6c, 0x65, 0x73, 0x73, 0x54, 0x68,
	0x61, 0x6e, 0x4f, 0x72, 0x45, 0x71, 0x75, 0x61, 0x6c, 0x12, 0x47, 0x0a, 0x12, 0x67, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x72, 0x54, 0x68, 0x61, 0x6e, 0x4f, 0x72, 0x45, 0x71, 0x75, 0x61, 0x6c, 0x18,
	0x08, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x50, 0x61,
	0x69, 0x72, 0x45, 0x78, 0x70, 0x72, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x48, 0x00, 0x52, 0x12,
	0x67, 0x72, 0x65, 0x61, 0x74, 0x65, 0x72, 0x54, 0x68, 0x61, 0x6e, 0x4f, 0x72, 0x45, 0x71, 0x75,
	0x61, 0x6c, 0x12, 0x27, 0x0a, 0x02, 0x69, 0x6e, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x15,
	0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x50, 0x61, 0x69, 0x72, 0x45, 0x78, 0x70, 0x72, 0x65,
	0x73, 0x73, 0x69, 0x6f, 0x6e, 0x48, 0x00, 0x52, 0x02, 0x69, 0x6e, 0x12, 0x2b, 0x0a, 0x06, 0x69,
	0x73, 0x4e, 0x75, 0x6c, 0x6c, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x6d, 0x6f,
	0x64, 0x65, 0x6c, 0x2e, 0x45, 0x78, 0x70, 0x72, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x48, 0x00,
	0x52, 0x06, 0x69, 0x73, 0x4e, 0x75, 0x6c, 0x6c, 0x12, 0x3d, 0x0a, 0x0a, 0x72, 0x65, 0x67, 0x65,
	0x78, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x6d,
	0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x52, 0x65, 0x67, 0x65, 0x78, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x45,
	0x78, 0x70, 0x72, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x48, 0x00, 0x52, 0x0a, 0x72, 0x65, 0x67,
	0x65, 0x78, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x12, 0x3f, 0x0a, 0x07, 0x66, 0x69, 0x6c, 0x74, 0x65,
	0x72, 0x73, 0x18, 0x0c, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x25, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c,
	0x2e, 0x42, 0x6f, 0x6f, 0x6c, 0x65, 0x61, 0x6e, 0x45, 0x78, 0x70, 0x72, 0x65, 0x73, 0x73, 0x69,
	0x6f, 0x6e, 0x2e, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52,
	0x07, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x73, 0x1a, 0x52, 0x0a, 0x0c, 0x46, 0x69, 0x6c, 0x74,
	0x65, 0x72, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x2c, 0x0a, 0x05, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x56, 0x61, 0x6c, 0x75,
	0x65, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x42, 0x0c, 0x0a, 0x0a,
	0x65, 0x78, 0x70, 0x72, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x42, 0x26, 0x5a, 0x24, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x61, 0x70, 0x69, 0x62, 0x72, 0x65, 0x77,
	0x2f, 0x61, 0x70, 0x69, 0x62, 0x72, 0x65, 0x77, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x6d, 0x6f, 0x64,
	0x65, 0x6c, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_model_query_proto_rawDescOnce sync.Once
	file_model_query_proto_rawDescData = file_model_query_proto_rawDesc
)

func file_model_query_proto_rawDescGZIP() []byte {
	file_model_query_proto_rawDescOnce.Do(func() {
		file_model_query_proto_rawDescData = protoimpl.X.CompressGZIP(file_model_query_proto_rawDescData)
	})
	return file_model_query_proto_rawDescData
}

var file_model_query_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_model_query_proto_goTypes = []interface{}{
	(*CompoundBooleanExpression)(nil), // 0: model.CompoundBooleanExpression
	(*Expression)(nil),                // 1: model.Expression
	(*PairExpression)(nil),            // 2: model.PairExpression
	(*Filters)(nil),                   // 3: model.Filters
	(*RegexMatchExpression)(nil),      // 4: model.RegexMatchExpression
	(*BooleanExpression)(nil),         // 5: model.BooleanExpression
	nil,                               // 6: model.Filters.FiltersEntry
	nil,                               // 7: model.BooleanExpression.FiltersEntry
	(*structpb.Value)(nil),            // 8: google.protobuf.Value
}
var file_model_query_proto_depIdxs = []int32{
	5,  // 0: model.CompoundBooleanExpression.expressions:type_name -> model.BooleanExpression
	8,  // 1: model.Expression.value:type_name -> google.protobuf.Value
	1,  // 2: model.PairExpression.left:type_name -> model.Expression
	1,  // 3: model.PairExpression.right:type_name -> model.Expression
	6,  // 4: model.Filters.filters:type_name -> model.Filters.FiltersEntry
	1,  // 5: model.RegexMatchExpression.expression:type_name -> model.Expression
	0,  // 6: model.BooleanExpression.and:type_name -> model.CompoundBooleanExpression
	0,  // 7: model.BooleanExpression.or:type_name -> model.CompoundBooleanExpression
	5,  // 8: model.BooleanExpression.not:type_name -> model.BooleanExpression
	2,  // 9: model.BooleanExpression.equal:type_name -> model.PairExpression
	2,  // 10: model.BooleanExpression.lessThan:type_name -> model.PairExpression
	2,  // 11: model.BooleanExpression.greaterThan:type_name -> model.PairExpression
	2,  // 12: model.BooleanExpression.lessThanOrEqual:type_name -> model.PairExpression
	2,  // 13: model.BooleanExpression.greaterThanOrEqual:type_name -> model.PairExpression
	2,  // 14: model.BooleanExpression.in:type_name -> model.PairExpression
	1,  // 15: model.BooleanExpression.isNull:type_name -> model.Expression
	4,  // 16: model.BooleanExpression.regexMatch:type_name -> model.RegexMatchExpression
	7,  // 17: model.BooleanExpression.filters:type_name -> model.BooleanExpression.FiltersEntry
	8,  // 18: model.Filters.FiltersEntry.value:type_name -> google.protobuf.Value
	8,  // 19: model.BooleanExpression.FiltersEntry.value:type_name -> google.protobuf.Value
	20, // [20:20] is the sub-list for method output_type
	20, // [20:20] is the sub-list for method input_type
	20, // [20:20] is the sub-list for extension type_name
	20, // [20:20] is the sub-list for extension extendee
	0,  // [0:20] is the sub-list for field type_name
}

func init() { file_model_query_proto_init() }
func file_model_query_proto_init() {
	if File_model_query_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_model_query_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CompoundBooleanExpression); i {
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
		file_model_query_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Expression); i {
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
		file_model_query_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PairExpression); i {
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
		file_model_query_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Filters); i {
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
		file_model_query_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RegexMatchExpression); i {
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
		file_model_query_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BooleanExpression); i {
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
	file_model_query_proto_msgTypes[1].OneofWrappers = []interface{}{
		(*Expression_Property)(nil),
		(*Expression_Value)(nil),
	}
	file_model_query_proto_msgTypes[5].OneofWrappers = []interface{}{
		(*BooleanExpression_And)(nil),
		(*BooleanExpression_Or)(nil),
		(*BooleanExpression_Not)(nil),
		(*BooleanExpression_Equal)(nil),
		(*BooleanExpression_LessThan)(nil),
		(*BooleanExpression_GreaterThan)(nil),
		(*BooleanExpression_LessThanOrEqual)(nil),
		(*BooleanExpression_GreaterThanOrEqual)(nil),
		(*BooleanExpression_In)(nil),
		(*BooleanExpression_IsNull)(nil),
		(*BooleanExpression_RegexMatch)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_model_query_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_model_query_proto_goTypes,
		DependencyIndexes: file_model_query_proto_depIdxs,
		MessageInfos:      file_model_query_proto_msgTypes,
	}.Build()
	File_model_query_proto = out.File
	file_model_query_proto_rawDesc = nil
	file_model_query_proto_goTypes = nil
	file_model_query_proto_depIdxs = nil
}
