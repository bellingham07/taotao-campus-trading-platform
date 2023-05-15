package router

import (
	"com.xpdj/go-gin/api"
	"com.xpdj/go-gin/router/middleware"
	"github.com/gin-gonic/gin"
)

type ChatRouter struct {
}

func (*ChatRouter) InitChatApiRouter(g *gin.RouterGroup) {
	chatApi := api.SystemApis.ChatApi

	rg := g.Group("/room", middleware.JWTAuthenticate)
	{
		rg.POST("", chatApi.RoomApi.CreateRoom)
		rg.GET("/:id", chatApi.RoomApi.Chat)

	}

	g = g.Group("/msg", middleware.JWTAuthenticate)
	{
		g.GET("", chatApi.MessageApi.ListMsgsByRoomId)
	}

}
