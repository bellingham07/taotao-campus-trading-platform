package logic

import (
	"context"
	"errors"
	"github.com/yitter/idgenerator-go/idgen"
	"go-go-zero/service/user/model"
	"golang.org/x/crypto/bcrypt"
	"math/rand"
	"strconv"
	"strings"

	"go-go-zero/service/user/cmd/api/internal/svc"
	"go-go-zero/service/user/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type RegisterLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RegisterLogic {
	return &RegisterLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RegisterLogic) Register(req *types.RegisterReq) (resp *types.BaseResp, err error) {
	password1 := strings.TrimSpace(req.Password1)
	password2 := strings.TrimSpace(req.Password2)
	if equal := strings.Compare(password1, password2); equal != 0 {
		return nil, errors.New("两次输入的密码不一样！")
	}
	password, err := bcrypt.GenerateFromPassword([]byte(password1), bcrypt.DefaultCost)
	randNum := rand.Int31()
	ui := &model.UserInfo{
		Id:       idgen.NextId(),
		Username: req.Username,
		Password: string(password),
		Name:     "user" + strconv.Itoa(int(randNum)),
	}
	result, err := model.UserInfoModel.Insert(l.svcCtx.UserInfo, l.ctx, ui)
	if r, err := result.RowsAffected(); r == 1 && err == nil {
		resp = &types.BaseResp{
			Code: 1,
			Msg:  "注册成功😊",
		}
		return
	}
	return nil, errors.New("注册失败🥲")
}
