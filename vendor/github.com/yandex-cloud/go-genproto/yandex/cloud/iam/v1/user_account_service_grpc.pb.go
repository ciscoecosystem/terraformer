// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package iam

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

// UserAccountServiceClient is the client API for UserAccountService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type UserAccountServiceClient interface {
	// Returns the specified UserAccount resource.
	Get(ctx context.Context, in *GetUserAccountRequest, opts ...grpc.CallOption) (*UserAccount, error)
}

type userAccountServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewUserAccountServiceClient(cc grpc.ClientConnInterface) UserAccountServiceClient {
	return &userAccountServiceClient{cc}
}

func (c *userAccountServiceClient) Get(ctx context.Context, in *GetUserAccountRequest, opts ...grpc.CallOption) (*UserAccount, error) {
	out := new(UserAccount)
	err := c.cc.Invoke(ctx, "/yandex.cloud.iam.v1.UserAccountService/Get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UserAccountServiceServer is the server API for UserAccountService service.
// All implementations should embed UnimplementedUserAccountServiceServer
// for forward compatibility
type UserAccountServiceServer interface {
	// Returns the specified UserAccount resource.
	Get(context.Context, *GetUserAccountRequest) (*UserAccount, error)
}

// UnimplementedUserAccountServiceServer should be embedded to have forward compatible implementations.
type UnimplementedUserAccountServiceServer struct {
}

func (UnimplementedUserAccountServiceServer) Get(context.Context, *GetUserAccountRequest) (*UserAccount, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}

// UnsafeUserAccountServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to UserAccountServiceServer will
// result in compilation errors.
type UnsafeUserAccountServiceServer interface {
	mustEmbedUnimplementedUserAccountServiceServer()
}

func RegisterUserAccountServiceServer(s grpc.ServiceRegistrar, srv UserAccountServiceServer) {
	s.RegisterService(&UserAccountService_ServiceDesc, srv)
}

func _UserAccountService_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUserAccountRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserAccountServiceServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/yandex.cloud.iam.v1.UserAccountService/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserAccountServiceServer).Get(ctx, req.(*GetUserAccountRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// UserAccountService_ServiceDesc is the grpc.ServiceDesc for UserAccountService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var UserAccountService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "yandex.cloud.iam.v1.UserAccountService",
	HandlerType: (*UserAccountServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Get",
			Handler:    _UserAccountService_Get_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "yandex/cloud/iam/v1/user_account_service.proto",
}
