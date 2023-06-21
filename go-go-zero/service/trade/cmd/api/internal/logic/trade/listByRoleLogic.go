package trade

import (
	"context"
	"go-go-zero/service/trade/model"

	"go-go-zero/service/trade/cmd/api/internal/svc"
	"go-go-zero/service/trade/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ListByRoleLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewListByRoleLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ListByRoleLogic {
	return &ListByRoleLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ListByRoleLogic) ListByRole(req *types.RoleReq) (interface{}, error) {
	var (
		userId int64 = 408301323265285
		tis          = make([]*model.TradeInfo, 0)
		tds          = make([]*model.TradeDone, 0)
	)
	if req.Role == 1 {
		// 1 先查询还在进行中的交易
		err := l.svcCtx.TradeInfo.Where("`buyer_id` = ?", userId).
			Desc("create_at").Find(tis)
		if err != nil {
			logx.Infof("[DB ERROR] ListByType 查询买家为%v的未完成记录错误 %v\n", userId, err)
			return nil, err
		}
		// 2 再查已完成的交易
		err = l.svcCtx.TradeDone.Where("`buyer_id` = ?", userId).
			Desc("done_at").Find(tds)
		if err != nil {
			logx.Infof("[DB ERROR] ListByType 查询买家为%v的完成记录错误 %v\n", userId, err)
			return tis, err
		}
		return l.appendTisAndTds(tis, tds)
	} else {
		// 1 先查询还在进行中的交易
		err := l.svcCtx.TradeInfo.Where("`seller_id` = ?", userId).
			Desc("create_at").Find(tis)
		if err != nil {
			logx.Infof("[DB ERROR] ListByType 查询买家为%v的未完成记录错误 %v\n", userId, err)
			return nil, err
		}
		// 2 再查已完成的交易
		err = l.svcCtx.TradeDone.Where("`seller_id` = ?", userId).
			Desc("done_at").Find(tds)
		if err != nil {
			logx.Infof("[DB ERROR] ListByType 查询买家为%v的完成记录错误 %v\n", userId, err)
			return tis, err
		}
		return l.appendTisAndTds(tis, tds)
	}
}

// 3 拼成一个slice
func (l *ListByRoleLogic) appendTisAndTds(tis []*model.TradeInfo, tds []*model.TradeDone) (interface{}, error) {
	data := make([]interface{}, len(tis)+len(tds))
	for idx, ti := range tis {
		data[idx] = ti
	}
	for idx, td := range tds {
		data[len(tis)+idx] = td
	}
	return data, nil
}
