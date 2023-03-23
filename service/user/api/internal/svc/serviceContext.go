package svc

import (
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
	"tik-tok-brief/common/middleware"
	"tik-tok-brief/service/user/api/internal/config"
	"tik-tok-brief/service/user/rpc/user"
)

type ServiceContext struct {
	Config config.Config
	UserRPC	user.UserZrpcClient
	JwtAuthMiddleware rest.Middleware
}

func NewServiceContext(c config.Config) *ServiceContext {

	return &ServiceContext{
		Config: c,
		UserRPC: user.NewUserZrpcClient(zrpc.MustNewClient(c.UserRPC)),
		JwtAuthMiddleware: middleware.NewJwtAuth(c.JWTAuth.AccessSecret).Handle,
	}
}
