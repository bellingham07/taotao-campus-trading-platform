package mq

import (
	"github.com/streadway/amqp"
	"github.com/zeromicro/go-zero/core/logx"
	"go-go-zero/common/utils"
	"time"
)

func CollectUpdatePublisher(redisKey string, member int64, isCollect bool, mqLogic *RabbitMQLogic) {
	now := time.Now()
	ticker := time.NewTicker(time.Second * 30)
	message := &utils.CcMessage{
		RedisKey:  redisKey,
		UserId:    member,
		Time:      now,
		IsCollect: isCollect,
	}
	body, _ := Json.Marshal(message)
	publisher := newCcPublisher(mqLogic)
	// 假如无法使用mq
	if publisher == nil {
		select {
		case <-ticker.C:
			if isCollect {
				go mqLogic.CollectCheck(message)
			} else {
				go mqLogic.UncollectCheck(message)
			}
			return
		}
	}
	err := publisher.Channel.Publish(publisher.Exchange, publisher.Key, false, false,
		amqp.Publishing{
			DeliveryMode: amqp.Persistent,
			ContentType:  "application/json",
			Body:         body,
		})
	if err != nil {
		logx.Infof("[RABBITMQ ERROR] : %v\n", err.Error())
		select {
		case <-ticker.C:
			if isCollect {
				go mqLogic.CollectCheck(message)
			} else {
				go mqLogic.UncollectCheck(message)
			}
			return
		}
	}
}

func newCcPublisher(mqLogic *RabbitMQLogic) *utils.RabbitMQ {
	return utils.NewRabbitMQ(utils.CmdtyCollectQueue, utils.CmdtyCollectExchange, "cc", mqLogic.svcCtx.RmqCore.Conn, mqLogic.svcCtx.RmqCore.Channel)
}
