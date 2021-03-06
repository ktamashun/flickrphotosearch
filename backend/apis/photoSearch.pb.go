// Code generated by protoc-gen-go. DO NOT EDIT.
// source: photoSearch.proto

/*
Package apis is a generated protocol buffer package.

It is generated from these files:
	photoSearch.proto

It has these top-level messages:
	PhotoSearchRequest
	PhotoSearchResponse
*/
package apis

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

type PhotoSearchRequest struct {
	Query string `protobuf:"bytes,1,opt,name=query" json:"query,omitempty"`
	Page  int32  `protobuf:"varint,2,opt,name=page" json:"page,omitempty"`
}

func (m *PhotoSearchRequest) Reset()                    { *m = PhotoSearchRequest{} }
func (m *PhotoSearchRequest) String() string            { return proto.CompactTextString(m) }
func (*PhotoSearchRequest) ProtoMessage()               {}
func (*PhotoSearchRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *PhotoSearchRequest) GetQuery() string {
	if m != nil {
		return m.Query
	}
	return ""
}

func (m *PhotoSearchRequest) GetPage() int32 {
	if m != nil {
		return m.Page
	}
	return 0
}

type PhotoSearchResponse struct {
	PageCount int32                        `protobuf:"varint,1,opt,name=pageCount" json:"pageCount,omitempty"`
	Total     string                       `protobuf:"bytes,2,opt,name=total" json:"total,omitempty"`
	Photos    []*PhotoSearchResponse_Photo `protobuf:"bytes,3,rep,name=photos" json:"photos,omitempty"`
}

func (m *PhotoSearchResponse) Reset()                    { *m = PhotoSearchResponse{} }
func (m *PhotoSearchResponse) String() string            { return proto.CompactTextString(m) }
func (*PhotoSearchResponse) ProtoMessage()               {}
func (*PhotoSearchResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *PhotoSearchResponse) GetPageCount() int32 {
	if m != nil {
		return m.PageCount
	}
	return 0
}

func (m *PhotoSearchResponse) GetTotal() string {
	if m != nil {
		return m.Total
	}
	return ""
}

func (m *PhotoSearchResponse) GetPhotos() []*PhotoSearchResponse_Photo {
	if m != nil {
		return m.Photos
	}
	return nil
}

type PhotoSearchResponse_Photo struct {
	ID           string `protobuf:"bytes,1,opt,name=ID" json:"ID,omitempty"`
	ThumbnailUrl string `protobuf:"bytes,2,opt,name=thumbnailUrl" json:"thumbnailUrl,omitempty"`
	MediumUrl    string `protobuf:"bytes,3,opt,name=mediumUrl" json:"mediumUrl,omitempty"`
	LargeUrl     string `protobuf:"bytes,4,opt,name=largeUrl" json:"largeUrl,omitempty"`
	Owner        string `protobuf:"bytes,5,opt,name=owner" json:"owner,omitempty"`
	Title        string `protobuf:"bytes,6,opt,name=title" json:"title,omitempty"`
}

func (m *PhotoSearchResponse_Photo) Reset()                    { *m = PhotoSearchResponse_Photo{} }
func (m *PhotoSearchResponse_Photo) String() string            { return proto.CompactTextString(m) }
func (*PhotoSearchResponse_Photo) ProtoMessage()               {}
func (*PhotoSearchResponse_Photo) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1, 0} }

func (m *PhotoSearchResponse_Photo) GetID() string {
	if m != nil {
		return m.ID
	}
	return ""
}

func (m *PhotoSearchResponse_Photo) GetThumbnailUrl() string {
	if m != nil {
		return m.ThumbnailUrl
	}
	return ""
}

func (m *PhotoSearchResponse_Photo) GetMediumUrl() string {
	if m != nil {
		return m.MediumUrl
	}
	return ""
}

func (m *PhotoSearchResponse_Photo) GetLargeUrl() string {
	if m != nil {
		return m.LargeUrl
	}
	return ""
}

func (m *PhotoSearchResponse_Photo) GetOwner() string {
	if m != nil {
		return m.Owner
	}
	return ""
}

func (m *PhotoSearchResponse_Photo) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func init() {
	proto.RegisterType((*PhotoSearchRequest)(nil), "flickrphotosearch.PhotoSearchRequest")
	proto.RegisterType((*PhotoSearchResponse)(nil), "flickrphotosearch.PhotoSearchResponse")
	proto.RegisterType((*PhotoSearchResponse_Photo)(nil), "flickrphotosearch.PhotoSearchResponse.Photo")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Photo service

type PhotoClient interface {
	// Search photos
	Search(ctx context.Context, in *PhotoSearchRequest, opts ...grpc.CallOption) (*PhotoSearchResponse, error)
}

type photoClient struct {
	cc *grpc.ClientConn
}

func NewPhotoClient(cc *grpc.ClientConn) PhotoClient {
	return &photoClient{cc}
}

func (c *photoClient) Search(ctx context.Context, in *PhotoSearchRequest, opts ...grpc.CallOption) (*PhotoSearchResponse, error) {
	out := new(PhotoSearchResponse)
	err := grpc.Invoke(ctx, "/flickrphotosearch.Photo/Search", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Photo service

type PhotoServer interface {
	// Search photos
	Search(context.Context, *PhotoSearchRequest) (*PhotoSearchResponse, error)
}

func RegisterPhotoServer(s *grpc.Server, srv PhotoServer) {
	s.RegisterService(&_Photo_serviceDesc, srv)
}

func _Photo_Search_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PhotoSearchRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PhotoServer).Search(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/flickrphotosearch.Photo/Search",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PhotoServer).Search(ctx, req.(*PhotoSearchRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Photo_serviceDesc = grpc.ServiceDesc{
	ServiceName: "flickrphotosearch.Photo",
	HandlerType: (*PhotoServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Search",
			Handler:    _Photo_Search_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "photoSearch.proto",
}

func init() { proto.RegisterFile("photoSearch.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 296 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x51, 0x31, 0x4f, 0xb4, 0x40,
	0x10, 0xfd, 0xe0, 0x0e, 0xf2, 0x31, 0x1a, 0x93, 0x1b, 0x2d, 0x08, 0xb1, 0x20, 0x24, 0x1a, 0x0a,
	0x43, 0x71, 0xf6, 0x16, 0x4a, 0x73, 0x9d, 0x59, 0x63, 0xa1, 0xdd, 0x72, 0xae, 0x07, 0x11, 0x58,
	0x6e, 0x59, 0x62, 0xfc, 0x39, 0xfe, 0x06, 0xff, 0xa0, 0xd9, 0x59, 0x4e, 0x72, 0xd1, 0x44, 0x3b,
	0xde, 0x7b, 0xc3, 0x9b, 0xf7, 0x76, 0x60, 0xd1, 0x95, 0x52, 0xcb, 0x3b, 0xc1, 0xd5, 0xba, 0xcc,
	0x3a, 0x25, 0xb5, 0xc4, 0xc5, 0x73, 0x5d, 0xad, 0x5f, 0x14, 0x09, 0x3d, 0x09, 0xc9, 0x15, 0xe0,
	0xed, 0x34, 0xc7, 0xc4, 0x76, 0x10, 0xbd, 0xc6, 0x13, 0xf0, 0xb6, 0x83, 0x50, 0x6f, 0xa1, 0x13,
	0x3b, 0x69, 0xc0, 0x2c, 0x40, 0x84, 0x79, 0xc7, 0x37, 0x22, 0x74, 0x63, 0x27, 0xf5, 0x18, 0x7d,
	0x27, 0x1f, 0x2e, 0x1c, 0xef, 0x19, 0xf4, 0x9d, 0x6c, 0x7b, 0x81, 0xa7, 0x10, 0x18, 0xfd, 0x46,
	0x0e, 0xad, 0x26, 0x17, 0x8f, 0x4d, 0x84, 0xf1, 0xd7, 0x52, 0xf3, 0x9a, 0xac, 0x02, 0x66, 0x01,
	0xe6, 0xe0, 0xdb, 0x68, 0xe1, 0x2c, 0x9e, 0xa5, 0x07, 0xcb, 0x8b, 0xec, 0x5b, 0xde, 0xec, 0x87,
	0x5d, 0x96, 0x63, 0xe3, 0xbf, 0xd1, 0xbb, 0x03, 0x1e, 0x31, 0x78, 0x04, 0xee, 0x2a, 0x1f, 0x2b,
	0xb8, 0xab, 0x1c, 0x13, 0x38, 0xd4, 0xe5, 0xd0, 0x14, 0x2d, 0xaf, 0xea, 0x7b, 0xb5, 0x5b, 0xbe,
	0xc7, 0x99, 0xdc, 0x8d, 0x78, 0xaa, 0x86, 0xc6, 0x0c, 0xcc, 0x68, 0x60, 0x22, 0x30, 0x82, 0xff,
	0x35, 0x57, 0x1b, 0x61, 0xc4, 0x39, 0x89, 0x5f, 0xd8, 0x74, 0x92, 0xaf, 0xad, 0x50, 0xa1, 0x67,
	0x3b, 0x11, 0xa0, 0xa6, 0x95, 0xae, 0x45, 0xe8, 0x8f, 0x4d, 0x0d, 0x58, 0x16, 0xbb, 0x88, 0x0f,
	0xe0, 0xdb, 0x32, 0x78, 0xf6, 0x5b, 0x59, 0xba, 0x4c, 0x74, 0xfe, 0xb7, 0x37, 0x49, 0xfe, 0x5d,
	0xfb, 0x8f, 0x73, 0xde, 0x55, 0x7d, 0xe1, 0xd3, 0xed, 0x2f, 0x3f, 0x03, 0x00, 0x00, 0xff, 0xff,
	0xc7, 0x5d, 0xf2, 0x32, 0x10, 0x02, 0x00, 0x00,
}
