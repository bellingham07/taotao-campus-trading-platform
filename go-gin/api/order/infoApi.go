package orderApi

import (
	orderLogic "com.xpdj/go-gin/logic/order"
	"com.xpdj/go-gin/model/response"
	"com.xpdj/go-gin/router/middleware"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type InfoApi struct {
}

func (*InfoApi) List(c *gin.Context) {
	userId := middleware.GetUserId(c)
	c.JSON(http.StatusOK, orderLogic.InfoLogic.ListByUserId(userId))
}

func (*InfoApi) GetById(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.GenH(response.FAIL, "请求参数错误！"))
	}
	c.JSON(http.StatusOK, orderLogic.InfoLogic.GetById(id))
}

func (*InfoApi) Buy(c *gin.Context) {
	//c.ShouldBind()
}
