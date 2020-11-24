// @Description  TODO
// @Author  	 jiangyang  
// @Created  	 2020/11/21 2:51 下午
package rabbitmq

import (
	"github.com/streadway/amqp"
	"log"
)

var conn *amqp.Connection

type Config struct {
	Addr      string `json:"addr"`
	QueueName string `json:"queue_name" mapstructure:"queue_name"`
}

func Init(cfg Config) {
	_c, err := amqp.Dial(cfg.Addr)
	if err != nil {
		log.Fatal(err)
	}
	conn = _c
}

func Conn() *amqp.Connection {
	return conn
}

func Close() {
	if conn != nil {
		conn.Close()
	}
}

// 生产者
func Produce(queueName string, body []byte) error {
	ch, err := conn.Channel()
	if err != nil {
		return err
	}
	defer ch.Close()
	q, err := ch.QueueDeclare(queueName, true, false, false, false, nil)
	if err != nil {
		return err
	}
	if err := ch.Publish("", q.Name, false, false, amqp.Publishing{
		DeliveryMode: amqp.Persistent,
		ContentType:  "text/json",
		Body:         body,
	}); err != nil {
		return err
	}
	return nil
}

// 消费者
func Consume(queueName string, msgHandler func(amqp.Delivery)) error {
	ch, err := conn.Channel()
	if err != nil {
		return err
	}
	q, err := ch.QueueDeclare(queueName, true, false, false, false, nil)
	if err != nil {
		return err
	}
	if err := ch.Qos(2, 0, false); err != nil {
		return err
	}
	consume, err := ch.Consume(q.Name, "", false, false, false, false, nil)
	if err != nil {
		return err
	}
	go func() {
		for msg := range consume {
			msgHandler(msg)
		}
	}()
	return nil
}

// 发布
func publish(excName string,key string,kind string, body []byte) error {
	ch, err := conn.Channel()
	if err != nil {
		return err
	}
	if err := ch.ExchangeDeclare(excName, kind, true, false, false, false, nil); err != nil {
		return err
	}
	if err := ch.Publish(excName, key, false, false, amqp.Publishing{
		DeliveryMode: amqp.Persistent,
		ContentType:  "text/json",
		Body:         body,
	}); err != nil {
		return err
	}
	return nil
}

// 订阅
// queueName: 为空时为临时队列不保存消息不需要ack确认，若想增加并发能力，则创建多个同名（队列名）订阅者
func subscribe(excName string,queueName string,key string,kind string, msgHandle func(amqp.Delivery)) error {
	ch, err := conn.Channel()
	if err != nil {
		return err
	}
	if err:=ch.ExchangeDeclare(excName,kind,true,false,false,false,nil);err!=nil{
		return err
	}
	q, err := ch.QueueDeclare(queueName, true, false, false, false, nil)
	if err != nil {
		return err
	}
	if err := ch.QueueBind(q.Name, key, excName, false, nil); err != nil {
		return err
	}
	consume, err := ch.Consume(q.Name, "", false, false, false, false, nil)
	if err != nil {
		return err
	}
	go func() {
		for msg := range consume {
			msgHandle(msg)
		}
	}()
	return nil
}



func FanoutPublish(excName string, body []byte) error {
	return publish(excName,"","fanout",body)
}


func FanoutSubscribe(excName string,queueName string, msgHandle func(amqp.Delivery)) error {
	return subscribe(excName,queueName,"","fanout",msgHandle)
}

func DirectPublish(excName string,key string, body []byte) error {
	return publish(excName,key,"direct",body)
}

func DirectSubscribe(excName string,queueName string,key string, msgHandle func(amqp.Delivery)) error {
	return subscribe(excName,queueName,key,"direct",msgHandle)
}

func TopicPublish(excName string,key string, body []byte) error {
	return publish(excName,key,"topic",body)
}

func TopicSubscribe(excName string,queueName string,key string, msgHandle func(amqp.Delivery)) error {
	return subscribe(excName,queueName,key,"topic",msgHandle)
}