package model

import "time"

type CommodityInfo struct {
	Id        int64     `json:"id"`     // bigint自增，id
	UserId    int64     `json:"userId"` // uuid，用户id
	Brand     string    `json:"brand"`  // 品牌
	Model     string    `json:"model"`  // 型号
	Name      string    `json:"name"`   // 名称
	Price     float64   `json:"price"`  // 商品价格
	ArticleId int64     `json:"articleId"`
	Status    int64     `json:"status"`    // 商品状态，默认1为草稿，2为发布，0为下架，-1为审核未通过需修改
	Stock     int64     `json:"stock"`     // 库存
	Tag       string    `json:"tag"`       // 分类名
	Type      int64     `json:"type"`      // 1为售卖商品，2为收商品
	CreateAt  time.Time `json:"createAt"`  // datetime，创建时间
	PublishAt time.Time `json:"publishAt"` // datetime，发布时间
	Collect   int64     `json:"collect"`   // 收藏数
	View      int64     `json:"view"`      // 查看数量
}

type CommodityCollect struct {
	Id          int64
	UserId      int64
	CommodityId int64
	Status      int8
	CreateAt    time.Time
}

type CommodityHistory struct {
	Id          int64
	UserId      int64
	CommodityId int64
	Time        time.Time
}

type CommodityTag struct {
	Id       int64
	Name     string
	CreateBy int64
	CreateAt time.Time
	UpdateBy int64
	UpdateAt time.Time
}
