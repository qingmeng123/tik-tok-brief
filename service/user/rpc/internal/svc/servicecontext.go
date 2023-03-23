package svc

import (
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"tik-tok-brief/service/user/model"
	"tik-tok-brief/service/user/rpc/internal/config"
)

type ServiceContext struct {
	Config config.Config
	UserModel model.UserModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	sqlxConn:=sqlx.NewMysql(c.Mysql.DataSource)

	return &ServiceContext{
		Config: c,
		UserModel: model.NewUserModel(sqlxConn,c.CacheRedis),
	}
}

