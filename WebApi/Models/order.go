package Models

import (
	"WebApi/Pb/order"
	"WebApi/Services"
	"context"
	"github.com/gomodule/redigo/redis"
	"strconv"
)

type OrderHandle interface {
	OrderLineUp(key string) error                                   //排队
	OrderIsArrive(key string) (bool, error)                         //查询redis是否有下单成功的标志
	DecrInventory(key int64) (bool, error)                          //验证库存时候足够，足够便减1,在redis中给个成功抢到，等待下单的标志
	CreateOrder(ctx context.Context, req *order.OrderInfoReq) error //下单处理
	PayHandle(ctx context.Context, req *order.OrderInfoReq) error   //支付处理
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

func (c *OrderModel) OrderLineUp(key string) error {
	return nil
}

func (c *OrderModel) OrderIsArrive(key string) (bool, error) {
	return true, nil
}

func (c *OrderModel) DecrInventory(bookId int64) (bool, error) {
	inventory, err := redis.Int64(c.CachedConn.Do("DECR", "Inventory:BookId:"+strconv.FormatInt(bookId, 10)))
	if err != nil {
		return false, err
	}
	if inventory < 0 {
		return false, nil
	}
	return true, nil
}

func (c *OrderModel) CreateOrder(ctx context.Context, req *order.OrderInfoReq) error {
	if _, err := c.OrderGrpc.CreateOrderInfo(ctx, req); err != nil {
		return err
	} else {
		return nil
	}
}

func (c *OrderModel) PayHandle(ctx context.Context, req *order.OrderInfoReq) error {
	return nil
}
