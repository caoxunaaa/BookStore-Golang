package book

import (
	"WebApi/Pb/book"
	"WebApi/Svc"
	"context"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func GetBookInventoryByBookIdHandler(c *gin.Context) {
	bookId, err := strconv.ParseInt(c.Param("bookId"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx := context.Background()
	res, err := Svc.SvcContext.Grpc.BookGrpc.FindBookInventoryByBookId(ctx, &book.BookInventoryReq{BookId: bookId})
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, res)
}

func CreateBookInventoryHandler(c *gin.Context) {
	bookId, err := strconv.ParseInt(c.PostForm("bookId"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	inventory, err := strconv.ParseInt(c.PostForm("inventory"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx := context.Background()
	res, err := Svc.SvcContext.Grpc.BookGrpc.CreateBookInventory(ctx, &book.BookInventoryReq{BookId: bookId, Inventory: inventory})
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, res)
}

func UpdateBookInventoryHandler(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	bookId, err := strconv.ParseInt(c.PostForm("bookId"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	inventory, err := strconv.ParseInt(c.PostForm("inventory"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx := context.Background()
	res, err := Svc.SvcContext.Grpc.BookGrpc.UpdateBookInventory(ctx, &book.BookInventoryReq{Id: id, BookId: bookId, Inventory: inventory})
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	_, err = Svc.SvcContext.Redis.Do("SET", "Inventory:BookId:"+strconv.FormatInt(bookId, 10), inventory)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, res)
}
