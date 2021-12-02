package Databases

import (
	"WebApi/Services"
	"errors"
	"fmt"
	"github.com/gomodule/redigo/redis"
	"time"
)

// redigo

func RedisPollInit(c *Services.Config) *redis.Pool {
	return &redis.Pool{
		MaxIdle:     5,
		MaxActive:   0,
		Wait:        false,
		IdleTimeout: time.Duration(5) * time.Second,
		Dial: func() (redis.Conn, error) {
			if len(c.Redis) != 1 {
				return nil, errors.New("没有使用redis或者非单点")
			}
			r, err := redis.Dial(
				"tcp",
				c.Redis[0].Host,
			)
			if err != nil {
				fmt.Println(err)
				return nil, err
			}
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

func RedisClientInit(c *Services.Config) redis.Conn {
	conn, _ := redis.Dial("tcp", "172.20.3.234:6379")
	return conn
}

func RedisClose(c *Services.Config) {
	_ = RedisPollInit(c).Get().Close()
}
