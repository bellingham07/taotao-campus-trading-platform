package model

import (
	"context"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ UserInfoModel = (*customUserInfoModel)(nil)

type (
	// UserInfoModel is an interface to be customized, add more methods here,
	// and implement the added methods in customUserInfoModel.
	UserInfoModel interface {
		userInfoModel
		QueryInfoByUsername(username string) *UserInfo
	}

	customUserInfoModel struct {
		*defaultUserInfoModel
	}
)

// NewUserInfoModel returns a model for the database table.
func NewUserInfoModel(conn sqlx.SqlConn) UserInfoModel {
	return &customUserInfoModel{
		defaultUserInfoModel: newUserInfoModel(conn),
	}
}

func (m *defaultUserInfoModel) QueryInfoByUsername(username string) *UserInfo {
	ui := new(UserInfo)
	query := "select * from user_info where username = ?"
	err := m.conn.QueryRowCtx(context.Background(), ui, query, username)
	if err != nil {
		logx.Debugf("[DB ERROR] : %v\n", err)
		return nil
	}
	return ui
}
