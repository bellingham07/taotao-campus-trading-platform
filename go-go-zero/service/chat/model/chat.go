package model

import "time"

type (
	ChatRoom struct {
		Id       int64     `xorm:"id" json:"id"`              // id
		CmdtyId  int64     `xorm:"cmdty_id" json:"cmdtyId"`   // 对应的商品信息
		SellerId int64     `xorm:"seller_id" json:"sellerId"` // 卖家id
		Seller   string    `xorm:"seller" json:"seller"`
		BuyerId  int64     `xorm:"buyer_id" json:"buyerId"`
		Buyer    string    `xorm:"buyer" json:"buyer"`
		Cover    string    `xorm:"cover" json:"cover"`
		CreateAt time.Time `xorm:"create_at" json:"createAt"`
		Status   int64     `xorm:"status" json:"status"`
	}
)
