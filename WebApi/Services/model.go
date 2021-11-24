package Services

import (
	"WebApi/Models"
	"github.com/gomodule/redigo/redis"
)

type ModelContext struct {
	Order *Models.OrderModel
}

func NewModel(grpc *GrpcContext, conn redis.Conn) *ModelContext {
	var m ModelContext
	m.Order = Models.NewOrderModel(grpc, conn)
	return &m
}
