package commodityApi

import (
	commodityLogic "com.xpdj/go-gin/logic/commodity"
	"com.xpdj/go-gin/model/response"
	"com.xpdj/go-gin/router/middleware"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
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
	idStr := c.Param("id")
	_, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.OkMsg("请求参数错误！"))
	}
	c.JSON(http.StatusOK, commodityLogic.CollectLogic.Uncollect(idStr, userIdStr))
}

func (a *CollectApi) List(c *gin.Context) {
	userIdStr := middleware.GetUserIdStr(c)
	c.JSON(http.StatusOK, commodityLogic.CollectLogic.List(userIdStr))
}
