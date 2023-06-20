package logic

import (
	"context"
	"go-go-zero/service/user/model"

	"go-go-zero/service/user/cmd/rpc/internal/svc"
	"go-go-zero/service/user/cmd/rpc/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type IncrLikeLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewIncrLikeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *IncrLikeLogic {
	return &IncrLikeLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *IncrLikeLogic) IncrLike(in *__.IdReq) (*__.CodeResp, error) {
	var resp = new(__.CodeResp)
	_, err := l.svcCtx.UserInfo.Incr("like").ID(in.GetId()).Update(&model.UserInfo{})
	if err != nil {
		logx.Infof("[DB ERROR] IncrLike 更新用户点赞数失败 %v\n", err)
		resp.Code = -1
		return resp, err
	}
	resp.Code = 0
	return resp, nil
}
