package main

import (
	"log"
	"rabbitmq/common"
	"utils"
)

func main() {
	mq := new(common.MqResource)

	mq.ConnectToMQ()
	channel := mq.OpenChannel()
	defer mq.CloseResource()

	queue, err := channel.QueueDeclare(
		"",
		false,
		false,
		true,
		false,
		nil)
	utils.FailOnError(err, "Failed to declare queue")

	err = channel.QueueBind(
		queue.Name,
		"",
		"logs",
		false,
		nil)
	utils.FailOnError(err, "Failed to bind queue")

	msgs, err := channel.Consume(
		queue.Name,
		"",
		false,
		false,
		false,
		false,
		nil)
	utils.FailOnError(err, "Failed to register a consumer")

	forever := make(chan bool)

	go func() {
		for d := range msgs {
			log.Printf(" [x] %s", d.Body)
		}
	}()

	log.Printf(" [*] Waiting for message. To exit press CTRL+C")
	<-forever
}