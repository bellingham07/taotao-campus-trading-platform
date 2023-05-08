package request

type OrderMake struct {
	SellerId      int64   `json:"sellerId"`
	SellerName    string  `json:"sellerName"`
	CommodityId   int64   `json:"commodityId"`
	CommodityName string  `json:"commodityName"`
	Location      string  `json:"location"`
	Payment       float64 `json:"payment"`
}
