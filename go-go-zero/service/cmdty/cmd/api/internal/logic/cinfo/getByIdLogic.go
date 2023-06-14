package cinfo

import (
	"context"
	"errors"
	jsoniter "github.com/json-iterator/go"
	"go-go-zero/common/utils"
	"go-go-zero/service/cmdty/cmd/api/internal/svc"
	"go-go-zero/service/cmdty/cmd/api/internal/types"
	"go-go-zero/service/cmdty/model"
	"strconv"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetByIdLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetByIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetByIdLogic {
	return &GetByIdLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

var json = jsoniter.ConfigCompatibleWithStandardLibrary

func (l *GetByIdLogic) GetById(req *types.IdTypeReq) (*model.CmdtyInfo, error) {
	var (
		id    = req.Id
		IdStr = strconv.FormatInt(id, 10)
	)
	key := utils.CmdtySellingPrepared
	if req.Type == 2 {
		key = utils.CmdtyWantPrepared
	}
	ci := new(model.CmdtyInfo)
	result, err := l.svcCtx.Redis.HGet(l.ctx, key, IdStr).Result()
	if result != "" && err == nil {
		err := json.Unmarshal([]byte(result), ci)
		if err != nil {
			return nil, errors.New("出错了！😢")
		}
		return ci, nil
	}
	key = utils.CmdtyInfo + IdStr
	ci.Id = id
	has, err := l.svcCtx.Xorm.Table("cmdty_info").Get(ci)
	if err != nil {
		return nil, errors.New("出错啦😥")
	}
	return
}
