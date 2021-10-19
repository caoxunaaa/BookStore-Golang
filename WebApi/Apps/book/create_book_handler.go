package book

import (
	"WebApi/Pb/book"
	"WebApi/Services"
	"WebApi/Utils"
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func CreateBookHandler(c *gin.Context) {
	name := c.PostForm("name")
	author := c.PostForm("author")
	storageTime := c.PostForm("storageTime")
	storageUserId, err := strconv.ParseInt(c.PostForm("storageUserId"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	_, image, _ := Utils.UploadFile(c.Request, "image")
	if image == "" {
		image = "Assets/无法加载.png"
	}
	fmt.Println(name, author, image, storageUserId, storageTime)
	rep, err := Services.Grpc.BookGrpc.CreateBook(context.Background(), &book.BookBasicInfoReq{
		Name:          name,
		Author:        author,
		Image:         image,
		StorageUserId: storageUserId,
		StorageTime:   storageTime,
	})
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, rep)
}
