package chatApi

import (
	chatLogic "com.xpdj/go-gin/logic/chat"
	"com.xpdj/go-gin/model/response"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type MessageApi struct {
}

func (*MessageApi) ListMsgsByRoomId(c *gin.Context) {
	roomIdStr := c.Query("roomId")
	roomId, err := strconv.ParseInt(roomIdStr, 10, 64)
	offsetStr := c.Query("offset")
	offset, err2 := strconv.ParseInt(offsetStr, 10, 64)
	if err != nil || err2 != nil {
		c.JSON(http.StatusBadRequest, response.ErrorMsg("请求参数错误！🤡"))
		return
	}
	c.JSON(http.StatusOK, chatLogic.ChatMessage.ListByRoomId(roomId, offset))
}
