package repository

import (
	"com.xpwk/go-gin/model"
	"log"
)

var User = new(UserRepository)

type UserRepository struct {
}

func (*UserRepository) QueryByUsername(user *model.User) *model.User {
	db := getDB().Find(user)
	if db.Error != nil {
		log.Println("query login users error: ", db.Error)
	}
	return user
}

func (*UserRepository) QueryById(id int64) (user *model.User) {
	db := getDB().Find(user)
	if db.Error != nil {
		log.Println("query user by id error: ", db.Error)
	}
	return user
}
