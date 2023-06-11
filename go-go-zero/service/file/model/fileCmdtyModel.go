package model

import (
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ FileCmdtyModel = (*customFileCmdtyModel)(nil)

type (
	// FileCmdtyModel is an interface to be customized, add more methods here,
	// and implement the added methods in customFileCmdtyModel.
	FileCmdtyModel interface {
		fileCmdtyModel
	}

	customFileCmdtyModel struct {
		*defaultFileCmdtyModel
	}
)

// NewFileCmdtyModel returns a model for the database table.
func NewFileCmdtyModel(conn sqlx.SqlConn) FileCmdtyModel {
	return &customFileCmdtyModel{
		defaultFileCmdtyModel: newFileCmdtyModel(conn),
	}
}
