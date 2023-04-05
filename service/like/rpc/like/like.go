// Code generated by goctl. DO NOT EDIT.
// Source: like.proto

package like

import (
	"context"

	"tik-tok-brief/service/like/rpc/proto/pb"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	IsLikeReq    = pb.IsLikeReq
	IsLikeResp   = pb.IsLikeResp
	Like         = pb.Like
	LikeListReq  = pb.LikeListReq
	LikeListResp = pb.LikeListResp
	LikeReq      = pb.LikeReq
	LikeResp     = pb.LikeResp

	LikeZrpcClient interface {
		// 赞操作
		Like(ctx context.Context, in *LikeReq, opts ...grpc.CallOption) (*LikeResp, error)
		// 获取点赞视频列表
		LikeList(ctx context.Context, in *LikeListReq, opts ...grpc.CallOption) (*LikeListResp, error)
		// 获取UID是否对VID点赞
		IsLike(ctx context.Context, in *IsLikeReq, opts ...grpc.CallOption) (*IsLikeResp, error)
	}

	defaultLikeZrpcClient struct {
		cli zrpc.Client
	}
)

func NewLikeZrpcClient(cli zrpc.Client) LikeZrpcClient {
	return &defaultLikeZrpcClient{
		cli: cli,
	}
}

// 赞操作
func (m *defaultLikeZrpcClient) Like(ctx context.Context, in *LikeReq, opts ...grpc.CallOption) (*LikeResp, error) {
	client := pb.NewLikeClient(m.cli.Conn())
	return client.Like(ctx, in, opts...)
}

// 获取点赞视频列表
func (m *defaultLikeZrpcClient) LikeList(ctx context.Context, in *LikeListReq, opts ...grpc.CallOption) (*LikeListResp, error) {
	client := pb.NewLikeClient(m.cli.Conn())
	return client.LikeList(ctx, in, opts...)
}

// 获取UID是否对VID点赞
func (m *defaultLikeZrpcClient) IsLike(ctx context.Context, in *IsLikeReq, opts ...grpc.CallOption) (*IsLikeResp, error) {
	client := pb.NewLikeClient(m.cli.Conn())
	return client.IsLike(ctx, in, opts...)
}
