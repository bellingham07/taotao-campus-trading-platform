package mq

import (
	"github.com/streadway/amqp"
	"github.com/zeromicro/go-zero/core/logx"
	"go-go-zero/common/utils"
	"time"
)

func CollectUpdatePublisher(redisKey string, member int64, isCollect bool, mqLogic *RabbitMQLogic) {
	ticker := time.NewTicker(time.Second * 30)
	msg := &utils.CcMessage{
		RedisKey:  redisKey,
		UserId:    member,
		Time:      time.Now().Local(),
		IsCollect: isCollect,
	}
	body, _ := Json.Marshal(msg)
	publisher := utils.NewRabbitMQ(utils.CmdtyCollectQueue, utils.CmdtyCollectExchange, "cc", mqLogic.svcCtx.RmqCore.Conn, mqLogic.svcCtx.RmqCore.Channel)
	err := publisher.Channel.Publish(publisher.Exchange, publisher.Key, false, false,
		amqp.Publishing{
			DeliveryMode: amqp.Persistent,
			ContentType:  "application/json",
			Body:         body,
		})
	if err != nil {
		logx.Infof("[RABBITMQ ERROR] CollectUpdatePublisher 发送收藏消息失败 %v\n", err)
		select {
		case <-ticker.C:
			if isCollect {
				go mqLogic.CollectCheck(msg)
			} else {
				go mqLogic.UncollectCheck(msg)
			}
		}
	}
}

func LikeCheckUpdate(redisKey string, member int64, mqLogic *RabbitMQLogic) {
	ticker := time.NewTicker(time.Second * 30)
	msg := &utils.LMessage{
		RedisKey: redisKey,
		Time:     time.Now().Local(),
		UserId:   member,
	}
	body, _ := Json.Marshal(msg)
	publisher := utils.NewRabbitMQ(utils.AtclLikeQueue, utils.AtclLikeExchange, "cl", mqLogic.svcCtx.RmqCore.Conn, mqLogic.svcCtx.RmqCore.Channel)
	err := publisher.Channel.Publish(publisher.Exchange, publisher.Key, false, false,
		amqp.Publishing{
			DeliveryMode: amqp.Persistent,
			ContentType:  "application/json",
			Body:         body,
		})
	if err != nil {
		logx.Infof("[RABBITMQ ERROR] LikeCheckUpdate 发送收藏消息失败 %v\n", err)
		select {
		case <-ticker.C:
			go mqLogic.LikeCheckUpdate(msg)
		}
	}
}
