package history

import (
	"context"
	"github.com/redis/go-redis/v9"
	"go-go-zero/common/utils"
	"strconv"
	"time"

	"github.com/zeromicro/go-zero/core/logx"
	"go-go-zero/service/cmdty/cmd/api/internal/svc"
)

type UpdateHistoryLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateHistoryLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateHistoryLogic {
	return &UpdateHistoryLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateHistoryLogic) UpdateHistory(id, userId int64) {
	key := utils.CmdtyHistory + strconv.FormatInt(userId, 10)
	member := redis.Z{
		Score:  float64(time.Now().Unix()),
		Member: id,
	}
	err := l.svcCtx.Redis.ZAdd(l.ctx, key, member).Err()
	if err != nil {
		logx.Infof("[DB ERROR] 更新redis足迹失败，userId：%d\n", userId)
	}
}
