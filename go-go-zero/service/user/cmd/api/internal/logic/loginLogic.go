package logic

import (
	"context"
	"errors"
	"go-go-zero/common/utils"
	"go-go-zero/service/user/model"
	"golang.org/x/crypto/bcrypt"

	"go-go-zero/service/user/cmd/api/internal/svc"
	"go-go-zero/service/user/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type LoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LoginLogic) Login(req *types.LoginReq) (string, error) {
	var ui *model.UserInfo
	has, err := l.svcCtx.UserInfo.Cols("username", "password").
		Where("username = ?", req.Username).Get(ui)
	if !has || err != nil {
		return "", errors.New("账号或密码错误！🥲")
	}
	err = bcrypt.CompareHashAndPassword([]byte(ui.Password), []byte(req.Password))
	if err != nil {
		return "", errors.New("账号或密码错误！🥲")
	}
	token, err := utils.GenToken(ui)
	if err != nil {
		return "", errors.New("登录错误！请稍后🥲")
	}
	return token, nil
}
