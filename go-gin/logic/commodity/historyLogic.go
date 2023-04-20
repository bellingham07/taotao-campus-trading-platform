package commodityLogic

import (
	"com.xpwk/go-gin/cache"
	commodityRepository "com.xpwk/go-gin/repository/commodity"
	"github.com/gin-gonic/gin"
	"strconv"
	"time"
)

var HistoryLogic = new(CommodityHistoryLogic)

type CommodityHistoryLogic struct {
}

func (*CommodityHistoryLogic) ListByUserId(userId int64) gin.H {
	commodityRepository.CommodityHistory.ListByUserId(userId)
	return nil
}

func (*CommodityHistoryLogic) UpdateHistory(id int64, userId int64) gin.H {
	key := cache.COMMODITYHISOTRY + strconv.FormatInt(userId, 10) + ":" + strconv.FormatInt(id, 10)
	_ = cache.RedisClient.SET(key, time.Now(), 30*24*time.Hour)
	return nil
}
