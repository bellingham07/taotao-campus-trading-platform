package router

import (
	"com.xpdj/go-gin/api"
	"com.xpdj/go-gin/router/middleware"
	"github.com/gin-gonic/gin"
)

type OrderRouter struct {
}

func (*OrderRouter) InitOrderApiRouter(g *gin.RouterGroup) {
	orderApi := new(api.OrderApi)
	og := g.Group("/trade", middleware.JWTAuthenticate)
	{
		og.GET("/:id", orderApi.InfoApi.GetById)
		og.GET("/list/", orderApi.InfoApi.List)
		og.POST("/buy", orderApi.InfoApi.Buy)
		og.POST("/sell", orderApi.InfoApi.Sell)
		og.PUT("/cancel", orderApi.InfoApi.Cancel)
		og.PUT("/sell/:id", orderApi.InfoApi.SellConfirm)
		og.PUT("/:id/:gb", orderApi.InfoApi.Done)
	}
}
