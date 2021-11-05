package Databases

import (
	"WebApi/Services"
	"errors"
	"fmt"
	"github.com/gomodule/redigo/redis"
	"time"
)

func RedisInit(c *Services.Config) redis.Conn {
	return RedisPollInit(c).Get()
}

func RedisPollInit(c *Services.Config) *redis.Pool {
	return &redis.Pool{
		MaxIdle:     5,
		MaxActive:   0,
		Wait:        true,
		IdleTimeout: time.Duration(1) * time.Second,
		Dial: func() (redis.Conn, error) {
			if len(c.Redis) != 1 {
				return nil, errors.New("没有使用redis或者非单点")
			}
			r, err := redis.Dial("tcp", c.Redis[0].Host)
			if err != nil {
				fmt.Println(err)
				return nil, err
			}
			fmt.Println("no password")
			if c.Redis[0].PassWord != "" {
				if _, err := r.Do("AUTH", c.Redis[0].PassWord); err != nil {
					_ = r.Close()
					return nil, err
				}
			}
			redis.DialDatabase(0)
			return r, err
		},
	}
}

func RedisClose(c *Services.Config) {
	_ = RedisPollInit(c).Get().Close()
}
