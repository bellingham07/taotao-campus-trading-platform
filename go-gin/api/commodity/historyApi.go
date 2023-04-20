package commodityApi

import (
	commodityLogic "com.xpwk/go-gin/logic/commodity"
	"github.com/gin-gonic/gin"
)

type HistoryApi struct {
}

func (*HistoryApi) List(c *gin.Context) {
	userIdStr := c.Param("userId")
	commodityLogic.HistoryLogic.ListByUserId(userIdStr)
}
