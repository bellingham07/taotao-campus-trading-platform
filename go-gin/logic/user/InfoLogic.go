package userLogic

import (
	"com.xpdj/go-gin/model"
	"com.xpdj/go-gin/model/request"
	"com.xpdj/go-gin/model/response"
	userRepository "com.xpdj/go-gin/repository/user"
	"com.xpdj/go-gin/utils"
	"com.xpdj/go-gin/utils/cache"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/yitter/idgenerator-go/idgen"
	"golang.org/x/crypto/bcrypt"
	"log"
	"math/rand"
	"strconv"
	"time"
)

var (
	UserInfo = new(UserInfoLogic)
)

type UserInfoLogic struct {
}

func (*UserInfoLogic) Login(loginUser *request.LoginUser) gin.H {
	username := loginUser.Username
	userDB, err := userRepository.UserInfo.QueryByUsername(username)
	err2 := bcrypt.CompareHashAndPassword([]byte(userDB.Password), []byte(loginUser.Password))
	if err != nil || err2 != nil {
		return response.ErrorMsg("账号或密码错误！")
	}
	userStr, _ := json.Marshal(userDB)
	key := cache.UserLogin + strconv.FormatInt(userDB.Id, 10)
	err = cache.RedisUtil.SET2JSON(key, userStr, 7*24*time.Hour)
	if err != nil {
		return response.ErrorMsg("服务器繁忙，请稍后")
	}
	token, _ := utils.GenerateToken(userDB)
	log.Println(token)
	return response.OkMsgData("登录成功😊", token)
}

func (*UserInfoLogic) GetUserById(id int64) gin.H {
	key := cache.UserInfo + strconv.FormatInt(id, 10)
	userStr, _ := cache.RedisUtil.GET(key)
	// 内容为“”，代表数据库中没有
	if userStr == "" {
		return response.ErrorMsg("非法参数")
	}
	// 数据库有
	user, err := userRepository.UserInfo.QueryById(id)
	if err != nil {
		err := cache.RedisUtil.SET2JSON(key, "", 5*time.Minute)
		if err != nil {
			log.Println("GetUserById 保存至redis失败：" + err.Error())
		}
		return response.Error()
	}
	_ = cache.RedisUtil.SET2JSON(key, user, 5*time.Minute)
	_ = json.Unmarshal([]byte(userStr), user)
	return response.OkData(user)
}

func (*UserInfoLogic) Register(userRegister *request.RegisterUser) gin.H {
	password1 := userRegister.Password1
	password2 := userRegister.Password2
	if password1 != password2 {
		return response.ErrorMsg("两次输入的密码不一致！")
	}
	password, err := bcrypt.GenerateFromPassword([]byte(password1), bcrypt.DefaultCost)
	if err != nil {
		return response.ErrorMsg("请求参数错误！")
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
		return response.ErrorMsg("帐号已存在！")
	}
	return response.OkMsg("恭喜你，注册成功。😊")
}

func (*UserInfoLogic) UpdateInfo(info *model.UserInfo) gin.H {
	if err := userRepository.UserInfo.UpdateById(info); err != nil {
		return response.ErrorMsg("更新失败！")
	}
	return response.OkMsg("更新成功！")
}
