package history

import (
	"context"
	"errors"
	"github.com/redis/go-redis/v9"
	"go-go-zero/common/utils"
	"go-go-zero/service/cmdty/model"
	"strconv"
	"time"

	"go-go-zero/service/cmdty/cmd/api/internal/svc"
	"go-go-zero/service/cmdty/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ListHistoryLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewListHistoryLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ListHistoryLogic {
	return &ListHistoryLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ListHistoryLogic) ListHistory() ([]*types.HistoryResp, error) {
	var (
		zqualified []redis.Z
		resp       = make([]*types.HistoryResp, 0)
		key        = utils.CmdtyHistory + "408301323265285"
	)
	// 1 先从redis中取出浏览记录的id
	zs, err := l.svcCtx.Redis.ZRevRangeWithScores(l.ctx, key, 0, -1).Result()
	if err != nil || zs == nil {
		return nil, errors.New("出错啦，请刷新😊")
	}
	// 2 筛选出符合时间的记录
	l.findQualifiedZs(&zqualified, &zs, key)
	// 3 先从redis中获取，没有，就去mysql取
	for _, z := range zqualified {
		id := z.Member.(int64)
		key = utils.CmdtyInfo + strconv.FormatInt(id, 10)
		ciMap, err := l.svcCtx.Redis.HGetAll(l.ctx, key).Result()
		if err == nil {
			price, _ := strconv.ParseFloat(ciMap["price"], 64)
			hr := &types.HistoryResp{
				Id:    id,
				Cover: ciMap["cover"],
				Price: price,
			}
			resp = append(resp, hr)
		}
		ci := &model.CmdtyInfo{Id: id}
		has, err := l.svcCtx.CmdtyInfo.Cols("price", "cover").Get(ci)
		if !has || err != nil {
			logx.Infof("[DB ERROR] ListHistory 获取商品信息失败 %v\n", err)
			continue
		}
		hr := &types.HistoryResp{
			Id:    id,
			Cover: ci.Cover,
			Price: ci.Price,
		}
		resp = append(resp, hr)
	}
	return resp, nil
}

func (l *ListHistoryLogic) findQualifiedZs(zqualified, zs *[]redis.Z, key string) {
	var thirtyDaysBefore = time.Now().Local().AddDate(0, 0, -30)
	for idx, z := range *zs {
		createTime := time.Unix(int64(z.Score), 0).Local()
		if createTime.Before(thirtyDaysBefore) {
			*zqualified = (*zs)[:idx]
			go l.removeExpiredRecords((*zs)[idx:], key)
			break
		}
	}
}

func (l *ListHistoryLogic) removeExpiredRecords(zs []redis.Z, key string) {
	var members = make([]interface{}, 0)
	for _, z := range zs {
		members = append(members, z.Member)
	}
	err := l.svcCtx.Redis.ZRem(l.ctx, key, members...).Err()
	if err != nil {
		logx.Debugf("[REDIS ERROR] ListHistory 删除过期足迹错误 %v\n", err)
	}
}
