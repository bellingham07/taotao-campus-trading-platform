package commodityApi

import (
	commodityLogic "com.xpdj/go-gin/logic/commodity"
	"com.xpdj/go-gin/model"
	"com.xpdj/go-gin/model/response"
	"com.xpdj/go-gin/router/middleware"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type InfoApi struct {
}

func (*InfoApi) GetInfoById(c *gin.Context) {
	idStr := c.Param("id")
	if idStr == "" {
		c.JSON(http.StatusBadRequest, response.ErrorMsg("请求参数错误！"))
		return
	}
	id, _ := strconv.ParseInt(idStr, 10, 64)
	userIdAny, isLogin := c.Get("userid")
	var userIdStr = ""
	if isLogin {
		userIdStr = userIdAny.(string)
	}
	userId, _ := strconv.ParseInt(userIdStr, 10, 64)
	c.JSON(http.StatusOK, commodityLogic.CommodityInfo.GetById(id, userId))
}

func (*InfoApi) ListByOption(c *gin.Context) {
	optionStr := c.Param("option")
	option, _ := strconv.Atoi(optionStr)
	c.JSON(http.StatusOK, commodityLogic.CommodityInfo.RandomListByType(option))
}

func (*InfoApi) SellSave(c *gin.Context) {
	var draft = new(model.CommodityInfo)
	err := c.ShouldBind(draft)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.ErrorMsg("请求参数错误！"))
		return
	}
	// 第一次保存
	if draft.Id == 0 {
		c.JSON(http.StatusOK, commodityLogic.CommodityInfo.SaveOrPublishInfo(draft, middleware.GetUserId(c), 1, false))
	}
	// 之前保存过
	c.JSON(http.StatusOK, commodityLogic.CommodityInfo.UpdateInfo(draft, false))
}

func (*InfoApi) SellPublish(c *gin.Context) {
	var info = new(model.CommodityInfo)
	err := c.ShouldBind(info)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.ErrorMsg("请求参数错误！"))
		return
	}
	// 未保存就发布
	if info.Id == 0 {
		c.JSON(http.StatusOK, commodityLogic.CommodityInfo.SaveOrPublishInfo(info, middleware.GetUserId(c), 1, true))
		return
	}
	// 发布（更新）
	c.JSON(http.StatusOK, commodityLogic.CommodityInfo.UpdateInfo(info, true))
}

func (*InfoApi) WantSave(c *gin.Context) {
	var draft = new(model.CommodityInfo)
	err := c.ShouldBind(draft)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.ErrorMsg("请求参数错误！"))
		return
	}
	// 第一次保存
	if draft.Id == 0 {
		c.JSON(http.StatusOK, commodityLogic.CommodityInfo.SaveOrPublishInfo(draft, middleware.GetUserId(c), 2, false))
	}
	// 之前保存过
	c.JSON(http.StatusOK, commodityLogic.CommodityInfo.UpdateInfo(draft, false))
}

func (*InfoApi) WantPublish(c *gin.Context) {
	var info = new(model.CommodityInfo)
	err := c.ShouldBind(info)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.ErrorMsg("请求参数错误！"))
		return
	}
	// 未保存就发布
	if info.Id == 0 {
		c.JSON(http.StatusOK, commodityLogic.CommodityInfo.SaveOrPublishInfo(info, middleware.GetUserId(c), 2, true))
		return
	}
	// 发布（更新）
	c.JSON(http.StatusOK, commodityLogic.CommodityInfo.UpdateInfo(info, true))
}

func (a *InfoApi) Like(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.ErrorMsg("请求参数错误！"))
		return
	}
	c.JSON(http.StatusOK, commodityLogic.CommodityInfo.Like(id, middleware.GetUserId(c)))
}

func (a *InfoApi) Unlike(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.ErrorMsg("请求参数错误！"))
		return
	}
	c.JSON(http.StatusOK, commodityLogic.CommodityInfo.Unlike(id, middleware.GetUserId(c)))
}
