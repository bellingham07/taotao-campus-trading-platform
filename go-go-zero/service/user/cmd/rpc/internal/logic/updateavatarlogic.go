package logic

import (
	"context"
	"go-go-zero/service/user/model"

	"go-go-zero/service/user/cmd/rpc/internal/svc"
	"go-go-zero/service/user/cmd/rpc/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateAvatarLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateAvatarLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateAvatarLogic {
	return &UpdateAvatarLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateAvatarLogic) UpdateAvatar(in *__.AvatarReq) (*__.CodeResp, error) {
	var ui = &model.UserInfo{
		Id:     in.Id,
		Avatar: in.Avatar,
	}
	if _, err := l.svcCtx.UserInfo.Update(ui); err != nil {
		logx.Infof("[DB ERROR] UpdateAvatar 更新头像失败 %v\n", err)
		return &__.CodeResp{Code: -1}, nil
	}
	return &__.CodeResp{Code: 0}, nil
}
