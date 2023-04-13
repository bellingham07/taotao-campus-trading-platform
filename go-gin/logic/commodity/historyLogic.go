package commodityLogic

import "github.com/gin-gonic/gin"

var HistoryLogic = new(CommodityHistoryLogic)

type CommodityHistoryLogic struct {
}

func (*CommodityHistoryLogic) ListByUserId() gin.H {

}
