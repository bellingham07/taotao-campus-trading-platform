package model

import "github.com/zeromicro/go-zero/core/stores/sqlx"

var _ CmdtyCollectModel = (*customCmdtyCollectModel)(nil)

type (
	// CmdtyCollectModel is an interface to be customized, add more methods here,
	// and implement the added methods in customCmdtyCollectModel.
	CmdtyCollectModel interface {
		cmdtyCollectModel
	}

	customCmdtyCollectModel struct {
		*defaultCmdtyCollectModel
	}
)

// NewCmdtyCollectModel returns a model for the database table.
func NewCmdtyCollectModel(conn sqlx.SqlConn) CmdtyCollectModel {
	return &customCmdtyCollectModel{
		defaultCmdtyCollectModel: newCmdtyCollectModel(conn),
	}
}
