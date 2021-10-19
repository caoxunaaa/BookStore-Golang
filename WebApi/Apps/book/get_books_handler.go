package book

import (
	"WebApi/Pb/book"
	"WebApi/Services"
	"context"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func GetAllBooksHandler(c *gin.Context) {
	reps, err := Services.Grpc.BookGrpc.FindAllBooks(context.Background(), &book.Request{})
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, reps)
}

func GetMyBooksHandler(c *gin.Context) {
	storageUserId, err := strconv.ParseInt(c.Query("storageUserId"), 10, 64)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	reps, err := Services.Grpc.BookGrpc.FindBooksByStorageUserId(context.Background(), &book.BookBasicInfoReq{StorageUserId: storageUserId})
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, reps)
}
