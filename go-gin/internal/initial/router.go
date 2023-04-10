package initial

import "github.com/gin-gonic/gin"
import myRtr "com.xpwk/go-gin/internal/router"

func Routers() *gin.Engine {
	router := gin.New()

	myRouter := new(myRtr.SystemGroup)

	groupRegistry := router.Group("/")
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
