// Code generated by protoc-gen-go. DO NOT EDIT.
// source: api.proto

package proto_server

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

type ReloadReq struct {
	Name                 string   `protobuf:"bytes,1,opt,name=Name,proto3" json:"Name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ReloadReq) Reset()         { *m = ReloadReq{} }
func (m *ReloadReq) String() string { return proto.CompactTextString(m) }
func (*ReloadReq) ProtoMessage()    {}
func (*ReloadReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_00212fb1f9d3bf1c, []int{0}
}

func (m *ReloadReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReloadReq.Unmarshal(m, b)
}
func (m *ReloadReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReloadReq.Marshal(b, m, deterministic)
}
func (m *ReloadReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReloadReq.Merge(m, src)
}
func (m *ReloadReq) XXX_Size() int {
	return xxx_messageInfo_ReloadReq.Size(m)
}
func (m *ReloadReq) XXX_DiscardUnknown() {
	xxx_messageInfo_ReloadReq.DiscardUnknown(m)
}

var xxx_messageInfo_ReloadReq proto.InternalMessageInfo

func (m *ReloadReq) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type ReloadRep struct {
	Ok                   bool     `protobuf:"varint,1,opt,name=Ok,proto3" json:"Ok,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ReloadRep) Reset()         { *m = ReloadRep{} }
func (m *ReloadRep) String() string { return proto.CompactTextString(m) }
func (*ReloadRep) ProtoMessage()    {}
func (*ReloadRep) Descriptor() ([]byte, []int) {
	return fileDescriptor_00212fb1f9d3bf1c, []int{1}
}

func (m *ReloadRep) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReloadRep.Unmarshal(m, b)
}
func (m *ReloadRep) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReloadRep.Marshal(b, m, deterministic)
}
func (m *ReloadRep) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReloadRep.Merge(m, src)
}
func (m *ReloadRep) XXX_Size() int {
	return xxx_messageInfo_ReloadRep.Size(m)
}
func (m *ReloadRep) XXX_DiscardUnknown() {
	xxx_messageInfo_ReloadRep.DiscardUnknown(m)
}

var xxx_messageInfo_ReloadRep proto.InternalMessageInfo

func (m *ReloadRep) GetOk() bool {
	if m != nil {
		return m.Ok
	}
	return false
}

func init() {
	proto.RegisterType((*ReloadReq)(nil), "ReloadReq")
	proto.RegisterType((*ReloadRep)(nil), "ReloadRep")
}

func init() { proto.RegisterFile("api.proto", fileDescriptor_00212fb1f9d3bf1c) }

var fileDescriptor_00212fb1f9d3bf1c = []byte{
	// 118 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x4c, 0x2c, 0xc8, 0xd4,
	0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x57, 0x92, 0xe7, 0xe2, 0x0c, 0x4a, 0xcd, 0xc9, 0x4f, 0x4c, 0x09,
	0x4a, 0x2d, 0x14, 0x12, 0xe2, 0x62, 0xf1, 0x4b, 0xcc, 0x4d, 0x95, 0x60, 0x54, 0x60, 0xd4, 0xe0,
	0x0c, 0x02, 0xb3, 0x95, 0xa4, 0x11, 0x0a, 0x0a, 0x84, 0xf8, 0xb8, 0x98, 0xfc, 0xb3, 0xc1, 0xd2,
	0x1c, 0x41, 0x4c, 0xfe, 0xd9, 0x46, 0xfa, 0x5c, 0x9c, 0x8e, 0x05, 0x99, 0xc1, 0xa9, 0x45, 0x65,
	0xa9, 0x45, 0x42, 0x4a, 0x5c, 0x6c, 0x10, 0x95, 0x42, 0x5c, 0x7a, 0x70, 0x33, 0xa5, 0x10, 0xec,
	0x02, 0x25, 0x86, 0x24, 0x36, 0xb0, 0xad, 0xc6, 0x80, 0x00, 0x00, 0x00, 0xff, 0xff, 0x91, 0x9f,
	0x23, 0x4c, 0x82, 0x00, 0x00, 0x00,
}
