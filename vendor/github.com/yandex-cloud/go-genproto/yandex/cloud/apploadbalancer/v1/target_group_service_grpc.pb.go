// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package apploadbalancer

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

// TargetGroupServiceClient is the client API for TargetGroupService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type TargetGroupServiceClient interface {
	// Returns the specified target group.
	//
	// To get the list of all available target groups, make a [List] request.
	Get(ctx context.Context, in *GetTargetGroupRequest, opts ...grpc.CallOption) (*TargetGroup, error)
	// Lists target groups in the specified folder.
	List(ctx context.Context, in *ListTargetGroupsRequest, opts ...grpc.CallOption) (*ListTargetGroupsResponse, error)
	// Creates a target group in the specified folder.
	Create(ctx context.Context, in *CreateTargetGroupRequest, opts ...grpc.CallOption) (*operation.Operation, error)
	// Updates the specified target group.
	Update(ctx context.Context, in *UpdateTargetGroupRequest, opts ...grpc.CallOption) (*operation.Operation, error)
	// Deletes the specified target group.
	Delete(ctx context.Context, in *DeleteTargetGroupRequest, opts ...grpc.CallOption) (*operation.Operation, error)
	// Adds targets to the specified target group.
	AddTargets(ctx context.Context, in *AddTargetsRequest, opts ...grpc.CallOption) (*operation.Operation, error)
	// Removes targets from the specified target group.
	RemoveTargets(ctx context.Context, in *RemoveTargetsRequest, opts ...grpc.CallOption) (*operation.Operation, error)
	// Lists operations for the specified target group.
	ListOperations(ctx context.Context, in *ListTargetGroupOperationsRequest, opts ...grpc.CallOption) (*ListTargetGroupOperationsResponse, error)
}

type targetGroupServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewTargetGroupServiceClient(cc grpc.ClientConnInterface) TargetGroupServiceClient {
	return &targetGroupServiceClient{cc}
}

func (c *targetGroupServiceClient) Get(ctx context.Context, in *GetTargetGroupRequest, opts ...grpc.CallOption) (*TargetGroup, error) {
	out := new(TargetGroup)
	err := c.cc.Invoke(ctx, "/yandex.cloud.apploadbalancer.v1.TargetGroupService/Get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *targetGroupServiceClient) List(ctx context.Context, in *ListTargetGroupsRequest, opts ...grpc.CallOption) (*ListTargetGroupsResponse, error) {
	out := new(ListTargetGroupsResponse)
	err := c.cc.Invoke(ctx, "/yandex.cloud.apploadbalancer.v1.TargetGroupService/List", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *targetGroupServiceClient) Create(ctx context.Context, in *CreateTargetGroupRequest, opts ...grpc.CallOption) (*operation.Operation, error) {
	out := new(operation.Operation)
	err := c.cc.Invoke(ctx, "/yandex.cloud.apploadbalancer.v1.TargetGroupService/Create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *targetGroupServiceClient) Update(ctx context.Context, in *UpdateTargetGroupRequest, opts ...grpc.CallOption) (*operation.Operation, error) {
	out := new(operation.Operation)
	err := c.cc.Invoke(ctx, "/yandex.cloud.apploadbalancer.v1.TargetGroupService/Update", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *targetGroupServiceClient) Delete(ctx context.Context, in *DeleteTargetGroupRequest, opts ...grpc.CallOption) (*operation.Operation, error) {
	out := new(operation.Operation)
	err := c.cc.Invoke(ctx, "/yandex.cloud.apploadbalancer.v1.TargetGroupService/Delete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *targetGroupServiceClient) AddTargets(ctx context.Context, in *AddTargetsRequest, opts ...grpc.CallOption) (*operation.Operation, error) {
	out := new(operation.Operation)
	err := c.cc.Invoke(ctx, "/yandex.cloud.apploadbalancer.v1.TargetGroupService/AddTargets", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *targetGroupServiceClient) RemoveTargets(ctx context.Context, in *RemoveTargetsRequest, opts ...grpc.CallOption) (*operation.Operation, error) {
	out := new(operation.Operation)
	err := c.cc.Invoke(ctx, "/yandex.cloud.apploadbalancer.v1.TargetGroupService/RemoveTargets", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *targetGroupServiceClient) ListOperations(ctx context.Context, in *ListTargetGroupOperationsRequest, opts ...grpc.CallOption) (*ListTargetGroupOperationsResponse, error) {
	out := new(ListTargetGroupOperationsResponse)
	err := c.cc.Invoke(ctx, "/yandex.cloud.apploadbalancer.v1.TargetGroupService/ListOperations", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TargetGroupServiceServer is the server API for TargetGroupService service.
// All implementations should embed UnimplementedTargetGroupServiceServer
// for forward compatibility
type TargetGroupServiceServer interface {
	// Returns the specified target group.
	//
	// To get the list of all available target groups, make a [List] request.
	Get(context.Context, *GetTargetGroupRequest) (*TargetGroup, error)
	// Lists target groups in the specified folder.
	List(context.Context, *ListTargetGroupsRequest) (*ListTargetGroupsResponse, error)
	// Creates a target group in the specified folder.
	Create(context.Context, *CreateTargetGroupRequest) (*operation.Operation, error)
	// Updates the specified target group.
	Update(context.Context, *UpdateTargetGroupRequest) (*operation.Operation, error)
	// Deletes the specified target group.
	Delete(context.Context, *DeleteTargetGroupRequest) (*operation.Operation, error)
	// Adds targets to the specified target group.
	AddTargets(context.Context, *AddTargetsRequest) (*operation.Operation, error)
	// Removes targets from the specified target group.
	RemoveTargets(context.Context, *RemoveTargetsRequest) (*operation.Operation, error)
	// Lists operations for the specified target group.
	ListOperations(context.Context, *ListTargetGroupOperationsRequest) (*ListTargetGroupOperationsResponse, error)
}

// UnimplementedTargetGroupServiceServer should be embedded to have forward compatible implementations.
type UnimplementedTargetGroupServiceServer struct {
}

func (UnimplementedTargetGroupServiceServer) Get(context.Context, *GetTargetGroupRequest) (*TargetGroup, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (UnimplementedTargetGroupServiceServer) List(context.Context, *ListTargetGroupsRequest) (*ListTargetGroupsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method List not implemented")
}
func (UnimplementedTargetGroupServiceServer) Create(context.Context, *CreateTargetGroupRequest) (*operation.Operation, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedTargetGroupServiceServer) Update(context.Context, *UpdateTargetGroupRequest) (*operation.Operation, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (UnimplementedTargetGroupServiceServer) Delete(context.Context, *DeleteTargetGroupRequest) (*operation.Operation, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (UnimplementedTargetGroupServiceServer) AddTargets(context.Context, *AddTargetsRequest) (*operation.Operation, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddTargets not implemented")
}
func (UnimplementedTargetGroupServiceServer) RemoveTargets(context.Context, *RemoveTargetsRequest) (*operation.Operation, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RemoveTargets not implemented")
}
func (UnimplementedTargetGroupServiceServer) ListOperations(context.Context, *ListTargetGroupOperationsRequest) (*ListTargetGroupOperationsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListOperations not implemented")
}

// UnsafeTargetGroupServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to TargetGroupServiceServer will
// result in compilation errors.
type UnsafeTargetGroupServiceServer interface {
	mustEmbedUnimplementedTargetGroupServiceServer()
}

func RegisterTargetGroupServiceServer(s grpc.ServiceRegistrar, srv TargetGroupServiceServer) {
	s.RegisterService(&TargetGroupService_ServiceDesc, srv)
}

func _TargetGroupService_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetTargetGroupRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TargetGroupServiceServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/yandex.cloud.apploadbalancer.v1.TargetGroupService/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TargetGroupServiceServer).Get(ctx, req.(*GetTargetGroupRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TargetGroupService_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListTargetGroupsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TargetGroupServiceServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/yandex.cloud.apploadbalancer.v1.TargetGroupService/List",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TargetGroupServiceServer).List(ctx, req.(*ListTargetGroupsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TargetGroupService_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateTargetGroupRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TargetGroupServiceServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/yandex.cloud.apploadbalancer.v1.TargetGroupService/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TargetGroupServiceServer).Create(ctx, req.(*CreateTargetGroupRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TargetGroupService_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateTargetGroupRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TargetGroupServiceServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/yandex.cloud.apploadbalancer.v1.TargetGroupService/Update",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TargetGroupServiceServer).Update(ctx, req.(*UpdateTargetGroupRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TargetGroupService_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteTargetGroupRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TargetGroupServiceServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/yandex.cloud.apploadbalancer.v1.TargetGroupService/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TargetGroupServiceServer).Delete(ctx, req.(*DeleteTargetGroupRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TargetGroupService_AddTargets_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddTargetsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TargetGroupServiceServer).AddTargets(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/yandex.cloud.apploadbalancer.v1.TargetGroupService/AddTargets",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TargetGroupServiceServer).AddTargets(ctx, req.(*AddTargetsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TargetGroupService_RemoveTargets_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RemoveTargetsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TargetGroupServiceServer).RemoveTargets(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/yandex.cloud.apploadbalancer.v1.TargetGroupService/RemoveTargets",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TargetGroupServiceServer).RemoveTargets(ctx, req.(*RemoveTargetsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TargetGroupService_ListOperations_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListTargetGroupOperationsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TargetGroupServiceServer).ListOperations(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/yandex.cloud.apploadbalancer.v1.TargetGroupService/ListOperations",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TargetGroupServiceServer).ListOperations(ctx, req.(*ListTargetGroupOperationsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// TargetGroupService_ServiceDesc is the grpc.ServiceDesc for TargetGroupService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var TargetGroupService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "yandex.cloud.apploadbalancer.v1.TargetGroupService",
	HandlerType: (*TargetGroupServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Get",
			Handler:    _TargetGroupService_Get_Handler,
		},
		{
			MethodName: "List",
			Handler:    _TargetGroupService_List_Handler,
		},
		{
			MethodName: "Create",
			Handler:    _TargetGroupService_Create_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _TargetGroupService_Update_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _TargetGroupService_Delete_Handler,
		},
		{
			MethodName: "AddTargets",
			Handler:    _TargetGroupService_AddTargets_Handler,
		},
		{
			MethodName: "RemoveTargets",
			Handler:    _TargetGroupService_RemoveTargets_Handler,
		},
		{
			MethodName: "ListOperations",
			Handler:    _TargetGroupService_ListOperations_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "yandex/cloud/apploadbalancer/v1/target_group_service.proto",
}
