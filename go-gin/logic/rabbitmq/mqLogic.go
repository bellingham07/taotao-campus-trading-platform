package mqLogic

import (
	"com.xpdj/go-gin/model"
	articleRepository "com.xpdj/go-gin/repository/article"
	commodityRepository "com.xpdj/go-gin/repository/commodity"
	userRepository "com.xpdj/go-gin/repository/user"
	"com.xpdj/go-gin/utils/cache"
	"strconv"
	"strings"
	"time"
)

func CollectCheckUpdate(ccMessage *CcMessage) {
	redisKey := ccMessage.RedisKey
	userId := ccMessage.UserId
	isMember := cache.RedisUtil.SISMEMBER(redisKey, userId)
	if isMember {
		commodityId := getIdByRedisKey(redisKey)
		collect := &model.CommodityCollect{
			CommodityId: commodityId,
			UserId:      userId,
			CreateAt:    ccMessage.Time,
			Status:      1,
		}
		_ = commodityRepository.CommodityCollect.Insert(collect)
	}
}

func CollectCheckDelete(ccMessage *CcMessage) {
	redisKey := ccMessage.RedisKey
	userId := ccMessage.UserId
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

func ViewCheckUpdate(vMessage *VMessage) {
	redisKey := vMessage.RedisKey
	isCommodity := vMessage.IsCommodity
	viewCountMap := cache.RedisUtil.HGETALL(redisKey)
	countStr, ok := viewCountMap["count"]
	now := time.Now()
	// 如果取不到就丢弃该次更改
	if !ok {
		return
	}
	count, _ := strconv.ParseInt(countStr, 10, 64)
	viewTimeStr, ok := viewCountMap["time"]
	var viewTime time.Time
	// 第一次没有设置过时间，直接拿会报错，这里解决一下
	if !ok {
		viewTime = now.Add(-5 * time.Minute)
	}
	viewTime, _ = time.Parse("2006-01-02 15:04:05", viewTimeStr)
	id := getIdByRedisKey(redisKey)
	// 如果已经累计了50个浏览量了，或者过去了1分钟，那我们就更新库
	if count >= 50 || viewTime.Before(now.Add(-time.Minute)) {
		// 更新商品的浏览量
		if isCommodity {
			_ = commodityRepository.CommodityInfo.UpdateViewById(id, count)
			_ = cache.RedisUtil.HSET1(redisKey, "time", time.Now())
			// 更新redis
			_ = cache.RedisUtil.HINCRBY(cache.CommodityInfo+strconv.FormatInt(id, 10), "view", count)
			_ = cache.RedisUtil.HDEL(redisKey, "count")
		} else {
			// 更新文章的浏览量，同理
			_ = articleRepository.ArticleContent.UpdateViewById(id, count)
			_ = cache.RedisUtil.HSET1(redisKey, "time", time.Now())
			_ = cache.RedisUtil.HINCRBY(cache.ArticleContent+strconv.FormatInt(id, 10), "view", count)
			_ = cache.RedisUtil.HDEL(redisKey, "count")
		}
	}
}

func LikeCheckUpdate(lMessage *LMessage) {
	redisKey := lMessage.RedisKey
	userId := lMessage.UserId
	isMember := cache.RedisUtil.SISMEMBER(redisKey, userId)
	if !isMember {
		return
	}
	id := getIdByRedisKey(redisKey)
	if lMessage.IsArticle {
		_ = articleRepository.ArticleContent.UpdateLikeById(id)
	} else {
		_ = commodityRepository.CommodityInfo.UpdateLikeById(id)
	}
	_ = userRepository.UserInfo.UpdateLikeById(lMessage.UserId)
}
