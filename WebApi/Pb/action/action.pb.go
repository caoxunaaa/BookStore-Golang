// Code generated by protoc-gen-go. DO NOT EDIT.
// source: action.proto

package action

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

type Request struct {
	Ping                 string   `protobuf:"bytes,1,opt,name=ping,proto3" json:"ping,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Request) Reset()         { *m = Request{} }
func (m *Request) String() string { return proto.CompactTextString(m) }
func (*Request) ProtoMessage()    {}
func (*Request) Descriptor() ([]byte, []int) {
	return fileDescriptor_59885c909ad4dfd3, []int{0}
}

func (m *Request) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Request.Unmarshal(m, b)
}
func (m *Request) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Request.Marshal(b, m, deterministic)
}
func (m *Request) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Request.Merge(m, src)
}
func (m *Request) XXX_Size() int {
	return xxx_messageInfo_Request.Size(m)
}
func (m *Request) XXX_DiscardUnknown() {
	xxx_messageInfo_Request.DiscardUnknown(m)
}

var xxx_messageInfo_Request proto.InternalMessageInfo

func (m *Request) GetPing() string {
	if m != nil {
		return m.Ping
	}
	return ""
}

type Response struct {
	Ok                   bool     `protobuf:"varint,1,opt,name=ok,proto3" json:"ok,omitempty"`
	Message              string   `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Response) Reset()         { *m = Response{} }
func (m *Response) String() string { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()    {}
func (*Response) Descriptor() ([]byte, []int) {
	return fileDescriptor_59885c909ad4dfd3, []int{1}
}

func (m *Response) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Response.Unmarshal(m, b)
}
func (m *Response) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Response.Marshal(b, m, deterministic)
}
func (m *Response) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Response.Merge(m, src)
}
func (m *Response) XXX_Size() int {
	return xxx_messageInfo_Response.Size(m)
}
func (m *Response) XXX_DiscardUnknown() {
	xxx_messageInfo_Response.DiscardUnknown(m)
}

var xxx_messageInfo_Response proto.InternalMessageInfo

func (m *Response) GetOk() bool {
	if m != nil {
		return m.Ok
	}
	return false
}

func (m *Response) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

type TrafficStatisticReq struct {
	Id                   int64    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	BookId               int64    `protobuf:"varint,2,opt,name=BookId,proto3" json:"BookId,omitempty"`
	ChapterNum           int64    `protobuf:"varint,3,opt,name=ChapterNum,proto3" json:"ChapterNum,omitempty"`
	TrafficNumber        int64    `protobuf:"varint,4,opt,name=TrafficNumber,proto3" json:"TrafficNumber,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TrafficStatisticReq) Reset()         { *m = TrafficStatisticReq{} }
func (m *TrafficStatisticReq) String() string { return proto.CompactTextString(m) }
func (*TrafficStatisticReq) ProtoMessage()    {}
func (*TrafficStatisticReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_59885c909ad4dfd3, []int{2}
}

func (m *TrafficStatisticReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TrafficStatisticReq.Unmarshal(m, b)
}
func (m *TrafficStatisticReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TrafficStatisticReq.Marshal(b, m, deterministic)
}
func (m *TrafficStatisticReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TrafficStatisticReq.Merge(m, src)
}
func (m *TrafficStatisticReq) XXX_Size() int {
	return xxx_messageInfo_TrafficStatisticReq.Size(m)
}
func (m *TrafficStatisticReq) XXX_DiscardUnknown() {
	xxx_messageInfo_TrafficStatisticReq.DiscardUnknown(m)
}

var xxx_messageInfo_TrafficStatisticReq proto.InternalMessageInfo

func (m *TrafficStatisticReq) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *TrafficStatisticReq) GetBookId() int64 {
	if m != nil {
		return m.BookId
	}
	return 0
}

func (m *TrafficStatisticReq) GetChapterNum() int64 {
	if m != nil {
		return m.ChapterNum
	}
	return 0
}

func (m *TrafficStatisticReq) GetTrafficNumber() int64 {
	if m != nil {
		return m.TrafficNumber
	}
	return 0
}

type TrafficStatisticResp struct {
	Id                   int64    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	BookId               int64    `protobuf:"varint,2,opt,name=BookId,proto3" json:"BookId,omitempty"`
	ChapterNum           int64    `protobuf:"varint,3,opt,name=ChapterNum,proto3" json:"ChapterNum,omitempty"`
	TrafficNumber        int64    `protobuf:"varint,4,opt,name=TrafficNumber,proto3" json:"TrafficNumber,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TrafficStatisticResp) Reset()         { *m = TrafficStatisticResp{} }
func (m *TrafficStatisticResp) String() string { return proto.CompactTextString(m) }
func (*TrafficStatisticResp) ProtoMessage()    {}
func (*TrafficStatisticResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_59885c909ad4dfd3, []int{3}
}

func (m *TrafficStatisticResp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TrafficStatisticResp.Unmarshal(m, b)
}
func (m *TrafficStatisticResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TrafficStatisticResp.Marshal(b, m, deterministic)
}
func (m *TrafficStatisticResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TrafficStatisticResp.Merge(m, src)
}
func (m *TrafficStatisticResp) XXX_Size() int {
	return xxx_messageInfo_TrafficStatisticResp.Size(m)
}
func (m *TrafficStatisticResp) XXX_DiscardUnknown() {
	xxx_messageInfo_TrafficStatisticResp.DiscardUnknown(m)
}

var xxx_messageInfo_TrafficStatisticResp proto.InternalMessageInfo

func (m *TrafficStatisticResp) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *TrafficStatisticResp) GetBookId() int64 {
	if m != nil {
		return m.BookId
	}
	return 0
}

func (m *TrafficStatisticResp) GetChapterNum() int64 {
	if m != nil {
		return m.ChapterNum
	}
	return 0
}

func (m *TrafficStatisticResp) GetTrafficNumber() int64 {
	if m != nil {
		return m.TrafficNumber
	}
	return 0
}

type TrafficStatisticsResp struct {
	TrafficStatistics    []*TrafficStatisticResp `protobuf:"bytes,1,rep,name=TrafficStatistics,proto3" json:"TrafficStatistics,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                `json:"-"`
	XXX_unrecognized     []byte                  `json:"-"`
	XXX_sizecache        int32                   `json:"-"`
}

func (m *TrafficStatisticsResp) Reset()         { *m = TrafficStatisticsResp{} }
func (m *TrafficStatisticsResp) String() string { return proto.CompactTextString(m) }
func (*TrafficStatisticsResp) ProtoMessage()    {}
func (*TrafficStatisticsResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_59885c909ad4dfd3, []int{4}
}

func (m *TrafficStatisticsResp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TrafficStatisticsResp.Unmarshal(m, b)
}
func (m *TrafficStatisticsResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TrafficStatisticsResp.Marshal(b, m, deterministic)
}
func (m *TrafficStatisticsResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TrafficStatisticsResp.Merge(m, src)
}
func (m *TrafficStatisticsResp) XXX_Size() int {
	return xxx_messageInfo_TrafficStatisticsResp.Size(m)
}
func (m *TrafficStatisticsResp) XXX_DiscardUnknown() {
	xxx_messageInfo_TrafficStatisticsResp.DiscardUnknown(m)
}

var xxx_messageInfo_TrafficStatisticsResp proto.InternalMessageInfo

func (m *TrafficStatisticsResp) GetTrafficStatistics() []*TrafficStatisticResp {
	if m != nil {
		return m.TrafficStatistics
	}
	return nil
}

func init() {
	proto.RegisterType((*Request)(nil), "action.Request")
	proto.RegisterType((*Response)(nil), "action.Response")
	proto.RegisterType((*TrafficStatisticReq)(nil), "action.TrafficStatisticReq")
	proto.RegisterType((*TrafficStatisticResp)(nil), "action.TrafficStatisticResp")
	proto.RegisterType((*TrafficStatisticsResp)(nil), "action.TrafficStatisticsResp")
}

func init() { proto.RegisterFile("action.proto", fileDescriptor_59885c909ad4dfd3) }

var fileDescriptor_59885c909ad4dfd3 = []byte{
	// 333 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xbc, 0x93, 0x4f, 0x53, 0xf2, 0x30,
	0x10, 0xc6, 0xa7, 0x2d, 0x53, 0x78, 0xf7, 0xf5, 0x6f, 0x54, 0xcc, 0xa0, 0x38, 0x4c, 0xc7, 0x03,
	0x27, 0x0e, 0xe8, 0x17, 0x00, 0x0e, 0x1d, 0x3c, 0x70, 0x88, 0x7a, 0xf1, 0x16, 0xda, 0x05, 0x33,
	0x40, 0x13, 0x9a, 0x70, 0xf0, 0xac, 0x1f, 0xcf, 0x0f, 0xe5, 0x18, 0xc2, 0x88, 0x54, 0xb8, 0x38,
	0xe3, 0xad, 0x9b, 0x27, 0xfb, 0xf4, 0xb7, 0xbb, 0x59, 0xd8, 0xe3, 0x89, 0x11, 0x32, 0x6b, 0xa9,
	0x5c, 0x1a, 0x49, 0xc2, 0x65, 0x14, 0xd5, 0xa1, 0xcc, 0x70, 0xbe, 0x40, 0x6d, 0x08, 0x81, 0x92,
	0x12, 0xd9, 0x98, 0x7a, 0x0d, 0xaf, 0xf9, 0x8f, 0xd9, 0xef, 0xe8, 0x16, 0x2a, 0x0c, 0xb5, 0x92,
	0x99, 0x46, 0x72, 0x00, 0xbe, 0x9c, 0x58, 0xb5, 0xc2, 0x7c, 0x39, 0x21, 0x14, 0xca, 0x33, 0xd4,
	0x9a, 0x8f, 0x91, 0xfa, 0x36, 0x65, 0x15, 0x46, 0xaf, 0x1e, 0x9c, 0x3c, 0xe4, 0x7c, 0x34, 0x12,
	0xc9, 0xbd, 0xe1, 0x46, 0x68, 0x23, 0x12, 0x86, 0xf3, 0x4f, 0x07, 0x91, 0x5a, 0x87, 0x80, 0xf9,
	0x22, 0x25, 0x55, 0x08, 0xbb, 0x52, 0x4e, 0xfa, 0xa9, 0x35, 0x08, 0x98, 0x8b, 0xc8, 0x15, 0x40,
	0xef, 0x99, 0x2b, 0x83, 0xf9, 0x60, 0x31, 0xa3, 0x81, 0xd5, 0xd6, 0x4e, 0xc8, 0x35, 0xec, 0x3b,
	0xfb, 0xc1, 0x62, 0x36, 0xc4, 0x9c, 0x96, 0xec, 0x95, 0xef, 0x87, 0xd1, 0x9b, 0x07, 0xa7, 0x45,
	0x0a, 0xad, 0xfe, 0x18, 0x23, 0x81, 0xb3, 0x4d, 0x0a, 0x6d, 0x31, 0xee, 0xe0, 0xb8, 0x20, 0x50,
	0xaf, 0x11, 0x34, 0xff, 0xb7, 0x2f, 0x5b, 0x6e, 0x58, 0x3f, 0xf1, 0xb3, 0x62, 0x5a, 0xfb, 0xdd,
	0x87, 0xb0, 0x63, 0x53, 0x48, 0x1f, 0xce, 0x63, 0x34, 0x9d, 0xe9, 0xb4, 0x70, 0x8b, 0x1c, 0xae,
	0x6c, 0xdd, 0xc8, 0x6b, 0xf5, 0x6d, 0xff, 0x59, 0x12, 0x22, 0x34, 0x63, 0x34, 0x9b, 0x5a, 0xf7,
	0x65, 0xd9, 0x9f, 0x4e, 0x96, 0xae, 0x35, 0xe3, 0x62, 0x3b, 0xf2, 0xbc, 0xb6, 0xb3, 0x1e, 0x12,
	0x43, 0xb5, 0x97, 0x23, 0x37, 0xb8, 0xa9, 0xee, 0x36, 0x3d, 0xfa, 0xaa, 0xc6, 0xbd, 0xd0, 0x18,
	0xaa, 0x8f, 0x2a, 0xfd, 0xbd, 0x51, 0xb7, 0xf2, 0xe4, 0xf6, 0x63, 0x18, 0xda, 0x75, 0xb9, 0xf9,
	0x08, 0x00, 0x00, 0xff, 0xff, 0x19, 0xdc, 0x7e, 0xac, 0x3e, 0x03, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// ActionClient is the client API for Action service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ActionClient interface {
	//TrafficStatistic
	GetAllTrafficStatistics(ctx context.Context, in *Request, opts ...grpc.CallOption) (*TrafficStatisticsResp, error)
	GetTrafficStatisticByBookIdAndChapterNum(ctx context.Context, in *TrafficStatisticReq, opts ...grpc.CallOption) (*TrafficStatisticResp, error)
	CreateTrafficStatistic(ctx context.Context, in *TrafficStatisticReq, opts ...grpc.CallOption) (*Response, error)
	UpdateTrafficStatistic(ctx context.Context, in *TrafficStatisticReq, opts ...grpc.CallOption) (*Response, error)
}

type actionClient struct {
	cc *grpc.ClientConn
}

func NewActionClient(cc *grpc.ClientConn) ActionClient {
	return &actionClient{cc}
}

func (c *actionClient) GetAllTrafficStatistics(ctx context.Context, in *Request, opts ...grpc.CallOption) (*TrafficStatisticsResp, error) {
	out := new(TrafficStatisticsResp)
	err := c.cc.Invoke(ctx, "/action.Action/GetAllTrafficStatistics", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *actionClient) GetTrafficStatisticByBookIdAndChapterNum(ctx context.Context, in *TrafficStatisticReq, opts ...grpc.CallOption) (*TrafficStatisticResp, error) {
	out := new(TrafficStatisticResp)
	err := c.cc.Invoke(ctx, "/action.Action/GetTrafficStatisticByBookIdAndChapterNum", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *actionClient) CreateTrafficStatistic(ctx context.Context, in *TrafficStatisticReq, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/action.Action/CreateTrafficStatistic", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *actionClient) UpdateTrafficStatistic(ctx context.Context, in *TrafficStatisticReq, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/action.Action/UpdateTrafficStatistic", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ActionServer is the server API for Action service.
type ActionServer interface {
	//TrafficStatistic
	GetAllTrafficStatistics(context.Context, *Request) (*TrafficStatisticsResp, error)
	GetTrafficStatisticByBookIdAndChapterNum(context.Context, *TrafficStatisticReq) (*TrafficStatisticResp, error)
	CreateTrafficStatistic(context.Context, *TrafficStatisticReq) (*Response, error)
	UpdateTrafficStatistic(context.Context, *TrafficStatisticReq) (*Response, error)
}

// UnimplementedActionServer can be embedded to have forward compatible implementations.
type UnimplementedActionServer struct {
}

func (*UnimplementedActionServer) GetAllTrafficStatistics(ctx context.Context, req *Request) (*TrafficStatisticsResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAllTrafficStatistics not implemented")
}
func (*UnimplementedActionServer) GetTrafficStatisticByBookIdAndChapterNum(ctx context.Context, req *TrafficStatisticReq) (*TrafficStatisticResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTrafficStatisticByBookIdAndChapterNum not implemented")
}
func (*UnimplementedActionServer) CreateTrafficStatistic(ctx context.Context, req *TrafficStatisticReq) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateTrafficStatistic not implemented")
}
func (*UnimplementedActionServer) UpdateTrafficStatistic(ctx context.Context, req *TrafficStatisticReq) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateTrafficStatistic not implemented")
}

func RegisterActionServer(s *grpc.Server, srv ActionServer) {
	s.RegisterService(&_Action_serviceDesc, srv)
}

func _Action_GetAllTrafficStatistics_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ActionServer).GetAllTrafficStatistics(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/action.Action/GetAllTrafficStatistics",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ActionServer).GetAllTrafficStatistics(ctx, req.(*Request))
	}
	return interceptor(ctx, in, info, handler)
}

func _Action_GetTrafficStatisticByBookIdAndChapterNum_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TrafficStatisticReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ActionServer).GetTrafficStatisticByBookIdAndChapterNum(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/action.Action/GetTrafficStatisticByBookIdAndChapterNum",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ActionServer).GetTrafficStatisticByBookIdAndChapterNum(ctx, req.(*TrafficStatisticReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Action_CreateTrafficStatistic_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TrafficStatisticReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ActionServer).CreateTrafficStatistic(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/action.Action/CreateTrafficStatistic",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ActionServer).CreateTrafficStatistic(ctx, req.(*TrafficStatisticReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Action_UpdateTrafficStatistic_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TrafficStatisticReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ActionServer).UpdateTrafficStatistic(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/action.Action/UpdateTrafficStatistic",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ActionServer).UpdateTrafficStatistic(ctx, req.(*TrafficStatisticReq))
	}
	return interceptor(ctx, in, info, handler)
}

var _Action_serviceDesc = grpc.ServiceDesc{
	ServiceName: "action.Action",
	HandlerType: (*ActionServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetAllTrafficStatistics",
			Handler:    _Action_GetAllTrafficStatistics_Handler,
		},
		{
			MethodName: "GetTrafficStatisticByBookIdAndChapterNum",
			Handler:    _Action_GetTrafficStatisticByBookIdAndChapterNum_Handler,
		},
		{
			MethodName: "CreateTrafficStatistic",
			Handler:    _Action_CreateTrafficStatistic_Handler,
		},
		{
			MethodName: "UpdateTrafficStatistic",
			Handler:    _Action_UpdateTrafficStatistic_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "action.proto",
}
