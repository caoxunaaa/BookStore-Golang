package svc

import (
	"Order/internal/config"
	"Order/model"
	"github.com/tal-tech/go-zero/core/stores/sqlx"
)

type ServiceContext struct {
	Config          config.Config
	OrderInfoModel  model.OrderInfoModel
	OrderGoodsModel model.OrderGoodsModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	con := sqlx.NewMysql(c.Mysql.DataSource)
	return &ServiceContext{
		Config:          c,
		OrderInfoModel:  model.NewOrderInfoModel(con, c.CacheRedis),
		OrderGoodsModel: model.NewOrderGoodsModel(con, c.CacheRedis),
	}
}
