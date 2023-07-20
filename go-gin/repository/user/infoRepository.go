package userRepository

import (
	"com.xpdj/go-gin/model"
	"com.xpdj/go-gin/repository"
	"log"
)

var UserInfo = new(UserInfoRepository)

type UserInfoRepository struct {
}

func user_info() string {
	return "user_info"
}

func (*UserInfoRepository) QueryByUsername(username string) (*model.UserInfo, error) {
	var userInfo = &model.UserInfo{Username: username}
	if err := repository.GetDB().Table(user_info()).Find(&userInfo).Error; err != nil {
		log.Println("[GORM ERROR] UserInfo QueryById NOT FOUND: " + err.Error())
		return userInfo, err
	}
	return userInfo, nil
}

func (*UserInfoRepository) QueryById(id int64) (user *model.UserInfo, err error) {
	user = &model.UserInfo{
		Id: id,
	}
	if err := repository.GetDB().Table(user_info()).First(user).Error; err != nil {
		log.Println("[GORM ERROR] UserInfo QueryById NOT FOUND: " + err.Error())
		return user, err
	}
	return user, nil
}

func (*UserInfoRepository) InsertInfoRegister(userInfo *model.UserInfo) error {
	if err := repository.GetDB().Table(user_info()).Create(&userInfo).Error; err != nil {
		log.Println("[GORM ERROR] UserInfo InsertInfoRegister Duplicate Key: " + err.Error())
		return err
	}
	return nil
}

func (*UserInfoRepository) UpdateById(info *model.UserInfo) error {
	if err := repository.GetDB().Table(user_info()).Updates(info).Error; err != nil {
		log.Println("[GORM ERROR] UserInfo UpdateById UpdateById Fail, Error: " + err.Error())
		return err
	}
	return nil
}

func (*UserInfoRepository) UpdateLikeById(id int64) error {
	log.Println(123123)
	if err := repository.GetDB().Raw("update ? set like = like + 1 where id = ?", user_info(), id).Error; err != nil {
		log.Println("[GORM ERROR] UserInfo UpdateLikeById Fail, Error: " + err.Error())
		return err
	}
	return nil
}
