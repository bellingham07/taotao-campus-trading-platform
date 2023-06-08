package model

import "github.com/zeromicro/go-zero/core/stores/sqlx"

var _ UserLocationModel = (*customUserLocationModel)(nil)

type (
	// UserLocationModel is an interface to be customized, add more methods here,
	// and implement the added methods in customUserLocationModel.
	UserLocationModel interface {
		userLocationModel
	}

	customUserLocationModel struct {
		*defaultUserLocationModel
	}
)

// NewUserLocationModel returns a model for the database table.
func NewUserLocationModel(conn sqlx.SqlConn) UserLocationModel {
	return &customUserLocationModel{
		defaultUserLocationModel: newUserLocationModel(conn),
	}
}
