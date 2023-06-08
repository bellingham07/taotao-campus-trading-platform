package fileRepository

import (
	"com.xpdj/go-gin/model"
	"com.xpdj/go-gin/repository"
	"log"
)

var FileCommodity = new(FileCommodityRepository)

type FileCommodityRepository struct {
}

func file_commodity() string {
	return "file_article"
}

func (*FileCommodityRepository) Insert(asset *model.FileCommodity) error {
	if err := repository.GetDB().Table(file_commodity()).Create(asset).Error; err != nil {
		log.Println("[GORM ERROR] FileCommodity Insert Fail, Error: " + err.Error())
		return err
	}
	return nil
}

func (*FileCommodityRepository) MultiInsert(assets *[]model.FileCommodity) error {
	if err := repository.GetDB().Table(file_commodity()).Create(assets).Error; err != nil {
		return err
	}
	return nil
}

func (*FileCommodityRepository) DeleteById(id int64) error {
	if err := repository.GetDB().Table(file_commodity()).Where("id = ?", id).Delete(&model.FileCommodity{}).Error; err != nil {
		log.Println("[GORM ERROR] FileCommodity DeleteById FAIL，请管理员手动删除！Error: ", err.Error())
		return err
	}
	return nil
}

func (*FileCommodityRepository) DeleteByUserIdAndIsCover(userId int64) error {
	if err := repository.GetDB().Table(file_commodity()).Where("user_id = ? AND is_cover = 1", userId).Delete(&model.FileCommodity{}).Error; err != nil {
		log.Println("[GORM ERROR] FileArticle DeleteByUserIdAndIsCover FAIL! Error: ", err.Error())
		return err
	}
	return nil
}

//func (*FileCommodityRepository) DeleteByUrl(fc *model.FileCommodity) error {
//	if err := repository.GetDB().Table(file_commodity()).Where("user_id = ? AND is_cover = ?", fc.UserId, fc.IsCover).Delete(&model.FileAsset{}).Error; err != nil {
//		log.Println("[GORM ERROR] FileCommodity DeleteById FAIL，请管理员手动删除！Error: ", err.Error())
//		return err
//	}
//	return nil
//}
//
//func (*FileCommodityRepository) QueryUrlByUserIdAndType(asset *model.FileCommodity) error {
//	if err := repository.GetDB().Table(file_commodity()).Select("id, url").Where("user_id = ? AND type = ?", asset.UserId, asset.Type).First(asset).Error; err != nil {
//		log.Println("[GORM ERROR] FileCommodity DeleteById FAIL，请管理员手动删除！Error: ", err.Error())
//		return err
//	}
//	return nil
//}
