package main

import (
	"log"
	"rabbitmq/common"
	"utils"
)



func main() {
	conn := common.ConnectToMQ()
	defer conn.Close()
	channel := common.OpenChannel(conn)
	defer channel.Close()
	queue := common.DeclareQueue(channel, "hello", false)

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