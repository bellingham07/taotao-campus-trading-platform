package logic

import (
	"com.xpwk/go-gin/model"
	"com.xpwk/go-gin/model/request"
	"com.xpwk/go-gin/model/response"
	"com.xpwk/go-gin/repository"
)

var User = new(UserLogic)

type UserLogic struct {
}

func (*UserLogic) Login(loginUser request.LoginUser) *response.Result {
	user := &model.User{
		Username: loginUser.Username,
	}
	userDB := repository.User.QueryByUsername(user)
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
	userDB := repository.User.QueryById(user)
}
