package mqLogic

import (
	"encoding/json"
	"github.com/streadway/amqp"
	"log"
)

func initConsumers() {
	go CcConsumer()
	go VConsumer()
	go LConsumer()

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
		err = json.Unmarshal(msg.Body, ccMessage)
		if err != nil {
			log.Printf("[RABBITMQ COMMODITYCOLLECT CONSUMER FAIL] Failed to unmarshal message: %v\n", err)
			msg.Nack(false, false)
			continue
		}
		if ccMessage.IsCollect {
			CollectCheckUpdate(ccMessage)
		} else {
			CollectCheckDelete(ccMessage)
		}
		msg.Ack(false)
	}
	<-forever
}

func VConsumer() {
	r := NewRabbitMQ(ViewDeadQueue, ViewDeadExchange, "v")

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
	msgs, err := r.Channel.Consume(ViewDeadQueue, "", false, false, false, false, nil)
	if err != nil {
		panic(err)
	}
	forever := make(chan int, 0)
	for msg := range msgs {
		log.Println("接受成功咕咕咕咕咕咕过过过过过过过过过过过")
		vMessage := new(VMessage)
		err = json.Unmarshal(msg.Body, vMessage)
		if err != nil {
			log.Printf("[RABBITMQ VIEWCOUNT CONSUMER FAIL] Failed to unmarshal message: %v\n", err)
			msg.Nack(false, false)
			continue
		}
		ViewCheckUpdate(vMessage)
		msg.Ack(false)
	}
	<-forever
}

func LConsumer() {
	r := NewRabbitMQ(LikeDeadQueue, LikeDeadExchange, "l")

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
	msgs, err := r.Channel.Consume(LikeDeadQueue, "", false, false, false, false, nil)
	if err != nil {
		panic(err)
	}
	forever := make(chan int, 0)
	for msg := range msgs {
		log.Println("接受成功咕咕咕咕咕咕过过过过过过过过过过过")
		lMessage := new(LMessage)
		err = json.Unmarshal(msg.Body, lMessage)
		if err != nil {
			log.Printf("[RABBITMQ COMMODITYCOLLECT CONSUMER FAIL] Failed to unmarshal message: %v\n", err)
			msg.Nack(false, false)
			continue
		}
		LikeCheckUpdate(lMessage)
		msg.Ack(false)
	}
	<-forever
}
