// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.20.1
// source: proto/video.proto

package pb

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

// VideoClient is the client API for Video service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type VideoClient interface {
	//保存视频信息到数据库
	SaveVideo(ctx context.Context, in *SaveVideoReq, opts ...grpc.CallOption) (*SaveVideoResp, error)
	//查看某用户id的视频列表
	PublishList(ctx context.Context, in *PublishListReq, opts ...grpc.CallOption) (*PublishListResp, error)
	//视频流
	Feed(ctx context.Context, in *FeedReq, opts ...grpc.CallOption) (*FeedResp, error)
}

type videoClient struct {
	cc grpc.ClientConnInterface
}

func NewVideoClient(cc grpc.ClientConnInterface) VideoClient {
	return &videoClient{cc}
}

func (c *videoClient) SaveVideo(ctx context.Context, in *SaveVideoReq, opts ...grpc.CallOption) (*SaveVideoResp, error) {
	out := new(SaveVideoResp)
	err := c.cc.Invoke(ctx, "/pb.video/SaveVideo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *videoClient) PublishList(ctx context.Context, in *PublishListReq, opts ...grpc.CallOption) (*PublishListResp, error) {
	out := new(PublishListResp)
	err := c.cc.Invoke(ctx, "/pb.video/PublishList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *videoClient) Feed(ctx context.Context, in *FeedReq, opts ...grpc.CallOption) (*FeedResp, error) {
	out := new(FeedResp)
	err := c.cc.Invoke(ctx, "/pb.video/Feed", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// VideoServer is the server API for Video service.
// All implementations must embed UnimplementedVideoServer
// for forward compatibility
type VideoServer interface {
	//保存视频信息到数据库
	SaveVideo(context.Context, *SaveVideoReq) (*SaveVideoResp, error)
	//查看某用户id的视频列表
	PublishList(context.Context, *PublishListReq) (*PublishListResp, error)
	//视频流
	Feed(context.Context, *FeedReq) (*FeedResp, error)
	mustEmbedUnimplementedVideoServer()
}

// UnimplementedVideoServer must be embedded to have forward compatible implementations.
type UnimplementedVideoServer struct {
}

func (UnimplementedVideoServer) SaveVideo(context.Context, *SaveVideoReq) (*SaveVideoResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SaveVideo not implemented")
}
func (UnimplementedVideoServer) PublishList(context.Context, *PublishListReq) (*PublishListResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PublishList not implemented")
}
func (UnimplementedVideoServer) Feed(context.Context, *FeedReq) (*FeedResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Feed not implemented")
}
func (UnimplementedVideoServer) mustEmbedUnimplementedVideoServer() {}

// UnsafeVideoServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to VideoServer will
// result in compilation errors.
type UnsafeVideoServer interface {
	mustEmbedUnimplementedVideoServer()
}

func RegisterVideoServer(s grpc.ServiceRegistrar, srv VideoServer) {
	s.RegisterService(&Video_ServiceDesc, srv)
}

func _Video_SaveVideo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SaveVideoReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VideoServer).SaveVideo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.video/SaveVideo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VideoServer).SaveVideo(ctx, req.(*SaveVideoReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Video_PublishList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PublishListReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VideoServer).PublishList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.video/PublishList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VideoServer).PublishList(ctx, req.(*PublishListReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Video_Feed_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FeedReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VideoServer).Feed(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.video/Feed",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VideoServer).Feed(ctx, req.(*FeedReq))
	}
	return interceptor(ctx, in, info, handler)
}

// Video_ServiceDesc is the grpc.ServiceDesc for Video service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Video_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "pb.video",
	HandlerType: (*VideoServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SaveVideo",
			Handler:    _Video_SaveVideo_Handler,
		},
		{
			MethodName: "PublishList",
			Handler:    _Video_PublishList_Handler,
		},
		{
			MethodName: "Feed",
			Handler:    _Video_Feed_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/video.proto",
}
