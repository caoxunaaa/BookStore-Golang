package order

import (
	"WebApi/Pb/book"
	"WebApi/Svc"
	"context"
	"fmt"
	"github.com/Shopify/sarama"
	"github.com/gin-gonic/gin"
	"github.com/gomodule/redigo/redis"
	"net/http"
	"strconv"
	"time"
)

const MANOEUVRABLE = 100

var ALLOCATING = false //允许排队FLAG

func LineUpHandler(c *gin.Context) {
	ord := c.PostForm("order") //json {'user_id': '', 'user_name': '', 'book_id': '', 'cost': 100}
	fmt.Println(ord)
	if ALLOCATING {
		orderInfo, err := Svc.SvcContext.Model.Order.ParseOrder(ord)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
			return
		}
		ok, err := redis.Bool(Svc.SvcContext.Redis.Get().Do("EXISTS", "Inventory:BookId:"+strconv.FormatInt(orderInfo.BookId, 10)))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
			return
		}
		if ok {
			invertory, err := redis.Int64(Svc.SvcContext.Redis.Get().Do("GET", "Inventory:BookId:"+strconv.FormatInt(orderInfo.BookId, 10)))
			if err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
				return
			}
			if invertory <= 0 {
				c.JSON(http.StatusOK, gin.H{"code": 2003, "message": "没有库存"})
				return
			}
		} else {
			c.JSON(http.StatusOK, gin.H{"code": 2003, "message": "没有库存"})
			return
		}

		if _, offset, err := Svc.SvcContext.Model.Order.OrderLineUp(ord); err != nil {
			if err.Error() == "未完成的订单" {
				c.JSON(http.StatusOK, gin.H{"code": 2001, "message": err.Error()})
				return
			}
			c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		} else {
			c.JSON(http.StatusOK, gin.H{"code": 2000, "message": "排在" + strconv.FormatInt(offset, 10) + "位"})
		}
	} else {
		c.JSON(http.StatusOK, gin.H{"code": 2002, "message": "热卖还未开始，不允许排队"})
	}
}

func StartOrderHandler(c *gin.Context) {
	_, _ = Svc.SvcContext.Redis.Get().Do("SET", "Inventory:BookId:4", 3) //假数据
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

	bookId, err := strconv.ParseInt(c.Query("bookId"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	res, err := Svc.SvcContext.Model.Order.GetNotPaidOrder(context.Background(), buyerId, bookId)
	if err != nil {
		if err.Error() == "无状态" {
			c.JSON(http.StatusOK, gin.H{"code": 2001, "message": "库存不足，排队已超时"})
			return
		} else if err.Error() == "排队" {
			c.JSON(http.StatusOK, gin.H{"code": 2002, "message": err.Error()})
			return
		} else if err.Error() == "订单超时未处理" {
			c.JSON(http.StatusOK, gin.H{"code": 2003, "message": err.Error()})
			return
		} else if err.Error() == "redis出错" {
			c.JSON(http.StatusOK, gin.H{"code": 2005, "message": err.Error()})
			return
		} else {
			c.JSON(http.StatusBadRequest, gin.H{"code": 1001, "message": err.Error()})
			return
		}
	}
	bookInfo, err := Svc.SvcContext.Grpc.BookGrpc.FindOneBookById(context.Background(), &book.BookBasicInfoReq{Id: bookId})
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"code": 1001, "message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"code": 2000, "message": res, "bookName": bookInfo.Name})
}

func CancelOrderHandler(c *gin.Context) {
	orderNum := c.PostForm("orderNum")
	err := Svc.SvcContext.Model.Order.DeleteOrder(context.Background(), orderNum)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"code": 2001, "message": "没有查询到未支付订单"})
		return
	}
}

func PayHandler(c *gin.Context) {
	orderNum := c.PostForm("orderNum")
	//假支付，随机成功
	err := Svc.SvcContext.Model.Order.PayHandle(context.Background(), orderNum)
	if err != nil {
		if err.Error() == "支付失败" {
			c.JSON(http.StatusOK, gin.H{"code": 2001, "message": err.Error()})
			return
		}
		c.JSON(http.StatusBadRequest, err)
		return
	}
	c.JSON(http.StatusOK, gin.H{"code": 2000, "message": "支付成功"})
}

type msgConsumerGroupHandler struct {
	channel chan struct{}
}

func (msgConsumerGroupHandler) Setup(_ sarama.ConsumerGroupSession) error   { return nil }
func (msgConsumerGroupHandler) Cleanup(_ sarama.ConsumerGroupSession) error { return nil }
func (h msgConsumerGroupHandler) ConsumeClaim(sess sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	for msg := range claim.Messages() {
		fmt.Printf("Message topic:%q partition:%d offset:%d  value:%s\n", msg.Topic, msg.Partition, msg.Offset, string(msg.Value))

		//查询库存，创建订单, 设置订单过期时间, 排队超时，取消排队
		h.channel <- struct{}{}
		ctx, cancel := context.WithTimeout(context.TODO(), time.Minute*1)
		go func(ctx context.Context, v string) {
			defer cancel()
			ticker := time.NewTicker(1 * time.Second)
			for _ = range ticker.C {
				select {
				case <-ctx.Done():
					fmt.Println("order timeout")
					ord, err := Svc.SvcContext.Model.Order.ParseOrder(v)
					if err != nil {
						fmt.Println(err)
					}
					err = Svc.SvcContext.Model.Order.SetUserStatus(ord, 0)
					if err != nil {
						fmt.Println(err)
					}
					return
				default:
					if orderNum, err := Svc.SvcContext.Model.Order.CreateOrder(context.Background(), v); err == nil {
						fmt.Println("Create Order Ok: " + orderNum)
						<-h.channel
						return
					} else {
						fmt.Println(err)
					}
				}
			}
		}(ctx, string(msg.Value))

		// 标记，sarama会自动进行提交，默认间隔1秒
		sess.MarkMessage(msg, "")
	}
	return nil
}
