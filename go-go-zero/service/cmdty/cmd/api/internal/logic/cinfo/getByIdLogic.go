package cinfo

import (
	"context"
	"go-go-zero/service/cmdty/cmd/api/internal/svc"
	"go-go-zero/service/cmdty/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetByIdLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetByIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetByIdLogic {
	return &GetByIdLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetByIdLogic) GetById(req *types.IdReq) (resp *types.InfoResp, err error) {
	ci, err := l.svcCtx.CmdtyInfo.FindOne(l.ctx, req.Id)
	if err != nil {
		resp = &types.InfoResp{
			BaseResp: types.BaseResp{Code: 0, Msg: "出错啦😼"},
		}
		return resp, nil
	}
	resp = &types.InfoResp{
		BaseResp: types.BaseResp{Code: 1, Msg: "ok"},
		Info: types.Info{
			Id:        ci.Id,
			UserId:    ci.UserId,
			Cover:     ci.Cover,
			Tag:       ci.Tag,
			Price:     ci.Price,
			Brand:     ci.Brand,
			Model:     ci.Model,
			Intro:     ci.Intro,
			Old:       ci.Old,
			Status:    ci.Status,
			CreateAt:  ci.CreateAt.Format("2006-01-02 15:04:05"),
			PublishAt: ci.PublishAt.Format("2006-01-02 15:04:05"),
			View:      ci.View,
			Collect:   ci.Collect,
			Type:      ci.Type,
			Like:      ci.Like,
		},
	}
	return
}
