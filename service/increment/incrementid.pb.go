// Code generated by protoc-gen-go.
// source: incrementid.proto
// DO NOT EDIT!

/*
Package increment is a generated protocol buffer package.

It is generated from these files:
	incrementid.proto

It has these top-level messages:
	GetIncrIdRequest
	GetIncrIdReply
	CheckIncrKeyExistReply
	CreateIncrKeyRequest
	NoneReply
*/
package increment

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
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

type GetIncrIdRequest struct {
	Name string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
}

func (m *GetIncrIdRequest) Reset()                    { *m = GetIncrIdRequest{} }
func (m *GetIncrIdRequest) String() string            { return proto.CompactTextString(m) }
func (*GetIncrIdRequest) ProtoMessage()               {}
func (*GetIncrIdRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type GetIncrIdReply struct {
	IncId uint64 `protobuf:"varint,1,opt,name=inc_id,json=incId" json:"inc_id,omitempty"`
}

func (m *GetIncrIdReply) Reset()                    { *m = GetIncrIdReply{} }
func (m *GetIncrIdReply) String() string            { return proto.CompactTextString(m) }
func (*GetIncrIdReply) ProtoMessage()               {}
func (*GetIncrIdReply) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

type CheckIncrKeyExistReply struct {
	Exist bool `protobuf:"varint,1,opt,name=exist" json:"exist,omitempty"`
}

func (m *CheckIncrKeyExistReply) Reset()                    { *m = CheckIncrKeyExistReply{} }
func (m *CheckIncrKeyExistReply) String() string            { return proto.CompactTextString(m) }
func (*CheckIncrKeyExistReply) ProtoMessage()               {}
func (*CheckIncrKeyExistReply) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

type CreateIncrKeyRequest struct {
	Name         string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	InitialValue uint64 `protobuf:"varint,2,opt,name=initial_value,json=initialValue" json:"initial_value,omitempty"`
}

func (m *CreateIncrKeyRequest) Reset()                    { *m = CreateIncrKeyRequest{} }
func (m *CreateIncrKeyRequest) String() string            { return proto.CompactTextString(m) }
func (*CreateIncrKeyRequest) ProtoMessage()               {}
func (*CreateIncrKeyRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

type NoneReply struct {
}

func (m *NoneReply) Reset()                    { *m = NoneReply{} }
func (m *NoneReply) String() string            { return proto.CompactTextString(m) }
func (*NoneReply) ProtoMessage()               {}
func (*NoneReply) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func init() {
	proto.RegisterType((*GetIncrIdRequest)(nil), "increment.GetIncrIdRequest")
	proto.RegisterType((*GetIncrIdReply)(nil), "increment.GetIncrIdReply")
	proto.RegisterType((*CheckIncrKeyExistReply)(nil), "increment.CheckIncrKeyExistReply")
	proto.RegisterType((*CreateIncrKeyRequest)(nil), "increment.CreateIncrKeyRequest")
	proto.RegisterType((*NoneReply)(nil), "increment.NoneReply")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion3

// Client API for IncrementId service

type IncrementIdClient interface {
	GetIncrId(ctx context.Context, in *GetIncrIdRequest, opts ...grpc.CallOption) (*GetIncrIdReply, error)
	CheckIncrKeyExist(ctx context.Context, in *GetIncrIdRequest, opts ...grpc.CallOption) (*CheckIncrKeyExistReply, error)
	CreateIncrKey(ctx context.Context, in *CreateIncrKeyRequest, opts ...grpc.CallOption) (*NoneReply, error)
}

type incrementIdClient struct {
	cc *grpc.ClientConn
}

func NewIncrementIdClient(cc *grpc.ClientConn) IncrementIdClient {
	return &incrementIdClient{cc}
}

func (c *incrementIdClient) GetIncrId(ctx context.Context, in *GetIncrIdRequest, opts ...grpc.CallOption) (*GetIncrIdReply, error) {
	out := new(GetIncrIdReply)
	err := grpc.Invoke(ctx, "/increment.IncrementId/GetIncrId", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *incrementIdClient) CheckIncrKeyExist(ctx context.Context, in *GetIncrIdRequest, opts ...grpc.CallOption) (*CheckIncrKeyExistReply, error) {
	out := new(CheckIncrKeyExistReply)
	err := grpc.Invoke(ctx, "/increment.IncrementId/CheckIncrKeyExist", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *incrementIdClient) CreateIncrKey(ctx context.Context, in *CreateIncrKeyRequest, opts ...grpc.CallOption) (*NoneReply, error) {
	out := new(NoneReply)
	err := grpc.Invoke(ctx, "/increment.IncrementId/CreateIncrKey", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for IncrementId service

type IncrementIdServer interface {
	GetIncrId(context.Context, *GetIncrIdRequest) (*GetIncrIdReply, error)
	CheckIncrKeyExist(context.Context, *GetIncrIdRequest) (*CheckIncrKeyExistReply, error)
	CreateIncrKey(context.Context, *CreateIncrKeyRequest) (*NoneReply, error)
}

func RegisterIncrementIdServer(s *grpc.Server, srv IncrementIdServer) {
	s.RegisterService(&_IncrementId_serviceDesc, srv)
}

func _IncrementId_GetIncrId_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetIncrIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IncrementIdServer).GetIncrId(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/increment.IncrementId/GetIncrId",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IncrementIdServer).GetIncrId(ctx, req.(*GetIncrIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _IncrementId_CheckIncrKeyExist_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetIncrIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IncrementIdServer).CheckIncrKeyExist(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/increment.IncrementId/CheckIncrKeyExist",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IncrementIdServer).CheckIncrKeyExist(ctx, req.(*GetIncrIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _IncrementId_CreateIncrKey_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateIncrKeyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IncrementIdServer).CreateIncrKey(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/increment.IncrementId/CreateIncrKey",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IncrementIdServer).CreateIncrKey(ctx, req.(*CreateIncrKeyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _IncrementId_serviceDesc = grpc.ServiceDesc{
	ServiceName: "increment.IncrementId",
	HandlerType: (*IncrementIdServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetIncrId",
			Handler:    _IncrementId_GetIncrId_Handler,
		},
		{
			MethodName: "CheckIncrKeyExist",
			Handler:    _IncrementId_CheckIncrKeyExist_Handler,
		},
		{
			MethodName: "CreateIncrKey",
			Handler:    _IncrementId_CreateIncrKey_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: fileDescriptor0,
}

func init() { proto.RegisterFile("incrementid.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 271 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x7c, 0x91, 0x41, 0x4b, 0xc3, 0x40,
	0x10, 0x85, 0x13, 0x69, 0x8b, 0x99, 0x5a, 0xb1, 0x43, 0x94, 0x5a, 0x0f, 0xea, 0x0a, 0xea, 0x29,
	0x07, 0xfd, 0x09, 0xa5, 0xe8, 0x22, 0x28, 0x04, 0xf4, 0x5a, 0xe2, 0xee, 0x80, 0x83, 0xe9, 0x26,
	0xa6, 0x5b, 0x31, 0x3f, 0xdc, 0xbb, 0x64, 0x5b, 0x43, 0xb4, 0xa1, 0xb7, 0x9d, 0x99, 0xf7, 0x76,
	0xf6, 0x7d, 0x0b, 0x43, 0x36, 0xaa, 0xa0, 0x39, 0x19, 0xcb, 0x3a, 0xca, 0x8b, 0xcc, 0x66, 0x18,
	0xd4, 0x2d, 0x71, 0x09, 0x07, 0x77, 0x64, 0xa5, 0x51, 0x85, 0xd4, 0x31, 0x7d, 0x2c, 0x69, 0x61,
	0x11, 0xa1, 0x63, 0x92, 0x39, 0x8d, 0xfc, 0x33, 0xff, 0x3a, 0x88, 0xdd, 0x59, 0x5c, 0xc1, 0x7e,
	0x43, 0x97, 0xa7, 0x25, 0x1e, 0x42, 0x8f, 0x8d, 0x9a, 0xb1, 0x76, 0xba, 0x4e, 0xdc, 0x65, 0xa3,
	0xa4, 0x16, 0x11, 0x1c, 0x4d, 0xde, 0x48, 0xbd, 0x57, 0xd2, 0x07, 0x2a, 0xa7, 0x5f, 0xbc, 0xb0,
	0x2b, 0x43, 0x08, 0x5d, 0xaa, 0x2a, 0xa7, 0xdf, 0x8d, 0x57, 0x85, 0x78, 0x82, 0x70, 0x52, 0x50,
	0x62, 0x69, 0x6d, 0xd8, 0xf2, 0x08, 0xbc, 0x80, 0x01, 0x1b, 0xb6, 0x9c, 0xa4, 0xb3, 0xcf, 0x24,
	0x5d, 0xd2, 0x68, 0xc7, 0x6d, 0xde, 0x5b, 0x37, 0x5f, 0xaa, 0x9e, 0xe8, 0x43, 0xf0, 0x98, 0x19,
	0x72, 0x3b, 0x6f, 0xbe, 0x7d, 0xe8, 0xcb, 0xdf, 0xb0, 0x52, 0xe3, 0x14, 0x82, 0x3a, 0x06, 0x9e,
	0x44, 0x35, 0x87, 0xe8, 0x3f, 0x84, 0xf1, 0x71, 0xfb, 0x30, 0x4f, 0x4b, 0xe1, 0xe1, 0x33, 0x0c,
	0x37, 0x42, 0x6e, 0xbf, 0xee, 0xbc, 0x31, 0x6c, 0xe7, 0x23, 0x3c, 0xbc, 0x87, 0xc1, 0x1f, 0x16,
	0x78, 0xda, 0x74, 0xb5, 0x50, 0x1a, 0x87, 0x0d, 0x41, 0x9d, 0x5a, 0x78, 0xaf, 0x3d, 0xf7, 0xd1,
	0xb7, 0x3f, 0x01, 0x00, 0x00, 0xff, 0xff, 0x2e, 0xaf, 0xb9, 0x6b, 0xfd, 0x01, 0x00, 0x00,
}