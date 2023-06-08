package cinfo

import (
	"context"

	"go-go-zero/service/cmdty/cmd/api/internal/svc"
	"go-go-zero/service/cmdty/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type WantsaveLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewWantsaveLogic(ctx context.Context, svcCtx *svc.ServiceContext) *WantsaveLogic {
	return &WantsaveLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *WantsaveLogic) Wantsave(req *types.InfoReq) (resp *types.BaseResp, err error) {
	// todo: add your logic here and delete this line

	return
}
