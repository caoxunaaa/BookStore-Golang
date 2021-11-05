package book

import (
	"WebApi/Pb/book"
	"WebApi/Svc"
	"context"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func CreateBookContentHandler(c *gin.Context) {
	bookId, err := strconv.ParseInt(c.PostForm("bookId"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	chapterNum, err := strconv.ParseInt(c.PostForm("chapterNum"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ChapterName := c.PostForm("chapterName")
	ChapterContent := c.PostForm("chapterContent")
	CreateTime := c.PostForm("createTime")
	rep, err := Svc.SvcContext.Grpc.BookGrpc.CreateBookContent(context.Background(), &book.BookContentReq{
		BookId:         bookId,
		ChapterNum:     chapterNum,
		ChapterName:    ChapterName,
		ChapterContent: ChapterContent,
		CreateTime:     CreateTime,
	})
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, rep)
}
