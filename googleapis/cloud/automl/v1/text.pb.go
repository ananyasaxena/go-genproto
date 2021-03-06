// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/cloud/automl/v1/text.proto

package automl

import (
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

// Dataset metadata for classification.
type TextClassificationDatasetMetadata struct {
	// Required. Type of the classification problem.
	ClassificationType   ClassificationType `protobuf:"varint,1,opt,name=classification_type,json=classificationType,proto3,enum=google.cloud.automl.v1.ClassificationType" json:"classification_type,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *TextClassificationDatasetMetadata) Reset()         { *m = TextClassificationDatasetMetadata{} }
func (m *TextClassificationDatasetMetadata) String() string { return proto.CompactTextString(m) }
func (*TextClassificationDatasetMetadata) ProtoMessage()    {}
func (*TextClassificationDatasetMetadata) Descriptor() ([]byte, []int) {
	return fileDescriptor_ffc003fc1ed6094b, []int{0}
}

func (m *TextClassificationDatasetMetadata) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TextClassificationDatasetMetadata.Unmarshal(m, b)
}
func (m *TextClassificationDatasetMetadata) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TextClassificationDatasetMetadata.Marshal(b, m, deterministic)
}
func (m *TextClassificationDatasetMetadata) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TextClassificationDatasetMetadata.Merge(m, src)
}
func (m *TextClassificationDatasetMetadata) XXX_Size() int {
	return xxx_messageInfo_TextClassificationDatasetMetadata.Size(m)
}
func (m *TextClassificationDatasetMetadata) XXX_DiscardUnknown() {
	xxx_messageInfo_TextClassificationDatasetMetadata.DiscardUnknown(m)
}

var xxx_messageInfo_TextClassificationDatasetMetadata proto.InternalMessageInfo

func (m *TextClassificationDatasetMetadata) GetClassificationType() ClassificationType {
	if m != nil {
		return m.ClassificationType
	}
	return ClassificationType_CLASSIFICATION_TYPE_UNSPECIFIED
}

// Model metadata that is specific to text classification.
type TextClassificationModelMetadata struct {
	// Output only. Classification type of the dataset used to train this model.
	ClassificationType   ClassificationType `protobuf:"varint,3,opt,name=classification_type,json=classificationType,proto3,enum=google.cloud.automl.v1.ClassificationType" json:"classification_type,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *TextClassificationModelMetadata) Reset()         { *m = TextClassificationModelMetadata{} }
func (m *TextClassificationModelMetadata) String() string { return proto.CompactTextString(m) }
func (*TextClassificationModelMetadata) ProtoMessage()    {}
func (*TextClassificationModelMetadata) Descriptor() ([]byte, []int) {
	return fileDescriptor_ffc003fc1ed6094b, []int{1}
}

func (m *TextClassificationModelMetadata) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TextClassificationModelMetadata.Unmarshal(m, b)
}
func (m *TextClassificationModelMetadata) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TextClassificationModelMetadata.Marshal(b, m, deterministic)
}
func (m *TextClassificationModelMetadata) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TextClassificationModelMetadata.Merge(m, src)
}
func (m *TextClassificationModelMetadata) XXX_Size() int {
	return xxx_messageInfo_TextClassificationModelMetadata.Size(m)
}
func (m *TextClassificationModelMetadata) XXX_DiscardUnknown() {
	xxx_messageInfo_TextClassificationModelMetadata.DiscardUnknown(m)
}

var xxx_messageInfo_TextClassificationModelMetadata proto.InternalMessageInfo

func (m *TextClassificationModelMetadata) GetClassificationType() ClassificationType {
	if m != nil {
		return m.ClassificationType
	}
	return ClassificationType_CLASSIFICATION_TYPE_UNSPECIFIED
}

// Dataset metadata that is specific to text extraction
type TextExtractionDatasetMetadata struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TextExtractionDatasetMetadata) Reset()         { *m = TextExtractionDatasetMetadata{} }
func (m *TextExtractionDatasetMetadata) String() string { return proto.CompactTextString(m) }
func (*TextExtractionDatasetMetadata) ProtoMessage()    {}
func (*TextExtractionDatasetMetadata) Descriptor() ([]byte, []int) {
	return fileDescriptor_ffc003fc1ed6094b, []int{2}
}

func (m *TextExtractionDatasetMetadata) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TextExtractionDatasetMetadata.Unmarshal(m, b)
}
func (m *TextExtractionDatasetMetadata) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TextExtractionDatasetMetadata.Marshal(b, m, deterministic)
}
func (m *TextExtractionDatasetMetadata) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TextExtractionDatasetMetadata.Merge(m, src)
}
func (m *TextExtractionDatasetMetadata) XXX_Size() int {
	return xxx_messageInfo_TextExtractionDatasetMetadata.Size(m)
}
func (m *TextExtractionDatasetMetadata) XXX_DiscardUnknown() {
	xxx_messageInfo_TextExtractionDatasetMetadata.DiscardUnknown(m)
}

var xxx_messageInfo_TextExtractionDatasetMetadata proto.InternalMessageInfo

// Model metadata that is specific to text extraction.
type TextExtractionModelMetadata struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TextExtractionModelMetadata) Reset()         { *m = TextExtractionModelMetadata{} }
func (m *TextExtractionModelMetadata) String() string { return proto.CompactTextString(m) }
func (*TextExtractionModelMetadata) ProtoMessage()    {}
func (*TextExtractionModelMetadata) Descriptor() ([]byte, []int) {
	return fileDescriptor_ffc003fc1ed6094b, []int{3}
}

func (m *TextExtractionModelMetadata) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TextExtractionModelMetadata.Unmarshal(m, b)
}
func (m *TextExtractionModelMetadata) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TextExtractionModelMetadata.Marshal(b, m, deterministic)
}
func (m *TextExtractionModelMetadata) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TextExtractionModelMetadata.Merge(m, src)
}
func (m *TextExtractionModelMetadata) XXX_Size() int {
	return xxx_messageInfo_TextExtractionModelMetadata.Size(m)
}
func (m *TextExtractionModelMetadata) XXX_DiscardUnknown() {
	xxx_messageInfo_TextExtractionModelMetadata.DiscardUnknown(m)
}

var xxx_messageInfo_TextExtractionModelMetadata proto.InternalMessageInfo

// Dataset metadata for text sentiment.
type TextSentimentDatasetMetadata struct {
	// Required. A sentiment is expressed as an integer ordinal, where higher
	// value means a more positive sentiment. The range of sentiments that will be
	// used is between 0 and sentiment_max (inclusive on both ends), and all the
	// values in the range must be represented in the dataset before a model can
	// be created. sentiment_max value must be between 1 and 10 (inclusive).
	SentimentMax         int32    `protobuf:"varint,1,opt,name=sentiment_max,json=sentimentMax,proto3" json:"sentiment_max,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TextSentimentDatasetMetadata) Reset()         { *m = TextSentimentDatasetMetadata{} }
func (m *TextSentimentDatasetMetadata) String() string { return proto.CompactTextString(m) }
func (*TextSentimentDatasetMetadata) ProtoMessage()    {}
func (*TextSentimentDatasetMetadata) Descriptor() ([]byte, []int) {
	return fileDescriptor_ffc003fc1ed6094b, []int{4}
}

func (m *TextSentimentDatasetMetadata) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TextSentimentDatasetMetadata.Unmarshal(m, b)
}
func (m *TextSentimentDatasetMetadata) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TextSentimentDatasetMetadata.Marshal(b, m, deterministic)
}
func (m *TextSentimentDatasetMetadata) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TextSentimentDatasetMetadata.Merge(m, src)
}
func (m *TextSentimentDatasetMetadata) XXX_Size() int {
	return xxx_messageInfo_TextSentimentDatasetMetadata.Size(m)
}
func (m *TextSentimentDatasetMetadata) XXX_DiscardUnknown() {
	xxx_messageInfo_TextSentimentDatasetMetadata.DiscardUnknown(m)
}

var xxx_messageInfo_TextSentimentDatasetMetadata proto.InternalMessageInfo

func (m *TextSentimentDatasetMetadata) GetSentimentMax() int32 {
	if m != nil {
		return m.SentimentMax
	}
	return 0
}

// Model metadata that is specific to text sentiment.
type TextSentimentModelMetadata struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TextSentimentModelMetadata) Reset()         { *m = TextSentimentModelMetadata{} }
func (m *TextSentimentModelMetadata) String() string { return proto.CompactTextString(m) }
func (*TextSentimentModelMetadata) ProtoMessage()    {}
func (*TextSentimentModelMetadata) Descriptor() ([]byte, []int) {
	return fileDescriptor_ffc003fc1ed6094b, []int{5}
}

func (m *TextSentimentModelMetadata) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TextSentimentModelMetadata.Unmarshal(m, b)
}
func (m *TextSentimentModelMetadata) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TextSentimentModelMetadata.Marshal(b, m, deterministic)
}
func (m *TextSentimentModelMetadata) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TextSentimentModelMetadata.Merge(m, src)
}
func (m *TextSentimentModelMetadata) XXX_Size() int {
	return xxx_messageInfo_TextSentimentModelMetadata.Size(m)
}
func (m *TextSentimentModelMetadata) XXX_DiscardUnknown() {
	xxx_messageInfo_TextSentimentModelMetadata.DiscardUnknown(m)
}

var xxx_messageInfo_TextSentimentModelMetadata proto.InternalMessageInfo

func init() {
	proto.RegisterType((*TextClassificationDatasetMetadata)(nil), "google.cloud.automl.v1.TextClassificationDatasetMetadata")
	proto.RegisterType((*TextClassificationModelMetadata)(nil), "google.cloud.automl.v1.TextClassificationModelMetadata")
	proto.RegisterType((*TextExtractionDatasetMetadata)(nil), "google.cloud.automl.v1.TextExtractionDatasetMetadata")
	proto.RegisterType((*TextExtractionModelMetadata)(nil), "google.cloud.automl.v1.TextExtractionModelMetadata")
	proto.RegisterType((*TextSentimentDatasetMetadata)(nil), "google.cloud.automl.v1.TextSentimentDatasetMetadata")
	proto.RegisterType((*TextSentimentModelMetadata)(nil), "google.cloud.automl.v1.TextSentimentModelMetadata")
}

func init() { proto.RegisterFile("google/cloud/automl/v1/text.proto", fileDescriptor_ffc003fc1ed6094b) }

var fileDescriptor_ffc003fc1ed6094b = []byte{
	// 354 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x92, 0xcd, 0x4a, 0xc3, 0x40,
	0x14, 0x85, 0x49, 0x45, 0xc1, 0x41, 0x5d, 0x54, 0x28, 0x1a, 0x5b, 0x6a, 0xe3, 0x46, 0x14, 0x26,
	0x44, 0x77, 0xd1, 0x8d, 0xad, 0xe2, 0xc6, 0x40, 0xa9, 0xa5, 0x0b, 0x2d, 0x94, 0x6b, 0x3a, 0x86,
	0xc0, 0x64, 0x26, 0x34, 0xb7, 0x25, 0xdd, 0x88, 0xcf, 0xe3, 0xde, 0x97, 0xf0, 0x51, 0x7c, 0x0a,
	0xc9, 0x4c, 0x2b, 0x4c, 0x7f, 0x70, 0xe3, 0x2e, 0xcc, 0x77, 0xce, 0x3d, 0xf7, 0x27, 0xa4, 0x11,
	0x49, 0x19, 0x71, 0xe6, 0x86, 0x5c, 0x8e, 0x87, 0x2e, 0x8c, 0x51, 0x26, 0xdc, 0x9d, 0x78, 0x2e,
	0xb2, 0x1c, 0x69, 0x3a, 0x92, 0x28, 0xcb, 0x15, 0x2d, 0xa1, 0x4a, 0x42, 0xb5, 0x84, 0x4e, 0x3c,
	0xbb, 0x3a, 0xb3, 0x42, 0x1a, 0xbb, 0x20, 0x84, 0x44, 0xc0, 0x58, 0x8a, 0x4c, 0xbb, 0xec, 0xf3,
	0x35, 0x85, 0x43, 0x0e, 0x59, 0x16, 0xbf, 0xc6, 0xa1, 0x52, 0x6b, 0xb1, 0xf3, 0x6e, 0x91, 0x46,
	0x97, 0xe5, 0xd8, 0x32, 0xe0, 0x2d, 0x20, 0x64, 0x0c, 0x03, 0x86, 0x30, 0x04, 0x84, 0xf2, 0x33,
	0xd9, 0x37, 0xdd, 0x03, 0x9c, 0xa6, 0xec, 0xc0, 0x3a, 0xb6, 0x4e, 0xf7, 0x2e, 0xce, 0xe8, 0xea,
	0x36, 0xa9, 0x59, 0xb3, 0x3b, 0x4d, 0x59, 0xa7, 0x1c, 0x2e, 0xbd, 0x39, 0x6f, 0xa4, 0xbe, 0xdc,
	0x41, 0x20, 0x87, 0x8c, 0xff, 0x95, 0xbf, 0xf1, 0x2f, 0xf9, 0x75, 0x52, 0x2b, 0xf2, 0xef, 0x72,
	0x1c, 0x41, 0xb8, 0x62, 0x7a, 0xa7, 0x46, 0x8e, 0x4c, 0x81, 0xd1, 0x9c, 0xd3, 0x22, 0xd5, 0x02,
	0x3f, 0x32, 0x81, 0x71, 0xc2, 0x04, 0x2e, 0x2e, 0xef, 0x84, 0xec, 0x66, 0x73, 0x36, 0x48, 0x20,
	0x57, 0x6b, 0xdb, 0xec, 0xec, 0xfc, 0x3e, 0x06, 0x90, 0x3b, 0x55, 0x62, 0x1b, 0x45, 0x8c, 0x88,
	0xe6, 0xa7, 0x45, 0xec, 0x50, 0x26, 0x6b, 0x06, 0x6d, 0x6e, 0x17, 0xd6, 0x76, 0x71, 0xcf, 0xb6,
	0xf5, 0x74, 0x3d, 0x13, 0x45, 0x92, 0x83, 0x88, 0xa8, 0x1c, 0x45, 0x6e, 0xc4, 0x84, 0xba, 0xb6,
	0xab, 0x11, 0xa4, 0x71, 0xb6, 0xf8, 0x77, 0x5c, 0xe9, 0xaf, 0x8f, 0x52, 0xe5, 0x5e, 0xdb, 0x5b,
	0x2a, 0xe3, 0x66, 0x8c, 0x32, 0x78, 0xa0, 0x3d, 0xef, 0x6b, 0x0e, 0xfa, 0x0a, 0xf4, 0x15, 0xe0,
	0xfd, 0x9e, 0xf7, 0x5d, 0x3a, 0xd4, 0xc0, 0xf7, 0x15, 0xf1, 0x7d, 0xed, 0xf1, 0xfd, 0x9e, 0xf7,
	0xb2, 0xa5, 0x62, 0x2f, 0x7f, 0x02, 0x00, 0x00, 0xff, 0xff, 0x59, 0x5c, 0x49, 0x6e, 0xec, 0x02,
	0x00, 0x00,
}
