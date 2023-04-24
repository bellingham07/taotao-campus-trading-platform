package fileApi

import (
	ossLogic "com.xpwk/go-gin/logic/oss"
	"com.xpwk/go-gin/model/response"
	"fmt"
	"github.com/gin-gonic/gin"
	"mime/multipart"
	"net/http"
)

type InfoApi struct {
}

func (*InfoApi) Upload(c *gin.Context) {

	fileHeaders, exist := c.Get("files")
	if !exist {
		c.JSON(http.StatusBadRequest, gin.H{"code": response.FAIL, "msg": "文件错误！"})
	}
	files, ok := fileHeaders.([]*multipart.FileHeader)
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{"code": response.FAIL, "msg": "文件错误！"})
	}
	value, _ := c.Get("userid")
	userId := value.(string)
	for _, file := range files {
		url, err := ossLogic.OSSClient.Upload(file, userId)
		fmt.Println(url, err)

	}
}
