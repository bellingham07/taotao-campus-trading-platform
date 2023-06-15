package cmt

import (
	"context"

	"go-go-zero/service/trade/cmd/api/internal/svc"
	"go-go-zero/service/trade/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CmtLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCmtLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CmtLogic {
	return &CmtLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CmtLogic) Cmt(req *types.CmtReq) error {
	// todo: add your logic here and delete this line

	return nil
}
