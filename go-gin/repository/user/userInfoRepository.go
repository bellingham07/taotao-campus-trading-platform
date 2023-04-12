package userRepository

import (
	"com.xpwk/go-gin/model"
	"com.xpwk/go-gin/repository"
	"log"
)

var UserInfo = new(UserInfoRepository)

type UserInfoRepository struct {
}

func tableName() string {
	return "user_info"
}

func (*UserInfoRepository) QueryByUsername(username string) (user model.User, err error) {
	user.Username = username
	if err := repository.GetDB().Table(tableName()).Find(&user).Error; err != nil {
		log.Println("QueryById NOT FOUND: " + err.Error())
		return user, err
	}
	return user, nil
}

func (*UserInfoRepository) QueryById(id int64) (user model.User, err error) {
	user = model.User{
		Id: id,
	}
	if err := repository.GetDB().Table(tableName()).First(&user).Error; err != nil {
		log.Println("QueryById NOT FOUND: " + err.Error())
		return user, err
	}
	return user, nil
}
