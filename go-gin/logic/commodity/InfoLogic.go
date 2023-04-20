package commodityLogic

import (
	"com.xpwk/go-gin/cache"
	"com.xpwk/go-gin/model"
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
	var commodityInfo model.CommodityInfo
	key := cache.COMMODITYINFO + strconv.FormatInt(id, 10)
	commodityInfoStr, err := cache.RedisClient.HGETALL(key)
	if commodityInfoStr["id"] == "" {
		_ = cache.RedisClient.HSET(key, commodityInfo)
		_ = cache.RedisClient.EXPIRE(key, 30*time.Second)
		return gin.H{"code": response.FAIL, "msg": "没有此商品信息"}
	}
	// redis没有数据
	if err != nil {
		commodityInfo, err = commodityRepository.CommodityInfo.QueryById(id)
		// 数据无，设置空
		if err != nil {
			_ = cache.RedisClient.HSET(key, commodityInfo)
			_ = cache.RedisClient.EXPIRE(key, 30*time.Second)
			return gin.H{"code": response.FAIL, "msg": "没有此商品信息"}
		}
		// jwt中存在用户，判断是在访问自己的商品还是别人的
		if exist {
			// redis取出的值不为空则说明，redis中有
			if commodityInfo.UserId != userId {
				HistoryLogic.UpdateHistory(id, userId)
			}
		}
		return gin.H{"code": response.OK, "msg": response.SUCCESS, "data": commodityInfo}
	}
	// redis有数据
	return nil
}

func L() {

}
