package uinfo

import (
	"context"
	"errors"
	"go-go-zero/service/user/model"

	"go-go-zero/service/user/cmd/api/internal/svc"
	"go-go-zero/service/user/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateByIdLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateByIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateByIdLogic {
	return &UpdateByIdLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateByIdLogic) UpdateById(req *types.UserInfoReq) error {
	var ui = &model.UserInfo{
		Id:       req.Id,
		Name:     req.Name,
		Gender:   req.Gender,
		Phone:    req.Phone,
		Intro:    req.Intro,
		Location: req.Location,
	}
	_, err := l.svcCtx.UserInfo.Update(ui)
	if err != nil {
		return errors.New("出错啦！😭")
	}
	return nil
}
