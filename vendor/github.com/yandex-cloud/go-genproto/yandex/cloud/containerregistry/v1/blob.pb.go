// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.17.3
// source: yandex/cloud/containerregistry/v1/blob.proto

package containerregistry

import (
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

// A Blob resource.
type Blob struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Output only. ID of the blob.
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// Content-addressable identifier of the blob.
	Digest string `protobuf:"bytes,2,opt,name=digest,proto3" json:"digest,omitempty"`
	// Size of the blob, specified in bytes.
	Size int64 `protobuf:"varint,3,opt,name=size,proto3" json:"size,omitempty"`
	// List of blob urls.
	Urls []string `protobuf:"bytes,4,rep,name=urls,proto3" json:"urls,omitempty"`
}

func (x *Blob) Reset() {
	*x = Blob{}
	if protoimpl.UnsafeEnabled {
		mi := &file_yandex_cloud_containerregistry_v1_blob_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Blob) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Blob) ProtoMessage() {}

func (x *Blob) ProtoReflect() protoreflect.Message {
	mi := &file_yandex_cloud_containerregistry_v1_blob_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Blob.ProtoReflect.Descriptor instead.
func (*Blob) Descriptor() ([]byte, []int) {
	return file_yandex_cloud_containerregistry_v1_blob_proto_rawDescGZIP(), []int{0}
}

func (x *Blob) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Blob) GetDigest() string {
	if x != nil {
		return x.Digest
	}
	return ""
}

func (x *Blob) GetSize() int64 {
	if x != nil {
		return x.Size
	}
	return 0
}

func (x *Blob) GetUrls() []string {
	if x != nil {
		return x.Urls
	}
	return nil
}

var File_yandex_cloud_containerregistry_v1_blob_proto protoreflect.FileDescriptor

var file_yandex_cloud_containerregistry_v1_blob_proto_rawDesc = []byte{
	0x0a, 0x2c, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x63,
	0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x79,
	0x2f, 0x76, 0x31, 0x2f, 0x62, 0x6c, 0x6f, 0x62, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x21,
	0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x63, 0x6f, 0x6e,
	0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x79, 0x2e, 0x76,
	0x31, 0x22, 0x56, 0x0a, 0x04, 0x42, 0x6c, 0x6f, 0x62, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x64, 0x69, 0x67,
	0x65, 0x73, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x64, 0x69, 0x67, 0x65, 0x73,
	0x74, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x04, 0x73, 0x69, 0x7a, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x75, 0x72, 0x6c, 0x73, 0x18, 0x04, 0x20,
	0x03, 0x28, 0x09, 0x52, 0x04, 0x75, 0x72, 0x6c, 0x73, 0x42, 0x80, 0x01, 0x0a, 0x25, 0x79, 0x61,
	0x6e, 0x64, 0x65, 0x78, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63,
	0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x79,
	0x2e, 0x76, 0x31, 0x5a, 0x57, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2d, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x67, 0x6f, 0x2d,
	0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2f,
	0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x72,
	0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x79, 0x2f, 0x76, 0x31, 0x3b, 0x63, 0x6f, 0x6e, 0x74, 0x61,
	0x69, 0x6e, 0x65, 0x72, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x79, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_yandex_cloud_containerregistry_v1_blob_proto_rawDescOnce sync.Once
	file_yandex_cloud_containerregistry_v1_blob_proto_rawDescData = file_yandex_cloud_containerregistry_v1_blob_proto_rawDesc
)

func file_yandex_cloud_containerregistry_v1_blob_proto_rawDescGZIP() []byte {
	file_yandex_cloud_containerregistry_v1_blob_proto_rawDescOnce.Do(func() {
		file_yandex_cloud_containerregistry_v1_blob_proto_rawDescData = protoimpl.X.CompressGZIP(file_yandex_cloud_containerregistry_v1_blob_proto_rawDescData)
	})
	return file_yandex_cloud_containerregistry_v1_blob_proto_rawDescData
}

var file_yandex_cloud_containerregistry_v1_blob_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_yandex_cloud_containerregistry_v1_blob_proto_goTypes = []interface{}{
	(*Blob)(nil), // 0: yandex.cloud.containerregistry.v1.Blob
}
var file_yandex_cloud_containerregistry_v1_blob_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_yandex_cloud_containerregistry_v1_blob_proto_init() }
func file_yandex_cloud_containerregistry_v1_blob_proto_init() {
	if File_yandex_cloud_containerregistry_v1_blob_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_yandex_cloud_containerregistry_v1_blob_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Blob); i {
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
			RawDescriptor: file_yandex_cloud_containerregistry_v1_blob_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_yandex_cloud_containerregistry_v1_blob_proto_goTypes,
		DependencyIndexes: file_yandex_cloud_containerregistry_v1_blob_proto_depIdxs,
		MessageInfos:      file_yandex_cloud_containerregistry_v1_blob_proto_msgTypes,
	}.Build()
	File_yandex_cloud_containerregistry_v1_blob_proto = out.File
	file_yandex_cloud_containerregistry_v1_blob_proto_rawDesc = nil
	file_yandex_cloud_containerregistry_v1_blob_proto_goTypes = nil
	file_yandex_cloud_containerregistry_v1_blob_proto_depIdxs = nil
}
