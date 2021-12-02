package order

import (
	"WebApi/Pb/order"
	"WebApi/Svc"
	"context"
	"fmt"
	"github.com/Shopify/sarama"
	"github.com/gin-gonic/gin"
	"github.com/gomodule/redigo/redis"
	"math/rand"
	"net/http"
	"strconv"
	"time"
)

const MANOEUVRABLE = 100

var ALLOCATING = false //允许排队FLAG

func LineUpHandler(c *gin.Context) {
	buyerId := c.PostForm("buyerId")
	order := c.PostForm("order") //json {'user_id': '', 'user_name': '', 'book_id': '', 'cost': 100}
	fmt.Println(buyerId, order)
	if ALLOCATING {
		if _, offset, err := Svc.SvcContext.Model.Order.OrderLineUp(buyerId, order); err != nil {
			c.JSON(http.StatusNotFound, err)
		} else {
			c.JSON(http.StatusOK, gin.H{"message": "排在" + strconv.FormatInt(offset, 10) + "位"})
		}
	} else {
		c.JSON(http.StatusOK, gin.H{"error": "热卖还未开始，不允许排队"})
	}
}

func StartOrderHandler(c *gin.Context) {
	_, _ = Svc.SvcContext.Redis.Get().Do("SET", "Inventory:BookId:4", 10) //假数据
	ALLOCATING = true
	c.JSON(http.StatusOK, gin.H{"message": "start"})
}

func EndOrderHandler(c *gin.Context) {
	ALLOCATING = false
	c.JSON(http.StatusOK, gin.H{"message": "end"})
}

func HotSaleHandler() {
	ch := make(chan struct{}, MANOEUVRABLE) //允许同时处理多少个订单
	go Svc.SvcContext.Model.Order.StartOrderHandle(context.Background(), msgConsumerGroupHandler{channel: ch}, ch)
}

func GetNotPaidOrderInfoHandler(c *gin.Context) {
	buyerId, err := strconv.ParseInt(c.Query("buyerId"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}
	res, err := Svc.SvcContext.Model.Order.GetNotPaidOrder(context.Background(), buyerId)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}
	timeout, err := redis.Bool(Svc.SvcContext.Model.Order.CachedConn.Get().Do("EXISTS", res.OrderNum))
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}
	if timeout {
		c.JSON(http.StatusOK, res)
	} else {
		// 返还库存1,  并且关闭订单
		_, err = Svc.SvcContext.Redis.Get().Do("INCR", "Inventory:BookId:"+strconv.Itoa(int(res.BookId)))
		if err != nil {
			c.JSON(http.StatusBadRequest, err)
			return
		}
		err = Svc.SvcContext.Model.Order.SetUserStatus(strconv.Itoa(int(res.BuyerId)), 3)
		if err != nil {
			c.JSON(http.StatusBadRequest, err)
			return
		}
		_, err = Svc.SvcContext.Model.Order.OrderGrpc.UpdateOrderInfo(context.Background(), &order.OrderInfoReq{
			Id:          res.Id,
			BuyerId:     res.BuyerId,
			OrderNum:    res.OrderNum,
			Cost:        res.Cost,
			IsPaid:      false,
			OrderStatus: "关闭",
			BookId:      res.BookId,
		})

		c.JSON(http.StatusOK, gin.H{"error": "订单已经超时"})
	}

}

func PayHandler(c *gin.Context) {
	orderNum := c.PostForm("orderNum")
	//假支付，随机成功
	rand.Seed(time.Now().UnixNano())
	num := rand.Intn(1000)
	if num < 900 {
		err := Svc.SvcContext.Model.Order.PayHandle(context.Background(), orderNum)
		if err != nil {
			c.JSON(http.StatusBadRequest, err)
		}
		c.JSON(http.StatusOK, gin.H{"message": "支付成功"})
	} else {
		c.JSON(http.StatusOK, gin.H{"message": "支付失败"})
	}
}

type msgConsumerGroupHandler struct {
	channel chan struct{}
}

func (msgConsumerGroupHandler) Setup(_ sarama.ConsumerGroupSession) error   { return nil }
func (msgConsumerGroupHandler) Cleanup(_ sarama.ConsumerGroupSession) error { return nil }
func (h msgConsumerGroupHandler) ConsumeClaim(sess sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	for msg := range claim.Messages() {
		fmt.Printf("Message topic:%q partition:%d offset:%d  value:%s\n", msg.Topic, msg.Partition, msg.Offset, string(msg.Value))

		sess.MarkMessage(msg, "")

		// 标记，sarama会自动进行提交，默认间隔1秒
		//查询库存，创建订单, 设置订单过期时间, 排队超时，取消排队
		h.channel <- struct{}{}

		go func(v string) {
			ctx, cancel := context.WithTimeout(context.Background(), time.Minute*5)
			defer cancel()
			for {
				select {
				case <-ctx.Done():
					fmt.Println("order timeout")
					return
				default:
					if orderNum, err := Svc.SvcContext.Model.Order.CreateOrder(context.Background(), v); err == nil {
						fmt.Println("Create Order Ok")
						_, err = Svc.SvcContext.Model.Order.CachedConn.Get().Do("SET", orderNum, 1, "EX", 60*5)
						if err != nil {
							fmt.Println(err)
							return
						}
						<-h.channel
						return
					} else {
						fmt.Println(err)
					}
				}
				time.Sleep(time.Second * 1)
			}
		}(string(msg.Value))

	}
	return nil
}
