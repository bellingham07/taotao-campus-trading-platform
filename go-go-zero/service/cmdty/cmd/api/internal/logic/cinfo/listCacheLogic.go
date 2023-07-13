package cinfo

import (
	"context"
	"go-go-zero/common/utils"
	"go-go-zero/service/cmdty/cmd/api/internal/svc"
	"go-go-zero/service/cmdty/model"
	"sync"

	"github.com/zeromicro/go-zero/core/logx"
)

type ListCacheLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewListCacheByTypeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ListCacheLogic {
	return &ListCacheLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ListCacheLogic) ListCacheByType() ([]model.CmdtyInfo, error) {
	var (
		wg      sync.WaitGroup
		data    = make(map[string][]model.CmdtyInfo)
		sellKey = utils.CmdtySellingPrepared
		wantKey = utils.CmdtyWantPrepared
	)

	wg.Add(2)
	go func() {
		err := l.FetchFromRedis(sellKey)
		if err != nil {

		}
		wg.Done()
	}()

	go func() {
		err := l.FetchFromRedis(sellKey)
		if err != nil {

		}
		wg.Done()
	}()

	cis := make([]model.CmdtyInfo, 0)
	for _, ciStr := range cisMap {
		var ci model.CmdtyInfo
		err = l.svcCtx.Json.Unmarshal([]byte(ciStr), &ci)
		if err != nil {
			continue
		}
		cis = append(cis, ci)
	}
	return cis, nil
}

func (l *ListCacheLogic) FetchFromRedis(key string) map[string]string {
	cache, err := l.svcCtx.Redis.HGetAll(l.ctx, key).Result()
	if err != nil {
		return nil
	}
	return cache
}
