package middleware

import (
	"com.xpdj/go-gin/model/response"
	"github.com/gin-gonic/gin"
	"log"
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
		log.Println("File Check 上传图片失败（中间件）" + err.Error())
		c.JSON(http.StatusBadRequest, response.GenH(response.FAIL, "文件错误！"))
		c.Abort()
		return
	}
	fileMap := multipartForm.File["pics"]
	for _, fileHeader := range fileMap {
		suffix := path.Ext(fileHeader.Filename)
		if _, ok := allowExtMap[suffix]; !ok {
			c.JSON(http.StatusBadRequest, response.GenH(response.FAIL, "上传文件格式不支持！"))
			c.Abort()
			return
		}
	}
	c.Set("files", fileMap)
	c.Next()
}
