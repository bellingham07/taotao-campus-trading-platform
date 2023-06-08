package cinfo

import (
	"context"

	"go-go-zero/service/cmdty/cmd/api/internal/svc"
	"go-go-zero/service/cmdty/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type SellpublishLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSellpublishLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SellpublishLogic {
	return &SellpublishLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SellpublishLogic) Sellpublish(req *types.InfoReq) (resp *types.BaseResp, err error) {
	// todo: add your logic here and delete this line

	return
}
