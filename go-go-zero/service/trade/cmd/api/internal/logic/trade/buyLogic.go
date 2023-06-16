package trade

import (
	"context"
	"errors"
	"go-go-zero/service/cmdty/cmd/rpc/cmdtyservice"
	"go-go-zero/service/trade/model"
	"time"

	"go-go-zero/service/trade/cmd/api/internal/svc"
	"go-go-zero/service/trade/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type BuyLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewBuyLogic(ctx context.Context, svcCtx *svc.ServiceContext) *BuyLogic {
	return &BuyLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *BuyLogic) Buy(req *types.BuyReq) (int64, error) {
	var (
		cmdtyId        = req.CmdtyId
		sellerId int64 = 408301323265285
		seller         = "小胖"
	)
	ti := &model.TradeInfo{
		BuyerId:           req.BuyerId,
		Buyer:             req.Buyer,
		SellerId:          sellerId,
		Seller:            seller,
		CmdtyId:           cmdtyId,
		Type:              1,
		Payment:           req.Payment,
		Status:            0,
		CreateAt:          time.Now().Local(),
		IsSellerConfirmed: 0,
		IsBuyerConfirmed:  1,
	}

	id, err := l.svcCtx.TradeInfo.Insert(ti)
	if err != nil {
		return 0, errors.New("开启交易失败！")
	}
	go l.updateCoverAndInfo(id, cmdtyId)
	return id, nil
}

func (l *BuyLogic) updateCoverAndInfo(id, cmdtyId int64) {
	resp, err := l.svcCtx.CmdtyRpc.GetCoverInfoById(l.ctx, &cmdtyservice.IdReq{Id: cmdtyId})
	if resp.GetCode() == -1 || err == nil {
		logx.Infof("[DB ERROR] updateCoverAndInfo 远程获取封面和简介失败 %v\n", err)
		return
	}
	ti := &model.TradeInfo{
		Id:         id,
		BriefIntro: resp.Info,
		Cover:      resp.Cover,
	}
	if _, err = l.svcCtx.TradeInfo.Update(ti); err != nil {
		logx.Infof("[DB ERROR] updateCoverAndInfo 更新交易的封面和简介失败 %v\n", err)
	}
}
