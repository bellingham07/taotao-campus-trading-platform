package mq

import (
	"com.xpdj/go-gin/config"
	"fmt"
	"github.com/streadway/amqp"
	"log"
)

var MQURL string

func InitRabbitMQ(config *config.RabbitMQConfig) {
	url := config.Url
	username := config.Username
	password := config.Password
	MQURL = "amqp://" + username + ":" + password + "@" + url
}

// RabbitMQ rabbitMQ结构体
type RabbitMQ struct {
	conn      *amqp.Connection
	channel   *amqp.Channel
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
	rabbitmq.channel, err = rabbitmq.conn.Channel()
	rabbitmq.failOnErr(err, "failed to open a channel")
	return rabbitmq
}

// PublishTopic 话题模式发送消息
func (r *RabbitMQ) PublishTopic(message string) error {
	// 发送消息
	err := r.channel.Publish(r.Exchange, r.Key, false, false, amqp.Publishing{
		ContentType: "text/plain",
		Body:        []byte(message),
	})
	if err != nil {
		return err
	}
	return nil
}

// ReceiveTopic 话题模式接受消息
// 要注意key,规则
// 其中“*”用于匹配一个单词，“#”用于匹配多个单词（可以是零个）
// 匹配 kuteng.* 表示匹配 kuteng.hello, kuteng.hello.one需要用kuteng.#才能匹配到
func (r *RabbitMQ) ReceiveTopic(queueName string) error {
	// 绑定队列到 exchange 中
	// 在pub/sub模式下，这里的key要为空
	err := r.channel.QueueBind(queueName, r.Key, r.Exchange, false, nil)
	if err != nil {
		return err
	}
	// 消费消息
	messges, err := r.channel.Consume(queueName, "", true, false, false, false, nil)
	forever := make(chan bool)
	go func() {
		for d := range messges {
			log.Printf("Received a message: %s", d.Body)
		}
	}()
	fmt.Println("退出请按 CTRL+C\n")
	<-forever
}

// 错误处理函数
func (r *RabbitMQ) failOnErr(err error, message string) {
	if err != nil {
		log.Fatalf("%s:%s", message, err)
		panic(fmt.Sprintf("%s:%s", message, err))
	}
}
