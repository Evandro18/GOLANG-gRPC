// Code generated by protoc-gen-go. DO NOT EDIT.
// source: service.proto

package protos

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
	Sale                 *Sale    `protobuf:"bytes,1,opt,name=Sale,proto3" json:"Sale,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Request) Reset()         { *m = Request{} }
func (m *Request) String() string { return proto.CompactTextString(m) }
func (*Request) ProtoMessage()    {}
func (*Request) Descriptor() ([]byte, []int) {
	return fileDescriptor_a0b84a42fa06f626, []int{0}
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

func (m *Request) GetSale() *Sale {
	if m != nil {
		return m.Sale
	}
	return nil
}

type RequestById struct {
	Id                   string   `protobuf:"bytes,1,opt,name=Id,proto3" json:"Id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RequestById) Reset()         { *m = RequestById{} }
func (m *RequestById) String() string { return proto.CompactTextString(m) }
func (*RequestById) ProtoMessage()    {}
func (*RequestById) Descriptor() ([]byte, []int) {
	return fileDescriptor_a0b84a42fa06f626, []int{1}
}

func (m *RequestById) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RequestById.Unmarshal(m, b)
}
func (m *RequestById) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RequestById.Marshal(b, m, deterministic)
}
func (m *RequestById) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RequestById.Merge(m, src)
}
func (m *RequestById) XXX_Size() int {
	return xxx_messageInfo_RequestById.Size(m)
}
func (m *RequestById) XXX_DiscardUnknown() {
	xxx_messageInfo_RequestById.DiscardUnknown(m)
}

var xxx_messageInfo_RequestById proto.InternalMessageInfo

func (m *RequestById) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

type Response struct {
	Sale                 *Sale    `protobuf:"bytes,1,opt,name=Sale,proto3" json:"Sale,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Response) Reset()         { *m = Response{} }
func (m *Response) String() string { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()    {}
func (*Response) Descriptor() ([]byte, []int) {
	return fileDescriptor_a0b84a42fa06f626, []int{2}
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

func (m *Response) GetSale() *Sale {
	if m != nil {
		return m.Sale
	}
	return nil
}

type Sale struct {
	Id                   string   `protobuf:"bytes,1,opt,name=Id,proto3" json:"Id,omitempty"`
	Product              string   `protobuf:"bytes,2,opt,name=Product,proto3" json:"Product,omitempty"`
	Quantity             int64    `protobuf:"varint,3,opt,name=Quantity,proto3" json:"Quantity,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Sale) Reset()         { *m = Sale{} }
func (m *Sale) String() string { return proto.CompactTextString(m) }
func (*Sale) ProtoMessage()    {}
func (*Sale) Descriptor() ([]byte, []int) {
	return fileDescriptor_a0b84a42fa06f626, []int{3}
}

func (m *Sale) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Sale.Unmarshal(m, b)
}
func (m *Sale) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Sale.Marshal(b, m, deterministic)
}
func (m *Sale) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Sale.Merge(m, src)
}
func (m *Sale) XXX_Size() int {
	return xxx_messageInfo_Sale.Size(m)
}
func (m *Sale) XXX_DiscardUnknown() {
	xxx_messageInfo_Sale.DiscardUnknown(m)
}

var xxx_messageInfo_Sale proto.InternalMessageInfo

func (m *Sale) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Sale) GetProduct() string {
	if m != nil {
		return m.Product
	}
	return ""
}

func (m *Sale) GetQuantity() int64 {
	if m != nil {
		return m.Quantity
	}
	return 0
}

type ListSalesRes struct {
	Sale                 *Sale    `protobuf:"bytes,1,opt,name=Sale,proto3" json:"Sale,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListSalesRes) Reset()         { *m = ListSalesRes{} }
func (m *ListSalesRes) String() string { return proto.CompactTextString(m) }
func (*ListSalesRes) ProtoMessage()    {}
func (*ListSalesRes) Descriptor() ([]byte, []int) {
	return fileDescriptor_a0b84a42fa06f626, []int{4}
}

func (m *ListSalesRes) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListSalesRes.Unmarshal(m, b)
}
func (m *ListSalesRes) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListSalesRes.Marshal(b, m, deterministic)
}
func (m *ListSalesRes) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListSalesRes.Merge(m, src)
}
func (m *ListSalesRes) XXX_Size() int {
	return xxx_messageInfo_ListSalesRes.Size(m)
}
func (m *ListSalesRes) XXX_DiscardUnknown() {
	xxx_messageInfo_ListSalesRes.DiscardUnknown(m)
}

var xxx_messageInfo_ListSalesRes proto.InternalMessageInfo

func (m *ListSalesRes) GetSale() *Sale {
	if m != nil {
		return m.Sale
	}
	return nil
}

type ListSalesReq struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListSalesReq) Reset()         { *m = ListSalesReq{} }
func (m *ListSalesReq) String() string { return proto.CompactTextString(m) }
func (*ListSalesReq) ProtoMessage()    {}
func (*ListSalesReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_a0b84a42fa06f626, []int{5}
}

func (m *ListSalesReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListSalesReq.Unmarshal(m, b)
}
func (m *ListSalesReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListSalesReq.Marshal(b, m, deterministic)
}
func (m *ListSalesReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListSalesReq.Merge(m, src)
}
func (m *ListSalesReq) XXX_Size() int {
	return xxx_messageInfo_ListSalesReq.Size(m)
}
func (m *ListSalesReq) XXX_DiscardUnknown() {
	xxx_messageInfo_ListSalesReq.DiscardUnknown(m)
}

var xxx_messageInfo_ListSalesReq proto.InternalMessageInfo

type Success struct {
	Sucess               bool     `protobuf:"varint,1,opt,name=Sucess,proto3" json:"Sucess,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Success) Reset()         { *m = Success{} }
func (m *Success) String() string { return proto.CompactTextString(m) }
func (*Success) ProtoMessage()    {}
func (*Success) Descriptor() ([]byte, []int) {
	return fileDescriptor_a0b84a42fa06f626, []int{6}
}

func (m *Success) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Success.Unmarshal(m, b)
}
func (m *Success) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Success.Marshal(b, m, deterministic)
}
func (m *Success) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Success.Merge(m, src)
}
func (m *Success) XXX_Size() int {
	return xxx_messageInfo_Success.Size(m)
}
func (m *Success) XXX_DiscardUnknown() {
	xxx_messageInfo_Success.DiscardUnknown(m)
}

var xxx_messageInfo_Success proto.InternalMessageInfo

func (m *Success) GetSucess() bool {
	if m != nil {
		return m.Sucess
	}
	return false
}

func init() {
	proto.RegisterType((*Request)(nil), "protos.Request")
	proto.RegisterType((*RequestById)(nil), "protos.RequestById")
	proto.RegisterType((*Response)(nil), "protos.Response")
	proto.RegisterType((*Sale)(nil), "protos.Sale")
	proto.RegisterType((*ListSalesRes)(nil), "protos.ListSalesRes")
	proto.RegisterType((*ListSalesReq)(nil), "protos.ListSalesReq")
	proto.RegisterType((*Success)(nil), "protos.Success")
}

func init() { proto.RegisterFile("service.proto", fileDescriptor_a0b84a42fa06f626) }

var fileDescriptor_a0b84a42fa06f626 = []byte{
	// 272 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x51, 0x41, 0x4b, 0xf3, 0x40,
	0x10, 0x65, 0xd3, 0x8f, 0x26, 0x9d, 0xe6, 0xab, 0x32, 0x8a, 0x84, 0x80, 0x10, 0xf7, 0x54, 0xb0,
	0x96, 0x52, 0x4f, 0x1e, 0xf5, 0x16, 0xe8, 0x41, 0x37, 0xbf, 0xa0, 0x66, 0xe7, 0x10, 0x90, 0xa6,
	0xcd, 0xec, 0x0a, 0xfd, 0x43, 0xfe, 0x4e, 0xd9, 0x4d, 0x52, 0x6d, 0xf1, 0xd0, 0xd3, 0xee, 0xbc,
	0x79, 0xf3, 0xde, 0xce, 0x5b, 0xf8, 0xcf, 0xd4, 0x7c, 0x56, 0x25, 0xcd, 0xb7, 0x4d, 0x6d, 0x6a,
	0x1c, 0xfa, 0x83, 0xe5, 0x3d, 0x84, 0x8a, 0x76, 0x96, 0xd8, 0x60, 0x06, 0xff, 0x8a, 0xf5, 0x07,
	0x25, 0x22, 0x13, 0xd3, 0xf1, 0x32, 0x6e, 0x89, 0x3c, 0x77, 0x98, 0xf2, 0x1d, 0x79, 0x0b, 0xe3,
	0x8e, 0xfc, 0xb2, 0xcf, 0x35, 0x4e, 0x20, 0xc8, 0xb5, 0xa7, 0x8f, 0x54, 0x90, 0x6b, 0x39, 0x83,
	0x48, 0x11, 0x6f, 0xeb, 0x0d, 0xd3, 0x19, 0x62, 0xab, 0x96, 0x71, 0xaa, 0x82, 0x09, 0x84, 0xaf,
	0x4d, 0xad, 0x6d, 0x69, 0x92, 0xc0, 0x83, 0x7d, 0x89, 0x29, 0x44, 0x6f, 0x76, 0xbd, 0x31, 0x95,
	0xd9, 0x27, 0x83, 0x4c, 0x4c, 0x07, 0xea, 0x50, 0xcb, 0x05, 0xc4, 0xab, 0x8a, 0x8d, 0x53, 0x64,
	0x45, 0x7c, 0x86, 0xff, 0xe4, 0x68, 0x62, 0x27, 0xef, 0x20, 0x2c, 0x6c, 0x59, 0x12, 0x33, 0xde,
	0xc0, 0xb0, 0xb0, 0xee, 0xe6, 0xc7, 0x23, 0xd5, 0x55, 0xcb, 0x2f, 0x01, 0xb1, 0xe7, 0x17, 0x6d,
	0x96, 0x38, 0x83, 0xf0, 0x59, 0x6b, 0xbf, 0xc6, 0x45, 0x6f, 0xd1, 0x25, 0x94, 0x5e, 0xfe, 0x00,
	0x5d, 0x26, 0x4f, 0x30, 0x3a, 0x38, 0xe2, 0x75, 0xdf, 0xfe, 0xfd, 0x88, 0xf4, 0x2f, 0x94, 0x17,
	0x02, 0x1f, 0x20, 0x72, 0x88, 0x8f, 0xfd, 0xea, 0xc4, 0xc9, 0x81, 0xe9, 0xd1, 0x86, 0xef, 0xed,
	0xef, 0x3e, 0x7e, 0x07, 0x00, 0x00, 0xff, 0xff, 0xd0, 0xc3, 0xc0, 0x33, 0xf5, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// SalesServiceClient is the client API for SalesService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type SalesServiceClient interface {
	AddSale(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error)
	ListSales(ctx context.Context, in *ListSalesReq, opts ...grpc.CallOption) (SalesService_ListSalesClient, error)
	ListById(ctx context.Context, in *RequestById, opts ...grpc.CallOption) (*Sale, error)
}

type salesServiceClient struct {
	cc *grpc.ClientConn
}

func NewSalesServiceClient(cc *grpc.ClientConn) SalesServiceClient {
	return &salesServiceClient{cc}
}

func (c *salesServiceClient) AddSale(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/protos.SalesService/AddSale", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *salesServiceClient) ListSales(ctx context.Context, in *ListSalesReq, opts ...grpc.CallOption) (SalesService_ListSalesClient, error) {
	stream, err := c.cc.NewStream(ctx, &_SalesService_serviceDesc.Streams[0], "/protos.SalesService/ListSales", opts...)
	if err != nil {
		return nil, err
	}
	x := &salesServiceListSalesClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type SalesService_ListSalesClient interface {
	Recv() (*ListSalesRes, error)
	grpc.ClientStream
}

type salesServiceListSalesClient struct {
	grpc.ClientStream
}

func (x *salesServiceListSalesClient) Recv() (*ListSalesRes, error) {
	m := new(ListSalesRes)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *salesServiceClient) ListById(ctx context.Context, in *RequestById, opts ...grpc.CallOption) (*Sale, error) {
	out := new(Sale)
	err := c.cc.Invoke(ctx, "/protos.SalesService/ListById", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SalesServiceServer is the server API for SalesService service.
type SalesServiceServer interface {
	AddSale(context.Context, *Request) (*Response, error)
	ListSales(*ListSalesReq, SalesService_ListSalesServer) error
	ListById(context.Context, *RequestById) (*Sale, error)
}

// UnimplementedSalesServiceServer can be embedded to have forward compatible implementations.
type UnimplementedSalesServiceServer struct {
}

func (*UnimplementedSalesServiceServer) AddSale(ctx context.Context, req *Request) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddSale not implemented")
}
func (*UnimplementedSalesServiceServer) ListSales(req *ListSalesReq, srv SalesService_ListSalesServer) error {
	return status.Errorf(codes.Unimplemented, "method ListSales not implemented")
}
func (*UnimplementedSalesServiceServer) ListById(ctx context.Context, req *RequestById) (*Sale, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListById not implemented")
}

func RegisterSalesServiceServer(s *grpc.Server, srv SalesServiceServer) {
	s.RegisterService(&_SalesService_serviceDesc, srv)
}

func _SalesService_AddSale_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SalesServiceServer).AddSale(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protos.SalesService/AddSale",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SalesServiceServer).AddSale(ctx, req.(*Request))
	}
	return interceptor(ctx, in, info, handler)
}

func _SalesService_ListSales_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(ListSalesReq)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(SalesServiceServer).ListSales(m, &salesServiceListSalesServer{stream})
}

type SalesService_ListSalesServer interface {
	Send(*ListSalesRes) error
	grpc.ServerStream
}

type salesServiceListSalesServer struct {
	grpc.ServerStream
}

func (x *salesServiceListSalesServer) Send(m *ListSalesRes) error {
	return x.ServerStream.SendMsg(m)
}

func _SalesService_ListById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RequestById)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SalesServiceServer).ListById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protos.SalesService/ListById",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SalesServiceServer).ListById(ctx, req.(*RequestById))
	}
	return interceptor(ctx, in, info, handler)
}

var _SalesService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "protos.SalesService",
	HandlerType: (*SalesServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AddSale",
			Handler:    _SalesService_AddSale_Handler,
		},
		{
			MethodName: "ListById",
			Handler:    _SalesService_ListById_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "ListSales",
			Handler:       _SalesService_ListSales_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "service.proto",
}
