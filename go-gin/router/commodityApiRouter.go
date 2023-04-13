package router

import (
	"com.xpwk/go-gin/api"
	"github.com/gin-gonic/gin"
)

type CommodityRouter struct {
}

func (*CommodityRouter) InitCommodityApiRouter(g *gin.RouterGroup) {
	commodityApi := api.SystemApis.CommodityApi
	g.Group("/cmdty")
	{

		hg := g.Group("/history")
		{
			hg.GET("/:userid", commodityApi.CommodityHistoryApi.List)
		}
	}
}
