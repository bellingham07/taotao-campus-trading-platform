package logic

import (
	"context"
	"errors"
	"go-go-zero/service/user/model"

	"go-go-zero/service/user/cmd/rpc/internal/svc"
	"go-go-zero/service/user/cmd/rpc/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type RetrieveNameAndAvatarLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewRetrieveNameAndAvatarLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RetrieveNameAndAvatarLogic {
	return &RetrieveNameAndAvatarLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *RetrieveNameAndAvatarLogic) RetrieveNameAndAvatar(in *__.IdReq) (*__.NameAvatarResp, error) {
	var (
		ui   = &model.UserInfo{Id: in.Id}
		resp = &__.NameAvatarResp{Code: -1}
	)
	has, err := l.svcCtx.UserInfo.Cols("name", "avatar").Get(ui)
	if !has || err != nil {
		return resp, errors.New("查询昵称和头像失败！")
	}
	resp.Code = 0
	resp.Name = ui.Name
	resp.Avatar = ui.Avatar
	return resp, nil
}
