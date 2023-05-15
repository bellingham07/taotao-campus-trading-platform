package chatLogic

import (
	"com.xpdj/go-gin/model/response"
	chatRepository "com.xpdj/go-gin/repository/chat"
	"github.com/gin-gonic/gin"
)

var ChatMessage = new(ChatMessageLogic)

type ChatMessageLogic struct {
}

func (*ChatMessageLogic) ListByRoomId(roomId, offset int64) gin.H {
	cms := chatRepository.ChatMessage.ListByRoomId(roomId, int(offset))
	return response.OkData(cms)
}
