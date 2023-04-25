package router

import (
	"com.xpwk/go-gin/api"
	"com.xpwk/go-gin/router/middleware"
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
