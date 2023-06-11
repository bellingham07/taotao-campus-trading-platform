package model

import (
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"time"
)

var _ FileAvatarModel = (*customFileAvatarModel)(nil)

type (
	// FileAvatarModel is an interface to be customized, add more methods here,
	// and implement the added methods in customFileAvatarModel.
	FileAvatarModel interface {
		fileAvatarModel
		SaveOrUpdateByUserId(url string, name string, id int64) error
		DeleteByUserId(id int64) error
	}

	customFileAvatarModel struct {
		*defaultFileAvatarModel
	}
)

// NewFileAvatarModel returns a model for the database table.
func NewFileAvatarModel(conn sqlx.SqlConn) FileAvatarModel {
	return &customFileAvatarModel{
		defaultFileAvatarModel: newFileAvatarModel(conn),
	}
}

func (c customFileAvatarModel) SaveOrUpdateByUserId(url string, objectname string, userId int64) error {
	err := c.conn.Transact(func(session sqlx.Session) error {
		query := "select id from file_avatar where user_id = ?"
		fa := new(FileAvatar)
		err := session.QueryRowPartial(fa, query, userId)
		// 查询不到，就执行插入
		if fa.Id == 0 || err != nil {
			stmtInsert := "insert into `file_avatar` (`url`, `objectname`, `user_id`, `upload_at`) values (?, ?, ?, ?)"
			_, err = session.Exec(stmtInsert, url, objectname, userId, time.Now().Format("2006-01-02 15:04:05"))
			if err != nil {
				logx.Debugf("[DB ERROR] SaveOrUpdateByUserId 添加头像数据错误 " + err.Error())
				return err
			}
			return nil
		}
		// 查询到了，就执行更新
		stmtUpdate := "update `file_avatar` set `url` = ?, `objectname` = ?, `upload_at` = ? where `user_id` = ?"
		_, err = session.Exec(stmtUpdate, url, objectname, time.Now().Format("2006-01-02 15:04:05"), userId)
		if err != nil {
			logx.Debugf("[DB ERROR] SaveOrUpdateByUserId 修改头像数据错误 " + err.Error())
			return err
		}
		return nil
	})
	return err
}

func (c customFileAvatarModel) DeleteByUserId(userId int64) error {
	stmt := "delete from `file_avatar` where user_id = ?"
	if _, err := c.conn.Exec(stmt, userId); err != nil {
		logx.Debugf("[DB ERROR] DeleteByUserId 删除头像记录失败 " + err.Error())
		return err
	}
	return nil
}
