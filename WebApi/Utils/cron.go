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

	//凌晨5点更新DB中书籍的访问量
	if err := c.AddFunc("0 0 5 * * ?", func() {
		_ = TrafficStatisticsImportDB()
	}); err != nil {
		fmt.Println(err)
	}
	fmt.Println("cron start")
	c.Start()
}

//定时在缓存和数据库之间进行访问量的同步
func TrafficStatisticsImportDB() error {
	var err error
	//找到所有redis访问量的内容
	res, err := redis.StringMap(Svc.SvcContext.Redis.Do("ZRANGE", "traffic_statistic", 0, -1, "WITHSCORES"))
	if err != nil {
		return err
	}
	fmt.Println(res)
	// Update to DB
	var bookId, chapterNum, trafficNumber int64
	for key, _ := range res {
		bookId, err = strconv.ParseInt(strings.Split(key, ":")[0], 10, 64)
		if err != nil {
			return err
		}

		chapterNum, err = strconv.ParseInt(strings.Split(key, ":")[1], 10, 64)
		if err != nil {
			return err
		}

		trafficNumber, err = strconv.ParseInt(res[key], 10, 64)
		if err != nil {
			return err
		}

		rep, err := Svc.SvcContext.Grpc.ActionGrpc.GetTrafficStatisticByBookIdAndChapterNum(context.Background(),
			&action.TrafficStatisticReq{
				BookId:     bookId,
				ChapterNum: chapterNum})
		fmt.Println(rep, err)

		if err != nil {
			if err.Error() == "rpc error: code = Unknown desc = sql: no rows in result set" {
				_, err := Svc.SvcContext.Grpc.ActionGrpc.CreateTrafficStatistic(context.Background(), &action.TrafficStatisticReq{
					BookId:        bookId,
					ChapterNum:    chapterNum,
					TrafficNumber: trafficNumber,
				})
				if err != nil {
					return err
				}
			} else {
				return err
			}
		} else {
			if trafficNumber > rep.TrafficNumber {
				_, err := Svc.SvcContext.Grpc.ActionGrpc.UpdateTrafficStatistic(context.Background(), &action.TrafficStatisticReq{
					Id:            rep.Id,
					BookId:        bookId,
					ChapterNum:    chapterNum,
					TrafficNumber: trafficNumber,
				})
				if err != nil {
					return err
				}
			}
		}
	}

	//DownLoad to Redis(防止Redis数据丢失)
	if len(res) == 0 || res == nil {
		resp, err := Svc.SvcContext.Grpc.ActionGrpc.GetAllTrafficStatistics(context.Background(), &action.Request{})
		if err != nil {
			return err
		}
		ts := resp.TrafficStatistics
		key := "traffic_statistic"
		for i, _ := range ts {
			_, err = Svc.SvcContext.Redis.Do("ZADD", key, ts[i].TrafficNumber, strconv.FormatInt(ts[i].BookId, 10)+":"+strconv.FormatInt(ts[i].ChapterNum, 10))
			if err != nil {
				return err
			}
		}
	}
	return nil
}
