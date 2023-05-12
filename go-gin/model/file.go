package model

import "time"

// FileCommodity 商品信息相关图片
type FileCommodity struct {
	Id          int64     `json:"id"` // bigint自增
	CommodityId int64     `json:"commodityId"`
	UserId      int64     `json:"userId"`
	Url         string    `json:"url"`
	ObjectName  string    `json:"objectName"`
	IsCover     int64     `json:"isCover"` // 默认为null，封面为1
	CreateAt    time.Time `json:"createAt"`
}

// FileArticle 文章内容相关图片
type FileArticle struct {
	Id         int64     `json:"id"` // bigint自增
	ArticleId  int64     `json:"articleId"`
	UserId     int64     `json:"userId"`
	Url        string    `json:"url"`
	ObjectName string    `json:"objectName"`
	IsCover    int64     `json:"isCover"` // 默认为null，封面为1
	CreateAt   time.Time `json:"createAt"`
}
