package chatRepository

import (
	"com.xpdj/go-gin/model"
	"com.xpdj/go-gin/repository"
)

var ChatRoom = new(ChatRoomRepository)

type ChatRoomRepository struct {
}

func chat_room() string {
	return "chat_room"
}

func (*ChatRoomRepository) Insert(cr *model.ChatRoom) error {
	if err := repository.GetDB().Table(chat_room()).Create(cr).Error; err != nil {
		return err
	}
	return nil
}

func (*ChatRoomRepository) QueryByCidSidBid(cid, sid, bid int64) *model.ChatRoom {
	crDB := new(model.ChatRoom)
	if err := repository.GetDB().Table(chat_room()).
		Where("commodity_id = ? AND seller_id = ? AND buyer_id = ?", cid, sid, bid).
		First(crDB).Error; err != nil {
		return nil
	}
	return crDB
}

func (*ChatRoomRepository) QueryById(cr *model.ChatRoom) error {
	if err := repository.GetDB().Table(chat_room()).First(cr).Error; err != nil {
		return err
	}
	return nil
}
