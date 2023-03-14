// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.9
// source: mio-api.proto

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

// MioApiClient is the client API for MioApi service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MioApiClient interface {
	GetSecretKey(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error)
}

type mioApiClient struct {
	cc grpc.ClientConnInterface
}

func NewMioApiClient(cc grpc.ClientConnInterface) MioApiClient {
	return &mioApiClient{cc}
}

func (c *mioApiClient) GetSecretKey(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/mioApi/GetSecretKey", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MioApiServer is the server API for MioApi service.
// All implementations must embed UnimplementedMioApiServer
// for forward compatibility
type MioApiServer interface {
	GetSecretKey(context.Context, *Request) (*Response, error)
	mustEmbedUnimplementedMioApiServer()
}

// UnimplementedMioApiServer must be embedded to have forward compatible implementations.
type UnimplementedMioApiServer struct {
}

func (UnimplementedMioApiServer) GetSecretKey(context.Context, *Request) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetSecretKey not implemented")
}
func (UnimplementedMioApiServer) mustEmbedUnimplementedMioApiServer() {}

// UnsafeMioApiServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MioApiServer will
// result in compilation errors.
type UnsafeMioApiServer interface {
	mustEmbedUnimplementedMioApiServer()
}

func RegisterMioApiServer(s grpc.ServiceRegistrar, srv MioApiServer) {
	s.RegisterService(&MioApi_ServiceDesc, srv)
}

func _MioApi_GetSecretKey_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MioApiServer).GetSecretKey(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mioApi/GetSecretKey",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MioApiServer).GetSecretKey(ctx, req.(*Request))
	}
	return interceptor(ctx, in, info, handler)
}

// MioApi_ServiceDesc is the grpc.ServiceDesc for MioApi service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var MioApi_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "mioApi",
	HandlerType: (*MioApiServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetSecretKey",
			Handler:    _MioApi_GetSecretKey_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "mio-api.proto",
}
