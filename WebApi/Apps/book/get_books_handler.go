package book

import (
	"WebApi/Pb/book"
	"WebApi/Services"
	"context"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetAllBooksHandler(c *gin.Context) {
	reps, err := Services.Grpc.BookGrpc.FindAllBooks(context.Background(), &book.Request{})
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, reps)
}
