// @Description  mongo
// @Author  	 jiangyang
// @Created  	 2020/11/17 4:12 下午

// Example Config:
// mongodb:
//   addr: 127.0.0.1:27017
//   database: demo
//   username:
//   password:

package mongodb

import (
	"context"
	"fmt"
	"sync"
	"time"

	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	database *mongo.Database
	client   *mongo.Client
	once     sync.Once
)

func GetConn(collectionName string) *mongo.Collection {
	return database.Collection(collectionName)
}

type Config struct {
	Username string `json:"username" yaml:"username"`
	Password string `json:"password" yaml:"password"`
	Addr     string `json:"addr" yaml:"addr"`
	Database string `json:"database" yaml:"database"`
}

func Init(cfg Config) {

	once.Do(func() {
		format := `mongodb://%s:%s@%s/%s`

		if cfg.Username == "" || cfg.Password == "" {
			format = `mongodb://%s%s%s/%s`
		}

		uri := fmt.Sprintf(format,
			cfg.Username,
			cfg.Password,
			cfg.Addr,
			cfg.Database,
		)
		opt := options.Client().ApplyURI(uri)

		var err error
		client, err = mongo.NewClient(opt)
		if err != nil {
			logrus.Fatalf("couldn't connect to mongo: %v", err)
		}

		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		err = client.Connect(ctx)
		if err != nil {
			logrus.Fatalf("mongo client couldn't connect with background context: %v", err)
		}

		database = client.Database(cfg.Database)
		logrus.Info("mongo connect successfully")
	})
}

func Close() error {
	if client != nil {
		if err := client.Disconnect(nil); err != nil {
			return errors.WithStack(err)
		}
	}
	logrus.Info("mongo connect disconnected")
	return nil
}
