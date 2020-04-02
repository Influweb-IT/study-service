// Code generated by protoc-gen-go. DO NOT EDIT.
// source: survey-response.proto

package api

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
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

type SurveyResponse struct {
	Key                  string                `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	SubmittedBy          string                `protobuf:"bytes,2,opt,name=submitted_by,json=submittedBy,proto3" json:"submitted_by,omitempty"`
	SubmittedFor         string                `protobuf:"bytes,3,opt,name=submitted_for,json=submittedFor,proto3" json:"submitted_for,omitempty"`
	SubmittedAt          int64                 `protobuf:"varint,4,opt,name=submitted_at,json=submittedAt,proto3" json:"submitted_at,omitempty"`
	Responses            []*SurveyItemResponse `protobuf:"bytes,5,rep,name=responses,proto3" json:"responses,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *SurveyResponse) Reset()         { *m = SurveyResponse{} }
func (m *SurveyResponse) String() string { return proto.CompactTextString(m) }
func (*SurveyResponse) ProtoMessage()    {}
func (*SurveyResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_9f7d4a107e829578, []int{0}
}

func (m *SurveyResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SurveyResponse.Unmarshal(m, b)
}
func (m *SurveyResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SurveyResponse.Marshal(b, m, deterministic)
}
func (m *SurveyResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SurveyResponse.Merge(m, src)
}
func (m *SurveyResponse) XXX_Size() int {
	return xxx_messageInfo_SurveyResponse.Size(m)
}
func (m *SurveyResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_SurveyResponse.DiscardUnknown(m)
}

var xxx_messageInfo_SurveyResponse proto.InternalMessageInfo

func (m *SurveyResponse) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func (m *SurveyResponse) GetSubmittedBy() string {
	if m != nil {
		return m.SubmittedBy
	}
	return ""
}

func (m *SurveyResponse) GetSubmittedFor() string {
	if m != nil {
		return m.SubmittedFor
	}
	return ""
}

func (m *SurveyResponse) GetSubmittedAt() int64 {
	if m != nil {
		return m.SubmittedAt
	}
	return 0
}

func (m *SurveyResponse) GetResponses() []*SurveyItemResponse {
	if m != nil {
		return m.Responses
	}
	return nil
}

type SurveyItemResponse struct {
	Key  string        `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	Meta *ResponseMeta `protobuf:"bytes,2,opt,name=meta,proto3" json:"meta,omitempty"`
	// for item groups:
	Items []*SurveyItemResponse `protobuf:"bytes,3,rep,name=items,proto3" json:"items,omitempty"`
	// for single items:
	Response             *ResponseItem `protobuf:"bytes,4,opt,name=response,proto3" json:"response,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *SurveyItemResponse) Reset()         { *m = SurveyItemResponse{} }
func (m *SurveyItemResponse) String() string { return proto.CompactTextString(m) }
func (*SurveyItemResponse) ProtoMessage()    {}
func (*SurveyItemResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_9f7d4a107e829578, []int{1}
}

func (m *SurveyItemResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SurveyItemResponse.Unmarshal(m, b)
}
func (m *SurveyItemResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SurveyItemResponse.Marshal(b, m, deterministic)
}
func (m *SurveyItemResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SurveyItemResponse.Merge(m, src)
}
func (m *SurveyItemResponse) XXX_Size() int {
	return xxx_messageInfo_SurveyItemResponse.Size(m)
}
func (m *SurveyItemResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_SurveyItemResponse.DiscardUnknown(m)
}

var xxx_messageInfo_SurveyItemResponse proto.InternalMessageInfo

func (m *SurveyItemResponse) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func (m *SurveyItemResponse) GetMeta() *ResponseMeta {
	if m != nil {
		return m.Meta
	}
	return nil
}

func (m *SurveyItemResponse) GetItems() []*SurveyItemResponse {
	if m != nil {
		return m.Items
	}
	return nil
}

func (m *SurveyItemResponse) GetResponse() *ResponseItem {
	if m != nil {
		return m.Response
	}
	return nil
}

type ResponseItem struct {
	Key   string `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	Value string `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
	Dtype string `protobuf:"bytes,3,opt,name=dtype,proto3" json:"dtype,omitempty"`
	// For response option groups:
	Items                []*ResponseItem `protobuf:"bytes,4,rep,name=items,proto3" json:"items,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *ResponseItem) Reset()         { *m = ResponseItem{} }
func (m *ResponseItem) String() string { return proto.CompactTextString(m) }
func (*ResponseItem) ProtoMessage()    {}
func (*ResponseItem) Descriptor() ([]byte, []int) {
	return fileDescriptor_9f7d4a107e829578, []int{2}
}

func (m *ResponseItem) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ResponseItem.Unmarshal(m, b)
}
func (m *ResponseItem) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ResponseItem.Marshal(b, m, deterministic)
}
func (m *ResponseItem) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ResponseItem.Merge(m, src)
}
func (m *ResponseItem) XXX_Size() int {
	return xxx_messageInfo_ResponseItem.Size(m)
}
func (m *ResponseItem) XXX_DiscardUnknown() {
	xxx_messageInfo_ResponseItem.DiscardUnknown(m)
}

var xxx_messageInfo_ResponseItem proto.InternalMessageInfo

func (m *ResponseItem) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func (m *ResponseItem) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

func (m *ResponseItem) GetDtype() string {
	if m != nil {
		return m.Dtype
	}
	return ""
}

func (m *ResponseItem) GetItems() []*ResponseItem {
	if m != nil {
		return m.Items
	}
	return nil
}

type ResponseMeta struct {
	Position   int32  `protobuf:"varint,1,opt,name=position,proto3" json:"position,omitempty"`
	LocaleCode string `protobuf:"bytes,2,opt,name=locale_code,json=localeCode,proto3" json:"locale_code,omitempty"`
	Version    int32  `protobuf:"varint,3,opt,name=version,proto3" json:"version,omitempty"`
	// timestamps:
	Rendered             []int64  `protobuf:"varint,4,rep,packed,name=rendered,proto3" json:"rendered,omitempty"`
	Displayed            []int64  `protobuf:"varint,5,rep,packed,name=displayed,proto3" json:"displayed,omitempty"`
	Responded            []int64  `protobuf:"varint,6,rep,packed,name=responded,proto3" json:"responded,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ResponseMeta) Reset()         { *m = ResponseMeta{} }
func (m *ResponseMeta) String() string { return proto.CompactTextString(m) }
func (*ResponseMeta) ProtoMessage()    {}
func (*ResponseMeta) Descriptor() ([]byte, []int) {
	return fileDescriptor_9f7d4a107e829578, []int{3}
}

func (m *ResponseMeta) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ResponseMeta.Unmarshal(m, b)
}
func (m *ResponseMeta) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ResponseMeta.Marshal(b, m, deterministic)
}
func (m *ResponseMeta) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ResponseMeta.Merge(m, src)
}
func (m *ResponseMeta) XXX_Size() int {
	return xxx_messageInfo_ResponseMeta.Size(m)
}
func (m *ResponseMeta) XXX_DiscardUnknown() {
	xxx_messageInfo_ResponseMeta.DiscardUnknown(m)
}

var xxx_messageInfo_ResponseMeta proto.InternalMessageInfo

func (m *ResponseMeta) GetPosition() int32 {
	if m != nil {
		return m.Position
	}
	return 0
}

func (m *ResponseMeta) GetLocaleCode() string {
	if m != nil {
		return m.LocaleCode
	}
	return ""
}

func (m *ResponseMeta) GetVersion() int32 {
	if m != nil {
		return m.Version
	}
	return 0
}

func (m *ResponseMeta) GetRendered() []int64 {
	if m != nil {
		return m.Rendered
	}
	return nil
}

func (m *ResponseMeta) GetDisplayed() []int64 {
	if m != nil {
		return m.Displayed
	}
	return nil
}

func (m *ResponseMeta) GetResponded() []int64 {
	if m != nil {
		return m.Responded
	}
	return nil
}

func init() {
	proto.RegisterType((*SurveyResponse)(nil), "inf.survey_response.SurveyResponse")
	proto.RegisterType((*SurveyItemResponse)(nil), "inf.survey_response.SurveyItemResponse")
	proto.RegisterType((*ResponseItem)(nil), "inf.survey_response.ResponseItem")
	proto.RegisterType((*ResponseMeta)(nil), "inf.survey_response.ResponseMeta")
}

func init() {
	proto.RegisterFile("survey-response.proto", fileDescriptor_9f7d4a107e829578)
}

var fileDescriptor_9f7d4a107e829578 = []byte{
	// 386 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x53, 0xcf, 0x6a, 0xdb, 0x30,
	0x18, 0xc7, 0x53, 0x9c, 0x25, 0x9f, 0xb3, 0x31, 0xb4, 0x0d, 0xc4, 0x18, 0x2c, 0xf1, 0x0e, 0xcd,
	0xa5, 0x3e, 0xa4, 0x94, 0x9e, 0x72, 0x68, 0x4a, 0x0b, 0x3d, 0xf4, 0xe2, 0xde, 0x7a, 0x09, 0x4e,
	0xf4, 0x05, 0x4c, 0x6d, 0xcb, 0x48, 0x4a, 0xc0, 0x0f, 0xd0, 0xc7, 0xea, 0x5b, 0xf4, 0xd6, 0x97,
	0x29, 0x96, 0x22, 0x3b, 0xa5, 0x86, 0xb6, 0xb7, 0x7c, 0xbf, 0x3f, 0x9f, 0x7e, 0x3f, 0x29, 0x86,
	0xdf, 0x6a, 0x2b, 0x77, 0x58, 0x1d, 0x4b, 0x54, 0xa5, 0x28, 0x14, 0x46, 0xa5, 0x14, 0x5a, 0xd0,
	0x9f, 0x69, 0xb1, 0x89, 0x2c, 0xb5, 0x74, 0x54, 0xf8, 0xe4, 0xc1, 0xf7, 0x5b, 0x83, 0xc5, 0x7b,
	0x88, 0xfe, 0x00, 0x72, 0x8f, 0x15, 0xf3, 0xc6, 0xde, 0x74, 0x18, 0xd7, 0x3f, 0xe9, 0x04, 0x46,
	0x6a, 0xbb, 0xca, 0x53, 0xad, 0x91, 0x2f, 0x57, 0x15, 0xfb, 0x62, 0xa8, 0xa0, 0xc1, 0x16, 0x15,
	0xfd, 0x0f, 0xdf, 0x5a, 0xc9, 0x46, 0x48, 0x46, 0x8c, 0xa6, 0xf5, 0x5d, 0x09, 0xf9, 0x7a, 0x4f,
	0xa2, 0x59, 0x6f, 0xec, 0x4d, 0xc9, 0xc1, 0x9e, 0x73, 0x4d, 0x2f, 0x61, 0xe8, 0xb2, 0x29, 0xe6,
	0x8f, 0xc9, 0x34, 0x98, 0x1d, 0x45, 0x1d, 0xc1, 0x23, 0x1b, 0xfa, 0x5a, 0x63, 0xee, 0x82, 0xc7,
	0xad, 0x33, 0x7c, 0xf6, 0x80, 0xbe, 0x55, 0x74, 0x54, 0x3b, 0x85, 0x5e, 0x8e, 0x3a, 0x31, 0x95,
	0x82, 0xd9, 0xa4, 0xf3, 0x28, 0x67, 0xbf, 0x41, 0x9d, 0xc4, 0x46, 0x4e, 0xe7, 0xe0, 0xa7, 0x1a,
	0x73, 0xc5, 0xc8, 0xe7, 0x22, 0x5a, 0x17, 0x9d, 0xc3, 0xc0, 0xa9, 0xcc, 0x25, 0xbc, 0x77, 0xb2,
	0xd9, 0xd1, 0x58, 0xc2, 0x07, 0x0f, 0x46, 0x87, 0x54, 0x47, 0xaf, 0x5f, 0xe0, 0xef, 0x92, 0x6c,
	0x8b, 0xfb, 0xb7, 0xb2, 0x43, 0x8d, 0x72, 0x5d, 0x95, 0xb8, 0x7f, 0x1d, 0x3b, 0xd0, 0x33, 0x57,
	0xa6, 0x67, 0xca, 0x7c, 0x20, 0x8a, 0xd5, 0x87, 0x8f, 0x07, 0x39, 0xea, 0xcb, 0xa1, 0x7f, 0x60,
	0x50, 0x0a, 0x95, 0xea, 0x54, 0x14, 0x26, 0x8c, 0x1f, 0x37, 0x33, 0xfd, 0x07, 0x41, 0x26, 0xd6,
	0x49, 0x86, 0xcb, 0xb5, 0xe0, 0x2e, 0x17, 0x58, 0xe8, 0x42, 0x70, 0xa4, 0x0c, 0xbe, 0xee, 0x50,
	0xaa, 0xda, 0x4b, 0x8c, 0xd7, 0x8d, 0xf5, 0x5a, 0x89, 0x05, 0x47, 0x89, 0xdc, 0x64, 0x24, 0x71,
	0x33, 0xd3, 0xbf, 0x30, 0xe4, 0xa9, 0x2a, 0xb3, 0xa4, 0x42, 0x6e, 0xfe, 0x30, 0x24, 0x6e, 0x81,
	0x9a, 0xb5, 0x0d, 0x38, 0x72, 0xd6, 0xb7, 0x6c, 0x03, 0x2c, 0xfc, 0x3b, 0x92, 0x94, 0xe9, 0xaa,
	0x6f, 0xbe, 0x8f, 0x93, 0x97, 0x00, 0x00, 0x00, 0xff, 0xff, 0xd0, 0x09, 0x4a, 0x2f, 0x38, 0x03,
	0x00, 0x00,
}