package userRepository

import (
	"com.xpwk/go-gin/model"
	"com.xpwk/go-gin/repository"
	"log"
)

var UserInfo = new(UserInfoRepository)

type UserInfoRepository struct {
}

func userInfo() string {
	return "user_info"
}

func (*UserInfoRepository) QueryByUsername(username string) (user model.UserInfo, err error) {
	user.Username = username
	if err := repository.GetDB().Table(userInfo()).Find(&user).Error; err != nil {
		log.Println("QueryById NOT FOUND: " + err.Error())
		return user, err
	}
	return user, nil
}

func (*UserInfoRepository) QueryById(id int64) (user model.UserInfo, err error) {
	user = model.UserInfo{
		Id: id,
	}
	if err := repository.GetDB().Table(userInfo()).First(&user).Error; err != nil {
		log.Println("QueryById NOT FOUND: " + err.Error())
		return user, err
	}
	return user, nil
}
