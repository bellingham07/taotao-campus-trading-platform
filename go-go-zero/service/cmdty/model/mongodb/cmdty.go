package mongodb

import "time"

type CmdtyCmt struct {
	Id       int64      `bson:"_id"`        // id
	CmdtyId  int64      `bson:"cmdty_id"`   // 对应的商品id
	UserId   int64      `bson:"user_id"`    // 留言的用户id
	Content  string     `bson:"content"`    // 留言内容
	RootId   int64      `bson:"root_id"`    // 根留言
	ToUserId int64      `bson:"to_user_id"` // 回复给到用户
	CreateAt time.Time  `bson:"create_at"`  // 评论时间
	SubCmt   []CmdtyCmt `bson:"sub_cmt, -"`
}
