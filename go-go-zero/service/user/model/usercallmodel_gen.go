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
	userCallFieldNames          = builder.RawFieldNames(&UserCall{})
	userCallRows                = strings.Join(userCallFieldNames, ",")
	userCallRowsExpectAutoSet   = strings.Join(stringx.Remove(userCallFieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), ",")
	userCallRowsWithPlaceHolder = strings.Join(stringx.Remove(userCallFieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), "=?,") + "=?"
)

type (
	userCallModel interface {
		Insert(ctx context.Context, data *UserCall) (sql.Result, error)
		FindOne(ctx context.Context, id int64) (*UserCall, error)
		Update(ctx context.Context, data *UserCall) error
		Delete(ctx context.Context, id int64) error
	}

	defaultUserCallModel struct {
		conn  sqlx.SqlConn
		table string
	}

	UserCall struct {
		Id       int64     `db:"id"`        // id
		Name     string    `db:"name"`      // 称号名字
		CreateBy time.Time `db:"create_by"` // 管理员的id
		CreateAt time.Time `db:"create_at"` // 创建时间
		UpdateBy time.Time `db:"update_by"` // 管理员的id
		UpdateAt time.Time `db:"update_at"` // 更新时间
	}
)

func newUserCallModel(conn sqlx.SqlConn) *defaultUserCallModel {
	return &defaultUserCallModel{
		conn:  conn,
		table: "`user_call`",
	}
}

func (m *defaultUserCallModel) Delete(ctx context.Context, id int64) error {
	query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
	_, err := m.conn.ExecCtx(ctx, query, id)
	return err
}

func (m *defaultUserCallModel) FindOne(ctx context.Context, id int64) (*UserCall, error) {
	query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", userCallRows, m.table)
	var resp UserCall
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

func (m *defaultUserCallModel) Insert(ctx context.Context, data *UserCall) (sql.Result, error) {
	query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?)", m.table, userCallRowsExpectAutoSet)
	ret, err := m.conn.ExecCtx(ctx, query, data.Name, data.CreateBy, data.UpdateBy)
	return ret, err
}

func (m *defaultUserCallModel) Update(ctx context.Context, data *UserCall) error {
	query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, userCallRowsWithPlaceHolder)
	_, err := m.conn.ExecCtx(ctx, query, data.Name, data.CreateBy, data.UpdateBy, data.Id)
	return err
}

func (m *defaultUserCallModel) tableName() string {
	return m.table
}
