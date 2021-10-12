package Router

import (
	"WebApi/Apps/user"
	"WebApi/Middlewares"
	"github.com/gin-gonic/gin"
)

func Init() *gin.Engine {
	r := gin.Default()
	//r. Use(Middlewares.Cors())
	r.Static("/Assets", "./Assets")
	r.StaticFile("/favicon.ico", "./Assets/favicon.ico")

	userGroup := r.Group("/user")
	{
		userGroup.POST("/login", user.LoginHandler)
		userGroup.POST("/register", user.RegisterHandler)
		userGroup.GET("/",Middlewares.JWTSuperuserMiddleware(), user.GetAllUsersHandler)
	}

	return r
}
