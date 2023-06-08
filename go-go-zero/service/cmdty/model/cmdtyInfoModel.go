package model

import "github.com/zeromicro/go-zero/core/stores/sqlx"

var _ CmdtyInfoModel = (*customCmdtyInfoModel)(nil)

type (
	// CmdtyInfoModel is an interface to be customized, add more methods here,
	// and implement the added methods in customCmdtyInfoModel.
	CmdtyInfoModel interface {
		cmdtyInfoModel
	}

	customCmdtyInfoModel struct {
		*defaultCmdtyInfoModel
	}
)

// NewCmdtyInfoModel returns a model for the database table.
func NewCmdtyInfoModel(conn sqlx.SqlConn) CmdtyInfoModel {
	return &customCmdtyInfoModel{
		defaultCmdtyInfoModel: newCmdtyInfoModel(conn),
	}
}
