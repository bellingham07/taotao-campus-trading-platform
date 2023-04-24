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
	g.GET("/avatar", middleware.JWTAuthenticate, middleware.FileCheck, fileApi.InfoApi.Upload)
}
