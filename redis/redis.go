// @Description  TODO
// @Author  	 jiangyang  
// @Created  	 2020/11/2 10:02 上午
package redis

import (
	"context"
	"github.com/go-redis/redis/v8"
	"github.com/sirupsen/logrus"
)

var (
	client *redis.Client
)

type Config struct {
	Addr     string `json:"addr" yaml:"addr"`
	Password string `json:"password" yaml:"password"`
	Db       int    `json:"db" yaml:"db"`
	PoolSize int    `json:"pool_size" yaml:"pool_size"`
}

func Init(cfg Config) {
	client = redis.NewClient(&redis.Options{
		Addr:     cfg.Addr,
		Password: cfg.Password,
		DB:       cfg.Db,
		PoolSize: cfg.PoolSize,
	})
	if _, err := client.Ping(context.Background()).Result(); err != nil {
		logrus.Fatal(err)
	}

}

func GetConn() *redis.Client {
	return client
}

func Close() {
	if client != nil {
		client.Close()
	}
}
