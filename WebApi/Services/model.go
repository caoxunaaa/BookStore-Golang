package Services

import (
	"WebApi/Models"
	"WebApi/Pb/order"
	"github.com/gomodule/redigo/redis"
)

type ModelContext struct {
	Order *Models.OrderModel
}

func NewModel(grpc order.OrderClient, conn redis.Conn, kafka *KafkaContext) *ModelContext {
	var m ModelContext
	m.Order = Models.NewOrderModel(grpc, conn, kafka.Producer, kafka.Consumer)
	return &m
}
