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

type UncollectLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUncollectLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UncollectLogic {
	return &UncollectLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UncollectLogic) Uncollect(req *types.IdReq, userId int64) error {
	var (
		key       = utils.CmdtyCollect + strconv.FormatInt(req.Id, 10)
		userIdStr = strconv.FormatInt(userId, 10)
	)
	isMember, _ := l.svcCtx.Redis.SIsMember(l.ctx, key, userIdStr).Result()
	if isMember {
		mqLogic := mq.NewRabbitMQLogic(l.ctx, l.svcCtx)
		go mq.CollectUpdatePublisher(key, userId, false, mqLogic)
		return nil
	}
	return errors.New("你本来就没收藏人家嘛！😫")
}
