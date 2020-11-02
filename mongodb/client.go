package mongodb

import (
	"context"
	"fmt"
	"github.com/pkg/errors"
	"sync"
	"time"

	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	database *mongo.Database
	client   *mongo.Client
	once     sync.Once
)

// GetCollection ...
func GetCollection(collectionName string) *mongo.Collection {
	return database.Collection(collectionName)
}

// Init ...
func Init() {

	once.Do(func() {

		uri := fmt.Sprintf(`mongodb://%s:%s@%s/%s`,
			viper.GetString("db.mongodb.username"),
			viper.GetString("db.mongodb.password"),
			viper.GetString("db.mongodb.addr"),
			viper.GetString("db.mongodb.database"),
		)

		var err error
		client, err = mongo.NewClient(options.Client().ApplyURI(uri))
		if err != nil {
			logrus.Fatalf("couldn't connect to mongo: %v", err)
		}

		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		err = client.Connect(ctx)
		if err != nil {
			logrus.Fatalf("mongo client couldn't connect with background context: %v", err)
		}

		database = client.Database(viper.GetString("db.mongodb.database"))
		logrus.Info("mongo connect successfully")
	})
}

//Disconnect method
func Disconnect() error {
	if client != nil {
		if err := client.Disconnect(nil); err != nil {
			return errors.WithStack(err)
		}
	}
	logrus.Info("mongo connect disconnected")
	return nil
}
