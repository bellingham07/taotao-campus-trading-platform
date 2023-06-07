package model

import "github.com/zeromicro/go-zero/core/stores/sqlx"

var _ UserOptModel = (*customUserOptModel)(nil)

type (
	// UserOptModel is an interface to be customized, add more methods here,
	// and implement the added methods in customUserOptModel.
	UserOptModel interface {
		userOptModel
	}

	customUserOptModel struct {
		*defaultUserOptModel
	}
)

// NewUserOptModel returns a model for the database table.
func NewUserOptModel(conn sqlx.SqlConn) UserOptModel {
	return &customUserOptModel{
		defaultUserOptModel: newUserOptModel(conn),
	}
}
