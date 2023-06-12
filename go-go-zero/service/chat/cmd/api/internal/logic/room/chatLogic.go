package room

import (
	"com.xpdj/go-gin/model/response"
	chatRepository "com.xpdj/go-gin/repository/chat"
	"context"
	"errors"
	"github.com/gorilla/websocket"
	"go-go-zero/service/chat/model"
	"log"
	"net/http"
	"time"

	"go-go-zero/service/chat/cmd/api/internal/svc"
	"go-go-zero/service/chat/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ChatLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewChatLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ChatLogic {
	return &ChatLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ChatLogic) Chat(req *types.IdReq, w http.ResponseWriter, r *http.Request) error {
	roomId := req.Id
	conn, err := l.svcCtx.Upgrader.Upgrade(w, r, nil)
	if err != nil {
		return errors.New("无法连接到聊天室！😭")
	}
	cr := &model.ChatRoom{Id: roomId}
	has, err := l.svcCtx.Xorm.Table("chat_room").Get(cr)
	if !has || err != nil {
		return errors.New("出错啦！😭")
	}
	for {
		msg := &model.ChatMessage{
			RoomId: id,
			Time:   time.Now(),
			UserId: userId,
		}
		err = myConn.readJSON(msg)
		if err != nil {
			log.Println("地欧弟3", err)
			c.JSON(http.StatusBadRequest, response.ErrorMsg("消息发送失败"))
			return
		}
		if err = chatRepository.ChatMessage.Insert(msg); err != nil {
			log.Println("地欧弟4", err)
			c.JSON(http.StatusBadRequest, response.ErrorMsg("消息发送失败"))
			return
		}
		if cc, ok := wc[msg.ToUserId]; ok {
			err = cc.WriteMessage(websocket.TextMessage, []byte(msg.Content))
			if err != nil {
				log.Printf("Write Message Error:%v\n\n", err)
				return
			}
		}
	}
	return nil
}
