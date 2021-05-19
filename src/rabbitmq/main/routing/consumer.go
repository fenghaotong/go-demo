package main

import (
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

	exchangeName := "logs_direct"
	exchangeType := "direct"
	err := channel.ExchangeDeclare(
		exchangeName,
		exchangeType,
		true,
		false,
		false,
		false,
		nil)
	utils.FailOnError(err, " Failed to declare to message")

	queue, err := channel.QueueDeclare(
		"",
		false,
		false,
		true,
		false,
		nil)
	utils.FailOnError(err, "Failed to declare to queue")

	if len(os.Args) < 2 {
		log.Printf("Usage: %s [info] [waring] [error]", os.Args[0])
		os.Exit(0)
	}

	for _, s := range os.Args[1:] {
		log.Printf("Binding queue %s to exchange %s with routing key %s",
			queue.Name, exchangeName, s)

		err = channel.QueueBind(
			queue.Name,
			s,
			exchangeName,
			false,
			nil)
		utils.FailOnError(err, "Failed to bind a queue")
	}

	msgs, err := channel.Consume(
		queue.Name,
		"",
		true,
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