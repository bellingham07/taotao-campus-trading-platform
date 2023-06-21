package trade

import (
	"context"
	"errors"
	"go-go-zero/service/trade/model"
	"time"
	"xorm.io/xorm"

	"go-go-zero/service/trade/cmd/api/internal/svc"
	"go-go-zero/service/trade/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CancelLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCancelLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CancelLogic {
	return &CancelLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CancelLogic) Cancel(req *types.IdReq) error {
	var ti = &model.TradeInfo{Id: req.Id}
	_, err := l.svcCtx.Xorm.Transaction(func(session *xorm.Session) (interface{}, error) {
		tx := session.Table("trade_info")
		has, err := tx.Get(ti)
		if !has || err != nil {
			return nil, errors.New("找不到这个交易！")
		}
		if ti.Status != 1 {
			return nil, errors.New("交易信息错误！")
		}
		td := &model.TradeDone{
			Id:         ti.Id,
			BuyerId:    ti.BuyerId,
			Buyer:      ti.Buyer,
			SellerId:   ti.SellerId,
			Seller:     ti.Seller,
			CmdtyId:    ti.CmdtyId,
			Type:       ti.Type,
			Location:   ti.Location,
			BriefIntro: ti.BriefIntro,
			Cover:      ti.Cover,
			Payment:    ti.Payment,
			Status:     3, // 3表示取消
			CreateAt:   ti.CreateAt,
			DoneAt:     time.Now().Local(),

			IsCmtDone: 0,
		}
		tcLogic := NewCommonLogic(l.ctx, l.svcCtx)
		err = tcLogic.save2DoneAndRemoveFromInfo(td, ti)
		if err != nil {
			return nil, errors.New("操作失败！")
		}
		return nil, nil
	})
	return err
}
