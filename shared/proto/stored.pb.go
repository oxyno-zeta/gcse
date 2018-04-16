// Code generated by protoc-gen-go. DO NOT EDIT.
// source: github.com/daviddengcn/gcse/shared/proto/stored.proto

package gcsepb

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

type PackageCrawlHistoryReq struct {
	Package string `protobuf:"bytes,1,opt,name=package" json:"package,omitempty"`
}

func (m *PackageCrawlHistoryReq) Reset()                    { *m = PackageCrawlHistoryReq{} }
func (m *PackageCrawlHistoryReq) String() string            { return proto.CompactTextString(m) }
func (*PackageCrawlHistoryReq) ProtoMessage()               {}
func (*PackageCrawlHistoryReq) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{0} }

func (m *PackageCrawlHistoryReq) GetPackage() string {
	if m != nil {
		return m.Package
	}
	return ""
}

type PackageCrawlHistoryResp struct {
	Info *HistoryInfo `protobuf:"bytes,1,opt,name=info" json:"info,omitempty"`
}

func (m *PackageCrawlHistoryResp) Reset()                    { *m = PackageCrawlHistoryResp{} }
func (m *PackageCrawlHistoryResp) String() string            { return proto.CompactTextString(m) }
func (*PackageCrawlHistoryResp) ProtoMessage()               {}
func (*PackageCrawlHistoryResp) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{1} }

func (m *PackageCrawlHistoryResp) GetInfo() *HistoryInfo {
	if m != nil {
		return m.Info
	}
	return nil
}

func init() {
	proto.RegisterType((*PackageCrawlHistoryReq)(nil), "gcse.PackageCrawlHistoryReq")
	proto.RegisterType((*PackageCrawlHistoryResp)(nil), "gcse.PackageCrawlHistoryResp")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for StoreService service

type StoreServiceClient interface {
	PackageCrawlHistory(ctx context.Context, in *PackageCrawlHistoryReq, opts ...grpc.CallOption) (*PackageCrawlHistoryResp, error)
}

type storeServiceClient struct {
	cc *grpc.ClientConn
}

func NewStoreServiceClient(cc *grpc.ClientConn) StoreServiceClient {
	return &storeServiceClient{cc}
}

func (c *storeServiceClient) PackageCrawlHistory(ctx context.Context, in *PackageCrawlHistoryReq, opts ...grpc.CallOption) (*PackageCrawlHistoryResp, error) {
	out := new(PackageCrawlHistoryResp)
	err := grpc.Invoke(ctx, "/gcse.StoreService/PackageCrawlHistory", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for StoreService service

type StoreServiceServer interface {
	PackageCrawlHistory(context.Context, *PackageCrawlHistoryReq) (*PackageCrawlHistoryResp, error)
}

func RegisterStoreServiceServer(s *grpc.Server, srv StoreServiceServer) {
	s.RegisterService(&_StoreService_serviceDesc, srv)
}

func _StoreService_PackageCrawlHistory_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PackageCrawlHistoryReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StoreServiceServer).PackageCrawlHistory(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/gcse.StoreService/PackageCrawlHistory",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StoreServiceServer).PackageCrawlHistory(ctx, req.(*PackageCrawlHistoryReq))
	}
	return interceptor(ctx, in, info, handler)
}

var _StoreService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "gcse.StoreService",
	HandlerType: (*StoreServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "PackageCrawlHistory",
			Handler:    _StoreService_PackageCrawlHistory_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "github.com/daviddengcn/gcse/shared/proto/stored.proto",
}

func init() {
	proto.RegisterFile("github.com/daviddengcn/gcse/shared/proto/stored.proto", fileDescriptor2)
}

var fileDescriptor2 = []byte{
	// 210 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x32, 0x4d, 0xcf, 0x2c, 0xc9,
	0x28, 0x4d, 0xd2, 0x4b, 0xce, 0xcf, 0xd5, 0x4f, 0x49, 0x2c, 0xcb, 0x4c, 0x49, 0x49, 0xcd, 0x4b,
	0x4f, 0xce, 0xd3, 0x4f, 0x4f, 0x2e, 0x4e, 0xd5, 0x2f, 0xce, 0x48, 0x2c, 0x4a, 0x4d, 0xd1, 0x2f,
	0x28, 0xca, 0x2f, 0xc9, 0xd7, 0x2f, 0x2e, 0xc9, 0x2f, 0x4a, 0x4d, 0xd1, 0x03, 0x73, 0x84, 0x58,
	0x40, 0xf2, 0x52, 0x24, 0x68, 0x2e, 0xc8, 0x4c, 0x49, 0x2d, 0x82, 0x68, 0x56, 0x32, 0xe2, 0x12,
	0x0b, 0x48, 0x4c, 0xce, 0x4e, 0x4c, 0x4f, 0x75, 0x2e, 0x4a, 0x2c, 0xcf, 0xf1, 0xc8, 0x04, 0x19,
	0x5d, 0x19, 0x94, 0x5a, 0x28, 0x24, 0xc1, 0xc5, 0x5e, 0x00, 0x91, 0x91, 0x60, 0x54, 0x60, 0xd4,
	0xe0, 0x0c, 0x82, 0x71, 0x95, 0x1c, 0xb8, 0xc4, 0xb1, 0xea, 0x29, 0x2e, 0x10, 0x52, 0xe5, 0x62,
	0xc9, 0xcc, 0x4b, 0xcb, 0x07, 0xeb, 0xe0, 0x36, 0x12, 0xd4, 0x03, 0xd9, 0xae, 0x07, 0x55, 0xe0,
	0x99, 0x97, 0x96, 0x1f, 0x04, 0x96, 0x36, 0x4a, 0xe2, 0xe2, 0x09, 0x06, 0x79, 0x21, 0x38, 0xb5,
	0xa8, 0x2c, 0x33, 0x39, 0x55, 0x28, 0x88, 0x4b, 0x18, 0x8b, 0x89, 0x42, 0x32, 0x10, 0xfd, 0xd8,
	0x1d, 0x28, 0x25, 0x8b, 0x47, 0xb6, 0xb8, 0xc0, 0x89, 0x23, 0x8a, 0x0d, 0x24, 0x5f, 0x90, 0x94,
	0xc4, 0x06, 0xf6, 0xaa, 0x31, 0x20, 0x00, 0x00, 0xff, 0xff, 0x4d, 0x64, 0x41, 0x27, 0x60, 0x01,
	0x00, 0x00,
}
