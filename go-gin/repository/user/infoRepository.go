package userRepository

import (
	"com.xpwk/go-gin/model"
	"com.xpwk/go-gin/repository"
	"log"
)

var UserInfo = new(UserInfoRepository)

type UserInfoRepository struct {
}

func (*UserInfoRepository) tableName() string {
	return "user_info"
}

func (*UserInfoRepository) QueryByUsername(username string) (userInfo model.UserInfo, err error) {
	userInfo.Username = username
	if err := repository.GetDB().Table(UserInfo.tableName()).Find(&userInfo).Error; err != nil {
		log.Println("[GORM-WRONG] QueryById NOT FOUND: " + err.Error())
		return userInfo, err
	}
	return userInfo, nil
}

func (*UserInfoRepository) QueryById(id int64) (user model.UserInfo, err error) {
	user = model.UserInfo{
		Id: id,
	}
	if err := repository.GetDB().Table(UserInfo.tableName()).First(&user).Error; err != nil {
		log.Println("[GORM-WRONG] QueryById NOT FOUND: " + err.Error())
		return user, err
	}
	return user, nil
}
