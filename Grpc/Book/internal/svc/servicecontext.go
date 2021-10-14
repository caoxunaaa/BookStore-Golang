package svc

import (
	"Book/internal/config"
	"Book/model"
	"github.com/tal-tech/go-zero/core/stores/sqlx"
)

type ServiceContext struct {
	Config config.Config
	Model  model.BookBasicInfoModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	con := sqlx.NewMysql(c.Mysql.DataSource)
	return &ServiceContext{
		Config: c,
		Model:  model.NewBookBasicInfoModel(con, c.CacheRedis),
	}
}
