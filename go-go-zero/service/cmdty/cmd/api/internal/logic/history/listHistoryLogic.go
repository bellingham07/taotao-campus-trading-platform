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
	resp := make([]*types.HistoryResp, 0)
	key := utils.CmdtyHistory + "408301323265285"
	zs, err := l.svcCtx.Redis.ZRevRangeWithScores(l.ctx, key, 0, -1).Result()
	now := time.Now()
	var zqualified []redis.Z
	for idx, z := range zs {
		createTime := time.Unix(int64(z.Score), 0)
		if createTime.Before(now.AddDate(0, 0, -30)) {
			zqualified = zs[:idx]
			go func() {
				var ids []interface{}
				for _, z := range zs[idx:] {
					ids = append(ids, z.Member)
				}
				err = l.svcCtx.Redis.ZRem(l.ctx, key, ids).Err()
				if err != nil {
					logx.Debugf("[REDIS ERROR] ListHistory 删除过期足迹错误 " + err.Error())
				}
			}()
			break
		}
	}
	// 从redis中获取，缓存中没有，就去mysql取
	for _, z := range zqualified {
		idStr := z.Member.(string)
		id, _ := strconv.ParseInt(idStr, 10, 64)
		key := utils.CmdtyInfo + idStr
		cir, err := l.svcCtx.Redis.HGetAll(l.ctx, key).Result()
		if err != nil {
			// 缓存中没有了
			cidb, err := l.svcCtx.CmdtyInfo.FindOne(l.ctx, id)
			if err != nil {
				logx.Debugf("[DB ERROR] ListHistory mysql 获取 cmdtyInfo 错误 " + err.Error())
			} else {
				resp = append(resp, &types.HistoryResp{
					Id:    cidb.Id,
					Cover: cidb.Cover,
					Price: cidb.Price,
				})
			}
		} else {
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
