package cinfo

import (
	"context"

	"go-go-zero/service/cmdty/cmd/api/internal/svc"
	"go-go-zero/service/cmdty/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type SellsaveLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSellsaveLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SellsaveLogic {
	return &SellsaveLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SellsaveLogic) Sellsave(req *types.InfoReq) error {
	infoCommonLogic := NewInfoCommonLogic(l.ctx, l.svcCtx)
	ci := infoCommonLogic.CopyPartial(req)
	// 第一次保存
	if req.Id == 0 {
		if err := infoCommonLogic.SaveOrPublishInfo(ci, 1, false); err != nil {
			return err
		}
		return nil
	}
	// 已经保存过，进行更新
	if err := infoCommonLogic.UpdateInfo(ci, false); err != nil {
		return err
	}
	return nil
}
