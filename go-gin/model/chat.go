package model

import "time"

type ChatRoom struct {
	Id          int64     `json:"id" form:"id"`
	CommodityId int64     `json:"commodityId" form:"commodityId"`
	SellerId    int64     `json:"sellerId" form:"sellerId"`
	SellerName  string    `json:"sellerName" form:"sellerName"`
	BuyerId     int64     `json:"buyerId" form:"buyerId"`
	BuyerName   string    `json:"buyerName" form:"buyerName"`
	Cover       string    `json:"cover" form:"cover"`
	CreateAt    time.Time `json:"createAt" form:"createAt"`
}

type ChatMessage struct {
	Id       int64     `json:"id" form:"id"`
	RoomId   int64     `json:"roomId" form:"roomId"`
	Content  string    `json:"content" form:"content"`
	Time     time.Time `json:"time" form:"time"`
	UserId   int64     `json:"userId" form:"userId"`
	ToUserId int64     `json:"toUserId" form:"toUserId"`
}
