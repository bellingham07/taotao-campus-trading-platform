package model

import "github.com/zeromicro/go-zero/core/stores/sqlx"

var _ TradeDoneModel = (*customTradeDoneModel)(nil)

type (
	// TradeDoneModel is an interface to be customized, add more methods here,
	// and implement the added methods in customTradeDoneModel.
	TradeDoneModel interface {
		tradeDoneModel
	}

	customTradeDoneModel struct {
		*defaultTradeDoneModel
	}
)

// NewTradeDoneModel returns a model for the database table.
func NewTradeDoneModel(conn sqlx.SqlConn) TradeDoneModel {
	return &customTradeDoneModel{
		defaultTradeDoneModel: newTradeDoneModel(conn),
	}
}
