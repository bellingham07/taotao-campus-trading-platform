package mqLogic

import (
	"com.xpdj/go-gin/model"
	articleRepository "com.xpdj/go-gin/repository/article"
	commodityRepository "com.xpdj/go-gin/repository/commodity"
	"com.xpdj/go-gin/utils/cache"
	"strconv"
	"strings"
	"time"
)

func CollectCheckUpdate(redisKey string, userId int64, now time.Time) {
	isMember := cache.RedisUtil.SISMEMBER(redisKey, userId)
	if isMember {
		commodityId := getIdByRedisKey(redisKey)
		collect := &model.CommodityCollect{
			CommodityId: commodityId,
			UserId:      userId,
			CreateAt:    now,
			Status:      1,
		}
		_ = commodityRepository.CommodityCollect.Insert(collect)
	}
}

func CollectCheckDelete(redisKey string, userId int64) {
	isMember := cache.RedisUtil.SISMEMBER(redisKey, strconv.FormatInt(userId, 10))
	if !isMember {
		commodityId := getIdByRedisKey(redisKey)
		_ = commodityRepository.CommodityCollect.DeleteByCmdtyIdAndUserId(commodityId, userId)
	}
}

func getIdByRedisKey(redisKey string) int64 {
	split := strings.LastIndex(redisKey, ":")
	idStr := redisKey[split+1:]
	id, _ := strconv.ParseInt(idStr, 10, 64)
	return id
}

func ViewCheckUpdate(redisKey string, isCommodity bool) {
	viewCountMap := cache.RedisUtil.HGETALL(redisKey)
	countStr, ok := viewCountMap["count"]
	// 如果取不到就丢弃该次更改
	if !ok {
		return
	}
	count, _ := strconv.ParseInt(countStr, 10, 64)
	viewTimeStr := viewCountMap["time"]
	viewTime, _ := time.Parse("2006-01-02 15:04:05", viewTimeStr)
	now := time.Now()
	id := getIdByRedisKey(redisKey)
	// 如果已经累计了50个浏览量了，或者过去了五分钟，那我们就更新库
	if count >= 50 || viewTime.Before(now.Add(-time.Minute*5)) {
		if isCommodity {
			err := commodityRepository.CommodityInfo.UpdateViewById(id, count)
			if err != nil {
				_ = cache.RedisUtil.HSET1(redisKey, "time", time.Now())
			}
			_ = cache.RedisUtil.HDEL(redisKey, "count")
		} else {
			err := articleRepository.ArticleContent.UpdateViewById(id, count)
			if err != nil {
				_ = cache.RedisUtil.HSET1(redisKey, "time", time.Now())
			}
			_ = cache.RedisUtil.HDEL(redisKey, "count")
		}
	}
}
