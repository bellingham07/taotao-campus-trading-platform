package commodityLogic

import "github.com/gin-gonic/gin"

var HistoryLogic = new(CommodityHistoryLogic)

type CommodityHistoryLogic struct {
}

func (*CommodityHistoryLogic) ListByUserId(userId int64) gin.H {
	return nil
}
