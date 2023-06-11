package model

import "github.com/zeromicro/go-zero/core/stores/sqlx"

var _ FileAtclModel = (*customFileAtclModel)(nil)

type (
	// FileAtclModel is an interface to be customized, add more methods here,
	// and implement the added methods in customFileAtclModel.
	FileAtclModel interface {
		fileAtclModel
	}

	customFileAtclModel struct {
		*defaultFileAtclModel
	}
)

// NewFileAtclModel returns a model for the database table.
func NewFileAtclModel(conn sqlx.SqlConn) FileAtclModel {
	return &customFileAtclModel{
		defaultFileAtclModel: newFileAtclModel(conn),
	}
}
