package userLogic

import (
	"com.xpwk/go-gin/cache"
	"com.xpwk/go-gin/model/request"
	"com.xpwk/go-gin/model/response"
	"com.xpwk/go-gin/repository/user"
	"com.xpwk/go-gin/utils"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"log"
	"reflect"
	"strconv"
	"time"
)

var (
	UserInfo = new(UserInfoLogic)
)

type UserInfoLogic struct {
}

func (ul UserInfoLogic) isEmpty() bool {
	return reflect.DeepEqual(ul, UserInfoLogic{})
}

func (*UserInfoLogic) Login(loginUser request.LoginUser) gin.H {
	username := loginUser.Username
	userDB, err := userRepository.UserInfo.QueryByUsername(username)
	err2 := bcrypt.CompareHashAndPassword([]byte(userDB.Password), []byte(loginUser.Password))
	if err != nil || err2 != nil {
		return gin.H{
			"code": response.FAIL,
			"msg":  "账号或密码错误！",
		}
	}
	userStr, _ := json.Marshal(userDB)
	key := cache.USERLOGIN + strconv.FormatInt(userDB.Id, 10)
	err = cache.RedisClient.Set(key, userStr, 7*24*time.Hour)
	if err != nil {
		return gin.H{
			"code": response.FAIL,
			"msg":  "服务器繁忙，请稍后",
		}
	}
	token, _ := utils.GenerateToken(userDB)
	log.Println(token)
	return gin.H{
		"code": response.OK,
		"msg":  "登录成功",
		"data": token,
	}
}

func (*UserInfoLogic) GetUserById(id int64) gin.H {
	key := cache.USERINFO + strconv.FormatInt(id, 10)
	userStr, _ := cache.RedisClient.Get(key)
	// 内容为“nil”，代表数据库中没有
	if userStr == "nil" {
		return gin.H{
			"code": response.FAIL,
			"msg":  response.ERROR,
		}
	}
	// 数据库有
	user, err := userRepository.UserInfo.QueryById(id)
	if err != nil {
		err := cache.RedisClient.Set(key, "nil", 5*time.Minute)
		if err != nil {
			log.Println("GetUserById 保存至redis失败：" + err.Error())
		}

		return gin.H{
			"code": response.FAIL,
			"msg":  response.ERROR,
		}
	}
	_ = cache.RedisClient.Set(key, user, 5*time.Minute)
	_ = json.Unmarshal([]byte(userStr), &user)
	return gin.H{
		"code": response.OK,
		"msg":  response.SUCCESS,
		"data": user,
	}

}

func (*UserInfoLogic) Register() gin.H {

	return gin.H{
		"code": response.OK,
		"msg":  response.SUCCESS,
		"data": "",
	}
}
