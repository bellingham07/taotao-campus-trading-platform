package fileRepository

import (
	"com.xpdj/go-gin/model"
	"com.xpdj/go-gin/repository"
	"log"
)

var FileArticle = new(FileArticleRepository)

type FileArticleRepository struct {
}

func file_article() string {
	return "file_article"
}

func (*FileArticleRepository) Insert(asset *model.FileArticle) error {
	if err := repository.GetDB().Table(file_article()).Create(asset).Error; err != nil {
		log.Println("[GORM ERROR] FileArticle Insert Fail, Error: " + err.Error())
		return err
	}
	return nil
}

func (*FileArticleRepository) MultiInsert(assets *[]model.FileArticle) error {
	if err := repository.GetDB().Table(file_article()).Create(assets).Error; err != nil {
		return err
	}
	return nil
}

func (*FileArticleRepository) DeleteById(id int64) error {
	if err := repository.GetDB().Table(file_article()).Where("id = ?", id).Delete(&model.FileArticle{}).Error; err != nil {
		log.Println("[GORM ERROR] FileArticle DeleteById FAIL，请管理员手动删除! Error: ", err.Error())
		return err
	}
	return nil
}

func (*FileArticleRepository) DeleteByUserIdAndIsCover(userId int64) error {
	if err := repository.GetDB().Table(file_article()).Where("user_id = ? AND is_cover = 1", userId).Delete(&model.FileArticle{}).Error; err != nil {
		log.Println("[GORM ERROR] FileArticle DeleteByUserIdAndIsCover FAIL! Error: ", err.Error())
		return err
	}
	return nil
}

//func (*FileArticleRepository) DeleteByUrl(fa *model.FileArticle) error {
//	if err := repository.GetDB().Table(file_article()).Where("user_id = ? AND type = ?", fa.UserId, fa.IsCover).Delete(&model.FileAsset{}).Error; err != nil {
//		log.Println("[GORM ERROR] FileArticle DeleteById FAIL，请管理员手动删除! Error: ", err.Error())
//		return err
//	}
//	return nil
//}
//
//func (*FileArticleRepository) QueryUrlByUserIdAndType(asset *model.FileArticle) error {
//	if err := repository.GetDB().Table(file_article()).Select("id, url").Where("user_id = ? AND type = ?", asset.UserId, asset.Type).First(asset).Error; err != nil {
//		log.Println("[GORM ERROR] FileArticle DeleteById FAIL，请管理员手动删除! Error: ", err.Error())
//		return err
//	}
//	return nil
//}
