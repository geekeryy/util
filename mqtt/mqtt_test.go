// @Description  TODO
// @Author  	 jiangyang  
// @Created  	 2020/11/19 10:53 上午
package mqtt_test

import (
	"github.com/comeonjy/util/config"
	mqttx "github.com/comeonjy/util/mqtt"
	mqtt "github.com/eclipse/paho.mqtt.golang"
	"github.com/sirupsen/logrus"
	"testing"
)

func TestSubscribe(t *testing.T) {
	mqttx.Init(config.GetConfig().Mqtt)
	logrus.SetFormatter(&logrus.TextFormatter{})

	defer mqttx.Close()

	if err := mqttx.Subscribe(onMessage, 2, "demo/#"); err != nil {
		t.Error(err)
	}

	select {}
}

func onMessage(client mqtt.Client, msg mqtt.Message) {
	logrus.Info(msg.MessageID(), msg.Topic(), string(msg.Payload()), msg.Qos(), msg.Retained())
}

func TestPublish1(t *testing.T) {
	cfg := config.GetConfig().Mqtt
	cfg.ClientID = "1"
	mqttx.Init(cfg)
	defer mqttx.Close()
	msg := mqttx.Message{
		ClientID: cfg.ClientID,
		Data:     "hello1",
	}
	if err := mqttx.Publish("demo/1", 2, false, msg); err != nil {
		t.Error(err)
	}
	select {

	}

}

func TestPublish2(t *testing.T) {
	cfg := config.GetConfig().Mqtt
	cfg.ClientID = "2"
	mqttx.Init(cfg)
	defer mqttx.Close()
	msg := mqttx.Message{
		ClientID: cfg.ClientID,
		Data:     "hello2",
	}
	if err := mqttx.Publish("demo/2", 2, false, msg); err != nil {
		t.Error(err)
	}
	select {

}
}
