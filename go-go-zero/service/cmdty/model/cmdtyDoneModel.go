package model

import "github.com/zeromicro/go-zero/core/stores/sqlx"

var _ CmdtyDoneModel = (*customCmdtyDoneModel)(nil)

type (
	// CmdtyDoneModel is an interface to be customized, add more methods here,
	// and implement the added methods in customCmdtyDoneModel.
	CmdtyDoneModel interface {
		cmdtyDoneModel
	}

	customCmdtyDoneModel struct {
		*defaultCmdtyDoneModel
	}
)

// NewCmdtyDoneModel returns a model for the database table.
func NewCmdtyDoneModel(conn sqlx.SqlConn) CmdtyDoneModel {
	return &customCmdtyDoneModel{
		defaultCmdtyDoneModel: newCmdtyDoneModel(conn),
	}
}
