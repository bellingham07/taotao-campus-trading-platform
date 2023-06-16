package trade

import (
	"context"

	"go-go-zero/service/trade/cmd/api/internal/svc"
	"go-go-zero/service/trade/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetByIdAndDoneLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetByIdAndDoneLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetByIdAndDoneLogic {
	return &GetByIdAndDoneLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetByIdAndDoneLogic) GetByIdAndDone(req *types.IdDoneReq) error {
	// todo: add your logic here and delete this line

	return nil
}
