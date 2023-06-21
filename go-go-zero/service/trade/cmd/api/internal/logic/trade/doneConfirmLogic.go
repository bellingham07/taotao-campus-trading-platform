package trade

import (
	"context"

	"go-go-zero/service/trade/cmd/api/internal/svc"
	"go-go-zero/service/trade/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DoneConfirmLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDoneConfirmLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DoneConfirmLogic {
	return &DoneConfirmLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DoneConfirmLogic) DoneConfirm(req *types.IdReq) error {
	// todo: add your logic here and delete this line

	return nil
}
