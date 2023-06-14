package model

import "github.com/zeromicro/go-zero/core/stores/sqlx"

var _ TradeInfoModel = (*customTradeInfoModel)(nil)

type (
	// TradeInfoModel is an interface to be customized, add more methods here,
	// and implement the added methods in customTradeInfoModel.
	TradeInfoModel interface {
		tradeInfoModel
	}

	customTradeInfoModel struct {
		*defaultTradeInfoModel
	}
)

// NewTradeInfoModel returns a model for the database table.
func NewTradeInfoModel(conn sqlx.SqlConn) TradeInfoModel {
	return &customTradeInfoModel{
		defaultTradeInfoModel: newTradeInfoModel(conn),
	}
}
