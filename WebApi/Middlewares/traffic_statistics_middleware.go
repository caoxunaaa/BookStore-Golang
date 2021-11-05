package Middlewares

import (
	"WebApi/Svc"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gomodule/redigo/redis"
)

var Expire int = 60 * 10

func TrafficStatisticsMiddleware() func(c *gin.Context) {
	return func(c *gin.Context) {
		ip := c.ClientIP()
		url := c.Request.URL
		//redis 出错的情况下不记录阅读量并通知工作人员
		if repeat, err := IsRepeat(ip + url.String()); err == nil {
			if !repeat {
				err = TrafficStatistics(ip + url.String())
				if err != nil {
					fmt.Println(err)
				}
			}
		} else {
			fmt.Println(err)
		}

		c.Next()
	}
}

//判斷訪問是否在指定時間內重複 true为重复,反之false
func IsRepeat(key string) (bool, error) {
	ok, err := redis.Bool(Svc.SvcContext.Redis.Do("EXISTS", key))
	if err != nil {
		return false, err
	}
	if !ok {
		_, err = Svc.SvcContext.Redis.Do("SET", key, []byte{}, "NX", "EX", Expire)
		if err != nil {
			return false, err
		}
		return ok, nil
	} else {
		return !ok, nil
	}

}

//在redis记录访问量
func TrafficStatistics(key string) error {
	_, err := Svc.SvcContext.Redis.Do("INCR", key)
	if err != nil {
		return err
	}
	return nil
}
