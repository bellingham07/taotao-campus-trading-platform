package room

import (
	"com.xpdj/go-gin/model/response"
	chatRepository "com.xpdj/go-gin/repository/chat"
	"context"
	"errors"
	"github.com/gorilla/websocket"
	"go-go-zero/service/chat/model"
	"go-go-zero/service/chat/model/mongo"
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
	var userId int64 = 408301323265285
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
		cm := mongo.ChatMessage{
			RoomId: 0,
			Time:   time.Now(),
			UserId: userId,
		}
		if err = conn.ReadJSON(cm); err != nil {
			logx.Debugf("[WEBSOCKET ERROR] Chat 解析websocket消息错误 " + err.Error())
			return errors.New("消息发送失败！")
		}
		if one, err := l.svcCtx.ChatMessage.InsertOne(l.ctx, cm); err != nil {
			return err
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
