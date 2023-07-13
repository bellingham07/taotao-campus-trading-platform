package noauth

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

func NewListCacheLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ListCacheLogic {
	return &ListCacheLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ListCacheLogic) ListCache() map[string][]model.CmdtyInfo {
	var (
		wg      sync.WaitGroup
		sellKey = utils.CmdtySellingPrepared
		wantKey = utils.CmdtyWantPrepared
		resp    = make(map[string][]model.CmdtyInfo)
	)

	wg.Add(2)
	go func() {
		var cisForSale = l.fetchCache(sellKey)
		if cisForSale == nil {
			cisForSale = l.queryFromDB(1)
		}
		resp["sell"] = cisForSale
		wg.Done()
	}()
	go func() {
		var cisBeWanted = l.fetchCache(wantKey)
		if cisBeWanted == nil {
			cisBeWanted = l.queryFromDB(2)
		}
		resp["want"] = cisBeWanted
		wg.Done()
	}()
	wg.Wait()

	return resp
}

func (l *ListCacheLogic) fetchCache(key string) []model.CmdtyInfo {
	var cisMap, err = l.svcCtx.Redis.HGetAll(l.ctx, key).Result()
	if cisMap == nil || len(cisMap) == 0 || err != nil {
		logx.Infof("[REDIS ERROR] ListCache %v", err)
		return nil
	}

	cis := make([]model.CmdtyInfo, 0)
	for _, ciStr := range cisMap {
		var ci model.CmdtyInfo
		err = l.svcCtx.Json.Unmarshal([]byte(ciStr), &ci)
		if err != nil {
			continue
		}
		cis = append(cis, ci)
	}
	return cis
}

func (l *ListCacheLogic) queryFromDB(ciType int8) []model.CmdtyInfo {
	var cis = make([]model.CmdtyInfo, 0)
	err := l.svcCtx.CmdtyInfo.Desc("publish_at").
		Where("`type` = ? AND `status` = ?", ciType, 2).
		Limit(100, 0).Find(&cis)
	if err != nil {
		logx.Infof("[DB ERROR] queryFromDB %v", err)
		return nil
	}
	return cis
}
