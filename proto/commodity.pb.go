// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/commodity.proto

package proto

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

type GetCommodityListsReq struct {
	// 返回数量
	Limit int64 `protobuf:"varint,1,opt,name=limit,proto3" json:"limit,omitempty"`
	// 返回数据开始数
	Offset int64 `protobuf:"varint,2,opt,name=offset,proto3" json:"offset,omitempty"`
	// 检索商品ID
	ID                   int64    `protobuf:"varint,3,opt,name=ID,proto3" json:"ID,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetCommodityListsReq) Reset()         { *m = GetCommodityListsReq{} }
func (m *GetCommodityListsReq) String() string { return proto.CompactTextString(m) }
func (*GetCommodityListsReq) ProtoMessage()    {}
func (*GetCommodityListsReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_9d5890442b0a2fd1, []int{0}
}

func (m *GetCommodityListsReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetCommodityListsReq.Unmarshal(m, b)
}
func (m *GetCommodityListsReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetCommodityListsReq.Marshal(b, m, deterministic)
}
func (m *GetCommodityListsReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetCommodityListsReq.Merge(m, src)
}
func (m *GetCommodityListsReq) XXX_Size() int {
	return xxx_messageInfo_GetCommodityListsReq.Size(m)
}
func (m *GetCommodityListsReq) XXX_DiscardUnknown() {
	xxx_messageInfo_GetCommodityListsReq.DiscardUnknown(m)
}

var xxx_messageInfo_GetCommodityListsReq proto.InternalMessageInfo

func (m *GetCommodityListsReq) GetLimit() int64 {
	if m != nil {
		return m.Limit
	}
	return 0
}

func (m *GetCommodityListsReq) GetOffset() int64 {
	if m != nil {
		return m.Offset
	}
	return 0
}

func (m *GetCommodityListsReq) GetID() int64 {
	if m != nil {
		return m.ID
	}
	return 0
}

type GetCommodityListsResp struct {
	// 返回码 200:成功 500:失败
	Code int32 `protobuf:"varint,1,opt,name=Code,proto3" json:"Code,omitempty"`
	// 错误信息
	Err string `protobuf:"bytes,2,opt,name=Err,proto3" json:"Err,omitempty"`
	// 检索数据总数
	Allnum int64 `protobuf:"varint,3,opt,name=Allnum,proto3" json:"Allnum,omitempty"`
	// 数据数列
	Data                 []*CommodityInfo `protobuf:"bytes,4,rep,name=Data,proto3" json:"Data,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *GetCommodityListsResp) Reset()         { *m = GetCommodityListsResp{} }
func (m *GetCommodityListsResp) String() string { return proto.CompactTextString(m) }
func (*GetCommodityListsResp) ProtoMessage()    {}
func (*GetCommodityListsResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_9d5890442b0a2fd1, []int{1}
}

func (m *GetCommodityListsResp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetCommodityListsResp.Unmarshal(m, b)
}
func (m *GetCommodityListsResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetCommodityListsResp.Marshal(b, m, deterministic)
}
func (m *GetCommodityListsResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetCommodityListsResp.Merge(m, src)
}
func (m *GetCommodityListsResp) XXX_Size() int {
	return xxx_messageInfo_GetCommodityListsResp.Size(m)
}
func (m *GetCommodityListsResp) XXX_DiscardUnknown() {
	xxx_messageInfo_GetCommodityListsResp.DiscardUnknown(m)
}

var xxx_messageInfo_GetCommodityListsResp proto.InternalMessageInfo

func (m *GetCommodityListsResp) GetCode() int32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *GetCommodityListsResp) GetErr() string {
	if m != nil {
		return m.Err
	}
	return ""
}

func (m *GetCommodityListsResp) GetAllnum() int64 {
	if m != nil {
		return m.Allnum
	}
	return 0
}

func (m *GetCommodityListsResp) GetData() []*CommodityInfo {
	if m != nil {
		return m.Data
	}
	return nil
}

type CommodityInfo struct {
	// 商品ID
	ID int64 `protobuf:"varint,1,opt,name=ID,proto3" json:"ID,omitempty"`
	// 商品名称
	Name string `protobuf:"bytes,2,opt,name=Name,proto3" json:"Name,omitempty"`
	// 商品原价
	OldPrice string `protobuf:"bytes,3,opt,name=OldPrice,proto3" json:"OldPrice,omitempty"`
	// 商品打折价
	NewPrice string `protobuf:"bytes,4,opt,name=NewPrice,proto3" json:"NewPrice,omitempty"`
	// 打折价说明
	PriceRemark string `protobuf:"bytes,5,opt,name=PriceRemark,proto3" json:"PriceRemark,omitempty"`
	// 更新时间
	UpdateTime           int64    `protobuf:"varint,6,opt,name=updateTime,proto3" json:"updateTime,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CommodityInfo) Reset()         { *m = CommodityInfo{} }
func (m *CommodityInfo) String() string { return proto.CompactTextString(m) }
func (*CommodityInfo) ProtoMessage()    {}
func (*CommodityInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_9d5890442b0a2fd1, []int{2}
}

func (m *CommodityInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CommodityInfo.Unmarshal(m, b)
}
func (m *CommodityInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CommodityInfo.Marshal(b, m, deterministic)
}
func (m *CommodityInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CommodityInfo.Merge(m, src)
}
func (m *CommodityInfo) XXX_Size() int {
	return xxx_messageInfo_CommodityInfo.Size(m)
}
func (m *CommodityInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_CommodityInfo.DiscardUnknown(m)
}

var xxx_messageInfo_CommodityInfo proto.InternalMessageInfo

func (m *CommodityInfo) GetID() int64 {
	if m != nil {
		return m.ID
	}
	return 0
}

func (m *CommodityInfo) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *CommodityInfo) GetOldPrice() string {
	if m != nil {
		return m.OldPrice
	}
	return ""
}

func (m *CommodityInfo) GetNewPrice() string {
	if m != nil {
		return m.NewPrice
	}
	return ""
}

func (m *CommodityInfo) GetPriceRemark() string {
	if m != nil {
		return m.PriceRemark
	}
	return ""
}

func (m *CommodityInfo) GetUpdateTime() int64 {
	if m != nil {
		return m.UpdateTime
	}
	return 0
}

func init() {
	proto.RegisterType((*GetCommodityListsReq)(nil), "proto.GetCommodityListsReq")
	proto.RegisterType((*GetCommodityListsResp)(nil), "proto.GetCommodityListsResp")
	proto.RegisterType((*CommodityInfo)(nil), "proto.CommodityInfo")
}

func init() { proto.RegisterFile("proto/commodity.proto", fileDescriptor_9d5890442b0a2fd1) }

var fileDescriptor_9d5890442b0a2fd1 = []byte{
	// 334 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x51, 0x5d, 0x4a, 0xf3, 0x40,
	0x14, 0x25, 0x7f, 0xe5, 0xeb, 0x2d, 0x9f, 0xe8, 0x25, 0x95, 0x10, 0x8b, 0x94, 0x3c, 0xf5, 0xa9,
	0x81, 0xba, 0x02, 0x69, 0x45, 0x0a, 0x52, 0x65, 0xe8, 0x06, 0xc6, 0x66, 0x5a, 0x46, 0x33, 0x99,
	0x98, 0x4c, 0x29, 0x3e, 0x09, 0x6e, 0xc1, 0x75, 0xb8, 0x1a, 0xb7, 0xe0, 0x42, 0x24, 0x37, 0x69,
	0xa9, 0xb4, 0x3e, 0xe5, 0x9e, 0x73, 0x66, 0xce, 0x39, 0x37, 0x03, 0xdd, 0xbc, 0xd0, 0x46, 0xc7,
	0x0b, 0xad, 0x94, 0x4e, 0xa4, 0x79, 0x1d, 0x12, 0x46, 0x8f, 0x3e, 0x61, 0x6f, 0xa5, 0xf5, 0x2a,
	0x15, 0x31, 0xcf, 0x65, 0xcc, 0xb3, 0x4c, 0x1b, 0x6e, 0xa4, 0xce, 0xca, 0xfa, 0x50, 0x34, 0x07,
	0xff, 0x56, 0x98, 0xf1, 0xf6, 0xea, 0x9d, 0x2c, 0x4d, 0xc9, 0xc4, 0x0b, 0xfa, 0xe0, 0xa5, 0x52,
	0x49, 0x13, 0x58, 0x7d, 0x6b, 0xe0, 0xb0, 0x1a, 0xe0, 0x39, 0xb4, 0xf4, 0x72, 0x59, 0x0a, 0x13,
	0xd8, 0x44, 0x37, 0x08, 0x4f, 0xc0, 0x9e, 0x4e, 0x02, 0x87, 0x38, 0x7b, 0x3a, 0x89, 0xde, 0xa0,
	0x7b, 0xc4, 0xb5, 0xcc, 0x11, 0xc1, 0x1d, 0xeb, 0x44, 0x90, 0xab, 0xc7, 0x68, 0xc6, 0x53, 0x70,
	0x6e, 0x8a, 0x82, 0x1c, 0xdb, 0xac, 0x1a, 0xab, 0x98, 0xeb, 0x34, 0xcd, 0xd6, 0xaa, 0xb1, 0x6c,
	0x10, 0x0e, 0xc0, 0x9d, 0x70, 0xc3, 0x03, 0xb7, 0xef, 0x0c, 0x3a, 0x23, 0xbf, 0x5e, 0x61, 0xb8,
	0x8b, 0x99, 0x66, 0x4b, 0xcd, 0xe8, 0x44, 0xf4, 0x69, 0xc1, 0xff, 0x5f, 0x7c, 0x53, 0xd1, 0xda,
	0x56, 0xac, 0x9a, 0xcc, 0xb8, 0x12, 0x4d, 0x2c, 0xcd, 0x18, 0xc2, 0xbf, 0xfb, 0x34, 0x79, 0x28,
	0xe4, 0x42, 0x50, 0x72, 0x9b, 0xed, 0x70, 0xa5, 0xcd, 0xc4, 0xa6, 0xd6, 0xdc, 0x5a, 0xdb, 0x62,
	0xec, 0x43, 0x87, 0x06, 0x26, 0x14, 0x2f, 0x9e, 0x03, 0x8f, 0xe4, 0x7d, 0x0a, 0x2f, 0x01, 0xd6,
	0x79, 0xc2, 0x8d, 0x98, 0x4b, 0x25, 0x82, 0x16, 0xb5, 0xd8, 0x63, 0x46, 0x1b, 0x68, 0xef, 0xea,
	0xe2, 0x13, 0x9c, 0x1d, 0xfc, 0x3d, 0xbc, 0x68, 0xb6, 0x3d, 0xf6, 0x5a, 0x61, 0xef, 0x6f, 0xb1,
	0xcc, 0xa3, 0xf0, 0xfd, 0xeb, 0xfb, 0xc3, 0xf6, 0x11, 0xe3, 0x03, 0xfd, 0xb1, 0x45, 0x17, 0xaf,
	0x7e, 0x02, 0x00, 0x00, 0xff, 0xff, 0x80, 0xf7, 0xcd, 0x81, 0x44, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// CommodityClient is the client API for Commodity service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type CommodityClient interface {
	//获取商品列表
	GetCommodityLists(ctx context.Context, in *GetCommodityListsReq, opts ...grpc.CallOption) (*GetCommodityListsResp, error)
}

type commodityClient struct {
	cc *grpc.ClientConn
}

func NewCommodityClient(cc *grpc.ClientConn) CommodityClient {
	return &commodityClient{cc}
}

func (c *commodityClient) GetCommodityLists(ctx context.Context, in *GetCommodityListsReq, opts ...grpc.CallOption) (*GetCommodityListsResp, error) {
	out := new(GetCommodityListsResp)
	err := c.cc.Invoke(ctx, "/proto.Commodity/GetCommodityLists", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CommodityServer is the server API for Commodity service.
type CommodityServer interface {
	//获取商品列表
	GetCommodityLists(context.Context, *GetCommodityListsReq) (*GetCommodityListsResp, error)
}

func RegisterCommodityServer(s *grpc.Server, srv CommodityServer) {
	s.RegisterService(&_Commodity_serviceDesc, srv)
}

func _Commodity_GetCommodityLists_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetCommodityListsReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CommodityServer).GetCommodityLists(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Commodity/GetCommodityLists",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CommodityServer).GetCommodityLists(ctx, req.(*GetCommodityListsReq))
	}
	return interceptor(ctx, in, info, handler)
}

var _Commodity_serviceDesc = grpc.ServiceDesc{
	ServiceName: "proto.Commodity",
	HandlerType: (*CommodityServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetCommodityLists",
			Handler:    _Commodity_GetCommodityLists_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/commodity.proto",
}
