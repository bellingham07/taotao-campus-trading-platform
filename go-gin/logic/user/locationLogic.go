package userLogic

import (
	"com.xpdj/go-gin/model/response"
	userRepository "com.xpdj/go-gin/repository/user"
	"com.xpdj/go-gin/utils"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
)

var UserLocation = new(UserLocationLogic)

type UserLocationLogic struct {
}

func (*UserLocationLogic) List() gin.H {
	userLocationsStr, err := utils.RedisUtil.GET(utils.USERLOCATION)
	if err == redis.Nil {
		userLocations := userRepository.UserLocation.QueryAll()
		if userLocations == nil {
			return gin.H{"code": response.FAIL, "msg": response.ERROR}
		}
		userLocationsStr, _ := json.Marshal(userLocations)
		_ = utils.RedisUtil.SET(utils.USERLOCATION, userLocationsStr, 0)
		return gin.H{"code": response.OK, "msg": response.SUCCESS, "data": userLocations}
	}
	userLocations := json.Unmarshal([]byte(userLocationsStr), &userLocationsStr)
	return gin.H{"code": response.OK, "msg": response.SUCCESS, "data": userLocations}
}
