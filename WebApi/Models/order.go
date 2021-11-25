package Models

import (
	"WebApi/Pb/order"
	"context"
	"github.com/Shopify/sarama"
	"github.com/gomodule/redigo/redis"
	"strconv"
	"time"
)

const (
	topic = "order"
	key   = "order"
)

type OrderHandle interface {
	OrderLineUp(key string) (int32, int64, error)                   //排队
	OrderIsArrive(v string) (bool, error)                           //查询redis是否有下单成功的标志
	DecrInventory(key int64) (bool, error)                          //验证库存时候足够，足够便减1,在redis中给个成功抢到，等待下单的标志
	CreateOrder(ctx context.Context, req *order.OrderInfoReq) error //下单处理
	PayHandle(ctx context.Context, orderNum string) error           //支付处理,模拟支付花费2秒，默认都支付成功
}

type OrderModel struct {
	OrderGrpc     order.OrderClient
	CachedConn    redis.Conn
	KafkaProducer sarama.SyncProducer
	KafkaConsumer sarama.Consumer
}

func NewOrderModel(g order.OrderClient, c redis.Conn, kp sarama.SyncProducer, kc sarama.Consumer) *OrderModel {
	return &OrderModel{
		OrderGrpc:     g,
		CachedConn:    c,
		KafkaProducer: kp,
		KafkaConsumer: kc,
	}
}

func (c *OrderModel) OrderLineUp(v string) (int32, int64, error) {
	msg := &sarama.ProducerMessage{
		Topic: topic,
		Key:   sarama.StringEncoder(key),
		Value: sarama.StringEncoder(v),
	}
	if partition, offset, err := c.KafkaProducer.SendMessage(msg); err != nil {
		return partition, offset, err
	} else {
		return partition, offset, err
	}
}

func (c *OrderModel) OrderIsArrive(kv string) (bool, error) {
	ok, err := redis.Bool(c.CachedConn.Do("EXISTS", kv))
	if err != nil {
		return false, err
	}
	if ok {
		return true, nil
	} else {
		return false, nil
	}

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

func (c *OrderModel) PayHandle(ctx context.Context, orderNum string) error {
	time.Sleep(time.Second * 1)
	// Pay Ok
	orderInfo, err := c.OrderGrpc.GetOrderInfoByOrderNum(ctx, &order.OrderInfoReq{
		OrderNum: orderNum,
	})
	if err != nil {
		return err
	}
	_, err = c.OrderGrpc.UpdateOrderInfo(ctx, &order.OrderInfoReq{
		Id:          orderInfo.Id,
		BuyerId:     orderInfo.BuyerId,
		OrderNum:    orderInfo.OrderNum,
		Cost:        orderInfo.Cost,
		IsPaid:      true,
		OrderStatus: "已支付",
	})
	if err != nil {
		return err
	}
	return nil
}
