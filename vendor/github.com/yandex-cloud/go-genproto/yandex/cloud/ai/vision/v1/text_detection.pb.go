// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.17.3
// source: yandex/cloud/ai/vision/v1/text_detection.proto

package vision

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

type TextAnnotation struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Pages of the recognized file.
	//
	// For JPEG and PNG files contains only 1 page.
	Pages []*Page `protobuf:"bytes,1,rep,name=pages,proto3" json:"pages,omitempty"`
}

func (x *TextAnnotation) Reset() {
	*x = TextAnnotation{}
	if protoimpl.UnsafeEnabled {
		mi := &file_yandex_cloud_ai_vision_v1_text_detection_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TextAnnotation) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TextAnnotation) ProtoMessage() {}

func (x *TextAnnotation) ProtoReflect() protoreflect.Message {
	mi := &file_yandex_cloud_ai_vision_v1_text_detection_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TextAnnotation.ProtoReflect.Descriptor instead.
func (*TextAnnotation) Descriptor() ([]byte, []int) {
	return file_yandex_cloud_ai_vision_v1_text_detection_proto_rawDescGZIP(), []int{0}
}

func (x *TextAnnotation) GetPages() []*Page {
	if x != nil {
		return x.Pages
	}
	return nil
}

type Page struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Page width in pixels.
	Width int64 `protobuf:"varint,1,opt,name=width,proto3" json:"width,omitempty"`
	// Page height in pixels.
	Height int64 `protobuf:"varint,2,opt,name=height,proto3" json:"height,omitempty"`
	// Recognized text blocks in this page.
	Blocks []*Block `protobuf:"bytes,3,rep,name=blocks,proto3" json:"blocks,omitempty"`
	// Recognized entities
	Entities []*Entity `protobuf:"bytes,4,rep,name=entities,proto3" json:"entities,omitempty"`
}

func (x *Page) Reset() {
	*x = Page{}
	if protoimpl.UnsafeEnabled {
		mi := &file_yandex_cloud_ai_vision_v1_text_detection_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Page) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Page) ProtoMessage() {}

func (x *Page) ProtoReflect() protoreflect.Message {
	mi := &file_yandex_cloud_ai_vision_v1_text_detection_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Page.ProtoReflect.Descriptor instead.
func (*Page) Descriptor() ([]byte, []int) {
	return file_yandex_cloud_ai_vision_v1_text_detection_proto_rawDescGZIP(), []int{1}
}

func (x *Page) GetWidth() int64 {
	if x != nil {
		return x.Width
	}
	return 0
}

func (x *Page) GetHeight() int64 {
	if x != nil {
		return x.Height
	}
	return 0
}

func (x *Page) GetBlocks() []*Block {
	if x != nil {
		return x.Blocks
	}
	return nil
}

func (x *Page) GetEntities() []*Entity {
	if x != nil {
		return x.Entities
	}
	return nil
}

type Entity struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Entity name
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// Recognized entity text
	Text string `protobuf:"bytes,2,opt,name=text,proto3" json:"text,omitempty"`
}

func (x *Entity) Reset() {
	*x = Entity{}
	if protoimpl.UnsafeEnabled {
		mi := &file_yandex_cloud_ai_vision_v1_text_detection_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Entity) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Entity) ProtoMessage() {}

func (x *Entity) ProtoReflect() protoreflect.Message {
	mi := &file_yandex_cloud_ai_vision_v1_text_detection_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Entity.ProtoReflect.Descriptor instead.
func (*Entity) Descriptor() ([]byte, []int) {
	return file_yandex_cloud_ai_vision_v1_text_detection_proto_rawDescGZIP(), []int{2}
}

func (x *Entity) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Entity) GetText() string {
	if x != nil {
		return x.Text
	}
	return ""
}

type Block struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Area on the page where the text block is located.
	BoundingBox *Polygon `protobuf:"bytes,1,opt,name=bounding_box,json=boundingBox,proto3" json:"bounding_box,omitempty"`
	// Recognized lines in this block.
	Lines []*Line `protobuf:"bytes,2,rep,name=lines,proto3" json:"lines,omitempty"`
}

func (x *Block) Reset() {
	*x = Block{}
	if protoimpl.UnsafeEnabled {
		mi := &file_yandex_cloud_ai_vision_v1_text_detection_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Block) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Block) ProtoMessage() {}

func (x *Block) ProtoReflect() protoreflect.Message {
	mi := &file_yandex_cloud_ai_vision_v1_text_detection_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Block.ProtoReflect.Descriptor instead.
func (*Block) Descriptor() ([]byte, []int) {
	return file_yandex_cloud_ai_vision_v1_text_detection_proto_rawDescGZIP(), []int{3}
}

func (x *Block) GetBoundingBox() *Polygon {
	if x != nil {
		return x.BoundingBox
	}
	return nil
}

func (x *Block) GetLines() []*Line {
	if x != nil {
		return x.Lines
	}
	return nil
}

type Line struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Area on the page where the line is located.
	BoundingBox *Polygon `protobuf:"bytes,1,opt,name=bounding_box,json=boundingBox,proto3" json:"bounding_box,omitempty"`
	// Recognized words in this line.
	Words []*Word `protobuf:"bytes,2,rep,name=words,proto3" json:"words,omitempty"`
	// Confidence of the OCR results for the line. Range [0, 1].
	Confidence float64 `protobuf:"fixed64,3,opt,name=confidence,proto3" json:"confidence,omitempty"`
}

func (x *Line) Reset() {
	*x = Line{}
	if protoimpl.UnsafeEnabled {
		mi := &file_yandex_cloud_ai_vision_v1_text_detection_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Line) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Line) ProtoMessage() {}

func (x *Line) ProtoReflect() protoreflect.Message {
	mi := &file_yandex_cloud_ai_vision_v1_text_detection_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Line.ProtoReflect.Descriptor instead.
func (*Line) Descriptor() ([]byte, []int) {
	return file_yandex_cloud_ai_vision_v1_text_detection_proto_rawDescGZIP(), []int{4}
}

func (x *Line) GetBoundingBox() *Polygon {
	if x != nil {
		return x.BoundingBox
	}
	return nil
}

func (x *Line) GetWords() []*Word {
	if x != nil {
		return x.Words
	}
	return nil
}

func (x *Line) GetConfidence() float64 {
	if x != nil {
		return x.Confidence
	}
	return 0
}

type Word struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Area on the page where the word is located.
	BoundingBox *Polygon `protobuf:"bytes,1,opt,name=bounding_box,json=boundingBox,proto3" json:"bounding_box,omitempty"`
	// Recognized word value.
	Text string `protobuf:"bytes,2,opt,name=text,proto3" json:"text,omitempty"`
	// Confidence of the OCR results for the word. Range [0, 1].
	Confidence float64 `protobuf:"fixed64,3,opt,name=confidence,proto3" json:"confidence,omitempty"`
	// A list of detected languages together with confidence.
	Languages []*Word_DetectedLanguage `protobuf:"bytes,4,rep,name=languages,proto3" json:"languages,omitempty"`
	// Id of recognized word in entities array
	EntityIndex int64 `protobuf:"varint,5,opt,name=entity_index,json=entityIndex,proto3" json:"entity_index,omitempty"`
}

func (x *Word) Reset() {
	*x = Word{}
	if protoimpl.UnsafeEnabled {
		mi := &file_yandex_cloud_ai_vision_v1_text_detection_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Word) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Word) ProtoMessage() {}

func (x *Word) ProtoReflect() protoreflect.Message {
	mi := &file_yandex_cloud_ai_vision_v1_text_detection_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Word.ProtoReflect.Descriptor instead.
func (*Word) Descriptor() ([]byte, []int) {
	return file_yandex_cloud_ai_vision_v1_text_detection_proto_rawDescGZIP(), []int{5}
}

func (x *Word) GetBoundingBox() *Polygon {
	if x != nil {
		return x.BoundingBox
	}
	return nil
}

func (x *Word) GetText() string {
	if x != nil {
		return x.Text
	}
	return ""
}

func (x *Word) GetConfidence() float64 {
	if x != nil {
		return x.Confidence
	}
	return 0
}

func (x *Word) GetLanguages() []*Word_DetectedLanguage {
	if x != nil {
		return x.Languages
	}
	return nil
}

func (x *Word) GetEntityIndex() int64 {
	if x != nil {
		return x.EntityIndex
	}
	return 0
}

type Word_DetectedLanguage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Detected language code.
	LanguageCode string `protobuf:"bytes,1,opt,name=language_code,json=languageCode,proto3" json:"language_code,omitempty"`
	// Confidence of detected language. Range [0, 1].
	Confidence float64 `protobuf:"fixed64,2,opt,name=confidence,proto3" json:"confidence,omitempty"`
}

func (x *Word_DetectedLanguage) Reset() {
	*x = Word_DetectedLanguage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_yandex_cloud_ai_vision_v1_text_detection_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Word_DetectedLanguage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Word_DetectedLanguage) ProtoMessage() {}

func (x *Word_DetectedLanguage) ProtoReflect() protoreflect.Message {
	mi := &file_yandex_cloud_ai_vision_v1_text_detection_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Word_DetectedLanguage.ProtoReflect.Descriptor instead.
func (*Word_DetectedLanguage) Descriptor() ([]byte, []int) {
	return file_yandex_cloud_ai_vision_v1_text_detection_proto_rawDescGZIP(), []int{5, 0}
}

func (x *Word_DetectedLanguage) GetLanguageCode() string {
	if x != nil {
		return x.LanguageCode
	}
	return ""
}

func (x *Word_DetectedLanguage) GetConfidence() float64 {
	if x != nil {
		return x.Confidence
	}
	return 0
}

var File_yandex_cloud_ai_vision_v1_text_detection_proto protoreflect.FileDescriptor

var file_yandex_cloud_ai_vision_v1_text_detection_proto_rawDesc = []byte{
	0x0a, 0x2e, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x61,
	0x69, 0x2f, 0x76, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x2f, 0x76, 0x31, 0x2f, 0x74, 0x65, 0x78, 0x74,
	0x5f, 0x64, 0x65, 0x74, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x19, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x61,
	0x69, 0x2e, 0x76, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x1a, 0x2a, 0x79, 0x61, 0x6e,
	0x64, 0x65, 0x78, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x61, 0x69, 0x2f, 0x76, 0x69, 0x73,
	0x69, 0x6f, 0x6e, 0x2f, 0x76, 0x31, 0x2f, 0x70, 0x72, 0x69, 0x6d, 0x69, 0x74, 0x69, 0x76, 0x65,
	0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x47, 0x0a, 0x0e, 0x54, 0x65, 0x78, 0x74, 0x41,
	0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x35, 0x0a, 0x05, 0x70, 0x61, 0x67,
	0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x79, 0x61, 0x6e, 0x64, 0x65,
	0x78, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x61, 0x69, 0x2e, 0x76, 0x69, 0x73, 0x69, 0x6f,
	0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x61, 0x67, 0x65, 0x52, 0x05, 0x70, 0x61, 0x67, 0x65, 0x73,
	0x22, 0xad, 0x01, 0x0a, 0x04, 0x50, 0x61, 0x67, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x77, 0x69, 0x64,
	0x74, 0x68, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x77, 0x69, 0x64, 0x74, 0x68, 0x12,
	0x16, 0x0a, 0x06, 0x68, 0x65, 0x69, 0x67, 0x68, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x06, 0x68, 0x65, 0x69, 0x67, 0x68, 0x74, 0x12, 0x38, 0x0a, 0x06, 0x62, 0x6c, 0x6f, 0x63, 0x6b,
	0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x20, 0x2e, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78,
	0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x61, 0x69, 0x2e, 0x76, 0x69, 0x73, 0x69, 0x6f, 0x6e,
	0x2e, 0x76, 0x31, 0x2e, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x52, 0x06, 0x62, 0x6c, 0x6f, 0x63, 0x6b,
	0x73, 0x12, 0x3d, 0x0a, 0x08, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x69, 0x65, 0x73, 0x18, 0x04, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x21, 0x2e, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2e, 0x63, 0x6c, 0x6f,
	0x75, 0x64, 0x2e, 0x61, 0x69, 0x2e, 0x76, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e,
	0x45, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x52, 0x08, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x69, 0x65, 0x73,
	0x22, 0x30, 0x0a, 0x06, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x12,
	0x0a, 0x04, 0x74, 0x65, 0x78, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x65,
	0x78, 0x74, 0x22, 0x85, 0x01, 0x0a, 0x05, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x12, 0x45, 0x0a, 0x0c,
	0x62, 0x6f, 0x75, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x5f, 0x62, 0x6f, 0x78, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x22, 0x2e, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2e, 0x63, 0x6c, 0x6f, 0x75,
	0x64, 0x2e, 0x61, 0x69, 0x2e, 0x76, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x50,
	0x6f, 0x6c, 0x79, 0x67, 0x6f, 0x6e, 0x52, 0x0b, 0x62, 0x6f, 0x75, 0x6e, 0x64, 0x69, 0x6e, 0x67,
	0x42, 0x6f, 0x78, 0x12, 0x35, 0x0a, 0x05, 0x6c, 0x69, 0x6e, 0x65, 0x73, 0x18, 0x02, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2e, 0x63, 0x6c, 0x6f, 0x75,
	0x64, 0x2e, 0x61, 0x69, 0x2e, 0x76, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x4c,
	0x69, 0x6e, 0x65, 0x52, 0x05, 0x6c, 0x69, 0x6e, 0x65, 0x73, 0x22, 0xa4, 0x01, 0x0a, 0x04, 0x4c,
	0x69, 0x6e, 0x65, 0x12, 0x45, 0x0a, 0x0c, 0x62, 0x6f, 0x75, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x5f,
	0x62, 0x6f, 0x78, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x22, 0x2e, 0x79, 0x61, 0x6e, 0x64,
	0x65, 0x78, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x61, 0x69, 0x2e, 0x76, 0x69, 0x73, 0x69,
	0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x6f, 0x6c, 0x79, 0x67, 0x6f, 0x6e, 0x52, 0x0b, 0x62,
	0x6f, 0x75, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x42, 0x6f, 0x78, 0x12, 0x35, 0x0a, 0x05, 0x77, 0x6f,
	0x72, 0x64, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x79, 0x61, 0x6e, 0x64,
	0x65, 0x78, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x61, 0x69, 0x2e, 0x76, 0x69, 0x73, 0x69,
	0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x57, 0x6f, 0x72, 0x64, 0x52, 0x05, 0x77, 0x6f, 0x72, 0x64,
	0x73, 0x12, 0x1e, 0x0a, 0x0a, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x64, 0x65, 0x6e, 0x63, 0x65, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x01, 0x52, 0x0a, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x64, 0x65, 0x6e, 0x63,
	0x65, 0x22, 0xcd, 0x02, 0x0a, 0x04, 0x57, 0x6f, 0x72, 0x64, 0x12, 0x45, 0x0a, 0x0c, 0x62, 0x6f,
	0x75, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x5f, 0x62, 0x6f, 0x78, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x22, 0x2e, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e,
	0x61, 0x69, 0x2e, 0x76, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x6f, 0x6c,
	0x79, 0x67, 0x6f, 0x6e, 0x52, 0x0b, 0x62, 0x6f, 0x75, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x42, 0x6f,
	0x78, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x65, 0x78, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x74, 0x65, 0x78, 0x74, 0x12, 0x1e, 0x0a, 0x0a, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x64, 0x65,
	0x6e, 0x63, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x01, 0x52, 0x0a, 0x63, 0x6f, 0x6e, 0x66, 0x69,
	0x64, 0x65, 0x6e, 0x63, 0x65, 0x12, 0x4e, 0x0a, 0x09, 0x6c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67,
	0x65, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x30, 0x2e, 0x79, 0x61, 0x6e, 0x64, 0x65,
	0x78, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x61, 0x69, 0x2e, 0x76, 0x69, 0x73, 0x69, 0x6f,
	0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x57, 0x6f, 0x72, 0x64, 0x2e, 0x44, 0x65, 0x74, 0x65, 0x63, 0x74,
	0x65, 0x64, 0x4c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x52, 0x09, 0x6c, 0x61, 0x6e, 0x67,
	0x75, 0x61, 0x67, 0x65, 0x73, 0x12, 0x21, 0x0a, 0x0c, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x5f,
	0x69, 0x6e, 0x64, 0x65, 0x78, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0b, 0x65, 0x6e, 0x74,
	0x69, 0x74, 0x79, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x1a, 0x57, 0x0a, 0x10, 0x44, 0x65, 0x74, 0x65,
	0x63, 0x74, 0x65, 0x64, 0x4c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x12, 0x23, 0x0a, 0x0d,
	0x6c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0c, 0x6c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x43, 0x6f, 0x64,
	0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x64, 0x65, 0x6e, 0x63, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x01, 0x52, 0x0a, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x64, 0x65, 0x6e, 0x63,
	0x65, 0x42, 0x65, 0x0a, 0x1d, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2e, 0x63, 0x6c, 0x6f, 0x75,
	0x64, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x61, 0x69, 0x2e, 0x76, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x2e,
	0x76, 0x31, 0x5a, 0x44, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x79,
	0x61, 0x6e, 0x64, 0x65, 0x78, 0x2d, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x67, 0x6f, 0x2d, 0x67,
	0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2f, 0x63,
	0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x61, 0x69, 0x2f, 0x76, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x2f, 0x76,
	0x31, 0x3b, 0x76, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_yandex_cloud_ai_vision_v1_text_detection_proto_rawDescOnce sync.Once
	file_yandex_cloud_ai_vision_v1_text_detection_proto_rawDescData = file_yandex_cloud_ai_vision_v1_text_detection_proto_rawDesc
)

func file_yandex_cloud_ai_vision_v1_text_detection_proto_rawDescGZIP() []byte {
	file_yandex_cloud_ai_vision_v1_text_detection_proto_rawDescOnce.Do(func() {
		file_yandex_cloud_ai_vision_v1_text_detection_proto_rawDescData = protoimpl.X.CompressGZIP(file_yandex_cloud_ai_vision_v1_text_detection_proto_rawDescData)
	})
	return file_yandex_cloud_ai_vision_v1_text_detection_proto_rawDescData
}

var file_yandex_cloud_ai_vision_v1_text_detection_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_yandex_cloud_ai_vision_v1_text_detection_proto_goTypes = []interface{}{
	(*TextAnnotation)(nil),        // 0: yandex.cloud.ai.vision.v1.TextAnnotation
	(*Page)(nil),                  // 1: yandex.cloud.ai.vision.v1.Page
	(*Entity)(nil),                // 2: yandex.cloud.ai.vision.v1.Entity
	(*Block)(nil),                 // 3: yandex.cloud.ai.vision.v1.Block
	(*Line)(nil),                  // 4: yandex.cloud.ai.vision.v1.Line
	(*Word)(nil),                  // 5: yandex.cloud.ai.vision.v1.Word
	(*Word_DetectedLanguage)(nil), // 6: yandex.cloud.ai.vision.v1.Word.DetectedLanguage
	(*Polygon)(nil),               // 7: yandex.cloud.ai.vision.v1.Polygon
}
var file_yandex_cloud_ai_vision_v1_text_detection_proto_depIdxs = []int32{
	1, // 0: yandex.cloud.ai.vision.v1.TextAnnotation.pages:type_name -> yandex.cloud.ai.vision.v1.Page
	3, // 1: yandex.cloud.ai.vision.v1.Page.blocks:type_name -> yandex.cloud.ai.vision.v1.Block
	2, // 2: yandex.cloud.ai.vision.v1.Page.entities:type_name -> yandex.cloud.ai.vision.v1.Entity
	7, // 3: yandex.cloud.ai.vision.v1.Block.bounding_box:type_name -> yandex.cloud.ai.vision.v1.Polygon
	4, // 4: yandex.cloud.ai.vision.v1.Block.lines:type_name -> yandex.cloud.ai.vision.v1.Line
	7, // 5: yandex.cloud.ai.vision.v1.Line.bounding_box:type_name -> yandex.cloud.ai.vision.v1.Polygon
	5, // 6: yandex.cloud.ai.vision.v1.Line.words:type_name -> yandex.cloud.ai.vision.v1.Word
	7, // 7: yandex.cloud.ai.vision.v1.Word.bounding_box:type_name -> yandex.cloud.ai.vision.v1.Polygon
	6, // 8: yandex.cloud.ai.vision.v1.Word.languages:type_name -> yandex.cloud.ai.vision.v1.Word.DetectedLanguage
	9, // [9:9] is the sub-list for method output_type
	9, // [9:9] is the sub-list for method input_type
	9, // [9:9] is the sub-list for extension type_name
	9, // [9:9] is the sub-list for extension extendee
	0, // [0:9] is the sub-list for field type_name
}

func init() { file_yandex_cloud_ai_vision_v1_text_detection_proto_init() }
func file_yandex_cloud_ai_vision_v1_text_detection_proto_init() {
	if File_yandex_cloud_ai_vision_v1_text_detection_proto != nil {
		return
	}
	file_yandex_cloud_ai_vision_v1_primitives_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_yandex_cloud_ai_vision_v1_text_detection_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TextAnnotation); i {
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
		file_yandex_cloud_ai_vision_v1_text_detection_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Page); i {
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
		file_yandex_cloud_ai_vision_v1_text_detection_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Entity); i {
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
		file_yandex_cloud_ai_vision_v1_text_detection_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Block); i {
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
		file_yandex_cloud_ai_vision_v1_text_detection_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Line); i {
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
		file_yandex_cloud_ai_vision_v1_text_detection_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Word); i {
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
		file_yandex_cloud_ai_vision_v1_text_detection_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Word_DetectedLanguage); i {
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
			RawDescriptor: file_yandex_cloud_ai_vision_v1_text_detection_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_yandex_cloud_ai_vision_v1_text_detection_proto_goTypes,
		DependencyIndexes: file_yandex_cloud_ai_vision_v1_text_detection_proto_depIdxs,
		MessageInfos:      file_yandex_cloud_ai_vision_v1_text_detection_proto_msgTypes,
	}.Build()
	File_yandex_cloud_ai_vision_v1_text_detection_proto = out.File
	file_yandex_cloud_ai_vision_v1_text_detection_proto_rawDesc = nil
	file_yandex_cloud_ai_vision_v1_text_detection_proto_goTypes = nil
	file_yandex_cloud_ai_vision_v1_text_detection_proto_depIdxs = nil
}
