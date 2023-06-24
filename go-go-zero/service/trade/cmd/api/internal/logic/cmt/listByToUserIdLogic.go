package cmt

import (
	"context"
	"errors"
	"go-go-zero/service/trade/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"

	"go-go-zero/service/trade/cmd/api/internal/svc"
	"go-go-zero/service/trade/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ListByToUserIdLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewListByToUserIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ListByToUserIdLogic {
	return &ListByToUserIdLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ListByToUserIdLogic) ListByToUserId(req *types.UserIdReq) ([]model.TradeCmt, error) {
	var (
		tcs         = make([]model.TradeCmt, 0)
		filter      = bson.M{"to_user_id": req.UserId}
		findOptions = &options.FindOptions{Sort: bson.M{"create_at": -1}}
	)
	cursor, err := l.svcCtx.TradeCmt.Find(l.ctx, filter, findOptions)
	if err != nil {
		logx.Infof("[MONGO ERROR] ListByToUserId 获取被评列表失败 %v\n", err)
		return nil, errors.New("出错啦！😭")
	}
	for cursor.Next(l.ctx) {
		tc := model.TradeCmt{}
		if err = cursor.Decode(&tc); err != nil {
			logx.Infof("[MONGO ERROR] ListByToUserId 解析结果错误 %v\n", err)
			continue
		}
		tcs = append(tcs, tc)
	}
	return tcs, nil
}
