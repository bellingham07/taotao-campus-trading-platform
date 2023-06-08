package model

import (
	"context"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ CmdtyCollectModel = (*customCmdtyCollectModel)(nil)

type (
	// CmdtyCollectModel is an interface to be customized, add more methods here,
	// and implement the added methods in customCmdtyCollectModel.
	CmdtyCollectModel interface {
		cmdtyCollectModel
		DeleteByCmdtyIdAndUserId(id int64, userId int64) error
		ListByUserId(userId int64) []*CmdtyCollect
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

func (c *customCmdtyCollectModel) DeleteByCmdtyIdAndUserId(cmdtyId int64, userId int64) error {
	query := "delete from cmdty_collect where cmdty_id = ? AND user_id = ?"
	_, err := c.conn.ExecCtx(context.Background(), query, cmdtyId, userId)
	if err != nil {
		logx.Errorf("[GO-ZERO ORM ERROR] DeleteByCmdtyIdAndUserId Fail : %s \n", err.Error())
		return err
	}
	return nil
}

func (c *customCmdtyCollectModel) ListByUserId(userId int64) []*CmdtyCollect {
	cc := make([]*CmdtyCollect, 0)
	query := "select * from cmdty_collect where user_id = ?"
	err := c.conn.QueryRowCtx(context.Background(), cc, query, userId)
	if err != nil {
		logx.Errorf("[GO-ZERO ORM ERROR] ListByUserId Fail : %s \n", err.Error())
		return nil
	}
	return cc
}
