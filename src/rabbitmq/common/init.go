package common

import (
	"github.com/streadway/amqp"
	"utils"
)

type MqResource struct {
	conn *amqp.Connection
	channel *amqp.Channel
}


func (self *MqResource) ConnectToMQ() {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	utils.FailOnError(err, "Failed to connect to RabbitMQ")
	self.conn = conn
}

func (self *MqResource) OpenChannel() *amqp.Channel {
	ch, err := self.conn.Channel()
	utils.FailOnError(err, "Failed to open a channel")
	self.channel = ch
	return self.channel
}

func (self *MqResource) CloseResource() {
	self.conn.Close()
	self.channel.Close()
}