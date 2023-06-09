// Code generated by goctl. DO NOT EDIT.
// Source: video.proto

package video

import (
	"context"

	"tik-tok-brief/service/video/rpc/proto/pb"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	FeedReq                 = pb.FeedReq
	FeedResp                = pb.FeedResp
	GetVideoListByIdsReq    = pb.GetVideoListByIdsReq
	GetVideoListByIdsResp   = pb.GetVideoListByIdsResp
	PublishListReq          = pb.PublishListReq
	PublishListResp         = pb.PublishListResp
	SaveVideoReq            = pb.SaveVideoReq
	SaveVideoResp           = pb.SaveVideoResp
	UpdateCommentCountReq   = pb.UpdateCommentCountReq
	UpdateCommentCountResp  = pb.UpdateCommentCountResp
	UpdateFavoriteCountReq  = pb.UpdateFavoriteCountReq
	UpdateFavoriteCountResp = pb.UpdateFavoriteCountResp
	Video                   = pb.Video

	VideoZrpcClient interface {
		// 保存视频信息到数据库
		SaveVideo(ctx context.Context, in *SaveVideoReq, opts ...grpc.CallOption) (*SaveVideoResp, error)
		// 查看某用户id的视频列表
		PublishList(ctx context.Context, in *PublishListReq, opts ...grpc.CallOption) (*PublishListResp, error)
		// 视频流
		Feed(ctx context.Context, in *FeedReq, opts ...grpc.CallOption) (*FeedResp, error)
		// 更新视频点赞数
		UpdateFavoriteCount(ctx context.Context, in *UpdateFavoriteCountReq, opts ...grpc.CallOption) (*UpdateFavoriteCountResp, error)
		// 更新视频评论数
		UpdateCommentCount(ctx context.Context, in *UpdateCommentCountReq, opts ...grpc.CallOption) (*UpdateFavoriteCountResp, error)
		// 根据videoIds获取视频列表
		GetVideoListByIds(ctx context.Context, in *GetVideoListByIdsReq, opts ...grpc.CallOption) (*GetVideoListByIdsResp, error)
	}

	defaultVideoZrpcClient struct {
		cli zrpc.Client
	}
)

func NewVideoZrpcClient(cli zrpc.Client) VideoZrpcClient {
	return &defaultVideoZrpcClient{
		cli: cli,
	}
}

// 保存视频信息到数据库
func (m *defaultVideoZrpcClient) SaveVideo(ctx context.Context, in *SaveVideoReq, opts ...grpc.CallOption) (*SaveVideoResp, error) {
	client := pb.NewVideoClient(m.cli.Conn())
	return client.SaveVideo(ctx, in, opts...)
}

// 查看某用户id的视频列表
func (m *defaultVideoZrpcClient) PublishList(ctx context.Context, in *PublishListReq, opts ...grpc.CallOption) (*PublishListResp, error) {
	client := pb.NewVideoClient(m.cli.Conn())
	return client.PublishList(ctx, in, opts...)
}

// 视频流
func (m *defaultVideoZrpcClient) Feed(ctx context.Context, in *FeedReq, opts ...grpc.CallOption) (*FeedResp, error) {
	client := pb.NewVideoClient(m.cli.Conn())
	return client.Feed(ctx, in, opts...)
}

// 更新视频点赞数
func (m *defaultVideoZrpcClient) UpdateFavoriteCount(ctx context.Context, in *UpdateFavoriteCountReq, opts ...grpc.CallOption) (*UpdateFavoriteCountResp, error) {
	client := pb.NewVideoClient(m.cli.Conn())
	return client.UpdateFavoriteCount(ctx, in, opts...)
}

// 更新视频评论数
func (m *defaultVideoZrpcClient) UpdateCommentCount(ctx context.Context, in *UpdateCommentCountReq, opts ...grpc.CallOption) (*UpdateFavoriteCountResp, error) {
	client := pb.NewVideoClient(m.cli.Conn())
	return client.UpdateCommentCount(ctx, in, opts...)
}

// 根据videoIds获取视频列表
func (m *defaultVideoZrpcClient) GetVideoListByIds(ctx context.Context, in *GetVideoListByIdsReq, opts ...grpc.CallOption) (*GetVideoListByIdsResp, error) {
	client := pb.NewVideoClient(m.cli.Conn())
	return client.GetVideoListByIds(ctx, in, opts...)
}
