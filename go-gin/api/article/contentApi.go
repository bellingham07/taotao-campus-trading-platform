package articleApi

import (
	articleLogic "com.xpdj/go-gin/logic/article"
	"com.xpdj/go-gin/model"
	"com.xpdj/go-gin/model/response"
	"com.xpdj/go-gin/router/middleware"
	"github.com/gin-gonic/gin"
	"net/http"
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
	userIdAny, exists := c.Get("userid")
	if !exists {

	}
}
