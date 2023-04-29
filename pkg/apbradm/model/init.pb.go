// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        (unknown)
// source: model/init.proto

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

type AppConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Host                  string `protobuf:"bytes,1,opt,name=host,proto3" json:"host,omitempty"`
	Port                  int32  `protobuf:"varint,2,opt,name=port,proto3" json:"port,omitempty"`
	JwtPrivateKey         string `protobuf:"bytes,3,opt,name=jwtPrivateKey,proto3" json:"jwtPrivateKey,omitempty"`
	JwtPublicKey          string `protobuf:"bytes,4,opt,name=jwtPublicKey,proto3" json:"jwtPublicKey,omitempty"`
	DisableAuthentication bool   `protobuf:"varint,5,opt,name=disableAuthentication,proto3" json:"disableAuthentication,omitempty"`
	DisableCache          bool   `protobuf:"varint,6,opt,name=disableCache,proto3" json:"disableCache,omitempty"`
	PluginsPath           string `protobuf:"bytes,7,opt,name=pluginsPath,proto3" json:"pluginsPath,omitempty"`
}

func (x *AppConfig) Reset() {
	*x = AppConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_model_init_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AppConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AppConfig) ProtoMessage() {}

func (x *AppConfig) ProtoReflect() protoreflect.Message {
	mi := &file_model_init_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AppConfig.ProtoReflect.Descriptor instead.
func (*AppConfig) Descriptor() ([]byte, []int) {
	return file_model_init_proto_rawDescGZIP(), []int{0}
}

func (x *AppConfig) GetHost() string {
	if x != nil {
		return x.Host
	}
	return ""
}

func (x *AppConfig) GetPort() int32 {
	if x != nil {
		return x.Port
	}
	return 0
}

func (x *AppConfig) GetJwtPrivateKey() string {
	if x != nil {
		return x.JwtPrivateKey
	}
	return ""
}

func (x *AppConfig) GetJwtPublicKey() string {
	if x != nil {
		return x.JwtPublicKey
	}
	return ""
}

func (x *AppConfig) GetDisableAuthentication() bool {
	if x != nil {
		return x.DisableAuthentication
	}
	return false
}

func (x *AppConfig) GetDisableCache() bool {
	if x != nil {
		return x.DisableCache
	}
	return false
}

func (x *AppConfig) GetPluginsPath() string {
	if x != nil {
		return x.PluginsPath
	}
	return ""
}

type InitData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Config *AppConfig `protobuf:"bytes,1,opt,name=config,proto3" json:"config,omitempty"`
}

func (x *InitData) Reset() {
	*x = InitData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_model_init_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InitData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InitData) ProtoMessage() {}

func (x *InitData) ProtoReflect() protoreflect.Message {
	mi := &file_model_init_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InitData.ProtoReflect.Descriptor instead.
func (*InitData) Descriptor() ([]byte, []int) {
	return file_model_init_proto_rawDescGZIP(), []int{1}
}

func (x *InitData) GetConfig() *AppConfig {
	if x != nil {
		return x.Config
	}
	return nil
}

var File_model_init_proto protoreflect.FileDescriptor

var file_model_init_proto_rawDesc = []byte{
	0x0a, 0x10, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2f, 0x69, 0x6e, 0x69, 0x74, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x05, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x22, 0xf9, 0x01, 0x0a, 0x09, 0x41, 0x70,
	0x70, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x12, 0x0a, 0x04, 0x68, 0x6f, 0x73, 0x74, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x68, 0x6f, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x70,
	0x6f, 0x72, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x70, 0x6f, 0x72, 0x74, 0x12,
	0x24, 0x0a, 0x0d, 0x6a, 0x77, 0x74, 0x50, 0x72, 0x69, 0x76, 0x61, 0x74, 0x65, 0x4b, 0x65, 0x79,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x6a, 0x77, 0x74, 0x50, 0x72, 0x69, 0x76, 0x61,
	0x74, 0x65, 0x4b, 0x65, 0x79, 0x12, 0x22, 0x0a, 0x0c, 0x6a, 0x77, 0x74, 0x50, 0x75, 0x62, 0x6c,
	0x69, 0x63, 0x4b, 0x65, 0x79, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x6a, 0x77, 0x74,
	0x50, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x4b, 0x65, 0x79, 0x12, 0x34, 0x0a, 0x15, 0x64, 0x69, 0x73,
	0x61, 0x62, 0x6c, 0x65, 0x41, 0x75, 0x74, 0x68, 0x65, 0x6e, 0x74, 0x69, 0x63, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x18, 0x05, 0x20, 0x01, 0x28, 0x08, 0x52, 0x15, 0x64, 0x69, 0x73, 0x61, 0x62, 0x6c,
	0x65, 0x41, 0x75, 0x74, 0x68, 0x65, 0x6e, 0x74, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12,
	0x22, 0x0a, 0x0c, 0x64, 0x69, 0x73, 0x61, 0x62, 0x6c, 0x65, 0x43, 0x61, 0x63, 0x68, 0x65, 0x18,
	0x06, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0c, 0x64, 0x69, 0x73, 0x61, 0x62, 0x6c, 0x65, 0x43, 0x61,
	0x63, 0x68, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x73, 0x50, 0x61,
	0x74, 0x68, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e,
	0x73, 0x50, 0x61, 0x74, 0x68, 0x22, 0x34, 0x0a, 0x08, 0x49, 0x6e, 0x69, 0x74, 0x44, 0x61, 0x74,
	0x61, 0x12, 0x28, 0x0a, 0x06, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x10, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x41, 0x70, 0x70, 0x43, 0x6f, 0x6e,
	0x66, 0x69, 0x67, 0x52, 0x06, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x42, 0x25, 0x5a, 0x23, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x74, 0x69, 0x73, 0x6c, 0x69, 0x62,
	0x2f, 0x61, 0x70, 0x69, 0x62, 0x72, 0x65, 0x77, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x6d, 0x6f, 0x64,
	0x65, 0x6c, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_model_init_proto_rawDescOnce sync.Once
	file_model_init_proto_rawDescData = file_model_init_proto_rawDesc
)

func file_model_init_proto_rawDescGZIP() []byte {
	file_model_init_proto_rawDescOnce.Do(func() {
		file_model_init_proto_rawDescData = protoimpl.X.CompressGZIP(file_model_init_proto_rawDescData)
	})
	return file_model_init_proto_rawDescData
}

var file_model_init_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_model_init_proto_goTypes = []interface{}{
	(*AppConfig)(nil), // 0: model.AppConfig
	(*InitData)(nil),  // 1: model.InitData
}
var file_model_init_proto_depIdxs = []int32{
	0, // 0: model.InitData.config:type_name -> model.AppConfig
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_model_init_proto_init() }
func file_model_init_proto_init() {
	if File_model_init_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_model_init_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AppConfig); i {
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
		file_model_init_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*InitData); i {
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
			RawDescriptor: file_model_init_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_model_init_proto_goTypes,
		DependencyIndexes: file_model_init_proto_depIdxs,
		MessageInfos:      file_model_init_proto_msgTypes,
	}.Build()
	File_model_init_proto = out.File
	file_model_init_proto_rawDesc = nil
	file_model_init_proto_goTypes = nil
	file_model_init_proto_depIdxs = nil
}
