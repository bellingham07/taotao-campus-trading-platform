package noauth

import (
	"context"
	"errors"
	"go-go-zero/common/utils"
	"go-go-zero/service/cmdty/cmd/api/internal/logic/history"
	"go-go-zero/service/cmdty/model"
	"strconv"
	"time"

	"go-go-zero/service/cmdty/cmd/api/internal/svc"
	"go-go-zero/service/cmdty/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetByIdTypeLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetByIdTypeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetByIdTypeLogic {
	return &GetByIdTypeLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetByIdTypeLogic) GetByIdTypeLogic(req *types.IdTypeReq, userId int64) (interface{}, error) {
	var (
		id               = req.Id
		IdStr            = strconv.FormatInt(id, 10)
		ci               = new(model.CmdtyInfo)
		resp             = map[string]interface{}{"isCollected": 0}
		cmdtyInfoKey     = utils.CmdtyInfo + IdStr
		preparedCacheKey = utils.CmdtySellingPrepared
	)

	if req.Type == 2 {
		preparedCacheKey = utils.CmdtyWantPrepared
	}

	// 1 先检查是不是前100条已缓存的数据
	result, err := l.svcCtx.Redis.HGet(l.ctx, preparedCacheKey, IdStr).Result()
	if result != "" && err == nil {
		err = l.svcCtx.Json.Unmarshal([]byte(result), ci)
		if err != nil {
			return nil, errors.New("出错了！😢")
		}
		return ci, nil
	}

	// 2 不是，先从redis中查hash再查mysql
	data, err := l.getById(cmdtyInfoKey, id)

	// 当前访问用户已登录，更新其足迹，并判断是否收藏该商品
	if userId != 0 {
		go l.updateHistory(id, userId)
		resp["isCollected"] = l.isCollected(userId, utils.CmdtyCollect+IdStr)
	}
	resp["cmdty"] = data
	return resp, err
}

func (l *GetByIdTypeLogic) getById(key string, id int64) (interface{}, error) {
	// 2 不是，先从redis中查hash
	cimap, err := l.svcCtx.Redis.HGetAll(l.ctx, key).Result()
	if cimap != nil && err == nil {
		// 2.1 解决缓存穿透
		if _, ok := cimap["nil"]; ok {
			return nil, errors.New("没有这个物品！😶‍🌫️")
		}
		// 2.2 查到了返回
		if idMap, ok := cimap["id"]; ok && idMap == strconv.FormatInt(id, 10) {
			go l.svcCtx.Redis.Expire(l.ctx, key, 5*time.Minute)
			return cimap, nil
		}
	}

	// 3 redis中没有，查数据库
	return l.getFromCmdtyInfo(key, id)
}

func (l *GetByIdTypeLogic) getFromCmdtyInfo(key string, id int64) (interface{}, error) {
	var ci = &model.CmdtyInfo{Id: id}
	has, err := l.svcCtx.CmdtyInfo.Get(ci)
	if !has && err == nil {
		go l.svcCtx.Redis.HSet(l.ctx, key, map[string]string{"nil": ""})
		return nil, errors.New("没有这个物品！😶‍🌫️")
	}
	data := map[string]interface{}{
		"id":          ci.Id,
		"userId":      ci.UserId,
		"brief_intro": ci.BriefIntro,
		"cover":       ci.Cover,
		"tag":         ci.Tag,
		"price":       ci.Price,
		"brand":       ci.Brand,
		"model":       ci.Model,
		"intro":       ci.Intro,
		"old":         ci.Old,
		"status":      ci.Status,
		"createAt":    ci.CreateAt,
		"publishAt":   ci.PublishAt,
		"view":        ci.View,
		"collect":     ci.Collect,
		"type":        ci.Type,
		"like":        ci.Like,
	}
	go l.svcCtx.Redis.HSet(l.ctx, key, data)
	return data, nil
}

func (l *GetByIdTypeLogic) isCollected(userId int64, collectKey string) int8 {

	// 如果是别人的商品，就需要判断有没有收藏过
	isMember, err := l.svcCtx.Redis.SIsMember(l.ctx, collectKey, userId).Result()
	if isMember && err == nil {
		return 1
	}
	return 0
}

func (l *GetByIdTypeLogic) updateHistory(id, userId int64) {
	uhLogic := history.NewUpdateHistoryLogic(l.ctx, l.svcCtx)
	uhLogic.UpdateHistory(id, userId)
}
