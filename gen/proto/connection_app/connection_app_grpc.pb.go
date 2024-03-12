// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             (unknown)
// source: proto/connection_app/connection_app.proto

package pb

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	ConnectionApp_PostConnectionApp_FullMethodName = "/proto.connection_app.ConnectionApp/PostConnectionApp"
	ConnectionApp_GetConnectionApp_FullMethodName  = "/proto.connection_app.ConnectionApp/GetConnectionApp"
)

// ConnectionAppClient is the client API for ConnectionApp service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ConnectionAppClient interface {
	PostConnectionApp(ctx context.Context, in *ConnectionAppBodyRequest, opts ...grpc.CallOption) (*ConnectionAppResponse, error)
	GetConnectionApp(ctx context.Context, in *ConnectionAppRequest, opts ...grpc.CallOption) (*ConnectionAppStructResponse, error)
}

type connectionAppClient struct {
	cc grpc.ClientConnInterface
}

func NewConnectionAppClient(cc grpc.ClientConnInterface) ConnectionAppClient {
	return &connectionAppClient{cc}
}

func (c *connectionAppClient) PostConnectionApp(ctx context.Context, in *ConnectionAppBodyRequest, opts ...grpc.CallOption) (*ConnectionAppResponse, error) {
	out := new(ConnectionAppResponse)
	err := c.cc.Invoke(ctx, ConnectionApp_PostConnectionApp_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *connectionAppClient) GetConnectionApp(ctx context.Context, in *ConnectionAppRequest, opts ...grpc.CallOption) (*ConnectionAppStructResponse, error) {
	out := new(ConnectionAppStructResponse)
	err := c.cc.Invoke(ctx, ConnectionApp_GetConnectionApp_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ConnectionAppServer is the server API for ConnectionApp service.
// All implementations should embed UnimplementedConnectionAppServer
// for forward compatibility
type ConnectionAppServer interface {
	PostConnectionApp(context.Context, *ConnectionAppBodyRequest) (*ConnectionAppResponse, error)
	GetConnectionApp(context.Context, *ConnectionAppRequest) (*ConnectionAppStructResponse, error)
}

// UnimplementedConnectionAppServer should be embedded to have forward compatible implementations.
type UnimplementedConnectionAppServer struct {
}

func (UnimplementedConnectionAppServer) PostConnectionApp(context.Context, *ConnectionAppBodyRequest) (*ConnectionAppResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PostConnectionApp not implemented")
}
func (UnimplementedConnectionAppServer) GetConnectionApp(context.Context, *ConnectionAppRequest) (*ConnectionAppStructResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetConnectionApp not implemented")
}

// UnsafeConnectionAppServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ConnectionAppServer will
// result in compilation errors.
type UnsafeConnectionAppServer interface {
	mustEmbedUnimplementedConnectionAppServer()
}

func RegisterConnectionAppServer(s grpc.ServiceRegistrar, srv ConnectionAppServer) {
	s.RegisterService(&ConnectionApp_ServiceDesc, srv)
}

func _ConnectionApp_PostConnectionApp_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ConnectionAppBodyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ConnectionAppServer).PostConnectionApp(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ConnectionApp_PostConnectionApp_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ConnectionAppServer).PostConnectionApp(ctx, req.(*ConnectionAppBodyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ConnectionApp_GetConnectionApp_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ConnectionAppRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ConnectionAppServer).GetConnectionApp(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ConnectionApp_GetConnectionApp_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ConnectionAppServer).GetConnectionApp(ctx, req.(*ConnectionAppRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ConnectionApp_ServiceDesc is the grpc.ServiceDesc for ConnectionApp service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ConnectionApp_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "proto.connection_app.ConnectionApp",
	HandlerType: (*ConnectionAppServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "PostConnectionApp",
			Handler:    _ConnectionApp_PostConnectionApp_Handler,
		},
		{
			MethodName: "GetConnectionApp",
			Handler:    _ConnectionApp_GetConnectionApp_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/connection_app/connection_app.proto",
}
