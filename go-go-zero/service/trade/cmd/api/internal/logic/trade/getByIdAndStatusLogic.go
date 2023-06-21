package trade

import (
	"context"

	"go-go-zero/service/trade/cmd/api/internal/svc"
	"go-go-zero/service/trade/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetByIdAndStatusLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetByIdAndStatusLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetByIdAndStatusLogic {
	return &GetByIdAndStatusLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetByIdAndStatusLogic) GetByIdAndStatus(req *types.IdStatusReq) error {
	// todo: add your logic here and delete this line

	return nil
}
