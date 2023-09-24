// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             (unknown)
// source: stub/event-channel.proto

package stub

import (
	context "context"
	model "github.com/apibrew/apibrew/pkg/model"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	EventChannel_Poll_FullMethodName  = "/stub.EventChannel/Poll"
	EventChannel_Write_FullMethodName = "/stub.EventChannel/Write"
)

// EventChannelClient is the client API for EventChannel service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type EventChannelClient interface {
	// Sends a greeting
	Poll(ctx context.Context, in *EventPollRequest, opts ...grpc.CallOption) (EventChannel_PollClient, error)
	Write(ctx context.Context, in *EventWriteRequest, opts ...grpc.CallOption) (*EventWriteResponse, error)
}

type eventChannelClient struct {
	cc grpc.ClientConnInterface
}

func NewEventChannelClient(cc grpc.ClientConnInterface) EventChannelClient {
	return &eventChannelClient{cc}
}

func (c *eventChannelClient) Poll(ctx context.Context, in *EventPollRequest, opts ...grpc.CallOption) (EventChannel_PollClient, error) {
	stream, err := c.cc.NewStream(ctx, &EventChannel_ServiceDesc.Streams[0], EventChannel_Poll_FullMethodName, opts...)
	if err != nil {
		return nil, err
	}
	x := &eventChannelPollClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type EventChannel_PollClient interface {
	Recv() (*model.Event, error)
	grpc.ClientStream
}

type eventChannelPollClient struct {
	grpc.ClientStream
}

func (x *eventChannelPollClient) Recv() (*model.Event, error) {
	m := new(model.Event)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *eventChannelClient) Write(ctx context.Context, in *EventWriteRequest, opts ...grpc.CallOption) (*EventWriteResponse, error) {
	out := new(EventWriteResponse)
	err := c.cc.Invoke(ctx, EventChannel_Write_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// EventChannelServer is the server API for EventChannel service.
// All implementations must embed UnimplementedEventChannelServer
// for forward compatibility
type EventChannelServer interface {
	// Sends a greeting
	Poll(*EventPollRequest, EventChannel_PollServer) error
	Write(context.Context, *EventWriteRequest) (*EventWriteResponse, error)
	mustEmbedUnimplementedEventChannelServer()
}

// UnimplementedEventChannelServer must be embedded to have forward compatible implementations.
type UnimplementedEventChannelServer struct {
}

func (UnimplementedEventChannelServer) Poll(*EventPollRequest, EventChannel_PollServer) error {
	return status.Errorf(codes.Unimplemented, "method Poll not implemented")
}
func (UnimplementedEventChannelServer) Write(context.Context, *EventWriteRequest) (*EventWriteResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Write not implemented")
}
func (UnimplementedEventChannelServer) mustEmbedUnimplementedEventChannelServer() {}

// UnsafeEventChannelServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to EventChannelServer will
// result in compilation errors.
type UnsafeEventChannelServer interface {
	mustEmbedUnimplementedEventChannelServer()
}

func RegisterEventChannelServer(s grpc.ServiceRegistrar, srv EventChannelServer) {
	s.RegisterService(&EventChannel_ServiceDesc, srv)
}

func _EventChannel_Poll_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(EventPollRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(EventChannelServer).Poll(m, &eventChannelPollServer{stream})
}

type EventChannel_PollServer interface {
	Send(*model.Event) error
	grpc.ServerStream
}

type eventChannelPollServer struct {
	grpc.ServerStream
}

func (x *eventChannelPollServer) Send(m *model.Event) error {
	return x.ServerStream.SendMsg(m)
}

func _EventChannel_Write_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EventWriteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EventChannelServer).Write(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: EventChannel_Write_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EventChannelServer).Write(ctx, req.(*EventWriteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// EventChannel_ServiceDesc is the grpc.ServiceDesc for EventChannel service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var EventChannel_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "stub.EventChannel",
	HandlerType: (*EventChannelServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Write",
			Handler:    _EventChannel_Write_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Poll",
			Handler:       _EventChannel_Poll_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "stub/event-channel.proto",
}