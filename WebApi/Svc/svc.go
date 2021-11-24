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
	Model *Services.ModelContext
	Kafka *Services.KafkaContext
}

func NewContext(c *Services.Config) *ServiceContext {
	grpc := Services.GrpcInit(c)
	conn := Databases.RedisInit(c)
	return &ServiceContext{
		Grpc:  grpc,
		Redis: conn,
		Model: Services.NewModel(grpc, conn),
		Kafka: Services.NewKafka(c),
	}
}
