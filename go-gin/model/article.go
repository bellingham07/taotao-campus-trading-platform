package model

import "time"

type ArticleContent struct {
	Id       int64     `json:"id"`      // bigint自增，id
	UserId   int64     `json:"userId"`  // uuid，用户的id
	Content  string    `json:"content"` // 帖子内容
	Status   int64     `json:"status"`  // 1为草稿，2为发布，-1为审核不通过
	Title    string    `json:"title"`
	Type     int64     `json:"type"`     // 1为商品介绍文章，2为文章
	UpdateAt time.Time `json:"updateAt"` // datetime，更新时间
	CreateAt time.Time `json:"createAt"` // datetime，创建时间
}

type ArticleComment struct {
	Id        int64  `json:"id"`        // bigint自增，id
	ArticleId string `json:"articleId"` // bigint自增
	Content   string `json:"content"`   // 留言内容
	RootId    int64  `json:"rootId"`    // 根留言
	ToUserId  int64  `json:"toUserId"`  // 雪花算法，回复给到用户
	UserId    int64  `json:"userId"`    // 雪花算法，留言的用户id
	CreateAt  string `json:"createAt"`  // 评论时间
}

type ArticleBulletin struct {
	Id       int64  `json:"id"`       // 自增，id
	AdminId  string `json:"adminId"`  // uuid，发布管理员id
	Content  string `json:"content"`  // 内容
	CreateAt string `json:"createAt"` // datetime，发布时间
	UpdateAt string `json:"updateAt"` // datetime，更新时间
}
