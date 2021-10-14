package Router

import (
	"WebApi/Apps/book"
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
		userGroup.GET("/", Middlewares.JWTSuperuserMiddleware(), user.GetAllUsersHandler)
	}
	bookGroup := r.Group("/book", Middlewares.JWTSuperuserMiddleware())
	{
		bookGroup.POST("/", book.CreateBookHandler)
		bookGroup.GET("/", book.GetAllBooksHandler)
	}

	return r
}
