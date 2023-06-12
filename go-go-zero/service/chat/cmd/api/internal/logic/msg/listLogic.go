package msg

import (
	"context"
	"errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"

	"go-go-zero/service/chat/cmd/api/internal/svc"
	"go-go-zero/service/chat/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ListLogic {
	return &ListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ListLogic) List(req *types.ListMessageReq) ([]*types.ChatMessageResp, error) {
	filter := bson.M{"room_id": req.Id}
	page := req.Page - 1
	var pageSize int64 = 20
	skip := (page - 1) * pageSize
	sort := bson.M{"time": -1} // 按消息发出倒序排序
	// 执行查询
	cursor, err := l.svcCtx.ChatMessage.Find(l.ctx, filter,
		options.Find().SetSkip(skip).SetLimit(pageSize).SetSort(sort))
	if err != nil {
		return nil, errors.New("聊天记录加载错误！😥")
	}
	defer cursor.Close(l.ctx)

	resp := make([]*types.ChatMessageResp, 0)
	for cursor.Next(l.ctx) {
		cm := new(types.ChatMessageResp)
		if err = cursor.Decode(cm); err != nil {
			return resp, errors.New("出错啦！😥")
		}
		resp = append(resp, cm)
	}
	return resp, nil
}
