package common

import (
	"bytes"
	"github.com/streadway/amqp"
	"log"
	"sync"
	"time"
	"utils"
)

const MQURL = "amqp://guest:guest@1.117.115.165:5672/"

var (
	rabbitmq *RabbitMQ
	once sync.Once
)

type RabbitMQ struct {
	conn *amqp.Connection
	channel *amqp.Channel
	QueueName string
	ExchangeName string
	ExchangeType string
	Key string
	Mqyrl string
}

// 单例模式创建RabbitMQ实例
func NewRabbitMQ(queueName string, exchangeName string, exchangeType string, key string) *RabbitMQ {
	once.Do(func() {
		log.Printf("--------------init rabbitmq connection and channel-------------------")
		rabbitmq = &RabbitMQ{QueueName: queueName, ExchangeName: exchangeName, ExchangeType: exchangeType, Key: key, Mqyrl: MQURL}
		var err error
		rabbitmq.conn, err = amqp.Dial(rabbitmq.Mqyrl)
		utils.FailOnError(err, "Failed to connect to RabbitMQ")
		rabbitmq.channel, err = rabbitmq.conn.Channel()
		utils.FailOnError(err, "Failed to open a channel")
	})
	return rabbitmq
}

// 关闭channel和connect
func (r *RabbitMQ) Destory()  {
	r.channel.Close()
	r.conn.Close()
}

func NewRabbitMQSimple(queueName string) *RabbitMQ  {
	return NewRabbitMQ(queueName, "", "", "")
}

func (r *RabbitMQ) PublishSimple(message string)  {
	_, err := r.channel.QueueDeclare(
		r.QueueName,
		false,
		false,
		false,
		false,
		nil)
	utils.FailOnError(err, "Failed to decalre a queue")

	r.channel.Publish(
		r.ExchangeName,
		r.QueueName,
		false,
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body: []byte(message),
		})
	log.Printf(" [x] Sent %s", message)
}

func (r *RabbitMQ) ConsumeSimple()  {
	err := r.channel.Qos(1, 0, false)
	utils.FailOnError(err, "Failed to set Qos")

	_, err = r.channel.QueueDeclare(
		r.QueueName,
		false,
		false,
		false,
		false,
		nil)
	utils.FailOnError(err, "Failed to decalre a queue")

	msgs, err := r.channel.Consume(
		r.QueueName,
		"",
		false,
		false,
		false,
		false,
		nil)
	forever := make(chan bool)

	go func() {
		for d := range msgs {
			log.Printf("Received a message: %s ", d.Body)
			dotCount := bytes.Count(d.Body, []byte("."))
			t := time.Duration(dotCount) + 2
			time.Sleep(t * time.Second)
			log.Printf("Done")
			d.Ack(false)   // 开启acknowledge
		}
	}()

	log.Printf(" [*] Waiting for message. To exit press CTRL+C")
	<-forever
}


