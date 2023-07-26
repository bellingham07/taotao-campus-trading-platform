package cron

import (
	"context"
	"fmt"
	jsoniter "github.com/json-iterator/go"
	"github.com/redis/go-redis/v9"
	"github.com/robfig/cron"
	"github.com/zeromicro/go-zero/core/logx"
	"go-go-zero/common/utils"
	"go-go-zero/service/cmdty/cmd/api/internal/svc"
	"go-go-zero/service/cmdty/model"
	"strconv"
	"time"
	"xorm.io/xorm"
)

func ci2CacheLogic(svcCtx *svc.ServiceContext) {
	var c = cron.New()
	err := c.AddFunc("1 0,30 6-23 * *", exec(svcCtx))
	if err != nil {
		logx.Info("[CRON ERROR] ci2CacheLogic 任务执行失败 %v\n", err)
		panic("[CRON ERROR] ci2CacheLogic 任务执行失败")
	}

	c.Start()

	select {}
}

func exec(svcCtx *svc.ServiceContext) func() {
	return func() {
		var db2Cache = func(ciType int8) {
			var cis = queryFromDB(svcCtx.CmdtyInfo, ciType)
			if cis != nil {
				send2Cache(svcCtx.Redis, svcCtx.Json, cis, ciType)
			}
		}

		db2Cache(1)
		db2Cache(2)

		logx.Infof("[CRON SUCCESS] %v 最新商品缓存", time.Now().Format("2006-01-02 03:04:05"))
	}
}

func queryFromDB(cmdtyInfo *xorm.Session, ciType int8) [][]model.CmdtyInfo {
	var cis = make([]model.CmdtyInfo, 0)
	err := cmdtyInfo.Where("`status` = ? AND `type` = ?", 2, ciType).
		Desc("publish_at").Limit(200, 0).Find(&cis)
	if err != nil {
		logx.Infof("[DB ERROR] CRON queryFromDB 查询最新数据失败 %v\n", err)
		return nil
	}

	fmt.Println(len(cis))

	res := make([][]model.CmdtyInfo, 0)
	offset := 0
	for i := 0; i < 10; i++ {
		res = append(res, cis[offset:offset+20])
		offset += 20
	}
	return res
}

func send2Cache(redis *redis.Client, json jsoniter.API, cisPaged [][]model.CmdtyInfo, ciType int8) {
	var prefix = utils.CmdtySellNewest
	if ciType == 2 {
		prefix = utils.CmdtyWantNewest
	}

	pipeline := redis.Pipeline()
	ctx := context.Background()
	for page, cis := range cisPaged {
		key := prefix + strconv.FormatInt(int64(page), 10)
		data := make(map[string]interface{})
		for i, ci := range cis {
			ciStr, err := json.Marshal(ci)
			if err != nil {
				continue
			}
			no := strconv.FormatInt(int64(page*20+i), 10)
			data[no] = ciStr
		}
		pipeline.HSet(ctx, key, data)
		pipeline.Expire(ctx, key, 6*time.Minute)
	}
	_, err := pipeline.Exec(ctx)
	if err != nil {
		logx.Infof("[REDIS ERROR] CRON send2Cache 发送数据到redis失败 %v\n", err)
	}
}
