package model

import "time"

type (
	AtclContent struct {
		Id       int64     `xorm:"id" json:"id"`          // id
		UserId   int64     `xorm:"user_id" json:"userId"` // 用户的id
		Title    string    `xorm:"title" json:"title"`
		Content  string    `xorm:"content" json:"content"`    // 帖子内容
		Cover    string    `xorm:"cover" json:"cover"`        // 封面
		Status   int64     `xorm:"status" json:"status"`      // 1为草稿，2为发布，-1为审核不通过
		CreateAt time.Time `xorm:"create_at" json:"createAt"` // 创建时间
		UpdateAt time.Time `xorm:"update_at" json:"updateAt"` // 更新时间
		Collect  int64     `xorm:"collect" json:"collect"`
		Like     int64     `xorm:"like" json:"like"`
		View     int64     `xorm:"view" json:"view"`
	}

	AtclCmt struct {
		Id       int64     `bson:"id" json:"id"` // id
		AtclId   int64     `bson:"article_id" json:"articleId"`
		UserId   int64     `bson:"user_id" json:"userId"`      // 留言的用户id
		Content  string    `bson:"content" json:"content"`     // 留言内容
		RootId   int64     `bson:"root_id" json:"rootId"`      // 根留言
		ToUserId int64     `bson:"to_user_id" json:"toUserId"` // 回复给到用户
		CreateAt time.Time `bson:"create_at" json:"createAt"`  // 评论时间
	}

	AtclCollect struct {
		Id       int64     `xorm:"id" json:"id"`                // id
		UserId   int64     `xorm:"user_id" json:"userId"`       // 用户id
		AtclId   int64     `xorm:"article_id" json:"articleId"` // 商品id
		Status   int64     `xorm:"status" json:"status"`        // 1存在，0失效
		CreateAt time.Time `xorm:"create_at" json:"createAt"`   // 创建时间
	}

	AtclBulletin struct {
		Id       int64     `xorm:"id" json:"id"`              // id
		Content  string    `xorm:"content" json:"content"`    // 内容
		AdminId  int64     `xorm:"admin_id" json:"adminId"`   // 发布管理员id
		CreateAt time.Time `xorm:"create_at" json:"createAt"` // 发布时间
		UpdateAt time.Time `xorm:"update_at" json:"updateAt"` // 更新时间
	}
)
