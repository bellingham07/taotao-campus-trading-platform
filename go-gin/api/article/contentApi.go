package articleApi

import (
	articleLogic "com.xpdj/go-gin/logic/article"
	"com.xpdj/go-gin/model"
	"com.xpdj/go-gin/model/response"
	"com.xpdj/go-gin/router/middleware"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type ContentApi struct {
}

func (*ContentApi) Publish(c *gin.Context) {
	var draft = new(model.ArticleContent)
	err := c.ShouldBind(draft)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.ErrorMsg("请求参数错误！"))
		return
	}
	// 直接发布
	if draft.Id == 0 {
		c.JSON(http.StatusOK, articleLogic.ContentLogic.SavaOrPublish(draft, middleware.GetUserId(c), true))
	}
	// 更新已发布的
	c.JSON(http.StatusOK, articleLogic.ContentLogic.Update(draft, true))
}

func (*ContentApi) Save(c *gin.Context) {
	var draft = new(model.ArticleContent)
	err := c.ShouldBind(draft)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.ErrorMsg("请求参数错误！"))
		return
	}
	if draft.Id == 0 {
		c.JSON(http.StatusOK, articleLogic.ContentLogic.SavaOrPublish(draft, middleware.GetUserId(c), false))
	}
	// 更新已保存的
	c.JSON(http.StatusOK, articleLogic.ContentLogic.Update(draft, false))
}

func (*ContentApi) GetById(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.ErrorMsg("请求参数错误！"))
		return
	}
	userIdAny, isLogin := c.Get("userid")
	var userIdStr = ""
	if isLogin {
		userIdStr = userIdAny.(string)
	}
	userId, _ := strconv.ParseInt(userIdStr, 10, 64)
	c.JSON(http.StatusOK, articleLogic.ContentLogic.GetById(id, userId))

}

func (a *ContentApi) Like(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.ErrorMsg("请求参数错误！"))
		return
	}
	c.JSON(http.StatusOK, articleLogic.ContentLogic.Like(id, middleware.GetUserId(c)))
}

func (a *ContentApi) Unlike(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.ErrorMsg("请求参数错误！"))
		return
	}
	c.JSON(http.StatusOK, articleLogic.ContentLogic.Unlike(id, middleware.GetUserId(c)))
}
