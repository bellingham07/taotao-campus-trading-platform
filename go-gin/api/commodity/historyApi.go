package commodityApi

import (
	commodityLogic "com.xpdj/go-gin/logic/commodity"
	"github.com/gin-gonic/gin"
)

type HistoryApi struct {
}

func (*HistoryApi) List(c *gin.Context) {
	userIdAny, _ := c.Get("userid")
	userIdStr := userIdAny.(string)
	commodityLogic.HistoryLogic.ListByUserId(userIdStr)
}

func (*HistoryApi) Delete(c *gin.Context) {

}
