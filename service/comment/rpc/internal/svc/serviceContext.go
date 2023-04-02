package svc

import (
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"tik-tok-brief/service/comment/model"
	"tik-tok-brief/service/comment/rpc/internal/config"
)

type ServiceContext struct {
	Config       config.Config
	CommentModel model.CommentModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	sqlxConn := sqlx.NewMysql(c.Mysql.DataSource)
	return &ServiceContext{
		Config:       c,
		CommentModel: model.NewCommentModel(sqlxConn, c.CacheRedis),
	}
}
