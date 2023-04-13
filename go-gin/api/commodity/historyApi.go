package commodityApi

import (
	commodityLogic "com.xpwk/go-gin/logic/commodity"
	"github.com/gin-gonic/gin"
)

type CommodityHistoryApi struct {
}

func (*CommodityHistoryApi) List(ctx *gin.Context) {
	userid := ctx.Param("userid")
	commodityLogic.HistoryLogic.ListByUserId()
}
