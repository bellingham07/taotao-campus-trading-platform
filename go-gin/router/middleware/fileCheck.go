package middleware

import (
	"com.xpdj/go-gin/model/response"
	"github.com/gin-gonic/gin"
	"net/http"
	"path"
)

var allowExtMap = map[string]bool{
	".jpg":  true,
	".png":  true,
	".gif":  true,
	".jpeg": true,
}

func FileCheck(c *gin.Context) {
	multipartForm, err := c.MultipartForm()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"code": response.FAIL, "msg": "文件错误！"})
		return
	}
	fileMap := multipartForm.File["pics"]
	for _, fileHeader := range fileMap {
		suffix := path.Ext(fileHeader.Filename)
		if _, ok := allowExtMap[suffix]; !ok {
			c.JSON(http.StatusBadRequest, gin.H{"code": response.FAIL, "message": "上传文件格式不支持！"})
			return
		}
	}
	c.Set("files", fileMap)
	c.Next()
}
