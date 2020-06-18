// Code generated by protoc-gen-go. DO NOT EDIT.
// source: cache/cache_service.proto

package cache

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
	wrappers "github.com/golang/protobuf/ptypes/wrappers"
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

type CacheMessage struct {
	Key                  []byte               `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	Value                []byte               `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
	Expiration           *timestamp.Timestamp `protobuf:"bytes,3,opt,name=expiration,proto3" json:"expiration,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *CacheMessage) Reset()         { *m = CacheMessage{} }
func (m *CacheMessage) String() string { return proto.CompactTextString(m) }
func (*CacheMessage) ProtoMessage()    {}
func (*CacheMessage) Descriptor() ([]byte, []int) {
	return fileDescriptor_7933906ce719ddba, []int{0}
}

func (m *CacheMessage) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CacheMessage.Unmarshal(m, b)
}
func (m *CacheMessage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CacheMessage.Marshal(b, m, deterministic)
}
func (m *CacheMessage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CacheMessage.Merge(m, src)
}
func (m *CacheMessage) XXX_Size() int {
	return xxx_messageInfo_CacheMessage.Size(m)
}
func (m *CacheMessage) XXX_DiscardUnknown() {
	xxx_messageInfo_CacheMessage.DiscardUnknown(m)
}

var xxx_messageInfo_CacheMessage proto.InternalMessageInfo

func (m *CacheMessage) GetKey() []byte {
	if m != nil {
		return m.Key
	}
	return nil
}

func (m *CacheMessage) GetValue() []byte {
	if m != nil {
		return m.Value
	}
	return nil
}

func (m *CacheMessage) GetExpiration() *timestamp.Timestamp {
	if m != nil {
		return m.Expiration
	}
	return nil
}

type CacheListRequest struct {
	Limit                *wrappers.UInt32Value `protobuf:"bytes,1,opt,name=limit,proto3" json:"limit,omitempty"`
	Offset               *wrappers.UInt32Value `protobuf:"bytes,2,opt,name=offset,proto3" json:"offset,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *CacheListRequest) Reset()         { *m = CacheListRequest{} }
func (m *CacheListRequest) String() string { return proto.CompactTextString(m) }
func (*CacheListRequest) ProtoMessage()    {}
func (*CacheListRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_7933906ce719ddba, []int{1}
}

func (m *CacheListRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CacheListRequest.Unmarshal(m, b)
}
func (m *CacheListRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CacheListRequest.Marshal(b, m, deterministic)
}
func (m *CacheListRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CacheListRequest.Merge(m, src)
}
func (m *CacheListRequest) XXX_Size() int {
	return xxx_messageInfo_CacheListRequest.Size(m)
}
func (m *CacheListRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CacheListRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CacheListRequest proto.InternalMessageInfo

func (m *CacheListRequest) GetLimit() *wrappers.UInt32Value {
	if m != nil {
		return m.Limit
	}
	return nil
}

func (m *CacheListRequest) GetOffset() *wrappers.UInt32Value {
	if m != nil {
		return m.Offset
	}
	return nil
}

type CacheListResponse struct {
	Items                []*CacheMessage `protobuf:"bytes,1,rep,name=items,proto3" json:"items,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *CacheListResponse) Reset()         { *m = CacheListResponse{} }
func (m *CacheListResponse) String() string { return proto.CompactTextString(m) }
func (*CacheListResponse) ProtoMessage()    {}
func (*CacheListResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_7933906ce719ddba, []int{2}
}

func (m *CacheListResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CacheListResponse.Unmarshal(m, b)
}
func (m *CacheListResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CacheListResponse.Marshal(b, m, deterministic)
}
func (m *CacheListResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CacheListResponse.Merge(m, src)
}
func (m *CacheListResponse) XXX_Size() int {
	return xxx_messageInfo_CacheListResponse.Size(m)
}
func (m *CacheListResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CacheListResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CacheListResponse proto.InternalMessageInfo

func (m *CacheListResponse) GetItems() []*CacheMessage {
	if m != nil {
		return m.Items
	}
	return nil
}

func init() {
	proto.RegisterType((*CacheMessage)(nil), "cache.v1.CacheMessage")
	proto.RegisterType((*CacheListRequest)(nil), "cache.v1.CacheListRequest")
	proto.RegisterType((*CacheListResponse)(nil), "cache.v1.CacheListResponse")
}

func init() { proto.RegisterFile("cache/cache_service.proto", fileDescriptor_7933906ce719ddba) }

var fileDescriptor_7933906ce719ddba = []byte{
	// 449 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x53, 0xdd, 0x6a, 0xd4, 0x40,
	0x14, 0x26, 0x8d, 0xbb, 0xe8, 0xd9, 0x82, 0xdd, 0xb1, 0xc8, 0x76, 0x5a, 0x74, 0xc9, 0xd5, 0x52,
	0x74, 0x62, 0x53, 0xaf, 0x7a, 0x67, 0x15, 0x44, 0xa8, 0x20, 0xd1, 0x8a, 0x78, 0x23, 0xb3, 0xcb,
	0xd9, 0xdd, 0xb1, 0x49, 0x26, 0x66, 0x4e, 0xa2, 0x8b, 0xf6, 0xc6, 0x57, 0xf0, 0x49, 0x7c, 0x16,
	0x5f, 0xc1, 0x07, 0x91, 0xcc, 0x24, 0x1a, 0xfa, 0xa3, 0x60, 0x6f, 0x42, 0xe6, 0x3b, 0xdf, 0x7c,
	0xdf, 0xe1, 0x9b, 0x73, 0x60, 0x6b, 0x26, 0x67, 0x4b, 0x0c, 0xed, 0xf7, 0x9d, 0xc1, 0xa2, 0x52,
	0x33, 0x14, 0x79, 0xa1, 0x49, 0xb3, 0xeb, 0x16, 0x14, 0xd5, 0x1e, 0xbf, 0xbb, 0xd0, 0x7a, 0x91,
	0x60, 0x68, 0xf1, 0x69, 0x39, 0x0f, 0x49, 0xa5, 0x68, 0x48, 0xa6, 0xb9, 0xa3, 0xf2, 0x3b, 0x67,
	0x09, 0x1f, 0x0b, 0x99, 0xe7, 0x58, 0x98, 0xa6, 0xbe, 0xd3, 0xd4, 0x65, 0xae, 0x42, 0x99, 0x65,
	0x9a, 0x24, 0x29, 0x9d, 0x35, 0xd5, 0xa0, 0x80, 0xf5, 0xc7, 0xb5, 0xd5, 0x73, 0x34, 0x46, 0x2e,
	0x90, 0x6d, 0x80, 0x7f, 0x82, 0xab, 0x91, 0x37, 0xf6, 0x26, 0xeb, 0x71, 0xfd, 0xcb, 0x36, 0xa1,
	0x57, 0xc9, 0xa4, 0xc4, 0xd1, 0x9a, 0xc5, 0xdc, 0x81, 0x1d, 0x00, 0xe0, 0xa7, 0x5c, 0x15, 0x56,
	0x6c, 0xe4, 0x8f, 0xbd, 0xc9, 0x20, 0xe2, 0xc2, 0x59, 0x89, 0xb6, 0x15, 0xf1, 0xaa, 0xed, 0x35,
	0xee, 0xb0, 0x83, 0x2f, 0xb0, 0x61, 0x3d, 0x8f, 0x94, 0xa1, 0x18, 0x3f, 0x94, 0x68, 0x88, 0x45,
	0xd0, 0x4b, 0x54, 0xaa, 0xc8, 0x3a, 0x0f, 0xa2, 0x9d, 0x73, 0x52, 0xc7, 0xcf, 0x32, 0xda, 0x8f,
	0x5e, 0xd7, 0xe6, 0xb1, 0xa3, 0xb2, 0x87, 0xd0, 0xd7, 0xf3, 0xb9, 0x41, 0xb2, 0xad, 0xfd, 0xeb,
	0x52, 0xc3, 0x0d, 0x1e, 0xc1, 0xb0, 0xe3, 0x6e, 0x72, 0x9d, 0x19, 0x64, 0xf7, 0xa0, 0xa7, 0x08,
	0x53, 0x33, 0xf2, 0xc6, 0xfe, 0x64, 0x10, 0xdd, 0x16, 0x6d, 0xfe, 0xa2, 0x9b, 0x4e, 0xec, 0x48,
	0xd1, 0x77, 0x1f, 0x7a, 0x16, 0x67, 0x47, 0xe0, 0xbf, 0x28, 0x89, 0x5d, 0xc2, 0xe7, 0x7f, 0x49,
	0x24, 0x18, 0x7e, 0xfd, 0xf1, 0xf3, 0xdb, 0xda, 0x20, 0xb8, 0x11, 0x56, 0x7b, 0x6e, 0x06, 0xd8,
	0x31, 0xf8, 0x2f, 0xf1, 0xff, 0xd4, 0xb8, 0x55, 0xdb, 0xe4, 0x37, 0x7f, 0xab, 0x85, 0x9f, 0x4f,
	0x70, 0x75, 0x7a, 0xe0, 0xed, 0xd6, 0xb2, 0x4f, 0x91, 0xd8, 0xf6, 0xb9, 0xeb, 0x87, 0x2b, 0x42,
	0x63, 0xd3, 0xe1, 0x97, 0x78, 0x06, 0x5b, 0x56, 0xf7, 0x16, 0x1b, 0x76, 0x74, 0xed, 0x04, 0x9c,
	0xb2, 0x37, 0xd0, 0x7f, 0x82, 0x09, 0x12, 0x5e, 0x49, 0x79, 0xf7, 0x02, 0xe5, 0x18, 0xae, 0xd5,
	0xaf, 0xc3, 0xf8, 0x99, 0xab, 0x9d, 0x81, 0xe1, 0xdb, 0x17, 0xd6, 0xdc, 0x73, 0xb6, 0xd9, 0xb2,
	0x3f, 0xd9, 0x1e, 0x46, 0x6f, 0x1f, 0x2c, 0x14, 0x2d, 0xcb, 0xa9, 0x98, 0xe9, 0x34, 0x7c, 0xaf,
	0x97, 0xd9, 0x14, 0x93, 0x44, 0x67, 0xcd, 0xfe, 0xdd, 0x6f, 0xf6, 0xcf, 0xad, 0x91, 0xc3, 0xa6,
	0x7d, 0x7b, 0xd8, 0xff, 0x15, 0x00, 0x00, 0xff, 0xff, 0xab, 0xf5, 0xa2, 0x5e, 0xa9, 0x03, 0x00,
	0x00,
}