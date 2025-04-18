// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v3.21.12
// source: user.proto

package userpb

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	SendData_Send_FullMethodName = "/user.Send_data/Send"
	SendData_Get_FullMethodName  = "/user.Send_data/Get"
)

// SendDataClient is the client API for SendData service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type SendDataClient interface {
	Send(ctx context.Context, in *DataSend, opts ...grpc.CallOption) (*Return, error)
	Get(ctx context.Context, in *GetDataSend, opts ...grpc.CallOption) (*GetDataReturn, error)
}

type sendDataClient struct {
	cc grpc.ClientConnInterface
}

func NewSendDataClient(cc grpc.ClientConnInterface) SendDataClient {
	return &sendDataClient{cc}
}

func (c *sendDataClient) Send(ctx context.Context, in *DataSend, opts ...grpc.CallOption) (*Return, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Return)
	err := c.cc.Invoke(ctx, SendData_Send_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sendDataClient) Get(ctx context.Context, in *GetDataSend, opts ...grpc.CallOption) (*GetDataReturn, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetDataReturn)
	err := c.cc.Invoke(ctx, SendData_Get_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SendDataServer is the server API for SendData service.
// All implementations must embed UnimplementedSendDataServer
// for forward compatibility.
type SendDataServer interface {
	Send(context.Context, *DataSend) (*Return, error)
	Get(context.Context, *GetDataSend) (*GetDataReturn, error)
	mustEmbedUnimplementedSendDataServer()
}

// UnimplementedSendDataServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedSendDataServer struct{}

func (UnimplementedSendDataServer) Send(context.Context, *DataSend) (*Return, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Send not implemented")
}
func (UnimplementedSendDataServer) Get(context.Context, *GetDataSend) (*GetDataReturn, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (UnimplementedSendDataServer) mustEmbedUnimplementedSendDataServer() {}
func (UnimplementedSendDataServer) testEmbeddedByValue()                  {}

// UnsafeSendDataServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to SendDataServer will
// result in compilation errors.
type UnsafeSendDataServer interface {
	mustEmbedUnimplementedSendDataServer()
}

func RegisterSendDataServer(s grpc.ServiceRegistrar, srv SendDataServer) {
	// If the following call pancis, it indicates UnimplementedSendDataServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&SendData_ServiceDesc, srv)
}

func _SendData_Send_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DataSend)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SendDataServer).Send(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SendData_Send_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SendDataServer).Send(ctx, req.(*DataSend))
	}
	return interceptor(ctx, in, info, handler)
}

func _SendData_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetDataSend)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SendDataServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SendData_Get_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SendDataServer).Get(ctx, req.(*GetDataSend))
	}
	return interceptor(ctx, in, info, handler)
}

// SendData_ServiceDesc is the grpc.ServiceDesc for SendData service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var SendData_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "user.Send_data",
	HandlerType: (*SendDataServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Send",
			Handler:    _SendData_Send_Handler,
		},
		{
			MethodName: "Get",
			Handler:    _SendData_Get_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "user.proto",
}
