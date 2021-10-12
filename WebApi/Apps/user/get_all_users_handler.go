package user

import (
	"WebApi/Pb/user"
	"WebApi/Services"
	"context"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetAllUsersHandler(c *gin.Context) {
	rep, err := Services.UserGrpc.FindAllUser(context.Background(), &user.Request{})
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, rep)
}