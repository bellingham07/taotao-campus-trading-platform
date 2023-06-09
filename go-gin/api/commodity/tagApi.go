package commodityApi

import (
	commodityLogic "com.xpdj/go-gin/logic/commodity"
	"com.xpdj/go-gin/model"
	"com.xpdj/go-gin/model/response"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type TagApi struct {
}

func (*TagApi) List(c *gin.Context) {
	c.JSON(http.StatusOK, commodityLogic.CategoryLogic.List)
}

func (*TagApi) Add(c *gin.Context) {
	var tag = new(model.CommodityTag)
	err := c.ShouldBind(tag)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.ErrorMsg("请求参数错误！"))
	}
	c.JSON(http.StatusOK, commodityLogic.CategoryLogic.Add(tag))
}

func (*TagApi) Remove(c *gin.Context) {
	idStr := c.Param("id")
	if idStr == "" {
		c.JSON(http.StatusBadRequest, response.ErrorMsg("请求参数错误！"))
	}
	id, err := strconv.ParseInt(idStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.ErrorMsg("请求参数错误！"))
	}
	c.JSON(http.StatusOK, commodityLogic.CategoryLogic.RemoveById(id))
}
