package like

import (
	"context"
	"errors"
	"go-go-zero/common/utils"
	"go-go-zero/service/atcl/cmd/api/internal/logic/mq"
	"go-go-zero/service/atcl/cmd/api/internal/types"
	"strconv"

	"github.com/zeromicro/go-zero/core/logx"
	"go-go-zero/service/atcl/cmd/api/internal/svc"
)

type LikeLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewLikeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LikeLogic {
	return &LikeLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LikeLogic) Like(req *types.IdReq) error {
	var (
		userId    int64 = 408301323265285
		userIdStr       = "408301323265285"
	)
	key := utils.AtclLike + strconv.FormatInt(req.Id, 10)
	added := l.svcCtx.Redis.SAdd(l.ctx, key, userIdStr).Val()
	if added == 0 {
		return errors.New("不能重复点赞哦😊")
	}
	mqLogic := mq.NewRabbitMQLogic(l.ctx, l.svcCtx)
	go mq.LikeCheckUpdate(key, userId, mqLogic)
	return nil
}
