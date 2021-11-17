package svc

import (
	"Action/internal/config"
	"Action/model"
	"github.com/tal-tech/go-zero/core/stores/sqlx"
)

type ServiceContext struct {
	Config                config.Config
	TrafficStatisticModel model.TrafficStatisticModel
	CommentModel          model.CommentModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	conn := sqlx.NewMysql(c.Mysql.DataSource)
	return &ServiceContext{
		Config:                c,
		TrafficStatisticModel: model.NewTrafficStatisticModel(conn),
		CommentModel:          model.NewCommentModel(conn),
	}
}
