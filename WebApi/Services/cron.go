package Services

import (
	"fmt"
	"github.com/robfig/cron"
)

func init() {
	c := cron.New()

	if err := c.AddFunc("*/10 * * * * ?", func() {
		fmt.Println("HelloWorld")
	}); err != nil {
		fmt.Println(err)
	}

	c.Start()
}

//定时在缓存和数据库之间进行访问量的同步dgdfgfd
func TrafficStatisticsImportDB() {

}
