package router

import (
	"com.xpdj/go-gin/api"
	"com.xpdj/go-gin/router/middleware"
	"github.com/gin-gonic/gin"
)

type CommodityRouter struct {
}

func (*CommodityRouter) InitCommodityApiRouter(g *gin.RouterGroup) {
	commodityApi := api.SystemApis.CommodityApi

	g.GET("/:id", commodityApi.InfoApi.GetInfoById)
	g.GET("/list/:option", commodityApi.InfoApi.ListByOption)

	ig := g.Group("", middleware.JWTAuthenticate)
	{
		ig.POST("/sellsave", commodityApi.InfoApi.SellSave)
		ig.POST("/sellpublish", commodityApi.InfoApi.SellPublish)
		ig.POST("/wantsave", commodityApi.InfoApi.WantSave)
		ig.POST("/wantpublish", commodityApi.InfoApi.WantPublish)
	}

	// 足迹
	hg := g.Group("/history", middleware.JWTAuthenticate)
	{
		hg.GET("", commodityApi.HistoryApi.List)
		hg.DELETE("", commodityApi.HistoryApi.Delete)
	}

	// 商品收藏
	cog := g.Group("/collect", middleware.JWTAuthenticate)
	{
		cog.GET("/:id", commodityApi.CollectApi.Collect)
		cog.GET("/list", commodityApi.CollectApi.List)
		cog.DELETE("/:id", commodityApi.CollectApi.Uncollect)
	}

	// 商品标签
	cag := g.Group("/tag")
	{
		cag.GET("", commodityApi.TagApi.List)
		cag.POST("", commodityApi.TagApi.Add)
		cag.DELETE("/:id", commodityApi.TagApi.Remove)
	}

}
