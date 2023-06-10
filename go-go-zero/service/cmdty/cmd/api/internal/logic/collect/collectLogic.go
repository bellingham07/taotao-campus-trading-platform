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

func (l *CollectLogic) Collect(req *types.IdReq) error {
	key := utils.CmdtyCollect + strconv.FormatInt(req.Id, 10)
	var userId int64 = 408301323265285
	userIdStr := "408301323265285"
	r, err := l.svcCtx.RedisClient.SAdd(l.ctx, key, userIdStr).Result()
	if r == 0 {
		logx.Debugf("[REDIS ERROR] collect 收藏失败 " + err.Error())
		return errors.New("好啦好啦，知道你喜欢了！但不能再次收藏哦😚")
	}
	mqLogic := mq.NewRabbitMQLogic(l.ctx, l.svcCtx)
	mq.CollectUpdatePublisher(key, userId, true, mqLogic)
	return nil
}
