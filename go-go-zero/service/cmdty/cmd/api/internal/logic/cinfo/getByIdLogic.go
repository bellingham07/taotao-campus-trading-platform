package cinfo

import (
	"context"
	"errors"
	"go-go-zero/common/utils"
	"go-go-zero/service/cmdty/cmd/api/internal/svc"
	"go-go-zero/service/cmdty/cmd/api/internal/types"
	"go-go-zero/service/cmdty/model"
	"strconv"
	"time"

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

func (l *GetByIdLogic) GetById(req *types.IdTypeReq) (interface{}, error) {
	var (
		id    = req.Id
		IdStr = strconv.FormatInt(id, 10)
		key   = utils.CmdtySellingPrepared
	)
	if req.Type == 2 {
		key = utils.CmdtyWantPrepared
	}
	// 1 先检查是不是前100条已缓存的数据
	ci := new(model.CmdtyInfo)
	result, err := l.svcCtx.Redis.HGet(l.ctx, key, IdStr).Result()
	if result != "" && err == nil {
		err = l.svcCtx.Json.Unmarshal([]byte(result), ci)
		if err != nil {
			return nil, errors.New("出错了！😢")
		}
		return ci, nil
	}
	// 2 不是，先从redis中查hash
	key = utils.CmdtyInfo + IdStr
	cimap, err := l.svcCtx.Redis.HGetAll(l.ctx, key).Result()
	if cimap != nil && err == nil {
		// 2.1 解决缓存穿透
		if _, ok := cimap["nil"]; ok {
			return nil, errors.New("没有这个物品！😶‍🌫️")
		}
		// 2.2 查到了返回
		if id, ok := cimap["id"]; ok && id == IdStr {
			go l.svcCtx.Redis.Expire(l.ctx, key, 5*time.Minute)
			return cimap, nil
		}
	}
	// 3 redis中没有，查数据库
	ci.Id = id
	has, err := l.svcCtx.Xorm.Table("cmdty_info").Get(ci)
	if !has && err == nil {
		l.svcCtx.Redis.HSet(l.ctx, key, map[string]string{"nil": ""})
		return nil, errors.New("没有这个物品！😶‍🌫️")
	}
	data := map[string]interface{}{
		"id":        ci.Id,
		"userId":    ci.UserId,
		"cover":     ci.Cover,
		"tag":       ci.Tag,
		"price":     ci.Price,
		"brand":     ci.Brand,
		"model":     ci.Model,
		"intro":     ci.Intro,
		"old":       ci.Old,
		"status":    ci.Status,
		"createAt":  ci.CreateAt,
		"publishAt": ci.PublishAt,
		"view":      ci.View,
		"collect":   ci.Collect,
		"type":      ci.Type,
		"like":      ci.Like,
	}
	go l.svcCtx.Redis.HSet(l.ctx, key, data)
	return data, nil
}
