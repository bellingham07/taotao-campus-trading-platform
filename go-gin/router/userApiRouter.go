package router

import (
	"com.xpwk/go-gin/api"
	"github.com/gin-gonic/gin"
)

type UserRouter struct {
}

func (*UserRouter) InitUserApiRouter(g *gin.RouterGroup) {
	userApi := api.SystemApis.UserApi
	ug := g.Group("/user")
	{
		ug.POST("/login", userApi.UserInfoApi.UserLogin)
		ug.GET("/info/:id", userApi.UserInfoApi.GetUserInfo)
		ug.POST("/info", userApi.UserInfoApi.UpdateUserInfo)

		ulg := ug.Group("/location")
		{
			ulg.GET("/list", userApi.UserLocationApi.ListLocations)
		}
	}
}
