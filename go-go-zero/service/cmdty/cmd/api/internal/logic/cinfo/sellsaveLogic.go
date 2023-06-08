package cinfo

import (
	"context"

	"go-go-zero/service/cmdty/cmd/api/internal/svc"
	"go-go-zero/service/cmdty/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type SellsaveLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSellsaveLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SellsaveLogic {
	return &SellsaveLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SellsaveLogic) Sellsave(req *types.InfoReq) (resp *types.BaseResp, err error) {
	// todo: add your logic here and delete this line

	return
}
