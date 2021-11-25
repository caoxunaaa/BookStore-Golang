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
	kafka := Services.NewKafka(c)
	return &ServiceContext{
		Grpc:  grpc,
		Redis: conn,
		Kafka: kafka,
		Model: Services.NewModel(grpc.OrderGrpc, conn, kafka),
	}
}
