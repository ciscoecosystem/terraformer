// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.17.3
// source: yandex/cloud/mdb/kafka/v1/user.proto

package kafka

import (
	_ "github.com/yandex-cloud/go-genproto/yandex/cloud"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Permission_AccessRole int32

const (
	Permission_ACCESS_ROLE_UNSPECIFIED Permission_AccessRole = 0
	// producer role for the user.
	Permission_ACCESS_ROLE_PRODUCER Permission_AccessRole = 1
	// consumer role for the user.
	Permission_ACCESS_ROLE_CONSUMER Permission_AccessRole = 2
	// admin role for the user.
	Permission_ACCESS_ROLE_ADMIN Permission_AccessRole = 3
)

// Enum value maps for Permission_AccessRole.
var (
	Permission_AccessRole_name = map[int32]string{
		0: "ACCESS_ROLE_UNSPECIFIED",
		1: "ACCESS_ROLE_PRODUCER",
		2: "ACCESS_ROLE_CONSUMER",
		3: "ACCESS_ROLE_ADMIN",
	}
	Permission_AccessRole_value = map[string]int32{
		"ACCESS_ROLE_UNSPECIFIED": 0,
		"ACCESS_ROLE_PRODUCER":    1,
		"ACCESS_ROLE_CONSUMER":    2,
		"ACCESS_ROLE_ADMIN":       3,
	}
)

func (x Permission_AccessRole) Enum() *Permission_AccessRole {
	p := new(Permission_AccessRole)
	*p = x
	return p
}

func (x Permission_AccessRole) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Permission_AccessRole) Descriptor() protoreflect.EnumDescriptor {
	return file_yandex_cloud_mdb_kafka_v1_user_proto_enumTypes[0].Descriptor()
}

func (Permission_AccessRole) Type() protoreflect.EnumType {
	return &file_yandex_cloud_mdb_kafka_v1_user_proto_enumTypes[0]
}

func (x Permission_AccessRole) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Permission_AccessRole.Descriptor instead.
func (Permission_AccessRole) EnumDescriptor() ([]byte, []int) {
	return file_yandex_cloud_mdb_kafka_v1_user_proto_rawDescGZIP(), []int{2, 0}
}

// A Kafka user.
// For more information, see the [Operations -> Accounts](/docs/managed-kafka/operations/cluster-accounts) section of the documentation.
type User struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Name of the Kafka user.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// ID of the Apache Kafka® cluster the user belongs to.
	//
	// To get the Apache Kafka® cluster ID, make a [ClusterService.List] request.
	ClusterId string `protobuf:"bytes,2,opt,name=cluster_id,json=clusterId,proto3" json:"cluster_id,omitempty"`
	// Set of permissions granted to this user.
	Permissions []*Permission `protobuf:"bytes,3,rep,name=permissions,proto3" json:"permissions,omitempty"`
}

func (x *User) Reset() {
	*x = User{}
	if protoimpl.UnsafeEnabled {
		mi := &file_yandex_cloud_mdb_kafka_v1_user_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *User) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*User) ProtoMessage() {}

func (x *User) ProtoReflect() protoreflect.Message {
	mi := &file_yandex_cloud_mdb_kafka_v1_user_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use User.ProtoReflect.Descriptor instead.
func (*User) Descriptor() ([]byte, []int) {
	return file_yandex_cloud_mdb_kafka_v1_user_proto_rawDescGZIP(), []int{0}
}

func (x *User) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *User) GetClusterId() string {
	if x != nil {
		return x.ClusterId
	}
	return ""
}

func (x *User) GetPermissions() []*Permission {
	if x != nil {
		return x.Permissions
	}
	return nil
}

type UserSpec struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Name of the Kafka user.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// Password of the Kafka user.
	Password string `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
	// Set of permissions granted to the user.
	Permissions []*Permission `protobuf:"bytes,3,rep,name=permissions,proto3" json:"permissions,omitempty"`
}

func (x *UserSpec) Reset() {
	*x = UserSpec{}
	if protoimpl.UnsafeEnabled {
		mi := &file_yandex_cloud_mdb_kafka_v1_user_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserSpec) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserSpec) ProtoMessage() {}

func (x *UserSpec) ProtoReflect() protoreflect.Message {
	mi := &file_yandex_cloud_mdb_kafka_v1_user_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserSpec.ProtoReflect.Descriptor instead.
func (*UserSpec) Descriptor() ([]byte, []int) {
	return file_yandex_cloud_mdb_kafka_v1_user_proto_rawDescGZIP(), []int{1}
}

func (x *UserSpec) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *UserSpec) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

func (x *UserSpec) GetPermissions() []*Permission {
	if x != nil {
		return x.Permissions
	}
	return nil
}

type Permission struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Name or prefix-pattern with wildcard for the topic that the permission grants access to.
	//
	// To get the topic name, make a [TopicService.List] request.
	TopicName string `protobuf:"bytes,1,opt,name=topic_name,json=topicName,proto3" json:"topic_name,omitempty"`
	// Access role type to grant to the user.
	Role Permission_AccessRole `protobuf:"varint,2,opt,name=role,proto3,enum=yandex.cloud.mdb.kafka.v1.Permission_AccessRole" json:"role,omitempty"`
}

func (x *Permission) Reset() {
	*x = Permission{}
	if protoimpl.UnsafeEnabled {
		mi := &file_yandex_cloud_mdb_kafka_v1_user_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Permission) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Permission) ProtoMessage() {}

func (x *Permission) ProtoReflect() protoreflect.Message {
	mi := &file_yandex_cloud_mdb_kafka_v1_user_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Permission.ProtoReflect.Descriptor instead.
func (*Permission) Descriptor() ([]byte, []int) {
	return file_yandex_cloud_mdb_kafka_v1_user_proto_rawDescGZIP(), []int{2}
}

func (x *Permission) GetTopicName() string {
	if x != nil {
		return x.TopicName
	}
	return ""
}

func (x *Permission) GetRole() Permission_AccessRole {
	if x != nil {
		return x.Role
	}
	return Permission_ACCESS_ROLE_UNSPECIFIED
}

var File_yandex_cloud_mdb_kafka_v1_user_proto protoreflect.FileDescriptor

var file_yandex_cloud_mdb_kafka_v1_user_proto_rawDesc = []byte{
	0x0a, 0x24, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x6d,
	0x64, 0x62, 0x2f, 0x6b, 0x61, 0x66, 0x6b, 0x61, 0x2f, 0x76, 0x31, 0x2f, 0x75, 0x73, 0x65, 0x72,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x19, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2e, 0x63,
	0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x6d, 0x64, 0x62, 0x2e, 0x6b, 0x61, 0x66, 0x6b, 0x61, 0x2e, 0x76,
	0x31, 0x1a, 0x1d, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f,
	0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0x82, 0x01, 0x0a, 0x04, 0x55, 0x73, 0x65, 0x72, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1d, 0x0a,
	0x0a, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x09, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x49, 0x64, 0x12, 0x47, 0x0a, 0x0b,
	0x70, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x25, 0x2e, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64,
	0x2e, 0x6d, 0x64, 0x62, 0x2e, 0x6b, 0x61, 0x66, 0x6b, 0x61, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x65,
	0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x0b, 0x70, 0x65, 0x72, 0x6d, 0x69, 0x73,
	0x73, 0x69, 0x6f, 0x6e, 0x73, 0x22, 0xb2, 0x01, 0x0a, 0x08, 0x55, 0x73, 0x65, 0x72, 0x53, 0x70,
	0x65, 0x63, 0x12, 0x32, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x42, 0x1e, 0xe8, 0xc7, 0x31, 0x01, 0x8a, 0xc8, 0x31, 0x05, 0x31, 0x2d, 0x32, 0x35, 0x36, 0xf2,
	0xc7, 0x31, 0x0d, 0x5b, 0x61, 0x2d, 0x7a, 0x41, 0x2d, 0x5a, 0x30, 0x2d, 0x39, 0x5f, 0x5d, 0x2a,
	0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x29, 0x0a, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f,
	0x72, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x0d, 0xe8, 0xc7, 0x31, 0x01, 0x8a, 0xc8,
	0x31, 0x05, 0x38, 0x2d, 0x31, 0x32, 0x38, 0x52, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72,
	0x64, 0x12, 0x47, 0x0a, 0x0b, 0x70, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x73,
	0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x25, 0x2e, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2e,
	0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x6d, 0x64, 0x62, 0x2e, 0x6b, 0x61, 0x66, 0x6b, 0x61, 0x2e,
	0x76, 0x31, 0x2e, 0x50, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x0b, 0x70,
	0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x22, 0xe7, 0x01, 0x0a, 0x0a, 0x50,
	0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x1d, 0x0a, 0x0a, 0x74, 0x6f, 0x70,
	0x69, 0x63, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x74,
	0x6f, 0x70, 0x69, 0x63, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x44, 0x0a, 0x04, 0x72, 0x6f, 0x6c, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x30, 0x2e, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2e,
	0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x6d, 0x64, 0x62, 0x2e, 0x6b, 0x61, 0x66, 0x6b, 0x61, 0x2e,
	0x76, 0x31, 0x2e, 0x50, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x2e, 0x41, 0x63,
	0x63, 0x65, 0x73, 0x73, 0x52, 0x6f, 0x6c, 0x65, 0x52, 0x04, 0x72, 0x6f, 0x6c, 0x65, 0x22, 0x74,
	0x0a, 0x0a, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x52, 0x6f, 0x6c, 0x65, 0x12, 0x1b, 0x0a, 0x17,
	0x41, 0x43, 0x43, 0x45, 0x53, 0x53, 0x5f, 0x52, 0x4f, 0x4c, 0x45, 0x5f, 0x55, 0x4e, 0x53, 0x50,
	0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x18, 0x0a, 0x14, 0x41, 0x43, 0x43,
	0x45, 0x53, 0x53, 0x5f, 0x52, 0x4f, 0x4c, 0x45, 0x5f, 0x50, 0x52, 0x4f, 0x44, 0x55, 0x43, 0x45,
	0x52, 0x10, 0x01, 0x12, 0x18, 0x0a, 0x14, 0x41, 0x43, 0x43, 0x45, 0x53, 0x53, 0x5f, 0x52, 0x4f,
	0x4c, 0x45, 0x5f, 0x43, 0x4f, 0x4e, 0x53, 0x55, 0x4d, 0x45, 0x52, 0x10, 0x02, 0x12, 0x15, 0x0a,
	0x11, 0x41, 0x43, 0x43, 0x45, 0x53, 0x53, 0x5f, 0x52, 0x4f, 0x4c, 0x45, 0x5f, 0x41, 0x44, 0x4d,
	0x49, 0x4e, 0x10, 0x03, 0x42, 0x64, 0x0a, 0x1d, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2e, 0x63,
	0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x6d, 0x64, 0x62, 0x2e, 0x6b, 0x61, 0x66,
	0x6b, 0x61, 0x2e, 0x76, 0x31, 0x5a, 0x43, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2d, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x67,
	0x6f, 0x2d, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x79, 0x61, 0x6e, 0x64, 0x65,
	0x78, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x6d, 0x64, 0x62, 0x2f, 0x6b, 0x61, 0x66, 0x6b,
	0x61, 0x2f, 0x76, 0x31, 0x3b, 0x6b, 0x61, 0x66, 0x6b, 0x61, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_yandex_cloud_mdb_kafka_v1_user_proto_rawDescOnce sync.Once
	file_yandex_cloud_mdb_kafka_v1_user_proto_rawDescData = file_yandex_cloud_mdb_kafka_v1_user_proto_rawDesc
)

func file_yandex_cloud_mdb_kafka_v1_user_proto_rawDescGZIP() []byte {
	file_yandex_cloud_mdb_kafka_v1_user_proto_rawDescOnce.Do(func() {
		file_yandex_cloud_mdb_kafka_v1_user_proto_rawDescData = protoimpl.X.CompressGZIP(file_yandex_cloud_mdb_kafka_v1_user_proto_rawDescData)
	})
	return file_yandex_cloud_mdb_kafka_v1_user_proto_rawDescData
}

var file_yandex_cloud_mdb_kafka_v1_user_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_yandex_cloud_mdb_kafka_v1_user_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_yandex_cloud_mdb_kafka_v1_user_proto_goTypes = []interface{}{
	(Permission_AccessRole)(0), // 0: yandex.cloud.mdb.kafka.v1.Permission.AccessRole
	(*User)(nil),               // 1: yandex.cloud.mdb.kafka.v1.User
	(*UserSpec)(nil),           // 2: yandex.cloud.mdb.kafka.v1.UserSpec
	(*Permission)(nil),         // 3: yandex.cloud.mdb.kafka.v1.Permission
}
var file_yandex_cloud_mdb_kafka_v1_user_proto_depIdxs = []int32{
	3, // 0: yandex.cloud.mdb.kafka.v1.User.permissions:type_name -> yandex.cloud.mdb.kafka.v1.Permission
	3, // 1: yandex.cloud.mdb.kafka.v1.UserSpec.permissions:type_name -> yandex.cloud.mdb.kafka.v1.Permission
	0, // 2: yandex.cloud.mdb.kafka.v1.Permission.role:type_name -> yandex.cloud.mdb.kafka.v1.Permission.AccessRole
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_yandex_cloud_mdb_kafka_v1_user_proto_init() }
func file_yandex_cloud_mdb_kafka_v1_user_proto_init() {
	if File_yandex_cloud_mdb_kafka_v1_user_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_yandex_cloud_mdb_kafka_v1_user_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*User); i {
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
		file_yandex_cloud_mdb_kafka_v1_user_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserSpec); i {
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
		file_yandex_cloud_mdb_kafka_v1_user_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Permission); i {
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
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_yandex_cloud_mdb_kafka_v1_user_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_yandex_cloud_mdb_kafka_v1_user_proto_goTypes,
		DependencyIndexes: file_yandex_cloud_mdb_kafka_v1_user_proto_depIdxs,
		EnumInfos:         file_yandex_cloud_mdb_kafka_v1_user_proto_enumTypes,
		MessageInfos:      file_yandex_cloud_mdb_kafka_v1_user_proto_msgTypes,
	}.Build()
	File_yandex_cloud_mdb_kafka_v1_user_proto = out.File
	file_yandex_cloud_mdb_kafka_v1_user_proto_rawDesc = nil
	file_yandex_cloud_mdb_kafka_v1_user_proto_goTypes = nil
	file_yandex_cloud_mdb_kafka_v1_user_proto_depIdxs = nil
}
