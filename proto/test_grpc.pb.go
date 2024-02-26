// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             (unknown)
// source: proto/test.proto

package proto

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
	AService_Test_FullMethodName = "/AService/Test"
)

// AServiceClient is the client API for AService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AServiceClient interface {
	Test(ctx context.Context, in *A, opts ...grpc.CallOption) (*A, error)
}

type aServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewAServiceClient(cc grpc.ClientConnInterface) AServiceClient {
	return &aServiceClient{cc}
}

func (c *aServiceClient) Test(ctx context.Context, in *A, opts ...grpc.CallOption) (*A, error) {
	out := new(A)
	err := c.cc.Invoke(ctx, AService_Test_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AServiceServer is the server API for AService service.
// All implementations must embed UnimplementedAServiceServer
// for forward compatibility
type AServiceServer interface {
	Test(context.Context, *A) (*A, error)
	mustEmbedUnimplementedAServiceServer()
}

// UnimplementedAServiceServer must be embedded to have forward compatible implementations.
type UnimplementedAServiceServer struct {
}

func (UnimplementedAServiceServer) Test(context.Context, *A) (*A, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Test not implemented")
}
func (UnimplementedAServiceServer) mustEmbedUnimplementedAServiceServer() {}

// UnsafeAServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AServiceServer will
// result in compilation errors.
type UnsafeAServiceServer interface {
	mustEmbedUnimplementedAServiceServer()
}

func RegisterAServiceServer(s grpc.ServiceRegistrar, srv AServiceServer) {
	s.RegisterService(&AService_ServiceDesc, srv)
}

func _AService_Test_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(A)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AServiceServer).Test(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AService_Test_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AServiceServer).Test(ctx, req.(*A))
	}
	return interceptor(ctx, in, info, handler)
}

// AService_ServiceDesc is the grpc.ServiceDesc for AService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var AService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "AService",
	HandlerType: (*AServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Test",
			Handler:    _AService_Test_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/test.proto",
}
