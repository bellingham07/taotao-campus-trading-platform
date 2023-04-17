package commodityLogic

import (
	"com.xpwk/go-gin/model/response"
	commodityRepository "com.xpwk/go-gin/repository/commodity"
	"github.com/gin-gonic/gin"
)

var CommodityInfo = new(CommodityInfoLogic)

type CommodityInfoLogic struct {
}

func (*CommodityInfoLogic) ListCategory() gin.H {
	return nil
}

func (*CommodityInfoLogic) SaveCommodity() gin.H {
	return nil
}

func (*CommodityInfoLogic) GetById(id int64) gin.H {
	commodityInfo, err := commodityRepository.CommodityInfo.QueryById(id)
	if err != nil {
		return gin.H{"code": response.FAIL, "msg": response.ERROR}
	}
	return gin.H{"code": response.OK, "msg": response.SUCCESS, "data": commodityInfo}
}

func L() {

}
