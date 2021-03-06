// Code generated by protoc-gen-go. DO NOT EDIT.
// source: math.proto

package math

import (
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
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

// message for Math parameters, only support two parameters  both of them need to be integers
type Parameters struct {
	Param1               int32    `protobuf:"varint,1,opt,name=param1,proto3" json:"param1,omitempty"`
	Param2               int32    `protobuf:"varint,2,opt,name=param2,proto3" json:"param2,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Parameters) Reset()         { *m = Parameters{} }
func (m *Parameters) String() string { return proto.CompactTextString(m) }
func (*Parameters) ProtoMessage()    {}
func (*Parameters) Descriptor() ([]byte, []int) {
	return fileDescriptor_f139a3799a86a974, []int{0}
}

func (m *Parameters) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Parameters.Unmarshal(m, b)
}
func (m *Parameters) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Parameters.Marshal(b, m, deterministic)
}
func (m *Parameters) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Parameters.Merge(m, src)
}
func (m *Parameters) XXX_Size() int {
	return xxx_messageInfo_Parameters.Size(m)
}
func (m *Parameters) XXX_DiscardUnknown() {
	xxx_messageInfo_Parameters.DiscardUnknown(m)
}

var xxx_messageInfo_Parameters proto.InternalMessageInfo

func (m *Parameters) GetParam1() int32 {
	if m != nil {
		return m.Param1
	}
	return 0
}

func (m *Parameters) GetParam2() int32 {
	if m != nil {
		return m.Param2
	}
	return 0
}

// message for Math response, will be float
type Result struct {
	Result               float32  `protobuf:"fixed32,1,opt,name=result,proto3" json:"result,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Result) Reset()         { *m = Result{} }
func (m *Result) String() string { return proto.CompactTextString(m) }
func (*Result) ProtoMessage()    {}
func (*Result) Descriptor() ([]byte, []int) {
	return fileDescriptor_f139a3799a86a974, []int{1}
}

func (m *Result) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Result.Unmarshal(m, b)
}
func (m *Result) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Result.Marshal(b, m, deterministic)
}
func (m *Result) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Result.Merge(m, src)
}
func (m *Result) XXX_Size() int {
	return xxx_messageInfo_Result.Size(m)
}
func (m *Result) XXX_DiscardUnknown() {
	xxx_messageInfo_Result.DiscardUnknown(m)
}

var xxx_messageInfo_Result proto.InternalMessageInfo

func (m *Result) GetResult() float32 {
	if m != nil {
		return m.Result
	}
	return 0
}

func init() {
	proto.RegisterType((*Parameters)(nil), "math.Parameters")
	proto.RegisterType((*Result)(nil), "math.Result")
}

func init() { proto.RegisterFile("math.proto", fileDescriptor_f139a3799a86a974) }

var fileDescriptor_f139a3799a86a974 = []byte{
	// 168 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xe2, 0xe2, 0xca, 0x4d, 0x2c, 0xc9,
	0xd0, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x01, 0xb1, 0x95, 0x6c, 0xb8, 0xb8, 0x02, 0x12,
	0x8b, 0x12, 0x73, 0x53, 0x4b, 0x52, 0x8b, 0x8a, 0x85, 0xc4, 0xb8, 0xd8, 0x0a, 0x40, 0x3c, 0x43,
	0x09, 0x46, 0x05, 0x46, 0x0d, 0xd6, 0x20, 0x28, 0x0f, 0x2e, 0x6e, 0x24, 0xc1, 0x84, 0x24, 0x6e,
	0xa4, 0xa4, 0xc0, 0xc5, 0x16, 0x94, 0x5a, 0x5c, 0x9a, 0x53, 0x02, 0x52, 0x51, 0x04, 0x66, 0x81,
	0x75, 0x32, 0x05, 0x41, 0x79, 0x46, 0xeb, 0x19, 0xb9, 0x58, 0x7c, 0x13, 0x4b, 0x32, 0x84, 0x54,
	0xb9, 0x98, 0x1d, 0x53, 0x52, 0x84, 0x04, 0xf4, 0xc0, 0x4e, 0x40, 0xd8, 0x29, 0xc5, 0x03, 0x11,
	0x81, 0x9a, 0xa3, 0xc5, 0xc5, 0x11, 0x5c, 0x9a, 0x54, 0x52, 0x94, 0x98, 0x5c, 0x42, 0x8c, 0x5a,
	0xdf, 0xd2, 0x9c, 0x92, 0xcc, 0x82, 0x9c, 0x4a, 0x82, 0x6a, 0x35, 0xb8, 0xd8, 0x5c, 0x32, 0xcb,
	0x32, 0x53, 0x52, 0x09, 0xa9, 0x4c, 0x62, 0x03, 0x07, 0x8f, 0x31, 0x20, 0x00, 0x00, 0xff, 0xff,
	0x62, 0x56, 0xf4, 0x80, 0x2c, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// MathClient is the client API for Math service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type MathClient interface {
	Add(ctx context.Context, in *Parameters, opts ...grpc.CallOption) (*Result, error)
	Subtract(ctx context.Context, in *Parameters, opts ...grpc.CallOption) (*Result, error)
	Multiply(ctx context.Context, in *Parameters, opts ...grpc.CallOption) (*Result, error)
	Divide(ctx context.Context, in *Parameters, opts ...grpc.CallOption) (*Result, error)
}

type mathClient struct {
	cc *grpc.ClientConn
}

func NewMathClient(cc *grpc.ClientConn) MathClient {
	return &mathClient{cc}
}

func (c *mathClient) Add(ctx context.Context, in *Parameters, opts ...grpc.CallOption) (*Result, error) {
	out := new(Result)
	err := c.cc.Invoke(ctx, "/math.Math/Add", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mathClient) Subtract(ctx context.Context, in *Parameters, opts ...grpc.CallOption) (*Result, error) {
	out := new(Result)
	err := c.cc.Invoke(ctx, "/math.Math/Subtract", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mathClient) Multiply(ctx context.Context, in *Parameters, opts ...grpc.CallOption) (*Result, error) {
	out := new(Result)
	err := c.cc.Invoke(ctx, "/math.Math/Multiply", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mathClient) Divide(ctx context.Context, in *Parameters, opts ...grpc.CallOption) (*Result, error) {
	out := new(Result)
	err := c.cc.Invoke(ctx, "/math.Math/Divide", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MathServer is the server API for Math service.
type MathServer interface {
	Add(context.Context, *Parameters) (*Result, error)
	Subtract(context.Context, *Parameters) (*Result, error)
	Multiply(context.Context, *Parameters) (*Result, error)
	Divide(context.Context, *Parameters) (*Result, error)
}

func RegisterMathServer(s *grpc.Server, srv MathServer) {
	s.RegisterService(&_Math_serviceDesc, srv)
}

func _Math_Add_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Parameters)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MathServer).Add(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/math.Math/Add",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MathServer).Add(ctx, req.(*Parameters))
	}
	return interceptor(ctx, in, info, handler)
}

func _Math_Subtract_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Parameters)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MathServer).Subtract(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/math.Math/Subtract",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MathServer).Subtract(ctx, req.(*Parameters))
	}
	return interceptor(ctx, in, info, handler)
}

func _Math_Multiply_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Parameters)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MathServer).Multiply(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/math.Math/Multiply",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MathServer).Multiply(ctx, req.(*Parameters))
	}
	return interceptor(ctx, in, info, handler)
}

func _Math_Divide_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Parameters)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MathServer).Divide(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/math.Math/Divide",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MathServer).Divide(ctx, req.(*Parameters))
	}
	return interceptor(ctx, in, info, handler)
}

var _Math_serviceDesc = grpc.ServiceDesc{
	ServiceName: "math.Math",
	HandlerType: (*MathServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Add",
			Handler:    _Math_Add_Handler,
		},
		{
			MethodName: "Subtract",
			Handler:    _Math_Subtract_Handler,
		},
		{
			MethodName: "Multiply",
			Handler:    _Math_Multiply_Handler,
		},
		{
			MethodName: "Divide",
			Handler:    _Math_Divide_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "math.proto",
}
