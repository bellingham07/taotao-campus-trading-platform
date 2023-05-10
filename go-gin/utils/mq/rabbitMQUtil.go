package mq

import (
	"com.xpdj/go-gin/config"
	"fmt"
	"github.com/streadway/amqp"
	"log"
	"time"
)

var MQURL string

func InitRabbitMQ(config *config.RabbitMQConfig) {
	url := config.Url
	username := config.Username
	password := config.Password
	MQURL = "amqp://" + username + ":" + password + "@" + url
	initConsumers()
}

// RabbitMQ rabbitMQ结构体
type RabbitMQ struct {
	conn      *amqp.Connection
	Channel   *amqp.Channel
	QueueName string // 队列名称
	Exchange  string // 交换机名称
	Key       string // bind Key 名称
	Mqurl     string // 连接信息
}

// NewRabbitMQ 创建结构体实例
func NewRabbitMQ(queueName string, exchange string, key string) *RabbitMQ {
	return &RabbitMQ{QueueName: queueName, Exchange: exchange, Key: key, Mqurl: MQURL}
}

// NewRabbitMQTopic 创建RabbitMQ实例
func NewRabbitMQTopic(queueName string, exchangeName string, routingKey string) *RabbitMQ {
	// 创建RabbitMQ实例
	rabbitmq := NewRabbitMQ(queueName, exchangeName, routingKey)
	var err error
	// 获取connection
	rabbitmq.conn, err = amqp.Dial(rabbitmq.Mqurl)
	rabbitmq.failOnErr(err, "failed to connect rabbitmq!")
	// 获取channel
	rabbitmq.Channel, err = rabbitmq.conn.Channel()
	rabbitmq.failOnErr(err, "failed to open a channel")
	return rabbitmq
}

// PublishTopic 话题模式发送消息
func (r *RabbitMQ) PublishTopic(message string) error {
	// 发送消息
	err := r.Channel.Publish(r.Exchange, r.Key, false, false, amqp.Publishing{
		ContentType: "text/plain",
		Body:        []byte(message),
	})
	if err != nil {
		return err
	}
	return nil
}

// 错误处理函数
func (r *RabbitMQ) failOnErr(err error, message string) {
	if err != nil {
		log.Fatalf("%s:%s", message, err)
		panic(fmt.Sprintf("%s:%s", message, err))
	}
}

const (
	CommodityCollectExchange     = "taotao_commodity_collect_exchange"
	CommodityCollectDeadExchange = "taotao_commodity_collect_exchange_dead"
	CommodityCollectQueue        = "taotao_commodity_collect"
	CommodityCollectDeadQueue    = "taotao_commodity_collect_dead"

	UserCollectQueue = "taotao_delay_user_collect"

	LikeQueue = "taotao_delay_like"
)

type CcMessage struct {
	RedisKey  string
	Time      time.Time
	UserId    string
	IsCollect bool
}
