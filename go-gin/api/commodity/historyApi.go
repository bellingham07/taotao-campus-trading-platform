package commodityApi

import (
	commodityLogic "com.xpwk/go-gin/logic/commodity"
	"github.com/gin-gonic/gin"
	"strconv"
)

type HistoryApi struct {
}

func (*HistoryApi) List(c *gin.Context) {
	userIdStr := c.Param("userId")
	userId, _ := strconv.ParseInt(userIdStr, 10, 64)
	commodityLogic.HistoryLogic.ListByUserId(userId)
}
