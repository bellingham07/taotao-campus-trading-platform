package room

import (
	"context"
	"errors"
	"github.com/yitter/idgenerator-go/idgen"
	"go-go-zero/service/chat/model"
	"go-go-zero/service/chat/model/mongo"
	"net/http"
	"strconv"
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

const pongWait = 60 * time.Second

func (l *ChatLogic) Chat(req *types.ChatReq, w http.ResponseWriter, r *http.Request, userId int64) error {
	var (
		roomId    = req.RoomId
		roomIdStr = strconv.FormatInt(roomId, 10)
		sellerId  = req.SellerId
		buyerId   = req.BuyerId
		isBuyer   = false
	)
	if sellerId == 0 {
		isBuyer = true
	}

	// 然后开始进行websocket的连接
	conn, err := l.svcCtx.Upgrader.Upgrade(w, r, nil)
	if err != nil {
		conn.Close()
		return errors.New("无法连接到聊天室！😭")
	}

	if userId != sellerId && userId != buyerId {
		return errors.New("身份验证错误！")
	}

	// 先找有没有房间，没有房间直接返回
	cr := &model.ChatRoom{Id: roomId}
	has, err := l.svcCtx.ChatRoom.Get(cr)
	if !has || err != nil {
		logx.Debugf("[DB ERROR] Chat 查询聊天室信息失败 %v\n", err.Error())
		return errors.New("聊天室不存在！🫠")
	}

	conn.SetReadDeadline(time.Now().Add(pongWait)) // 连接进来先给一个60秒的超时时间
	conn.SetPongHandler(func(string) error {       // 每次收到 Pong 消息，更新连接的超时时间
		conn.SetReadDeadline(time.Now().Add(pongWait))
		return nil
	})
	defer conn.Close()

	icon := roomIdStr + strconv.FormatInt(userId, 10)
	l.svcCtx.Conn.Lock.Lock()
	l.svcCtx.Conn.ConnPool[icon] = conn
	l.svcCtx.Conn.Lock.Unlock()

	for {
		cm := new(mongo.ChatMessage)
		if err = conn.ReadJSON(cm); err != nil {
			return errors.New("消息未能发送成功！")
		}
		cm.Id = idgen.NextId()
		cm.RoomId = roomId
		cm.UserId = userId
		cm.Time = time.Now()

		// 检测超时时间是否已经过期
		if err = conn.SetReadDeadline(time.Now().Add(pongWait)); err != nil {
			go l.deleteConn(icon)
			conn.Close()
			return errors.New("连接已断开！")
		}

		if _, err = l.svcCtx.ChatMessage.InsertOne(l.ctx, cm); err != nil {
			go l.deleteConn(icon)
			logx.Debugf("[MONGO ERROR] Chat 插入聊天信息失败 %v\n", err)
			return errors.New("消息保存失败！")
		}
		msg, err := l.svcCtx.Json.Marshal(cm)
		if err != nil {
			go l.deleteConn(icon)
			logx.Debugf("[JSON MARSHAL ERROR] Chat 序列化消息错误 %v\n", err)
			return errors.New("未知错误！😭程序员大哭")
		}

		iconUser2 := roomIdStr
		if isBuyer {
			iconUser2 += strconv.FormatInt(cr.BuyerId, 10)
		} else {
			iconUser2 += strconv.FormatInt(cr.SellerId, 10)
		}
		l.svcCtx.Conn.Lock.RLock()
		conn, _ := l.svcCtx.Conn.ConnPool[iconUser2]
		l.svcCtx.Conn.Lock.RUnlock()

		// map取不到也不return，可能第一次创建对方还没有连过
		_ = conn.WriteJSON(msg)
	}
}

func (l *ChatLogic) deleteConn(icon string) {
	l.svcCtx.Conn.Lock.Lock()
	delete(l.svcCtx.Conn.ConnPool, icon)
	l.svcCtx.Conn.Lock.Unlock()
}
