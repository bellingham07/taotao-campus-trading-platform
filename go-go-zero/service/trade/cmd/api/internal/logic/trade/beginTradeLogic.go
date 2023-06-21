package trade

import (
	"context"
	"errors"
	"github.com/yitter/idgenerator-go/idgen"
	"go-go-zero/service/trade/cmd/api/internal/svc"
	"go-go-zero/service/trade/cmd/api/internal/types"
	"go-go-zero/service/trade/model"
	"time"

	"github.com/zeromicro/go-zero/core/logx"
)

type BeginTradeLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewBeginTradeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *BeginTradeLogic {
	return &BeginTradeLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *BeginTradeLogic) BeginTrade(req *types.TradeReq) (int64, error) {
	var (
		cmdtyId       = req.CmdtyId
		userId  int64 = 408301323265285
		user          = "小胖"
	)
	ti := &model.TradeInfo{
		Id:       idgen.NextId(),
		CmdtyId:  cmdtyId,
		Payment:  req.Payment,
		Type:     req.Type,
		Status:   0,
		CreateAt: time.Now().Local(),
	}
	if req.Type == 1 {
		if req.OwnerId != userId {
			return 0, errors.New("交易信息错误！")
		}
		ti.BuyerId = req.OtherId
		ti.Buyer = req.Other
		ti.SellerId = userId
		ti.Seller = user
		ti.IsSellerConfirmed = 0
		ti.IsBuyerConfirmed = 1
	} else {
		if req.OtherId != userId {
			return 0, errors.New("交易信息错误！")
		}
		ti.BuyerId = userId
		ti.Buyer = user
		ti.SellerId = req.OtherId
		ti.Seller = req.Other
		ti.IsSellerConfirmed = 1
		ti.IsBuyerConfirmed = 0
	}

	id, err := l.svcCtx.TradeInfo.Insert(ti)
	if err != nil {
		return 0, errors.New("开启交易失败！")
	}
	tcLogic := NewCommonLogic(l.ctx, l.svcCtx)
	go tcLogic.updateCoverAndInfo(id, cmdtyId)
	return id, nil
}
