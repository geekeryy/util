// @Description  TODO
// @Author  	 jiangyang  
// @Created  	 2020/11/19 10:53 上午
package mqtt_test

import (
	"encoding/json"
	"github.com/comeonjy/util/config"
	"github.com/comeonjy/util/elastic"
	"github.com/comeonjy/util/log"
	mqttx "github.com/comeonjy/util/mqtt"
	mqtt "github.com/eclipse/paho.mqtt.golang"
	"github.com/sirupsen/logrus"
	"testing"
	"time"
)

func TestSubscribe(t *testing.T) {
	elastic.Init(config.GetConfig().Elastic)
	log.Init(config.GetConfig().Log)
	mqttx.Init(config.GetConfig().Mqtt)

	defer mqttx.Close()

	if err := mqttx.Subscribe(onMessage, 2, "demo/1"); err != nil {
		t.Error(err)
	}

	if err := mqttx.Subscribe(onMessage, 2, "demo/2"); err != nil {
		t.Error(err)
	}

	time.Sleep(5 * time.Second)
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
		Time:     time.Now().Unix(),
	}
	data, err := json.Marshal(msg)
	if err != nil {
		t.Error(err)
	}
	if err := mqttx.Publish("demo/1", 2, false, data); err != nil {
		t.Error(err)
	}
	select {}

}

func TestPublish2(t *testing.T) {
	cfg := config.GetConfig().Mqtt
	cfg.ClientID = "2"
	mqttx.Init(cfg)
	defer mqttx.Close()
	msg := mqttx.Message{
		ClientID: cfg.ClientID,
		Data:     "hello2",
		Time:     time.Now().Unix(),
	}
	data, err := json.Marshal(msg)
	if err != nil {
		t.Error(err)
	}
	if err := mqttx.Publish("demo/2", 2, false, data); err != nil {
		t.Error(err)
	}
	select {}
}
