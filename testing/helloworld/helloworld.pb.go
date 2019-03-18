// Code generated by protoc-gen-go. DO NOT EDIT.
// source: helloworld.proto

package helloworld

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
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

// The request message containing the user's name.
type HelloRequest struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *HelloRequest) Reset()         { *m = HelloRequest{} }
func (m *HelloRequest) String() string { return proto.CompactTextString(m) }
func (*HelloRequest) ProtoMessage()    {}
func (*HelloRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_17b8c58d586b62f2, []int{0}
}

func (m *HelloRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HelloRequest.Unmarshal(m, b)
}
func (m *HelloRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HelloRequest.Marshal(b, m, deterministic)
}
func (m *HelloRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HelloRequest.Merge(m, src)
}
func (m *HelloRequest) XXX_Size() int {
	return xxx_messageInfo_HelloRequest.Size(m)
}
func (m *HelloRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_HelloRequest.DiscardUnknown(m)
}

var xxx_messageInfo_HelloRequest proto.InternalMessageInfo

func (m *HelloRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

// The response message containing the greetings
type HelloReply struct {
	Message              string   `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *HelloReply) Reset()         { *m = HelloReply{} }
func (m *HelloReply) String() string { return proto.CompactTextString(m) }
func (*HelloReply) ProtoMessage()    {}
func (*HelloReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_17b8c58d586b62f2, []int{1}
}

func (m *HelloReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HelloReply.Unmarshal(m, b)
}
func (m *HelloReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HelloReply.Marshal(b, m, deterministic)
}
func (m *HelloReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HelloReply.Merge(m, src)
}
func (m *HelloReply) XXX_Size() int {
	return xxx_messageInfo_HelloReply.Size(m)
}
func (m *HelloReply) XXX_DiscardUnknown() {
	xxx_messageInfo_HelloReply.DiscardUnknown(m)
}

var xxx_messageInfo_HelloReply proto.InternalMessageInfo

func (m *HelloReply) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

// The request message containing the user's name.
type NotifyReq struct {
	Nick                 string   `protobuf:"bytes,1,opt,name=nick,proto3" json:"nick,omitempty"`
	Addr                 string   `protobuf:"bytes,2,opt,name=addr,proto3" json:"addr,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *NotifyReq) Reset()         { *m = NotifyReq{} }
func (m *NotifyReq) String() string { return proto.CompactTextString(m) }
func (*NotifyReq) ProtoMessage()    {}
func (*NotifyReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_17b8c58d586b62f2, []int{2}
}

func (m *NotifyReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NotifyReq.Unmarshal(m, b)
}
func (m *NotifyReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NotifyReq.Marshal(b, m, deterministic)
}
func (m *NotifyReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NotifyReq.Merge(m, src)
}
func (m *NotifyReq) XXX_Size() int {
	return xxx_messageInfo_NotifyReq.Size(m)
}
func (m *NotifyReq) XXX_DiscardUnknown() {
	xxx_messageInfo_NotifyReq.DiscardUnknown(m)
}

var xxx_messageInfo_NotifyReq proto.InternalMessageInfo

func (m *NotifyReq) GetNick() string {
	if m != nil {
		return m.Nick
	}
	return ""
}

func (m *NotifyReq) GetAddr() string {
	if m != nil {
		return m.Addr
	}
	return ""
}

// The response message containing the greetings
type NotifyReply struct {
	Message              string   `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	Phone                string   `protobuf:"bytes,2,opt,name=phone,proto3" json:"phone,omitempty"`
	Car                  string   `protobuf:"bytes,3,opt,name=car,proto3" json:"car,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *NotifyReply) Reset()         { *m = NotifyReply{} }
func (m *NotifyReply) String() string { return proto.CompactTextString(m) }
func (*NotifyReply) ProtoMessage()    {}
func (*NotifyReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_17b8c58d586b62f2, []int{3}
}

func (m *NotifyReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NotifyReply.Unmarshal(m, b)
}
func (m *NotifyReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NotifyReply.Marshal(b, m, deterministic)
}
func (m *NotifyReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NotifyReply.Merge(m, src)
}
func (m *NotifyReply) XXX_Size() int {
	return xxx_messageInfo_NotifyReply.Size(m)
}
func (m *NotifyReply) XXX_DiscardUnknown() {
	xxx_messageInfo_NotifyReply.DiscardUnknown(m)
}

var xxx_messageInfo_NotifyReply proto.InternalMessageInfo

func (m *NotifyReply) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func (m *NotifyReply) GetPhone() string {
	if m != nil {
		return m.Phone
	}
	return ""
}

func (m *NotifyReply) GetCar() string {
	if m != nil {
		return m.Car
	}
	return ""
}

type SimpleData struct {
	Msg                  string   `protobuf:"bytes,1,opt,name=msg,proto3" json:"msg,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SimpleData) Reset()         { *m = SimpleData{} }
func (m *SimpleData) String() string { return proto.CompactTextString(m) }
func (*SimpleData) ProtoMessage()    {}
func (*SimpleData) Descriptor() ([]byte, []int) {
	return fileDescriptor_17b8c58d586b62f2, []int{4}
}

func (m *SimpleData) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SimpleData.Unmarshal(m, b)
}
func (m *SimpleData) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SimpleData.Marshal(b, m, deterministic)
}
func (m *SimpleData) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SimpleData.Merge(m, src)
}
func (m *SimpleData) XXX_Size() int {
	return xxx_messageInfo_SimpleData.Size(m)
}
func (m *SimpleData) XXX_DiscardUnknown() {
	xxx_messageInfo_SimpleData.DiscardUnknown(m)
}

var xxx_messageInfo_SimpleData proto.InternalMessageInfo

func (m *SimpleData) GetMsg() string {
	if m != nil {
		return m.Msg
	}
	return ""
}

type ServerStreamData struct {
	Msg                  string   `protobuf:"bytes,1,opt,name=msg,proto3" json:"msg,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ServerStreamData) Reset()         { *m = ServerStreamData{} }
func (m *ServerStreamData) String() string { return proto.CompactTextString(m) }
func (*ServerStreamData) ProtoMessage()    {}
func (*ServerStreamData) Descriptor() ([]byte, []int) {
	return fileDescriptor_17b8c58d586b62f2, []int{5}
}

func (m *ServerStreamData) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ServerStreamData.Unmarshal(m, b)
}
func (m *ServerStreamData) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ServerStreamData.Marshal(b, m, deterministic)
}
func (m *ServerStreamData) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ServerStreamData.Merge(m, src)
}
func (m *ServerStreamData) XXX_Size() int {
	return xxx_messageInfo_ServerStreamData.Size(m)
}
func (m *ServerStreamData) XXX_DiscardUnknown() {
	xxx_messageInfo_ServerStreamData.DiscardUnknown(m)
}

var xxx_messageInfo_ServerStreamData proto.InternalMessageInfo

func (m *ServerStreamData) GetMsg() string {
	if m != nil {
		return m.Msg
	}
	return ""
}

func init() {
	proto.RegisterType((*HelloRequest)(nil), "helloworld.HelloRequest")
	proto.RegisterType((*HelloReply)(nil), "helloworld.HelloReply")
	proto.RegisterType((*NotifyReq)(nil), "helloworld.NotifyReq")
	proto.RegisterType((*NotifyReply)(nil), "helloworld.NotifyReply")
	proto.RegisterType((*SimpleData)(nil), "helloworld.SimpleData")
	proto.RegisterType((*ServerStreamData)(nil), "helloworld.ServerStreamData")
}

func init() { proto.RegisterFile("helloworld.proto", fileDescriptor_17b8c58d586b62f2) }

var fileDescriptor_17b8c58d586b62f2 = []byte{
	// 365 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x53, 0x5d, 0x4b, 0xeb, 0x40,
	0x10, 0x6d, 0x6e, 0x6f, 0x9b, 0xdb, 0xb9, 0xf7, 0xa1, 0x77, 0xfc, 0x0a, 0x41, 0x44, 0x16, 0x91,
	0x3e, 0x95, 0x92, 0x3e, 0x89, 0x5f, 0xa0, 0x42, 0x0b, 0x82, 0x4a, 0x22, 0xf8, 0xbc, 0x4d, 0xc6,
	0x76, 0x69, 0xd2, 0xa4, 0x9b, 0xa8, 0xe4, 0x1f, 0xf8, 0xab, 0xfc, 0x6d, 0xb2, 0x49, 0xd3, 0xa6,
	0x52, 0x0b, 0xf6, 0x6d, 0xe6, 0xcc, 0x39, 0x87, 0xd9, 0x33, 0x2c, 0x34, 0x47, 0xe4, 0xfb, 0xe1,
	0x5b, 0x28, 0x7d, 0xaf, 0x1d, 0xc9, 0x30, 0x09, 0x11, 0x16, 0x08, 0x63, 0xf0, 0xaf, 0xaf, 0x3a,
	0x9b, 0xa6, 0x2f, 0x14, 0x27, 0x88, 0xf0, 0x7b, 0xc2, 0x03, 0x32, 0xb4, 0x43, 0xad, 0xd5, 0xb0,
	0xb3, 0x9a, 0x1d, 0x03, 0xcc, 0x38, 0x91, 0x9f, 0xa2, 0x01, 0x7a, 0x40, 0x71, 0xcc, 0x87, 0x05,
	0xa9, 0x68, 0x59, 0x17, 0x1a, 0x77, 0x61, 0x22, 0x9e, 0x53, 0x9b, 0xa6, 0x99, 0x91, 0x70, 0xc7,
	0x73, 0x23, 0xe1, 0x8e, 0x15, 0xc6, 0x3d, 0x4f, 0x1a, 0xbf, 0x72, 0x4c, 0xd5, 0xec, 0x1e, 0xfe,
	0x16, 0xa2, 0xb5, 0xee, 0xb8, 0x0d, 0xb5, 0x68, 0x14, 0x4e, 0x68, 0xa6, 0xce, 0x1b, 0x6c, 0x42,
	0xd5, 0xe5, 0xd2, 0xa8, 0x66, 0x98, 0x2a, 0xd9, 0x01, 0x80, 0x23, 0x82, 0xc8, 0xa7, 0x1b, 0x9e,
	0x70, 0x35, 0x0f, 0xe2, 0xe1, 0xcc, 0x4b, 0x95, 0xec, 0x08, 0x9a, 0x0e, 0xc9, 0x57, 0x92, 0x4e,
	0x22, 0x89, 0x07, 0xab, 0x59, 0xd6, 0xbb, 0x06, 0x7a, 0x4f, 0x12, 0x25, 0x24, 0xf1, 0x02, 0xfe,
	0x38, 0x3c, 0xcd, 0x22, 0x40, 0xa3, 0x5d, 0x8a, 0xb3, 0x9c, 0x9c, 0xb9, 0xbb, 0x62, 0x12, 0xf9,
	0x29, 0xab, 0xe0, 0x39, 0xe8, 0x0e, 0x4f, 0x7b, 0x42, 0xfa, 0x9b, 0xc8, 0xad, 0x0f, 0x0d, 0xea,
	0x79, 0x44, 0x78, 0x96, 0x6d, 0xf2, 0xa4, 0x38, 0xb8, 0x53, 0x16, 0xcc, 0x73, 0x37, 0xf7, 0x56,
	0xc1, 0xf9, 0x1e, 0xa7, 0xd9, 0x1e, 0x7d, 0x3e, 0xe2, 0x1b, 0x88, 0x4f, 0xa0, 0xa6, 0xc4, 0xe2,
	0xe7, 0x52, 0xeb, 0x11, 0xfe, 0x5f, 0x09, 0x4f, 0xe4, 0x79, 0xab, 0xec, 0x85, 0x4b, 0x78, 0x09,
	0xba, 0x02, 0xed, 0x87, 0x6b, 0x5c, 0x7a, 0xfa, 0xe2, 0x76, 0xe6, 0x37, 0x38, 0xab, 0xb4, 0xb4,
	0x8e, 0x66, 0x0d, 0x60, 0xab, 0x7c, 0xc7, 0xc2, 0xf7, 0x16, 0x1a, 0x39, 0x60, 0x47, 0x2e, 0xee,
	0x2f, 0x39, 0x7c, 0xb9, 0xba, 0xb9, 0x76, 0xca, 0x2a, 0x1d, 0x6d, 0x50, 0xcf, 0x3e, 0x4c, 0xf7,
	0x33, 0x00, 0x00, 0xff, 0xff, 0x3d, 0xb3, 0xda, 0xc6, 0x44, 0x03, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// GreeterClient is the client API for Greeter service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type GreeterClient interface {
	// Sends a greeting
	SayHello(ctx context.Context, in *HelloRequest, opts ...grpc.CallOption) (*HelloReply, error)
	SayGirl(ctx context.Context, in *HelloRequest, opts ...grpc.CallOption) (*HelloReply, error)
}

type greeterClient struct {
	cc *grpc.ClientConn
}

func NewGreeterClient(cc *grpc.ClientConn) GreeterClient {
	return &greeterClient{cc}
}

func (c *greeterClient) SayHello(ctx context.Context, in *HelloRequest, opts ...grpc.CallOption) (*HelloReply, error) {
	out := new(HelloReply)
	err := c.cc.Invoke(ctx, "/helloworld.Greeter/SayHello", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *greeterClient) SayGirl(ctx context.Context, in *HelloRequest, opts ...grpc.CallOption) (*HelloReply, error) {
	out := new(HelloReply)
	err := c.cc.Invoke(ctx, "/helloworld.Greeter/SayGirl", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GreeterServer is the server API for Greeter service.
type GreeterServer interface {
	// Sends a greeting
	SayHello(context.Context, *HelloRequest) (*HelloReply, error)
	SayGirl(context.Context, *HelloRequest) (*HelloReply, error)
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
		FullMethod: "/helloworld.Greeter/SayHello",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GreeterServer).SayHello(ctx, req.(*HelloRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Greeter_SayGirl_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HelloRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GreeterServer).SayGirl(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/helloworld.Greeter/SayGirl",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GreeterServer).SayGirl(ctx, req.(*HelloRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Greeter_serviceDesc = grpc.ServiceDesc{
	ServiceName: "helloworld.Greeter",
	HandlerType: (*GreeterServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SayHello",
			Handler:    _Greeter_SayHello_Handler,
		},
		{
			MethodName: "SayGirl",
			Handler:    _Greeter_SayGirl_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "helloworld.proto",
}

// NotifyClient is the client API for Notify service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type NotifyClient interface {
	// Sends notify
	SayWorld(ctx context.Context, in *NotifyReq, opts ...grpc.CallOption) (*NotifyReply, error)
	SayHaha(ctx context.Context, in *NotifyReq, opts ...grpc.CallOption) (*NotifyReply, error)
	SayHi(ctx context.Context, in *NotifyReq, opts ...grpc.CallOption) (*NotifyReply, error)
}

type notifyClient struct {
	cc *grpc.ClientConn
}

func NewNotifyClient(cc *grpc.ClientConn) NotifyClient {
	return &notifyClient{cc}
}

func (c *notifyClient) SayWorld(ctx context.Context, in *NotifyReq, opts ...grpc.CallOption) (*NotifyReply, error) {
	out := new(NotifyReply)
	err := c.cc.Invoke(ctx, "/helloworld.Notify/SayWorld", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *notifyClient) SayHaha(ctx context.Context, in *NotifyReq, opts ...grpc.CallOption) (*NotifyReply, error) {
	out := new(NotifyReply)
	err := c.cc.Invoke(ctx, "/helloworld.Notify/SayHaha", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *notifyClient) SayHi(ctx context.Context, in *NotifyReq, opts ...grpc.CallOption) (*NotifyReply, error) {
	out := new(NotifyReply)
	err := c.cc.Invoke(ctx, "/helloworld.Notify/SayHi", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// NotifyServer is the server API for Notify service.
type NotifyServer interface {
	// Sends notify
	SayWorld(context.Context, *NotifyReq) (*NotifyReply, error)
	SayHaha(context.Context, *NotifyReq) (*NotifyReply, error)
	SayHi(context.Context, *NotifyReq) (*NotifyReply, error)
}

func RegisterNotifyServer(s *grpc.Server, srv NotifyServer) {
	s.RegisterService(&_Notify_serviceDesc, srv)
}

func _Notify_SayWorld_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NotifyReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NotifyServer).SayWorld(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/helloworld.Notify/SayWorld",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NotifyServer).SayWorld(ctx, req.(*NotifyReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Notify_SayHaha_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NotifyReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NotifyServer).SayHaha(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/helloworld.Notify/SayHaha",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NotifyServer).SayHaha(ctx, req.(*NotifyReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Notify_SayHi_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NotifyReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NotifyServer).SayHi(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/helloworld.Notify/SayHi",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NotifyServer).SayHi(ctx, req.(*NotifyReq))
	}
	return interceptor(ctx, in, info, handler)
}

var _Notify_serviceDesc = grpc.ServiceDesc{
	ServiceName: "helloworld.Notify",
	HandlerType: (*NotifyServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SayWorld",
			Handler:    _Notify_SayWorld_Handler,
		},
		{
			MethodName: "SayHaha",
			Handler:    _Notify_SayHaha_Handler,
		},
		{
			MethodName: "SayHi",
			Handler:    _Notify_SayHi_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "helloworld.proto",
}

// BidiStreamServiceClient is the client API for BidiStreamService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type BidiStreamServiceClient interface {
	BidiRPC(ctx context.Context, opts ...grpc.CallOption) (BidiStreamService_BidiRPCClient, error)
}

type bidiStreamServiceClient struct {
	cc *grpc.ClientConn
}

func NewBidiStreamServiceClient(cc *grpc.ClientConn) BidiStreamServiceClient {
	return &bidiStreamServiceClient{cc}
}

func (c *bidiStreamServiceClient) BidiRPC(ctx context.Context, opts ...grpc.CallOption) (BidiStreamService_BidiRPCClient, error) {
	stream, err := c.cc.NewStream(ctx, &_BidiStreamService_serviceDesc.Streams[0], "/helloworld.BidiStreamService/BidiRPC", opts...)
	if err != nil {
		return nil, err
	}
	x := &bidiStreamServiceBidiRPCClient{stream}
	return x, nil
}

type BidiStreamService_BidiRPCClient interface {
	Send(*SimpleData) error
	Recv() (*SimpleData, error)
	grpc.ClientStream
}

type bidiStreamServiceBidiRPCClient struct {
	grpc.ClientStream
}

func (x *bidiStreamServiceBidiRPCClient) Send(m *SimpleData) error {
	return x.ClientStream.SendMsg(m)
}

func (x *bidiStreamServiceBidiRPCClient) Recv() (*SimpleData, error) {
	m := new(SimpleData)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// BidiStreamServiceServer is the server API for BidiStreamService service.
type BidiStreamServiceServer interface {
	BidiRPC(BidiStreamService_BidiRPCServer) error
}

func RegisterBidiStreamServiceServer(s *grpc.Server, srv BidiStreamServiceServer) {
	s.RegisterService(&_BidiStreamService_serviceDesc, srv)
}

func _BidiStreamService_BidiRPC_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(BidiStreamServiceServer).BidiRPC(&bidiStreamServiceBidiRPCServer{stream})
}

type BidiStreamService_BidiRPCServer interface {
	Send(*SimpleData) error
	Recv() (*SimpleData, error)
	grpc.ServerStream
}

type bidiStreamServiceBidiRPCServer struct {
	grpc.ServerStream
}

func (x *bidiStreamServiceBidiRPCServer) Send(m *SimpleData) error {
	return x.ServerStream.SendMsg(m)
}

func (x *bidiStreamServiceBidiRPCServer) Recv() (*SimpleData, error) {
	m := new(SimpleData)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _BidiStreamService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "helloworld.BidiStreamService",
	HandlerType: (*BidiStreamServiceServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "BidiRPC",
			Handler:       _BidiStreamService_BidiRPC_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "helloworld.proto",
}

// ServerStreamServiceClient is the client API for ServerStreamService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ServerStreamServiceClient interface {
	StreamRpc(ctx context.Context, in *ServerStreamData, opts ...grpc.CallOption) (ServerStreamService_StreamRpcClient, error)
}

type serverStreamServiceClient struct {
	cc *grpc.ClientConn
}

func NewServerStreamServiceClient(cc *grpc.ClientConn) ServerStreamServiceClient {
	return &serverStreamServiceClient{cc}
}

func (c *serverStreamServiceClient) StreamRpc(ctx context.Context, in *ServerStreamData, opts ...grpc.CallOption) (ServerStreamService_StreamRpcClient, error) {
	stream, err := c.cc.NewStream(ctx, &_ServerStreamService_serviceDesc.Streams[0], "/helloworld.ServerStreamService/StreamRpc", opts...)
	if err != nil {
		return nil, err
	}
	x := &serverStreamServiceStreamRpcClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type ServerStreamService_StreamRpcClient interface {
	Recv() (*ServerStreamData, error)
	grpc.ClientStream
}

type serverStreamServiceStreamRpcClient struct {
	grpc.ClientStream
}

func (x *serverStreamServiceStreamRpcClient) Recv() (*ServerStreamData, error) {
	m := new(ServerStreamData)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// ServerStreamServiceServer is the server API for ServerStreamService service.
type ServerStreamServiceServer interface {
	StreamRpc(*ServerStreamData, ServerStreamService_StreamRpcServer) error
}

func RegisterServerStreamServiceServer(s *grpc.Server, srv ServerStreamServiceServer) {
	s.RegisterService(&_ServerStreamService_serviceDesc, srv)
}

func _ServerStreamService_StreamRpc_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(ServerStreamData)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(ServerStreamServiceServer).StreamRpc(m, &serverStreamServiceStreamRpcServer{stream})
}

type ServerStreamService_StreamRpcServer interface {
	Send(*ServerStreamData) error
	grpc.ServerStream
}

type serverStreamServiceStreamRpcServer struct {
	grpc.ServerStream
}

func (x *serverStreamServiceStreamRpcServer) Send(m *ServerStreamData) error {
	return x.ServerStream.SendMsg(m)
}

var _ServerStreamService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "helloworld.ServerStreamService",
	HandlerType: (*ServerStreamServiceServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "StreamRpc",
			Handler:       _ServerStreamService_StreamRpc_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "helloworld.proto",
}
