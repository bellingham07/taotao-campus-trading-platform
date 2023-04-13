package commodityLogic

import "github.com/gin-gonic/gin"

var CommodityInfo = new(CommodityInfoLogic)

type CommodityInfoLogic struct {
}

func (*CommodityInfoLogic) ListCategory() gin.H {
	return nil
}

func (*CommodityInfoLogic) SaveCommodity() gin.H {
	return nil
}

func L() {

}
