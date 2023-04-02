package svc

import (
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"tik-tok-brief/service/like/model"
	"tik-tok-brief/service/like/rpc/internal/config"
)

type ServiceContext struct {
	Config    config.Config
	LikeModel model.LikeModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	sqlxConn := sqlx.NewMysql(c.Mysql.DataSource)
	return &ServiceContext{
		Config:    c,
		LikeModel: model.NewLikeModel(sqlxConn, c.CacheRedis),
	}
}
