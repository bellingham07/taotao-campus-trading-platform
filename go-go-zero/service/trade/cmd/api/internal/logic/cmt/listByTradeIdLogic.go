package cmt

import (
	"context"
	"errors"
	"go-go-zero/service/trade/model"
	"go.mongodb.org/mongo-driver/bson"

	"go-go-zero/service/trade/cmd/api/internal/svc"
	"go-go-zero/service/trade/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ListByTradeIdLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewListByTradeIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ListByTradeIdLogic {
	return &ListByTradeIdLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ListByTradeIdLogic) ListByTradeId(req *types.TradeIdReq) ([]*model.TradeCmt, error) {
	var (
		tcs    = make([]*model.TradeCmt, 0)
		filter = bson.M{"trade_id": req.TradeId}
	)
	cursor, err := l.svcCtx.TradeCmt.Find(l.ctx, filter)
	if err != nil {
		logx.Infof("[MONGO ERROR] ListByTradeId 获取被评列表失败 %v\n", err)
		return nil, errors.New("出错啦！😭")
	}
	for cursor.Next(l.ctx) {
		tc := new(model.TradeCmt)
		if err = cursor.Decode(tc); err != nil {
			logx.Infof("[MONGO ERROR] ListByTradeId 解析结果错误 %v\n", err)
			continue
		}
		tcs = append(tcs, tc)
	}
	return tcs, nil
}
