package model

import (
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ CmdtyTagModel = (*customCmdtyTagModel)(nil)

type (
	// CmdtyTagModel is an interface to be customized, add more methods here,
	// and implement the added methods in customCmdtyTagModel.
	CmdtyTagModel interface {
		cmdtyTagModel
		List() []*CmdtyTag
		DeleteByIds(ids []int64) error
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

func (c customCmdtyTagModel) List() (cts []*CmdtyTag) {
	cts = make([]*CmdtyTag, 0)
	query := "select * from cmdty_tag"
	err := c.conn.QueryRows(cts, query)
	if err != nil {
		logx.Debugf("[DB ERROR] List 查询错误 " + err.Error())
		return nil
	}
	return
}

func (c customCmdtyTagModel) DeleteByIds(ids []int64) error {
	stmt := "delete from cmdty_tag where id in (?)"
	_, err := c.conn.Exec(stmt, ids)
	if err != nil {
		logx.Debugf("[DB ERROR] DeleteByIds 删除错误 " + err.Error())
		return err
	}
	return nil
}
