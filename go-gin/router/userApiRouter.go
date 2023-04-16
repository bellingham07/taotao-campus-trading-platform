package router

import (
	"com.xpwk/go-gin/api"
	"com.xpwk/go-gin/router/middleware"
	"github.com/gin-gonic/gin"
)

type UserRouter struct {
}

func (*UserRouter) InitUserApiRouter(g *gin.RouterGroup) {
	userApi := api.SystemApis.UserApi

	g.POST("/login", userApi.InfoApi.UserLogin)
	g.GET("/logout", middleware.JWTAuthenticate(), userApi.InfoApi.Logout)

	ig := g.Group("/info")
	{
		ig.GET("/:id", middleware.JWTAuthenticate(), userApi.InfoApi.GetInfoById)
		ig.POST("/", userApi.InfoApi.UpdateInfo)
	}

	lg := g.Group("/location")
	{
		lg.GET("/list", userApi.LocationApi.List)
	}

}
