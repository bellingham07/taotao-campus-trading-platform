package model

import "github.com/zeromicro/go-zero/core/stores/sqlx"

var _ ChatRoomModel = (*customChatRoomModel)(nil)

type (
	// ChatRoomModel is an interface to be customized, add more methods here,
	// and implement the added methods in customChatRoomModel.
	ChatRoomModel interface {
		chatRoomModel
	}

	customChatRoomModel struct {
		*defaultChatRoomModel
	}
)

// NewChatRoomModel returns a model for the database table.
func NewChatRoomModel(conn sqlx.SqlConn) ChatRoomModel {
	return &customChatRoomModel{
		defaultChatRoomModel: newChatRoomModel(conn),
	}
}
