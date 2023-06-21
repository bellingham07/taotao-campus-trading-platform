package trade

import (
	"context"
	"errors"
	"github.com/zeromicro/go-zero/core/logx"
	"go-go-zero/service/cmdty/cmd/rpc/cmdtyservice"
	"go-go-zero/service/trade/cmd/api/internal/svc"
	"go-go-zero/service/trade/model"
)

type CommonLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCommonLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CommonLogic {
	return &CommonLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CommonLogic) updateCoverAndInfo(id, cmdtyId int64) {
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

func (l *CommonLogic) save2DoneAndRemoveFromInfo(td *model.TradeDone, ti *model.TradeInfo) error {
	_, err := l.svcCtx.TradeDone.Insert(td)
	if err != nil {
		logx.Infof("[DB ERROR] Save2DoneRecord 插入交易完成记录失败 %v\n", err)
		return errors.New("操作失败！😢")
	}
	go func() {
		_, err = l.svcCtx.TradeInfo.Delete(ti)
		if err != nil {
			logx.Infof("[DB ERROR] Save2DoneRecord 删除交易记录失败 %v\n", err)
		}
	}()
	return nil
}
