// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        (unknown)
// source: stub/openapi.proto

package stub

import (
	_ "github.com/google/gnostic/openapiv3"
	_ "github.com/tislib/data-handler/pkg/model"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

var File_stub_openapi_proto protoreflect.FileDescriptor

var file_stub_openapi_proto_rawDesc = []byte{
	0x0a, 0x12, 0x73, 0x74, 0x75, 0x62, 0x2f, 0x6f, 0x70, 0x65, 0x6e, 0x61, 0x70, 0x69, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04, 0x73, 0x74, 0x75, 0x62, 0x1a, 0x11, 0x6d, 0x6f, 0x64, 0x65,
	0x6c, 0x2f, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x11, 0x6d,
	0x6f, 0x64, 0x65, 0x6c, 0x2f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e,
	0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b,
	0x6f, 0x70, 0x65, 0x6e, 0x61, 0x70, 0x69, 0x76, 0x33, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x42, 0xc8, 0x02, 0x5a, 0x27,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x74, 0x69, 0x73, 0x6c, 0x69,
	0x62, 0x2f, 0x64, 0x61, 0x74, 0x61, 0x2d, 0x68, 0x61, 0x6e, 0x64, 0x6c, 0x65, 0x72, 0x2f, 0x70,
	0x6b, 0x67, 0x2f, 0x73, 0x74, 0x75, 0x62, 0xba, 0x47, 0x9b, 0x02, 0x12, 0xdd, 0x01, 0x0a, 0x15,
	0x54, 0x69, 0x74, 0x6c, 0x65, 0x20, 0x66, 0x72, 0x6f, 0x6d, 0x20, 0x61, 0x6e, 0x6e, 0x6f, 0x74,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1b, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x20, 0x66, 0x72, 0x6f, 0x6d, 0x20, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x22, 0x45, 0x0a, 0x0c, 0x43, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x20, 0x4e, 0x61,
	0x6d, 0x65, 0x12, 0x21, 0x68, 0x74, 0x74, 0x70, 0x73, 0x3a, 0x2f, 0x2f, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x67, 0x6e,
	0x6f, 0x73, 0x74, 0x69, 0x63, 0x1a, 0x12, 0x67, 0x6e, 0x6f, 0x73, 0x74, 0x69, 0x63, 0x40, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6f, 0x6d, 0x2a, 0x47, 0x0a, 0x0e, 0x41, 0x70, 0x61,
	0x63, 0x68, 0x65, 0x20, 0x4c, 0x69, 0x63, 0x65, 0x6e, 0x73, 0x65, 0x12, 0x35, 0x68, 0x74, 0x74,
	0x70, 0x73, 0x3a, 0x2f, 0x2f, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x67, 0x6e, 0x6f, 0x73, 0x74, 0x69, 0x63, 0x2f, 0x62,
	0x6c, 0x6f, 0x62, 0x2f, 0x6d, 0x61, 0x73, 0x74, 0x65, 0x72, 0x2f, 0x4c, 0x49, 0x43, 0x45, 0x4e,
	0x53, 0x45, 0x32, 0x17, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x20, 0x66, 0x72, 0x6f, 0x6d,
	0x20, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2a, 0x27, 0x3a, 0x25, 0x0a,
	0x23, 0x0a, 0x0a, 0x42, 0x65, 0x61, 0x72, 0x65, 0x72, 0x41, 0x75, 0x74, 0x68, 0x12, 0x15, 0x0a,
	0x13, 0x0a, 0x04, 0x68, 0x74, 0x74, 0x70, 0x2a, 0x06, 0x62, 0x65, 0x61, 0x72, 0x65, 0x72, 0x32,
	0x03, 0x4a, 0x57, 0x54, 0x32, 0x10, 0x0a, 0x0e, 0x0a, 0x0a, 0x62, 0x65, 0x61, 0x72, 0x65, 0x72,
	0x41, 0x75, 0x74, 0x68, 0x12, 0x00, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var file_stub_openapi_proto_goTypes = []interface{}{}
var file_stub_openapi_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_stub_openapi_proto_init() }
func file_stub_openapi_proto_init() {
	if File_stub_openapi_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_stub_openapi_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_stub_openapi_proto_goTypes,
		DependencyIndexes: file_stub_openapi_proto_depIdxs,
	}.Build()
	File_stub_openapi_proto = out.File
	file_stub_openapi_proto_rawDesc = nil
	file_stub_openapi_proto_goTypes = nil
	file_stub_openapi_proto_depIdxs = nil
}
