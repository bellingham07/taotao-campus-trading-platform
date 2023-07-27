package logic

import (
	"context"
	"go-go-zero/service/file/model"

	"go-go-zero/service/file/cmd/rpc/internal/svc"
	"go-go-zero/service/file/cmd/rpc/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetCmdtyPicsByOrderLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetCmdtyPicsByOrderLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetCmdtyPicsByOrderLogic {
	return &GetCmdtyPicsByOrderLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetCmdtyPicsByOrderLogic) GetCmdtyPicsByOrder(in *__.CmdtyPicsReq) (*__.CmdtyPicsResp, error) {
	var resp = &__.CmdtyPicsResp{Code: -1, Pics: nil}

	cfs := make([]model.FileCmdty, 0)
	err := l.svcCtx.FileCmdty.Where("cmdty_id = ?", in.Id).Asc("order").Find(&cfs)
	if err != nil {
		return resp, err
	}

	pics := make([]*__.Pic, 0)
	for _, cf := range cfs {
		pic := &__.Pic{
			Id:    cf.Id,
			Order: cf.Order,
			Url:   cf.Url,
		}
		pics = append(pics, pic)
	}
	resp.Code = 0
	resp.Pics = pics
	return resp, nil
}
