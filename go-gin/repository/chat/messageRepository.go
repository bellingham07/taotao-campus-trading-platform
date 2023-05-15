package chatRepository

import (
	"com.xpdj/go-gin/model"
	"com.xpdj/go-gin/repository"
)

var ChatMessage = new(ChatMessageRepository)

type ChatMessageRepository struct {
}

func chat_message() string {
	return "chat_message"
}

func (*ChatMessageRepository) ListByRoomId(roomId int64, offset int) []*model.ChatMessage {
	cms := make([]*model.ChatMessage, 20)
	if err := repository.GetDB().Table(chat_message()).
		Where("room_id = ?", roomId).Find(&cms).
		Limit(20).Offset(offset).
		Order("time desc").Error; err != nil {
		return nil
	}
	return cms
}

func (*ChatMessageRepository) Insert(msg *model.ChatMessage) error {
	if err := repository.GetDB().Table(chat_message()).Create(msg).Error; err != nil {
		return err
	}
	return nil
}
