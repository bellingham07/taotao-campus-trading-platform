package commodityApi

import (
	commodityLogic "com.xpwk/go-gin/logic/commodity"
	"com.xpwk/go-gin/model/response"
	"github.com/gin-gonic/gin"
	"net/http"
	"reflect"
	"strconv"
)

type InfoApi struct {
}

func (*InfoApi) GetInfoById(c *gin.Context) {
	idStr := c.Param("id")
	if idStr == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"code": response.FAIL, "msg": response.ERROR})
		return
	}
	id, _ := strconv.ParseInt(idStr, 10, 64)
	userIdStr, exist := c.Get("userid")
	if !exist {
		c.JSON(http.StatusOK, commodityLogic.CommodityInfo.GetById(id, 0, exist))
		return
	}
	userId := reflect.ValueOf(userIdStr).Int()
	c.JSON(http.StatusOK, commodityLogic.CommodityInfo.GetById(id, userId, exist))
}
