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

func (*InfoApi) GetInfoById(ctx *gin.Context) {
	idStr := ctx.Param("id")
	if idStr == "" {
		ctx.JSON(http.StatusUnauthorized, gin.H{"code": response.FAIL, "msg": response.ERROR})
		return
	}
	id, _ := strconv.ParseInt(idStr, 10, 64)
	ctx.JSON(http.StatusOK, commodityLogic.CommodityInfo.GetById(id))
}
