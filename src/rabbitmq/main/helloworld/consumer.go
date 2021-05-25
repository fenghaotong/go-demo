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
		"hello",
		false,     // 队列持久化
		false,
		false,
		false,
		nil)
	utils.FailOnError(err, "Failed to declare a queue")

	msgs, err := channel.Consume(
		queue.Name,
		"",
		true,
		false,
		false,
		false,
		nil)
	utils.FailOnError(err, "Failed t register a consumer")

	forever := make(chan bool)

	go func() {
		for d := range msgs {
			log.Printf("Received a message: %s ", d.Body)
		}
	}()

	log.Printf(" [*] Waiting for message. To exit press CTRL+C")
	<-forever
}