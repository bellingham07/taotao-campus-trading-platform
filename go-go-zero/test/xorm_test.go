package test

import (
	"fmt"
	"go-go-zero/common/utils"
	userModel "go-go-zero/service/user/model"
	"testing"
)

func TestUserInfo(t *testing.T) {
	engine := utils.InitXorm("mysql", utils.Mysql{Dsn: "root:123456@tcp(43.143.241.157:3306)/taotao_trading_user?charset=utf8mb4&parseTime=True&loc=Local"})

	s := engine.Table("user_info")
	ui := &userModel.UserInfo{Username: "qjdlk", Password: "sadmsdnfjk"}

	insert, err := s.Insert(ui)
	fmt.Println(insert, err)
}
