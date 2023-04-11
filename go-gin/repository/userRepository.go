package repository

import (
	"com.xpwk/go-gin/model"
	"log"
)

var User = new(UserRepository)

type UserRepository struct {
}

func (*UserRepository) QueryByUsername(username string) (user *model.User) {
	user = &model.User{
		Username: username,
	}
	if err := getDB().Find(user).Error; err != nil {
		log.Println("QueryById NOT FOUND: " + err.Error())
		return nil
	}
	return user
}

func (*UserRepository) QueryById(id int64) (user *model.User) {
	user = &model.User{
		Id: id,
	}
	if err := getDB().First(user).Error; err != nil {
		log.Println("QueryById NOT FOUND: " + err.Error())
		return nil
	}
	return user
}
