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

type ConfirmLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewConfirmLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ConfirmLogic {
	return &ConfirmLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ConfirmLogic) Confirm(req *types.IdStatusReq) (interface{}, error) {
	var (
		id           = req.Id
		userId int64 = 408301323265285
		ti           = &model.TradeInfo{Id: id}
	)
	// 判断当前交易是在那个状态
	if req.Status == 0 {
		// 1 交易状态为0，说明有一方为确认进行交易
		return l.status0confirm(userId, ti)
	} else if req.Status == 1 {
		// 2 交易状态为1，说明已经开始进行交易，可以进行完成确认了
		return l.status1confirm(userId, ti)
	}
	return nil, errors.New("参数错误！")
}

// 交易状态为0，进行开始确认的逻辑
func (l *ConfirmLogic) status0confirm(userId int64, ti *model.TradeInfo) (interface{}, error) {
	status, err := l.svcCtx.Xorm.Transaction(func(session *xorm.Session) (interface{}, error) {
		tx := session.Table("trade_info")
		has, err := l.svcCtx.TradeInfo.Get(ti)
		if !has || err != nil {
			return 0, errors.New("没有这个交易！😢")
		}
		if userId == ti.SellerId {
			ti.IsSellerConfirmed = 1
			if ti.IsBuyerConfirmed == 1 {
				ti.Status = 1
			} else {
				return nil, errors.New("交易信息错误！")
			}
			_, err = tx.Cols("is_seller_confirmed", "status").ID(ti.Id).Update(ti)
			if err != nil {
				logx.Infof("[DB ERROR] Confirm 更新交易信息错误 %v\n", err)
				return 0, errors.New("操作失败！😢")
			}
			return ti.Status, nil
		} else if userId == ti.BuyerId {
			ti.IsBuyerConfirmed = 1
			if ti.IsSellerConfirmed == 1 {
				ti.Status = 1
			} else {
				return nil, errors.New("交易信息错误！")
			}
			_, err = tx.Cols("is_buyer_confirmed", "status").ID(ti.Id).Update(ti)
			if err != nil {
				logx.Infof("[DB ERROR] Confirm 更新交易信息错误 %v\n", err)
				return 0, errors.New("操作失败！😢")
			}
			return ti.Status, nil
		}
		return 0, errors.New("交易信息错误！")
	})
	return status, err
}

// 交易状态为1，进行完成确认的逻辑
func (l *ConfirmLogic) status1confirm(userId int64, ti *model.TradeInfo) (interface{}, error) {
	status, err := l.svcCtx.Xorm.Transaction(func(session *xorm.Session) (interface{}, error) {
		tx := session.Table("trade_info")
		has, err := l.svcCtx.TradeInfo.Get(ti)
		if !has || err != nil {
			return 0, errors.New("没有这个交易！😢")
		}
		now := time.Now().Local()
		if ti.Status != 1 {
			return 0, errors.New("交易信息错误！")
		}
		if userId == ti.BuyerId {
			ti.IsBuyerDone = 1
			ti.BuyerDoneAt = now
			if ti.IsSellerDone == 1 {
				err = l.saveRecord2Done(ti, now)
				if err != nil {
					return 1, err
				}
				return 2, nil
			}
			_, err = tx.Cols("is_buyer_done", "buyer_done_at", "status").ID(ti.Id).Update(ti)
			if err != nil {
				return 1, errors.New("操作失败！😢")
			}
			return 2, nil
		} else if userId == ti.SellerId {
			ti.IsSellerDone = 1
			ti.SellerDoneAt = now
			if ti.IsBuyerDone == 1 {
				err = l.saveRecord2Done(ti, now)
				if err != nil {
					return 1, err
				}
				return 2, nil
			}
			_, err = tx.Cols("is_buyer_done", "buyer_done_at", "status").ID(ti.Id).Update(ti)
			if err != nil {
				return 1, errors.New("操作失败！😢")
			}
			return 2, nil
		}
		return 1, errors.New("交易信息错误！")
	})
	return status, err
}

// 交易完成，保存至done表
func (l *ConfirmLogic) saveRecord2Done(ti *model.TradeInfo, now time.Time) error {
	var td = &model.TradeDone{
		Id:           ti.Id,
		BuyerId:      ti.BuyerId,
		Buyer:        ti.Buyer,
		SellerId:     ti.SellerId,
		Seller:       ti.Seller,
		CmdtyId:      ti.CmdtyId,
		Type:         ti.Type,
		BriefIntro:   ti.BriefIntro,
		Cover:        ti.Cover,
		Payment:      ti.Payment,
		Location:     ti.Location,
		Status:       2,
		CreateAt:     ti.CreateAt,
		SellerDoneAt: ti.SellerDoneAt,
		BuyerDoneAt:  ti.BuyerDoneAt,
		DoneAt:       now,
		IsSellerCmt:  0,
		IsBuyerCmt:   0,
	}
	tcLogic := NewCommonLogic(l.ctx, l.svcCtx)
	return tcLogic.save2DoneAndRemoveFromInfo(td, ti)
}
