package history

import (
	"context"
	"github.com/redis/go-redis/v9"
	"go-go-zero/common/utils"
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
		resp       = make([]*types.HistoryResp, 0)
		key        = utils.CmdtyHistory + "408301323265285"
		zqualified []redis.Z
	)
	zs, err := l.svcCtx.Redis.ZRevRangeWithScores(l.ctx, key, 0, -1).Result()
	if err == nil {
		l.findQualifiedZs(zqualified, zs, key)
		// 从redis中获取，缓存中没有，就去mysql取
		for _, z := range zqualified {
			id := z.Member.(int64)
			cidb, err := l.svcCtx.CmdtyInfo.Get()
			if err != nil {
				logx.Debugf("[DB ERROR] ListHistory mysql 获取 cmdtyInfo 错误 %v\n", err)
			} else {
				resp = append(resp, &types.HistoryResp{
					Id:    cidb.Id,
					Cover: cidb.Cover,
					Price: cidb.Price,
				})
			}

			id, _ := strconv.ParseInt(cir["Id"], 10, 64)
			price, _ := strconv.ParseFloat(cir["Price"], 64)
			hr := &types.HistoryResp{
				Id:    id,
				Cover: cir["Cover"],
				Price: price,
			}
			resp = append(resp, hr)

		}
	}

	return resp, nil
}

func (l *ListHistoryLogic) findQualifiedZs(zqualified, zs []redis.Z, key string) {
	var thirtyDaysBefore = time.Now().Local().AddDate(0, 0, -30)
	for idx, z := range zs {
		createTime := time.Unix(int64(z.Score), 0).Local()
		if createTime.Before(thirtyDaysBefore) {
			zqualified = zs[:idx]
			go l.removeExpiredRecords(zs[idx:], key)
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
