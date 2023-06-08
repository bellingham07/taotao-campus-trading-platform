package model

import "time"

// ArticleContent 文章内容
type ArticleContent struct {
	Id       int64     `json:"id"`     // bigint自增，id
	UserId   int64     `json:"userId"` // uuid，用户的id
	Title    string    `json:"title"`
	Content  string    `json:"content"`  // 帖子内容
	Cover    string    `json:"cover"`    // 封面
	Status   int64     `json:"status"`   // 1为草稿，2为发布，-1为审核不通过
	CreateAt time.Time `json:"createAt"` // datetime，创建时间
	UpdateAt time.Time `json:"updateAt"` // datetime，更新时间
	Like     int64     `json:"like"`
	Collect  int64     `json:"collect"`
}

// ArticleAsset 文章中相关的图片
type ArticleAsset struct {
	Id        int64 `json:"id"` // bigint自增
	ArticleId int64 `json:"article_id"`
	AssetId   int64 `json:"asset_id"`
	UserId    int64 `json:"user_id"`
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
