package commodityLogic

import (
	"com.xpdj/go-gin/model/response"
	"com.xpdj/go-gin/utils"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	"time"
)

var CollectLogic = new(CommodityCollectLogic)

type CommodityCollectLogic struct {
}

func (*CommodityCollectLogic) Collect(id, userId string) gin.H {
	key := utils.COMMODITYCOLLECT + userId
	z := redis.Z{
		Score:  float64(time.Now().Unix()),
		Member: id,
	}
	err := utils.RedisUtil.ZADDNX(key, &z)
	if err != nil {
		return gin.H{"code": response.FAIL, "msg": response.ERROR}
	}
	return gin.H{"code": response.OK, "msg": response.SUCCESS}
}

func (*CommodityCollectLogic) Uncollect(id, userId string) gin.H {
	key := utils.COMMODITYCOLLECT + userId

	if err := utils.RedisUtil.ZREM(key, id); err != nil {
		return gin.H{"code": response.FAIL, "msg": response.ERROR}
	}
	return gin.H{"code": response.OK, "msg": response.SUCCESS}
}

func (*CommodityCollectLogic) List(userId string) gin.H {
	key := utils.COMMODITYCOLLECT + userId
	ids := utils.RedisUtil.ZREVRANGE(key, 0, -1)
	if ids == nil {
		return gin.H{"code": response.FAIL, "msg": response.ERROR}
	}
	var commodityInfosMap []map[string]string
	for _, id := range ids {
		key := utils.COMMODITYINFO + id
		infoMap, _ := utils.RedisUtil.HGETALL(key)
		commodityInfosMap = append(commodityInfosMap, infoMap)
	}
	return gin.H{"code": response.OK, "msg": response.SUCCESS, "data": commodityInfosMap}
}
