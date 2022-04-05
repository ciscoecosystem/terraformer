// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package compute

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

// DiskPlacementGroupServiceClient is the client API for DiskPlacementGroupService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type DiskPlacementGroupServiceClient interface {
	// Returns the specified placement group.
	Get(ctx context.Context, in *GetDiskPlacementGroupRequest, opts ...grpc.CallOption) (*DiskPlacementGroup, error)
	// Retrieves the list of placement groups in the specified folder.
	List(ctx context.Context, in *ListDiskPlacementGroupsRequest, opts ...grpc.CallOption) (*ListDiskPlacementGroupsResponse, error)
	// Creates a placement group in the specified folder.
	Create(ctx context.Context, in *CreateDiskPlacementGroupRequest, opts ...grpc.CallOption) (*operation.Operation, error)
	// Updates the specified placement group.
	Update(ctx context.Context, in *UpdateDiskPlacementGroupRequest, opts ...grpc.CallOption) (*operation.Operation, error)
	// Deletes the specified placement group.
	Delete(ctx context.Context, in *DeleteDiskPlacementGroupRequest, opts ...grpc.CallOption) (*operation.Operation, error)
	// Lists disks for the specified placement group.
	ListDisks(ctx context.Context, in *ListDiskPlacementGroupDisksRequest, opts ...grpc.CallOption) (*ListDiskPlacementGroupDisksResponse, error)
	// Lists operations for the specified placement group.
	ListOperations(ctx context.Context, in *ListDiskPlacementGroupOperationsRequest, opts ...grpc.CallOption) (*ListDiskPlacementGroupOperationsResponse, error)
}

type diskPlacementGroupServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewDiskPlacementGroupServiceClient(cc grpc.ClientConnInterface) DiskPlacementGroupServiceClient {
	return &diskPlacementGroupServiceClient{cc}
}

func (c *diskPlacementGroupServiceClient) Get(ctx context.Context, in *GetDiskPlacementGroupRequest, opts ...grpc.CallOption) (*DiskPlacementGroup, error) {
	out := new(DiskPlacementGroup)
	err := c.cc.Invoke(ctx, "/yandex.cloud.compute.v1.DiskPlacementGroupService/Get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *diskPlacementGroupServiceClient) List(ctx context.Context, in *ListDiskPlacementGroupsRequest, opts ...grpc.CallOption) (*ListDiskPlacementGroupsResponse, error) {
	out := new(ListDiskPlacementGroupsResponse)
	err := c.cc.Invoke(ctx, "/yandex.cloud.compute.v1.DiskPlacementGroupService/List", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *diskPlacementGroupServiceClient) Create(ctx context.Context, in *CreateDiskPlacementGroupRequest, opts ...grpc.CallOption) (*operation.Operation, error) {
	out := new(operation.Operation)
	err := c.cc.Invoke(ctx, "/yandex.cloud.compute.v1.DiskPlacementGroupService/Create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *diskPlacementGroupServiceClient) Update(ctx context.Context, in *UpdateDiskPlacementGroupRequest, opts ...grpc.CallOption) (*operation.Operation, error) {
	out := new(operation.Operation)
	err := c.cc.Invoke(ctx, "/yandex.cloud.compute.v1.DiskPlacementGroupService/Update", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *diskPlacementGroupServiceClient) Delete(ctx context.Context, in *DeleteDiskPlacementGroupRequest, opts ...grpc.CallOption) (*operation.Operation, error) {
	out := new(operation.Operation)
	err := c.cc.Invoke(ctx, "/yandex.cloud.compute.v1.DiskPlacementGroupService/Delete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *diskPlacementGroupServiceClient) ListDisks(ctx context.Context, in *ListDiskPlacementGroupDisksRequest, opts ...grpc.CallOption) (*ListDiskPlacementGroupDisksResponse, error) {
	out := new(ListDiskPlacementGroupDisksResponse)
	err := c.cc.Invoke(ctx, "/yandex.cloud.compute.v1.DiskPlacementGroupService/ListDisks", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *diskPlacementGroupServiceClient) ListOperations(ctx context.Context, in *ListDiskPlacementGroupOperationsRequest, opts ...grpc.CallOption) (*ListDiskPlacementGroupOperationsResponse, error) {
	out := new(ListDiskPlacementGroupOperationsResponse)
	err := c.cc.Invoke(ctx, "/yandex.cloud.compute.v1.DiskPlacementGroupService/ListOperations", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DiskPlacementGroupServiceServer is the server API for DiskPlacementGroupService service.
// All implementations should embed UnimplementedDiskPlacementGroupServiceServer
// for forward compatibility
type DiskPlacementGroupServiceServer interface {
	// Returns the specified placement group.
	Get(context.Context, *GetDiskPlacementGroupRequest) (*DiskPlacementGroup, error)
	// Retrieves the list of placement groups in the specified folder.
	List(context.Context, *ListDiskPlacementGroupsRequest) (*ListDiskPlacementGroupsResponse, error)
	// Creates a placement group in the specified folder.
	Create(context.Context, *CreateDiskPlacementGroupRequest) (*operation.Operation, error)
	// Updates the specified placement group.
	Update(context.Context, *UpdateDiskPlacementGroupRequest) (*operation.Operation, error)
	// Deletes the specified placement group.
	Delete(context.Context, *DeleteDiskPlacementGroupRequest) (*operation.Operation, error)
	// Lists disks for the specified placement group.
	ListDisks(context.Context, *ListDiskPlacementGroupDisksRequest) (*ListDiskPlacementGroupDisksResponse, error)
	// Lists operations for the specified placement group.
	ListOperations(context.Context, *ListDiskPlacementGroupOperationsRequest) (*ListDiskPlacementGroupOperationsResponse, error)
}

// UnimplementedDiskPlacementGroupServiceServer should be embedded to have forward compatible implementations.
type UnimplementedDiskPlacementGroupServiceServer struct {
}

func (UnimplementedDiskPlacementGroupServiceServer) Get(context.Context, *GetDiskPlacementGroupRequest) (*DiskPlacementGroup, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (UnimplementedDiskPlacementGroupServiceServer) List(context.Context, *ListDiskPlacementGroupsRequest) (*ListDiskPlacementGroupsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method List not implemented")
}
func (UnimplementedDiskPlacementGroupServiceServer) Create(context.Context, *CreateDiskPlacementGroupRequest) (*operation.Operation, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedDiskPlacementGroupServiceServer) Update(context.Context, *UpdateDiskPlacementGroupRequest) (*operation.Operation, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (UnimplementedDiskPlacementGroupServiceServer) Delete(context.Context, *DeleteDiskPlacementGroupRequest) (*operation.Operation, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (UnimplementedDiskPlacementGroupServiceServer) ListDisks(context.Context, *ListDiskPlacementGroupDisksRequest) (*ListDiskPlacementGroupDisksResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListDisks not implemented")
}
func (UnimplementedDiskPlacementGroupServiceServer) ListOperations(context.Context, *ListDiskPlacementGroupOperationsRequest) (*ListDiskPlacementGroupOperationsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListOperations not implemented")
}

// UnsafeDiskPlacementGroupServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to DiskPlacementGroupServiceServer will
// result in compilation errors.
type UnsafeDiskPlacementGroupServiceServer interface {
	mustEmbedUnimplementedDiskPlacementGroupServiceServer()
}

func RegisterDiskPlacementGroupServiceServer(s grpc.ServiceRegistrar, srv DiskPlacementGroupServiceServer) {
	s.RegisterService(&DiskPlacementGroupService_ServiceDesc, srv)
}

func _DiskPlacementGroupService_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetDiskPlacementGroupRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DiskPlacementGroupServiceServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/yandex.cloud.compute.v1.DiskPlacementGroupService/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DiskPlacementGroupServiceServer).Get(ctx, req.(*GetDiskPlacementGroupRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DiskPlacementGroupService_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListDiskPlacementGroupsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DiskPlacementGroupServiceServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/yandex.cloud.compute.v1.DiskPlacementGroupService/List",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DiskPlacementGroupServiceServer).List(ctx, req.(*ListDiskPlacementGroupsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DiskPlacementGroupService_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateDiskPlacementGroupRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DiskPlacementGroupServiceServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/yandex.cloud.compute.v1.DiskPlacementGroupService/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DiskPlacementGroupServiceServer).Create(ctx, req.(*CreateDiskPlacementGroupRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DiskPlacementGroupService_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateDiskPlacementGroupRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DiskPlacementGroupServiceServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/yandex.cloud.compute.v1.DiskPlacementGroupService/Update",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DiskPlacementGroupServiceServer).Update(ctx, req.(*UpdateDiskPlacementGroupRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DiskPlacementGroupService_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteDiskPlacementGroupRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DiskPlacementGroupServiceServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/yandex.cloud.compute.v1.DiskPlacementGroupService/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DiskPlacementGroupServiceServer).Delete(ctx, req.(*DeleteDiskPlacementGroupRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DiskPlacementGroupService_ListDisks_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListDiskPlacementGroupDisksRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DiskPlacementGroupServiceServer).ListDisks(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/yandex.cloud.compute.v1.DiskPlacementGroupService/ListDisks",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DiskPlacementGroupServiceServer).ListDisks(ctx, req.(*ListDiskPlacementGroupDisksRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DiskPlacementGroupService_ListOperations_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListDiskPlacementGroupOperationsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DiskPlacementGroupServiceServer).ListOperations(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/yandex.cloud.compute.v1.DiskPlacementGroupService/ListOperations",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DiskPlacementGroupServiceServer).ListOperations(ctx, req.(*ListDiskPlacementGroupOperationsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// DiskPlacementGroupService_ServiceDesc is the grpc.ServiceDesc for DiskPlacementGroupService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var DiskPlacementGroupService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "yandex.cloud.compute.v1.DiskPlacementGroupService",
	HandlerType: (*DiskPlacementGroupServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Get",
			Handler:    _DiskPlacementGroupService_Get_Handler,
		},
		{
			MethodName: "List",
			Handler:    _DiskPlacementGroupService_List_Handler,
		},
		{
			MethodName: "Create",
			Handler:    _DiskPlacementGroupService_Create_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _DiskPlacementGroupService_Update_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _DiskPlacementGroupService_Delete_Handler,
		},
		{
			MethodName: "ListDisks",
			Handler:    _DiskPlacementGroupService_ListDisks_Handler,
		},
		{
			MethodName: "ListOperations",
			Handler:    _DiskPlacementGroupService_ListOperations_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "yandex/cloud/compute/v1/disk_placement_group_service.proto",
}
