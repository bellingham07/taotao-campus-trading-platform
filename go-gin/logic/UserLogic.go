package logic

import (
	"com.xpwk/go-gin/cache"
	"com.xpwk/go-gin/model/request"
	"com.xpwk/go-gin/model/response"
	"com.xpwk/go-gin/repository"
	"encoding/json"
	"strconv"
)

var (
	User = new(UserLogic)
)

type UserLogic struct {
}

func (*UserLogic) Login(loginUser request.LoginUser) *response.Result {
	username := loginUser.Username

	userDB := repository.User.QueryByUsername(username)
	if loginUser.Password != userDB.Password {
		return &response.Result{
			Code: -1,
			Msg:  "账号或密码错误！",
			Data: "",
		}
	}
	return &response.Result{
		Code: 1,
		Msg:  "登录成功",
		Data: "authorization",
	}
}

func (*UserLogic) GetUserById(id int64) *response.Result {
	key := cache.USERID + strconv.FormatInt(id, 10)
	userStr, _ := cache.RedisClient.Get(key)
	// 内容为“nil”，代表数据库中没有
	if userStr == "nil" {
		return &response.Result{
			Code: response.FAIL,
			Msg:  "内容不存在",
		}
	}
	user := repository.User.QueryById(id)
	if user == nil {
		_ = cache.RedisClient.Set(key, "nil", 30)
	}
	_ = cache.RedisClient.Set(key, user, 30)
	_ = json.Unmarshal([]byte(userStr), user)
	return &response.Result{
		Code: response.OK,
		Msg:  response.SUCCESS,
		Data: user,
	}

}
