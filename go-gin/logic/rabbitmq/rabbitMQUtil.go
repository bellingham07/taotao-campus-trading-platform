package mqLogic

import (
	"com.xpdj/go-gin/config"
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

// DialRabbitMq 创建RabbitMQ实例
func DialRabbitMq(queueName string, exchangeName string, routingKey string) *RabbitMQ {
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
		log.Fatalf(message + ":" + err.Error())
	}
}

const (
	// 商品收藏相关
	CommodityCollectExchange     = "taotao_commodity_collect_exchange"
	CommodityCollectDeadExchange = "taotao_commodity_collect_exchange_dead"
	CommodityCollectQueue        = "taotao_commodity_collect"
	CommodityCollectDeadQueue    = "taotao_commodity_collect_dead"

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
