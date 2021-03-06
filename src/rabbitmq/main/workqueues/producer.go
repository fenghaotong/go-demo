package main

import (
	"github.com/streadway/amqp"
	"log"
	"os"
	"rabbitmq/common"
	"utils"
)

func main() {
	mq := new(common.MqResource)
	mq.ConnectToMQ()
	channel := mq.OpenChannel()
	defer mq.CloseResource()
	queue, err := channel.QueueDeclare(
		"task_queue",
		true,     // 队列持久化
		false,
		false,
		false,
		nil)
	utils.FailOnError(err, "Failed to declare a queue")

	body := common.BodyForm(os.Args)
	err = channel.Publish(
		"",
		queue.Name,
		false,
		false,
		amqp.Publishing{
			DeliveryMode: amqp.Persistent,
			ContentType: "text/plain",
			Body: []byte(body),
		})
	utils.FailOnError(err, "Failed to publish to message")

	log.Printf(" [x] produce %s", body)
}