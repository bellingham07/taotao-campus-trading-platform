package mongo

import "time"

// ChatMessage mongodb
type ChatMessage struct {
	Id      int64     `json:"id" bson:"_id"` // bigint自增
	Content string    `json:"content" bson:"content"`
	RoomId  int64     `json:"room_id" bson:"room_id"`
	Time    time.Time `json:"time" bson:"time"`
	UserId  int64     `json:"user_id" bson:"user_id"`
}
