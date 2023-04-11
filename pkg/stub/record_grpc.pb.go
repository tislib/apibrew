// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             (unknown)
// source: stub/record.proto

package stub

import (
	context "context"
	model "github.com/tislib/data-handler/pkg/model"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// RecordClient is the client API for Record service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type RecordClient interface {
	Create(ctx context.Context, in *CreateRecordRequest, opts ...grpc.CallOption) (*CreateRecordResponse, error)
	Update(ctx context.Context, in *UpdateRecordRequest, opts ...grpc.CallOption) (*UpdateRecordResponse, error)
	UpdateMulti(ctx context.Context, in *UpdateMultiRecordRequest, opts ...grpc.CallOption) (*UpdateMultiRecordResponse, error)
	Delete(ctx context.Context, in *DeleteRecordRequest, opts ...grpc.CallOption) (*DeleteRecordResponse, error)
	List(ctx context.Context, in *ListRecordRequest, opts ...grpc.CallOption) (*ListRecordResponse, error)
	Search(ctx context.Context, in *SearchRecordRequest, opts ...grpc.CallOption) (*SearchRecordResponse, error)
	ReadStream(ctx context.Context, in *ReadStreamRequest, opts ...grpc.CallOption) (Record_ReadStreamClient, error)
	WriteStream(ctx context.Context, opts ...grpc.CallOption) (Record_WriteStreamClient, error)
	Get(ctx context.Context, in *GetRecordRequest, opts ...grpc.CallOption) (*GetRecordResponse, error)
}

type recordClient struct {
	cc grpc.ClientConnInterface
}

func NewRecordClient(cc grpc.ClientConnInterface) RecordClient {
	return &recordClient{cc}
}

func (c *recordClient) Create(ctx context.Context, in *CreateRecordRequest, opts ...grpc.CallOption) (*CreateRecordResponse, error) {
	out := new(CreateRecordResponse)
	err := c.cc.Invoke(ctx, "/stub.Record/Create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *recordClient) Update(ctx context.Context, in *UpdateRecordRequest, opts ...grpc.CallOption) (*UpdateRecordResponse, error) {
	out := new(UpdateRecordResponse)
	err := c.cc.Invoke(ctx, "/stub.Record/Update", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *recordClient) UpdateMulti(ctx context.Context, in *UpdateMultiRecordRequest, opts ...grpc.CallOption) (*UpdateMultiRecordResponse, error) {
	out := new(UpdateMultiRecordResponse)
	err := c.cc.Invoke(ctx, "/stub.Record/UpdateMulti", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *recordClient) Delete(ctx context.Context, in *DeleteRecordRequest, opts ...grpc.CallOption) (*DeleteRecordResponse, error) {
	out := new(DeleteRecordResponse)
	err := c.cc.Invoke(ctx, "/stub.Record/Delete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *recordClient) List(ctx context.Context, in *ListRecordRequest, opts ...grpc.CallOption) (*ListRecordResponse, error) {
	out := new(ListRecordResponse)
	err := c.cc.Invoke(ctx, "/stub.Record/List", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *recordClient) Search(ctx context.Context, in *SearchRecordRequest, opts ...grpc.CallOption) (*SearchRecordResponse, error) {
	out := new(SearchRecordResponse)
	err := c.cc.Invoke(ctx, "/stub.Record/Search", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *recordClient) ReadStream(ctx context.Context, in *ReadStreamRequest, opts ...grpc.CallOption) (Record_ReadStreamClient, error) {
	stream, err := c.cc.NewStream(ctx, &Record_ServiceDesc.Streams[0], "/stub.Record/ReadStream", opts...)
	if err != nil {
		return nil, err
	}
	x := &recordReadStreamClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Record_ReadStreamClient interface {
	Recv() (*model.Record, error)
	grpc.ClientStream
}

type recordReadStreamClient struct {
	grpc.ClientStream
}

func (x *recordReadStreamClient) Recv() (*model.Record, error) {
	m := new(model.Record)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *recordClient) WriteStream(ctx context.Context, opts ...grpc.CallOption) (Record_WriteStreamClient, error) {
	stream, err := c.cc.NewStream(ctx, &Record_ServiceDesc.Streams[1], "/stub.Record/WriteStream", opts...)
	if err != nil {
		return nil, err
	}
	x := &recordWriteStreamClient{stream}
	return x, nil
}

type Record_WriteStreamClient interface {
	Send(*model.Record) error
	CloseAndRecv() (*WriteStreamResponse, error)
	grpc.ClientStream
}

type recordWriteStreamClient struct {
	grpc.ClientStream
}

func (x *recordWriteStreamClient) Send(m *model.Record) error {
	return x.ClientStream.SendMsg(m)
}

func (x *recordWriteStreamClient) CloseAndRecv() (*WriteStreamResponse, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(WriteStreamResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *recordClient) Get(ctx context.Context, in *GetRecordRequest, opts ...grpc.CallOption) (*GetRecordResponse, error) {
	out := new(GetRecordResponse)
	err := c.cc.Invoke(ctx, "/stub.Record/Get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RecordServer is the server API for Record service.
// All implementations must embed UnimplementedRecordServer
// for forward compatibility
type RecordServer interface {
	Create(context.Context, *CreateRecordRequest) (*CreateRecordResponse, error)
	Update(context.Context, *UpdateRecordRequest) (*UpdateRecordResponse, error)
	UpdateMulti(context.Context, *UpdateMultiRecordRequest) (*UpdateMultiRecordResponse, error)
	Delete(context.Context, *DeleteRecordRequest) (*DeleteRecordResponse, error)
	List(context.Context, *ListRecordRequest) (*ListRecordResponse, error)
	Search(context.Context, *SearchRecordRequest) (*SearchRecordResponse, error)
	ReadStream(*ReadStreamRequest, Record_ReadStreamServer) error
	WriteStream(Record_WriteStreamServer) error
	Get(context.Context, *GetRecordRequest) (*GetRecordResponse, error)
	mustEmbedUnimplementedRecordServer()
}

// UnimplementedRecordServer must be embedded to have forward compatible implementations.
type UnimplementedRecordServer struct {
}

func (UnimplementedRecordServer) Create(context.Context, *CreateRecordRequest) (*CreateRecordResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedRecordServer) Update(context.Context, *UpdateRecordRequest) (*UpdateRecordResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (UnimplementedRecordServer) UpdateMulti(context.Context, *UpdateMultiRecordRequest) (*UpdateMultiRecordResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateMulti not implemented")
}
func (UnimplementedRecordServer) Delete(context.Context, *DeleteRecordRequest) (*DeleteRecordResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (UnimplementedRecordServer) List(context.Context, *ListRecordRequest) (*ListRecordResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method List not implemented")
}
func (UnimplementedRecordServer) Search(context.Context, *SearchRecordRequest) (*SearchRecordResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Search not implemented")
}
func (UnimplementedRecordServer) ReadStream(*ReadStreamRequest, Record_ReadStreamServer) error {
	return status.Errorf(codes.Unimplemented, "method ReadStream not implemented")
}
func (UnimplementedRecordServer) WriteStream(Record_WriteStreamServer) error {
	return status.Errorf(codes.Unimplemented, "method WriteStream not implemented")
}
func (UnimplementedRecordServer) Get(context.Context, *GetRecordRequest) (*GetRecordResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (UnimplementedRecordServer) mustEmbedUnimplementedRecordServer() {}

// UnsafeRecordServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to RecordServer will
// result in compilation errors.
type UnsafeRecordServer interface {
	mustEmbedUnimplementedRecordServer()
}

func RegisterRecordServer(s grpc.ServiceRegistrar, srv RecordServer) {
	s.RegisterService(&Record_ServiceDesc, srv)
}

func _Record_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateRecordRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RecordServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/stub.Record/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RecordServer).Create(ctx, req.(*CreateRecordRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Record_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateRecordRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RecordServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/stub.Record/Update",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RecordServer).Update(ctx, req.(*UpdateRecordRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Record_UpdateMulti_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateMultiRecordRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RecordServer).UpdateMulti(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/stub.Record/UpdateMulti",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RecordServer).UpdateMulti(ctx, req.(*UpdateMultiRecordRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Record_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteRecordRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RecordServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/stub.Record/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RecordServer).Delete(ctx, req.(*DeleteRecordRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Record_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListRecordRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RecordServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/stub.Record/List",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RecordServer).List(ctx, req.(*ListRecordRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Record_Search_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SearchRecordRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RecordServer).Search(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/stub.Record/Search",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RecordServer).Search(ctx, req.(*SearchRecordRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Record_ReadStream_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(ReadStreamRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(RecordServer).ReadStream(m, &recordReadStreamServer{stream})
}

type Record_ReadStreamServer interface {
	Send(*model.Record) error
	grpc.ServerStream
}

type recordReadStreamServer struct {
	grpc.ServerStream
}

func (x *recordReadStreamServer) Send(m *model.Record) error {
	return x.ServerStream.SendMsg(m)
}

func _Record_WriteStream_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(RecordServer).WriteStream(&recordWriteStreamServer{stream})
}

type Record_WriteStreamServer interface {
	SendAndClose(*WriteStreamResponse) error
	Recv() (*model.Record, error)
	grpc.ServerStream
}

type recordWriteStreamServer struct {
	grpc.ServerStream
}

func (x *recordWriteStreamServer) SendAndClose(m *WriteStreamResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *recordWriteStreamServer) Recv() (*model.Record, error) {
	m := new(model.Record)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _Record_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRecordRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RecordServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/stub.Record/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RecordServer).Get(ctx, req.(*GetRecordRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Record_ServiceDesc is the grpc.ServiceDesc for Record service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Record_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "stub.Record",
	HandlerType: (*RecordServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _Record_Create_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _Record_Update_Handler,
		},
		{
			MethodName: "UpdateMulti",
			Handler:    _Record_UpdateMulti_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _Record_Delete_Handler,
		},
		{
			MethodName: "List",
			Handler:    _Record_List_Handler,
		},
		{
			MethodName: "Search",
			Handler:    _Record_Search_Handler,
		},
		{
			MethodName: "Get",
			Handler:    _Record_Get_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "ReadStream",
			Handler:       _Record_ReadStream_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "WriteStream",
			Handler:       _Record_WriteStream_Handler,
			ClientStreams: true,
		},
	},
	Metadata: "stub/record.proto",
}
