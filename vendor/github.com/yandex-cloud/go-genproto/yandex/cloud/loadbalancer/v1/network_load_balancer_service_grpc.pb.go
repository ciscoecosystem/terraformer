// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package loadbalancer

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

// NetworkLoadBalancerServiceClient is the client API for NetworkLoadBalancerService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type NetworkLoadBalancerServiceClient interface {
	// Returns the specified NetworkLoadBalancer resource.
	//
	// Get the list of available NetworkLoadBalancer resources by making a [List] request.
	Get(ctx context.Context, in *GetNetworkLoadBalancerRequest, opts ...grpc.CallOption) (*NetworkLoadBalancer, error)
	// Retrieves the list of NetworkLoadBalancer resources in the specified folder.
	List(ctx context.Context, in *ListNetworkLoadBalancersRequest, opts ...grpc.CallOption) (*ListNetworkLoadBalancersResponse, error)
	// Creates a network load balancer in the specified folder using the data specified in the request.
	Create(ctx context.Context, in *CreateNetworkLoadBalancerRequest, opts ...grpc.CallOption) (*operation.Operation, error)
	// Updates the specified network load balancer.
	Update(ctx context.Context, in *UpdateNetworkLoadBalancerRequest, opts ...grpc.CallOption) (*operation.Operation, error)
	// Deletes the specified network load balancer.
	Delete(ctx context.Context, in *DeleteNetworkLoadBalancerRequest, opts ...grpc.CallOption) (*operation.Operation, error)
	// Starts load balancing and health checking with the specified network load balancer with specified settings.
	// Changes network load balancer status to `` ACTIVE ``.
	Start(ctx context.Context, in *StartNetworkLoadBalancerRequest, opts ...grpc.CallOption) (*operation.Operation, error)
	// Stops load balancing and health checking with the specified network load balancer.
	// Changes load balancer status to `` STOPPED ``.
	Stop(ctx context.Context, in *StopNetworkLoadBalancerRequest, opts ...grpc.CallOption) (*operation.Operation, error)
	// Attaches a target group to the specified network load balancer.
	AttachTargetGroup(ctx context.Context, in *AttachNetworkLoadBalancerTargetGroupRequest, opts ...grpc.CallOption) (*operation.Operation, error)
	// Detaches the target group from the specified network load balancer.
	DetachTargetGroup(ctx context.Context, in *DetachNetworkLoadBalancerTargetGroupRequest, opts ...grpc.CallOption) (*operation.Operation, error)
	// Gets states of target resources in the attached target group.
	GetTargetStates(ctx context.Context, in *GetTargetStatesRequest, opts ...grpc.CallOption) (*GetTargetStatesResponse, error)
	// Adds a listener to the specified network load balancer.
	AddListener(ctx context.Context, in *AddNetworkLoadBalancerListenerRequest, opts ...grpc.CallOption) (*operation.Operation, error)
	// Removes the listener from the specified network load balancer.
	RemoveListener(ctx context.Context, in *RemoveNetworkLoadBalancerListenerRequest, opts ...grpc.CallOption) (*operation.Operation, error)
	// Lists operations for the specified network load balancer.
	ListOperations(ctx context.Context, in *ListNetworkLoadBalancerOperationsRequest, opts ...grpc.CallOption) (*ListNetworkLoadBalancerOperationsResponse, error)
}

type networkLoadBalancerServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewNetworkLoadBalancerServiceClient(cc grpc.ClientConnInterface) NetworkLoadBalancerServiceClient {
	return &networkLoadBalancerServiceClient{cc}
}

func (c *networkLoadBalancerServiceClient) Get(ctx context.Context, in *GetNetworkLoadBalancerRequest, opts ...grpc.CallOption) (*NetworkLoadBalancer, error) {
	out := new(NetworkLoadBalancer)
	err := c.cc.Invoke(ctx, "/yandex.cloud.loadbalancer.v1.NetworkLoadBalancerService/Get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *networkLoadBalancerServiceClient) List(ctx context.Context, in *ListNetworkLoadBalancersRequest, opts ...grpc.CallOption) (*ListNetworkLoadBalancersResponse, error) {
	out := new(ListNetworkLoadBalancersResponse)
	err := c.cc.Invoke(ctx, "/yandex.cloud.loadbalancer.v1.NetworkLoadBalancerService/List", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *networkLoadBalancerServiceClient) Create(ctx context.Context, in *CreateNetworkLoadBalancerRequest, opts ...grpc.CallOption) (*operation.Operation, error) {
	out := new(operation.Operation)
	err := c.cc.Invoke(ctx, "/yandex.cloud.loadbalancer.v1.NetworkLoadBalancerService/Create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *networkLoadBalancerServiceClient) Update(ctx context.Context, in *UpdateNetworkLoadBalancerRequest, opts ...grpc.CallOption) (*operation.Operation, error) {
	out := new(operation.Operation)
	err := c.cc.Invoke(ctx, "/yandex.cloud.loadbalancer.v1.NetworkLoadBalancerService/Update", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *networkLoadBalancerServiceClient) Delete(ctx context.Context, in *DeleteNetworkLoadBalancerRequest, opts ...grpc.CallOption) (*operation.Operation, error) {
	out := new(operation.Operation)
	err := c.cc.Invoke(ctx, "/yandex.cloud.loadbalancer.v1.NetworkLoadBalancerService/Delete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *networkLoadBalancerServiceClient) Start(ctx context.Context, in *StartNetworkLoadBalancerRequest, opts ...grpc.CallOption) (*operation.Operation, error) {
	out := new(operation.Operation)
	err := c.cc.Invoke(ctx, "/yandex.cloud.loadbalancer.v1.NetworkLoadBalancerService/Start", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *networkLoadBalancerServiceClient) Stop(ctx context.Context, in *StopNetworkLoadBalancerRequest, opts ...grpc.CallOption) (*operation.Operation, error) {
	out := new(operation.Operation)
	err := c.cc.Invoke(ctx, "/yandex.cloud.loadbalancer.v1.NetworkLoadBalancerService/Stop", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *networkLoadBalancerServiceClient) AttachTargetGroup(ctx context.Context, in *AttachNetworkLoadBalancerTargetGroupRequest, opts ...grpc.CallOption) (*operation.Operation, error) {
	out := new(operation.Operation)
	err := c.cc.Invoke(ctx, "/yandex.cloud.loadbalancer.v1.NetworkLoadBalancerService/AttachTargetGroup", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *networkLoadBalancerServiceClient) DetachTargetGroup(ctx context.Context, in *DetachNetworkLoadBalancerTargetGroupRequest, opts ...grpc.CallOption) (*operation.Operation, error) {
	out := new(operation.Operation)
	err := c.cc.Invoke(ctx, "/yandex.cloud.loadbalancer.v1.NetworkLoadBalancerService/DetachTargetGroup", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *networkLoadBalancerServiceClient) GetTargetStates(ctx context.Context, in *GetTargetStatesRequest, opts ...grpc.CallOption) (*GetTargetStatesResponse, error) {
	out := new(GetTargetStatesResponse)
	err := c.cc.Invoke(ctx, "/yandex.cloud.loadbalancer.v1.NetworkLoadBalancerService/GetTargetStates", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *networkLoadBalancerServiceClient) AddListener(ctx context.Context, in *AddNetworkLoadBalancerListenerRequest, opts ...grpc.CallOption) (*operation.Operation, error) {
	out := new(operation.Operation)
	err := c.cc.Invoke(ctx, "/yandex.cloud.loadbalancer.v1.NetworkLoadBalancerService/AddListener", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *networkLoadBalancerServiceClient) RemoveListener(ctx context.Context, in *RemoveNetworkLoadBalancerListenerRequest, opts ...grpc.CallOption) (*operation.Operation, error) {
	out := new(operation.Operation)
	err := c.cc.Invoke(ctx, "/yandex.cloud.loadbalancer.v1.NetworkLoadBalancerService/RemoveListener", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *networkLoadBalancerServiceClient) ListOperations(ctx context.Context, in *ListNetworkLoadBalancerOperationsRequest, opts ...grpc.CallOption) (*ListNetworkLoadBalancerOperationsResponse, error) {
	out := new(ListNetworkLoadBalancerOperationsResponse)
	err := c.cc.Invoke(ctx, "/yandex.cloud.loadbalancer.v1.NetworkLoadBalancerService/ListOperations", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// NetworkLoadBalancerServiceServer is the server API for NetworkLoadBalancerService service.
// All implementations should embed UnimplementedNetworkLoadBalancerServiceServer
// for forward compatibility
type NetworkLoadBalancerServiceServer interface {
	// Returns the specified NetworkLoadBalancer resource.
	//
	// Get the list of available NetworkLoadBalancer resources by making a [List] request.
	Get(context.Context, *GetNetworkLoadBalancerRequest) (*NetworkLoadBalancer, error)
	// Retrieves the list of NetworkLoadBalancer resources in the specified folder.
	List(context.Context, *ListNetworkLoadBalancersRequest) (*ListNetworkLoadBalancersResponse, error)
	// Creates a network load balancer in the specified folder using the data specified in the request.
	Create(context.Context, *CreateNetworkLoadBalancerRequest) (*operation.Operation, error)
	// Updates the specified network load balancer.
	Update(context.Context, *UpdateNetworkLoadBalancerRequest) (*operation.Operation, error)
	// Deletes the specified network load balancer.
	Delete(context.Context, *DeleteNetworkLoadBalancerRequest) (*operation.Operation, error)
	// Starts load balancing and health checking with the specified network load balancer with specified settings.
	// Changes network load balancer status to `` ACTIVE ``.
	Start(context.Context, *StartNetworkLoadBalancerRequest) (*operation.Operation, error)
	// Stops load balancing and health checking with the specified network load balancer.
	// Changes load balancer status to `` STOPPED ``.
	Stop(context.Context, *StopNetworkLoadBalancerRequest) (*operation.Operation, error)
	// Attaches a target group to the specified network load balancer.
	AttachTargetGroup(context.Context, *AttachNetworkLoadBalancerTargetGroupRequest) (*operation.Operation, error)
	// Detaches the target group from the specified network load balancer.
	DetachTargetGroup(context.Context, *DetachNetworkLoadBalancerTargetGroupRequest) (*operation.Operation, error)
	// Gets states of target resources in the attached target group.
	GetTargetStates(context.Context, *GetTargetStatesRequest) (*GetTargetStatesResponse, error)
	// Adds a listener to the specified network load balancer.
	AddListener(context.Context, *AddNetworkLoadBalancerListenerRequest) (*operation.Operation, error)
	// Removes the listener from the specified network load balancer.
	RemoveListener(context.Context, *RemoveNetworkLoadBalancerListenerRequest) (*operation.Operation, error)
	// Lists operations for the specified network load balancer.
	ListOperations(context.Context, *ListNetworkLoadBalancerOperationsRequest) (*ListNetworkLoadBalancerOperationsResponse, error)
}

// UnimplementedNetworkLoadBalancerServiceServer should be embedded to have forward compatible implementations.
type UnimplementedNetworkLoadBalancerServiceServer struct {
}

func (UnimplementedNetworkLoadBalancerServiceServer) Get(context.Context, *GetNetworkLoadBalancerRequest) (*NetworkLoadBalancer, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (UnimplementedNetworkLoadBalancerServiceServer) List(context.Context, *ListNetworkLoadBalancersRequest) (*ListNetworkLoadBalancersResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method List not implemented")
}
func (UnimplementedNetworkLoadBalancerServiceServer) Create(context.Context, *CreateNetworkLoadBalancerRequest) (*operation.Operation, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedNetworkLoadBalancerServiceServer) Update(context.Context, *UpdateNetworkLoadBalancerRequest) (*operation.Operation, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (UnimplementedNetworkLoadBalancerServiceServer) Delete(context.Context, *DeleteNetworkLoadBalancerRequest) (*operation.Operation, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (UnimplementedNetworkLoadBalancerServiceServer) Start(context.Context, *StartNetworkLoadBalancerRequest) (*operation.Operation, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Start not implemented")
}
func (UnimplementedNetworkLoadBalancerServiceServer) Stop(context.Context, *StopNetworkLoadBalancerRequest) (*operation.Operation, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Stop not implemented")
}
func (UnimplementedNetworkLoadBalancerServiceServer) AttachTargetGroup(context.Context, *AttachNetworkLoadBalancerTargetGroupRequest) (*operation.Operation, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AttachTargetGroup not implemented")
}
func (UnimplementedNetworkLoadBalancerServiceServer) DetachTargetGroup(context.Context, *DetachNetworkLoadBalancerTargetGroupRequest) (*operation.Operation, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DetachTargetGroup not implemented")
}
func (UnimplementedNetworkLoadBalancerServiceServer) GetTargetStates(context.Context, *GetTargetStatesRequest) (*GetTargetStatesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTargetStates not implemented")
}
func (UnimplementedNetworkLoadBalancerServiceServer) AddListener(context.Context, *AddNetworkLoadBalancerListenerRequest) (*operation.Operation, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddListener not implemented")
}
func (UnimplementedNetworkLoadBalancerServiceServer) RemoveListener(context.Context, *RemoveNetworkLoadBalancerListenerRequest) (*operation.Operation, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RemoveListener not implemented")
}
func (UnimplementedNetworkLoadBalancerServiceServer) ListOperations(context.Context, *ListNetworkLoadBalancerOperationsRequest) (*ListNetworkLoadBalancerOperationsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListOperations not implemented")
}

// UnsafeNetworkLoadBalancerServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to NetworkLoadBalancerServiceServer will
// result in compilation errors.
type UnsafeNetworkLoadBalancerServiceServer interface {
	mustEmbedUnimplementedNetworkLoadBalancerServiceServer()
}

func RegisterNetworkLoadBalancerServiceServer(s grpc.ServiceRegistrar, srv NetworkLoadBalancerServiceServer) {
	s.RegisterService(&NetworkLoadBalancerService_ServiceDesc, srv)
}

func _NetworkLoadBalancerService_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetNetworkLoadBalancerRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NetworkLoadBalancerServiceServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/yandex.cloud.loadbalancer.v1.NetworkLoadBalancerService/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NetworkLoadBalancerServiceServer).Get(ctx, req.(*GetNetworkLoadBalancerRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _NetworkLoadBalancerService_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListNetworkLoadBalancersRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NetworkLoadBalancerServiceServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/yandex.cloud.loadbalancer.v1.NetworkLoadBalancerService/List",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NetworkLoadBalancerServiceServer).List(ctx, req.(*ListNetworkLoadBalancersRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _NetworkLoadBalancerService_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateNetworkLoadBalancerRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NetworkLoadBalancerServiceServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/yandex.cloud.loadbalancer.v1.NetworkLoadBalancerService/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NetworkLoadBalancerServiceServer).Create(ctx, req.(*CreateNetworkLoadBalancerRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _NetworkLoadBalancerService_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateNetworkLoadBalancerRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NetworkLoadBalancerServiceServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/yandex.cloud.loadbalancer.v1.NetworkLoadBalancerService/Update",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NetworkLoadBalancerServiceServer).Update(ctx, req.(*UpdateNetworkLoadBalancerRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _NetworkLoadBalancerService_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteNetworkLoadBalancerRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NetworkLoadBalancerServiceServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/yandex.cloud.loadbalancer.v1.NetworkLoadBalancerService/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NetworkLoadBalancerServiceServer).Delete(ctx, req.(*DeleteNetworkLoadBalancerRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _NetworkLoadBalancerService_Start_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StartNetworkLoadBalancerRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NetworkLoadBalancerServiceServer).Start(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/yandex.cloud.loadbalancer.v1.NetworkLoadBalancerService/Start",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NetworkLoadBalancerServiceServer).Start(ctx, req.(*StartNetworkLoadBalancerRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _NetworkLoadBalancerService_Stop_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StopNetworkLoadBalancerRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NetworkLoadBalancerServiceServer).Stop(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/yandex.cloud.loadbalancer.v1.NetworkLoadBalancerService/Stop",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NetworkLoadBalancerServiceServer).Stop(ctx, req.(*StopNetworkLoadBalancerRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _NetworkLoadBalancerService_AttachTargetGroup_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AttachNetworkLoadBalancerTargetGroupRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NetworkLoadBalancerServiceServer).AttachTargetGroup(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/yandex.cloud.loadbalancer.v1.NetworkLoadBalancerService/AttachTargetGroup",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NetworkLoadBalancerServiceServer).AttachTargetGroup(ctx, req.(*AttachNetworkLoadBalancerTargetGroupRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _NetworkLoadBalancerService_DetachTargetGroup_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DetachNetworkLoadBalancerTargetGroupRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NetworkLoadBalancerServiceServer).DetachTargetGroup(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/yandex.cloud.loadbalancer.v1.NetworkLoadBalancerService/DetachTargetGroup",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NetworkLoadBalancerServiceServer).DetachTargetGroup(ctx, req.(*DetachNetworkLoadBalancerTargetGroupRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _NetworkLoadBalancerService_GetTargetStates_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetTargetStatesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NetworkLoadBalancerServiceServer).GetTargetStates(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/yandex.cloud.loadbalancer.v1.NetworkLoadBalancerService/GetTargetStates",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NetworkLoadBalancerServiceServer).GetTargetStates(ctx, req.(*GetTargetStatesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _NetworkLoadBalancerService_AddListener_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddNetworkLoadBalancerListenerRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NetworkLoadBalancerServiceServer).AddListener(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/yandex.cloud.loadbalancer.v1.NetworkLoadBalancerService/AddListener",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NetworkLoadBalancerServiceServer).AddListener(ctx, req.(*AddNetworkLoadBalancerListenerRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _NetworkLoadBalancerService_RemoveListener_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RemoveNetworkLoadBalancerListenerRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NetworkLoadBalancerServiceServer).RemoveListener(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/yandex.cloud.loadbalancer.v1.NetworkLoadBalancerService/RemoveListener",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NetworkLoadBalancerServiceServer).RemoveListener(ctx, req.(*RemoveNetworkLoadBalancerListenerRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _NetworkLoadBalancerService_ListOperations_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListNetworkLoadBalancerOperationsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NetworkLoadBalancerServiceServer).ListOperations(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/yandex.cloud.loadbalancer.v1.NetworkLoadBalancerService/ListOperations",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NetworkLoadBalancerServiceServer).ListOperations(ctx, req.(*ListNetworkLoadBalancerOperationsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// NetworkLoadBalancerService_ServiceDesc is the grpc.ServiceDesc for NetworkLoadBalancerService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var NetworkLoadBalancerService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "yandex.cloud.loadbalancer.v1.NetworkLoadBalancerService",
	HandlerType: (*NetworkLoadBalancerServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Get",
			Handler:    _NetworkLoadBalancerService_Get_Handler,
		},
		{
			MethodName: "List",
			Handler:    _NetworkLoadBalancerService_List_Handler,
		},
		{
			MethodName: "Create",
			Handler:    _NetworkLoadBalancerService_Create_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _NetworkLoadBalancerService_Update_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _NetworkLoadBalancerService_Delete_Handler,
		},
		{
			MethodName: "Start",
			Handler:    _NetworkLoadBalancerService_Start_Handler,
		},
		{
			MethodName: "Stop",
			Handler:    _NetworkLoadBalancerService_Stop_Handler,
		},
		{
			MethodName: "AttachTargetGroup",
			Handler:    _NetworkLoadBalancerService_AttachTargetGroup_Handler,
		},
		{
			MethodName: "DetachTargetGroup",
			Handler:    _NetworkLoadBalancerService_DetachTargetGroup_Handler,
		},
		{
			MethodName: "GetTargetStates",
			Handler:    _NetworkLoadBalancerService_GetTargetStates_Handler,
		},
		{
			MethodName: "AddListener",
			Handler:    _NetworkLoadBalancerService_AddListener_Handler,
		},
		{
			MethodName: "RemoveListener",
			Handler:    _NetworkLoadBalancerService_RemoveListener_Handler,
		},
		{
			MethodName: "ListOperations",
			Handler:    _NetworkLoadBalancerService_ListOperations_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "yandex/cloud/loadbalancer/v1/network_load_balancer_service.proto",
}
