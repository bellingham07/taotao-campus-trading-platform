package logic

import (
	"context"
	"errors"
	"github.com/yitter/idgenerator-go/idgen"
	"go-go-zero/service/user/cmd/api/internal/svc"
	"go-go-zero/service/user/cmd/api/internal/types"
	"go-go-zero/service/user/model"
	"golang.org/x/crypto/bcrypt"
	"math/rand"
	"strconv"
	"strings"

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

func (l *RegisterLogic) Register(req *types.RegisterReq) error {
	var (
		password1 = strings.TrimSpace(req.Password1)
		password2 = strings.TrimSpace(req.Password2)
	)
	equal := strings.Compare(password1, password2)
	if equal != 0 {
		return errors.New("两次输入的密码不一样！")
	}
	password, err := bcrypt.GenerateFromPassword([]byte(password1), bcrypt.DefaultCost)
	randNum := rand.Int31()
	ui := &model.UserInfo{
		Id:       idgen.NextId(),
		Username: req.Username,
		Password: string(password),
		Name:     "user" + strconv.Itoa(int(randNum)),
		//LastLogin: time.Now(),
	}
	_, err = l.svcCtx.Xorm.Table("user_info").Insert(ui)
	if err == nil {
		return nil
	} else if strings.Contains(err.Error(), "Duplicate") {
		return errors.New("来晚了一步，该账号已经被抢走了😭")
	}
	return errors.New("注册失败啦😥")
}
