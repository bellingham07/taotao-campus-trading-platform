package commodityLogic

import (
	"com.xpdj/go-gin/model/response"
	"com.xpdj/go-gin/utils/cache"
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
	zs := cache.RedisUtil.ZREVRANGEWITHSCORES(key, 0, -1)
	if zs == nil {
		return gin.H{"code": response.FAIL, "msg": "你还没有浏览过商品，快去看看有什么好物吧！😊"}
	}
	var zqualified []redis.Z
	now := time.Now()
	for index, z := range zs {
		createTimef := z.Score
		createTimei := int64(createTimef)
		createTime := time.Unix(createTimei, 0)
		// 计算 30 天前的时间
		if createTime.Before(now.AddDate(0, 0, -30)) {
			zqualified = zs[:index]
			go func() {
				var ids []interface{}
				for _, z := range zs[index:] {
					ids = append(ids, z.Member)
				}
				if err := cache.RedisUtil.ZREM(key, ids); err != nil {
					log.Printf("删除redis足迹出错，userId：%s\n", userId)
				}
			}()
			break
		}
	}
	var infos []map[string]string
	for _, z := range zqualified {
		id := z.Member.(string)
		key := cache.COMMODITYINFO + id
		info, _ := cache.RedisUtil.HGETALL(key)
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
	if err := cache.RedisUtil.ZADD(key, &member); err != nil {
		log.Printf("更新redis足迹失败，userId：%d\n", userId)
	}
}
