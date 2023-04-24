package fileApi

import (
	ossLogic "com.xpwk/go-gin/logic/oss"
	"com.xpwk/go-gin/model/response"
	"github.com/gin-gonic/gin"
	"net/http"
)

type InfoApi struct {
}

func (*InfoApi) Upload(c *gin.Context) {
	multipartForm, err := c.MultipartForm()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"code": response.FAIL, "msg": "文件错误！"})
	}
	file := multipartForm.File
	err = ossLogic.OSSClient.Upload()
}
