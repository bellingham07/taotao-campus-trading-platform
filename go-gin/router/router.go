package router

import (
	"com.xpwk/go-gin/router/middleware"
	"github.com/gin-gonic/gin"
)

func Routers() *gin.Engine {
	router := gin.Default()

	myRouter := new(SystemRouterGroup)

	groupRegistry := router.Group("/", middleware.Cors())
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
