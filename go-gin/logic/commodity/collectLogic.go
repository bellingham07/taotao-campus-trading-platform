package commodityLogic

import (
	"com.xpwk/go-gin/cache"
	"com.xpwk/go-gin/model/response"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	"time"
)

var CollectLogic = new(CommodityCollectLogic)

type CommodityCollectLogic struct {
}

func (*CommodityCollectLogic) Collect(id, userId string) gin.H {
	key := cache.COMMODITYCOLLECT + userId
	z := redis.Z{
		Score:  float64(time.Now().Unix()),
		Member: id,
	}
	err := cache.RedisClient.ZADDNX(key, &z)
	if err != nil {
		return gin.H{"code": response.FAIL, "msg": response.ERROR}
	}
	return gin.H{"code": response.OK, "msg": response.SUCCESS}
}

func (*CommodityCollectLogic) Uncollect(id, userId string) gin.H {
	key := cache.COMMODITYCOLLECT + userId

	if err := cache.RedisClient.ZREM(key, id); err != nil {
		return gin.H{"code": response.FAIL, "msg": response.ERROR}
	}
	return gin.H{"code": response.OK, "msg": response.SUCCESS}
}

func (*CommodityCollectLogic) List(userId string) gin.H {
	key := cache.COMMODITYCOLLECT + userId
	ids := cache.RedisClient.ZREVRANGE(key, 0, -1)
	if ids == nil {
		return gin.H{"code": response.FAIL, "msg": response.ERROR}
	}
	var commodityInfosMap []map[string]string
	for _, id := range ids {
		key := cache.COMMODITYINFO + id
		infoMap, _ := cache.RedisClient.HGETALL(key)
		commodityInfosMap = append(commodityInfosMap, infoMap)
	}
	return gin.H{"code": response.OK, "msg": response.SUCCESS, "data": commodityInfosMap}
}
