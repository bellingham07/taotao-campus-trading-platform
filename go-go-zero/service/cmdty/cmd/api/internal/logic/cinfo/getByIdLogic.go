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
	ci, _ := l.svcCtx.CmdtyInfo.FindOne(l.ctx, req.Id)
	resp = &types.InfoResp{
		Id:        0,
		UserId:    0,
		Cover:     "",
		Tag:       "",
		Price:     0,
		Brand:     "",
		Model:     "",
		Intro:     "",
		Old:       "",
		Status:    0,
		CreateAt:  "",
		PublishAt: "",
		View:      0,
		Collect:   0,
		Type:      0,
		Like:      0,
	}
	return
}
