// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.4.0
// - protoc             v3.6.1
// source: semaphore.proto

package __

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.62.0 or later.
const _ = grpc.SupportPackageIsVersion8

const (
	SemaphoreService_Increase_FullMethodName = "/semaphore.SemaphoreService/Increase"
	SemaphoreService_Decrease_FullMethodName = "/semaphore.SemaphoreService/Decrease"
)

// SemaphoreServiceClient is the client API for SemaphoreService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type SemaphoreServiceClient interface {
	Increase(ctx context.Context, in *IncreaseRequest, opts ...grpc.CallOption) (*SemaphoreResponse, error)
	Decrease(ctx context.Context, in *DecreaseRequest, opts ...grpc.CallOption) (*SemaphoreResponse, error)
}

type semaphoreServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewSemaphoreServiceClient(cc grpc.ClientConnInterface) SemaphoreServiceClient {
	return &semaphoreServiceClient{cc}
}

func (c *semaphoreServiceClient) Increase(ctx context.Context, in *IncreaseRequest, opts ...grpc.CallOption) (*SemaphoreResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(SemaphoreResponse)
	err := c.cc.Invoke(ctx, SemaphoreService_Increase_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *semaphoreServiceClient) Decrease(ctx context.Context, in *DecreaseRequest, opts ...grpc.CallOption) (*SemaphoreResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(SemaphoreResponse)
	err := c.cc.Invoke(ctx, SemaphoreService_Decrease_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SemaphoreServiceServer is the server API for SemaphoreService service.
// All implementations must embed UnimplementedSemaphoreServiceServer
// for forward compatibility
type SemaphoreServiceServer interface {
	Increase(context.Context, *IncreaseRequest) (*SemaphoreResponse, error)
	Decrease(context.Context, *DecreaseRequest) (*SemaphoreResponse, error)
	mustEmbedUnimplementedSemaphoreServiceServer()
}

// UnimplementedSemaphoreServiceServer must be embedded to have forward compatible implementations.
type UnimplementedSemaphoreServiceServer struct {
}

func (UnimplementedSemaphoreServiceServer) Increase(context.Context, *IncreaseRequest) (*SemaphoreResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Increase not implemented")
}
func (UnimplementedSemaphoreServiceServer) Decrease(context.Context, *DecreaseRequest) (*SemaphoreResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Decrease not implemented")
}
func (UnimplementedSemaphoreServiceServer) mustEmbedUnimplementedSemaphoreServiceServer() {}

// UnsafeSemaphoreServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to SemaphoreServiceServer will
// result in compilation errors.
type UnsafeSemaphoreServiceServer interface {
	mustEmbedUnimplementedSemaphoreServiceServer()
}

func RegisterSemaphoreServiceServer(s grpc.ServiceRegistrar, srv SemaphoreServiceServer) {
	s.RegisterService(&SemaphoreService_ServiceDesc, srv)
}

func _SemaphoreService_Increase_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IncreaseRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SemaphoreServiceServer).Increase(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SemaphoreService_Increase_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SemaphoreServiceServer).Increase(ctx, req.(*IncreaseRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SemaphoreService_Decrease_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DecreaseRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SemaphoreServiceServer).Decrease(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SemaphoreService_Decrease_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SemaphoreServiceServer).Decrease(ctx, req.(*DecreaseRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// SemaphoreService_ServiceDesc is the grpc.ServiceDesc for SemaphoreService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var SemaphoreService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "semaphore.SemaphoreService",
	HandlerType: (*SemaphoreServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Increase",
			Handler:    _SemaphoreService_Increase_Handler,
		},
		{
			MethodName: "Decrease",
			Handler:    _SemaphoreService_Decrease_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "semaphore.proto",
}
