package cmt

import (
	"context"

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

func (l *ListByTradeIdLogic) ListByTradeId(req *types.TradeIdReq) error {
	// todo: add your logic here and delete this line

	return nil
}
