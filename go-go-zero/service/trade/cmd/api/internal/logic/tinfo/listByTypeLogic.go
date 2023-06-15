package tinfo

import (
	"context"

	"go-go-zero/service/trade/cmd/api/internal/svc"
	"go-go-zero/service/trade/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ListByTypeLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewListByTypeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ListByTypeLogic {
	return &ListByTypeLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ListByTypeLogic) ListByType(req *types.TypeReq) error {
	// todo: add your logic here and delete this line

	return nil
}
