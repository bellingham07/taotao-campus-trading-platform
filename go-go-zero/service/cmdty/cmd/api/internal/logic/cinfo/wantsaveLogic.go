package cinfo

import (
	"context"
	"github.com/zeromicro/go-zero/core/logx"
	"go-go-zero/service/cmdty/cmd/api/internal/svc"
	"go-go-zero/service/cmdty/cmd/api/internal/types"
)

type WantsaveLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewWantsaveLogic(ctx context.Context, svcCtx *svc.ServiceContext) *WantsaveLogic {
	return &WantsaveLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *WantsaveLogic) Wantsave(req *types.InfoReq) error {
	infoCommonLogic := NewInfoCommonLogic(l.ctx, l.svcCtx)
	ci := infoCommonLogic.CopyPartial(req)
	// 第一次保存
	if req.Id == 0 {
		if err := infoCommonLogic.SaveOrPublishInfo(ci, 2, false); err != nil {
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
