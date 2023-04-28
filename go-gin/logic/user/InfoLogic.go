package userLogic

import (
	"com.xpdj/go-gin/model"
	"com.xpdj/go-gin/model/request"
	"com.xpdj/go-gin/model/response"
	userRepository "com.xpdj/go-gin/repository/user"
	"com.xpdj/go-gin/utils"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/yitter/idgenerator-go/idgen"
	"golang.org/x/crypto/bcrypt"
	"log"
	"math/rand"
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
	key := utils.USERLOGIN + strconv.FormatInt(userDB.Id, 10)
	err = utils.RedisUtil.SET(key, userStr, 7*24*time.Hour)
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
	key := utils.USERINFO + strconv.FormatInt(id, 10)
	userStr, _ := utils.RedisUtil.GET(key)
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
		err := utils.RedisUtil.SET(key, "nil", 5*time.Minute)
		if err != nil {
			log.Println("GetUserById 保存至redis失败：" + err.Error())
		}

		return gin.H{
			"code": response.FAIL,
			"msg":  response.ERROR,
		}
	}
	_ = utils.RedisUtil.SET(key, user, 5*time.Minute)
	_ = json.Unmarshal([]byte(userStr), &user)
	return gin.H{
		"code": response.OK,
		"msg":  response.SUCCESS,
		"data": user,
	}

}

func (*UserInfoLogic) Register(userRegister *request.RegisterUser) gin.H {
	password1 := userRegister.Password1
	password2 := userRegister.Password2
	if password1 != password2 {
		return gin.H{"code": response.FAIL, "msg": "两次输入的密码不一致！"}
	}
	password, err := bcrypt.GenerateFromPassword([]byte(password1), bcrypt.DefaultCost)
	if err != nil {
		return gin.H{"code": response.FAIL, "msg": "请求错误！"}
	}
	id := idgen.NextId()
	randNum := rand.Int31()
	randNumStr := "user" + strconv.Itoa(int(randNum))
	var userInfo = model.UserInfo{
		Id:        id,
		Username:  userRegister.Username,
		Password:  string(password),
		Name:      randNumStr,
		LastLogin: time.Now(),
	}
	err = userRepository.UserInfo.InsertInfoRegister(&userInfo)
	if err != nil {
		return gin.H{"code": response.FAIL, "msg": "帐号已存在！"}
	}
	return gin.H{"code": response.OK, "msg": "恭喜你，注册成功。😊"}
}
