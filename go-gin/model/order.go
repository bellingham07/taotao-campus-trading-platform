package model

import "time"

type OrderInfo struct {
	Id            int64     `json:"id"`
	SellerId      int64     `json:"sellerId"`
	SellerName    string    `json:"sellerName"`
	BuyerId       int64     `json:"buyerId"`
	BuyerName     string    `json:"buyerName"`
	CommodityId   int64     `json:"commodityId"`
	CommodityName string    `json:"commodityName"`
	Payment       float64   `json:"payment"`
	Location      string    `json:"location"`
	Status        int8      `json:"status"`
	CreateAt      time.Time `json:"createAt"`
	DoneAt        time.Time `json:"doneAt"`
	IsGood        string    `json:"isGood"`
}

type OrderComment struct {
	Id       int64
	UserId   int64
	ToUserId int64
	Content  string
	Type     int8
	CreateAt time.Time
}
