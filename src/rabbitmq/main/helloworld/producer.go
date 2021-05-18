package main

import (
	"github.com/streadway/amqp"
	"rabbitmq/common"
	"utils"
)

func main() {
	conn := common.ConnectToMQ()
	defer conn.Close()
	channel := common.OpenChannel(conn)
	defer channel.Close()
	queue := common.DeclareQueue(channel, "hello", false)

	body := "hello world"
	err := channel.Publish(
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