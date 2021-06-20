// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package __

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

// SubClient is the client API for Sub service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type SubClient interface {
	GetToken(ctx context.Context, in *GetTokenRequest, opts ...grpc.CallOption) (*GetTokenResponse, error)
}

type subClient struct {
	cc grpc.ClientConnInterface
}

func NewSubClient(cc grpc.ClientConnInterface) SubClient {
	return &subClient{cc}
}

func (c *subClient) GetToken(ctx context.Context, in *GetTokenRequest, opts ...grpc.CallOption) (*GetTokenResponse, error) {
	out := new(GetTokenResponse)
	err := c.cc.Invoke(ctx, "/Sub/GetToken", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SubServer is the server API for Sub service.
// All implementations must embed UnimplementedSubServer
// for forward compatibility
type SubServer interface {
	GetToken(context.Context, *GetTokenRequest) (*GetTokenResponse, error)
	mustEmbedUnimplementedSubServer()
}

// UnimplementedSubServer must be embedded to have forward compatible implementations.
type UnimplementedSubServer struct {
}

func (UnimplementedSubServer) GetToken(context.Context, *GetTokenRequest) (*GetTokenResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetToken not implemented")
}
func (UnimplementedSubServer) mustEmbedUnimplementedSubServer() {}

// UnsafeSubServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to SubServer will
// result in compilation errors.
type UnsafeSubServer interface {
	mustEmbedUnimplementedSubServer()
}

func RegisterSubServer(s grpc.ServiceRegistrar, srv SubServer) {
	s.RegisterService(&Sub_ServiceDesc, srv)
}

func _Sub_GetToken_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetTokenRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SubServer).GetToken(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Sub/GetToken",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SubServer).GetToken(ctx, req.(*GetTokenRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Sub_ServiceDesc is the grpc.ServiceDesc for Sub service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Sub_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "Sub",
	HandlerType: (*SubServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetToken",
			Handler:    _Sub_GetToken_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "protocol/sub.proto",
}