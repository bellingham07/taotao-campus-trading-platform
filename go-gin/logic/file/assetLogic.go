package fileLogic

import (
	ossLogic "com.xpwk/go-gin/logic/oss"
	"com.xpwk/go-gin/model"
	"com.xpwk/go-gin/model/response"
	fileRepository "com.xpwk/go-gin/repository/file"
	"github.com/gin-gonic/gin"
	"mime/multipart"
	"strconv"
	"time"
)

var AssetLogic = new(FileAssetLogic)

type FileAssetLogic struct {
}

func (*FileAssetLogic) SaveAvatar(fileHeader *multipart.FileHeader, userIdStr string) gin.H {
	url, objectName, err := ossLogic.OSSClient.Upload(fileHeader, userIdStr)
	if err != nil {
		return gin.H{"code": response.FAIL, "msg": "上传图片失败！"}
	}
	userId, _ := strconv.ParseInt(userIdStr, 10, 64)
	var fileAsset = &model.FileAsset{
		UserId:     userId,
		ObjectName: objectName,
		Url:        url,
		ArticleId:  0,
		CreateAt:   time.Now(),
	}
	if err := fileRepository.AssetRepository.Insert(fileAsset); err != nil {
		return gin.H{"code": response.FAIL, "msg": "上传图片失败！"}
	}
	return gin.H{"code": response.OK, "msg": response.SUCCESS}

}

func (*FileAssetLogic) SavePics(fileHeader []*multipart.FileHeader, userIdStr string, articleIdStr string) gin.H {
	// 1.先存到OSS
	urls, objectNames, err := ossLogic.OSSClient.MultiUpload(fileHeader, userIdStr)
	if err != nil {
		return gin.H{"code": response.FAIL, "msg": "上传图片失败！"}
	}
	userId, _ := strconv.ParseInt(userIdStr, 10, 64)
	articleId, _ := strconv.ParseInt(articleIdStr, 10, 64)
	var fileAssets []model.FileAsset
	for index, url := range urls {
		fileAsset := model.FileAsset{
			UserId:     userId,
			ObjectName: objectNames[index],
			Url:        url,
			ArticleId:  articleId,
			CreateAt:   time.Now(),
		}
		fileAssets = append(fileAssets, fileAsset)
	}
	// 2.存到数据库
	if err := fileRepository.AssetRepository.MultiInsert(&fileAssets); err != nil {
		return gin.H{"code": response.FAIL, "msg": "上传图片失败！"}
	}
	return gin.H{"code": response.OK, "msg": response.SUCCESS, "data": urls}
}
