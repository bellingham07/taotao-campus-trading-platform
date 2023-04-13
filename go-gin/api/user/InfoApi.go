package userApi

import (
	"com.xpwk/go-gin/logic/user"
	"com.xpwk/go-gin/model/request"
	"com.xpwk/go-gin/model/response"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type UserInfoApi struct {
}

func (*UserInfoApi) UserLogin(ctx *gin.Context) {
	var loginUser request.LoginUser
	err := ctx.ShouldBind(&loginUser)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, &response.Result{Msg: "验证码错误"})
	}
	ctx.JSON(http.StatusOK, userLogic.UserInfo.Login(loginUser))
}

func (*UserInfoApi) GetUserInfo(ctx *gin.Context) {
	idStr := ctx.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, &response.Result{
			Code: response.FAIL,
			Msg:  "请求错误",
		})
	}
	ctx.JSON(http.StatusOK, userLogic.UserInfo.GetUserById(id))
}

func (*UserInfoApi) UpdateUserInfo(ctx *gin.Context) {

}

func (*UserInfoApi) Register(ctx *gin.Context) {

	ctx.JSON(http.StatusOK, userLogic.UserInfo.Register())
}
