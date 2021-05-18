package common

import (
	"github.com/streadway/amqp"
	"utils"
)

func ConnectToMQ() *amqp.Connection{
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	utils.FailOnError(err, "Failed to connect to RabbitMQ")
	return conn
}

func OpenChannel(conn *amqp.Connection) *amqp.Channel {
	ch, err := conn.Channel()
	utils.FailOnError(err, "Failed to open a channel")
	return ch
}

func DeclareQueue(ch *amqp.Channel, queueName string, durable bool) amqp.Queue {
	q, err := ch.QueueDeclare(
		queueName,
		durable,     // 队列持久化
		false,
		false,
		false,
		nil)
	utils.FailOnError(err, "Failed to declare a queue")
	return q
}

