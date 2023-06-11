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
	fileCmdtyFieldNames          = builder.RawFieldNames(&FileCmdty{})
	fileCmdtyRows                = strings.Join(fileCmdtyFieldNames, ",")
	fileCmdtyRowsExpectAutoSet   = strings.Join(stringx.Remove(fileCmdtyFieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), ",")
	fileCmdtyRowsWithPlaceHolder = strings.Join(stringx.Remove(fileCmdtyFieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), "=?,") + "=?"
)

type (
	fileCmdtyModel interface {
		Insert(ctx context.Context, data *FileCmdty) (sql.Result, error)
		FindOne(ctx context.Context, id int64) (*FileCmdty, error)
		Update(ctx context.Context, data *FileCmdty) error
		Delete(ctx context.Context, id int64) error
	}

	defaultFileCmdtyModel struct {
		conn  sqlx.SqlConn
		table string
	}

	FileCmdty struct {
		Id         int64     `db:"id"` // bigint自增
		CmdtyId    int64     `db:"cmdty_id"`
		UserId     int64     `db:"user_id"`
		Url        string    `db:"url"`
		ObjectName string    `db:"objectName"`
		UploadAt   time.Time `db:"upload_at"`
		IsCover    int64     `db:"is_cover"` // 默认为0，封面为1
		Order      int64     `db:"order"`
	}
)

func newFileCmdtyModel(conn sqlx.SqlConn) *defaultFileCmdtyModel {
	return &defaultFileCmdtyModel{
		conn:  conn,
		table: "`file_cmdty`",
	}
}

func (m *defaultFileCmdtyModel) Delete(ctx context.Context, id int64) error {
	query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
	_, err := m.conn.ExecCtx(ctx, query, id)
	return err
}

func (m *defaultFileCmdtyModel) FindOne(ctx context.Context, id int64) (*FileCmdty, error) {
	query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", fileCmdtyRows, m.table)
	var resp FileCmdty
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

func (m *defaultFileCmdtyModel) Insert(ctx context.Context, data *FileCmdty) (sql.Result, error) {
	query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?, ?, ?, ?)", m.table, fileCmdtyRowsExpectAutoSet)
	ret, err := m.conn.ExecCtx(ctx, query, data.CmdtyId, data.UserId, data.Url, data.ObjectName, data.UploadAt, data.IsCover, data.Order)
	return ret, err
}

func (m *defaultFileCmdtyModel) Update(ctx context.Context, data *FileCmdty) error {
	query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, fileCmdtyRowsWithPlaceHolder)
	_, err := m.conn.ExecCtx(ctx, query, data.CmdtyId, data.UserId, data.Url, data.ObjectName, data.UploadAt, data.IsCover, data.Order, data.Id)
	return err
}

func (m *defaultFileCmdtyModel) tableName() string {
	return m.table
}
