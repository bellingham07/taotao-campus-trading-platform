package noauth

import (
	"context"
	"errors"
	"go-go-zero/common/utils"
	"go-go-zero/service/user/model"
	"golang.org/x/crypto/bcrypt"
	"strconv"
	"time"

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
	var ui = new(model.UserInfo)
	has, err := l.svcCtx.UserInfo.Cols("id", "name", "username", "password").
		Where("username = ?", req.Username).Get(ui)
	if !has || err != nil {
		return "", errors.New("账号或密码错误！🥲")
	}

	err = bcrypt.CompareHashAndPassword([]byte(ui.Password), []byte(req.Password))
	if err != nil {
		return "", errors.New("账号或密码错误2！🥲")
	}

	token, err := utils.GenToken(ui)
	if err != nil {
		return "", errors.New("登录错误！请稍后🥲")
	}

	key := utils.UserLogin + strconv.FormatInt(ui.Id, 10)
	l.svcCtx.Redis.Set(l.ctx, key, token, 7*24*time.Hour)

	return token, nil
}
