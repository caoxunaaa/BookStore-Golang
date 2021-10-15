package book

import (
	"WebApi/Pb/book"
	"WebApi/Services"
	"WebApi/Utils"
	"context"
	"github.com/gin-gonic/gin"
	"net/http"
)

func CreateBookContentHandler(c *gin.Context) {
	name := c.PostForm("name")
	author := c.PostForm("author")
	storageTime := c.PostForm("storageTime")
	_, image := Utils.UploadFile(c.Request, "image")
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
