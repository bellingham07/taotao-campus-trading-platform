package userApi

import (
	userLogic "com.xpdj/go-gin/logic/user"
	"com.xpdj/go-gin/model"
	"com.xpdj/go-gin/model/request"
	"com.xpdj/go-gin/model/response"
	"com.xpdj/go-gin/router/middleware"
	"com.xpdj/go-gin/utils"
	"com.xpdj/go-gin/utils/assist"
	"com.xpdj/go-gin/utils/cache"
	"context"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
	"time"
)

type InfoApi struct {
}

// GenerateLoginCode 生成验证码
// 用户点击发送验证码 将验证码存入到redis 过期时间为5min
func (*InfoApi) GenerateLoginCode(c *gin.Context) {
	rds := cache.RedisUtil
	rctx := context.Background()
	// 1 获取用户信息
	var loginUser = new(request.LoginUser)
	_ = c.ShouldBind(loginUser)

	if loginUser.Username == "" {
		c.JSON(500, gin.H{"msg": "用户名为空"})
		return
	}
	// 2 生成验证码
	code := assist.GenerateCode()
	// 3 存入redis 设置过期时间
	key := cache.UserLoginCode + loginUser.Username
	rds.Client.Set(rctx, key, code, 5*time.Minute)
}

// UserLoginWithCode 验证码登录
func (*InfoApi) UserLoginWithCode(c *gin.Context) {
	rds := cache.RedisUtil
	rctx := context.Background()
	// 1 获取用户信息
	var loginUser = new(request.LoginUser)
	_ = c.ShouldBind(loginUser)
	if loginUser.Username == "" {
		response.ErrorMsg("用户名为空")
	}
	// 2 校验验证码
	key := cache.UserLoginCode + loginUser.Username
	if loginUser.ValidCode != rds.Client.Get(rctx, key).Val() {
		response.ErrorMsg("验证码错误")
	}
	// 3 生成token
	userInfo := model.UserInfo{
		Username: loginUser.Username,
		Password: loginUser.Password,
	}
	token, err := utils.GenerateToken(&userInfo)
	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, gin.H{"token": token})
}

func (*InfoApi) UserLogin(c *gin.Context) {
	var loginUser = new(request.LoginUser)
	_ = c.ShouldBind(loginUser)

	// TODO
	//if err != nil || loginUser.ValidCode == "" {
	//	c.JSON(http.StatusBadRequest, gin.H{
	//		"code": response.ERROR,
	//		"msg":  "请输入正确验证码",
	//	})
	//	return
	//}
	c.JSON(http.StatusOK, userLogic.UserInfo.Login(loginUser))
}

func (*InfoApi) Logout(c *gin.Context) {
	userId := middleware.GetUserIdStr(c)
	key := cache.UserLogin + userId
	_ = cache.RedisUtil.DEL(key)
	c.JSON(http.StatusOK, response.OkMsg("期待下一次遇见！😊"))

}

func (*InfoApi) GetInfoById(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.ErrorMsg("请求参数错误！"))
		return
	}
	c.JSON(http.StatusOK, userLogic.UserInfo.GetUserById(id))
}

func (*InfoApi) UpdateInfo(c *gin.Context) {
	info := new(model.UserInfo)
	if err := c.ShouldBind(info); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, response.ErrorMsg("请求参数错误！"))
		return
	}
	userId := middleware.GetUserId(c)
	info.Id = userId
	c.JSON(http.StatusOK, userLogic.UserInfo.UpdateInfo(info))
}

func (*InfoApi) Register(c *gin.Context) {
	// TODO
	var registerUser = new(request.RegisterUser)
	err := c.ShouldBind(registerUser)
	if err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, response.ErrorMsg("请求参数错误！"))
		return
	}
	c.JSON(http.StatusOK, userLogic.UserInfo.Register(registerUser))
}
