package logic

import (
	"context"
	"go-go-zero/service/cmdty/model"

	"go-go-zero/service/cmdty/cmd/rpc/internal/svc"
	"go-go-zero/service/cmdty/cmd/rpc/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetCoverInfoByIdLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetCoverInfoByIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetCoverInfoByIdLogic {
	return &GetCoverInfoByIdLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetCoverInfoByIdLogic) GetCoverInfoById(in *__.IdReq) (*__.CoverInfoResp, error) {
	var (
		ci   = model.CmdtyInfo{Id: in.GetId()}
		resp = &__.CoverInfoResp{Code: -1}
	)
	has, err := l.svcCtx.CmdtyInfo.Get(ci)
	if has && err == nil {
		resp.Code = 0
		resp.Cover = ci.Cover
		resp.Info = ci.Intro[:20]
		return resp, nil
	}
	logx.Infof("[DB ERROR] GetCoverInfoById 查询封面和简介失败 %v\n", err)
	return resp, err
}
