package Router

import (
	"WebApi/Apps/user"
	"WebApi/Middlewares"
	"github.com/gin-gonic/gin"
	"time"
)

func Init() *gin.Engine {
	r := gin.Default()

	r.Static("/Assets", "./Assets")
	r.StaticFile("/favicon.ico", "./Assets/favicon.ico")
	r.Use(Middlewares.TimeoutMiddleware(time.Minute * 30))

	userGroup := r.Group("/user")
	{
		userGroup.POST("/login", user.LoginHandler)
	}
	return r
}
