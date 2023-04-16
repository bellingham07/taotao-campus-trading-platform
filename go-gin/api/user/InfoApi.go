package userApi

import (
	"com.xpwk/go-gin/logic/user"
	"com.xpwk/go-gin/model/request"
	"com.xpwk/go-gin/model/response"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type InfoApi struct {
}

func (*InfoApi) UserLogin(ctx *gin.Context) {
	var loginUser request.LoginUser
	_ = ctx.ShouldBind(&loginUser)

	// TODO
	//if err != nil || loginUser.ValidCode == "" {
	//	ctx.JSON(http.StatusBadRequest, gin.H{
	//		"code": response.FAIL,
	//		"msg":  "请输入正确验证码",
	//	})
	//	return
	//}

	ctx.JSON(http.StatusOK, userLogic.UserInfo.Login(loginUser))
}

func (*InfoApi) Logout(ctx *gin.Context) {
	// TODO
}

func (*InfoApi) GetInfoById(ctx *gin.Context) {
	idStr := ctx.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"code": response.FAIL,
			"msg":  "请求错误",
		})
	}
	ctx.JSON(http.StatusOK, userLogic.UserInfo.GetUserById(id))
}

func (*InfoApi) UpdateInfo(ctx *gin.Context) {
	//TODO
}

func (*InfoApi) Register(ctx *gin.Context) {
	// TODO
	ctx.JSON(http.StatusOK, userLogic.UserInfo.Register())
}
