package mq

import (
	"fmt"
	"log"
)

// ReceiveTopic 话题模式接受消息
// 要注意key,规则
// 其中“*”用于匹配一个单词，“#”用于匹配多个单词（可以是零个）
// 匹配 kuteng.* 表示匹配 kuteng.hello, kuteng.hello.one需要用kuteng.#才能匹配到
func (r *RabbitMQ) DelayCollect(queueName string) error {
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
