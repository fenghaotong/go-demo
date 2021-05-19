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
	utils.FailOnError(err, "Failed to declare exchange")

	body := common.BodyForm(os.Args)
	err = channel.Publish(
		exchangeName,
		common.SeverityFrom(os.Args),
		false,
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body: []byte(body),
		})

	utils.FailOnError(err, "Failed to publish a message")

	log.Printf(" [x] send %s ", body)

}