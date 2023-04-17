package orderApi

import (
	orderLogic "com.xpwk/go-gin/logic/order"
	"com.xpwk/go-gin/model/response"
	"github.com/gin-gonic/gin"
	"net/http"
	"reflect"
)

type InfoApi struct {
}

func (*InfoApi) List(c *gin.Context) {
	userIdStr, exist := c.Get("userid")
	if !exist {
		c.JSON(http.StatusUnauthorized, gin.H{"code": response.FAIL, "msg": "请先登录"})
		return
	}
	userId := reflect.ValueOf(userIdStr).Int()
	c.JSON(http.StatusOK, orderLogic.InfoLogic.ListByUserId(userId))
}
