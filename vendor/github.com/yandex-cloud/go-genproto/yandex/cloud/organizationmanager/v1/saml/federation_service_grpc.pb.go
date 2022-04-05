// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package saml

import (
	context "context"
	operation "github.com/yandex-cloud/go-genproto/yandex/cloud/operation"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// FederationServiceClient is the client API for FederationService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type FederationServiceClient interface {
	// Returns the specified federation.
	//
	// To get the list of available federations, make a [List] request.
	Get(ctx context.Context, in *GetFederationRequest, opts ...grpc.CallOption) (*Federation, error)
	// Retrieves the list of federations in the specified organization.
	List(ctx context.Context, in *ListFederationsRequest, opts ...grpc.CallOption) (*ListFederationsResponse, error)
	// Creates a federation in the specified organization.
	Create(ctx context.Context, in *CreateFederationRequest, opts ...grpc.CallOption) (*operation.Operation, error)
	// Updates the specified federation.
	Update(ctx context.Context, in *UpdateFederationRequest, opts ...grpc.CallOption) (*operation.Operation, error)
	// Deletes the specified federation.
	Delete(ctx context.Context, in *DeleteFederationRequest, opts ...grpc.CallOption) (*operation.Operation, error)
	// Adds users to the specified federation.
	AddUserAccounts(ctx context.Context, in *AddFederatedUserAccountsRequest, opts ...grpc.CallOption) (*operation.Operation, error)
	// Lists users for the specified federation.
	ListUserAccounts(ctx context.Context, in *ListFederatedUserAccountsRequest, opts ...grpc.CallOption) (*ListFederatedUserAccountsResponse, error)
	// Lists operations for the specified federation.
	ListOperations(ctx context.Context, in *ListFederationOperationsRequest, opts ...grpc.CallOption) (*ListFederationOperationsResponse, error)
}

type federationServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewFederationServiceClient(cc grpc.ClientConnInterface) FederationServiceClient {
	return &federationServiceClient{cc}
}

func (c *federationServiceClient) Get(ctx context.Context, in *GetFederationRequest, opts ...grpc.CallOption) (*Federation, error) {
	out := new(Federation)
	err := c.cc.Invoke(ctx, "/yandex.cloud.organizationmanager.v1.saml.FederationService/Get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *federationServiceClient) List(ctx context.Context, in *ListFederationsRequest, opts ...grpc.CallOption) (*ListFederationsResponse, error) {
	out := new(ListFederationsResponse)
	err := c.cc.Invoke(ctx, "/yandex.cloud.organizationmanager.v1.saml.FederationService/List", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *federationServiceClient) Create(ctx context.Context, in *CreateFederationRequest, opts ...grpc.CallOption) (*operation.Operation, error) {
	out := new(operation.Operation)
	err := c.cc.Invoke(ctx, "/yandex.cloud.organizationmanager.v1.saml.FederationService/Create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *federationServiceClient) Update(ctx context.Context, in *UpdateFederationRequest, opts ...grpc.CallOption) (*operation.Operation, error) {
	out := new(operation.Operation)
	err := c.cc.Invoke(ctx, "/yandex.cloud.organizationmanager.v1.saml.FederationService/Update", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *federationServiceClient) Delete(ctx context.Context, in *DeleteFederationRequest, opts ...grpc.CallOption) (*operation.Operation, error) {
	out := new(operation.Operation)
	err := c.cc.Invoke(ctx, "/yandex.cloud.organizationmanager.v1.saml.FederationService/Delete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *federationServiceClient) AddUserAccounts(ctx context.Context, in *AddFederatedUserAccountsRequest, opts ...grpc.CallOption) (*operation.Operation, error) {
	out := new(operation.Operation)
	err := c.cc.Invoke(ctx, "/yandex.cloud.organizationmanager.v1.saml.FederationService/AddUserAccounts", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *federationServiceClient) ListUserAccounts(ctx context.Context, in *ListFederatedUserAccountsRequest, opts ...grpc.CallOption) (*ListFederatedUserAccountsResponse, error) {
	out := new(ListFederatedUserAccountsResponse)
	err := c.cc.Invoke(ctx, "/yandex.cloud.organizationmanager.v1.saml.FederationService/ListUserAccounts", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *federationServiceClient) ListOperations(ctx context.Context, in *ListFederationOperationsRequest, opts ...grpc.CallOption) (*ListFederationOperationsResponse, error) {
	out := new(ListFederationOperationsResponse)
	err := c.cc.Invoke(ctx, "/yandex.cloud.organizationmanager.v1.saml.FederationService/ListOperations", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// FederationServiceServer is the server API for FederationService service.
// All implementations should embed UnimplementedFederationServiceServer
// for forward compatibility
type FederationServiceServer interface {
	// Returns the specified federation.
	//
	// To get the list of available federations, make a [List] request.
	Get(context.Context, *GetFederationRequest) (*Federation, error)
	// Retrieves the list of federations in the specified organization.
	List(context.Context, *ListFederationsRequest) (*ListFederationsResponse, error)
	// Creates a federation in the specified organization.
	Create(context.Context, *CreateFederationRequest) (*operation.Operation, error)
	// Updates the specified federation.
	Update(context.Context, *UpdateFederationRequest) (*operation.Operation, error)
	// Deletes the specified federation.
	Delete(context.Context, *DeleteFederationRequest) (*operation.Operation, error)
	// Adds users to the specified federation.
	AddUserAccounts(context.Context, *AddFederatedUserAccountsRequest) (*operation.Operation, error)
	// Lists users for the specified federation.
	ListUserAccounts(context.Context, *ListFederatedUserAccountsRequest) (*ListFederatedUserAccountsResponse, error)
	// Lists operations for the specified federation.
	ListOperations(context.Context, *ListFederationOperationsRequest) (*ListFederationOperationsResponse, error)
}

// UnimplementedFederationServiceServer should be embedded to have forward compatible implementations.
type UnimplementedFederationServiceServer struct {
}

func (UnimplementedFederationServiceServer) Get(context.Context, *GetFederationRequest) (*Federation, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (UnimplementedFederationServiceServer) List(context.Context, *ListFederationsRequest) (*ListFederationsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method List not implemented")
}
func (UnimplementedFederationServiceServer) Create(context.Context, *CreateFederationRequest) (*operation.Operation, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedFederationServiceServer) Update(context.Context, *UpdateFederationRequest) (*operation.Operation, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (UnimplementedFederationServiceServer) Delete(context.Context, *DeleteFederationRequest) (*operation.Operation, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (UnimplementedFederationServiceServer) AddUserAccounts(context.Context, *AddFederatedUserAccountsRequest) (*operation.Operation, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddUserAccounts not implemented")
}
func (UnimplementedFederationServiceServer) ListUserAccounts(context.Context, *ListFederatedUserAccountsRequest) (*ListFederatedUserAccountsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListUserAccounts not implemented")
}
func (UnimplementedFederationServiceServer) ListOperations(context.Context, *ListFederationOperationsRequest) (*ListFederationOperationsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListOperations not implemented")
}

// UnsafeFederationServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to FederationServiceServer will
// result in compilation errors.
type UnsafeFederationServiceServer interface {
	mustEmbedUnimplementedFederationServiceServer()
}

func RegisterFederationServiceServer(s grpc.ServiceRegistrar, srv FederationServiceServer) {
	s.RegisterService(&FederationService_ServiceDesc, srv)
}

func _FederationService_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetFederationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FederationServiceServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/yandex.cloud.organizationmanager.v1.saml.FederationService/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FederationServiceServer).Get(ctx, req.(*GetFederationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _FederationService_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListFederationsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FederationServiceServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/yandex.cloud.organizationmanager.v1.saml.FederationService/List",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FederationServiceServer).List(ctx, req.(*ListFederationsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _FederationService_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateFederationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FederationServiceServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/yandex.cloud.organizationmanager.v1.saml.FederationService/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FederationServiceServer).Create(ctx, req.(*CreateFederationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _FederationService_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateFederationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FederationServiceServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/yandex.cloud.organizationmanager.v1.saml.FederationService/Update",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FederationServiceServer).Update(ctx, req.(*UpdateFederationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _FederationService_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteFederationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FederationServiceServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/yandex.cloud.organizationmanager.v1.saml.FederationService/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FederationServiceServer).Delete(ctx, req.(*DeleteFederationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _FederationService_AddUserAccounts_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddFederatedUserAccountsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FederationServiceServer).AddUserAccounts(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/yandex.cloud.organizationmanager.v1.saml.FederationService/AddUserAccounts",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FederationServiceServer).AddUserAccounts(ctx, req.(*AddFederatedUserAccountsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _FederationService_ListUserAccounts_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListFederatedUserAccountsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FederationServiceServer).ListUserAccounts(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/yandex.cloud.organizationmanager.v1.saml.FederationService/ListUserAccounts",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FederationServiceServer).ListUserAccounts(ctx, req.(*ListFederatedUserAccountsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _FederationService_ListOperations_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListFederationOperationsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FederationServiceServer).ListOperations(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/yandex.cloud.organizationmanager.v1.saml.FederationService/ListOperations",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FederationServiceServer).ListOperations(ctx, req.(*ListFederationOperationsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// FederationService_ServiceDesc is the grpc.ServiceDesc for FederationService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var FederationService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "yandex.cloud.organizationmanager.v1.saml.FederationService",
	HandlerType: (*FederationServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Get",
			Handler:    _FederationService_Get_Handler,
		},
		{
			MethodName: "List",
			Handler:    _FederationService_List_Handler,
		},
		{
			MethodName: "Create",
			Handler:    _FederationService_Create_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _FederationService_Update_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _FederationService_Delete_Handler,
		},
		{
			MethodName: "AddUserAccounts",
			Handler:    _FederationService_AddUserAccounts_Handler,
		},
		{
			MethodName: "ListUserAccounts",
			Handler:    _FederationService_ListUserAccounts_Handler,
		},
		{
			MethodName: "ListOperations",
			Handler:    _FederationService_ListOperations_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "yandex/cloud/organizationmanager/v1/saml/federation_service.proto",
}
