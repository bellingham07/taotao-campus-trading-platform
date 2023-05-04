package router

import (
	"com.xpdj/go-gin/router/middleware"
	"github.com/gin-gonic/gin"
)

type SystemRouterGroup struct {
	UserRouter
	CommodityRouter
	OrderRouter
	ArticleRouter
	FileRouter
	MessageRouter
}

func Routers() *gin.Engine {
	router := gin.Default()

	myRouter := new(SystemRouterGroup)

	groupRegistry := router.Group("", middleware.Cors())
	{
		myRouter.UserRouter.InitUserApiRouter(groupRegistry.Group("/user"))
		myRouter.CommodityRouter.InitCommodityApiRouter(groupRegistry.Group("/cmdty"))
		myRouter.FileRouter.InitFileApiRouter(groupRegistry.Group("/file"))
		myRouter.OrderRouter.InitOrderApiRouter(groupRegistry.Group("/order"))
		myRouter.MessageRouter.InitMessageApiRouter(groupRegistry.Group("/msg"))
		myRouter.ArticleRouter.InitArticleApiRouter(groupRegistry.Group("/atcl"))
	}

	return router
}
