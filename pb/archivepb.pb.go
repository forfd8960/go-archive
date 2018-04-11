// Code generated by protoc-gen-go.
// source: archivepb.proto
// DO NOT EDIT!

/*
Package archivepb is a generated protocol buffer package.

It is generated from these files:
	archivepb.proto

It has these top-level messages:
	ArchiveItem
	GetArchiveListReq
	GetArchiveListRes
*/
package archivepb

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

type ArchiveItem struct {
	FromHost  string `protobuf:"bytes,1,opt,name=from_host,json=fromHost" json:"from_host,omitempty"`
	ItemName  string `protobuf:"bytes,2,opt,name=item_name,json=itemName" json:"item_name,omitempty"`
	ItemUrl   string `protobuf:"bytes,3,opt,name=item_url,json=itemUrl" json:"item_url,omitempty"`
	CollectAt string `protobuf:"bytes,4,opt,name=collect_at,json=collectAt" json:"collect_at,omitempty"`
}

func (m *ArchiveItem) Reset()                    { *m = ArchiveItem{} }
func (m *ArchiveItem) String() string            { return proto.CompactTextString(m) }
func (*ArchiveItem) ProtoMessage()               {}
func (*ArchiveItem) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *ArchiveItem) GetFromHost() string {
	if m != nil {
		return m.FromHost
	}
	return ""
}

func (m *ArchiveItem) GetItemName() string {
	if m != nil {
		return m.ItemName
	}
	return ""
}

func (m *ArchiveItem) GetItemUrl() string {
	if m != nil {
		return m.ItemUrl
	}
	return ""
}

func (m *ArchiveItem) GetCollectAt() string {
	if m != nil {
		return m.CollectAt
	}
	return ""
}

type GetArchiveListReq struct {
	Offset int32 `protobuf:"varint,1,opt,name=offset" json:"offset,omitempty"`
	Limit  int32 `protobuf:"varint,2,opt,name=limit" json:"limit,omitempty"`
}

func (m *GetArchiveListReq) Reset()                    { *m = GetArchiveListReq{} }
func (m *GetArchiveListReq) String() string            { return proto.CompactTextString(m) }
func (*GetArchiveListReq) ProtoMessage()               {}
func (*GetArchiveListReq) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *GetArchiveListReq) GetOffset() int32 {
	if m != nil {
		return m.Offset
	}
	return 0
}

func (m *GetArchiveListReq) GetLimit() int32 {
	if m != nil {
		return m.Limit
	}
	return 0
}

type GetArchiveListRes struct {
	ArchiveItems []*ArchiveItem `protobuf:"bytes,1,rep,name=archive_items,json=archiveItems" json:"archive_items,omitempty"`
	TotalCount   int32          `protobuf:"varint,2,opt,name=total_count,json=totalCount" json:"total_count,omitempty"`
}

func (m *GetArchiveListRes) Reset()                    { *m = GetArchiveListRes{} }
func (m *GetArchiveListRes) String() string            { return proto.CompactTextString(m) }
func (*GetArchiveListRes) ProtoMessage()               {}
func (*GetArchiveListRes) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *GetArchiveListRes) GetArchiveItems() []*ArchiveItem {
	if m != nil {
		return m.ArchiveItems
	}
	return nil
}

func (m *GetArchiveListRes) GetTotalCount() int32 {
	if m != nil {
		return m.TotalCount
	}
	return 0
}

func init() {
	proto.RegisterType((*ArchiveItem)(nil), "archivepb.ArchiveItem")
	proto.RegisterType((*GetArchiveListReq)(nil), "archivepb.GetArchiveListReq")
	proto.RegisterType((*GetArchiveListRes)(nil), "archivepb.GetArchiveListRes")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for GoArchive service

type GoArchiveClient interface {
	GetArchiveList(ctx context.Context, in *GetArchiveListReq, opts ...grpc.CallOption) (*GetArchiveListRes, error)
}

type goArchiveClient struct {
	cc *grpc.ClientConn
}

func NewGoArchiveClient(cc *grpc.ClientConn) GoArchiveClient {
	return &goArchiveClient{cc}
}

func (c *goArchiveClient) GetArchiveList(ctx context.Context, in *GetArchiveListReq, opts ...grpc.CallOption) (*GetArchiveListRes, error) {
	out := new(GetArchiveListRes)
	err := grpc.Invoke(ctx, "/archivepb.go_archive/GetArchiveList", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for GoArchive service

type GoArchiveServer interface {
	GetArchiveList(context.Context, *GetArchiveListReq) (*GetArchiveListRes, error)
}

func RegisterGoArchiveServer(s *grpc.Server, srv GoArchiveServer) {
	s.RegisterService(&_GoArchive_serviceDesc, srv)
}

func _GoArchive_GetArchiveList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetArchiveListReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GoArchiveServer).GetArchiveList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/archivepb.go_archive/GetArchiveList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GoArchiveServer).GetArchiveList(ctx, req.(*GetArchiveListReq))
	}
	return interceptor(ctx, in, info, handler)
}

var _GoArchive_serviceDesc = grpc.ServiceDesc{
	ServiceName: "archivepb.go_archive",
	HandlerType: (*GoArchiveServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetArchiveList",
			Handler:    _GoArchive_GetArchiveList_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "archivepb.proto",
}

func init() { proto.RegisterFile("archivepb.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 272 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x91, 0xcd, 0x4b, 0xc3, 0x40,
	0x14, 0xc4, 0x8d, 0x35, 0xd5, 0xbc, 0xf8, 0x81, 0x0f, 0x29, 0xf1, 0x0b, 0x4b, 0x4e, 0x3d, 0xf5,
	0x50, 0x8f, 0x9e, 0x82, 0x07, 0x15, 0xa4, 0x87, 0x80, 0x37, 0x61, 0xd9, 0x86, 0x8d, 0x0d, 0xec,
	0xf6, 0xb5, 0xd9, 0x57, 0xef, 0xfe, 0xe7, 0xb2, 0x9b, 0xd5, 0x56, 0x0a, 0xde, 0x32, 0xbf, 0x21,
	0x3b, 0x33, 0xbb, 0x70, 0x26, 0xdb, 0x6a, 0xde, 0x7c, 0xaa, 0xe5, 0x6c, 0xbc, 0x6c, 0x89, 0x09,
	0x93, 0x5f, 0x90, 0x7f, 0x45, 0x90, 0x16, 0x9d, 0x7a, 0x61, 0x65, 0xf0, 0x1a, 0x92, 0xba, 0x25,
	0x23, 0xe6, 0x64, 0x39, 0x8b, 0x86, 0xd1, 0x28, 0x29, 0x8f, 0x1c, 0x78, 0x26, 0xcb, 0xce, 0x6c,
	0x58, 0x19, 0xb1, 0x90, 0x46, 0x65, 0xfb, 0x9d, 0xe9, 0xc0, 0x54, 0x1a, 0x85, 0x97, 0xe0, 0xbf,
	0xc5, 0xba, 0xd5, 0x59, 0xcf, 0x7b, 0x87, 0x4e, 0xbf, 0xb5, 0x1a, 0x6f, 0x01, 0x2a, 0xd2, 0x5a,
	0x55, 0x2c, 0x24, 0x67, 0x07, 0xde, 0x4c, 0x02, 0x29, 0x38, 0x2f, 0xe0, 0xfc, 0x49, 0x71, 0x68,
	0xf1, 0xda, 0x58, 0x2e, 0xd5, 0x0a, 0x07, 0xd0, 0xa7, 0xba, 0xb6, 0xaa, 0x6b, 0x11, 0x97, 0x41,
	0xe1, 0x05, 0xc4, 0xba, 0x31, 0x0d, 0xfb, 0xfc, 0xb8, 0xec, 0x44, 0xbe, 0xda, 0x3d, 0xc2, 0xe2,
	0x03, 0x9c, 0x84, 0xa1, 0xc2, 0x35, 0xb1, 0x59, 0x34, 0xec, 0x8d, 0xd2, 0xc9, 0x60, 0xbc, 0xb9,
	0x8f, 0xad, 0xe9, 0xe5, 0xb1, 0xdc, 0x08, 0x8b, 0x77, 0x90, 0x32, 0xb1, 0xd4, 0xa2, 0xa2, 0xf5,
	0xe2, 0x27, 0x0d, 0x3c, 0x7a, 0x74, 0x64, 0xf2, 0x0e, 0xf0, 0x41, 0x22, 0xfc, 0x83, 0x53, 0x38,
	0xfd, 0x5b, 0x00, 0x6f, 0xb6, 0x62, 0x76, 0xe6, 0x5d, 0xfd, 0xe7, 0xda, 0x7c, 0x6f, 0xd6, 0xf7,
	0x2f, 0x75, 0xff, 0x1d, 0x00, 0x00, 0xff, 0xff, 0xba, 0xdb, 0x57, 0xa7, 0xbc, 0x01, 0x00, 0x00,
}
