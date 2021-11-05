package book

import (
	"WebApi/Pb/book"
	"WebApi/Svc"
	"context"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func GetBookByIdHandler(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	reps, err := Svc.SvcContext.Grpc.BookGrpc.FindOneBookById(context.Background(), &book.BookBasicInfoReq{Id: id})
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, reps)
}

func GetAllBooksHandler(c *gin.Context) {
	reps, err := Svc.SvcContext.Grpc.BookGrpc.FindAllBooks(context.Background(), &book.Request{})
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, reps)
}

func GetMyselfBooksHandler(c *gin.Context) {
	storageUserId, err := strconv.ParseInt(c.Param("storageUserId"), 10, 64)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	reps, err := Svc.SvcContext.Grpc.BookGrpc.FindBooksByStorageUserId(context.Background(), &book.BookBasicInfoReq{StorageUserId: storageUserId})
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, reps)
}
