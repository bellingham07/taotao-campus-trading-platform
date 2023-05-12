package router

import (
	"com.xpdj/go-gin/api"
	"com.xpdj/go-gin/router/middleware"
	"github.com/gin-gonic/gin"
)

type ArticleRouter struct {
}

func (*ArticleRouter) InitArticleApiRouter(g *gin.RouterGroup) {
	articleApi := api.SystemApis.ArticleApi

	g.GET("/:id", articleApi.ContentApi.GetById)
	ag := g.Group("", middleware.JWTAuthenticate)
	{
		ag.POST("/save", articleApi.ContentApi.Save)
		ag.POST("/publish", articleApi.ContentApi.Publish)
	}
}
