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
		return response.GenH(response.FAIL, response.ERROR)
	}
	return response.GenH(response.OK, response.SUCCESS)
}

func (*CommodityCollectLogic) Uncollect(id, userId string) gin.H {
	key := utils.COMMODITYCOLLECT + userId

	if err := utils.RedisUtil.ZREM(key, id); err != nil {
		return response.GenH(response.FAIL, response.ERROR)
	}
	return response.GenH(response.OK, response.SUCCESS)
}

func (*CommodityCollectLogic) List(userId string) gin.H {
	key := utils.COMMODITYCOLLECT + userId
	ids := utils.RedisUtil.ZREVRANGE(key, 0, -1)
	if ids == nil {
		return response.GenH(response.FAIL, response.ERROR)
	}
	var infosMap []map[string]string
	for _, id := range ids {
		key := utils.COMMODITYINFO + id
		infoMap, _ := utils.RedisUtil.HGETALL(key)
		infosMap = append(infosMap, infoMap)
	}
	return response.GenH(response.OK, response.SUCCESS, infosMap)
}
