package book

import (
	"WebApi/Pb/book"
	"WebApi/Services"
	"WebApi/Utils"
	"context"
	"github.com/gin-gonic/gin"
	"net/http"
)

func CreateBookHandler(c *gin.Context) {
	name := c.PostForm("name")
	author := c.PostForm("author")
	storageTime := c.PostForm("storageTime")
	_, image, err := Utils.UploadFile(c.Request, "image")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	rep, err := Services.Grpc.BookGrpc.CreateBook(context.Background(), &book.BookBasicInfoReq{
		Name:        name,
		Author:      author,
		Image:       image,
		StorageTime: storageTime,
	})
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, rep)
}
