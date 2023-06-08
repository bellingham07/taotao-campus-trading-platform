package logic

import (
	"context"
	"github.com/yitter/idgenerator-go/idgen"
	"go-go-zero/service/user/model"
	"golang.org/x/crypto/bcrypt"
	"math/rand"
	"strconv"
	"strings"
	"time"

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
		return &types.BaseResp{Code: 0, Msg: "两次输入的密码不一样！"}, nil
	}
	password, err := bcrypt.GenerateFromPassword([]byte(password1), bcrypt.DefaultCost)
	randNum := rand.Int31()
	ui := &model.UserInfo{
		Id:        idgen.NextId(),
		Username:  req.Username,
		Password:  string(password),
		Name:      "user" + strconv.Itoa(int(randNum)),
		LastLogin: time.Now(),
	}
	_, err = model.UserInfoModel.Insert(l.svcCtx.UserInfo, l.ctx, ui)
	if err == nil {
		resp = &types.BaseResp{Code: 1, Msg: "注册成功😊"}
		return
	} else if strings.Contains(err.Error(), "Duplicate") {
		resp = &types.BaseResp{Code: 0, Msg: "来晚了一步，该账号已经被抢走了😭"}
		return resp, nil
	}
	return &types.BaseResp{Code: 0, Msg: "注册失败啦😥"}, nil
}
