package commodityLogic

import (
	"com.xpdj/go-gin/model/response"
	"com.xpdj/go-gin/utils/cache"
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
	err := cache.RedisUtil.ZADDNX(key, &z)
	if err != nil {
		return response.GenH(response.FAIL, response.ERROR)
	}
	return response.GenH(response.OK, response.SUCCESS)
}

//func (*CommodityCollectLogic) CollectRedis(id, userId string) gin.H {
//	key := cache.COMMODITYCOLLECT + id
//	err := cache.RedisUtil.HSETNX(key, userId, "")
//	if err != nil {
//		return response.GenH(response.FAIL, "服务器繁忙，请稍后😊")
//	}
//	err = delayInsertCollect(id, userId)
//	if err != nil {
//		return response.GenH(response.FAIL, response.ERROR)
//	}
//	return response.GenH(response.OK, response.SUCCESS)
//}

func (*CommodityCollectLogic) Uncollect(id, userId string) gin.H {
	key := cache.COMMODITYCOLLECT + userId

	if err := cache.RedisUtil.ZREM(key, id); err != nil {
		return response.GenH(response.FAIL, response.ERROR)
	}
	return response.GenH(response.OK, response.SUCCESS)
}

func (*CommodityCollectLogic) List(userId string) gin.H {
	key := cache.COMMODITYCOLLECT + userId
	ids := cache.RedisUtil.ZREVRANGE(key, 0, -1)
	if ids == nil {
		return response.GenH(response.FAIL, response.ERROR)
	}
	var infosMap []map[string]string
	for _, id := range ids {
		key := cache.COMMODITYINFO + id
		infoMap, _ := cache.RedisUtil.HGETALL(key)
		infosMap = append(infosMap, infoMap)
	}
	return response.GenH(response.OK, response.SUCCESS, infosMap)
}

//func delayInsertCollect() error {
//
//}
