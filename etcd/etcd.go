// Package etcd @Description  TODO
// @Author  	 jiangyang  
// @Created  	 2021/5/4 12:23 下午
package etcd

import (
	"go.etcd.io/etcd/client/v3"
	"log"
	"time"
)

type Config struct {
	Endpoints   []string `json:"endpoints" yaml:"endpoints"`
	DialTimeOut int64    `json:"dial_time_out" yaml:"dial_time_out" mapstructure:"dial_time_out"`
}

var cfg Config

var client *clientv3.Client

func Init(cfg Config) {
	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   cfg.Endpoints,
		DialTimeout: time.Duration(cfg.DialTimeOut) * time.Second,
	})
	if err != nil {
		log.Fatal(err)
	}
	client = cli
}

func Conn() *clientv3.Client {
	return client
}
