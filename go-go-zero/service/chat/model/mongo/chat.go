package mongo

import "time"

// ChatMessage mongo
type ChatMessage struct {
	Id      int64     `json:"id"` // bigint自增
	Content string    `json:"content"`
	RoomId  int64     `json:"room_id"`
	Time    time.Time `json:"time"`
	UserId  int64     `json:"user_id"`
}
