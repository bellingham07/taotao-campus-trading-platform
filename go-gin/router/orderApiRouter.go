package router

import (
	"com.xpwk/go-gin/api"
	"com.xpwk/go-gin/router/middleware"
	"github.com/gin-gonic/gin"
)

type OrderRouter struct {
}

func (*OrderRouter) InitOrderApiRouter(g *gin.RouterGroup) {
	orderApi := new(api.OrderApi)
	og := g.Group("/order", middleware.JWTAuthenticate())
	{
		og.GET("/list/:id", orderApi.InfoApi.List)
	}
}
