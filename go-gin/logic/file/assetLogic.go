package fileLogic

import (
	ossLogic "com.xpdj/go-gin/logic/oss"
	"com.xpdj/go-gin/model"
	"com.xpdj/go-gin/model/response"
	"com.xpdj/go-gin/repository"
	articleRepository "com.xpdj/go-gin/repository/article"
	fileRepository "com.xpdj/go-gin/repository/file"
	userRepository "com.xpdj/go-gin/repository/user"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"log"
	"mime/multipart"
	"strconv"
	"strings"
	"time"
)

var AssetLogic = new(FileAssetLogic)

type FileAssetLogic struct {
}

func (*FileAssetLogic) SaveAvatar(fileHeader *multipart.FileHeader, userIdStr, oldUrl string) gin.H {
	// 1.先上传到OSS
	url, _, err := ossLogic.OSSClient.Upload(fileHeader, userIdStr)
	if err != nil {
		log.Println("Save Avatar 上传图片失败（OSS上传）" + err.Error())
		return response.ErrorMsg("上传图片失败！")
	}
	userId, _ := strconv.ParseInt(userIdStr, 10, 64)
	// 2.OSS上传成功，就先更新用户信息中的头像
	err = repository.GetDB().Transaction(func(tx *gorm.DB) error {
		// 2.1.先更新用户信息表
		userInfo := &model.UserInfo{
			Id:     userId,
			Avatar: url,
		}
		if err = userRepository.UserInfo.UpdateById(userInfo); err != nil {
			// 2.2.用户信息表的头像更新失败，我们就删除刚刚那张图片
			index := strings.Index(oldUrl, "images")
			go ossLogic.OSSClient.Delete(oldUrl[index:])
			return err
		}
		return nil
	})
	if err != nil {
		return response.ErrorMsg("上传图片失败！")
	}
	return response.OkMsg("头像更新成功")
}

func (*FileAssetLogic) SavePics(fileHeader []*multipart.FileHeader, userIdStr string, id int64, isArticle bool) gin.H {
	// 1.先存到OSS
	urls, objectNames, err := ossLogic.OSSClient.MultiUpload(fileHeader, userIdStr)
	if err != nil {
		log.Println("Save Pics 上传图片失败（OSS多上传）" + err.Error())
		return response.ErrorMsg("上传图片失败！")
	}
	userId, _ := strconv.ParseInt(userIdStr, 10, 64)
	if isArticle {
		var fa []model.FileArticle
		for index, url := range urls {
			fileAsset := model.FileArticle{
				UserId:     userId,
				ObjectName: objectNames[index],
				Url:        url,
				CreateAt:   time.Now(),
			}
			fa = append(fa, fileAsset)
		}
		// 2.存到数据库
		if err = fileRepository.FileArticle.MultiInsert(&fa); err != nil {
			log.Println("Save Pics 上传图片失败（数据库插入）" + err.Error())
			go ossLogic.OSSClient.MultiDelete(objectNames)
			return response.ErrorMsg("上传图片失败！")
		}
		return response.OkData(urls)
	} else {
		var fc []model.FileCommodity
		for index, url := range urls {
			fileAsset := model.FileCommodity{
				UserId:     userId,
				ObjectName: objectNames[index],
				Url:        url,
				CreateAt:   time.Now(),
			}
			fc = append(fc, fileAsset)
		}
		// 2.存到数据库
		if err = fileRepository.FileCommodity.MultiInsert(&fc); err != nil {
			log.Println("Save Pics 上传图片失败（数据库插入）" + err.Error())
			go ossLogic.OSSClient.MultiDelete(objectNames)
			return response.ErrorMsg("上传图片失败！")
		}
		return response.OkData(urls)
	}
}

func (*FileAssetLogic) SaveCover(fileHeader *multipart.FileHeader, userId, id int64, isArticle bool) gin.H {
	// 1.先存到OSS
	url, objectName, err := ossLogic.OSSClient.Upload(fileHeader, strconv.FormatInt(userId, 10))
	if err != nil {
		log.Println("Save Cover 上传图片失败（OSS多上传）" + err.Error())
		return response.ErrorMsg("上传图片失败！")
	}
	// 2.存到数据库
	// 2.1 如果是插入到article相关的表中
	if isArticle {
		err = repository.GetDB().Transaction(func(tx *gorm.DB) error {
			// 2.2 先删之前的cover
			err = fileRepository.FileArticle.DeleteByUserIdAndIsCover(userId)
			if err != nil {
				return err
			}
			// 2.3.先插入至file_article表中
			fc := &model.FileArticle{
				UserId:     userId,
				ObjectName: objectName,
				Url:        url,
				CreateAt:   time.Now(),
			}
			if err = fileRepository.FileArticle.Insert(fc); err != nil {
				log.Println("Save Cover 上传图片失败（数据库file_asset插入）" + err.Error())
				return err
			}
			// 2.4.更新file_article表的cover字段
			articleContent := &model.ArticleContent{
				Id:    id,
				Cover: url,
			}
			if err = articleRepository.ArticleContent.UpdateById(articleContent); err != nil {
				return err
			}
			return nil
		})
		if err != nil {
			go ossLogic.OSSClient.Delete(objectName)
			return response.ErrorMsg("上传图片失败！")
		}
		return response.OkMsg("上传成功😊")
	} else {
		err = repository.GetDB().Transaction(func(tx *gorm.DB) error {
			// 3.2 先删之前的cover
			err = fileRepository.FileCommodity.DeleteByUserIdAndIsCover(userId)
			if err != nil {
				return err
			}
			// 3.3.先插入至file_commodity表中
			fc := &model.FileCommodity{
				UserId:     userId,
				ObjectName: objectName,
				Url:        url,
				CreateAt:   time.Now(),
			}
			if err = fileRepository.FileCommodity.Insert(fc); err != nil {
				log.Println("Save Cover 上传图片失败（数据库file_asset插入）" + err.Error())
				return err
			}
			// 3.4.更新file_commodity表的cover字段
			articleContent := &model.ArticleContent{
				Id:    id,
				Cover: url,
			}
			if err = articleRepository.ArticleContent.UpdateById(articleContent); err != nil {
				return err
			}
			return nil
		})
		if err != nil {
			go ossLogic.OSSClient.Delete(objectName)
			return response.ErrorMsg("上传图片失败！")
		}

		return response.OkMsg("上传成功😊")
	}
}
