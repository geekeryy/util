// @Description  TODO
// @Author  	 jiangyang  
// @Created  	 2020/11/19 10:51 上午
package mqtt

import (
	"encoding/json"
	"github.com/sirupsen/logrus"
	"time"

	"github.com/eclipse/paho.mqtt.golang"
)

var cfg Config
var client mqtt.Client
var topicArr []string

type Config struct {
	Broker    string `json:"broker" yaml:"broker"`
	ClientID  string `json:"client_id" yaml:"client_id" mapstructure:"client_id"`
	KeepAlive int64  `json:"keep_alive" yaml:"keep_alive" mapstructure:"keep_alive"`
}

var f mqtt.MessageHandler = func(client mqtt.Client, msg mqtt.Message) {
	logrus.Info(msg.MessageID(),msg.Topic(),string(msg.Payload()),msg.Qos(),msg.Retained())
}

func Init(c Config) {
	cfg = c
	mqtt.ERROR = logrus.New()
	mqtt.CRITICAL = logrus.New()
	mqtt.WARN = logrus.New()
	//mqtt.DEBUG = logrus.New()

	opts := mqtt.NewClientOptions().AddBroker(cfg.Broker).SetClientID(cfg.ClientID)

	opts.SetKeepAlive(time.Duration(cfg.KeepAlive) * time.Second)
	opts.SetDefaultPublishHandler(f)
	opts.SetPingTimeout(5 * time.Second)
	opts.SetCleanSession(false)
	opts.SetAutoReconnect(true)

	client = mqtt.NewClient(opts)
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		logrus.Fatal(token.Error().Error())
	}
}

func Conn() mqtt.Client {
	return client
}

func Close() {
	client.Unsubscribe(topicArr...)
	client.Disconnect(1000)
}

type Message struct {
	ClientID string      `json:"client_id"`
	Data     interface{} `json:"data"`
	Time     int64       `json:"time"`
}

func Publish(topic string, qos byte, retained bool, msg Message) error {
	msg.ClientID = cfg.ClientID
	msg.Time = time.Now().Unix()
	data, err := json.Marshal(msg)
	if err != nil {
		return err
	}
	if token := client.Publish(topic, qos, retained, data); token.Wait() && token.Error() != nil {
		return token.Error()
	}
	return nil
}

func Subscribe(onMessage mqtt.MessageHandler, qos byte, topics ...string) error {
	filters := make(map[string]byte)
	for _, v := range topics {
		filters[v] = qos
	}
	if token := client.SubscribeMultiple(filters, onMessage); token.Wait() && token.Error() != nil {
		return token.Error()
	}
	topicArr=append(topicArr,topics...)
	return nil
}
