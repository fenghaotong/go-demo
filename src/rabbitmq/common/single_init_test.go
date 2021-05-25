package common

import (
	"strconv"
	"sync"
	"testing"
	"time"
)

func TestRabbitMQ_PublishSimple(t *testing.T) {
	var lock sync.Mutex
	simple := NewRabbitMQSimple("test")
	forever := make(chan bool)
	var count int

	go func() {
		for {
			lock.Lock()
			count += 1
			lock.Unlock()
			simple.PublishSimple("协程1：" + strconv.Itoa(count))
			time.Sleep(1*time.Second)
		}
	}()

	go func() {
		for {
			lock.Lock()
			count += 1
			lock.Unlock()
			simple.PublishSimple("协程2：" + strconv.Itoa(count))
			time.Sleep(1*time.Second)
		}
	}()
	<-forever

}

func TestRabbitMQ_ConsumeSimple(t *testing.T) {
	simple := NewRabbitMQSimple("test")
	simple.ConsumeSimple()
}