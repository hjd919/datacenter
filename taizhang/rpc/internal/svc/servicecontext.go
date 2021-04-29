package svc

import (
	"datacenter/taizhang/model"
	"datacenter/taizhang/rpc/internal/config"
	"github.com/tal-tech/go-zero/core/stores/sqlx"
)

type ServiceContext struct {
	Config           config.Config
	AppTaizhangModel model.AppTaizhangModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	conn := sqlx.NewMysql(c.Mysql.DataSource)
	atm := model.NewAppTaizhangModel(conn, c.CacheRedis)

	return &ServiceContext{
		Config:           c,
		AppTaizhangModel: atm,
	}
}
