package cmt

import (
	"context"
	"errors"
	"go-go-zero/service/cmdty/model"

	"go-go-zero/service/cmdty/cmd/api/internal/svc"
	"go-go-zero/service/cmdty/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ListByInfoIdLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewListByInfoIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ListByInfoIdLogic {
	return &ListByInfoIdLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ListByInfoIdLogic) ListByInfoId(req *types.CmdtyIdReq) (resp []*model.CmdtyCmt, err error) {
	resp = l.svcCtx.CmdtyCmt.ListByCmdtyId(req.Id)
	if resp == nil {
		return nil, errors.New("出错啦！😢")
	}
	return resp, nil
}
