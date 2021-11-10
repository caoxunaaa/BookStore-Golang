package svc

import (
	"Action/internal/config"
	"Action/model"
	"github.com/tal-tech/go-zero/core/stores/sqlx"
)

type ServiceContext struct {
	Config                config.Config
	TrafficStatisticModel model.TrafficStatisticModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:                c,
		TrafficStatisticModel: model.NewTrafficStatisticModel(sqlx.NewMysql(c.Mysql.DataSource)),
	}
}
