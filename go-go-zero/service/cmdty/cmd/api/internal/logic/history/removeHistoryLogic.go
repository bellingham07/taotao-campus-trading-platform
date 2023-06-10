package history

import (
	"context"

	"go-go-zero/service/cmdty/cmd/api/internal/svc"
	"go-go-zero/service/cmdty/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type RemoveHistoryLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRemoveHistoryLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RemoveHistoryLogic {
	return &RemoveHistoryLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RemoveHistoryLogic) RemoveHistory(req *types.IdsReq) (resp *types.BaseResp, err error) {
	// todo: add your logic here and delete this line

	return
}
