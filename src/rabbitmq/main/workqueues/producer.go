package main

import (
	"github.com/streadway/amqp"
	"log"
	"os"
	"rabbitmq/common"
	"strings"
	"utils"
)

func bodyForm(args []string) string {
	var s string
	if (len(args) < 2 || os.Args[1] == ""){
		s = "hello"
	} else {
		s = strings.Join(args[1:], "")
	}
	return s
}

func main() {
	conn := common.ConnectToMQ()
	defer conn.Close()
	channel := common.OpenChannel(conn)
	defer channel.Close()
	queue := common.DeclareQueue(channel, "task_queue", true)

	body := bodyForm(os.Args)
	err := channel.Publish(
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