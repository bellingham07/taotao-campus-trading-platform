package utils

import (
	"github.com/streadway/amqp"
	"time"
)

var MQUrl string

type (
	// RabbitMQConf rabbitMQ配置
	RabbitMQConf struct {
		RmqUrl string
	}

	RabbitmqCore struct {
		Conn    *amqp.Connection
		Channel *amqp.Channel
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

// NewRabbitMQ 创建结构体实例
func NewRabbitMQ(queueName string, exchange string, key string, conn *amqp.Connection, channel *amqp.Channel) *RabbitMQ {
	return &RabbitMQ{
		Conn:      conn,
		Channel:   channel,
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

const (
	// 商品收藏相关
	CmdtyCollectExchange     = "taotao_cmdty_collect_exchange"
	CmdtyCollectDeadExchange = "taotao_cmdty_collect_exchange_dead"
	CmdtyCollectQueue        = "taotao_cmdty_collect"
	CmdtyCollectDeadQueue    = "taotao_cmdty_collect_dead"

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
	AtclCollectExchange     = "taotao_atcl_collect_exchange"
	AtclCollectDeadExchange = "taotao_atcl_collect_exchange_dead"
	AtclCollectQueue        = "taotao_atcl_collect"
	AtclCollectDeadQueue    = "taotao_atcl_collect_dead"
	AtclLikeExchange        = "taotao_atcl_like_exchange"
	AtclLikeDeadExchange    = "taotao_atcl_like_exchange_dead"
	AtclLikeQueue           = "taotao_atcl_like"
	AtclLikeDeadQueue       = "taotao_atcl_like_dead"
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
	RedisKey string
	Time     time.Time
	UserId   int64
}
