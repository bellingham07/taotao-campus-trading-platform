package cinfo

import (
	"context"

	"go-go-zero/service/cmdty/cmd/api/internal/svc"
	"go-go-zero/service/cmdty/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ListByUidAndOptionLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewListByUidAndOptionLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ListByUidAndOptionLogic {
	return &ListByUidAndOptionLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ListByUidAndOptionLogic) ListByUidAndOption() (resp []types.InfoRespLite, err error) {
	// todo: add your logic here and delete this line

	return
}
