package model

import "time"

type CommodityInfo struct {
	Id            int64
	UserId        int64
	Name          string
	Category      int
	CategoryName  string
	Price         float32
	Brand         string
	Model         string
	Stock         int
	ArticlePostId int
	Status        int
	CreateAt      time.Time
	PublishAt     time.Time
	UpdateAt      time.Time
	Views         int
	Collect       int
}

type CommodityCollect struct {
	Id            int64
	UserId        int64
	CommodityId   int64
	CommodityName string
	Status        int8
	CrateAt       time.Time
}

type CommodityHistory struct {
	Id          int64
	UserId      int64
	CommodityId int64
	Time        time.Time
}
