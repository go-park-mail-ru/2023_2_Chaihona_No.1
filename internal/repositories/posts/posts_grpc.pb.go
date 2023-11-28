// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.12.4
// source: posts.proto

package posts

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	PostsService_CreateNewPostCtx_FullMethodName                 = "/posts.PostsService/CreateNewPostCtx"
	PostsService_DeletePostCtx_FullMethodName                    = "/posts.PostsService/DeletePostCtx"
	PostsService_ChangePostCtx_FullMethodName                    = "/posts.PostsService/ChangePostCtx"
	PostsService_GetPostByIdCtx_FullMethodName                   = "/posts.PostsService/GetPostByIdCtx"
	PostsService_GetPostsByAuthorIdForStrangerCtx_FullMethodName = "/posts.PostsService/GetPostsByAuthorIdForStrangerCtx"
	PostsService_GetOwnPostsByAuthorIdCtx_FullMethodName         = "/posts.PostsService/GetOwnPostsByAuthorIdCtx"
	PostsService_GetPostsByAuthorIdForFollowerCtx_FullMethodName = "/posts.PostsService/GetPostsByAuthorIdForFollowerCtx"
	PostsService_GetUsersFeedCtx_FullMethodName                  = "/posts.PostsService/GetUsersFeedCtx"
)

// PostsServiceClient is the client API for PostsService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type PostsServiceClient interface {
	CreateNewPostCtx(ctx context.Context, in *PostGRPC, opts ...grpc.CallOption) (*Int, error)
	DeletePostCtx(ctx context.Context, in *UInt, opts ...grpc.CallOption) (*Nothing, error)
	ChangePostCtx(ctx context.Context, in *PostGRPC, opts ...grpc.CallOption) (*Nothing, error)
	GetPostByIdCtx(ctx context.Context, in *UInt, opts ...grpc.CallOption) (*PostGRPC, error)
	GetPostsByAuthorIdForStrangerCtx(ctx context.Context, in *AuthorSubscriberId, opts ...grpc.CallOption) (*PostsMapGRPC, error)
	GetOwnPostsByAuthorIdCtx(ctx context.Context, in *AuthorSubscriberId, opts ...grpc.CallOption) (*PostsMapGRPC, error)
	GetPostsByAuthorIdForFollowerCtx(ctx context.Context, in *AuthorSubscriberId, opts ...grpc.CallOption) (*PostsMapGRPC, error)
	GetUsersFeedCtx(ctx context.Context, in *UInt, opts ...grpc.CallOption) (*PostsMapGRPC, error)
}

type postsServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewPostsServiceClient(cc grpc.ClientConnInterface) PostsServiceClient {
	return &postsServiceClient{cc}
}

func (c *postsServiceClient) CreateNewPostCtx(ctx context.Context, in *PostGRPC, opts ...grpc.CallOption) (*Int, error) {
	out := new(Int)
	err := c.cc.Invoke(ctx, PostsService_CreateNewPostCtx_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *postsServiceClient) DeletePostCtx(ctx context.Context, in *UInt, opts ...grpc.CallOption) (*Nothing, error) {
	out := new(Nothing)
	err := c.cc.Invoke(ctx, PostsService_DeletePostCtx_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *postsServiceClient) ChangePostCtx(ctx context.Context, in *PostGRPC, opts ...grpc.CallOption) (*Nothing, error) {
	out := new(Nothing)
	err := c.cc.Invoke(ctx, PostsService_ChangePostCtx_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *postsServiceClient) GetPostByIdCtx(ctx context.Context, in *UInt, opts ...grpc.CallOption) (*PostGRPC, error) {
	out := new(PostGRPC)
	err := c.cc.Invoke(ctx, PostsService_GetPostByIdCtx_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *postsServiceClient) GetPostsByAuthorIdForStrangerCtx(ctx context.Context, in *AuthorSubscriberId, opts ...grpc.CallOption) (*PostsMapGRPC, error) {
	out := new(PostsMapGRPC)
	err := c.cc.Invoke(ctx, PostsService_GetPostsByAuthorIdForStrangerCtx_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *postsServiceClient) GetOwnPostsByAuthorIdCtx(ctx context.Context, in *AuthorSubscriberId, opts ...grpc.CallOption) (*PostsMapGRPC, error) {
	out := new(PostsMapGRPC)
	err := c.cc.Invoke(ctx, PostsService_GetOwnPostsByAuthorIdCtx_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *postsServiceClient) GetPostsByAuthorIdForFollowerCtx(ctx context.Context, in *AuthorSubscriberId, opts ...grpc.CallOption) (*PostsMapGRPC, error) {
	out := new(PostsMapGRPC)
	err := c.cc.Invoke(ctx, PostsService_GetPostsByAuthorIdForFollowerCtx_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *postsServiceClient) GetUsersFeedCtx(ctx context.Context, in *UInt, opts ...grpc.CallOption) (*PostsMapGRPC, error) {
	out := new(PostsMapGRPC)
	err := c.cc.Invoke(ctx, PostsService_GetUsersFeedCtx_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PostsServiceServer is the server API for PostsService service.
// All implementations must embed UnimplementedPostsServiceServer
// for forward compatibility
type PostsServiceServer interface {
	CreateNewPostCtx(context.Context, *PostGRPC) (*Int, error)
	DeletePostCtx(context.Context, *UInt) (*Nothing, error)
	ChangePostCtx(context.Context, *PostGRPC) (*Nothing, error)
	GetPostByIdCtx(context.Context, *UInt) (*PostGRPC, error)
	GetPostsByAuthorIdForStrangerCtx(context.Context, *AuthorSubscriberId) (*PostsMapGRPC, error)
	GetOwnPostsByAuthorIdCtx(context.Context, *AuthorSubscriberId) (*PostsMapGRPC, error)
	GetPostsByAuthorIdForFollowerCtx(context.Context, *AuthorSubscriberId) (*PostsMapGRPC, error)
	GetUsersFeedCtx(context.Context, *UInt) (*PostsMapGRPC, error)
	mustEmbedUnimplementedPostsServiceServer()
}

// UnimplementedPostsServiceServer must be embedded to have forward compatible implementations.
type UnimplementedPostsServiceServer struct {
}

func (UnimplementedPostsServiceServer) CreateNewPostCtx(context.Context, *PostGRPC) (*Int, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateNewPostCtx not implemented")
}
func (UnimplementedPostsServiceServer) DeletePostCtx(context.Context, *UInt) (*Nothing, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeletePostCtx not implemented")
}
func (UnimplementedPostsServiceServer) ChangePostCtx(context.Context, *PostGRPC) (*Nothing, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ChangePostCtx not implemented")
}
func (UnimplementedPostsServiceServer) GetPostByIdCtx(context.Context, *UInt) (*PostGRPC, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPostByIdCtx not implemented")
}
func (UnimplementedPostsServiceServer) GetPostsByAuthorIdForStrangerCtx(context.Context, *AuthorSubscriberId) (*PostsMapGRPC, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPostsByAuthorIdForStrangerCtx not implemented")
}
func (UnimplementedPostsServiceServer) GetOwnPostsByAuthorIdCtx(context.Context, *AuthorSubscriberId) (*PostsMapGRPC, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetOwnPostsByAuthorIdCtx not implemented")
}
func (UnimplementedPostsServiceServer) GetPostsByAuthorIdForFollowerCtx(context.Context, *AuthorSubscriberId) (*PostsMapGRPC, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPostsByAuthorIdForFollowerCtx not implemented")
}
func (UnimplementedPostsServiceServer) GetUsersFeedCtx(context.Context, *UInt) (*PostsMapGRPC, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUsersFeedCtx not implemented")
}
func (UnimplementedPostsServiceServer) mustEmbedUnimplementedPostsServiceServer() {}

// UnsafePostsServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to PostsServiceServer will
// result in compilation errors.
type UnsafePostsServiceServer interface {
	mustEmbedUnimplementedPostsServiceServer()
}

func RegisterPostsServiceServer(s grpc.ServiceRegistrar, srv PostsServiceServer) {
	s.RegisterService(&PostsService_ServiceDesc, srv)
}

func _PostsService_CreateNewPostCtx_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PostGRPC)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PostsServiceServer).CreateNewPostCtx(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PostsService_CreateNewPostCtx_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PostsServiceServer).CreateNewPostCtx(ctx, req.(*PostGRPC))
	}
	return interceptor(ctx, in, info, handler)
}

func _PostsService_DeletePostCtx_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UInt)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PostsServiceServer).DeletePostCtx(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PostsService_DeletePostCtx_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PostsServiceServer).DeletePostCtx(ctx, req.(*UInt))
	}
	return interceptor(ctx, in, info, handler)
}

func _PostsService_ChangePostCtx_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PostGRPC)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PostsServiceServer).ChangePostCtx(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PostsService_ChangePostCtx_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PostsServiceServer).ChangePostCtx(ctx, req.(*PostGRPC))
	}
	return interceptor(ctx, in, info, handler)
}

func _PostsService_GetPostByIdCtx_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UInt)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PostsServiceServer).GetPostByIdCtx(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PostsService_GetPostByIdCtx_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PostsServiceServer).GetPostByIdCtx(ctx, req.(*UInt))
	}
	return interceptor(ctx, in, info, handler)
}

func _PostsService_GetPostsByAuthorIdForStrangerCtx_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AuthorSubscriberId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PostsServiceServer).GetPostsByAuthorIdForStrangerCtx(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PostsService_GetPostsByAuthorIdForStrangerCtx_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PostsServiceServer).GetPostsByAuthorIdForStrangerCtx(ctx, req.(*AuthorSubscriberId))
	}
	return interceptor(ctx, in, info, handler)
}

func _PostsService_GetOwnPostsByAuthorIdCtx_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AuthorSubscriberId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PostsServiceServer).GetOwnPostsByAuthorIdCtx(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PostsService_GetOwnPostsByAuthorIdCtx_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PostsServiceServer).GetOwnPostsByAuthorIdCtx(ctx, req.(*AuthorSubscriberId))
	}
	return interceptor(ctx, in, info, handler)
}

func _PostsService_GetPostsByAuthorIdForFollowerCtx_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AuthorSubscriberId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PostsServiceServer).GetPostsByAuthorIdForFollowerCtx(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PostsService_GetPostsByAuthorIdForFollowerCtx_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PostsServiceServer).GetPostsByAuthorIdForFollowerCtx(ctx, req.(*AuthorSubscriberId))
	}
	return interceptor(ctx, in, info, handler)
}

func _PostsService_GetUsersFeedCtx_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UInt)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PostsServiceServer).GetUsersFeedCtx(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PostsService_GetUsersFeedCtx_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PostsServiceServer).GetUsersFeedCtx(ctx, req.(*UInt))
	}
	return interceptor(ctx, in, info, handler)
}

// PostsService_ServiceDesc is the grpc.ServiceDesc for PostsService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var PostsService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "posts.PostsService",
	HandlerType: (*PostsServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateNewPostCtx",
			Handler:    _PostsService_CreateNewPostCtx_Handler,
		},
		{
			MethodName: "DeletePostCtx",
			Handler:    _PostsService_DeletePostCtx_Handler,
		},
		{
			MethodName: "ChangePostCtx",
			Handler:    _PostsService_ChangePostCtx_Handler,
		},
		{
			MethodName: "GetPostByIdCtx",
			Handler:    _PostsService_GetPostByIdCtx_Handler,
		},
		{
			MethodName: "GetPostsByAuthorIdForStrangerCtx",
			Handler:    _PostsService_GetPostsByAuthorIdForStrangerCtx_Handler,
		},
		{
			MethodName: "GetOwnPostsByAuthorIdCtx",
			Handler:    _PostsService_GetOwnPostsByAuthorIdCtx_Handler,
		},
		{
			MethodName: "GetPostsByAuthorIdForFollowerCtx",
			Handler:    _PostsService_GetPostsByAuthorIdForFollowerCtx_Handler,
		},
		{
			MethodName: "GetUsersFeedCtx",
			Handler:    _PostsService_GetUsersFeedCtx_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "posts.proto",
}
