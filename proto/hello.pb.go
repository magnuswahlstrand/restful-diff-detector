// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/hello.proto

/*
Package hello is a generated protocol buffer package.

It is generated from these files:
	proto/hello.proto

It has these top-level messages:
	HelloRequest
	HelloReply
	DiffNotification
	DiffSubscribe
*/
package hello

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import google_protobuf "github.com/golang/protobuf/ptypes/struct"

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

// The request message containing the user's name.
type HelloRequest struct {
	Name string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
}

func (m *HelloRequest) Reset()                    { *m = HelloRequest{} }
func (m *HelloRequest) String() string            { return proto.CompactTextString(m) }
func (*HelloRequest) ProtoMessage()               {}
func (*HelloRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *HelloRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

// The response message containing the greetings
type HelloReply struct {
	Message string `protobuf:"bytes,1,opt,name=message" json:"message,omitempty"`
}

func (m *HelloReply) Reset()                    { *m = HelloReply{} }
func (m *HelloReply) String() string            { return proto.CompactTextString(m) }
func (*HelloReply) ProtoMessage()               {}
func (*HelloReply) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *HelloReply) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

// The notification containing
type DiffNotification struct {
	ResponseData *google_protobuf.Struct `protobuf:"bytes,1,opt,name=responseData" json:"responseData,omitempty"`
}

func (m *DiffNotification) Reset()                    { *m = DiffNotification{} }
func (m *DiffNotification) String() string            { return proto.CompactTextString(m) }
func (*DiffNotification) ProtoMessage()               {}
func (*DiffNotification) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *DiffNotification) GetResponseData() *google_protobuf.Struct {
	if m != nil {
		return m.ResponseData
	}
	return nil
}

// The subscribe message containing the rest path, sub name.
type DiffSubscribe struct {
	Path           string `protobuf:"bytes,1,opt,name=path" json:"path,omitempty"`
	Period         int32  `protobuf:"varint,2,opt,name=period" json:"period,omitempty"`
	SubscriberName string `protobuf:"bytes,3,opt,name=subscriberName" json:"subscriberName,omitempty"`
}

func (m *DiffSubscribe) Reset()                    { *m = DiffSubscribe{} }
func (m *DiffSubscribe) String() string            { return proto.CompactTextString(m) }
func (*DiffSubscribe) ProtoMessage()               {}
func (*DiffSubscribe) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *DiffSubscribe) GetPath() string {
	if m != nil {
		return m.Path
	}
	return ""
}

func (m *DiffSubscribe) GetPeriod() int32 {
	if m != nil {
		return m.Period
	}
	return 0
}

func (m *DiffSubscribe) GetSubscriberName() string {
	if m != nil {
		return m.SubscriberName
	}
	return ""
}

func init() {
	proto.RegisterType((*HelloRequest)(nil), "hello.HelloRequest")
	proto.RegisterType((*HelloReply)(nil), "hello.HelloReply")
	proto.RegisterType((*DiffNotification)(nil), "hello.DiffNotification")
	proto.RegisterType((*DiffSubscribe)(nil), "hello.DiffSubscribe")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Greeter service

type GreeterClient interface {
	// Sends a greeting
	SayHello(ctx context.Context, in *HelloRequest, opts ...grpc.CallOption) (*HelloReply, error)
}

type greeterClient struct {
	cc *grpc.ClientConn
}

func NewGreeterClient(cc *grpc.ClientConn) GreeterClient {
	return &greeterClient{cc}
}

func (c *greeterClient) SayHello(ctx context.Context, in *HelloRequest, opts ...grpc.CallOption) (*HelloReply, error) {
	out := new(HelloReply)
	err := grpc.Invoke(ctx, "/hello.Greeter/SayHello", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Greeter service

type GreeterServer interface {
	// Sends a greeting
	SayHello(context.Context, *HelloRequest) (*HelloReply, error)
}

func RegisterGreeterServer(s *grpc.Server, srv GreeterServer) {
	s.RegisterService(&_Greeter_serviceDesc, srv)
}

func _Greeter_SayHello_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HelloRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GreeterServer).SayHello(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/hello.Greeter/SayHello",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GreeterServer).SayHello(ctx, req.(*HelloRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Greeter_serviceDesc = grpc.ServiceDesc{
	ServiceName: "hello.Greeter",
	HandlerType: (*GreeterServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SayHello",
			Handler:    _Greeter_SayHello_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/hello.proto",
}

// Client API for DiffSubscriber service

type DiffSubscriberClient interface {
	// Sends a greeting
	Subscribe(ctx context.Context, in *DiffSubscribe, opts ...grpc.CallOption) (*DiffNotification, error)
}

type diffSubscriberClient struct {
	cc *grpc.ClientConn
}

func NewDiffSubscriberClient(cc *grpc.ClientConn) DiffSubscriberClient {
	return &diffSubscriberClient{cc}
}

func (c *diffSubscriberClient) Subscribe(ctx context.Context, in *DiffSubscribe, opts ...grpc.CallOption) (*DiffNotification, error) {
	out := new(DiffNotification)
	err := grpc.Invoke(ctx, "/hello.DiffSubscriber/Subscribe", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for DiffSubscriber service

type DiffSubscriberServer interface {
	// Sends a greeting
	Subscribe(context.Context, *DiffSubscribe) (*DiffNotification, error)
}

func RegisterDiffSubscriberServer(s *grpc.Server, srv DiffSubscriberServer) {
	s.RegisterService(&_DiffSubscriber_serviceDesc, srv)
}

func _DiffSubscriber_Subscribe_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DiffSubscribe)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DiffSubscriberServer).Subscribe(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/hello.DiffSubscriber/Subscribe",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DiffSubscriberServer).Subscribe(ctx, req.(*DiffSubscribe))
	}
	return interceptor(ctx, in, info, handler)
}

var _DiffSubscriber_serviceDesc = grpc.ServiceDesc{
	ServiceName: "hello.DiffSubscriber",
	HandlerType: (*DiffSubscriberServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Subscribe",
			Handler:    _DiffSubscriber_Subscribe_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/hello.proto",
}

func init() { proto.RegisterFile("proto/hello.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 321 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x5c, 0x90, 0xc1, 0x4f, 0xfa, 0x30,
	0x1c, 0xc5, 0xd9, 0xef, 0x27, 0x20, 0x5f, 0x11, 0xa5, 0x1a, 0x21, 0xe8, 0x81, 0xec, 0x40, 0x38,
	0x15, 0x83, 0xde, 0x34, 0x31, 0x21, 0x24, 0x7a, 0x42, 0x32, 0x0e, 0x9e, 0xbb, 0xf1, 0xdd, 0x68,
	0x52, 0x68, 0x6d, 0xbb, 0x28, 0xff, 0xbd, 0x69, 0xc7, 0x22, 0xf3, 0xf6, 0xbe, 0x2f, 0xef, 0x6d,
	0x9f, 0x3e, 0xe8, 0x2a, 0x2d, 0xad, 0x9c, 0x6c, 0x50, 0x08, 0x49, 0xbd, 0x26, 0x75, 0x7f, 0x0c,
	0xee, 0x32, 0x29, 0x33, 0x81, 0x13, 0x6f, 0xc6, 0x79, 0x3a, 0x31, 0x56, 0xe7, 0x89, 0x2d, 0x42,
	0x61, 0x08, 0xed, 0x37, 0x17, 0x8b, 0xf0, 0x33, 0x47, 0x63, 0x09, 0x81, 0x93, 0x1d, 0xdb, 0x62,
	0x3f, 0x18, 0x06, 0xe3, 0x56, 0xe4, 0x75, 0x38, 0x02, 0x38, 0x64, 0x94, 0xd8, 0x93, 0x3e, 0x34,
	0xb7, 0x68, 0x0c, 0xcb, 0xca, 0x50, 0x79, 0x86, 0xef, 0x70, 0x39, 0xe7, 0x69, 0xba, 0x90, 0x96,
	0xa7, 0x3c, 0x61, 0x96, 0xcb, 0x1d, 0x79, 0x82, 0xb6, 0x46, 0xa3, 0xe4, 0xce, 0xe0, 0x9c, 0x59,
	0xe6, 0x2b, 0x67, 0xd3, 0x1e, 0x2d, 0xa0, 0x68, 0x09, 0x45, 0x57, 0x1e, 0x2a, 0xaa, 0x84, 0xc3,
	0x04, 0xce, 0xdd, 0x07, 0x57, 0x79, 0x6c, 0x12, 0xcd, 0x63, 0x74, 0x74, 0x8a, 0xd9, 0x4d, 0x49,
	0xe7, 0x34, 0xb9, 0x81, 0x86, 0x42, 0xcd, 0xe5, 0xba, 0xff, 0x6f, 0x18, 0x8c, 0xeb, 0xd1, 0xe1,
	0x22, 0x23, 0xe8, 0x98, 0xb2, 0xa8, 0x17, 0xee, 0x4d, 0xff, 0x7d, 0xeb, 0x8f, 0x3b, 0x7d, 0x81,
	0xe6, 0xab, 0x46, 0xb4, 0xa8, 0xc9, 0x23, 0x9c, 0xae, 0xd8, 0xde, 0xbf, 0x95, 0x5c, 0xd1, 0x62,
	0xcb, 0xe3, 0x75, 0x06, 0xdd, 0xaa, 0xa9, 0xc4, 0x3e, 0xac, 0x4d, 0x17, 0xd0, 0xa9, 0x50, 0x6a,
	0xf2, 0x0c, 0xad, 0x5f, 0xe6, 0xeb, 0x43, 0xa7, 0x92, 0x19, 0xf4, 0x8e, 0xdc, 0xe3, 0xc1, 0xc2,
	0xda, 0xec, 0x1e, 0x6e, 0xb9, 0xa4, 0x99, 0x56, 0x09, 0xc5, 0x6f, 0xb6, 0x55, 0x02, 0x4d, 0x11,
	0xfe, 0x92, 0x5a, 0xac, 0x67, 0x17, 0xfe, 0xe7, 0x1f, 0x4e, 0x2f, 0xdd, 0x7a, 0xcb, 0x20, 0x6e,
	0xf8, 0x19, 0x1f, 0x7e, 0x02, 0x00, 0x00, 0xff, 0xff, 0x10, 0x55, 0x75, 0x85, 0x05, 0x02, 0x00,
	0x00,
}
