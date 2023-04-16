package commodityLogic

import (
	commodityRepository "com.xpwk/go-gin/repository/commodity"
	"github.com/gin-gonic/gin"
)

var HistoryLogic = new(CommodityHistoryLogic)

type CommodityHistoryLogic struct {
}

func (*CommodityHistoryLogic) ListByUserId(userId int64) gin.H {
	commodityRepository.CommodityHistory.ListByUserId(userId)
	return nil
}
