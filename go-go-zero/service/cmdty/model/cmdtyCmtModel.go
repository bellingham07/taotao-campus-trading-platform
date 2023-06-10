package model

import (
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ CmdtyCmtModel = (*customCmdtyCmtModel)(nil)

type (
	// CmdtyCmtModel is an interface to be customized, add more methods here,
	// and implement the added methods in customCmdtyCmtModel.
	CmdtyCmtModel interface {
		cmdtyCmtModel
		ListByCmdtyId(id int64) []*CmdtyCmt
	}

	customCmdtyCmtModel struct {
		*defaultCmdtyCmtModel
	}
)

// NewCmdtyCmtModel returns a model for the database table.
func NewCmdtyCmtModel(conn sqlx.SqlConn) CmdtyCmtModel {
	return &customCmdtyCmtModel{
		defaultCmdtyCmtModel: newCmdtyCmtModel(conn),
	}
}

func (c *customCmdtyCmtModel) ListByCmdtyId(id int64) (ccs []*CmdtyCmt) {
	ccs = make([]*CmdtyCmt, 0)
	query := "select * from cmdty_cmt where `id` = ? limit 1"
	err := c.conn.QueryRows(ccs, query, id)
	if err != nil {
		logx.Debugf("[DB ERROR] ListByCmdtyId 查询错误 " + err.Error())
		return nil
	}
	return
}
