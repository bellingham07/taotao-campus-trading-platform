package fileLogic

import (
	ossLogic "com.xpdj/go-gin/logic/oss"
	"com.xpdj/go-gin/model"
	"com.xpdj/go-gin/model/response"
	fileRepository "com.xpdj/go-gin/repository/file"
	"github.com/gin-gonic/gin"
	"log"
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
		log.Println("Save Avatar 上传图片失败（OSS上传）" + err.Error())
		return response.GenH(response.FAIL, "上传图片失败！")
	}
	userId, _ := strconv.ParseInt(userIdStr, 10, 64)
	var fileAsset = &model.FileAsset{
		UserId:     userId,
		ObjectName: objectName,
		Url:        url,
		CreateAt:   time.Now(),
	}
	if err := fileRepository.AssetRepository.Insert(fileAsset); err != nil {
		log.Println("Save Avatar 上传图片失败（数据库插入）" + err.Error())
		go ossLogic.OSSClient.Delete(objectName)
		return response.GenH(response.FAIL, "上传图片失败！")
	}
	return response.GenH(response.OK, response.SUCCESS)
}

func (*FileAssetLogic) SavePics(fileHeader []*multipart.FileHeader, userIdStr string, articleIdStr string) gin.H {
	// 1.先存到OSS
	urls, objectNames, err := ossLogic.OSSClient.MultiUpload(fileHeader, userIdStr)
	if err != nil {
		log.Println("Save Pics 上传图片失败（OSS多上传）" + err.Error())
		return response.GenH(response.FAIL, "上传图片失败！")
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
		log.Println("Save Pics 上传图片失败（数据库插入）" + err.Error())
		go ossLogic.OSSClient.MultiDelete(objectNames)
		return response.GenH(response.FAIL, "上传图片失败！")
	}
	return response.GenH(response.OK, response.SUCCESS, urls)
}
