// Code generated by goctl. DO NOT EDIT.

package model

import (
	"context"
	"database/sql"
	"fmt"
	"strings"
	"time"

	"github.com/zeromicro/go-zero/core/stores/builder"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/core/stringx"
)

var (
	tradeCmtFieldNames          = builder.RawFieldNames(&TradeCmt{})
	tradeCmtRows                = strings.Join(tradeCmtFieldNames, ",")
	tradeCmtRowsExpectAutoSet   = strings.Join(stringx.Remove(tradeCmtFieldNames, "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), ",")
	tradeCmtRowsWithPlaceHolder = strings.Join(stringx.Remove(tradeCmtFieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), "=?,") + "=?"
)

type (
	tradeCmtModel interface {
		Insert(ctx context.Context, data *TradeCmt) (sql.Result, error)
		FindOne(ctx context.Context, id int64) (*TradeCmt, error)
		Update(ctx context.Context, data *TradeCmt) error
		Delete(ctx context.Context, id int64) error
	}

	defaultTradeCmtModel struct {
		conn  sqlx.SqlConn
		table string
	}

	TradeCmt struct {
		Id          int64     `db:"id"` // id
		TradeId     int64     `db:"trade_id"`
		BuyerId     int64     `db:"buyer_id"`
		Buyer       string    `db:"buyer"`
		BuyerCover  string    `db:"buyer_cover"`
		SellerId    int64     `db:"seller_id"`
		Seller      string    `db:"seller"`
		SellerCover string    `db:"seller_cover"`
		Content     string    `db:"content"`   // 评价内容
		Type        int64     `db:"type"`      // 差评或好评，0为差评，1为好评
		CreateAt    time.Time `db:"create_at"` // 创建时间
	}
)

func newTradeCmtModel(conn sqlx.SqlConn) *defaultTradeCmtModel {
	return &defaultTradeCmtModel{
		conn:  conn,
		table: "`trade_cmt`",
	}
}

func (m *defaultTradeCmtModel) Delete(ctx context.Context, id int64) error {
	query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
	_, err := m.conn.ExecCtx(ctx, query, id)
	return err
}

func (m *defaultTradeCmtModel) FindOne(ctx context.Context, id int64) (*TradeCmt, error) {
	query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", tradeCmtRows, m.table)
	var resp TradeCmt
	err := m.conn.QueryRowCtx(ctx, &resp, query, id)
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultTradeCmtModel) Insert(ctx context.Context, data *TradeCmt) (sql.Result, error) {
	query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?, ?, ?, ?, ?, ?, ?)", m.table, tradeCmtRowsExpectAutoSet)
	ret, err := m.conn.ExecCtx(ctx, query, data.Id, data.TradeId, data.BuyerId, data.Buyer, data.BuyerCover, data.SellerId, data.Seller, data.SellerCover, data.Content, data.Type)
	return ret, err
}

func (m *defaultTradeCmtModel) Update(ctx context.Context, data *TradeCmt) error {
	query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, tradeCmtRowsWithPlaceHolder)
	_, err := m.conn.ExecCtx(ctx, query, data.TradeId, data.BuyerId, data.Buyer, data.BuyerCover, data.SellerId, data.Seller, data.SellerCover, data.Content, data.Type, data.Id)
	return err
}

func (m *defaultTradeCmtModel) tableName() string {
	return m.table
}
