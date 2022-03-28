// Code generated by protoc-gen-go. DO NOT EDIT.
// source: yandex/cloud/containerregistry/v1/image.proto

package containerregistry

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

// An Image resource. For more information, see [Docker image](/docs/cloud/container-registry/docker-image).
type Image struct {
	// Output only. ID of the Docker image.
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// Name of the Docker image.
	// The name is unique within the registry.
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	// Content-addressable identifier of the Docker image.
	Digest string `protobuf:"bytes,3,opt,name=digest,proto3" json:"digest,omitempty"`
	// Compressed size of the Docker image, specified in bytes.
	CompressedSize int64 `protobuf:"varint,4,opt,name=compressed_size,json=compressedSize,proto3" json:"compressed_size,omitempty"`
	// Configuration of the Docker image.
	Config *Blob `protobuf:"bytes,5,opt,name=config,proto3" json:"config,omitempty"`
	// Layers of the Docker image.
	Layers []*Blob `protobuf:"bytes,6,rep,name=layers,proto3" json:"layers,omitempty"`
	// Tags of the Docker image.
	//
	// Each tag is unique within the repository.
	Tags []string `protobuf:"bytes,7,rep,name=tags,proto3" json:"tags,omitempty"`
	// Output only. Creation timestamp in [RFC3339](https://www.ietf.org/rfc/rfc3339.txt) text format.
	CreatedAt            *timestamp.Timestamp `protobuf:"bytes,8,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *Image) Reset()         { *m = Image{} }
func (m *Image) String() string { return proto.CompactTextString(m) }
func (*Image) ProtoMessage()    {}
func (*Image) Descriptor() ([]byte, []int) {
	return fileDescriptor_0fedf978b00ec0da, []int{0}
}

func (m *Image) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Image.Unmarshal(m, b)
}
func (m *Image) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Image.Marshal(b, m, deterministic)
}
func (m *Image) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Image.Merge(m, src)
}
func (m *Image) XXX_Size() int {
	return xxx_messageInfo_Image.Size(m)
}
func (m *Image) XXX_DiscardUnknown() {
	xxx_messageInfo_Image.DiscardUnknown(m)
}

var xxx_messageInfo_Image proto.InternalMessageInfo

func (m *Image) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Image) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Image) GetDigest() string {
	if m != nil {
		return m.Digest
	}
	return ""
}

func (m *Image) GetCompressedSize() int64 {
	if m != nil {
		return m.CompressedSize
	}
	return 0
}

func (m *Image) GetConfig() *Blob {
	if m != nil {
		return m.Config
	}
	return nil
}

func (m *Image) GetLayers() []*Blob {
	if m != nil {
		return m.Layers
	}
	return nil
}

func (m *Image) GetTags() []string {
	if m != nil {
		return m.Tags
	}
	return nil
}

func (m *Image) GetCreatedAt() *timestamp.Timestamp {
	if m != nil {
		return m.CreatedAt
	}
	return nil
}

func init() {
	proto.RegisterType((*Image)(nil), "yandex.cloud.containerregistry.v1.Image")
}

func init() {
	proto.RegisterFile("yandex/cloud/containerregistry/v1/image.proto", fileDescriptor_0fedf978b00ec0da)
}

var fileDescriptor_0fedf978b00ec0da = []byte{
	// 334 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x51, 0x41, 0x4b, 0x2b, 0x31,
	0x18, 0x64, 0x77, 0xdb, 0x7d, 0xaf, 0x79, 0xd0, 0x07, 0x39, 0x48, 0xe8, 0xc5, 0x55, 0x90, 0xf6,
	0x60, 0x13, 0xaa, 0x27, 0xf1, 0x20, 0xf6, 0xe6, 0x75, 0x15, 0x04, 0x2f, 0x25, 0xbb, 0xf9, 0x1a,
	0x03, 0xbb, 0x9b, 0x25, 0x49, 0x8b, 0xed, 0xc9, 0x3f, 0xe6, 0x7f, 0x93, 0x4d, 0x5a, 0x44, 0x2a,
	0x14, 0x6f, 0xf9, 0x86, 0x99, 0xc9, 0x7c, 0xf3, 0xa1, 0xe9, 0x86, 0x37, 0x02, 0xde, 0x58, 0x59,
	0xe9, 0x95, 0x60, 0xa5, 0x6e, 0x1c, 0x57, 0x0d, 0x18, 0x03, 0x52, 0x59, 0x67, 0x36, 0x6c, 0x3d,
	0x63, 0xaa, 0xe6, 0x12, 0x68, 0x6b, 0xb4, 0xd3, 0xf8, 0x2c, 0xd0, 0xa9, 0xa7, 0xd3, 0x03, 0x3a,
	0x5d, 0xcf, 0x46, 0x97, 0xc7, 0x1d, 0x8b, 0x4a, 0x17, 0xc1, 0x70, 0x74, 0x2a, 0xb5, 0x96, 0x15,
	0x30, 0x3f, 0x15, 0xab, 0x25, 0x73, 0xaa, 0x06, 0xeb, 0x78, 0xdd, 0x06, 0xc2, 0xf9, 0x47, 0x8c,
	0xfa, 0x0f, 0x5d, 0x02, 0x3c, 0x44, 0xb1, 0x12, 0x24, 0xca, 0xa2, 0xc9, 0x20, 0x8f, 0x95, 0xc0,
	0x18, 0xf5, 0x1a, 0x5e, 0x03, 0x89, 0x3d, 0xe2, 0xdf, 0xf8, 0x04, 0xa5, 0x42, 0x49, 0xb0, 0x8e,
	0x24, 0x1e, 0xdd, 0x4d, 0x78, 0x8c, 0xfe, 0x97, 0xba, 0x6e, 0x0d, 0x58, 0x0b, 0x62, 0x61, 0xd5,
	0x16, 0x48, 0x2f, 0x8b, 0x26, 0x49, 0x3e, 0xfc, 0x82, 0x1f, 0xd5, 0x16, 0xf0, 0x1d, 0x4a, 0x4b,
	0xdd, 0x2c, 0x95, 0x24, 0xfd, 0x2c, 0x9a, 0xfc, 0xbb, 0x1a, 0xd3, 0xa3, 0x1b, 0xd3, 0x79, 0xa5,
	0x8b, 0x7c, 0x27, 0xeb, 0x0c, 0x2a, 0xbe, 0x01, 0x63, 0x49, 0x9a, 0x25, 0xbf, 0x32, 0x08, 0xb2,
	0x6e, 0x2d, 0xc7, 0xa5, 0x25, 0x7f, 0xb2, 0xa4, 0x5b, 0xab, 0x7b, 0xe3, 0x1b, 0x84, 0x4a, 0x03,
	0xdc, 0x81, 0x58, 0x70, 0x47, 0xfe, 0xfa, 0x64, 0x23, 0x1a, 0xaa, 0xa3, 0xfb, 0xea, 0xe8, 0xd3,
	0xbe, 0xba, 0x7c, 0xb0, 0x63, 0xdf, 0xbb, 0xf9, 0x7b, 0x84, 0x2e, 0xbe, 0x25, 0xe0, 0xad, 0xfa,
	0x31, 0xc5, 0xcb, 0xb3, 0x54, 0xee, 0x75, 0x55, 0xd0, 0x52, 0xd7, 0x2c, 0x28, 0xa6, 0xe1, 0x86,
	0x52, 0x4f, 0x25, 0x34, 0xfe, 0x1b, 0x76, 0xf4, 0xb8, 0xb7, 0x07, 0x60, 0x91, 0x7a, 0xe9, 0xf5,
	0x67, 0x00, 0x00, 0x00, 0xff, 0xff, 0x31, 0xfb, 0x21, 0xfc, 0x6c, 0x02, 0x00, 0x00,
}