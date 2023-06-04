// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        (unknown)
// source: model/security.proto

package model

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type OperationType int32

const (
	OperationType_OPERATION_TYPE_READ   OperationType = 0
	OperationType_OPERATION_TYPE_CREATE OperationType = 1
	OperationType_OPERATION_TYPE_UPDATE OperationType = 2
	OperationType_OPERATION_TYPE_DELETE OperationType = 3
	OperationType_FULL                  OperationType = 4
)

// Enum value maps for OperationType.
var (
	OperationType_name = map[int32]string{
		0: "OPERATION_TYPE_READ",
		1: "OPERATION_TYPE_CREATE",
		2: "OPERATION_TYPE_UPDATE",
		3: "OPERATION_TYPE_DELETE",
		4: "FULL",
	}
	OperationType_value = map[string]int32{
		"OPERATION_TYPE_READ":   0,
		"OPERATION_TYPE_CREATE": 1,
		"OPERATION_TYPE_UPDATE": 2,
		"OPERATION_TYPE_DELETE": 3,
		"FULL":                  4,
	}
)

func (x OperationType) Enum() *OperationType {
	p := new(OperationType)
	*p = x
	return p
}

func (x OperationType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (OperationType) Descriptor() protoreflect.EnumDescriptor {
	return file_model_security_proto_enumTypes[0].Descriptor()
}

func (OperationType) Type() protoreflect.EnumType {
	return &file_model_security_proto_enumTypes[0]
}

func (x OperationType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use OperationType.Descriptor instead.
func (OperationType) EnumDescriptor() ([]byte, []int) {
	return file_model_security_proto_rawDescGZIP(), []int{0}
}

type PermitType int32

const (
	PermitType_PERMIT_TYPE_ALLOW   PermitType = 0
	PermitType_PERMIT_TYPE_REJECT  PermitType = 1
	PermitType_PERMIT_TYPE_UNKNOWN PermitType = 2
)

// Enum value maps for PermitType.
var (
	PermitType_name = map[int32]string{
		0: "PERMIT_TYPE_ALLOW",
		1: "PERMIT_TYPE_REJECT",
		2: "PERMIT_TYPE_UNKNOWN",
	}
	PermitType_value = map[string]int32{
		"PERMIT_TYPE_ALLOW":   0,
		"PERMIT_TYPE_REJECT":  1,
		"PERMIT_TYPE_UNKNOWN": 2,
	}
)

func (x PermitType) Enum() *PermitType {
	p := new(PermitType)
	*p = x
	return p
}

func (x PermitType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (PermitType) Descriptor() protoreflect.EnumDescriptor {
	return file_model_security_proto_enumTypes[1].Descriptor()
}

func (PermitType) Type() protoreflect.EnumType {
	return &file_model_security_proto_enumTypes[1]
}

func (x PermitType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use PermitType.Descriptor instead.
func (PermitType) EnumDescriptor() ([]byte, []int) {
	return file_model_security_proto_rawDescGZIP(), []int{1}
}

// SecurityConstraint is a rule
type SecurityConstraint struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Namespace string                 `protobuf:"bytes,1,opt,name=namespace,proto3" json:"namespace,omitempty"` //  namespace name where it will be applied
	Resource  string                 `protobuf:"bytes,2,opt,name=resource,proto3" json:"resource,omitempty"`   // resource name where it will be applied
	Property  string                 `protobuf:"bytes,3,opt,name=property,proto3" json:"property,omitempty"`   // property name where it will be applied
	Before    *timestamppb.Timestamp `protobuf:"bytes,5,opt,name=before,proto3" json:"before,omitempty"`       // before it is valid
	After     *timestamppb.Timestamp `protobuf:"bytes,6,opt,name=after,proto3" json:"after,omitempty"`         // after it is valid
	Principal string                 `protobuf:"bytes,7,opt,name=principal,proto3" json:"principal,omitempty"` // usernames which it is applied to
	// repeated string role = 8; // roles which it is applied to
	RecordIds []string      `protobuf:"bytes,9,rep,name=recordIds,proto3" json:"recordIds,omitempty"`                            // list of record ids which it is applied to
	Operation OperationType `protobuf:"varint,13,opt,name=operation,proto3,enum=model.OperationType" json:"operation,omitempty"` // operation name which it is applied to
	Permit    PermitType    `protobuf:"varint,14,opt,name=permit,proto3,enum=model.PermitType" json:"permit,omitempty"`          // permission
}

func (x *SecurityConstraint) Reset() {
	*x = SecurityConstraint{}
	if protoimpl.UnsafeEnabled {
		mi := &file_model_security_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SecurityConstraint) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SecurityConstraint) ProtoMessage() {}

func (x *SecurityConstraint) ProtoReflect() protoreflect.Message {
	mi := &file_model_security_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SecurityConstraint.ProtoReflect.Descriptor instead.
func (*SecurityConstraint) Descriptor() ([]byte, []int) {
	return file_model_security_proto_rawDescGZIP(), []int{0}
}

func (x *SecurityConstraint) GetNamespace() string {
	if x != nil {
		return x.Namespace
	}
	return ""
}

func (x *SecurityConstraint) GetResource() string {
	if x != nil {
		return x.Resource
	}
	return ""
}

func (x *SecurityConstraint) GetProperty() string {
	if x != nil {
		return x.Property
	}
	return ""
}

func (x *SecurityConstraint) GetBefore() *timestamppb.Timestamp {
	if x != nil {
		return x.Before
	}
	return nil
}

func (x *SecurityConstraint) GetAfter() *timestamppb.Timestamp {
	if x != nil {
		return x.After
	}
	return nil
}

func (x *SecurityConstraint) GetPrincipal() string {
	if x != nil {
		return x.Principal
	}
	return ""
}

func (x *SecurityConstraint) GetRecordIds() []string {
	if x != nil {
		return x.RecordIds
	}
	return nil
}

func (x *SecurityConstraint) GetOperation() OperationType {
	if x != nil {
		return x.Operation
	}
	return OperationType_OPERATION_TYPE_READ
}

func (x *SecurityConstraint) GetPermit() PermitType {
	if x != nil {
		return x.Permit
	}
	return PermitType_PERMIT_TYPE_ALLOW
}

type SecurityContext struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Constraints []*SecurityConstraint `protobuf:"bytes,1,rep,name=constraints,proto3" json:"constraints,omitempty"`
}

func (x *SecurityContext) Reset() {
	*x = SecurityContext{}
	if protoimpl.UnsafeEnabled {
		mi := &file_model_security_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SecurityContext) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SecurityContext) ProtoMessage() {}

func (x *SecurityContext) ProtoReflect() protoreflect.Message {
	mi := &file_model_security_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SecurityContext.ProtoReflect.Descriptor instead.
func (*SecurityContext) Descriptor() ([]byte, []int) {
	return file_model_security_proto_rawDescGZIP(), []int{1}
}

func (x *SecurityContext) GetConstraints() []*SecurityConstraint {
	if x != nil {
		return x.Constraints
	}
	return nil
}

var File_model_security_proto protoreflect.FileDescriptor

var file_model_security_proto_rawDesc = []byte{
	0x0a, 0x14, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2f, 0x73, 0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x1a, 0x1f, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0f,
	0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2f, 0x68, 0x63, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0xeb, 0x02, 0x0a, 0x12, 0x53, 0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x43, 0x6f, 0x6e, 0x73,
	0x74, 0x72, 0x61, 0x69, 0x6e, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70,
	0x61, 0x63, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x6e, 0x61, 0x6d, 0x65, 0x73,
	0x70, 0x61, 0x63, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x12, 0x1a, 0x0a, 0x08, 0x70, 0x72, 0x6f, 0x70, 0x65, 0x72, 0x74, 0x79, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x08, 0x70, 0x72, 0x6f, 0x70, 0x65, 0x72, 0x74, 0x79, 0x12, 0x32, 0x0a, 0x06,
	0x62, 0x65, 0x66, 0x6f, 0x72, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x06, 0x62, 0x65, 0x66, 0x6f, 0x72, 0x65,
	0x12, 0x30, 0x0a, 0x05, 0x61, 0x66, 0x74, 0x65, 0x72, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x05, 0x61, 0x66, 0x74,
	0x65, 0x72, 0x12, 0x1c, 0x0a, 0x09, 0x70, 0x72, 0x69, 0x6e, 0x63, 0x69, 0x70, 0x61, 0x6c, 0x18,
	0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x70, 0x72, 0x69, 0x6e, 0x63, 0x69, 0x70, 0x61, 0x6c,
	0x12, 0x1c, 0x0a, 0x09, 0x72, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x49, 0x64, 0x73, 0x18, 0x09, 0x20,
	0x03, 0x28, 0x09, 0x52, 0x09, 0x72, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x49, 0x64, 0x73, 0x12, 0x32,
	0x0a, 0x09, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x0d, 0x20, 0x01, 0x28,
	0x0e, 0x32, 0x14, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x52, 0x09, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x12, 0x29, 0x0a, 0x06, 0x70, 0x65, 0x72, 0x6d, 0x69, 0x74, 0x18, 0x0e, 0x20, 0x01,
	0x28, 0x0e, 0x32, 0x11, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x50, 0x65, 0x72, 0x6d, 0x69,
	0x74, 0x54, 0x79, 0x70, 0x65, 0x52, 0x06, 0x70, 0x65, 0x72, 0x6d, 0x69, 0x74, 0x22, 0x5d, 0x0a,
	0x0f, 0x53, 0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74,
	0x12, 0x4a, 0x0a, 0x0b, 0x63, 0x6f, 0x6e, 0x73, 0x74, 0x72, 0x61, 0x69, 0x6e, 0x74, 0x73, 0x18,
	0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x53, 0x65,
	0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x43, 0x6f, 0x6e, 0x73, 0x74, 0x72, 0x61, 0x69, 0x6e, 0x74,
	0x42, 0x0d, 0xc2, 0x47, 0x0a, 0x63, 0x6f, 0x6e, 0x73, 0x74, 0x72, 0x61, 0x69, 0x6e, 0x74, 0x52,
	0x0b, 0x63, 0x6f, 0x6e, 0x73, 0x74, 0x72, 0x61, 0x69, 0x6e, 0x74, 0x73, 0x2a, 0x83, 0x01, 0x0a,
	0x0d, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x12, 0x17,
	0x0a, 0x13, 0x4f, 0x50, 0x45, 0x52, 0x41, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x54, 0x59, 0x50, 0x45,
	0x5f, 0x52, 0x45, 0x41, 0x44, 0x10, 0x00, 0x12, 0x19, 0x0a, 0x15, 0x4f, 0x50, 0x45, 0x52, 0x41,
	0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x43, 0x52, 0x45, 0x41, 0x54, 0x45,
	0x10, 0x01, 0x12, 0x19, 0x0a, 0x15, 0x4f, 0x50, 0x45, 0x52, 0x41, 0x54, 0x49, 0x4f, 0x4e, 0x5f,
	0x54, 0x59, 0x50, 0x45, 0x5f, 0x55, 0x50, 0x44, 0x41, 0x54, 0x45, 0x10, 0x02, 0x12, 0x19, 0x0a,
	0x15, 0x4f, 0x50, 0x45, 0x52, 0x41, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f,
	0x44, 0x45, 0x4c, 0x45, 0x54, 0x45, 0x10, 0x03, 0x12, 0x08, 0x0a, 0x04, 0x46, 0x55, 0x4c, 0x4c,
	0x10, 0x04, 0x2a, 0x54, 0x0a, 0x0a, 0x50, 0x65, 0x72, 0x6d, 0x69, 0x74, 0x54, 0x79, 0x70, 0x65,
	0x12, 0x15, 0x0a, 0x11, 0x50, 0x45, 0x52, 0x4d, 0x49, 0x54, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f,
	0x41, 0x4c, 0x4c, 0x4f, 0x57, 0x10, 0x00, 0x12, 0x16, 0x0a, 0x12, 0x50, 0x45, 0x52, 0x4d, 0x49,
	0x54, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x52, 0x45, 0x4a, 0x45, 0x43, 0x54, 0x10, 0x01, 0x12,
	0x17, 0x0a, 0x13, 0x50, 0x45, 0x52, 0x4d, 0x49, 0x54, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x55,
	0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x10, 0x02, 0x42, 0x26, 0x5a, 0x24, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x61, 0x70, 0x69, 0x62, 0x72, 0x65, 0x77, 0x2f, 0x61,
	0x70, 0x69, 0x62, 0x72, 0x65, 0x77, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_model_security_proto_rawDescOnce sync.Once
	file_model_security_proto_rawDescData = file_model_security_proto_rawDesc
)

func file_model_security_proto_rawDescGZIP() []byte {
	file_model_security_proto_rawDescOnce.Do(func() {
		file_model_security_proto_rawDescData = protoimpl.X.CompressGZIP(file_model_security_proto_rawDescData)
	})
	return file_model_security_proto_rawDescData
}

var file_model_security_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_model_security_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_model_security_proto_goTypes = []interface{}{
	(OperationType)(0),            // 0: model.OperationType
	(PermitType)(0),               // 1: model.PermitType
	(*SecurityConstraint)(nil),    // 2: model.SecurityConstraint
	(*SecurityContext)(nil),       // 3: model.SecurityContext
	(*timestamppb.Timestamp)(nil), // 4: google.protobuf.Timestamp
}
var file_model_security_proto_depIdxs = []int32{
	4, // 0: model.SecurityConstraint.before:type_name -> google.protobuf.Timestamp
	4, // 1: model.SecurityConstraint.after:type_name -> google.protobuf.Timestamp
	0, // 2: model.SecurityConstraint.operation:type_name -> model.OperationType
	1, // 3: model.SecurityConstraint.permit:type_name -> model.PermitType
	2, // 4: model.SecurityContext.constraints:type_name -> model.SecurityConstraint
	5, // [5:5] is the sub-list for method output_type
	5, // [5:5] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_model_security_proto_init() }
func file_model_security_proto_init() {
	if File_model_security_proto != nil {
		return
	}
	file_model_hcl_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_model_security_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SecurityConstraint); i {
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
		file_model_security_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SecurityContext); i {
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
			RawDescriptor: file_model_security_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_model_security_proto_goTypes,
		DependencyIndexes: file_model_security_proto_depIdxs,
		EnumInfos:         file_model_security_proto_enumTypes,
		MessageInfos:      file_model_security_proto_msgTypes,
	}.Build()
	File_model_security_proto = out.File
	file_model_security_proto_rawDesc = nil
	file_model_security_proto_goTypes = nil
	file_model_security_proto_depIdxs = nil
}
