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
	og := g.Group("/order", middleware.JWTAuthenticate)
	{
		og.GET("/list/:id", orderApi.InfoApi.List)
	}
}
