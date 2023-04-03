package svc

import (
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/zrpc"
	"tik-tok-brief/service/follow/rpc/follow"
	"tik-tok-brief/service/user/model"
	"tik-tok-brief/service/user/rpc/internal/config"
)

type ServiceContext struct {
	Config      config.Config
	UserModel   model.UserModel
	FollowerRPC follow.FollowZrpcClient
}

func NewServiceContext(c config.Config) *ServiceContext {
	sqlxConn := sqlx.NewMysql(c.Mysql.DataSource)

	return &ServiceContext{
		Config:      c,
		UserModel:   model.NewUserModel(sqlxConn, c.CacheRedis),
		FollowerRPC: follow.NewFollowZrpcClient(zrpc.MustNewClient(c.FollowRPC)),
	}
}
