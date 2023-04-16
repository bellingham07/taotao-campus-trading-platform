package router

import (
	"com.xpwk/go-gin/api"
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

}
