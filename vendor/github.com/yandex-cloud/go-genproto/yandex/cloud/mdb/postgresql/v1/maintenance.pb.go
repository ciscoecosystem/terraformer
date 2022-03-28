// Code generated by protoc-gen-go. DO NOT EDIT.
// source: yandex/cloud/mdb/postgresql/v1/maintenance.proto

package postgresql

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

type WeeklyMaintenanceWindow_WeekDay int32

const (
	WeeklyMaintenanceWindow_WEEK_DAY_UNSPECIFIED WeeklyMaintenanceWindow_WeekDay = 0
	WeeklyMaintenanceWindow_MON                  WeeklyMaintenanceWindow_WeekDay = 1
	WeeklyMaintenanceWindow_TUE                  WeeklyMaintenanceWindow_WeekDay = 2
	WeeklyMaintenanceWindow_WED                  WeeklyMaintenanceWindow_WeekDay = 3
	WeeklyMaintenanceWindow_THU                  WeeklyMaintenanceWindow_WeekDay = 4
	WeeklyMaintenanceWindow_FRI                  WeeklyMaintenanceWindow_WeekDay = 5
	WeeklyMaintenanceWindow_SAT                  WeeklyMaintenanceWindow_WeekDay = 6
	WeeklyMaintenanceWindow_SUN                  WeeklyMaintenanceWindow_WeekDay = 7
)

var WeeklyMaintenanceWindow_WeekDay_name = map[int32]string{
	0: "WEEK_DAY_UNSPECIFIED",
	1: "MON",
	2: "TUE",
	3: "WED",
	4: "THU",
	5: "FRI",
	6: "SAT",
	7: "SUN",
}

var WeeklyMaintenanceWindow_WeekDay_value = map[string]int32{
	"WEEK_DAY_UNSPECIFIED": 0,
	"MON":                  1,
	"TUE":                  2,
	"WED":                  3,
	"THU":                  4,
	"FRI":                  5,
	"SAT":                  6,
	"SUN":                  7,
}

func (x WeeklyMaintenanceWindow_WeekDay) String() string {
	return proto.EnumName(WeeklyMaintenanceWindow_WeekDay_name, int32(x))
}

func (WeeklyMaintenanceWindow_WeekDay) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_47e5ef56e0290927, []int{2, 0}
}

type MaintenanceWindow struct {
	// Types that are valid to be assigned to Policy:
	//	*MaintenanceWindow_Anytime
	//	*MaintenanceWindow_WeeklyMaintenanceWindow
	Policy               isMaintenanceWindow_Policy `protobuf_oneof:"policy"`
	XXX_NoUnkeyedLiteral struct{}                   `json:"-"`
	XXX_unrecognized     []byte                     `json:"-"`
	XXX_sizecache        int32                      `json:"-"`
}

func (m *MaintenanceWindow) Reset()         { *m = MaintenanceWindow{} }
func (m *MaintenanceWindow) String() string { return proto.CompactTextString(m) }
func (*MaintenanceWindow) ProtoMessage()    {}
func (*MaintenanceWindow) Descriptor() ([]byte, []int) {
	return fileDescriptor_47e5ef56e0290927, []int{0}
}

func (m *MaintenanceWindow) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MaintenanceWindow.Unmarshal(m, b)
}
func (m *MaintenanceWindow) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MaintenanceWindow.Marshal(b, m, deterministic)
}
func (m *MaintenanceWindow) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MaintenanceWindow.Merge(m, src)
}
func (m *MaintenanceWindow) XXX_Size() int {
	return xxx_messageInfo_MaintenanceWindow.Size(m)
}
func (m *MaintenanceWindow) XXX_DiscardUnknown() {
	xxx_messageInfo_MaintenanceWindow.DiscardUnknown(m)
}

var xxx_messageInfo_MaintenanceWindow proto.InternalMessageInfo

type isMaintenanceWindow_Policy interface {
	isMaintenanceWindow_Policy()
}

type MaintenanceWindow_Anytime struct {
	Anytime *AnytimeMaintenanceWindow `protobuf:"bytes,1,opt,name=anytime,proto3,oneof"`
}

type MaintenanceWindow_WeeklyMaintenanceWindow struct {
	WeeklyMaintenanceWindow *WeeklyMaintenanceWindow `protobuf:"bytes,2,opt,name=weekly_maintenance_window,json=weeklyMaintenanceWindow,proto3,oneof"`
}

func (*MaintenanceWindow_Anytime) isMaintenanceWindow_Policy() {}

func (*MaintenanceWindow_WeeklyMaintenanceWindow) isMaintenanceWindow_Policy() {}

func (m *MaintenanceWindow) GetPolicy() isMaintenanceWindow_Policy {
	if m != nil {
		return m.Policy
	}
	return nil
}

func (m *MaintenanceWindow) GetAnytime() *AnytimeMaintenanceWindow {
	if x, ok := m.GetPolicy().(*MaintenanceWindow_Anytime); ok {
		return x.Anytime
	}
	return nil
}

func (m *MaintenanceWindow) GetWeeklyMaintenanceWindow() *WeeklyMaintenanceWindow {
	if x, ok := m.GetPolicy().(*MaintenanceWindow_WeeklyMaintenanceWindow); ok {
		return x.WeeklyMaintenanceWindow
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*MaintenanceWindow) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*MaintenanceWindow_Anytime)(nil),
		(*MaintenanceWindow_WeeklyMaintenanceWindow)(nil),
	}
}

type AnytimeMaintenanceWindow struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AnytimeMaintenanceWindow) Reset()         { *m = AnytimeMaintenanceWindow{} }
func (m *AnytimeMaintenanceWindow) String() string { return proto.CompactTextString(m) }
func (*AnytimeMaintenanceWindow) ProtoMessage()    {}
func (*AnytimeMaintenanceWindow) Descriptor() ([]byte, []int) {
	return fileDescriptor_47e5ef56e0290927, []int{1}
}

func (m *AnytimeMaintenanceWindow) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AnytimeMaintenanceWindow.Unmarshal(m, b)
}
func (m *AnytimeMaintenanceWindow) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AnytimeMaintenanceWindow.Marshal(b, m, deterministic)
}
func (m *AnytimeMaintenanceWindow) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AnytimeMaintenanceWindow.Merge(m, src)
}
func (m *AnytimeMaintenanceWindow) XXX_Size() int {
	return xxx_messageInfo_AnytimeMaintenanceWindow.Size(m)
}
func (m *AnytimeMaintenanceWindow) XXX_DiscardUnknown() {
	xxx_messageInfo_AnytimeMaintenanceWindow.DiscardUnknown(m)
}

var xxx_messageInfo_AnytimeMaintenanceWindow proto.InternalMessageInfo

type WeeklyMaintenanceWindow struct {
	Day WeeklyMaintenanceWindow_WeekDay `protobuf:"varint,1,opt,name=day,proto3,enum=yandex.cloud.mdb.postgresql.v1.WeeklyMaintenanceWindow_WeekDay" json:"day,omitempty"`
	// Hour of the day in UTC.
	Hour                 int64    `protobuf:"varint,2,opt,name=hour,proto3" json:"hour,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *WeeklyMaintenanceWindow) Reset()         { *m = WeeklyMaintenanceWindow{} }
func (m *WeeklyMaintenanceWindow) String() string { return proto.CompactTextString(m) }
func (*WeeklyMaintenanceWindow) ProtoMessage()    {}
func (*WeeklyMaintenanceWindow) Descriptor() ([]byte, []int) {
	return fileDescriptor_47e5ef56e0290927, []int{2}
}

func (m *WeeklyMaintenanceWindow) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_WeeklyMaintenanceWindow.Unmarshal(m, b)
}
func (m *WeeklyMaintenanceWindow) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_WeeklyMaintenanceWindow.Marshal(b, m, deterministic)
}
func (m *WeeklyMaintenanceWindow) XXX_Merge(src proto.Message) {
	xxx_messageInfo_WeeklyMaintenanceWindow.Merge(m, src)
}
func (m *WeeklyMaintenanceWindow) XXX_Size() int {
	return xxx_messageInfo_WeeklyMaintenanceWindow.Size(m)
}
func (m *WeeklyMaintenanceWindow) XXX_DiscardUnknown() {
	xxx_messageInfo_WeeklyMaintenanceWindow.DiscardUnknown(m)
}

var xxx_messageInfo_WeeklyMaintenanceWindow proto.InternalMessageInfo

func (m *WeeklyMaintenanceWindow) GetDay() WeeklyMaintenanceWindow_WeekDay {
	if m != nil {
		return m.Day
	}
	return WeeklyMaintenanceWindow_WEEK_DAY_UNSPECIFIED
}

func (m *WeeklyMaintenanceWindow) GetHour() int64 {
	if m != nil {
		return m.Hour
	}
	return 0
}

type MaintenanceOperation struct {
	Info                 string               `protobuf:"bytes,1,opt,name=info,proto3" json:"info,omitempty"`
	DelayedUntil         *timestamp.Timestamp `protobuf:"bytes,2,opt,name=delayed_until,json=delayedUntil,proto3" json:"delayed_until,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *MaintenanceOperation) Reset()         { *m = MaintenanceOperation{} }
func (m *MaintenanceOperation) String() string { return proto.CompactTextString(m) }
func (*MaintenanceOperation) ProtoMessage()    {}
func (*MaintenanceOperation) Descriptor() ([]byte, []int) {
	return fileDescriptor_47e5ef56e0290927, []int{3}
}

func (m *MaintenanceOperation) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MaintenanceOperation.Unmarshal(m, b)
}
func (m *MaintenanceOperation) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MaintenanceOperation.Marshal(b, m, deterministic)
}
func (m *MaintenanceOperation) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MaintenanceOperation.Merge(m, src)
}
func (m *MaintenanceOperation) XXX_Size() int {
	return xxx_messageInfo_MaintenanceOperation.Size(m)
}
func (m *MaintenanceOperation) XXX_DiscardUnknown() {
	xxx_messageInfo_MaintenanceOperation.DiscardUnknown(m)
}

var xxx_messageInfo_MaintenanceOperation proto.InternalMessageInfo

func (m *MaintenanceOperation) GetInfo() string {
	if m != nil {
		return m.Info
	}
	return ""
}

func (m *MaintenanceOperation) GetDelayedUntil() *timestamp.Timestamp {
	if m != nil {
		return m.DelayedUntil
	}
	return nil
}

func init() {
	proto.RegisterEnum("yandex.cloud.mdb.postgresql.v1.WeeklyMaintenanceWindow_WeekDay", WeeklyMaintenanceWindow_WeekDay_name, WeeklyMaintenanceWindow_WeekDay_value)
	proto.RegisterType((*MaintenanceWindow)(nil), "yandex.cloud.mdb.postgresql.v1.MaintenanceWindow")
	proto.RegisterType((*AnytimeMaintenanceWindow)(nil), "yandex.cloud.mdb.postgresql.v1.AnytimeMaintenanceWindow")
	proto.RegisterType((*WeeklyMaintenanceWindow)(nil), "yandex.cloud.mdb.postgresql.v1.WeeklyMaintenanceWindow")
	proto.RegisterType((*MaintenanceOperation)(nil), "yandex.cloud.mdb.postgresql.v1.MaintenanceOperation")
}

func init() {
	proto.RegisterFile("yandex/cloud/mdb/postgresql/v1/maintenance.proto", fileDescriptor_47e5ef56e0290927)
}

var fileDescriptor_47e5ef56e0290927 = []byte{
	// 483 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x53, 0xdd, 0x6e, 0xd3, 0x30,
	0x14, 0x5e, 0xd6, 0xac, 0x65, 0x06, 0xa6, 0x60, 0x4d, 0x5a, 0xa9, 0x18, 0xa0, 0x5c, 0x71, 0x53,
	0x9b, 0x94, 0x5f, 0x09, 0xd0, 0xd4, 0xd2, 0x4c, 0xab, 0x50, 0x3b, 0xc8, 0x1a, 0x55, 0x70, 0x13,
	0x39, 0xb5, 0x97, 0x59, 0x24, 0x76, 0x68, 0x9d, 0x96, 0xbc, 0x02, 0x4f, 0x05, 0x57, 0xf0, 0x34,
	0x88, 0x6b, 0xae, 0x50, 0x9c, 0x56, 0xdd, 0xb4, 0x1f, 0xa4, 0xdd, 0x9d, 0x1e, 0x7f, 0x3f, 0xe7,
	0x9c, 0x7e, 0x01, 0x8f, 0x73, 0x22, 0x28, 0xfb, 0x8a, 0xc7, 0xb1, 0xcc, 0x28, 0x4e, 0x68, 0x88,
	0x53, 0x39, 0x55, 0xd1, 0x84, 0x4d, 0xbf, 0xc4, 0x78, 0xe6, 0xe0, 0x84, 0x70, 0xa1, 0x98, 0x20,
	0x62, 0xcc, 0x50, 0x3a, 0x91, 0x4a, 0xc2, 0xfb, 0x25, 0x03, 0x69, 0x06, 0x4a, 0x68, 0x88, 0x56,
	0x0c, 0x34, 0x73, 0x1a, 0x0f, 0x22, 0x29, 0xa3, 0x98, 0x61, 0x8d, 0x0e, 0xb3, 0x63, 0xac, 0x78,
	0xc2, 0xa6, 0x8a, 0x24, 0x69, 0x29, 0xd0, 0xd8, 0x3d, 0x63, 0x39, 0x23, 0x31, 0xa7, 0x44, 0x71,
	0x29, 0xca, 0x67, 0xfb, 0x8f, 0x01, 0xee, 0xf4, 0x57, 0xae, 0x23, 0x2e, 0xa8, 0x9c, 0xc3, 0x21,
	0xa8, 0x11, 0x91, 0x17, 0x52, 0x75, 0xe3, 0xa1, 0xf1, 0xe8, 0x66, 0xeb, 0x25, 0xba, 0x7a, 0x0e,
	0xd4, 0x2e, 0xe1, 0xe7, 0xa4, 0x0e, 0xd6, 0xbc, 0xa5, 0x14, 0xcc, 0xc0, 0xdd, 0x39, 0x63, 0x9f,
	0xe3, 0x3c, 0x38, 0xb5, 0x67, 0x30, 0xd7, 0xb8, 0xfa, 0xba, 0xf6, 0x79, 0xf1, 0x3f, 0x9f, 0x91,
	0x16, 0xb8, 0xc8, 0x66, 0x67, 0x7e, 0xf1, 0x53, 0x67, 0x0b, 0x54, 0x53, 0x19, 0xf3, 0x71, 0x0e,
	0xcd, 0xef, 0x3f, 0x1c, 0xc3, 0x6e, 0x80, 0xfa, 0x65, 0xd3, 0xda, 0xbf, 0x0d, 0xb0, 0x73, 0x89,
	0x05, 0xfc, 0x00, 0x2a, 0x94, 0xe4, 0xfa, 0x20, 0x5b, 0xad, 0xbd, 0x6b, 0x0e, 0xaa, 0xfb, 0x5d,
	0x92, 0x7b, 0x85, 0x16, 0xbc, 0x07, 0xcc, 0x13, 0x99, 0x4d, 0xf4, 0xf2, 0x95, 0xce, 0x8d, 0xbf,
	0x3f, 0x1d, 0xd3, 0x69, 0xb6, 0x9e, 0x7a, 0xba, 0x6b, 0x87, 0xa0, 0xb6, 0x40, 0xc3, 0x3a, 0xd8,
	0x1e, 0xb9, 0xee, 0xbb, 0xa0, 0xdb, 0xfe, 0x18, 0xf8, 0x83, 0xa3, 0xf7, 0xee, 0xdb, 0xde, 0x7e,
	0xcf, 0xed, 0x5a, 0x6b, 0xb0, 0x06, 0x2a, 0xfd, 0xc3, 0x81, 0x65, 0x14, 0xc5, 0xd0, 0x77, 0xad,
	0xf5, 0xa2, 0x18, 0xb9, 0x5d, 0xab, 0xa2, 0x3b, 0x07, 0xbe, 0x65, 0x16, 0xc5, 0xbe, 0xd7, 0xb3,
	0x36, 0x8a, 0xe2, 0xa8, 0x3d, 0xb4, 0xaa, 0xba, 0xf0, 0x07, 0x56, 0xcd, 0x9e, 0x81, 0xed, 0x53,
	0x33, 0x1e, 0xa6, 0x6c, 0xa2, 0xd3, 0x01, 0x77, 0x81, 0xc9, 0xc5, 0xb1, 0xd4, 0xdb, 0x6e, 0x76,
	0x36, 0xbf, 0xfd, 0x72, 0x36, 0x5e, 0xbf, 0x69, 0x3d, 0x7b, 0xee, 0xe9, 0x36, 0xdc, 0x03, 0xb7,
	0x29, 0x8b, 0x49, 0xce, 0x68, 0x90, 0x09, 0xc5, 0xe3, 0xc5, 0xdf, 0xd7, 0x40, 0x65, 0x1c, 0xd1,
	0x32, 0x8e, 0x68, 0xb8, 0x8c, 0xa3, 0x77, 0x6b, 0x41, 0xf0, 0x0b, 0x7c, 0x67, 0x0a, 0xec, 0x33,
	0x07, 0x24, 0x29, 0x3f, 0x7f, 0xc4, 0x4f, 0xfd, 0x88, 0xab, 0x93, 0x2c, 0x44, 0x63, 0x99, 0xe0,
	0x12, 0xde, 0x2c, 0x73, 0x1c, 0xc9, 0x66, 0xc4, 0x84, 0x76, 0xc1, 0x57, 0x7f, 0x53, 0xaf, 0x56,
	0xbf, 0xc2, 0xaa, 0x26, 0x3c, 0xf9, 0x17, 0x00, 0x00, 0xff, 0xff, 0x46, 0x68, 0x3f, 0x47, 0x87,
	0x03, 0x00, 0x00,
}