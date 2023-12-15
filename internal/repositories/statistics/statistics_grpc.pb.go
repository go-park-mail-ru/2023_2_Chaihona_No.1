// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.12.4
// source: statistics.proto

package statistics

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
	Statistics_GetQuestionsCtx_FullMethodName   = "/statistics.Statistics/GetQuestionsCtx"
	Statistics_AnswerQuestionCtx_FullMethodName = "/statistics.Statistics/AnswerQuestionCtx"
	Statistics_GetStatisticCtx_FullMethodName   = "/statistics.Statistics/GetStatisticCtx"
)

// StatisticsClient is the client API for Statistics service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type StatisticsClient interface {
	GetQuestionsCtx(ctx context.Context, in *UserId, opts ...grpc.CallOption) (*QuestionsMap, error)
	AnswerQuestionCtx(ctx context.Context, in *Answer, opts ...grpc.CallOption) (*Nothing, error)
	GetStatisticCtx(ctx context.Context, in *Nothing, opts ...grpc.CallOption) (*QuestionsMap, error)
}

type statisticsClient struct {
	cc grpc.ClientConnInterface
}

func NewStatisticsClient(cc grpc.ClientConnInterface) StatisticsClient {
	return &statisticsClient{cc}
}

func (c *statisticsClient) GetQuestionsCtx(ctx context.Context, in *UserId, opts ...grpc.CallOption) (*QuestionsMap, error) {
	out := new(QuestionsMap)
	err := c.cc.Invoke(ctx, Statistics_GetQuestionsCtx_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *statisticsClient) AnswerQuestionCtx(ctx context.Context, in *Answer, opts ...grpc.CallOption) (*Nothing, error) {
	out := new(Nothing)
	err := c.cc.Invoke(ctx, Statistics_AnswerQuestionCtx_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *statisticsClient) GetStatisticCtx(ctx context.Context, in *Nothing, opts ...grpc.CallOption) (*QuestionsMap, error) {
	out := new(QuestionsMap)
	err := c.cc.Invoke(ctx, Statistics_GetStatisticCtx_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// StatisticsServer is the server API for Statistics service.
// All implementations must embed UnimplementedStatisticsServer
// for forward compatibility
type StatisticsServer interface {
	GetQuestionsCtx(context.Context, *UserId) (*QuestionsMap, error)
	AnswerQuestionCtx(context.Context, *Answer) (*Nothing, error)
	GetStatisticCtx(context.Context, *Nothing) (*QuestionsMap, error)
	mustEmbedUnimplementedStatisticsServer()
}

// UnimplementedStatisticsServer must be embedded to have forward compatible implementations.
type UnimplementedStatisticsServer struct {
}

func (UnimplementedStatisticsServer) GetQuestionsCtx(context.Context, *UserId) (*QuestionsMap, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetQuestionsCtx not implemented")
}
func (UnimplementedStatisticsServer) AnswerQuestionCtx(context.Context, *Answer) (*Nothing, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AnswerQuestionCtx not implemented")
}
func (UnimplementedStatisticsServer) GetStatisticCtx(context.Context, *Nothing) (*QuestionsMap, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetStatisticCtx not implemented")
}
func (UnimplementedStatisticsServer) mustEmbedUnimplementedStatisticsServer() {}

// UnsafeStatisticsServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to StatisticsServer will
// result in compilation errors.
type UnsafeStatisticsServer interface {
	mustEmbedUnimplementedStatisticsServer()
}

func RegisterStatisticsServer(s grpc.ServiceRegistrar, srv StatisticsServer) {
	s.RegisterService(&Statistics_ServiceDesc, srv)
}

func _Statistics_GetQuestionsCtx_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StatisticsServer).GetQuestionsCtx(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Statistics_GetQuestionsCtx_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StatisticsServer).GetQuestionsCtx(ctx, req.(*UserId))
	}
	return interceptor(ctx, in, info, handler)
}

func _Statistics_AnswerQuestionCtx_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Answer)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StatisticsServer).AnswerQuestionCtx(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Statistics_AnswerQuestionCtx_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StatisticsServer).AnswerQuestionCtx(ctx, req.(*Answer))
	}
	return interceptor(ctx, in, info, handler)
}

func _Statistics_GetStatisticCtx_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Nothing)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StatisticsServer).GetStatisticCtx(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Statistics_GetStatisticCtx_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StatisticsServer).GetStatisticCtx(ctx, req.(*Nothing))
	}
	return interceptor(ctx, in, info, handler)
}

// Statistics_ServiceDesc is the grpc.ServiceDesc for Statistics service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Statistics_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "statistics.Statistics",
	HandlerType: (*StatisticsServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetQuestionsCtx",
			Handler:    _Statistics_GetQuestionsCtx_Handler,
		},
		{
			MethodName: "AnswerQuestionCtx",
			Handler:    _Statistics_AnswerQuestionCtx_Handler,
		},
		{
			MethodName: "GetStatisticCtx",
			Handler:    _Statistics_GetStatisticCtx_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "statistics.proto",
}