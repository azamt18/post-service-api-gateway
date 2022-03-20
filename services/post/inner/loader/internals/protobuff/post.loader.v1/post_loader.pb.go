// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.19.4
// source: post_loader.proto

package post_loader_v1

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type LoadPostsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *LoadPostsRequest) Reset() {
	*x = LoadPostsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_post_loader_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LoadPostsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LoadPostsRequest) ProtoMessage() {}

func (x *LoadPostsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_post_loader_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LoadPostsRequest.ProtoReflect.Descriptor instead.
func (*LoadPostsRequest) Descriptor() ([]byte, []int) {
	return file_post_loader_proto_rawDescGZIP(), []int{0}
}

type LoadPostsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success          bool  `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
	LoadedPostsCount int64 `protobuf:"varint,2,opt,name=loaded_posts_count,json=loadedPostsCount,proto3" json:"loaded_posts_count,omitempty"`
}

func (x *LoadPostsResponse) Reset() {
	*x = LoadPostsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_post_loader_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LoadPostsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LoadPostsResponse) ProtoMessage() {}

func (x *LoadPostsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_post_loader_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LoadPostsResponse.ProtoReflect.Descriptor instead.
func (*LoadPostsResponse) Descriptor() ([]byte, []int) {
	return file_post_loader_proto_rawDescGZIP(), []int{1}
}

func (x *LoadPostsResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

func (x *LoadPostsResponse) GetLoadedPostsCount() int64 {
	if x != nil {
		return x.LoadedPostsCount
	}
	return 0
}

var File_post_loader_proto protoreflect.FileDescriptor

var file_post_loader_proto_rawDesc = []byte{
	0x0a, 0x11, 0x70, 0x6f, 0x73, 0x74, 0x5f, 0x6c, 0x6f, 0x61, 0x64, 0x65, 0x72, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x0e, 0x70, 0x6f, 0x73, 0x74, 0x2e, 0x6c, 0x6f, 0x61, 0x64, 0x65, 0x72,
	0x2e, 0x76, 0x31, 0x22, 0x12, 0x0a, 0x10, 0x4c, 0x6f, 0x61, 0x64, 0x50, 0x6f, 0x73, 0x74, 0x73,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x5b, 0x0a, 0x11, 0x4c, 0x6f, 0x61, 0x64, 0x50,
	0x6f, 0x73, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07,
	0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x73,
	0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x12, 0x2c, 0x0a, 0x12, 0x6c, 0x6f, 0x61, 0x64, 0x65, 0x64,
	0x5f, 0x70, 0x6f, 0x73, 0x74, 0x73, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x10, 0x6c, 0x6f, 0x61, 0x64, 0x65, 0x64, 0x50, 0x6f, 0x73, 0x74, 0x73, 0x43,
	0x6f, 0x75, 0x6e, 0x74, 0x32, 0x66, 0x0a, 0x10, 0x4c, 0x6f, 0x61, 0x64, 0x50, 0x6f, 0x73, 0x74,
	0x73, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x52, 0x0a, 0x09, 0x4c, 0x6f, 0x61, 0x64,
	0x50, 0x6f, 0x73, 0x74, 0x73, 0x12, 0x20, 0x2e, 0x70, 0x6f, 0x73, 0x74, 0x2e, 0x6c, 0x6f, 0x61,
	0x64, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x6f, 0x61, 0x64, 0x50, 0x6f, 0x73, 0x74, 0x73,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x21, 0x2e, 0x70, 0x6f, 0x73, 0x74, 0x2e, 0x6c,
	0x6f, 0x61, 0x64, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x6f, 0x61, 0x64, 0x50, 0x6f, 0x73,
	0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x10, 0x5a, 0x0e,
	0x70, 0x6f, 0x73, 0x74, 0x2e, 0x6c, 0x6f, 0x61, 0x64, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_post_loader_proto_rawDescOnce sync.Once
	file_post_loader_proto_rawDescData = file_post_loader_proto_rawDesc
)

func file_post_loader_proto_rawDescGZIP() []byte {
	file_post_loader_proto_rawDescOnce.Do(func() {
		file_post_loader_proto_rawDescData = protoimpl.X.CompressGZIP(file_post_loader_proto_rawDescData)
	})
	return file_post_loader_proto_rawDescData
}

var file_post_loader_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_post_loader_proto_goTypes = []interface{}{
	(*LoadPostsRequest)(nil),  // 0: post.loader.v1.LoadPostsRequest
	(*LoadPostsResponse)(nil), // 1: post.loader.v1.LoadPostsResponse
}
var file_post_loader_proto_depIdxs = []int32{
	0, // 0: post.loader.v1.LoadPostsService.LoadPosts:input_type -> post.loader.v1.LoadPostsRequest
	1, // 1: post.loader.v1.LoadPostsService.LoadPosts:output_type -> post.loader.v1.LoadPostsResponse
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_post_loader_proto_init() }
func file_post_loader_proto_init() {
	if File_post_loader_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_post_loader_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LoadPostsRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_post_loader_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LoadPostsResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_post_loader_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_post_loader_proto_goTypes,
		DependencyIndexes: file_post_loader_proto_depIdxs,
		MessageInfos:      file_post_loader_proto_msgTypes,
	}.Build()
	File_post_loader_proto = out.File
	file_post_loader_proto_rawDesc = nil
	file_post_loader_proto_goTypes = nil
	file_post_loader_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// LoadPostsServiceClient is the client API for LoadPostsService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type LoadPostsServiceClient interface {
	LoadPosts(ctx context.Context, in *LoadPostsRequest, opts ...grpc.CallOption) (*LoadPostsResponse, error)
}

type loadPostsServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewLoadPostsServiceClient(cc grpc.ClientConnInterface) LoadPostsServiceClient {
	return &loadPostsServiceClient{cc}
}

func (c *loadPostsServiceClient) LoadPosts(ctx context.Context, in *LoadPostsRequest, opts ...grpc.CallOption) (*LoadPostsResponse, error) {
	out := new(LoadPostsResponse)
	err := c.cc.Invoke(ctx, "/post.loader.v1.LoadPostsService/LoadPosts", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// LoadPostsServiceServer is the server API for LoadPostsService service.
type LoadPostsServiceServer interface {
	LoadPosts(context.Context, *LoadPostsRequest) (*LoadPostsResponse, error)
}

// UnimplementedLoadPostsServiceServer can be embedded to have forward compatible implementations.
type UnimplementedLoadPostsServiceServer struct {
}

func (*UnimplementedLoadPostsServiceServer) LoadPosts(context.Context, *LoadPostsRequest) (*LoadPostsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method LoadPosts not implemented")
}

func RegisterLoadPostsServiceServer(s *grpc.Server, srv LoadPostsServiceServer) {
	s.RegisterService(&_LoadPostsService_serviceDesc, srv)
}

func _LoadPostsService_LoadPosts_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LoadPostsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LoadPostsServiceServer).LoadPosts(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/post.loader.v1.LoadPostsService/LoadPosts",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LoadPostsServiceServer).LoadPosts(ctx, req.(*LoadPostsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _LoadPostsService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "post.loader.v1.LoadPostsService",
	HandlerType: (*LoadPostsServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "LoadPosts",
			Handler:    _LoadPostsService_LoadPosts_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "post_loader.proto",
}
