package router

import (
	"com.xpdj/go-gin/api"
	"com.xpdj/go-gin/router/middleware"
	"github.com/gin-gonic/gin"
)

type FileRouter struct {
}

func (*FileRouter) InitFileApiRouter(g *gin.RouterGroup) {
	fileApi := api.SystemApis.FileApi

	fg := g.Group("", middleware.JWTAuthenticate, middleware.FileCheck)
	{
		fg.POST("/avatar", fileApi.InfoApi.UploadAvatar)
		fg.POST("/pics", fileApi.InfoApi.UploadPics)
		fg.DELETE("", fileApi.InfoApi.Delete)
	}

}
