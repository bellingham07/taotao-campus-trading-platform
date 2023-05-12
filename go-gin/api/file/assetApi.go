package fileApi

import (
	fileLogic "com.xpdj/go-gin/logic/file"
	"com.xpdj/go-gin/model/response"
	"com.xpdj/go-gin/router/middleware"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type AssetApi struct {
}

func (*AssetApi) UploadAvatar(c *gin.Context) {
	files := middleware.GetFiles(c)
	userId := middleware.GetUserIdStr(c)
	c.JSON(http.StatusOK, fileLogic.AssetLogic.SaveAvatar(files[0], userId, middleware.GetAvatar(c)))
}

func (*AssetApi) UploadPics(c *gin.Context) {
	files := middleware.GetFiles(c)
	userId := middleware.GetUserIdStr(c)
	cmdtyIdStr := c.PostForm("commodityId")
	if cmdtyIdStr != "" {
		cmdtyId, err := strconv.ParseInt(cmdtyIdStr, 10, 64)
		if err != nil {
			c.JSON(http.StatusOK, response.ErrorMsg("请求参数错误！🤡"))
			return
		}
		c.JSON(http.StatusOK, fileLogic.AssetLogic.SavePics(files, userId, cmdtyId, false))
		return
	}
	articleIdStr := c.PostForm("articleId")
	articleId, err := strconv.ParseInt(articleIdStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusOK, response.ErrorMsg("请求参数错误！🤡"))
		return
	}
	c.JSON(http.StatusOK, fileLogic.AssetLogic.SavePics(files, userId, articleId, true))
}

func (*AssetApi) UploadCover(c *gin.Context) {
	files := middleware.GetFiles(c)
	userId := middleware.GetUserId(c)
	cmdtyIdStr := c.PostForm("commodityId")
	if cmdtyIdStr != "" {
		cmdtyId, err := strconv.ParseInt(cmdtyIdStr, 10, 64)
		if err != nil {
			c.JSON(http.StatusOK, response.ErrorMsg("请求参数错误！🤡"))
			return
		}
		c.JSON(http.StatusOK, fileLogic.AssetLogic.SaveCover(files[0], userId, cmdtyId, false))
		return
	}
	articleIdStr := c.PostForm("articleId")
	articleId, err := strconv.ParseInt(articleIdStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusOK, response.ErrorMsg("请求参数错误！🤡"))
		return
	}
	c.JSON(http.StatusOK, fileLogic.AssetLogic.SaveCover(files[0], userId, articleId, true))
}
