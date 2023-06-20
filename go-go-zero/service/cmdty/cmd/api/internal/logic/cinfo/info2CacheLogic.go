package cinfo

import (
	"context"
	"github.com/zeromicro/go-zero/core/logx"
	"go-go-zero/common/utils"
	"go-go-zero/service/cmdty/cmd/api/internal/svc"
	"go-go-zero/service/cmdty/model"
	"sync"
	"time"
)

type Cmdty2RedisLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCmdty2RedisLogic(ctx context.Context, svcCtx *svc.ServiceContext) *Cmdty2RedisLogic {
	return &Cmdty2RedisLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

var (
	ticker = time.NewTicker(30 * time.Minute)
	wg     sync.WaitGroup
)

func (l *Cmdty2RedisLogic) Cmdty2Redis() {
	for {
		select {
		case <-ticker.C:
			wg.Add(2)
			go func() {
				l.SellingCmdty2Redis()
				wg.Done()
			}()
			go func() {
				l.WantCmdty2Redis()
				wg.Done()
			}()
			wg.Wait()
			ticker.Reset(30 * time.Minute)
		}
	}
}

func (l *Cmdty2RedisLogic) SellingCmdty2Redis() {
	cmdtyPrepared := make(map[int64][]byte)
	sellingCmdty := make([]*model.CmdtyInfo, 0)
	// 查库
	err := l.svcCtx.CmdtyInfo.Where("`status` = ? AND `type` = ?", 2, 1).
		Desc("publish_at").Limit(100, 0).Find(sellingCmdty)
	if err != nil {
		logx.Infof("[DB ERROR] SellingCmdty2Redis 数据库查询错误 %v\n", err)
		return
	}
	// 序列化
	for _, v := range sellingCmdty {
		data, err := l.svcCtx.Json.Marshal(v)
		if err == nil {
			cmdtyPrepared[v.Id] = data
			continue
		}
		logx.Infof("[JSON MARSHAL ERROR] SellingCmdty2Redis 序列化数据错误 %v\n", err)
	}
	// 使用pipeline发给redis
	pipeline := l.svcCtx.Redis.Pipeline()
	pipeline.HSet(l.ctx, utils.CmdtySellingPrepared, cmdtyPrepared)
	pipeline.Expire(l.ctx, utils.CmdtySellingPrepared, 31*time.Minute)
	_, err = pipeline.Exec(l.ctx)
	if err != nil {
		logx.Infof("[REDIS ERROR] SellingCmdty2Redis redis执行错误 %v\n", err)
	}
}

func (l *Cmdty2RedisLogic) WantCmdty2Redis() {
	cmdtyPrepared := make(map[int64][]byte)
	wantCmdty := make([]*model.CmdtyInfo, 0)
	// 查库
	err := l.svcCtx.CmdtyInfo.Where("`status` = ? AND `type` = ?", 2, 2).
		Desc("publish_at").Limit(100, 0).Find(wantCmdty)
	if err != nil {
		logx.Infof("[DB ERROR] SellingCmdty2Redis 数据库查询错误 %v\n", err)
		return
	}
	// 序列化
	for _, v := range wantCmdty {
		data, _ := l.svcCtx.Json.Marshal(v)
		if err == nil {
			cmdtyPrepared[v.Id] = data
			continue
		}
		logx.Infof("[JSON MARSHAL ERROR] SellingCmdty2Redis 序列化数据错误 %v\n", err)
	}
	// 使用pipeline发给redis
	pipeline := l.svcCtx.Redis.Pipeline()
	pipeline.HSet(l.ctx, utils.CmdtyWantPrepared, cmdtyPrepared)
	pipeline.Expire(l.ctx, utils.CmdtyWantPrepared, 31*time.Minute)
	_, err = pipeline.Exec(l.ctx)
	if err != nil {
		logx.Infof("[REDIS ERROR] SellingCmdty2Redis redis执行错误 %v\n", err)
	}
}
