package Utils

import (
	"WebApi/Pb/action"
	"WebApi/Svc"
	"context"
	"fmt"
	"github.com/gomodule/redigo/redis"
	"github.com/robfig/cron"
	"strconv"
	"strings"
)

func init() {
	c := cron.New()

	if err := c.AddFunc("0 50 16 * * ?", func() {
		_ = TrafficStatisticsImportDB()
	}); err != nil {
		fmt.Println(err)
	}
	fmt.Println("cron start")
	c.Start()
}

//定时在缓存和数据库之间进行访问量的同步
func TrafficStatisticsImportDB() error {
	//找到所有redis访问量的内容
	res, err := redis.ByteSlices(Svc.SvcContext.Redis.Do("Keys", "traffic_statistic:*"))
	if err != nil {
		return err
	}
	// Update to DB
	for i := 0; i < len(res); i++ {
		bookId, err := strconv.ParseInt(strings.Split(string(res[i]), ":")[1], 10, 64)
		if err != nil {
			return err
		}
		chapterNum, err := strconv.ParseInt(strings.Split(string(res[i]), ":")[2], 10, 64)
		if err != nil {
			return err
		}
		trafficNumber, err := redis.Int64(Svc.SvcContext.Redis.Do("GET", string(res[i])))
		if err != nil {
			return err
		}
		_, err = Svc.SvcContext.Grpc.ActionGrpc.CreateTrafficStatistic(context.Background(), &action.TrafficStatisticReq{
			BookId:        bookId,
			ChapterNum:    chapterNum,
			TrafficNumber: trafficNumber,
		})
		if err != nil {
			return err
		}
	}

	//DownLoad to Redis(Redis数据丢失)
	if len(res) == 0 {
		resp, err := Svc.SvcContext.Grpc.ActionGrpc.GetAllTrafficStatistics(context.Background(), &action.Request{})
		if err != nil {
			return err
		}
		ts := resp.TrafficStatistics
		for i, _ := range ts {
			key := "traffic_statistic:" + strconv.FormatInt(ts[i].BookId, 10) + ":" + strconv.FormatInt(ts[i].ChapterNum, 10)
			_, err = Svc.SvcContext.Redis.Do("SET", key, ts[i].TrafficNumber)
			if err != nil {
				return err
			}
		}
	}

	return nil
}
