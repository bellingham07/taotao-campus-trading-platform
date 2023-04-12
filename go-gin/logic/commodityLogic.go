package logic

import "github.com/gin-gonic/gin"

var Commodity = new(CommodityLogic)

type CommodityLogic struct {
}

func (*CommodityLogic) ListCategory() gin.H {
	return nil
}

func (*CommodityLogic) SaveCommodity() gin.H {
	return nil
}

func L() {

}
