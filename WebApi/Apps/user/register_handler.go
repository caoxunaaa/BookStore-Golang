package user

import (
	"WebApi/Pb/user"
	"WebApi/Services"
	"context"
	"github.com/gin-gonic/gin"
	"net/http"
)

func RegisterHandler(c *gin.Context) {
	username := c.DefaultPostForm("username", "")
	password := c.DefaultPostForm("password", "")
	nickname := c.DefaultPostForm("nickname", "")
	pwAgain := c.DefaultPostForm("pwAgain", "")
	email := c.DefaultPostForm("email", "")
	phone := c.DefaultPostForm("phone", "")

	rep, err := Services.Grpc.UserGrpc.Register(context.Background(), &user.RegisterReq{
		Username:       username,
		Password:       password,
		Nickname:       nickname,
		Email:          email,
		Phone:          phone,
		RepeatPassword: pwAgain,
	})
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	} else {
		if rep.Ok == false {
			c.JSON(http.StatusBadRequest, gin.H{"error": rep.Code})
			return
		}
	}
	c.JSON(http.StatusOK, rep)
}
