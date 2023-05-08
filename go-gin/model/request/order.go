package request

type OrderMake struct {
	CommodityId int64   `json:"commodityId"`
	Quantity    float64 `json:"quantity"`
	Location    string  `json:"location"`
}
