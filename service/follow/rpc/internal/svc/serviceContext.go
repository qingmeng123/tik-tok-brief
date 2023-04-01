package svc

import (
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"tik-tok-brief/service/follow/model"
	"tik-tok-brief/service/follow/rpc/internal/config"
)

type ServiceContext struct {
	Config      config.Config
	FollowModel model.FollowModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	sqlxConn := sqlx.NewMysql(c.Mysql.DataSource)

	return &ServiceContext{
		Config:      c,
		FollowModel: model.NewFollowModel(sqlxConn, c.CacheRedis),
	}
}
