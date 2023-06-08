package collect

import (
	"context"
	"errors"
	"go-go-zero/common/utils"
	"go-go-zero/service/cmdty/cmd/api/internal/logic/mq"
	"strconv"

	"go-go-zero/service/cmdty/cmd/api/internal/svc"
	"go-go-zero/service/cmdty/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CollectLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCollectLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CollectLogic {
	return &CollectLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CollectLogic) Collect(req *types.IdReq) (resp *types.BaseResp, err error) {
	key := utils.CmdtyCollect + strconv.FormatInt(req.Id, 10)
	var userId int64 = 408301323265285
	userIdStr := "408301323265285"
	r, _ := l.svcCtx.RedisClient.SAdd(l.ctx, key, userIdStr).Result()
	if r == 0 {
		//logx.Error("[REDIS ERROR] collect " + err.Error())
		return nil, errors.New("好啦好啦，知道你喜欢了！但不能再次收藏哦😚")
	}
	mqLogic := mq.NewRabbitMQLogic(l.ctx, l.svcCtx)
	mq.CollectUpdatePublisher(key, userId, true, mqLogic)
	resp = &types.BaseResp{
		Code: 1,
		Msg:  "收藏成功😊",
	}
	return
}
