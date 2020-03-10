// Code generated by protoc-gen-go. DO NOT EDIT.
// source: uber/foo/v1/excited_api.proto

// Copyright (c) 2020 Uber Technologies, Inc.
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

package foov1

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

type ExclamationRequest struct {
	Value                string   `protobuf:"bytes,1,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ExclamationRequest) Reset()         { *m = ExclamationRequest{} }
func (m *ExclamationRequest) String() string { return proto.CompactTextString(m) }
func (*ExclamationRequest) ProtoMessage()    {}
func (*ExclamationRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_af997259370c9c75, []int{0}
}

func (m *ExclamationRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ExclamationRequest.Unmarshal(m, b)
}
func (m *ExclamationRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ExclamationRequest.Marshal(b, m, deterministic)
}
func (m *ExclamationRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ExclamationRequest.Merge(m, src)
}
func (m *ExclamationRequest) XXX_Size() int {
	return xxx_messageInfo_ExclamationRequest.Size(m)
}
func (m *ExclamationRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ExclamationRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ExclamationRequest proto.InternalMessageInfo

func (m *ExclamationRequest) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

type ExclamationResponse struct {
	Value                string   `protobuf:"bytes,1,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ExclamationResponse) Reset()         { *m = ExclamationResponse{} }
func (m *ExclamationResponse) String() string { return proto.CompactTextString(m) }
func (*ExclamationResponse) ProtoMessage()    {}
func (*ExclamationResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_af997259370c9c75, []int{1}
}

func (m *ExclamationResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ExclamationResponse.Unmarshal(m, b)
}
func (m *ExclamationResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ExclamationResponse.Marshal(b, m, deterministic)
}
func (m *ExclamationResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ExclamationResponse.Merge(m, src)
}
func (m *ExclamationResponse) XXX_Size() int {
	return xxx_messageInfo_ExclamationResponse.Size(m)
}
func (m *ExclamationResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ExclamationResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ExclamationResponse proto.InternalMessageInfo

func (m *ExclamationResponse) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

func init() {
	proto.RegisterType((*ExclamationRequest)(nil), "uber.foo.v1.ExclamationRequest")
	proto.RegisterType((*ExclamationResponse)(nil), "uber.foo.v1.ExclamationResponse")
}

func init() {
	proto.RegisterFile("uber/foo/v1/excited_api.proto", fileDescriptor_af997259370c9c75)
}

var fileDescriptor_af997259370c9c75 = []byte{
	// 262 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0x2d, 0x4d, 0x4a, 0x2d,
	0xd2, 0x4f, 0xcb, 0xcf, 0xd7, 0x2f, 0x33, 0xd4, 0x4f, 0xad, 0x48, 0xce, 0x2c, 0x49, 0x4d, 0x89,
	0x4f, 0x2c, 0xc8, 0xd4, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x06, 0x49, 0xeb, 0xa5, 0xe5,
	0xe7, 0xeb, 0x95, 0x19, 0x2a, 0x69, 0x71, 0x09, 0xb9, 0x56, 0x24, 0xe7, 0x24, 0xe6, 0x26, 0x96,
	0x64, 0xe6, 0xe7, 0x05, 0xa5, 0x16, 0x96, 0xa6, 0x16, 0x97, 0x08, 0x89, 0x70, 0xb1, 0x96, 0x25,
	0xe6, 0x94, 0xa6, 0x4a, 0x30, 0x2a, 0x30, 0x6a, 0x70, 0x06, 0x41, 0x38, 0x4a, 0xda, 0x5c, 0xc2,
	0x28, 0x6a, 0x8b, 0x0b, 0xf2, 0xf3, 0x8a, 0x53, 0xb1, 0x2b, 0x36, 0xfa, 0xc7, 0xc4, 0xc5, 0xe5,
	0x0a, 0xb1, 0xdb, 0x31, 0xc0, 0x53, 0x28, 0x80, 0x8b, 0x1b, 0x49, 0xaf, 0x90, 0xbc, 0x1e, 0x92,
	0x23, 0xf4, 0x30, 0x5d, 0x20, 0xa5, 0x80, 0x5b, 0x01, 0xd4, 0xda, 0x38, 0x2e, 0x71, 0x24, 0x61,
	0xe7, 0x9c, 0xcc, 0xd4, 0xbc, 0x92, 0xe0, 0x92, 0xa2, 0xd4, 0xc4, 0x5c, 0x2a, 0x98, 0xae, 0xc1,
	0x88, 0x66, 0x7e, 0x70, 0x6a, 0x51, 0x59, 0x6a, 0x11, 0xd5, 0xcc, 0x37, 0x00, 0x99, 0x2f, 0x8a,
	0x24, 0xe1, 0x94, 0x99, 0x92, 0x49, 0x45, 0xd7, 0x1b, 0x30, 0x3a, 0xf9, 0x72, 0xf1, 0x27, 0xe7,
	0xe7, 0x22, 0x2b, 0x75, 0xe2, 0x87, 0x45, 0x48, 0x41, 0x66, 0x00, 0x28, 0x29, 0x04, 0x30, 0x46,
	0xb1, 0xa6, 0xe5, 0xe7, 0x97, 0x19, 0x2e, 0x62, 0x62, 0x0e, 0x75, 0x8b, 0x58, 0xc5, 0xc4, 0x1d,
	0x0a, 0x52, 0xed, 0x96, 0x9f, 0xaf, 0x17, 0x66, 0x78, 0x0a, 0xc2, 0x8b, 0x71, 0xcb, 0xcf, 0x8f,
	0x09, 0x33, 0x4c, 0x62, 0x03, 0x27, 0x1e, 0x63, 0x40, 0x00, 0x00, 0x00, 0xff, 0xff, 0x9b, 0xc7,
	0x22, 0x83, 0x5d, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// ExcitedAPIClient is the client API for ExcitedAPI service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ExcitedAPIClient interface {
	// Exclamation adds an exclamation to the request value.
	Exclamation(ctx context.Context, in *ExclamationRequest, opts ...grpc.CallOption) (*ExclamationResponse, error)
	// ExclamationClientStream adds an exclamation to the combined request values.
	ExclamationClientStream(ctx context.Context, opts ...grpc.CallOption) (ExcitedAPI_ExclamationClientStreamClient, error)
	// ExclamationServerStream adds an exclamation to the request value
	// and streams each character as a response.
	ExclamationServerStream(ctx context.Context, in *ExclamationRequest, opts ...grpc.CallOption) (ExcitedAPI_ExclamationServerStreamClient, error)
	// ExclamationBidiStream adds an exclamation to the each request value.
	ExclamationBidiStream(ctx context.Context, opts ...grpc.CallOption) (ExcitedAPI_ExclamationBidiStreamClient, error)
}

type excitedAPIClient struct {
	cc grpc.ClientConnInterface
}

func NewExcitedAPIClient(cc grpc.ClientConnInterface) ExcitedAPIClient {
	return &excitedAPIClient{cc}
}

func (c *excitedAPIClient) Exclamation(ctx context.Context, in *ExclamationRequest, opts ...grpc.CallOption) (*ExclamationResponse, error) {
	out := new(ExclamationResponse)
	err := c.cc.Invoke(ctx, "/uber.foo.v1.ExcitedAPI/Exclamation", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *excitedAPIClient) ExclamationClientStream(ctx context.Context, opts ...grpc.CallOption) (ExcitedAPI_ExclamationClientStreamClient, error) {
	stream, err := c.cc.NewStream(ctx, &_ExcitedAPI_serviceDesc.Streams[0], "/uber.foo.v1.ExcitedAPI/ExclamationClientStream", opts...)
	if err != nil {
		return nil, err
	}
	x := &excitedAPIExclamationClientStreamClient{stream}
	return x, nil
}

type ExcitedAPI_ExclamationClientStreamClient interface {
	Send(*ExclamationRequest) error
	CloseAndRecv() (*ExclamationResponse, error)
	grpc.ClientStream
}

type excitedAPIExclamationClientStreamClient struct {
	grpc.ClientStream
}

func (x *excitedAPIExclamationClientStreamClient) Send(m *ExclamationRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *excitedAPIExclamationClientStreamClient) CloseAndRecv() (*ExclamationResponse, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(ExclamationResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *excitedAPIClient) ExclamationServerStream(ctx context.Context, in *ExclamationRequest, opts ...grpc.CallOption) (ExcitedAPI_ExclamationServerStreamClient, error) {
	stream, err := c.cc.NewStream(ctx, &_ExcitedAPI_serviceDesc.Streams[1], "/uber.foo.v1.ExcitedAPI/ExclamationServerStream", opts...)
	if err != nil {
		return nil, err
	}
	x := &excitedAPIExclamationServerStreamClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type ExcitedAPI_ExclamationServerStreamClient interface {
	Recv() (*ExclamationResponse, error)
	grpc.ClientStream
}

type excitedAPIExclamationServerStreamClient struct {
	grpc.ClientStream
}

func (x *excitedAPIExclamationServerStreamClient) Recv() (*ExclamationResponse, error) {
	m := new(ExclamationResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *excitedAPIClient) ExclamationBidiStream(ctx context.Context, opts ...grpc.CallOption) (ExcitedAPI_ExclamationBidiStreamClient, error) {
	stream, err := c.cc.NewStream(ctx, &_ExcitedAPI_serviceDesc.Streams[2], "/uber.foo.v1.ExcitedAPI/ExclamationBidiStream", opts...)
	if err != nil {
		return nil, err
	}
	x := &excitedAPIExclamationBidiStreamClient{stream}
	return x, nil
}

type ExcitedAPI_ExclamationBidiStreamClient interface {
	Send(*ExclamationRequest) error
	Recv() (*ExclamationResponse, error)
	grpc.ClientStream
}

type excitedAPIExclamationBidiStreamClient struct {
	grpc.ClientStream
}

func (x *excitedAPIExclamationBidiStreamClient) Send(m *ExclamationRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *excitedAPIExclamationBidiStreamClient) Recv() (*ExclamationResponse, error) {
	m := new(ExclamationResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// ExcitedAPIServer is the server API for ExcitedAPI service.
type ExcitedAPIServer interface {
	// Exclamation adds an exclamation to the request value.
	Exclamation(context.Context, *ExclamationRequest) (*ExclamationResponse, error)
	// ExclamationClientStream adds an exclamation to the combined request values.
	ExclamationClientStream(ExcitedAPI_ExclamationClientStreamServer) error
	// ExclamationServerStream adds an exclamation to the request value
	// and streams each character as a response.
	ExclamationServerStream(*ExclamationRequest, ExcitedAPI_ExclamationServerStreamServer) error
	// ExclamationBidiStream adds an exclamation to the each request value.
	ExclamationBidiStream(ExcitedAPI_ExclamationBidiStreamServer) error
}

// UnimplementedExcitedAPIServer can be embedded to have forward compatible implementations.
type UnimplementedExcitedAPIServer struct {
}

func (*UnimplementedExcitedAPIServer) Exclamation(ctx context.Context, req *ExclamationRequest) (*ExclamationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Exclamation not implemented")
}
func (*UnimplementedExcitedAPIServer) ExclamationClientStream(srv ExcitedAPI_ExclamationClientStreamServer) error {
	return status.Errorf(codes.Unimplemented, "method ExclamationClientStream not implemented")
}
func (*UnimplementedExcitedAPIServer) ExclamationServerStream(req *ExclamationRequest, srv ExcitedAPI_ExclamationServerStreamServer) error {
	return status.Errorf(codes.Unimplemented, "method ExclamationServerStream not implemented")
}
func (*UnimplementedExcitedAPIServer) ExclamationBidiStream(srv ExcitedAPI_ExclamationBidiStreamServer) error {
	return status.Errorf(codes.Unimplemented, "method ExclamationBidiStream not implemented")
}

func RegisterExcitedAPIServer(s *grpc.Server, srv ExcitedAPIServer) {
	s.RegisterService(&_ExcitedAPI_serviceDesc, srv)
}

func _ExcitedAPI_Exclamation_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ExclamationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ExcitedAPIServer).Exclamation(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/uber.foo.v1.ExcitedAPI/Exclamation",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ExcitedAPIServer).Exclamation(ctx, req.(*ExclamationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ExcitedAPI_ExclamationClientStream_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(ExcitedAPIServer).ExclamationClientStream(&excitedAPIExclamationClientStreamServer{stream})
}

type ExcitedAPI_ExclamationClientStreamServer interface {
	SendAndClose(*ExclamationResponse) error
	Recv() (*ExclamationRequest, error)
	grpc.ServerStream
}

type excitedAPIExclamationClientStreamServer struct {
	grpc.ServerStream
}

func (x *excitedAPIExclamationClientStreamServer) SendAndClose(m *ExclamationResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *excitedAPIExclamationClientStreamServer) Recv() (*ExclamationRequest, error) {
	m := new(ExclamationRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _ExcitedAPI_ExclamationServerStream_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(ExclamationRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(ExcitedAPIServer).ExclamationServerStream(m, &excitedAPIExclamationServerStreamServer{stream})
}

type ExcitedAPI_ExclamationServerStreamServer interface {
	Send(*ExclamationResponse) error
	grpc.ServerStream
}

type excitedAPIExclamationServerStreamServer struct {
	grpc.ServerStream
}

func (x *excitedAPIExclamationServerStreamServer) Send(m *ExclamationResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _ExcitedAPI_ExclamationBidiStream_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(ExcitedAPIServer).ExclamationBidiStream(&excitedAPIExclamationBidiStreamServer{stream})
}

type ExcitedAPI_ExclamationBidiStreamServer interface {
	Send(*ExclamationResponse) error
	Recv() (*ExclamationRequest, error)
	grpc.ServerStream
}

type excitedAPIExclamationBidiStreamServer struct {
	grpc.ServerStream
}

func (x *excitedAPIExclamationBidiStreamServer) Send(m *ExclamationResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *excitedAPIExclamationBidiStreamServer) Recv() (*ExclamationRequest, error) {
	m := new(ExclamationRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _ExcitedAPI_serviceDesc = grpc.ServiceDesc{
	ServiceName: "uber.foo.v1.ExcitedAPI",
	HandlerType: (*ExcitedAPIServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Exclamation",
			Handler:    _ExcitedAPI_Exclamation_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "ExclamationClientStream",
			Handler:       _ExcitedAPI_ExclamationClientStream_Handler,
			ClientStreams: true,
		},
		{
			StreamName:    "ExclamationServerStream",
			Handler:       _ExcitedAPI_ExclamationServerStream_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "ExclamationBidiStream",
			Handler:       _ExcitedAPI_ExclamationBidiStream_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "uber/foo/v1/excited_api.proto",
}
