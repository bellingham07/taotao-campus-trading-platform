package cinfo

import (
	"context"
	"errors"
	"github.com/zeromicro/go-zero/core/logx"
	"go-go-zero/service/cmdty/cmd/api/internal/svc"
	"go-go-zero/service/cmdty/cmd/api/internal/types"
	"go-go-zero/service/cmdty/model"
	"time"
)

type InfoCommonLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewInfoCommonLogic(ctx context.Context, svcCtx *svc.ServiceContext) *InfoCommonLogic {
	return &InfoCommonLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// UpdateInfo 更新草稿或者已发布，flag为标识（true为更新已发布，false为更新草稿）
func (l *InfoCommonLogic) UpdateInfo(ci *model.CmdtyInfo, isPublish bool) error {
	// 不是发布，即只更新内容
	if !isPublish {
		if err := l.svcCtx.CmdtyInfo.Update(l.ctx, ci); err != nil {
			return errors.New("操作失败，请重试！😢")
		}
		return nil
	}
	// 是发布，就更新状态和内容
	ci.Status = 2
	ci.PublishAt = time.Now()
	if err := l.svcCtx.CmdtyInfo.Update(l.ctx, ci); err != nil {
		return errors.New("操作失败，请重试！😢")
	}
	return nil
}

// SaveOrPublishInfo 保存并发布商品信息，区分出售和购买
func (l *InfoCommonLogic) SaveOrPublishInfo(ci *model.CmdtyInfo, cmdtyType int64, isPublish bool) error {
	now := time.Now()
	ci.CreateAt = now
	ci.Type = cmdtyType
	// 保存草稿
	if !isPublish {
		if _, err := l.svcCtx.CmdtyInfo.Insert(l.ctx, ci); err != nil {
			return errors.New("操作失败，请重试！😢")
		}
		return nil
	}
	// 直接发布
	ci.Status = 2
	ci.PublishAt = now
	if _, err := l.svcCtx.CmdtyInfo.Insert(l.ctx, ci); err != nil {
		return errors.New("操作失败，请重试！😢")
	}
	return nil
}

func (l *InfoCommonLogic) CopyPartial(req *types.InfoReq) *model.CmdtyInfo {
	return &model.CmdtyInfo{
		Id:     req.Id,
		UserId: req.UserId,
		Cover:  req.Cover,
		Tag:    req.Cover,
		Price:  req.Price,
		Brand:  req.Cover,
		Model:  req.Cover,
		Intro:  req.Cover,
		Old:    req.Cover,
		Status: req.Status,
	}
}
