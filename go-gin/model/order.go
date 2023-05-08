package model

import "time"

type OrderInfo struct {
	Id            int64
	Seller        int64
	SellerName    string
	Buyer         int64
	BuyerName     string
	CommodityId   int64
	CommodityName string
	Payment       float64
	Location      string
	Status        int8
	CreateAt      time.Time
	DoneAt        time.Time
}

type OrderComment struct {
	Id       int64
	UserId   int64
	ToUserId int64
	Content  string
	Type     int8
	CreateAt time.Time
}
