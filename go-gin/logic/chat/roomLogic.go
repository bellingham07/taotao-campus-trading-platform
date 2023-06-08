package chatLogic

import (
	"com.xpdj/go-gin/model"
	"com.xpdj/go-gin/model/response"
	chatRepository "com.xpdj/go-gin/repository/chat"
	"github.com/gin-gonic/gin"
	"time"
)

var ChatRoom = new(ChatRoomLogic)

type ChatRoomLogic struct {
}

func (*ChatRoomLogic) CreateRoom(cr *model.ChatRoom) gin.H {
	cr.CreateAt = time.Now()
	if err := chatRepository.ChatRoom.Insert(cr); err != nil {
		// 创建失败，则以前已经创建
		// 那就查询出room的id
		cr = chatRepository.ChatRoom.QueryByCidSidBid(cr.CommodityId, cr.SellerId, cr.BuyerId)
	}
	return response.OkData(cr)
}

func (*ChatRoomLogic) GetById(id int64) {

}
