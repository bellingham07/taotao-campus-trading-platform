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
		logx.Errorf("[GORM ERROR] DeleteByCmdtyIdAndUserId Fail : %s \n", err.Error())
		return err
	}
	return nil
}
