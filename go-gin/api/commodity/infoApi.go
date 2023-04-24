package commodityApi

import (
	commodityLogic "com.xpwk/go-gin/logic/commodity"
	"com.xpwk/go-gin/model/response"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type InfoApi struct {
}

func (*InfoApi) GetInfoById(c *gin.Context) {
	idStr := c.Param("id")
	if idStr == "" {
		c.JSON(http.StatusBadRequest, gin.H{"code": response.FAIL, "msg": response.ERROR})
		return
	}
	id, _ := strconv.ParseInt(idStr, 10, 64)
	userIdAny, exist := c.Get("userid")
	userIdStr := userIdAny.(string)
	userId, _ := strconv.ParseInt(userIdStr, 10, 64)
	c.JSON(http.StatusOK, commodityLogic.CommodityInfo.GetById(id, userId, exist))
}

func (a *InfoApi) ListByOption(c *gin.Context) {
	optionStr := c.Param("option")
	option, _ := strconv.Atoi(optionStr)
	c.JSON(http.StatusOK, commodityLogic.CommodityInfo.RandomListByType(option))
}
