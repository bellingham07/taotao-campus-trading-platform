package router

import (
	"com.xpwk/go-gin/router/middleware"
	"github.com/gin-gonic/gin"
)

func Routers() *gin.Engine {
	router := gin.New()

	myRouter := new(SystemRouterGroup)

	groupRegistry := router.Group("/", middleware.Cors())
	{
		myRouter.UserRouter.InitUserApiRouter(groupRegistry)
		myRouter.CommodityRouter.InitCommodityApiRouter(groupRegistry)
		myRouter.FileRouter.InitFileApiRouter(groupRegistry)
		myRouter.OrderRouter.InitOrderApiRouter(groupRegistry)
		myRouter.MessageRouter.InitMessageApiRouter(groupRegistry)
		myRouter.ArticleRouter.InitArticleApiRouter(groupRegistry)
	}

	return router
}
