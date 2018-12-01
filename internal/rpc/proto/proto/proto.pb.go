// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto.proto

package proto

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
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type GeneralRequest struct {
	Path                 string   `protobuf:"bytes,1,opt,name=path,proto3" json:"path,omitempty"`
	ProtoID              string   `protobuf:"bytes,2,opt,name=protoID,proto3" json:"protoID,omitempty"`
	Data                 []byte   `protobuf:"bytes,3,opt,name=data,proto3" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GeneralRequest) Reset()         { *m = GeneralRequest{} }
func (m *GeneralRequest) String() string { return proto.CompactTextString(m) }
func (*GeneralRequest) ProtoMessage()    {}
func (*GeneralRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_2fcc84b9998d60d8, []int{0}
}

func (m *GeneralRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GeneralRequest.Unmarshal(m, b)
}
func (m *GeneralRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GeneralRequest.Marshal(b, m, deterministic)
}
func (m *GeneralRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GeneralRequest.Merge(m, src)
}
func (m *GeneralRequest) XXX_Size() int {
	return xxx_messageInfo_GeneralRequest.Size(m)
}
func (m *GeneralRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GeneralRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GeneralRequest proto.InternalMessageInfo

func (m *GeneralRequest) GetPath() string {
	if m != nil {
		return m.Path
	}
	return ""
}

func (m *GeneralRequest) GetProtoID() string {
	if m != nil {
		return m.ProtoID
	}
	return ""
}

func (m *GeneralRequest) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

type GeneralResponse struct {
	Message              string   `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GeneralResponse) Reset()         { *m = GeneralResponse{} }
func (m *GeneralResponse) String() string { return proto.CompactTextString(m) }
func (*GeneralResponse) ProtoMessage()    {}
func (*GeneralResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_2fcc84b9998d60d8, []int{1}
}

func (m *GeneralResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GeneralResponse.Unmarshal(m, b)
}
func (m *GeneralResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GeneralResponse.Marshal(b, m, deterministic)
}
func (m *GeneralResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GeneralResponse.Merge(m, src)
}
func (m *GeneralResponse) XXX_Size() int {
	return xxx_messageInfo_GeneralResponse.Size(m)
}
func (m *GeneralResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GeneralResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GeneralResponse proto.InternalMessageInfo

func (m *GeneralResponse) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func init() {
	proto.RegisterType((*GeneralRequest)(nil), "proto.GeneralRequest")
	proto.RegisterType((*GeneralResponse)(nil), "proto.GeneralResponse")
}

func init() { proto.RegisterFile("proto.proto", fileDescriptor_2fcc84b9998d60d8) }

var fileDescriptor_2fcc84b9998d60d8 = []byte{
	// 205 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x2e, 0x28, 0xca, 0x2f,
	0xc9, 0xd7, 0x03, 0x93, 0x42, 0xac, 0x60, 0x4a, 0x29, 0x88, 0x8b, 0xcf, 0x3d, 0x35, 0x2f, 0xb5,
	0x28, 0x31, 0x27, 0x28, 0xb5, 0xb0, 0x34, 0xb5, 0xb8, 0x44, 0x48, 0x88, 0x8b, 0xa5, 0x20, 0xb1,
	0x24, 0x43, 0x82, 0x51, 0x81, 0x51, 0x83, 0x33, 0x08, 0xcc, 0x16, 0x92, 0xe0, 0x62, 0x07, 0x2b,
	0xf7, 0x74, 0x91, 0x60, 0x02, 0x0b, 0xc3, 0xb8, 0x20, 0xd5, 0x29, 0x89, 0x25, 0x89, 0x12, 0xcc,
	0x0a, 0x8c, 0x1a, 0x3c, 0x41, 0x60, 0xb6, 0x92, 0x36, 0x17, 0x3f, 0xdc, 0xcc, 0xe2, 0x82, 0xfc,
	0xbc, 0xe2, 0x54, 0x90, 0x01, 0xb9, 0xa9, 0xc5, 0xc5, 0x89, 0xe9, 0xa9, 0x50, 0x73, 0x61, 0x5c,
	0xa3, 0x6b, 0x8c, 0x5c, 0xac, 0x01, 0x60, 0x17, 0x39, 0x73, 0x09, 0xf8, 0xa5, 0x96, 0x83, 0xd9,
	0x49, 0xa5, 0x69, 0xee, 0xa5, 0x99, 0x29, 0xa9, 0x42, 0xa2, 0x10, 0xd7, 0xea, 0xa1, 0xba, 0x51,
	0x4a, 0x0c, 0x5d, 0x18, 0x62, 0x8d, 0x12, 0x83, 0x90, 0x1b, 0x97, 0x70, 0x50, 0x6a, 0x62, 0x0a,
	0x58, 0xb7, 0x5b, 0x51, 0x7e, 0xae, 0x6f, 0x6a, 0x6e, 0x7e, 0x51, 0x25, 0xe9, 0xe6, 0x38, 0x70,
	0xf1, 0x86, 0x17, 0x65, 0x96, 0xa4, 0x86, 0xe4, 0x93, 0x69, 0x42, 0x12, 0x1b, 0x58, 0xc2, 0x18,
	0x10, 0x00, 0x00, 0xff, 0xff, 0xe9, 0x12, 0x3e, 0xd4, 0x76, 0x01, 0x00, 0x00,
}