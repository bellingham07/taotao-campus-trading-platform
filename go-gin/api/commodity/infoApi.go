package commodityApi

import (
	commodityLogic "com.xpdj/go-gin/logic/commodity"
	"com.xpdj/go-gin/model/request"
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
		c.JSON(http.StatusBadRequest, gin.H{"code": response.FAIL, "msg": response.ERROR})
		return
	}
	id, _ := strconv.ParseInt(idStr, 10, 64)
	userIdAny, exist := c.Get("userid")
	userIdStr := userIdAny.(string)
	userId, _ := strconv.ParseInt(userIdStr, 10, 64)
	c.JSON(http.StatusOK, commodityLogic.CommodityInfo.GetById(id, userId, exist))
}

func (*InfoApi) ListByOption(c *gin.Context) {
	optionStr := c.Param("option")
	option, _ := strconv.Atoi(optionStr)
	c.JSON(http.StatusOK, commodityLogic.CommodityInfo.RandomListByType(option))
}

func (*InfoApi) SellSave(c *gin.Context) {
	var draft = new(request.CommodityArticleDraft)
	err := c.ShouldBind(draft)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.GenH(response.FAIL, "参数错误！"))
		return
	}
	// 第一次保存
	if draft.Id == 0 {
		c.JSON(http.StatusOK, commodityLogic.CommodityInfo.SaveOrPublishInfoAndArticle(draft, middleware.GetUserId(c), 1, false))
	}
	// 之前保存过
	c.JSON(http.StatusOK, commodityLogic.CommodityInfo.UpdateInfoAndArticle(draft, false))
}

func (*InfoApi) SellPublish(c *gin.Context) {
	var info = new(request.CommodityArticleDraft)
	err := c.ShouldBind(info)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.GenH(response.FAIL, "参数错误！"))
		return
	}
	// 未保存就发布
	if info.Id == 0 {
		c.JSON(http.StatusOK, commodityLogic.CommodityInfo.SaveOrPublishInfoAndArticle(info, middleware.GetUserId(c), 1, true))
		return
	}
	// 发布（更新）
	c.JSON(http.StatusOK, commodityLogic.CommodityInfo.UpdateInfoAndArticle(info, true))
}

func (*InfoApi) WantSave(c *gin.Context) {
	var draft = new(request.CommodityArticleDraft)
	err := c.ShouldBind(draft)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.GenH(response.FAIL, "参数错误！"))
		return
	}
	// 第一次保存
	if draft.Id == 0 {
		c.JSON(http.StatusOK, commodityLogic.CommodityInfo.SaveOrPublishInfoAndArticle(draft, middleware.GetUserId(c), 2, false))
	}
	// 之前保存过
	c.JSON(http.StatusOK, commodityLogic.CommodityInfo.UpdateInfoAndArticle(draft, false))
}

func (*InfoApi) WantPublish(c *gin.Context) {
	var info = new(request.CommodityArticleDraft)
	err := c.ShouldBind(info)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.GenH(response.FAIL, "参数错误！"))
		return
	}
	// 未保存就发布
	if info.Id == 0 {
		c.JSON(http.StatusOK, commodityLogic.CommodityInfo.SaveOrPublishInfoAndArticle(info, middleware.GetUserId(c), 2, true))
		return
	}
	// 发布（更新）
	c.JSON(http.StatusOK, commodityLogic.CommodityInfo.UpdateInfoAndArticle(info, true))
}
