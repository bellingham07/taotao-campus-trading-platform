package model

import "github.com/zeromicro/go-zero/core/stores/sqlx"

var _ UserDormModel = (*customUserDormModel)(nil)

type (
	// UserDormModel is an interface to be customized, add more methods here,
	// and implement the added methods in customUserDormModel.
	UserDormModel interface {
		userDormModel
	}

	customUserDormModel struct {
		*defaultUserDormModel
	}
)

// NewUserDormModel returns a model for the database table.
func NewUserDormModel(conn sqlx.SqlConn) UserDormModel {
	return &customUserDormModel{
		defaultUserDormModel: newUserDormModel(conn),
	}
}
