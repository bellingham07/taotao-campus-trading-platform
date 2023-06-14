package mq

import (
	"context"
	"github.com/json-iterator/go"
	"github.com/zeromicro/go-zero/core/logx"
	"go-go-zero/common/utils"
	"go-go-zero/service/cmdty/cmd/api/internal/svc"
	"go-go-zero/service/cmdty/model"
	"strconv"
	"strings"
)

type RabbitMQLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRabbitMQLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RabbitMQLogic {
	return &RabbitMQLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

var (
	RabbitMQ *RabbitMQLogic
	Json     = jsoniter.ConfigCompatibleWithStandardLibrary
)

func (l *RabbitMQLogic) CollectCheck(ccMessage *utils.CcMessage) {
	redisKey := ccMessage.RedisKey
	userId := ccMessage.UserId
	isMember, _ := l.svcCtx.Redis.SIsMember(l.ctx, redisKey, userId).Result()
	if isMember {
		commodityId := getIdByRedisKey(redisKey)
		collect := &model.CmdtyCollect{
			CmdtyId:  commodityId,
			UserId:   userId,
			CreateAt: ccMessage.Time,
			Status:   1,
		}
		_, _ = model.CmdtyCollectModel.Insert(l.svcCtx.CmdtyCollect, l.ctx, collect)
	}
}

func (l *RabbitMQLogic) UncollectCheck(ccMessage *utils.CcMessage) {
	redisKey := ccMessage.RedisKey
	userId := ccMessage.UserId
	isMember, _ := l.svcCtx.Redis.SIsMember(l.ctx, redisKey, strconv.FormatInt(userId, 10)).Result()
	if !isMember {
		cmdtyId := getIdByRedisKey(redisKey)
		_ = l.svcCtx.CmdtyCollect.DeleteByCmdtyIdAndUserId(cmdtyId, userId)
	}
}

func getIdByRedisKey(redisKey string) int64 {
	split := strings.LastIndex(redisKey, ":")
	idStr := redisKey[split+1:]
	id, _ := strconv.ParseInt(idStr, 10, 64)
	return id
}
