package userApi

import (
	"com.xpwk/go-gin/logic/user"
	"com.xpwk/go-gin/model/request"
	"com.xpwk/go-gin/model/response"
	"github.com/gin-gonic/gin"
	"log"
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

func (*InfoApi) GetInfoById(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": response.FAIL,
			"msg":  "请求错误",
		})
		return
	}
	c.JSON(http.StatusOK, userLogic.UserInfo.GetUserById(id))
}

func (*InfoApi) UpdateInfo(c *gin.Context) {
	//TODO
}

func (*InfoApi) Register(c *gin.Context) {
	// TODO
	var registerUser = new(request.RegisterUser)
	err := c.ShouldBind(registerUser)
	if err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"code": response.FAIL, "msg": "请求错误！"})
		return
	}
	c.JSON(http.StatusOK, userLogic.UserInfo.Register(registerUser))
}
