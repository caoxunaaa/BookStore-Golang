package Models

import (
	"WebApi/Pb/order"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/Shopify/sarama"
	"github.com/go-basic/uuid"
	"github.com/gomodule/redigo/redis"
	"math/rand"
	"strconv"
	"time"
)

const (
	topic                       = "order"
	key                         = "order"
	userOrderStatusRedisHashKey = "UserOrderStatus"
)

type OrderHandle interface {
	GetNotPaidOrder(ctx context.Context, buyerId int64) (*order.OrderInfoResp, error)
	GetUserStatus(userId string) (int, error)                                     //判断用户状态
	SetUserStatus(userId string, status int8) error                               //修改用户状态
	OrderLineUp(user string, v string) (partition int32, offset int64, err error) //排队
	OrderIsArrive(v string) (*order.OrderInfoReq, error)                          //查询redis是否有下单成功的标志
	ParseGoods(v string) (*struct {
		UserId   int64   `json:"user_id"`
		UserName string  `json:"user_name"`
		BookId   int64   `json:"book_id"`
		Cost     float64 `json:"cost"`
	}, error) //解析json
	DecrInventory(bookId int64) (bool, error)                  //验证库存时候足够，足够便减1,在redis中给个成功抢到，等待下单的标志
	CreateOrder(ctx context.Context, v string) (string, error) //下单处理
	PayHandle(ctx context.Context, orderNum string) error      //支付处理,模拟支付花费2秒，默认都支付成功
	StartOrderHandle(ctx context.Context, h sarama.ConsumerGroupHandler, ch chan struct{})
}

type OrderModel struct {
	OrderGrpc     order.OrderClient
	CachedConn    *redis.Pool
	KafkaProducer sarama.SyncProducer
	KafkaConsumer sarama.ConsumerGroup
}

func NewOrderModel(g order.OrderClient, c *redis.Pool, kp sarama.SyncProducer, kc sarama.ConsumerGroup) *OrderModel {
	return &OrderModel{
		OrderGrpc:     g,
		CachedConn:    c,
		KafkaProducer: kp,
		KafkaConsumer: kc,
	}
}

func (c *OrderModel) GetNotPaidOrder(ctx context.Context, buyerId int64) (*order.OrderInfoResp, error) {
	res, err := c.OrderGrpc.GetNotPaidOrderInfoByBuyerId(ctx, &order.OrderInfoReq{
		BuyerId: buyerId,
	})
	if err != nil {
		return nil, err
	}
	return res, err
}

func (c *OrderModel) GetUserStatus(userId string) (int, error) {
	//在redis中进行用户状态的存储，分别为初始：0,排队中：1，待支付：2，已支付or关闭：3, 错误：-1
	ok, err := redis.Bool(c.CachedConn.Get().Do("HEXISTS", userOrderStatusRedisHashKey, userId))
	if err != nil {
		return -1, err
	}
	if ok {
		if s, err := redis.Int(c.CachedConn.Get().Do("HGET", userOrderStatusRedisHashKey, userId)); err != nil {
			return -1, err
		} else {
			return s, nil
		}
	} else {
		return 0, nil
	}
}

func (c *OrderModel) SetUserStatus(userId string, status int8) error {
	//在redis中进行用户状态的存储，分别为初始：0,排队中：1，待支付：2，关闭：3, 错误：-1
	_, err := c.CachedConn.Get().Do("HSET", userOrderStatusRedisHashKey, userId, status)
	return err
}

func (c *OrderModel) OrderLineUp(userId string, v string) (partition int32, offset int64, err error) {
	s, err := c.GetUserStatus(userId)
	if err != nil {
		fmt.Println("GetUserStatus", err)
		return -1, -1, err
	}
	fmt.Println(s)
	switch s {
	case 0:
		break
	case 1:
		return -1, -1, errors.New("已经在排队的状态")
	case 2:
		return -1, -1, errors.New("有未完成的订单")
	default:
		err = c.SetUserStatus(userId, 0)
		if err != nil {
			return -1, -1, err
		}
		return c.OrderLineUp(userId, v)
	}

	//送入kafka中
	msg := &sarama.ProducerMessage{
		Topic: topic,
		Key:   sarama.StringEncoder(key),
		Value: sarama.StringEncoder(v),
	}
	partition, offset, err = c.KafkaProducer.SendMessage(msg)
	if err != nil {
		return partition, offset, err
	}
	//在redis中设置user排队中
	err = c.SetUserStatus(userId, 1)

	return partition, offset, err
}

func (c *OrderModel) ParseGoods(v string) (*struct {
	UserId   int64   `json:"user_id"`
	UserName string  `json:"user_name"`
	BookId   int64   `json:"book_id"`
	Cost     float64 `json:"cost"`
}, error) {
	var res = struct {
		UserId   int64   `json:"user_id"`
		UserName string  `json:"user_name"`
		BookId   int64   `json:"book_id"`
		Cost     float64 `json:"cost"`
	}{}
	err := json.Unmarshal([]byte(v), &res)
	if err != nil {
		return nil, err
	}
	return &res, nil
}

func (c *OrderModel) DecrInventory(bookId int64) (bool, error) {
	currentInventory, err := redis.Int64(c.CachedConn.Get().Do("GET", "Inventory:BookId:"+strconv.FormatInt(bookId, 10)))
	if err != nil {
		fmt.Println("DecrInventoryErr1", err)
		return false, err
	}
	if currentInventory <= 0 {
		return false, nil
	}
	_, err = c.CachedConn.Get().Do("SET", "Inventory:BookId:"+strconv.FormatInt(bookId, 10), currentInventory-1)
	if err != nil {
		fmt.Println("DecrInventoryErr2", err)
		return false, err
	}
	return true, nil
}

func (c *OrderModel) OrderIsArrive(v string) (*order.OrderInfoReq, error) {
	g, err := c.ParseGoods(v)
	if err != nil {
		return nil, err
	}
	ok, err := c.DecrInventory(g.BookId)
	if err != nil {
		return nil, err
	}
	if !ok {
		return nil, errors.New("没有库存")
	} else {
		return &order.OrderInfoReq{
			BuyerId:   g.UserId,
			OrderNum:  uuid.New(),
			OrderTime: time.Now().Format("2006-01-02 15:04:05"),
			Cost:      g.Cost,
			BookId:    g.BookId,
		}, nil
	}
}

func (c *OrderModel) CreateOrder(ctx context.Context, v string) (orderNum string, err error) {
	req, err := c.OrderIsArrive(v)
	if err != nil {
		return "", err
	}
	//订单时间随机
	rand.Seed(time.Now().UnixNano())
	num := rand.Intn(1000)
	time.Sleep(time.Millisecond * (time.Duration(num)))
	_, err = c.OrderGrpc.CreateOrderInfo(ctx, req)
	if err != nil {
		return "", err
	}
	return req.OrderNum, nil
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
	fmt.Println("orderInfo", orderInfo)
	_, err = c.OrderGrpc.UpdateOrderInfo(ctx, &order.OrderInfoReq{
		Id:          orderInfo.Id,
		BuyerId:     orderInfo.BuyerId,
		OrderNum:    orderInfo.OrderNum,
		Cost:        orderInfo.Cost,
		BookId:      orderInfo.BookId,
		IsPaid:      true,
		OrderStatus: "关闭",
	})
	if err != nil {
		return err
	}
	return nil
}

func (c *OrderModel) StartOrderHandle(ctx context.Context, h sarama.ConsumerGroupHandler, ch chan struct{}) {
	err := c.KafkaConsumer.Consume(context.Background(), []string{topic}, h)
	if err != nil {
		fmt.Println("StartOrderHandle ERR:", err.Error())
	}
}
