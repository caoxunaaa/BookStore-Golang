package action

import (
	"WebApi/Pb/action"
	"WebApi/Svc"
	"context"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func CreateCommentHandler(c *gin.Context) {
	parentId, err := strconv.ParseInt(c.PostForm("parentId"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	bookContentId, err := strconv.ParseInt(c.PostForm("bookContentId"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	comment := c.PostForm("comment")
	commentToUserId, err := strconv.ParseInt(c.PostForm("commentToUserId"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	commentToNickname := c.PostForm("commentToNickname")
	commentByUserId, err := strconv.ParseInt(c.PostForm("commentByUserId"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	commentByNickname := c.PostForm("commentByNickname")

	ctx := context.Background()
	res, err := Svc.SvcContext.Grpc.ActionGrpc.CreateComment(ctx, &action.CommentReq{
		ParentId:          parentId,
		BookContentId:     bookContentId,
		Comment:           comment,
		CommentToUserId:   commentToUserId,
		CommentToNickname: commentToNickname,
		CommentByUserId:   commentByUserId,
		CommentByNickname: commentByNickname,
	})
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	} else {
		c.JSON(http.StatusOK, res)
	}
}
