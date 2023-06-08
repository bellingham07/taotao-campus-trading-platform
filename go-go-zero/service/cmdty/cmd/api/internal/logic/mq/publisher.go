package mq

import (
	"github.com/streadway/amqp"
	"github.com/zeromicro/go-zero/core/logx"
	"go-go-zero/common/utils"
	"log"
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
	publisher := CcPublisher()
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
		amqp.Publishing{DeliveryMode: amqp.Persistent,
			ContentType: "application/json",
			Body:        body,
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

func CcPublisher() *utils.RabbitMQ {
	// 获取connection
	r := utils.DialRabbitMq(utils.CmdtyCollectQueue, utils.CmdtyCollectExchange, "cc")
	if r == nil {
		return nil
	}
	// 延迟队列配置
	delaySeconds := 1000
	exchangeName := r.Exchange
	queueName := r.QueueName
	key := r.Key
	// 声明ttl队列的交换机
	err := r.Channel.ExchangeDeclare(exchangeName, "direct", true, false, false, false, nil)
	if err != nil {
		log.Println("[RABBITMQ ERROR] ExchangeDeclare error : ", err.Error())
		return nil
	}
	args := amqp.Table{
		"x-dead-letter-exchange":    utils.CmdtyCollectDeadExchange,
		"x-dead-letter-routing-key": "cc",
		"x-message-ttl":             int32(delaySeconds * 30),
	}
	// 声明带有ttl的队列
	_, err = r.Channel.QueueDeclare(queueName, true, false, false, false, args)
	if err != nil {
		logx.Infof("[RABBITMQ ERROR] QueueDeclare error : ", err.Error())

		return nil
	}
	err = r.Channel.QueueBind(queueName, key, exchangeName, false, nil)
	if err != nil {
		logx.Infof("[RABBITMQ ERROR] QueueBinding error : ", err.Error())
		return nil
	}
	return r
}
