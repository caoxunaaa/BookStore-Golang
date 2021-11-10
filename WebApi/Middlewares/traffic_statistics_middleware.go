package Middlewares

import (
	"WebApi/Svc"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gomodule/redigo/redis"
	"regexp"
)

var Expire = 10

func TrafficStatisticsMiddleware() func(c *gin.Context) {
	return func(c *gin.Context) {
		ip := c.ClientIP()
		url := c.Request.URL
		//redis 出错的情况下不记录阅读量并通知工作人员
		if repeat, err := IsRepeat(ip + url.String()); err == nil {
			if !repeat {
				err = TrafficStatistics(url.String())
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
	}
	//重复
	return ok, nil

}

//在redis记录访问量
func TrafficStatistics(key string) error {
	re, err := regexp.Compile("[0-9]+") //解析出来哪本书哪个章节
	if err != nil {
		fmt.Println(err)
	}
	res := re.FindAll([]byte(key), -1)

	key = "traffic_statistic"
	if len(res) == 2 {
		member := string(res[0]) + ":" + string(res[1])
		_, err := Svc.SvcContext.Redis.Do("ZINCRBY", key, 1, member)
		if err != nil {
			return err
		}
		return nil
	} else {
		return errors.New("url不是正确的格式，无法用正则表达式匹配")
	}
}
