package main

import (
	"bytes"
	"log"
	"rabbitmq/common"
	"time"
	"utils"
)

func main() {
	conn := common.ConnectToMQ()
	defer conn.Close()
	channel := common.OpenChannel(conn)
	defer channel.Close()
	queue := common.DeclareQueue(channel, "task_queue", true)

	err := channel.Qos(1, 0, false)  // 设置prefetch_count之后，如果消费者没有对该消息进行ack，则不会在想改消费者投递信息
	utils.FailOnError(err, "Failed to set Qos")

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
			log.Printf("Received a message: %s ", d.Body)
			dotCount := bytes.Count(d.Body, []byte("."))
			t := time.Duration(dotCount)
			time.Sleep(t * time.Second)
			log.Printf("Done")
			d.Ack(false)   // 开启acknowledge
		}
	}()

	log.Printf(" [*] Waiting for message. To exit press CTRL+C")
	<-forever
}