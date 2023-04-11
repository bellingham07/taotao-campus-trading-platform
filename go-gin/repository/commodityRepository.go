package repository

import (
	"com.xpwk/go-gin/model"
	"log"
)

var Commodity = new(CommodityRepository)

type CommodityRepository struct {
}

func (*CommodityRepository) QueryById(user *model.User) *model.User {
	db := getDB().Find(user)
	if db.Error != nil {
		log.Println("query login users error: ", db.Error)
	}
	return user
}
