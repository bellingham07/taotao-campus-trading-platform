package fileApi

import (
	fileLogic "com.xpdj/go-gin/logic/file"
	"com.xpdj/go-gin/model/response"
	"github.com/gin-gonic/gin"
	"mime/multipart"
	"net/http"
)

type InfoApi struct {
}

func (*InfoApi) UploadAvatar(c *gin.Context) {

	fileHeaders, exist := c.Get("files")
	if !exist {
		c.JSON(http.StatusBadRequest, gin.H{"code": response.FAIL, "msg": "文件错误！"})
		return
	}
	files, ok := fileHeaders.([]*multipart.FileHeader)
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{"code": response.FAIL, "msg": "文件错误！"})
		return
	}
	value, _ := c.Get("userid")
	userId := value.(string)
	c.JSON(http.StatusOK, fileLogic.AssetLogic.SaveAvatar(files[0], userId))
}

func (*InfoApi) UploadPics(c *gin.Context) {
	fileHeaders, exist := c.Get("files")
	if !exist {
		c.JSON(http.StatusBadRequest, gin.H{"code": response.FAIL, "msg": "文件错误！"})
		return
	}
	files, ok := fileHeaders.([]*multipart.FileHeader)
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{"code": response.FAIL, "msg": "文件错误！"})
		return
	}
	value, _ := c.Get("userid")
	userId := value.(string)
	articleId := c.PostForm("articleid")
	if articleId == "" {
		c.JSON(http.StatusBadRequest, gin.H{"code": response.FAIL, "msg": "没有指定文章！"})
		return
	}
	c.JSON(http.StatusOK, fileLogic.AssetLogic.SavePics(files, userId, articleId))
}

// TODO
func (*InfoApi) Delete(c *gin.Context) {

}
