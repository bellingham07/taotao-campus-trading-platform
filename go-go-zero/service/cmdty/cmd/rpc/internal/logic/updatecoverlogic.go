package logic

import (
	"context"
	"go-go-zero/service/cmdty/model"

	"go-go-zero/service/cmdty/cmd/rpc/internal/svc"
	"go-go-zero/service/cmdty/cmd/rpc/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateCoverLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateCoverLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateCoverLogic {
	return &UpdateCoverLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateCoverLogic) UpdateCover(in *__.CoverReq) (*__.CodeResp, error) {
	var ci = &model.CmdtyInfo{Id: in.Id, Cover: in.Cover}
	if _, err := l.svcCtx.CmdtyInfo.Update(ci); err != nil {
		logx.Infof("[DB ERROR] RPC UpdateCover 更新商品封面失败 %v", err)
		return &__.CodeResp{Code: -1}, err
	}
	return &__.CodeResp{Code: 0}, nil
}
