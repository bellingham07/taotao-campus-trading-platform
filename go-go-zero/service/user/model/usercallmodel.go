package model

import "github.com/zeromicro/go-zero/core/stores/sqlx"

var _ UserCallModel = (*customUserCallModel)(nil)

type (
	// UserCallModel is an interface to be customized, add more methods here,
	// and implement the added methods in customUserCallModel.
	UserCallModel interface {
		userCallModel
	}

	customUserCallModel struct {
		*defaultUserCallModel
	}
)

// NewUserCallModel returns a model for the database table.
func NewUserCallModel(conn sqlx.SqlConn) UserCallModel {
	return &customUserCallModel{
		defaultUserCallModel: newUserCallModel(conn),
	}
}
