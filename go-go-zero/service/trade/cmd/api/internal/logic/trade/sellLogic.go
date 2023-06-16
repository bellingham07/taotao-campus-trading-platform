package trade

import (
	"context"

	"go-go-zero/service/trade/cmd/api/internal/svc"
	"go-go-zero/service/trade/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type SellLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSellLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SellLogic {
	return &SellLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SellLogic) Sell(req *types.BuyReq) error {
	// todo: add your logic here and delete this line

	return nil
}
