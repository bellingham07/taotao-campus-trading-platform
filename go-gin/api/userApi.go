package api

import (
	"com.xpwk/go-gin/logic"
	"com.xpwk/go-gin/model/request"
	"com.xpwk/go-gin/model/response"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type UserApi struct {
}

func (*UserApi) UserLogin(ctx *gin.Context) {
	var loginUser request.LoginUser
	err := ctx.ShouldBind(&loginUser)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, &response.Result{Msg: "验证码错误"})
	}
	ctx.JSON(http.StatusOK, logic.User.Login(loginUser))
}

func (*UserApi) GetInfo(ctx *gin.Context) {
	idStr := ctx.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		ctx.JSON(http.StatusOK, &response.Result{Code: -1, Msg: "没有用户信息"})
	}
	ctx.JSON(http.StatusOK, logic.User.GetUserById(id))
}

func (*UserApi) UpdateInfo(ctx *gin.Context) {

}
