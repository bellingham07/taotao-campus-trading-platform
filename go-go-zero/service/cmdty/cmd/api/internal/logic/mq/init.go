package mq

import (
	"context"
	"github.com/streadway/amqp"
	"go-go-zero/common/utils"
	"go-go-zero/service/cmdty/cmd/api/internal/svc"
)

func InitRabbitMQ(svcCtx *svc.ServiceContext) {
	RabbitMQ = NewRabbitMQLogic(context.Background(), svcCtx)

	go InitCcPublisher(svcCtx.RmqCore)

	go StartCcConsumer(svcCtx.RmqCore.Conn)
}

func InitCcPublisher(core *utils.RabbitmqCore) {
	// 获取connection
	r := utils.NewRabbitMQ(utils.CmdtyCollectQueue, utils.CmdtyCollectExchange, "cc", core.Conn, core.Channel)
	if r == nil {
		panic("[RABBITMQ ERROR] InitCcPublisher 初始化 cmdty collect publisher 错误！")
	}
	// 延迟队列配置
	delaySeconds := 1000
	exchangeName := r.Exchange
	queueName := r.QueueName
	key := r.Key
	// 声明ttl队列的交换机
	err := r.Channel.ExchangeDeclare(exchangeName, "direct", true, false, false, false, nil)
	if err != nil {
		panic("[RABBITMQ ERROR] InitCcPublisher ExchangeDeclare 错误 : " + err.Error())
		return
	}
	args := amqp.Table{
		"x-dead-letter-exchange":    utils.CmdtyCollectDeadExchange,
		"x-dead-letter-routing-key": "cc",
		"x-message-ttl":             int32(delaySeconds * 30),
	}
	// 声明带有ttl的队列
	_, err = r.Channel.QueueDeclare(queueName, true, false, false, false, args)
	if err != nil {
		panic("[RABBITMQ ERROR] QueueDeclare error : " + err.Error())
		return
	}
	err = r.Channel.QueueBind(queueName, key, exchangeName, false, nil)
	if err != nil {
		panic("[RABBITMQ ERROR] QueueBinding error : " + err.Error())
		return
	}
}
