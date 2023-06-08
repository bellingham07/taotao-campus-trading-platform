package userApi

import (
	userLogic "com.xpdj/go-gin/logic/user"
	"github.com/gin-gonic/gin"
	"net/http"
)

type LocationApi struct {
}

func (*LocationApi) List(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, userLogic.UserLocation.List())
}
