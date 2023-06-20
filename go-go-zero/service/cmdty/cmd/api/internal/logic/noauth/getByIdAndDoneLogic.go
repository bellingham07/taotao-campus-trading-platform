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

type GetByIdAndDoneLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetByIdDoneTypeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetByIdAndDoneLogic {
	return &GetByIdAndDoneLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetByIdAndDoneLogic) GetByIdDoneTypeLogic(req *types.IdDoneTypeReq, userId int64) (interface{}, error) {
	var (
		id    = req.Id
		done  = req.Done
		IdStr = strconv.FormatInt(id, 10)
		key1  = utils.CmdtySellingPrepared
		key2  = utils.CmdtyInfo + IdStr
		ci    = new(model.CmdtyInfo)
		resp  = make(map[string]interface{})
	)
	if done == 1 {
		data, err := l.getByIdAndDone(key2, id, done)
		if userId != 0 {
			resp = l.isCollected(id, userId, utils.CmdtyCollect+IdStr)
		}
		resp["data"] = data
		return resp, err
	}
	if req.Type == 2 {
		key1 = utils.CmdtyWantPrepared
	}
	// 1 先检查是不是前100条已缓存的数据
	result, err := l.svcCtx.Redis.HGet(l.ctx, key1, IdStr).Result()
	if result != "" && err == nil {
		err = l.svcCtx.Json.Unmarshal([]byte(result), ci)
		if err != nil {
			return nil, errors.New("出错了！😢")
		}
		return ci, nil
	}
	// 2 不是，先从redis中查hash再查mysql
	data, err := l.getByIdAndDone(key2, id, 0)
	if userId != 0 {
		l.updateHistory(id, userId)
		resp = l.isCollected(id, userId, utils.CmdtyCollect+IdStr)
	}
	resp["data"] = data
	return resp, err
}

func (l *GetByIdAndDoneLogic) getByIdAndDone(key string, id, done int64) (interface{}, error) {
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
	if done == 1 {
		return l.getFromCmdtyDone(key, id)
	}
	return l.getFromCmdtyInfo(key, id)
}

func (l *GetByIdAndDoneLogic) getFromCmdtyDone(key string, id int64) (interface{}, error) {
	var cd = &model.CmdtyInfo{Id: id}
	has, err := l.svcCtx.CmdtyDone.Get(cd)
	if !has && err == nil {
		go l.svcCtx.Redis.HSet(l.ctx, key, map[string]string{"nil": ""})
		return nil, errors.New("没有这个物品！😶‍🌫️")
	}
	data := map[string]interface{}{
		"id":        cd.Id,
		"userId":    cd.UserId,
		"cover":     cd.Cover,
		"tag":       cd.Tag,
		"price":     cd.Price,
		"brand":     cd.Brand,
		"model":     cd.Model,
		"intro":     cd.Intro,
		"old":       cd.Old,
		"status":    cd.Status,
		"createAt":  cd.CreateAt,
		"publishAt": cd.PublishAt,
		"view":      cd.View,
		"collect":   cd.Collect,
		"type":      cd.Type,
		"like":      cd.Like,
	}
	go l.svcCtx.Redis.HSet(l.ctx, key, data)
	return data, nil
}

func (l *GetByIdAndDoneLogic) getFromCmdtyInfo(key string, id int64) (interface{}, error) {
	var ci = &model.CmdtyInfo{Id: id}
	has, err := l.svcCtx.CmdtyInfo.Get(ci)
	if !has && err == nil {
		go l.svcCtx.Redis.HSet(l.ctx, key, map[string]string{"nil": ""})
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

func (l *GetByIdAndDoneLogic) isCollected(id, userId int64, collectKey string) map[string]interface{} {
	var resp = make(map[string]interface{})

	// 如果是别人的商品，就需要判断有没有收藏过
	isMember, err := l.svcCtx.Redis.SIsMember(l.ctx, collectKey, userId).Result()
	if isMember && err == nil {
		resp["isCollected"] = 1
		return resp
	}
	resp["isCollected"] = 0
	return resp
}

func (l *GetByIdAndDoneLogic) updateHistory(id, userId int64) {
	uhLogic := history.NewUpdateHistoryLogic(l.ctx, l.svcCtx)
	uhLogic.UpdateHistory(id, userId)
}
