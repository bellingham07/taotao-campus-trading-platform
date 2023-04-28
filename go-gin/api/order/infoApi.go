package orderApi

import (
	orderLogic "com.xpdj/go-gin/logic/order"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type InfoApi struct {
}

func (*InfoApi) List(c *gin.Context) {
	userIdAny, _ := c.Get("userid")
	userIdStr := userIdAny.(string)
	userId, _ := strconv.ParseInt(userIdStr, 10, 64)

	c.JSON(http.StatusOK, orderLogic.InfoLogic.ListByUserId(userId))
}
