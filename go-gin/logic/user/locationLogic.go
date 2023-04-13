package userLogic

import (
	"com.xpwk/go-gin/model/response"
	userRepository "com.xpwk/go-gin/repository/user"
	"github.com/gin-gonic/gin"
)

var UserLocation = new(UserLocationLogic)

type UserLocationLogic struct {
}

func (*UserLocationLogic) List(ctx *gin.Context) gin.H {
	userLocations := userRepository.UserLocation.QueryAll()
	if userLocations == nil {
		return gin.H{
			"code": response.FAIL,
			"msg":  response.ERROR,
		}
	}
	return gin.H{
		"code": response.OK,
		"msg":  response.SUCCESS,
		"data": userLocations,
	}
}
