package userRepository

import (
	"com.xpdj/go-gin/model"
	"com.xpdj/go-gin/repository"
	"log"
)

var UserLocation = new(UserLocationRepository)

type UserLocationRepository struct {
}

func user_location() string {
	return "user_location"
}

func (*UserLocationRepository) QueryAll() (locations []model.UserLocation) {
	if err := repository.GetDB().Table(user_location()).Find(&locations).Error; err != nil {
		log.Println("QueryAll User Location Fail：" + err.Error())
		return nil
	}
	return locations
}
