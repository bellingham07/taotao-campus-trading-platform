package router

import (
	"com.xpwk/go-gin/api"
	"com.xpwk/go-gin/router/middleware"
	"github.com/gin-gonic/gin"
)

type CommodityRouter struct {
}

func (*CommodityRouter) InitCommodityApiRouter(g *gin.RouterGroup) {
	commodityApi := api.SystemApis.CommodityApi

	ig := g.Group("/info")
	{
		ig.GET("/:id", commodityApi.InfoApi.GetInfoById)
	}

	hg := g.Group("/history")
	{
		hg.GET("/:userid", commodityApi.HistoryApi.List)
	}

	cog := g.Group("/collect", middleware.JWTAuthenticate())
	{
		cog.PUT("/:id", commodityApi.CollectApi.Collect)
		cog.GET("/collect", commodityApi.CollectApi.List)
		cog.DELETE("/:id", commodityApi.CollectApi.Uncollect)
	}

	cag := g.Group("/category")
	{
		cag.GET("/", commodityApi.CategoryApi.List)
	}

}
