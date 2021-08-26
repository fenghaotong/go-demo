package main

import (
	"fmt"
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

	exchangeName := "logs_topic"
	exchangeType := "topic"
	err := channel.ExchangeDeclare(
		exchangeName,
		exchangeType,
		true,
		false,
		false,
		false,
		nil)
	utils.FailOnError(err, "Failed to declare an exchange")

	body := common.BodyForm3(os.Args)
	fmt.Println(body)
	fmt.Println(common.SeverityFrom2(os.Args))
	err = channel.Publish(
		exchangeName,
		common.SeverityFrom2(os.Args),
		false,
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body: []byte(body),
		})
	utils.FailOnError(err, "Failed to publish a message")

	log.Printf(" [x] Sent %s", body)
}
