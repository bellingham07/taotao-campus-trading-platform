package model

import "github.com/zeromicro/go-zero/core/stores/sqlx"

var _ CmdtyTagModel = (*customCmdtyTagModel)(nil)

type (
	// CmdtyTagModel is an interface to be customized, add more methods here,
	// and implement the added methods in customCmdtyTagModel.
	CmdtyTagModel interface {
		cmdtyTagModel
	}

	customCmdtyTagModel struct {
		*defaultCmdtyTagModel
	}
)

// NewCmdtyTagModel returns a model for the database table.
func NewCmdtyTagModel(conn sqlx.SqlConn) CmdtyTagModel {
	return &customCmdtyTagModel{
		defaultCmdtyTagModel: newCmdtyTagModel(conn),
	}
}
