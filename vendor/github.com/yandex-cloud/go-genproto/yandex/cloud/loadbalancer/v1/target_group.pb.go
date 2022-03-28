// Code generated by protoc-gen-go. DO NOT EDIT.
// source: yandex/cloud/loadbalancer/v1/target_group.proto

package loadbalancer

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
	_ "github.com/yandex-cloud/go-genproto/yandex/cloud"
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

// A TargetGroup resource. For more information, see [Target groups and resources](/docs/load-balancer/target-resources).
type TargetGroup struct {
	// Output only. ID of the target group.
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// ID of the folder that the target group belongs to.
	FolderId string `protobuf:"bytes,2,opt,name=folder_id,json=folderId,proto3" json:"folder_id,omitempty"`
	// Output only. Creation timestamp in [RFC3339](https://www.ietf.org/rfc/rfc3339.txt) text format.
	CreatedAt *timestamp.Timestamp `protobuf:"bytes,3,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	// Name of the target group.
	// The name is unique within the folder. 3-63 characters long.
	Name string `protobuf:"bytes,4,opt,name=name,proto3" json:"name,omitempty"`
	// Description of the target group. 0-256 characters long.
	Description string `protobuf:"bytes,5,opt,name=description,proto3" json:"description,omitempty"`
	// Resource labels as `` key:value `` pairs. Мaximum of 64 per resource.
	Labels map[string]string `protobuf:"bytes,6,rep,name=labels,proto3" json:"labels,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// ID of the region where the target group resides.
	RegionId string `protobuf:"bytes,7,opt,name=region_id,json=regionId,proto3" json:"region_id,omitempty"`
	// A list of targets in the target group.
	Targets              []*Target `protobuf:"bytes,9,rep,name=targets,proto3" json:"targets,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *TargetGroup) Reset()         { *m = TargetGroup{} }
func (m *TargetGroup) String() string { return proto.CompactTextString(m) }
func (*TargetGroup) ProtoMessage()    {}
func (*TargetGroup) Descriptor() ([]byte, []int) {
	return fileDescriptor_eb82306a182641a6, []int{0}
}

func (m *TargetGroup) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TargetGroup.Unmarshal(m, b)
}
func (m *TargetGroup) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TargetGroup.Marshal(b, m, deterministic)
}
func (m *TargetGroup) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TargetGroup.Merge(m, src)
}
func (m *TargetGroup) XXX_Size() int {
	return xxx_messageInfo_TargetGroup.Size(m)
}
func (m *TargetGroup) XXX_DiscardUnknown() {
	xxx_messageInfo_TargetGroup.DiscardUnknown(m)
}

var xxx_messageInfo_TargetGroup proto.InternalMessageInfo

func (m *TargetGroup) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *TargetGroup) GetFolderId() string {
	if m != nil {
		return m.FolderId
	}
	return ""
}

func (m *TargetGroup) GetCreatedAt() *timestamp.Timestamp {
	if m != nil {
		return m.CreatedAt
	}
	return nil
}

func (m *TargetGroup) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *TargetGroup) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *TargetGroup) GetLabels() map[string]string {
	if m != nil {
		return m.Labels
	}
	return nil
}

func (m *TargetGroup) GetRegionId() string {
	if m != nil {
		return m.RegionId
	}
	return ""
}

func (m *TargetGroup) GetTargets() []*Target {
	if m != nil {
		return m.Targets
	}
	return nil
}

// A Target resource. For more information, see [Target groups and resources](/docs/load-balancer/concepts/target-resources).
type Target struct {
	// ID of the subnet that targets are connected to.
	// All targets in the target group must be connected to the same subnet within a single availability zone.
	SubnetId string `protobuf:"bytes,1,opt,name=subnet_id,json=subnetId,proto3" json:"subnet_id,omitempty"`
	// IP address of the target.
	Address              string   `protobuf:"bytes,2,opt,name=address,proto3" json:"address,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Target) Reset()         { *m = Target{} }
func (m *Target) String() string { return proto.CompactTextString(m) }
func (*Target) ProtoMessage()    {}
func (*Target) Descriptor() ([]byte, []int) {
	return fileDescriptor_eb82306a182641a6, []int{1}
}

func (m *Target) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Target.Unmarshal(m, b)
}
func (m *Target) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Target.Marshal(b, m, deterministic)
}
func (m *Target) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Target.Merge(m, src)
}
func (m *Target) XXX_Size() int {
	return xxx_messageInfo_Target.Size(m)
}
func (m *Target) XXX_DiscardUnknown() {
	xxx_messageInfo_Target.DiscardUnknown(m)
}

var xxx_messageInfo_Target proto.InternalMessageInfo

func (m *Target) GetSubnetId() string {
	if m != nil {
		return m.SubnetId
	}
	return ""
}

func (m *Target) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func init() {
	proto.RegisterType((*TargetGroup)(nil), "yandex.cloud.loadbalancer.v1.TargetGroup")
	proto.RegisterMapType((map[string]string)(nil), "yandex.cloud.loadbalancer.v1.TargetGroup.LabelsEntry")
	proto.RegisterType((*Target)(nil), "yandex.cloud.loadbalancer.v1.Target")
}

func init() {
	proto.RegisterFile("yandex/cloud/loadbalancer/v1/target_group.proto", fileDescriptor_eb82306a182641a6)
}

var fileDescriptor_eb82306a182641a6 = []byte{
	// 424 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x52, 0x4d, 0x8b, 0x14, 0x31,
	0x10, 0x65, 0x3e, 0x76, 0x66, 0xba, 0x1a, 0x44, 0x82, 0x87, 0x30, 0x2a, 0x36, 0x8b, 0xc2, 0x5c,
	0x36, 0xed, 0xac, 0x2c, 0xb8, 0x7e, 0x81, 0x0b, 0x22, 0x03, 0xee, 0xa5, 0xd9, 0x93, 0x97, 0x21,
	0xdd, 0xa9, 0x8d, 0xc1, 0x4c, 0x32, 0xa6, 0xd3, 0x83, 0xf3, 0x17, 0x3c, 0xfa, 0xab, 0xfc, 0x59,
	0xd2, 0x49, 0x37, 0xf4, 0x7a, 0x58, 0xbc, 0xa5, 0xaa, 0xf2, 0xea, 0xbd, 0x57, 0x55, 0x90, 0x1f,
	0xb9, 0x11, 0xf8, 0x33, 0xaf, 0xb4, 0x6d, 0x44, 0xae, 0x2d, 0x17, 0x25, 0xd7, 0xdc, 0x54, 0xe8,
	0xf2, 0xc3, 0x3a, 0xf7, 0xdc, 0x49, 0xf4, 0x5b, 0xe9, 0x6c, 0xb3, 0x67, 0x7b, 0x67, 0xbd, 0x25,
	0x4f, 0x22, 0x80, 0x05, 0x00, 0x1b, 0x02, 0xd8, 0x61, 0xbd, 0x7c, 0x26, 0xad, 0x95, 0x1a, 0xf3,
	0xf0, 0xb7, 0x6c, 0x6e, 0x73, 0xaf, 0x76, 0x58, 0x7b, 0xbe, 0xeb, 0xe0, 0xcb, 0xa7, 0x77, 0xf8,
	0x0e, 0x5c, 0x2b, 0xc1, 0xbd, 0xb2, 0x26, 0x96, 0x4f, 0x7f, 0x4f, 0x20, 0xbd, 0x09, 0xa4, 0x9f,
	0x5b, 0x4e, 0xf2, 0x00, 0xc6, 0x4a, 0xd0, 0x51, 0x36, 0x5a, 0x25, 0xc5, 0x58, 0x09, 0xf2, 0x18,
	0x92, 0x5b, 0xab, 0x05, 0xba, 0xad, 0x12, 0x74, 0x1c, 0xd2, 0x8b, 0x98, 0xd8, 0x08, 0x72, 0x09,
	0x50, 0x39, 0xe4, 0x1e, 0xc5, 0x96, 0x7b, 0x3a, 0xc9, 0x46, 0xab, 0xf4, 0x7c, 0xc9, 0xa2, 0x22,
	0xd6, 0x2b, 0x62, 0x37, 0xbd, 0xa2, 0x22, 0xe9, 0x7e, 0x7f, 0xf4, 0x84, 0xc0, 0xd4, 0xf0, 0x1d,
	0xd2, 0x69, 0x68, 0x19, 0xde, 0x24, 0x83, 0x54, 0x60, 0x5d, 0x39, 0xb5, 0x6f, 0x05, 0xd2, 0x93,
	0x50, 0x1a, 0xa6, 0xc8, 0x35, 0xcc, 0x34, 0x2f, 0x51, 0xd7, 0x74, 0x96, 0x4d, 0x56, 0xe9, 0xf9,
	0x05, 0xbb, 0x6f, 0x38, 0x6c, 0x60, 0x8c, 0x7d, 0x09, 0xb8, 0x4f, 0xc6, 0xbb, 0x63, 0xd1, 0x35,
	0x69, 0xcd, 0x39, 0x94, 0xca, 0x9a, 0xd6, 0xdc, 0x3c, 0x9a, 0x8b, 0x89, 0x8d, 0x20, 0x1f, 0x60,
	0x1e, 0xb7, 0x51, 0xd3, 0x24, 0x90, 0x3d, 0xff, 0x1f, 0xb2, 0xa2, 0x07, 0x2d, 0x2f, 0x21, 0x1d,
	0x70, 0x92, 0x87, 0x30, 0xf9, 0x8e, 0xc7, 0x6e, 0xb2, 0xed, 0x93, 0x3c, 0x82, 0x93, 0x03, 0xd7,
	0x0d, 0x76, 0x63, 0x8d, 0xc1, 0x9b, 0xf1, 0xeb, 0xd1, 0xe9, 0x06, 0x66, 0xb1, 0x1b, 0x79, 0x01,
	0x49, 0xdd, 0x94, 0x06, 0xfd, 0xb6, 0xdf, 0xca, 0xd5, 0xe2, 0xd7, 0x9f, 0xf5, 0xf4, 0xdd, 0xfb,
	0x8b, 0x97, 0xc5, 0x22, 0x96, 0x36, 0x82, 0x50, 0x98, 0x73, 0x21, 0x1c, 0xd6, 0x75, 0xd7, 0xac,
	0x0f, 0xaf, 0x7e, 0x40, 0x76, 0x47, 0x35, 0xdf, 0xab, 0x7f, 0x95, 0x7f, 0xbd, 0x96, 0xca, 0x7f,
	0x6b, 0x4a, 0x56, 0xd9, 0x5d, 0x77, 0x9d, 0x67, 0xf1, 0x5a, 0xa4, 0x3d, 0x93, 0x68, 0xc2, 0x22,
	0xef, 0x3d, 0xdb, 0xb7, 0xc3, 0xb8, 0x9c, 0x05, 0xc0, 0xab, 0xbf, 0x01, 0x00, 0x00, 0xff, 0xff,
	0xf7, 0x22, 0x64, 0x9d, 0xea, 0x02, 0x00, 0x00,
}