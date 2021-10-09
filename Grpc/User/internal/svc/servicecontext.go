package svc

import (
	"User/internal/config"
	"User/model"
	"github.com/tal-tech/go-zero/core/stores/sqlx"
)

type ServiceContext struct {
	Config config.Config
	Model  model.UserInfoModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	con := sqlx.NewMysql(c.Mysql.DataSource)
	return &ServiceContext{
		Config: c,
		Model:  model.NewUserInfoModel(con, c.CacheRedis),
	}
}
