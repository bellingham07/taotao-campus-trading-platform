package cinfo

import (
	"context"

	"go-go-zero/service/cmdty/cmd/api/internal/svc"
	"go-go-zero/service/cmdty/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type WantpublishLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewWantpublishLogic(ctx context.Context, svcCtx *svc.ServiceContext) *WantpublishLogic {
	return &WantpublishLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *WantpublishLogic) Wantpublish(req *types.InfoReq) error {
	// todo: add your logic here and delete this line

	return nil
}
