package logic

import (
	"context"
	"errors"
	"go-go-zero/common/utils"
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

func (l *LoginLogic) Login(req *types.LoginReq) (resp *types.LoginResp, err error) {
	ui := l.svcCtx.UserInfo.QueryInfoByUsername(req.Username)
	if err = bcrypt.CompareHashAndPassword([]byte(ui.Password), []byte(req.Password)); err != nil {
		return nil, errors.New("账号或密码错误！🥲")
	}
	token, err := utils.GenToken(ui)
	if err != nil {
		return nil, errors.New("登录错误！请稍后🥲")
	}
	resp = &types.LoginResp{
		BaseResp: types.BaseResp{
			Code: 1,
			Msg:  "登录成功😊",
		},
		Token: token,
	}
	return
}
