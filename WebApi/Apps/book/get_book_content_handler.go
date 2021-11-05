package book

import (
	"WebApi/Pb/book"
	"WebApi/Svc"
	"context"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

//某本书籍所有章节内容
func GetAllBookContentByBookIdHandler(c *gin.Context) {
	bookId, err := strconv.ParseInt(c.Query("bookId"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	reps, err := Svc.SvcContext.Grpc.BookGrpc.FindAllBookContentsByBookId(context.Background(), &book.BookContentReq{BookId: bookId})
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, reps)
}

//某本书籍某一章节内容
func GetOneBookContentByBookIdAndChapterNumHandler(c *gin.Context) {
	bookId, err := strconv.ParseInt(c.Query("bookId"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	chapterNumStr := c.Query("chapterNum")
	chapterNum, err := strconv.ParseInt(chapterNumStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	rep, err := Svc.SvcContext.Grpc.BookGrpc.FindOneBookContentByBookIdAndChapterNum(context.Background(), &book.BookContentReq{
		BookId:     bookId,
		ChapterNum: chapterNum,
	})
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, rep)

}
