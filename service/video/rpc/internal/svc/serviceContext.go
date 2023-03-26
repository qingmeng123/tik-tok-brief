package svc

import (
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"tik-tok-brief/service/video/model"
	"tik-tok-brief/service/video/rpc/internal/config"
)

type ServiceContext struct {
	Config     config.Config
	VideoModel model.VideoModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	sqlxConn := sqlx.NewMysql(c.Mysql.DataSource)
	return &ServiceContext{
		Config:     c,
		VideoModel: model.NewVideoModel(sqlxConn, c.CacheRedis),
	}
}
