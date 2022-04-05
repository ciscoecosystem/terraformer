// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.17.3
// source: yandex/cloud/vpc/v1/route_table.proto

package vpc

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// A RouteTable resource. For more information, see [Static Routes](/docs/vpc/concepts/static-routes).
type RouteTable struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// ID of the route table.
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// ID of the folder that the route table belongs to.
	FolderId string `protobuf:"bytes,2,opt,name=folder_id,json=folderId,proto3" json:"folder_id,omitempty"`
	// Creation timestamp in [RFC3339](https://www.ietf.org/rfc/rfc3339.txt) text format.
	CreatedAt *timestamppb.Timestamp `protobuf:"bytes,3,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	// Name of the route table. The name is unique within the project. 3-63 characters long.
	Name string `protobuf:"bytes,4,opt,name=name,proto3" json:"name,omitempty"`
	// Optional description of the route table. 0-256 characters long.
	Description string `protobuf:"bytes,5,opt,name=description,proto3" json:"description,omitempty"`
	// Resource labels as `` key:value `` pairs. Maximum of 64 per resource.
	Labels map[string]string `protobuf:"bytes,6,rep,name=labels,proto3" json:"labels,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// ID of the network the route table belongs to.
	NetworkId string `protobuf:"bytes,7,opt,name=network_id,json=networkId,proto3" json:"network_id,omitempty"`
	// List of static routes.
	StaticRoutes []*StaticRoute `protobuf:"bytes,8,rep,name=static_routes,json=staticRoutes,proto3" json:"static_routes,omitempty"`
}

func (x *RouteTable) Reset() {
	*x = RouteTable{}
	if protoimpl.UnsafeEnabled {
		mi := &file_yandex_cloud_vpc_v1_route_table_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RouteTable) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RouteTable) ProtoMessage() {}

func (x *RouteTable) ProtoReflect() protoreflect.Message {
	mi := &file_yandex_cloud_vpc_v1_route_table_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RouteTable.ProtoReflect.Descriptor instead.
func (*RouteTable) Descriptor() ([]byte, []int) {
	return file_yandex_cloud_vpc_v1_route_table_proto_rawDescGZIP(), []int{0}
}

func (x *RouteTable) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *RouteTable) GetFolderId() string {
	if x != nil {
		return x.FolderId
	}
	return ""
}

func (x *RouteTable) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

func (x *RouteTable) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *RouteTable) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *RouteTable) GetLabels() map[string]string {
	if x != nil {
		return x.Labels
	}
	return nil
}

func (x *RouteTable) GetNetworkId() string {
	if x != nil {
		return x.NetworkId
	}
	return ""
}

func (x *RouteTable) GetStaticRoutes() []*StaticRoute {
	if x != nil {
		return x.StaticRoutes
	}
	return nil
}

// A StaticRoute resource. For more information, see [Static Routes](/docs/vpc/concepts/static-routes).
type StaticRoute struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Destination:
	//	*StaticRoute_DestinationPrefix
	Destination isStaticRoute_Destination `protobuf_oneof:"destination"`
	// Types that are assignable to NextHop:
	//	*StaticRoute_NextHopAddress
	NextHop isStaticRoute_NextHop `protobuf_oneof:"next_hop"`
	// Resource labels as `` key:value `` pairs. Maximum of 64 per resource.
	Labels map[string]string `protobuf:"bytes,3,rep,name=labels,proto3" json:"labels,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *StaticRoute) Reset() {
	*x = StaticRoute{}
	if protoimpl.UnsafeEnabled {
		mi := &file_yandex_cloud_vpc_v1_route_table_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StaticRoute) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StaticRoute) ProtoMessage() {}

func (x *StaticRoute) ProtoReflect() protoreflect.Message {
	mi := &file_yandex_cloud_vpc_v1_route_table_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StaticRoute.ProtoReflect.Descriptor instead.
func (*StaticRoute) Descriptor() ([]byte, []int) {
	return file_yandex_cloud_vpc_v1_route_table_proto_rawDescGZIP(), []int{1}
}

func (m *StaticRoute) GetDestination() isStaticRoute_Destination {
	if m != nil {
		return m.Destination
	}
	return nil
}

func (x *StaticRoute) GetDestinationPrefix() string {
	if x, ok := x.GetDestination().(*StaticRoute_DestinationPrefix); ok {
		return x.DestinationPrefix
	}
	return ""
}

func (m *StaticRoute) GetNextHop() isStaticRoute_NextHop {
	if m != nil {
		return m.NextHop
	}
	return nil
}

func (x *StaticRoute) GetNextHopAddress() string {
	if x, ok := x.GetNextHop().(*StaticRoute_NextHopAddress); ok {
		return x.NextHopAddress
	}
	return ""
}

func (x *StaticRoute) GetLabels() map[string]string {
	if x != nil {
		return x.Labels
	}
	return nil
}

type isStaticRoute_Destination interface {
	isStaticRoute_Destination()
}

type StaticRoute_DestinationPrefix struct {
	// Destination subnet in CIDR notation
	DestinationPrefix string `protobuf:"bytes,1,opt,name=destination_prefix,json=destinationPrefix,proto3,oneof"`
}

func (*StaticRoute_DestinationPrefix) isStaticRoute_Destination() {}

type isStaticRoute_NextHop interface {
	isStaticRoute_NextHop()
}

type StaticRoute_NextHopAddress struct {
	// Next hop IP address
	NextHopAddress string `protobuf:"bytes,2,opt,name=next_hop_address,json=nextHopAddress,proto3,oneof"`
}

func (*StaticRoute_NextHopAddress) isStaticRoute_NextHop() {}

var File_yandex_cloud_vpc_v1_route_table_proto protoreflect.FileDescriptor

var file_yandex_cloud_vpc_v1_route_table_proto_rawDesc = []byte{
	0x0a, 0x25, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x76,
	0x70, 0x63, 0x2f, 0x76, 0x31, 0x2f, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x5f, 0x74, 0x61, 0x62, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x13, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2e,
	0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x76, 0x70, 0x63, 0x2e, 0x76, 0x31, 0x1a, 0x1f, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69,
	0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x90, 0x03,
	0x0a, 0x0a, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x54, 0x61, 0x62, 0x6c, 0x65, 0x12, 0x0e, 0x0a, 0x02,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1b, 0x0a, 0x09,
	0x66, 0x6f, 0x6c, 0x64, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x66, 0x6f, 0x6c, 0x64, 0x65, 0x72, 0x49, 0x64, 0x12, 0x39, 0x0a, 0x0a, 0x63, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x64, 0x41, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63,
	0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64,
	0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x43, 0x0a, 0x06, 0x6c, 0x61,
	0x62, 0x65, 0x6c, 0x73, 0x18, 0x06, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2b, 0x2e, 0x79, 0x61, 0x6e,
	0x64, 0x65, 0x78, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x76, 0x70, 0x63, 0x2e, 0x76, 0x31,
	0x2e, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x54, 0x61, 0x62, 0x6c, 0x65, 0x2e, 0x4c, 0x61, 0x62, 0x65,
	0x6c, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x06, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x12,
	0x1d, 0x0a, 0x0a, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x5f, 0x69, 0x64, 0x18, 0x07, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x09, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x49, 0x64, 0x12, 0x45,
	0x0a, 0x0d, 0x73, 0x74, 0x61, 0x74, 0x69, 0x63, 0x5f, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x73, 0x18,
	0x08, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x20, 0x2e, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2e, 0x63,
	0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x76, 0x70, 0x63, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x74, 0x61, 0x74,
	0x69, 0x63, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x52, 0x0c, 0x73, 0x74, 0x61, 0x74, 0x69, 0x63, 0x52,
	0x6f, 0x75, 0x74, 0x65, 0x73, 0x1a, 0x39, 0x0a, 0x0b, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x45,
	0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01,
	0x22, 0x86, 0x02, 0x0a, 0x0b, 0x53, 0x74, 0x61, 0x74, 0x69, 0x63, 0x52, 0x6f, 0x75, 0x74, 0x65,
	0x12, 0x2f, 0x0a, 0x12, 0x64, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f,
	0x70, 0x72, 0x65, 0x66, 0x69, 0x78, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x11,
	0x64, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x50, 0x72, 0x65, 0x66, 0x69,
	0x78, 0x12, 0x2a, 0x0a, 0x10, 0x6e, 0x65, 0x78, 0x74, 0x5f, 0x68, 0x6f, 0x70, 0x5f, 0x61, 0x64,
	0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x48, 0x01, 0x52, 0x0e, 0x6e,
	0x65, 0x78, 0x74, 0x48, 0x6f, 0x70, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x44, 0x0a,
	0x06, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2c, 0x2e,
	0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x76, 0x70, 0x63,
	0x2e, 0x76, 0x31, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x69, 0x63, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x2e,
	0x4c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x06, 0x6c, 0x61, 0x62,
	0x65, 0x6c, 0x73, 0x1a, 0x39, 0x0a, 0x0b, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x45, 0x6e, 0x74,
	0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x42, 0x0d,
	0x0a, 0x0b, 0x64, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x42, 0x0a, 0x0a,
	0x08, 0x6e, 0x65, 0x78, 0x74, 0x5f, 0x68, 0x6f, 0x70, 0x42, 0x56, 0x0a, 0x17, 0x79, 0x61, 0x6e,
	0x64, 0x65, 0x78, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x70,
	0x63, 0x2e, 0x76, 0x31, 0x5a, 0x3b, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2d, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x67, 0x6f,
	0x2d, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78,
	0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x76, 0x70, 0x63, 0x2f, 0x76, 0x31, 0x3b, 0x76, 0x70,
	0x63, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_yandex_cloud_vpc_v1_route_table_proto_rawDescOnce sync.Once
	file_yandex_cloud_vpc_v1_route_table_proto_rawDescData = file_yandex_cloud_vpc_v1_route_table_proto_rawDesc
)

func file_yandex_cloud_vpc_v1_route_table_proto_rawDescGZIP() []byte {
	file_yandex_cloud_vpc_v1_route_table_proto_rawDescOnce.Do(func() {
		file_yandex_cloud_vpc_v1_route_table_proto_rawDescData = protoimpl.X.CompressGZIP(file_yandex_cloud_vpc_v1_route_table_proto_rawDescData)
	})
	return file_yandex_cloud_vpc_v1_route_table_proto_rawDescData
}

var file_yandex_cloud_vpc_v1_route_table_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_yandex_cloud_vpc_v1_route_table_proto_goTypes = []interface{}{
	(*RouteTable)(nil),            // 0: yandex.cloud.vpc.v1.RouteTable
	(*StaticRoute)(nil),           // 1: yandex.cloud.vpc.v1.StaticRoute
	nil,                           // 2: yandex.cloud.vpc.v1.RouteTable.LabelsEntry
	nil,                           // 3: yandex.cloud.vpc.v1.StaticRoute.LabelsEntry
	(*timestamppb.Timestamp)(nil), // 4: google.protobuf.Timestamp
}
var file_yandex_cloud_vpc_v1_route_table_proto_depIdxs = []int32{
	4, // 0: yandex.cloud.vpc.v1.RouteTable.created_at:type_name -> google.protobuf.Timestamp
	2, // 1: yandex.cloud.vpc.v1.RouteTable.labels:type_name -> yandex.cloud.vpc.v1.RouteTable.LabelsEntry
	1, // 2: yandex.cloud.vpc.v1.RouteTable.static_routes:type_name -> yandex.cloud.vpc.v1.StaticRoute
	3, // 3: yandex.cloud.vpc.v1.StaticRoute.labels:type_name -> yandex.cloud.vpc.v1.StaticRoute.LabelsEntry
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_yandex_cloud_vpc_v1_route_table_proto_init() }
func file_yandex_cloud_vpc_v1_route_table_proto_init() {
	if File_yandex_cloud_vpc_v1_route_table_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_yandex_cloud_vpc_v1_route_table_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RouteTable); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_yandex_cloud_vpc_v1_route_table_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StaticRoute); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	file_yandex_cloud_vpc_v1_route_table_proto_msgTypes[1].OneofWrappers = []interface{}{
		(*StaticRoute_DestinationPrefix)(nil),
		(*StaticRoute_NextHopAddress)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_yandex_cloud_vpc_v1_route_table_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_yandex_cloud_vpc_v1_route_table_proto_goTypes,
		DependencyIndexes: file_yandex_cloud_vpc_v1_route_table_proto_depIdxs,
		MessageInfos:      file_yandex_cloud_vpc_v1_route_table_proto_msgTypes,
	}.Build()
	File_yandex_cloud_vpc_v1_route_table_proto = out.File
	file_yandex_cloud_vpc_v1_route_table_proto_rawDesc = nil
	file_yandex_cloud_vpc_v1_route_table_proto_goTypes = nil
	file_yandex_cloud_vpc_v1_route_table_proto_depIdxs = nil
}
