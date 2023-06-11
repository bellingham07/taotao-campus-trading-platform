package cinfo

import (
	"context"

	"go-go-zero/service/cmdty/cmd/api/internal/svc"
	"go-go-zero/service/cmdty/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type WantpublishLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewWantpublishLogic(ctx context.Context, svcCtx *svc.ServiceContext) *WantpublishLogic {
	return &WantpublishLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *WantpublishLogic) Wantpublish(req *types.InfoReq) error {
	infoCommonLogic := NewInfoCommonLogic(l.ctx, l.svcCtx)
	ci := infoCommonLogic.CopyPartial(req)
	// 未保存直接发布
	if req.Id == 0 {
		if err := infoCommonLogic.SaveOrPublishInfo(ci, 2, true); err != nil {
			return err
		}
		return nil
	}
	// 已经保存过，进行更新
	if err := infoCommonLogic.UpdateInfo(ci, true); err != nil {
		return err
	}
	return nil
}
