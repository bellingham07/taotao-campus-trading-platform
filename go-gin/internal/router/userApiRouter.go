package router

import (
	"com.xpwk/go-gin/internal/api"
	"github.com/gin-gonic/gin"
)

type UserRouter struct {
}

func (*UserRouter) InitUserApiRouter(g *gin.RouterGroup) {
	userApi := api.SystemApis.UserApi
	ug := g.Group("/user")
	{
		ug.POST("/login", userApi.UserLogin)
		ug.GET("/info", userApi.GetInfo)
		ug.POST("/info", userApi.UpdateInfo)
	}
}
