package logic

import (
	"com.xpwk/go-gin/cache"
	"com.xpwk/go-gin/model/request"
	"com.xpwk/go-gin/model/response"
	"com.xpwk/go-gin/repository/user"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"log"
	"reflect"
	"strconv"
	"time"
)

var (
	User = new(UserLogic)
)

type UserLogic struct {
}

func (ul UserLogic) isEmpty() bool {
	return reflect.DeepEqual(ul, UserLogic{})
}

func (*UserLogic) Login(loginUser request.LoginUser) gin.H {
	username := loginUser.Username
	password, err := bcrypt.GenerateFromPassword([]byte(loginUser.Password), bcrypt.DefaultCost)
	log.Printf("加密的密码是%v", string(password))

	if err != nil {
		return gin.H{
			"code": response.FAIL,
			"msg":  "非法请求！",
		}
	}
	log.Printf("帐号是%v", loginUser.Username)
	userDB, err := userRepository.UserInfo.QueryByUsername(username)
	err2 := bcrypt.CompareHashAndPassword(password, []byte(userDB.Password))
	if err != nil || err2 != nil {
		return gin.H{
			"code": response.FAIL,
			"msg":  "账号或密码错误！",
		}
	}
	return gin.H{
		"code": 1,
		"msg":  "登录成功",
		"data": "authorization",
	}
}

func (*UserLogic) GetUserById(id int64) gin.H {
	key := cache.USERID + strconv.FormatInt(id, 10)
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
		err := cache.RedisClient.Set(key, "nil", 2*time.Minute)
		if err != nil {
			log.Println("GetUserById 保存至redis失败：" + err.Error())
		}

		return gin.H{
			"code": response.FAIL,
			"msg":  response.ERROR,
		}
	}
	_ = cache.RedisClient.Set(key, user, 2*time.Minute)
	_ = json.Unmarshal([]byte(userStr), &user)
	return gin.H{
		"code": response.OK,
		"msg":  response.SUCCESS,
		"data": user,
	}

}

func (*UserLogic) Register() gin.H {

	return gin.H{
		"code": response.OK,
		"msg":  response.SUCCESS,
		"data": "",
	}
}
