package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type UserApi struct {
}

func (*UserApi) UserLogin(ctx *gin.Context) {

}

func (*UserApi) GetInfo(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, "asddsa")
}

func (*UserApi) UpdateInfo(ctx *gin.Context) {

}
