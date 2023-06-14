package model

import "github.com/zeromicro/go-zero/core/stores/sqlx"

var _ TradeCmtModel = (*customTradeCmtModel)(nil)

type (
	// TradeCmtModel is an interface to be customized, add more methods here,
	// and implement the added methods in customTradeCmtModel.
	TradeCmtModel interface {
		tradeCmtModel
	}

	customTradeCmtModel struct {
		*defaultTradeCmtModel
	}
)

// NewTradeCmtModel returns a model for the database table.
func NewTradeCmtModel(conn sqlx.SqlConn) TradeCmtModel {
	return &customTradeCmtModel{
		defaultTradeCmtModel: newTradeCmtModel(conn),
	}
}
