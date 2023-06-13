package logic

import (
	"context"
	jsoniter "github.com/json-iterator/go"
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
	json   = jsoniter.ConfigCompatibleWithStandardLibrary
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
			ticker.Reset(31 * time.Minute)
		}
	}
}

func (l *Cmdty2RedisLogic) SellingCmdty2Redis() {
	sellingCmdty := make([]*model.CmdtyInfo, 0)
	_ = l.svcCtx.Xorm.Table("cmdty_info").
		Where("status = ? AND type = ?", 2, 1).
		Desc("publish_at").Limit(100, 0).Find(sellingCmdty)
	data, _ := json.Marshal(sellingCmdty)
	l.svcCtx.RedisClient.Set(l.ctx, utils.CmdtySellingPrepared, data, 31*time.Minute)
}

func (l *Cmdty2RedisLogic) WantCmdty2Redis() {

	wantCmdty := make([]*model.CmdtyInfo, 0)
	_ = l.svcCtx.Xorm.Table("cmdty_info").
		Where("status = ? AND type = ?", 2, 2).
		Desc("publish_at").Limit(100, 0).Find(wantCmdty)
	data, _ := json.Marshal(wantCmdty)
	l.svcCtx.RedisClient.Set(l.ctx, utils.CmdtyWantPrepared, data, 31*time.Minute)
}
