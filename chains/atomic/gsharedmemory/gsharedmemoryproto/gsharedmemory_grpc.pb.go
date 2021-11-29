// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package gsharedmemoryproto

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

// SharedMemoryClient is the client API for SharedMemory service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type SharedMemoryClient interface {
	Get(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*GetResponse, error)
	Indexed(ctx context.Context, in *IndexedRequest, opts ...grpc.CallOption) (*IndexedResponse, error)
	Apply(ctx context.Context, in *ApplyRequest, opts ...grpc.CallOption) (*ApplyResponse, error)
}

type sharedMemoryClient struct {
	cc grpc.ClientConnInterface
}

func NewSharedMemoryClient(cc grpc.ClientConnInterface) SharedMemoryClient {
	return &sharedMemoryClient{cc}
}

func (c *sharedMemoryClient) Get(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*GetResponse, error) {
	out := new(GetResponse)
	err := c.cc.Invoke(ctx, "/gsharedmemoryproto.SharedMemory/Get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sharedMemoryClient) Indexed(ctx context.Context, in *IndexedRequest, opts ...grpc.CallOption) (*IndexedResponse, error) {
	out := new(IndexedResponse)
	err := c.cc.Invoke(ctx, "/gsharedmemoryproto.SharedMemory/Indexed", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sharedMemoryClient) Apply(ctx context.Context, in *ApplyRequest, opts ...grpc.CallOption) (*ApplyResponse, error) {
	out := new(ApplyResponse)
	err := c.cc.Invoke(ctx, "/gsharedmemoryproto.SharedMemory/Apply", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SharedMemoryServer is the server API for SharedMemory service.
// All implementations must embed UnimplementedSharedMemoryServer
// for forward compatibility
type SharedMemoryServer interface {
	Get(context.Context, *GetRequest) (*GetResponse, error)
	Indexed(context.Context, *IndexedRequest) (*IndexedResponse, error)
	Apply(context.Context, *ApplyRequest) (*ApplyResponse, error)
	mustEmbedUnimplementedSharedMemoryServer()
}

// UnimplementedSharedMemoryServer must be embedded to have forward compatible implementations.
type UnimplementedSharedMemoryServer struct{}

func (UnimplementedSharedMemoryServer) Get(context.Context, *GetRequest) (*GetResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}

func (UnimplementedSharedMemoryServer) Indexed(context.Context, *IndexedRequest) (*IndexedResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Indexed not implemented")
}

func (UnimplementedSharedMemoryServer) Apply(context.Context, *ApplyRequest) (*ApplyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Apply not implemented")
}
func (UnimplementedSharedMemoryServer) mustEmbedUnimplementedSharedMemoryServer() {}

// UnsafeSharedMemoryServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to SharedMemoryServer will
// result in compilation errors.
type UnsafeSharedMemoryServer interface {
	mustEmbedUnimplementedSharedMemoryServer()
}

func RegisterSharedMemoryServer(s grpc.ServiceRegistrar, srv SharedMemoryServer) {
	s.RegisterService(&SharedMemory_ServiceDesc, srv)
}

func _SharedMemory_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SharedMemoryServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/gsharedmemoryproto.SharedMemory/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SharedMemoryServer).Get(ctx, req.(*GetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SharedMemory_Indexed_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IndexedRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SharedMemoryServer).Indexed(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/gsharedmemoryproto.SharedMemory/Indexed",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SharedMemoryServer).Indexed(ctx, req.(*IndexedRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SharedMemory_Apply_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ApplyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SharedMemoryServer).Apply(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/gsharedmemoryproto.SharedMemory/Apply",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SharedMemoryServer).Apply(ctx, req.(*ApplyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// SharedMemory_ServiceDesc is the grpc.ServiceDesc for SharedMemory service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var SharedMemory_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "gsharedmemoryproto.SharedMemory",
	HandlerType: (*SharedMemoryServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Get",
			Handler:    _SharedMemory_Get_Handler,
		},
		{
			MethodName: "Indexed",
			Handler:    _SharedMemory_Indexed_Handler,
		},
		{
			MethodName: "Apply",
			Handler:    _SharedMemory_Apply_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "gsharedmemory.proto",
}
