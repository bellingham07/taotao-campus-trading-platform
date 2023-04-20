package commodityApi

import (
	commodityLogic "com.xpwk/go-gin/logic/commodity"
	"github.com/gin-gonic/gin"
	"net/http"
)

type CollectApi struct {
}

func (*CollectApi) Collect(c *gin.Context) {
	userIdAny, _ := c.Get("userid")
	userIdStr := userIdAny.(string)
	id := c.Param("id")
	c.JSON(http.StatusOK, commodityLogic.CollectLogic.Collect(id, userIdStr))

}

func (a *CollectApi) Uncollect(c *gin.Context) {
	userIdAny, _ := c.Get("userid")
	userIdStr := userIdAny.(string)
	id := c.Param("id")
	c.JSON(http.StatusOK, commodityLogic.CollectLogic.Uncollect(id, userIdStr))
}

func (a *CollectApi) List(c *gin.Context) {
	userIdAny, _ := c.Get("userid")
	userIdStr := userIdAny.(string)
	c.JSON(http.StatusOK, commodityLogic.CollectLogic.List(userIdStr))
}
