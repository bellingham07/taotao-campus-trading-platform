package main

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/streadway/amqp"
)

type RabbitMQ struct {
	conn *amqp.Connection
	ch   *amqp.Channel
}

type LikeMessage struct {
	UserID  int
	PostID  string
	Created time.Time
}

func NewRabbitMQ() RabbitMQ {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		panic(err)
	}
	ch, err := conn.Channel()
	if err != nil {
		panic(err)
	}
	return RabbitMQ{
		conn: conn,
		ch:   ch,
	}
}

func (r RabbitMQ) SendLikeMessage(userID int, postID string) error {
	message := LikeMessage{
		UserID:  userID,
		PostID:  postID,
		Created: time.Now(),
	}
	body, err := json.Marshal(message)
	if err != nil {
		return err
	}
	// 延迟队列配置
	delaySeconds := 30
	exchangeName := "like_exchange"
	queueName := "like_queue"
	err = r.ch.ExchangeDeclare(exchangeName, "direct", true, false, false, false, nil)
	if err != nil {
		return err
	}
	args := amqp.Table{
		"x-dead-letter-exchange":    "my-dead-exchange",
		"x-dead-letter-routing-key": "dx",
		"x-message-ttl":             int32(delaySeconds * 1000),
	}
	_, err = r.ch.QueueDeclare(queueName, true, false, false, false, args)
	if err != nil {
		return err
	}
	err = r.ch.QueueBind(queueName, "", exchangeName, false, nil)
	if err != nil {
		return err
	}
	err = r.ch.Publish(exchangeName, "", false, false, amqp.Publishing{DeliveryMode: amqp.Persistent, ContentType: "application/json", Body: body})
	if err != nil {
		return err
	}
	return nil
}

func (r RabbitMQ) CancelLikeMessage(userID int, postID string) error {
	message := LikeMessage{UserID: userID, PostID: postID, Created: time.Now()}
	body, err := json.Marshal(message)
	if err != nil {
		return err
	}
	// 直接交换机配置
	exchangeName := "like_exchange"
	err = r.ch.ExchangeDeclare(exchangeName, "direct", true, false, false, false, nil)
	if err != nil {
		return err
	}
	err = r.ch.Publish(exchangeName, "", false, false, amqp.Publishing{
		DeliveryMode: amqp.Persistent,
		ContentType:  "application/json",
		Body:         body,
	},
	)
	if err != nil {
		return err
	}
	return nil
}

func (r RabbitMQ) ConsumeLikeMessage(redisClient RedisClient) {
	exchangeName := "like_exchange"
	queueName := "like_queue"
	deadLetterExchangeName := "my-dead-exchange"
	deadLetterQueueName := "my-dead-queue"
	err := r.ch.ExchangeDeclare(exchangeName, "direct", true, false, false, false, nil)
	if err != nil {
		panic(err)
	}
	args := amqp.Table{
		"x-dead-letter-exchange":    deadLetterExchangeName,
		"x-dead-letter-routing-key": "dx",
	}
	_, err = r.ch.QueueDeclare(queueName, true, false, false, false, args)
	if err != nil {
		panic(err)
	}

	err = r.ch.QueueBind(queueName, "", exchangeName, false, nil)
	if err != nil {
		panic(err)
	}

	err = r.ch.ExchangeDeclare(deadLetterExchangeName, "direct", true, false, false, false, nil)
	if err != nil {
		panic(err)
	}

	_, err = r.ch.QueueDeclare(deadLetterQueueName, true, false, false, false, nil)
	if err != nil {
		panic(err)
	}

	err = r.ch.QueueBind(deadLetterQueueName, "dx", deadLetterExchangeName, false, nil)
	if err != nil {
		panic(err)
	}

	msgs, err := r.ch.Consume(queueName, "", false, false, false, false, nil)
	if err != nil {
		panic(err)
	}

	for msg := range msgs {
		var likeMessage LikeMessage
		err := json.Unmarshal(msg.Body, &likeMessage)
		if err != nil {
			fmt.Printf("Failed to unmarshal message: %v\n", err)
			msg.Nack(false, false)
			continue
		}

		isLike, err := redisClient.IsLike(likeMessage.UserID, likeMessage.PostID)
		if err != nil {
			fmt.Printf("Failed to check if user %d liked post %s: %v\n", likeMessage.UserID, likeMessage.PostID, err)
			msg.Nack(false, false)
			continue
		}

		if !isLike {
			fmt.Printf("User %d canceled like on post %s, discard like message\n", likeMessage.UserID, likeMessage.PostID)
			msg.Ack(false)
			continue
		}

		err = UpdateLikeCount(likeMessage.UserID, likeMessage.PostID)
		if err != nil {
			fmt.Printf("Failed to update like count for post %s: %v\n", likeMessage.PostID, err)
			msg.Nack(false, false)
			continue
		}

		fmt.Printf("Like count updated for post %s\n", likeMessage.PostID)
		msg.Ack(false)
	}
}
func UpdateLikeCount(userID int, postID string) error {
	// 更新点赞数量
	return nil
}
