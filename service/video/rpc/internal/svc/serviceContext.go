package svc

import (
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/zrpc"
	"tik-tok-brief/service/like/rpc/like"
	"tik-tok-brief/service/video/model"
	"tik-tok-brief/service/video/rpc/internal/config"
)

type ServiceContext struct {
	Config     config.Config
	VideoModel model.VideoModel
	LikeRPC    like.LikeZrpcClient
}

func NewServiceContext(c config.Config) *ServiceContext {
	sqlxConn := sqlx.NewMysql(c.Mysql.DataSource)
	return &ServiceContext{
		Config:     c,
		VideoModel: model.NewVideoModel(sqlxConn, c.CacheRedis),
		LikeRPC:    like.NewLikeZrpcClient(zrpc.MustNewClient(c.LikeRPC)),
	}
}
