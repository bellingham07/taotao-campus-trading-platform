package commodityLogic

import (
	"com.xpwk/go-gin/cache"
	"com.xpwk/go-gin/model/response"
	commodityRepository "com.xpwk/go-gin/repository/commodity"
	"github.com/gin-gonic/gin"
	"strconv"
	"time"
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

func (*CommodityInfoLogic) GetById(id int64, userId int64, exist bool) gin.H {
	key := cache.COMMODITYINFO + strconv.FormatInt(id, 10)
	commodityStr, err := cache.RedisClient.Get(key)
	if commodityStr == "" {
		return gin.H{"code": response.FAIL, "msg": "没有此商品信息"}
	}
	if exist {
		HistoryLogic.UpdateHistory(id, userId)
	}
	commodityInfo, err := commodityRepository.CommodityInfo.QueryById(id)
	if err != nil {
		_ = cache.RedisClient.Set(key, "", time.Minute)
		return gin.H{"code": response.FAIL, "msg": "没有此商品信息"}
	}
	return gin.H{"code": response.OK, "msg": response.SUCCESS, "data": commodityInfo}
}

func L() {

}
