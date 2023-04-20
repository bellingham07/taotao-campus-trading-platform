package commodityLogic

import (
	"com.xpwk/go-gin/cache"
	"github.com/gin-gonic/gin"
	"strconv"
)

var HistoryLogic = new(CommodityHistoryLogic)

type CommodityHistoryLogic struct {
}

func (*CommodityHistoryLogic) ListByUserId(userId string) gin.H {
	return nil
}

func (*CommodityHistoryLogic) UpdateHistory(id int64, userId int64) gin.H {
	_ = cache.COMMODITYHISOTRY + strconv.FormatInt(userId, 10) + ":" + strconv.FormatInt(id, 10)
	return nil
}
