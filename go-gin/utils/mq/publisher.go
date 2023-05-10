package mq

import (
	"github.com/streadway/amqp"
	"log"
)

func CcPublisher() *RabbitMQ {
	r := NewRabbitMQ(CommodityCollectQueue, CommodityCollectExchange, "cc")
	// 获取connection
	var err error
	r.conn, err = amqp.Dial(r.Mqurl)
	r.failOnErr(err, "failed to connect rabbitmq!")
	// 获取channel
	r.Channel, err = r.conn.Channel()
	r.failOnErr(err, "failed to open a channel")

	// 延迟队列配置
	delaySeconds := 1000
	exchangeName := r.Exchange
	queueName := r.QueueName
	key := r.Key
	// 声明ttl队列的交换机
	err = r.Channel.ExchangeDeclare(exchangeName, "direct", true, false, false, false, nil)
	if err != nil {
		log.Println("[RABBITMQ ERROR] ExchangeDeclare error", err.Error())
		return nil
	}
	args := amqp.Table{
		"x-dead-letter-exchange":    CommodityCollectDeadExchange,
		"x-dead-letter-routing-key": "cc",
		"x-message-ttl":             int32(delaySeconds * 30),
	}
	// 声明带有ttl的队列
	_, err = r.Channel.QueueDeclare(queueName, true, false, false, false, args)
	if err != nil {
		log.Println("[RABBITMQ ERROR] QueueDeclare error", err.Error())

		return nil
	}
	err = r.Channel.QueueBind(queueName, key, exchangeName, false, nil)
	if err != nil {
		log.Println("[RABBITMQ ERROR] QueueBinding error", err.Error())
		return nil
	}
	return r
}
