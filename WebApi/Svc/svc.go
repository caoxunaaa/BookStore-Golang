package Svc

import (
	"WebApi/Databases"
	"WebApi/Services"
	"github.com/gomodule/redigo/redis"
)

var SvcContext *ServiceContext

type ServiceContext struct {
	Grpc  *Services.GrpcContext
	Redis redis.Conn
}

func NewContext(c *Services.Config) *ServiceContext {
	return &ServiceContext{
		Grpc:  Services.GrpcInit(c),
		Redis: Databases.RedisInit(c),
	}
}
