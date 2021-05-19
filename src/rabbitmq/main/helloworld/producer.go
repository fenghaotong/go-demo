package main

import (
	"github.com/streadway/amqp"
	"rabbitmq/common"
	"utils"
)

func main() {
	mq := new(common.MqResource)
	mq.ConnectToMQ()
	channel := mq.OpenChannel()
	defer mq.CloseResource()

	queue, err := channel.QueueDeclare(
		"hello",
		false,     // 队列持久化
		false,
		false,
		false,
		nil)
	utils.FailOnError(err, "Failed to declare a queue")

	body := "hello world"
	err = channel.Publish(
		"",
		queue.Name,
		false,
		false,
		amqp.Publishing{
			ContentType: "test/plain",
			Body: []byte(body),
		})
	utils.FailOnError(err, "Failed to publish a message")
}