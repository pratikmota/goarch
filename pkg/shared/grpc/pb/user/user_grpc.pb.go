// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.19.4
// source: user/user.proto

package user

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

// UserCrudHandlerClient is the client API for UserCrudHandler service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type UserCrudHandlerClient interface {
	RegisterUser(ctx context.Context, in *RegisterRequest, opts ...grpc.CallOption) (*RegisterResponse, error)
}

type userCrudHandlerClient struct {
	cc grpc.ClientConnInterface
}

func NewUserCrudHandlerClient(cc grpc.ClientConnInterface) UserCrudHandlerClient {
	return &userCrudHandlerClient{cc}
}

func (c *userCrudHandlerClient) RegisterUser(ctx context.Context, in *RegisterRequest, opts ...grpc.CallOption) (*RegisterResponse, error) {
	out := new(RegisterResponse)
	err := c.cc.Invoke(ctx, "/com.goarch.grpc.auth.UserCrudHandler/RegisterUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UserCrudHandlerServer is the server API for UserCrudHandler service.
// All implementations must embed UnimplementedUserCrudHandlerServer
// for forward compatibility
type UserCrudHandlerServer interface {
	RegisterUser(context.Context, *RegisterRequest) (*RegisterResponse, error)
	mustEmbedUnimplementedUserCrudHandlerServer()
}

// UnimplementedUserCrudHandlerServer must be embedded to have forward compatible implementations.
type UnimplementedUserCrudHandlerServer struct {
}

func (UnimplementedUserCrudHandlerServer) RegisterUser(context.Context, *RegisterRequest) (*RegisterResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RegisterUser not implemented")
}
func (UnimplementedUserCrudHandlerServer) mustEmbedUnimplementedUserCrudHandlerServer() {}

// UnsafeUserCrudHandlerServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to UserCrudHandlerServer will
// result in compilation errors.
type UnsafeUserCrudHandlerServer interface {
	mustEmbedUnimplementedUserCrudHandlerServer()
}

func RegisterUserCrudHandlerServer(s grpc.ServiceRegistrar, srv UserCrudHandlerServer) {
	s.RegisterService(&UserCrudHandler_ServiceDesc, srv)
}

func _UserCrudHandler_RegisterUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RegisterRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserCrudHandlerServer).RegisterUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/com.goarch.grpc.auth.UserCrudHandler/RegisterUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserCrudHandlerServer).RegisterUser(ctx, req.(*RegisterRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// UserCrudHandler_ServiceDesc is the grpc.ServiceDesc for UserCrudHandler service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var UserCrudHandler_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "com.goarch.grpc.auth.UserCrudHandler",
	HandlerType: (*UserCrudHandlerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "RegisterUser",
			Handler:    _UserCrudHandler_RegisterUser_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "user/user.proto",
}