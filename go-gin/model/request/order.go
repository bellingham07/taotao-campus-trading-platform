package request

import "time"

type OrderDto struct {
	Id            int64
	OwnerId       int64  `json:"ownerId"`
	OwnerName     string `json:"ownerName"`
	SellerId      int64
	SellerName    string
	BuyerId       int64
	BuyerName     string
	CommodityId   int64
	CommodityName string
	Payment       float64
	Location      string
	Status        int8
	CreateAt      time.Time
	DoneAt        time.Time
	IsGood        string
}
