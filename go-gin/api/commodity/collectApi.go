package commodityApi

import (
	commodityLogic "com.xpwk/go-gin/logic/commodity"
	"com.xpwk/go-gin/router/middleware"
	"github.com/gin-gonic/gin"
	"net/http"
)

type CollectApi struct {
}

func (*CollectApi) Collect(c *gin.Context) {
	userIdStr := middleware.GetUserIdStr(c)
	id := c.Param("id")
	c.JSON(http.StatusOK, commodityLogic.CollectLogic.Collect(id, userIdStr))
}

func (a *CollectApi) Uncollect(c *gin.Context) {
	userIdStr := middleware.GetUserIdStr(c)
	id := c.Param("id")
	c.JSON(http.StatusOK, commodityLogic.CollectLogic.Uncollect(id, userIdStr))
}

func (a *CollectApi) List(c *gin.Context) {
	userIdStr := middleware.GetUserIdStr(c)
	c.JSON(http.StatusOK, commodityLogic.CollectLogic.List(userIdStr))
}
