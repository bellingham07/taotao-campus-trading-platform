package cmdty

import (
	"context"

	"github.com/zeromicro/go-zero/core/logx"
	"go-go-zero/service/file/cmd/api/internal/svc"
)

type RemoveLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRemoveLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RemoveLogic {
	return &RemoveLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RemoveLogic) Remove() error {
	// todo: add your logic here and delete this line

	return nil
}
