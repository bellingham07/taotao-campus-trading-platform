package mq

import (
	"context"
	jsoniter "github.com/json-iterator/go"
	"github.com/zeromicro/go-zero/core/logx"
	"go-go-zero/common/utils"
	"go-go-zero/service/atcl/cmd/api/internal/svc"
	"go-go-zero/service/atcl/model"
	"go-go-zero/service/user/cmd/rpc/userservice"
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

func (l *RabbitMQLogic) CollectCheck(acMessage *utils.CcMessage) {
	redisKey := acMessage.RedisKey
	userId := acMessage.UserId
	isMember, _ := l.svcCtx.Redis.SIsMember(l.ctx, redisKey, userId).Result()
	if isMember {
		atclId := getIdByRedisKey(redisKey)
		collect := &model.AtclCollect{
			AtclId:   atclId,
			UserId:   userId,
			CreateAt: acMessage.Time,
			Status:   1,
		}
		_, err := l.svcCtx.AtclCollect.Insert(collect)
		if err != nil {
			logx.Infof("[DB ERROR] CollectCheck 插入文章收藏记录失败 %v\n", err)
		}
	}
}

func (l *RabbitMQLogic) UncollectCheck(acMessage *utils.CcMessage) {
	redisKey := acMessage.RedisKey
	userId := acMessage.UserId
	isMember, _ := l.svcCtx.Redis.SIsMember(l.ctx, redisKey, strconv.FormatInt(userId, 10)).Result()
	if !isMember {
		cmdtyId := getIdByRedisKey(redisKey)
		cc := &model.AtclCollect{
			AtclId: cmdtyId,
			UserId: userId,
		}
		_, err := l.svcCtx.AtclCollect.Delete(cc)
		if err != nil {
			logx.Infof("[DB ERROR] UncollectCheck 删除文章收藏记录失败 %v\n", err)
		}
	}
}

func (l *RabbitMQLogic) LikeCheckUpdate(lMessage *utils.LMessage) {
	redisKey := lMessage.RedisKey
	userId := lMessage.UserId
	isMember, _ := l.svcCtx.Redis.SIsMember(l.ctx, redisKey, userId).Result()
	if isMember {
		atclId := getIdByRedisKey(redisKey)
		go l.IncrAtclLike(atclId)
		go l.IncrUserLike(userId)
	}
}

func (l *RabbitMQLogic) IncrAtclLike(atclId int64) {
	_, err := l.svcCtx.AtclCollect.ID(atclId).Incr("like").Update(&model.AtclCollect{})
	if err != nil {
		logx.Infof("[DB ERROR] IncrAtclLike 更新文章点赞数失败 %v\n", err)
	}
}

func (l *RabbitMQLogic) IncrUserLike(userId int64) {
	code, err := l.svcCtx.UserRpc.IncrLike(l.ctx, &userservice.IdReq{Id: userId})
	if code.GetCode() != 0 || err == nil {
		logx.Infof("[RPC ERROR] IncrUserLike 更新用户点赞数失败 %v\n", err)
	}
}

func getIdByRedisKey(redisKey string) int64 {
	split := strings.LastIndex(redisKey, ":")
	idStr := redisKey[split+1:]
	id, _ := strconv.ParseInt(idStr, 10, 64)
	return id
}
