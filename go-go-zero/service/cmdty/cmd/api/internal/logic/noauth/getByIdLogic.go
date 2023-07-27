package noauth

import (
	"context"
	"errors"
	"go-go-zero/common/utils"
	"go-go-zero/service/cmdty/cmd/api/internal/logic/history"
	"go-go-zero/service/cmdty/model"
	__cmdty "go-go-zero/service/file/cmd/rpc/types"
	__user "go-go-zero/service/user/cmd/rpc/types"
	"strconv"
	"sync"
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

func (l *GetByIdTypeLogic) GetByIdTypeLogic(req *types.IdReq, userId int64) (interface{}, error) {
	var (
		wg           sync.WaitGroup
		id           = req.Id
		IdStr        = strconv.FormatInt(id, 10)
		resp         = make(map[string]interface{})
		cmdtyInfoKey = utils.CmdtyInfo + IdStr
		userInfoMap  = make(map[string]interface{})
		pics         = make([]interface{}, 0)
	)

	data, ownerId, err := l.getById(cmdtyInfoKey, id)

	wg.Add(2)
	go func() {
		defer wg.Done()
		idReq := &__user.IdReq{Id: ownerId}
		resp, err := l.svcCtx.UserRpc.RetrieveNameAndAvatar(l.ctx, idReq)
		if err != nil || resp.Code == -1 {
			logx.Errorf("[RPC ERROR] GetByIdTypeLogic 调用rpc获取用户昵称和头像失败 %v\n", err)
			return
		}
		userInfoMap["id"] = ownerId
		userInfoMap["avatar"] = resp.Avatar
		userInfoMap["name"] = resp.Name
	}()

	go func() {
		defer wg.Done()
		cmdtyPicsReq := &__cmdty.CmdtyPicsReq{Id: id}
		resp, err := l.svcCtx.FileRpc.GetCmdtyPicsByOrder(l.ctx, cmdtyPicsReq)
		if err != nil || resp.Code == -1 {
			logx.Errorf("[RPC ERROR] GetByIdTypeLogic 调用rpc获取用户昵称和头像失败 %v\n", err)
			return
		}
		for _, pic := range resp.Pics {
			pics = append(pics, pic)
		}
	}()

	// 判断当前访问用户是否登录，更新其足迹，并判断是否收藏该商品
	if userId != 0 {
		go l.updateHistory(id, userId)
		resp["isCollected"] = l.isCollected(userId, utils.CmdtyCollect+IdStr)
	}
	resp["cmdtyInfo"] = data
	wg.Wait()
	resp["userInfo"] = userInfoMap
	resp["pics"] = pics
	return resp, err
}

func (l *GetByIdTypeLogic) getById(key string, id int64) (interface{}, int64, error) {
	// 2 不是，先从redis中查hash
	cimap, err := l.svcCtx.Redis.HGetAll(l.ctx, key).Result()
	if cimap != nil && err == nil {
		// 2.1 解决缓存穿透
		if _, ok := cimap["nil"]; ok {
			return nil, 0, errors.New("没有这个物品！😶‍🌫️")
		}
		// 2.2 查到了返回
		if idStr, ok := cimap["id"]; ok && idStr == strconv.FormatInt(id, 10) {
			go l.svcCtx.Redis.Expire(l.ctx, key, 5*time.Minute)
			userId, _ := strconv.ParseInt(cimap["userId"], 10, 64)
			return cimap, userId, nil
		}
	}

	// 3 redis中没有，查数据库
	return l.loadFromDB(key, id)
}

func (l *GetByIdTypeLogic) loadFromDB(key string, id int64) (interface{}, int64, error) {
	var ci = &model.CmdtyInfo{Id: id}
	has, err := l.svcCtx.CmdtyInfo.Get(ci)
	if !has && err == nil {
		go l.svcCtx.Redis.HSet(l.ctx, key, map[string]string{"nil": ""})
		return nil, 0, errors.New("没有这个物品！😶‍🌫️")
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
	return data, ci.UserId, nil
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
