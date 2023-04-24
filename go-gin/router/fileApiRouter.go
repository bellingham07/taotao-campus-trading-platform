package router

import (
	fileApi "com.xpwk/go-gin/api/file"
	"com.xpwk/go-gin/router/middleware"
	"github.com/gin-gonic/gin"
)

type FileRouter struct {
}

func (*FileRouter) InitFileApiRouter(g *gin.RouterGroup) {
	g.GET("/avatar", middleware.JWTAuthenticate(), fileApi.InfoApi.Upload)
}
