package utils

import (
	"github.com/streadway/amqp"
	"log"
	"time"
)

var MQUrl string

type (
	// RabbitMQConf rabbitMQ配置
	RabbitMQConf struct {
		MQUrl string
	}

	// RabbitMQ rabbitMQ结构体
	RabbitMQ struct {
		Conn      *amqp.Connection
		Channel   *amqp.Channel
		QueueName string // 队列名称
		Exchange  string // 交换机名称
		Key       string // bind Key 名称
		MQUrl     string // 连接信息
	}
)

//func InitRabbitMQ(config *RabbitMQConf) {
//	MQUrl = config.MQUrl
//	initConsumers()
//}

// NewRabbitMQ 创建结构体实例
func NewRabbitMQ(queueName string, exchange string, key string, conn *amqp.Connection) *RabbitMQ {
	return &RabbitMQ{
		Conn:      conn,
		QueueName: queueName,
		Exchange:  exchange,
		Key:       key,
		MQUrl:     MQUrl,
	}
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

// FailOnErr 错误处理函数
func (r *RabbitMQ) FailOnErr(err error, message string) {
	if err != nil {
		log.Fatalf(message + ":" + err.Error())
	}
}

const (
	// 商品收藏相关
	CmdtyCollectExchange     = "taotao_commodity_collect_exchange"
	CmdtyCollectDeadExchange = "taotao_commodity_collect_exchange_dead"
	CmdtyCollectQueue        = "taotao_commodity_collect"
	CmdtyCollectDeadQueue    = "taotao_commodity_collect_dead"

	// 浏览数相关（商品和文章）
	ViewExchange     = "taotao_view_exchange"
	ViewDeadExchange = "taotao_view_exchange_dead"
	ViewQueue        = "taotao_view"
	ViewDeadQueue    = "taotao_view_dead"

	// 用户收藏相关
	UserCollectExchange     = "taotao_user_collect_exchange"
	UserCollectDeadExchange = "taotao_user_collect_exchange_dead"
	UserCollectQueue        = "taotao_user_collect"
	UserCollectDeadQueue    = "taotao_user_collect_dead"

	// 点赞相关（商品和文章）
	LikeExchange     = "taotao_like_exchange"
	LikeDeadExchange = "taotao_like_exchange_dead"
	LikeQueue        = "taotao_like"
	LikeDeadQueue    = "taotao_like_dead"
)

type CcMessage struct {
	RedisKey  string
	Time      time.Time
	UserId    int64
	IsCollect bool
}

type VMessage struct {
	RedisKey    string
	Time        time.Time
	IsCommodity bool
}

type LMessage struct {
	RedisKey  string
	Time      time.Time
	UserId    int64
	IsArticle bool
}
