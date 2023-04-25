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
		ig.GET("/list/:option", commodityApi.InfoApi.ListByOption)
	}

	hg := g.Group("/history", middleware.JWTAuthenticate)
	{
		hg.GET("", commodityApi.HistoryApi.List)
		hg.DELETE("", commodityApi.HistoryApi.Delete)
	}

	cog := g.Group("/collect", middleware.JWTAuthenticate)
	{
		cog.GET("/:id", commodityApi.CollectApi.Collect)
		cog.GET("/list", commodityApi.CollectApi.List)
		cog.DELETE("/:id", commodityApi.CollectApi.Uncollect)
	}

	cag := g.Group("/category")
	{
		cag.GET("", commodityApi.CategoryApi.List)
	}

}
