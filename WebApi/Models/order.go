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
	orderTopic                  = "order"
	orderKey                    = "order"
	userOrderStatusRedisHashKey = "UserOrderStatus"
)

type OrderHandle interface {
	GetUserStatus(order *OrderInfo) (int, error)       //判断用户状态
	SetUserStatus(order *OrderInfo, status int8) error //修改用户状态

	OrderLineUp(orderJson string) (partition int32, offset int64, err error)        //排队
	CreateOrder(ctx context.Context, orderJson string) (orderNum string, err error) //下单处理
	PayHandle(ctx context.Context, orderNum string) error                           //支付处理,模拟支付花费2秒，默认都支付成功
	GetNotPaidOrder(ctx context.Context, buyerId int64, bookId int64) (*order.OrderInfoResp, error)
	DeleteOrder(ctx context.Context, orderNum string) error

	DecrInventory(bookId int64) (bool, error) //验证库存时候足够，足够便-1,在redis中给个成功抢到，等待下单的标志
	IncrInventory(bookId int64) (bool, error) //库存 +1
	ParseOrder(orderJson string) (*OrderInfo, error)

	StartOrderHandle(ctx context.Context, h sarama.ConsumerGroupHandler, ch chan struct{})
}

type OrderModel struct {
	OrderGrpc     order.OrderClient
	CachedConn    *redis.Pool
	KafkaProducer sarama.SyncProducer
	KafkaConsumer sarama.ConsumerGroup
}

type OrderInfo struct {
	UserId   int64   `json:"user_id"`
	UserName string  `json:"user_name"`
	BookId   int64   `json:"book_id"`
	Cost     float64 `json:"cost"`
}

func NewOrderModel(g order.OrderClient, c *redis.Pool, kp sarama.SyncProducer, kc sarama.ConsumerGroup) *OrderModel {
	return &OrderModel{
		OrderGrpc:     g,
		CachedConn:    c,
		KafkaProducer: kp,
		KafkaConsumer: kc,
	}
}

func (c *OrderModel) GetUserStatus(order *OrderInfo) (int, error) {
	//在redis中进行用户状态的存储，分别为初始：0,排队中：1，待支付：2，完成：3, 错误：-1
	//例如：  UserOrderStatus user:1:book:4 1  ->   userid=1用户 购买 bookid=4书籍，正在排队中
	key := "user:" + strconv.FormatInt(order.UserId, 10) + "book:" + strconv.FormatInt(order.BookId, 10)
	ok, err := redis.Bool(c.CachedConn.Get().Do("HEXISTS", userOrderStatusRedisHashKey, key))
	if err != nil {
		return -1, err
	}
	if ok {
		if s, err := redis.Int(c.CachedConn.Get().Do("HGET", userOrderStatusRedisHashKey, key)); err != nil {
			return -1, err
		} else {
			return s, nil
		}
	} else {
		return 0, nil
	}
}

func (c *OrderModel) SetUserStatus(order *OrderInfo, status int8) error {
	//在redis中进行用户状态的存储，分别为初始：0,排队中：1，待支付：2，完成：3, 错误：-1
	//例如：  UserOrderStatus user:1:book:4 1  ->   userid=1用户 购买 bookid=4书籍，正在排队中
	key := "user:" + strconv.FormatInt(order.UserId, 10) + "book:" + strconv.FormatInt(order.BookId, 10)
	_, err := c.CachedConn.Get().Do("HSET", userOrderStatusRedisHashKey, key, status)
	return err
}

func (c *OrderModel) OrderLineUp(orderJson string) (partition int32, offset int64, err error) {
	ord, err := c.ParseOrder(orderJson)
	if err != nil {
		return -1, -1, errors.New("ParseOrder  Error!")
	}
	s, err := c.GetUserStatus(ord)
	if err != nil {
		fmt.Println("GetUserStatus", err)
		return -1, -1, err
	}
	switch s {
	case 0:
		break
	case 1:
		return -1, -1, errors.New("排队")
	case 2:
		return -1, -1, errors.New("未完成的订单")
	case 3:
		err = c.SetUserStatus(ord, 0)
		if err != nil {
			return -1, -1, err
		}
		return c.OrderLineUp(orderJson)
	default:
		return -1, -1, errors.New("redis出错")
	}

	//送入kafka中
	msg := &sarama.ProducerMessage{
		Topic: orderTopic,
		Key:   sarama.StringEncoder(orderKey),
		Value: sarama.StringEncoder(orderJson),
	}
	partition, offset, err = c.KafkaProducer.SendMessage(msg)
	if err != nil {
		return partition, offset, err
	}
	//在redis中设置user排队中
	err = c.SetUserStatus(ord, 1)

	return partition, offset, err
}

func (c *OrderModel) DecrInventory(bookId int64) (bool, error) {
	ok, err := redis.Bool(c.CachedConn.Get().Do("EXISTS", "Inventory:BookId:"+strconv.FormatInt(bookId, 10)))
	if err != nil {
		fmt.Println("exists", err)
		return false, err
	}
	if !ok {
		return false, nil
	}

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

func (c *OrderModel) IncrInventory(bookId int64) (bool, error) {
	ok, err := redis.Bool(c.CachedConn.Get().Do("EXISTS", "Inventory:BookId:"+strconv.FormatInt(bookId, 10)))
	if err != nil {
		fmt.Println("exists", err)
		return false, err
	}
	if !ok {
		return false, nil
	}
	currentInventory, err := redis.Int64(c.CachedConn.Get().Do("GET", "Inventory:BookId:"+strconv.FormatInt(bookId, 10)))
	if err != nil {
		fmt.Println("IncrInventoryErr1", err)
		return false, err
	}
	_, err = c.CachedConn.Get().Do("SET", "Inventory:BookId:"+strconv.FormatInt(bookId, 10), currentInventory+1)
	if err != nil {
		fmt.Println("IncrInventoryErr2", err)
		return false, err
	}
	return true, nil
}

func (c *OrderModel) InventoryIsEnough(ord *OrderInfo) (*order.OrderInfoReq, error) {
	ok, err := c.DecrInventory(ord.BookId)
	if err != nil {
		return nil, err
	}
	if !ok {
		return nil, errors.New("没有库存")
	} else {
		return &order.OrderInfoReq{
			BuyerId:   ord.UserId,
			OrderNum:  uuid.New(),
			OrderTime: time.Now().Format("2006-01-02 15:04:05"),
			Cost:      ord.Cost,
			BookId:    ord.BookId,
		}, nil
	}
}

// 库存-1，生成订单号,返回订单号
func (c *OrderModel) CreateOrder(ctx context.Context, orderJson string) (orderNum string, err error) {
	ord, err := c.ParseOrder(orderJson)
	if err != nil {
		return "", err
	}

	s, err := c.GetUserStatus(ord)
	if err != nil {
		fmt.Println("GetUserStatus", err)
		return "", err
	}
	switch s {
	case 0:
		return "", errors.New("无状态")
	case 1:
		break
	case 2:
		return "", errors.New("未完成的订单")
	case 3:
		err = c.SetUserStatus(ord, 0)
		if err != nil {
			return "", errors.New("无状态")
		}
	default:
		return "", errors.New("redis出错")
	}

	orderReq, err := c.InventoryIsEnough(ord)
	if err != nil {
		return "", err
	}
	//订单时间随机
	rand.Seed(time.Now().UnixNano())
	num := rand.Intn(1000)
	time.Sleep(time.Millisecond * (time.Duration(num)))

	//设置订单超时时间
	_, err = c.CachedConn.Get().Do("SET", orderReq.OrderNum, 1, "EX", 60*1)
	if err != nil {
		return "", err
	}
	_, err = c.OrderGrpc.CreateOrderInfo(ctx, orderReq)
	if err != nil {
		_, err = c.CachedConn.Get().Do("DEL", orderReq.OrderNum)
		if err != nil {
			return "", err
		}
		return "", err
	}
	//在redis中设置user待支付
	err = c.SetUserStatus(ord, 2)

	return orderReq.OrderNum, nil
}

func (c *OrderModel) GetNotPaidOrder(ctx context.Context, buyerId int64, bookId int64) (*order.OrderInfoResp, error) {
	s, err := c.GetUserStatus(&OrderInfo{
		UserId: buyerId,
		BookId: bookId,
	})
	if err != nil {
		fmt.Println("GetUserStatus", err)
		return nil, err
	}

	fmt.Println("GetNotPaidOrder status", s)

	switch s {
	case 0:
		return nil, errors.New("无状态")
	case 1:
		return nil, errors.New("排队")
	case 2:
		break
	case 3:
		err = c.SetUserStatus(&OrderInfo{
			UserId: buyerId,
			BookId: bookId,
		}, 0)
		if err != nil {
			return nil, errors.New("无状态")
		}
	default:
		return nil, errors.New("redis出错")
	}

	res, err := c.OrderGrpc.GetNotPaidOrderInfoByBuyerId(ctx, &order.OrderInfoReq{
		BuyerId: buyerId,
	})
	if err != nil {
		return nil, err
	}

	//判断订单是否超时

	timeIn, err := redis.Bool(c.CachedConn.Get().Do("EXISTS", res.OrderNum))
	if err != nil {
		return nil, errors.New("redis出错")
	}
	if !timeIn {
		err = c.DeleteOrder(ctx, res.OrderNum)
		if err != nil {
			return nil, errors.New("redis出错")
		}
		return nil, errors.New("订单超时未处理")
	}

	return res, err
}

func (c *OrderModel) PayHandle(ctx context.Context, orderNum string) error {
	rand.Seed(time.Now().UnixNano())
	num := rand.Intn(1000)
	fmt.Println(orderNum, num)
	if num > 900 {
		//测试,随机支付失败
		return errors.New("支付失败")
	}

	time.Sleep(time.Second * 1)
	// Pay Ok
	orderInfo, err := c.OrderGrpc.GetOrderInfoByOrderNum(ctx, &order.OrderInfoReq{
		OrderNum: orderNum,
	})
	if err != nil {
		return errors.New("订单超时或因未知原因已丢失")
	}

	s, err := c.GetUserStatus(&OrderInfo{
		UserId: orderInfo.BuyerId,
		BookId: orderInfo.BookId,
	})
	if err != nil {
		fmt.Println("GetUserStatus", err)
		return err
	}
	switch s {
	case 0:
		return errors.New("无状态")
	case 1:
		return errors.New("排队")
	case 2:
		break
	case 3:
		err = c.SetUserStatus(&OrderInfo{
			UserId: orderInfo.BuyerId,
			BookId: orderInfo.BookId,
		}, 0)
		if err != nil {
			return errors.New("无状态")
		}
	default:
		return errors.New("redis出错")
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
	err = c.SetUserStatus(&OrderInfo{
		UserId: orderInfo.BuyerId,
		BookId: orderInfo.BookId,
	}, 3)
	if err != nil {
		return errors.New("完成")
	}

	return nil
}

func (c *OrderModel) StartOrderHandle(ctx context.Context, h sarama.ConsumerGroupHandler, ch chan struct{}) {
	err := c.KafkaConsumer.Consume(context.Background(), []string{orderTopic}, h)
	if err != nil {
		fmt.Println("StartOrderHandle ERR:", err.Error())
	}
}

func (c *OrderModel) DeleteOrder(ctx context.Context, orderNum string) error {
	// 返还库存1,  并且delete订单
	orderInfo, err := c.OrderGrpc.GetOrderInfoByOrderNum(ctx, &order.OrderInfoReq{
		OrderNum: orderNum,
	})
	if err != nil {
		return err
	}

	s, err := c.GetUserStatus(&OrderInfo{
		UserId: orderInfo.BuyerId,
		BookId: orderInfo.BookId,
	})
	if err != nil {
		fmt.Println("GetUserStatus", err)
		return err
	}
	switch s {
	case 0:
		return errors.New("无状态")
	case 1:
		return errors.New("排队")
	case 2:
		break
	case 3:
		err = c.SetUserStatus(&OrderInfo{
			UserId: orderInfo.BuyerId,
			BookId: orderInfo.BookId,
		}, 0)
		if err != nil {
			return errors.New("无状态")
		}
	default:
		return errors.New("redis出错")
	}

	_, err = c.OrderGrpc.DeleteOrderInfo(ctx, &order.OrderInfoReq{
		Id: orderInfo.Id,
	})
	if err != nil {
		return err
	}
	_, err = c.IncrInventory(orderInfo.BookId)
	if err != nil {
		return errors.New("redis出错")
	}

	err = c.SetUserStatus(&OrderInfo{UserId: orderInfo.BuyerId, BookId: orderInfo.BookId}, 0)
	if err != nil {
		return err
	}
	return nil
}

func (c *OrderModel) ParseOrder(orderJson string) (*OrderInfo, error) {
	var res = OrderInfo{}
	err := json.Unmarshal([]byte(orderJson), &res)
	if err != nil {
		return nil, err
	}
	return &res, nil
}
