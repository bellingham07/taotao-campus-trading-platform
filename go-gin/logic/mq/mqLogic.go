package mqLogic

import (
	"com.xpdj/go-gin/model"
	commodityRepository "com.xpdj/go-gin/repository/commodity"
	"com.xpdj/go-gin/utils/cache"
	"strconv"
	"strings"
	"time"
)

func CollectCheckUpdate(redisKey, member string, now time.Time) {
	isMember := cache.RedisUtil.SISMEMBER(redisKey, member)
	if isMember {
		commodityId, userId := getCmdtyIdAndUserIdByRedisKey(redisKey, member)
		collect := &model.CommodityCollect{
			CommodityId: commodityId,
			UserId:      userId,
			CreateAt:    now,
			Status:      1,
		}
		_ = commodityRepository.CommodityCollect.Insert(collect)
	}
}

func CollectCheckDelete(redisKey, member string) {
	isMember := cache.RedisUtil.SISMEMBER(redisKey, member)
	if !isMember {
		commodityId, userId := getCmdtyIdAndUserIdByRedisKey(redisKey, member)
		_ = commodityRepository.CommodityCollect.DeleteByCmdtyIdAndUserId(commodityId, userId)
	}
}

func getCmdtyIdAndUserIdByRedisKey(redisKey, member string) (int64, int64) {
	split := strings.LastIndex(redisKey, ":")
	commodityIdStr := redisKey[split+1:]
	commodityId, _ := strconv.ParseInt(commodityIdStr, 10, 64)
	userId, _ := strconv.ParseInt(member, 10, 64)
	return commodityId, userId
}
