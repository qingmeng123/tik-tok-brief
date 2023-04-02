package svc

import (
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"tik-tok-brief/service/chat/model"
	"tik-tok-brief/service/chat/rpc/internal/config"
)

type ServiceContext struct {
	Config    config.Config
	ChatModel model.ChatModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	sqlxConn := sqlx.NewMysql(c.Mysql.DataSource)
	return &ServiceContext{
		Config:    c,
		ChatModel: model.NewChatModel(sqlxConn, c.CacheRedis),
	}
}
