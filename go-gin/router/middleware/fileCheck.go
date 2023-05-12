package middleware

import (
	"com.xpdj/go-gin/model/response"
	"github.com/gin-gonic/gin"
	"log"
	"mime/multipart"
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
		c.JSON(http.StatusBadRequest, response.ErrorMsg("文件错误！"))
		c.Abort()
		return
	}
	fileMap := multipartForm.File["pics"]
	for _, fileHeader := range fileMap {
		suffix := path.Ext(fileHeader.Filename)
		if _, ok := allowExtMap[suffix]; !ok {
			c.JSON(http.StatusBadRequest, response.ErrorMsg("上传文件格式不支持！🥲你只能上传.jpg，.png，.gif，.jpeg格式的图片"))
			c.Abort()
			return
		}
	}
	c.Set("files", fileMap)
	c.Next()
}

func GetFiles(c *gin.Context) []*multipart.FileHeader {
	fileHeaders, exist := c.Get("files")
	if !exist {
		c.JSON(http.StatusBadRequest, response.ErrorMsg("文件错误！"))
		c.Abort()
	}
	files, ok := fileHeaders.([]*multipart.FileHeader)
	if !ok {
		c.JSON(http.StatusBadRequest, response.ErrorMsg("文件错误！"))
		c.Abort()
	}
	return files
}
