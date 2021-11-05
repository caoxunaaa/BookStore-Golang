package user

import (
	"WebApi/Pb/user"
	"WebApi/Svc"
	"context"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetAllUsersHandler(c *gin.Context) {
	rep, err := Svc.SvcContext.Grpc.UserGrpc.FindAllUser(context.Background(), &user.Request{})
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, rep)
}

func GetUserHandler(c *gin.Context) {
	var username string
	name, ok := c.Get("username")
	if !ok {
		c.JSON(http.StatusNotFound, gin.H{"error": "username not exists"})
		return
	} else {
		username = name.(string)
	}
	rep, err := Svc.SvcContext.Grpc.UserGrpc.FindOneUserByUsername(context.Background(), &user.UsernameReq{Username: username})
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, rep)
}
