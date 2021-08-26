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

	exchangeName := "cancel_fanout"
	exchangeType := "fanout"
	err := channel.ExchangeDeclare(
		exchangeName,
		exchangeType,
		true,
		false,
		false,
		false,
		nil)
	utils.FailOnError(err, "Failed to declare a Exchange")

	body := common.BodyForm(os.Args)
	err = channel.Publish(
		exchangeName,
		"",
		false,
		false,
		amqp.Publishing{
			ContentType: "text/palin",
			Body: []byte(body),
		})
	utils.FailOnError(err, "Failed to publish to message")

	log.Printf(" [x] Sent %s", body)
}