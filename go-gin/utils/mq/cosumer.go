package mq

import (
	"com.xpdj/go-gin/logic/mq"
	"encoding/json"
	"github.com/streadway/amqp"
	"log"
)

func initConsumers() {
	CcConsumer()
}

func CcConsumer() {
	r := NewRabbitMQ(CommodityCollectDeadQueue, CommodityCollectDeadExchange, "cc")

	// 获取connection
	var err error
	r.conn, err = amqp.Dial(r.Mqurl)
	r.failOnErr(err, "failed to connect rabbitmq!")
	// 获取channel
	r.Channel, err = r.conn.Channel()
	r.failOnErr(err, "failed to open a channel")

	exchangeName := r.Exchange
	queueName := r.QueueName
	key := r.Key
	// 声明死信交换机
	err = r.Channel.ExchangeDeclare(exchangeName, "direct", true, false, false, false, nil)
	if err != nil {
		panic(err)
	}
	// 声明有死信队列
	_, err = r.Channel.QueueDeclare(queueName, true, false, false, false, nil)
	if err != nil {
		panic(err)
	}
	// 将死信交换机和死信队列绑定
	err = r.Channel.QueueBind(queueName, key, exchangeName, false, nil)
	if err != nil {
		panic(err)
	}
	// 开始监听
	msgs, err := r.Channel.Consume(CommodityCollectDeadQueue, "", false, false, false, false, nil)
	if err != nil {
		panic(err)
	}
	forever := make(chan int, 0)
	for msg := range msgs {
		log.Println("接受成功咕咕咕咕咕咕过过过过过过过过过过过")
		ccMessage := new(CcMessage)
		err := json.Unmarshal(msg.Body, ccMessage)
		if err != nil {
			log.Printf("[RABBITMQ COMMODITYCOLLECT CONSUMER FAIL] Failed to unmarshal message: %v\n", err)
			msg.Nack(false, false)
			continue
		}
		if ccMessage.IsCollect {
			mqLogic.CollectCheckUpdate(ccMessage.RedisKey, ccMessage.UserId, ccMessage.Time)
		} else {
			mqLogic.CollectCheckDelete(ccMessage.RedisKey, ccMessage.UserId)
		}

		msg.Ack(false)
	}
	<-forever
}
