package Models

import (
	"WebApi/Pb/order"
	"WebApi/Services"
	"github.com/gomodule/redigo/redis"
)

type OrderHandle interface {
	ConfirmInventoryEnough() (bool, error)
	CreateOrder() error
}

type OrderModel struct {
	OrderGrpc  order.OrderClient
	CachedConn redis.Conn
}

func NewOrderModel(g *Services.GrpcContext, c redis.Conn) *OrderModel {
	return &OrderModel{
		OrderGrpc:  g.OrderGrpc,
		CachedConn: c,
	}
}

func (c *OrderModel) ConfirmInventoryEnough() (bool, error) {
	return true, nil
}

func (c *OrderModel) CreateOrder() error {
	return nil
}
