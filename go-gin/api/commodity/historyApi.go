package commodityApi

import (
	commodityLogic "com.xpwk/go-gin/logic/commodity"
	"github.com/gin-gonic/gin"
	"strconv"
)

type CommodityHistoryApi struct {
}

func (*CommodityHistoryApi) List(ctx *gin.Context) {
	userIdStr := ctx.Param("userId")
	userId, _ := strconv.ParseInt(userIdStr, 10, 10)
	commodityLogic.HistoryLogic.ListByUserId(userId)
}
