package uinfo

import (
	"context"
	"errors"
	"go-go-zero/common/utils"
	"go-go-zero/service/user/model"
	"strconv"
	"time"

	"go-go-zero/service/user/cmd/api/internal/svc"
	"go-go-zero/service/user/cmd/api/internal/types"

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

func (l *GetByIdLogic) GetById(req *types.IdReq) (interface{}, error) {
	var (
		id    = req.Id
		idStr = strconv.FormatInt(req.Id, 10)
		key   = utils.UserInfo + idStr
		ui    = &model.UserInfo{Id: id}
	)

	uiMap, err := l.svcCtx.Redis.HGetAll(l.ctx, key).Result()
	if uiMap != nil && err == nil {
		if v, ok := uiMap["nil"]; ok && v == "" {
			go l.svcCtx.Redis.Expire(l.ctx, key, time.Minute)
			return nil, errors.New("找不到这个用户！😢")
		}
		if v, ok := uiMap["id"]; ok && v == idStr {
			go l.svcCtx.Redis.Expire(l.ctx, key, 30*time.Minute)
			return uiMap, nil
		}
	}

	has, err := l.svcCtx.Xorm.Table("user_info").Get(ui)
	if !has && err != nil {
		go l.svcCtx.Redis.HSet(l.ctx, key, map[string]string{"nil": ""})
		return nil, errors.New("找不到这个用户！😶‍🌫️")
	}

	data := map[string]interface{}{
		"id":       ui.Id,
		"username": ui.Username,
		"password": ui.Password,
		"name":     ui.Name,
		"gender":   ui.Gender,
		"phone":    ui.Phone,
		"avatar":   ui.Avatar,
		"intro":    ui.Intro,
		"location": ui.Location,
		"like":     ui.Like,
		"status":   ui.Status,
		"done":     ui.Done,
		"call":     ui.Call,
		"fans":     ui.Fans,
		"follow":   ui.Follow,
		"positive": ui.Like,
		"negative": ui.Negative,
	}
	go l.svcCtx.Redis.HSet(l.ctx, key, data)
	return data, nil
}
