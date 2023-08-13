package router

import (
	"com.xpdj/go-gin/api"
	"com.xpdj/go-gin/router/middleware"
	"github.com/gin-gonic/gin"
)

type UserRouter struct {
}

func (*UserRouter) InitUserApiRouter(g *gin.RouterGroup) {
	userApi := api.SystemApis.UserApi

	g.POST("/login", userApi.InfoApi.UserLogin)
	g.POST("/code", userApi.InfoApi.GenerateLoginCode)
	g.GET("/logout", middleware.JWTAuthenticate, userApi.InfoApi.Logout)
	g.POST("/register", userApi.InfoApi.Register)

	ig := g.Group("/info", middleware.JWTAuthenticate)
	{
		ig.GET("/:id", userApi.InfoApi.GetInfoById)
		ig.POST("/", userApi.InfoApi.UpdateInfo)
	}

	lg := g.Group("/location")
	{
		lg.GET("/list", userApi.LocationApi.List)
	}

	cog := g.Group("/collect")
	{
		cog.GET("/:id")
		cog.DELETE("/:id")
		cog.GET("/list/:id")

	}
}
