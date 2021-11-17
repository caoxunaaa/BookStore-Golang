package Router

import (
	"WebApi/Apps/action"
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

	userGroup := r.Group("/user/")
	{
		userGroup.POST("/login", user.LoginHandler)
		userGroup.POST("/register", user.RegisterHandler)
		userGroup.GET("/", Middlewares.JWTSuperuserMiddleware(), user.GetAllUsersHandler)
	}
	bookGroup := r.Group("/book/", Middlewares.JWTAuthMiddleware())
	{
		bookGroup.POST("/", book.CreateBookHandler)
		bookGroup.GET("/", book.GetAllBooksHandler)
		bookGroup.GET("/id/:id", book.GetBookByIdHandler)
		bookGroup.GET("/username/:storageUserId", book.GetMyselfBooksHandler)
		content := bookGroup.Group("/content/")
		{
			content.GET("/", book.GetAllBookContentByBookIdHandler)
			content.GET("/chapterNum", Middlewares.TrafficStatisticsMiddleware(), book.GetOneBookContentByBookIdAndChapterNumHandler)
			content.POST("/", book.CreateBookContentHandler)
		}
	}
	actionGroup := r.Group("/action/")
	{
		trafficStatisticGroup := actionGroup.Group("/traffic-statistic/")
		{
			trafficStatisticGroup.GET("/by-bookId-and-chapterNum", action.GetTrafficStatisticByBookIdAndChapterNumHandler)
			trafficStatisticGroup.GET("/by-bookId", action.GetAllTrafficStatisticHandlerByBookId)
			trafficStatisticGroup.GET("/", action.GetAllTrafficStatisticHandler)
		}
		commentGroup := actionGroup.Group("/comment/")
		{
			commentGroup.GET("/by-book-content-id", action.GetCommentsByBookContentIdHandler)
		}
	}
	return r
}
