package cmt

import (
	"context"

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

func (l *ListByToUserIdLogic) ListByToUserId(req *types.UserIdReq) error {
	// todo: add your logic here and delete this line

	return nil
}
