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

	g.POST("/login", userApi.UserInfoApi.UserLogin)
	g.GET("/logout", middleware.JWTAuthenticate(), userApi.UserInfoApi.Logout)

	ig := g.Group("/info")
	{
		ig.GET("/:id", middleware.JWTAuthenticate(), userApi.UserInfoApi.GetInfoById)
		ig.POST("/", userApi.UserInfoApi.UpdateInfo)
	}

	lg := g.Group("/location")
	{
		lg.GET("/list", userApi.UserLocationApi.List)
	}

}
