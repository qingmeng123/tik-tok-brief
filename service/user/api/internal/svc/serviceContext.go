package svc

import (
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
	"tik-tok-brief/service/chat/rpc/chat"
	"tik-tok-brief/service/follow/rpc/follow"
	"tik-tok-brief/service/user/api/internal/config"
	"tik-tok-brief/service/user/api/internal/middleware"
	"tik-tok-brief/service/user/rpc/user"
)

type ServiceContext struct {
	Config            config.Config
	UserRPC           user.UserZrpcClient
	FollowRPC         follow.FollowZrpcClient
	ChatRPC           chat.Chat
	JwtAuthMiddleware rest.Middleware
}

func NewServiceContext(c config.Config) *ServiceContext {

	return &ServiceContext{
		Config:            c,
		UserRPC:           user.NewUserZrpcClient(zrpc.MustNewClient(c.UserRPC)),
		FollowRPC:         follow.NewFollowZrpcClient(zrpc.MustNewClient(c.FollowRPC)),
		ChatRPC:           chat.NewChat(zrpc.MustNewClient(c.ChatRPC)),
		JwtAuthMiddleware: middleware.NewJwtAuthMiddleware(c.JWTAuth.AccessSecret).Handle,
	}
}
