package commodityLogic

import (
	"com.xpwk/go-gin/cache"
	"com.xpwk/go-gin/model/response"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	"log"
	"strconv"
	"time"
)

var HistoryLogic = new(CommodityHistoryLogic)

type CommodityHistoryLogic struct {
}

func (*CommodityHistoryLogic) ListByUserId(userId string) gin.H {
	key := cache.COMMODITYHISOTRY + userId
	ids, err := cache.RedisClient.ZREVRANGE(key, 0, -1)
	if err != nil {
		return gin.H{"code": response.FAIL, "msg": "你还没有浏览过商品，快去看看有什么好物吧！😊"}
	}
	var infos []map[string]string
	for _, id := range ids {
		key := cache.COMMODITYINFO + id
		info, _ := cache.RedisClient.HGETALL(key)
		infos = append(infos, info)
	}
	return gin.H{"code": response.OK, "msg": response.SUCCESS, "data": infos}
}

func (*CommodityHistoryLogic) UpdateHistory(id int64, userId int64) {
	key := cache.COMMODITYHISOTRY + strconv.FormatInt(userId, 10)
	member := redis.Z{
		Score:  float64(time.Now().Unix()),
		Member: id,
	}
	if err := cache.RedisClient.ZADD(key, &member); err != nil {
		log.Println("Update History Fail!")
	}
}
