package svc

import (
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
	"tik-tok-brief/service/comment/rpc/comment"
	"tik-tok-brief/service/file/rpc/file"
	"tik-tok-brief/service/follow/rpc/follow"
	"tik-tok-brief/service/like/rpc/like"
	"tik-tok-brief/service/user/rpc/user"
	"tik-tok-brief/service/video/api/internal/config"
	"tik-tok-brief/service/video/api/internal/middleware"
	"tik-tok-brief/service/video/rpc/video"
)

type ServiceContext struct {
	Config                    config.Config
	ParseFormMiddleware       rest.Middleware
	JWTAuthMiddleware         rest.Middleware
	JWTOptionalAuthMiddleware rest.Middleware
	FileRPC                   file.File
	VideoRPC                  video.VideoZrpcClient
	UserRPC                   user.UserZrpcClient
	LikeRPC                   like.LikeZrpcClient
	CommentRPC                comment.CommentZrpcClient
	FollowRPC                 follow.FollowZrpcClient
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:                    c,
		ParseFormMiddleware:       middleware.NewParseFormMiddleware(c.JWTAuth.AccessSecret, c.MaxVideoSize).Handle,
		JWTAuthMiddleware:         middleware.NewJWTAuthMiddleware(c.JWTAuth.AccessSecret).Handle,
		JWTOptionalAuthMiddleware: middleware.NewJWTOptionalAuthMiddleware(c.JWTAuth.AccessSecret).Handle,
		FileRPC:                   file.NewFile(zrpc.MustNewClient(c.FileRPC)),
		VideoRPC:                  video.NewVideoZrpcClient(zrpc.MustNewClient(c.VideoRPC)),
		UserRPC:                   user.NewUserZrpcClient(zrpc.MustNewClient(c.UserRPC)),
		LikeRPC:                   like.NewLikeZrpcClient(zrpc.MustNewClient(c.LikeRPC)),
		CommentRPC:                comment.NewCommentZrpcClient(zrpc.MustNewClient(c.CommentRPC)),
		FollowRPC:                 follow.NewFollowZrpcClient(zrpc.MustNewClient(c.FollowRPC)),
	}
}
