// @Description  TODO
// @Author  	 jiangyang  
// @Created  	 2020/11/21 2:58 下午
package rabbitmq_test

import (
	"github.com/comeonjy/util/config"
	"github.com/comeonjy/util/rabbitmq"
	"github.com/streadway/amqp"
	"log"
	"strconv"
	"testing"
	"time"
)

func init() {
	rabbitmq.Init(config.GetConfig().Rabbitmq)
}

func TestConsume(t *testing.T) {
	if err := rabbitmq.Consume("hello_task", func(msg amqp.Delivery) {
		time.Sleep(time.Second)
		log.Println(string(msg.Body))
		msg.Ack(false)
	}); err != nil {
		t.Error(err)
	}
	select {}
}

func TestProduce(t *testing.T) {
	for i := 0; i < 100; i++ {
		time.Sleep(time.Second)
		if err := rabbitmq.Produce("hello_task", []byte("ok-"+strconv.Itoa(i))); err != nil {
			t.Error(err)
		}
	}
}

func TestFanoutPublish(t *testing.T) {
	if err := rabbitmq.FanoutPublish("log",[]byte("this is log")); err != nil {
		t.Error(err)
	}
}

func TestFanoutSubscribe(t *testing.T) {
	if err := rabbitmq.FanoutSubscribe("log", "33", func(msg amqp.Delivery) {
		log.Println(string(msg.Body))
	}); err != nil {
		t.Error(err)
	}
	select {}
}

func TestFanoutSubscribe2(t *testing.T) {
	if err := rabbitmq.FanoutSubscribe("log", "33",func(msg amqp.Delivery) {
		log.Println(string(msg.Body))
		//msg.Ack(false)
	}); err != nil {
		t.Error(err)
	}
	select {}
}

func TestDirectPublish(t *testing.T) {
	if err := rabbitmq.DirectPublish("log-d","info",[]byte("this is info log")); err != nil {
		t.Error(err)
	}
	if err := rabbitmq.DirectPublish("log-d","err",[]byte("this is err log")); err != nil {
		t.Error(err)
	}
}

func TestDirectSubscribe(t *testing.T) {
	if err := rabbitmq.DirectSubscribe("log-d", "","info", func(msg amqp.Delivery) {
		log.Println(string(msg.Body))
	}); err != nil {
		t.Error(err)
	}
	select {}
}

func TestDirectSubscribe2(t *testing.T) {
	if err := rabbitmq.DirectSubscribe("log-d", "","err",func(msg amqp.Delivery) {
		log.Println(string(msg.Body))
		//msg.Ack(false)
	}); err != nil {
		t.Error(err)
	}
	select {}
}

func TestTopicPublish(t *testing.T) {
	if err := rabbitmq.TopicPublish("log-to","a.info",[]byte("this is info log")); err != nil {
		t.Error(err)
	}
	if err := rabbitmq.TopicPublish("log-to","a.err",[]byte("this is err log")); err != nil {
		t.Error(err)
	}
	if err := rabbitmq.TopicPublish("log-to","b.info",[]byte("this is info log")); err != nil {
		t.Error(err)
	}
	if err := rabbitmq.TopicPublish("log-to","b.err",[]byte("this is err log")); err != nil {
		t.Error(err)
	}
}

func TestTopicSubscribe(t *testing.T) {
	if err := rabbitmq.TopicSubscribe("log-to", "","a.*", func(msg amqp.Delivery) {
		log.Println(string(msg.Body))
	}); err != nil {
		t.Error(err)
	}
	select {}
}

func TestTopicSubscribe2(t *testing.T) {
	if err := rabbitmq.TopicSubscribe("log-to", "","*.err",func(msg amqp.Delivery) {
		log.Println(string(msg.Body))
		//msg.Ack(false)
	}); err != nil {
		t.Error(err)
	}
	select {}
}
