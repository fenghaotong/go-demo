package common

import (
	"encoding/json"
	"fmt"
	"sync"
	"testing"
	"time"
)

func TestRabbitMQ_PublishSimple(t *testing.T) {
	var lock sync.Mutex
	simple := NewRabbitMQSimple("task_status_queue")
	forever := make(chan bool)
	var count int
	cancelTaskInfo := map[string]string{
		"taskid": "93659114-21f8-4b69-98e3-8ae5f66710a2",
		"status": "2",
		"host": "127.0.0.1",
		"port": "8888",
	}
	jsons, errs := json.Marshal(cancelTaskInfo)
	if errs != nil {
		fmt.Println(errs)
	}


	go func() {
		for {
			lock.Lock()
			count += 1
			lock.Unlock()
			//simple.PublishSimple("协程1：" + strconv.Itoa(count))
			simple.PublishSimple(string(jsons))
			time.Sleep(1*time.Second)
		}
	}()

	go func() {
		for {
			lock.Lock()
			count += 1
			lock.Unlock()
			//simple.PublishSimple("协程2：" + strconv.Itoa(count))
			simple.PublishSimple(string(jsons))
			time.Sleep(1*time.Second)
		}
	}()
	<-forever

}

func TestRabbitMQ_ConsumeSimple(t *testing.T) {
	simple := NewRabbitMQSimple("sdk_windows_queue")
	simple.ConsumeSimple()
}