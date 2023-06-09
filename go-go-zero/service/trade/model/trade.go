package model

import "time"

type (
	TradeInfo struct {
		Id                int64     `xorm:"id" json:"id"`                  // id
		BuyerId           int64     `xorm:"buyer_id" json:"buyerId"`       // 卖家id
		Buyer             string    `xorm:"buyer" json:"buyer"`            // 卖家名
		SellerId          int64     `xorm:"seller_id" json:"sellerId"`     // 买家id
		Seller            string    `xorm:"seller" json:"seller"`          // 买家名
		CmdtyId           int64     `xorm:"cmdty_id" json:"cmdtyId"`       // 商品id
		Type              int8      `xorm:"type" json:"type"`              // 和商品type一样，1表示“出售商品”，2表示“求购商品”
		BriefIntro        string    `xorm:"brief_intro" json:"briefIntro"` // 商品名
		Cover             string    `xorm:"cover" json:"cover"`
		Payment           float64   `xorm:"payment" json:"payment"`
		Location          string    `xorm:"location" json:"location"`
		Status            int8      `xorm:"status" json:"status"`
		CreateAt          time.Time `xorm:"create_at" json:"createAt"` // 创建时间
		IsSellerConfirmed int8      `xorm:"is_seller_confirmed" json:"isSellerConfirmed"`
		IsBuyerConfirmed  int8      `xorm:"is_buyer_confirmed" json:"isBuyerConfirmed"`
		IsSellerDone      int8      `xorm:"is_seller_done" json:"isSellerDone"` // 默认0，完成1
		IsBuyerDone       int8      `xorm:"is_buyer_done" json:"isBuyerDone"`   // 默认0，完成1
		SellerDoneAt      time.Time `xorm:"seller_done_at" json:"sellerDoneAt"`
		BuyerDoneAt       time.Time `xorm:"buyer_done_at" json:"buyerDoneAt"`
	}

	TradeDone struct {
		Id           int64     `xorm:"id" json:"id"`                  // id
		BuyerId      int64     `xorm:"buyer_id" json:"buyerId"`       // 卖家id
		Buyer        string    `xorm:"buyer" json:"buyer"`            // 卖家名
		SellerId     int64     `xorm:"seller_id" json:"sellerId"`     // 买家id
		Seller       string    `xorm:"seller" json:"seller"`          // 买家名
		CmdtyId      int64     `xorm:"cmdty_id" json:"cmdtyId"`       // 商品id
		Type         int8      `xorm:"type" json:"type"`              // 和商品type一样，1表示“出售商品”，2表示“求购商品”
		BriefIntro   string    `xorm:"brief_intro" json:"briefIntro"` // 商品名
		Cover        string    `xorm:"cover" json:"cover"`
		Payment      float64   `xorm:"payment" json:"payment"`
		Location     string    `xorm:"location" json:"location"`
		Status       int8      `xorm:"status" json:"status"`
		CreateAt     time.Time `xorm:"create_at" json:"createAt"`
		SellerDoneAt time.Time `xorm:"seller_done_at" json:"sellerDoneAt"`
		BuyerDoneAt  time.Time `xorm:"buyer_done_at" json:"buyerDoneAt"`
		DoneAt       time.Time `xorm:"done_at" json:"doneAt"` // 创建时间
		IsSellerCmt  int8      `xorm:"is_seller_cmt" json:"isSellerCmt"`
		IsBuyerCmt   int8      `xorm:"is_buyer_cmt" json:"isBuyerCmt"`
		SellerCmtAt  time.Time `xorm:"seller_cmt_at" json:"sellerCmtAt"`
		BuyerCmtAt   time.Time `xorm:"buyer_cmt_at" json:"buyerCmtAt"`
		IsCmtDone    int8      `xorm:"is_cmt_done" json:"is_cmt_done"`
	}

	TradeCmt struct {
		Id         int64     `bson:"_id"` // id
		TradeId    int64     `bson:"trade_id" json:"tradeId"`
		UserId     int64     `bson:"user_id" json:"userId"`
		User       string    `bson:"user" json:"user"`
		UserAvatar string    `bson:"user_avatar" json:"userAvatar"`
		ToUserId   int64     `bson:"to_user_id" json:"toUserId"`
		Content    string    `bson:"content" json:"content"`    // 评价内容
		Type       int       `bson:"type" json:"type"`          // 差评或好评，0为差评，1为好评
		CreateAt   time.Time `bson:"create_at" json:"createAt"` // 创建时间
	}
)
