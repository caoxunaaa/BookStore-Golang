package action

import (
	"WebApi/Svc"
	"github.com/gin-gonic/gin"
	"github.com/gomodule/redigo/redis"
	"net/http"
	"strconv"
	"strings"
)

func GetTrafficStatisticByBookIdAndChapterNumHandler(c *gin.Context) {
	bookId := c.Query("bookId")
	chapterNum := c.Query("chapterNum")

	//找到redis访问量的内容
	key := "traffic_statistic"
	res, err := redis.String(Svc.SvcContext.Redis.Get().Do("ZSCORE", key, bookId+":"+chapterNum))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	v, err := strconv.Atoi(res)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, v)
}

func GetAllTrafficStatisticHandler(c *gin.Context) {
	//找到redis访问量的内容
	key := "traffic_statistic"
	res, err := redis.StringMap(Svc.SvcContext.Redis.Get().Do("ZRANGE", key, 0, -1, "WITHSCORES"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, res)
}

func GetAllTrafficStatisticHandlerByBookId(c *gin.Context) {
	//通过书籍ID找到redis访问量的内容
	bookId := c.Query("bookId")
	key := "traffic_statistic"
	res, err := redis.StringMap(Svc.SvcContext.Redis.Get().Do("ZRANGE", key, 0, -1, "WITHSCORES"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	var count int64
	for k, _ := range res {
		if strings.Split(k, ":")[0] == bookId {
			n, err := strconv.ParseInt(res[k], 10, 64)
			if err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
				return
			}
			count += n
		}
	}

	c.JSON(http.StatusOK, count)
}
