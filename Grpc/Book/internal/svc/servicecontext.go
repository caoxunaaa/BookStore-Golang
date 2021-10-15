package svc

import (
	"Book/internal/config"
	"Book/model"
	"github.com/tal-tech/go-zero/core/stores/sqlx"
)

type ServiceContext struct {
	Config             config.Config
	BookBasicInfoModel model.BookBasicInfoModel
	BookContentModel   model.BookContentModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	con := sqlx.NewMysql(c.Mysql.DataSource)
	return &ServiceContext{
		Config:             c,
		BookBasicInfoModel: model.NewBookBasicInfoModel(con, c.CacheRedis),
		BookContentModel:   model.NewBookContentModel(con, c.CacheRedis),
	}
}
