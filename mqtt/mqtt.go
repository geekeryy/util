// @Description  mqtt消息服务
// @Author  	 jiangyang  
// @Created  	 2020/11/19 10:51 上午

// Example Config:
// mqtt:
//   broker: tcp://broker.emqx.io:1883
//   client_id: webhook_server
//   keep_alive: 60

package mqtt

import (
	"time"

	"github.com/eclipse/paho.mqtt.golang"
	"github.com/sirupsen/logrus"
)

var cfg Config
var client mqtt.Client
var topicArr []string
var subscribeArr []func() error

type Config struct {
	Broker    string `json:"broker" yaml:"broker"`
	ClientID  string `json:"client_id" yaml:"client_id" mapstructure:"client_id"`
	KeepAlive int64  `json:"keep_alive" yaml:"keep_alive" mapstructure:"keep_alive"`
}

var f mqtt.MessageHandler = func(client mqtt.Client, msg mqtt.Message) {
	logrus.Info(msg.MessageID(), msg.Topic(), string(msg.Payload()), msg.Qos(), msg.Retained())
}

func Init(c Config) {
	cfg = c
	mqtt.ERROR = logrus.StandardLogger()
	mqtt.CRITICAL = logrus.StandardLogger()
	mqtt.WARN = logrus.StandardLogger()

	opts := mqtt.NewClientOptions().AddBroker(cfg.Broker).SetClientID(cfg.ClientID)

	opts.SetKeepAlive(time.Duration(cfg.KeepAlive) * time.Second)
	opts.SetDefaultPublishHandler(f)
	opts.SetPingTimeout(5 * time.Second)
	opts.SetCleanSession(false)
	opts.SetAutoReconnect(true)
	opts.SetConnectionLostHandler(func(c mqtt.Client, err error) {
		logrus.Info("mqtt 连接断开")
		// 重连后重新订阅
		for _,v:=range subscribeArr {
			if err:=v();err!=nil{
				logrus.Error("mqtt 重新订阅失败：",err)
			}
		}

	})
	opts.SetOnConnectHandler(func(c mqtt.Client) {
		logrus.Info("mqtt 连接成功")
	})

	client = mqtt.NewClient(opts)
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		logrus.Fatal(token.Error().Error())
	}
	logrus.Info("mqtt connect successfully")
}

func Close() {
	client.Unsubscribe(topicArr...)
	logrus.Info("mqtt Unsubscribe: ",topicArr)
	client.Disconnect(1000)
	logrus.Info("mqtt connect closed")
}

type Message struct {
	ClientID string      `json:"client_id"`
	Data     interface{} `json:"data"`
	Time     int64       `json:"time"`
}

// 发布
func Publish(topic string, qos byte, retained bool, data []byte) error {

	if token := client.Publish(topic, qos, retained, data); token.Wait() && token.Error() != nil {
		return token.Error()
	}
	return nil
}


// 订阅
func Subscribe(onMessage mqtt.MessageHandler, qos byte, topics ...string) error {
	topicArr = append(topicArr, topics...)
	f := func() error {
		return subscribe(onMessage, qos, topics...)
	}
	subscribeArr= append(subscribeArr, f)
	return f()
}


func subscribe(onMessage mqtt.MessageHandler, qos byte, topics ...string) error {
	filters := make(map[string]byte)
	for _, v := range topics {
		filters[v] = qos
	}
	if token := client.SubscribeMultiple(filters, onMessage); token.Wait() && token.Error() != nil {
		return token.Error()
	}
	logrus.Info("mqtt Subscribe: ", topics, " success")
	return nil
}
