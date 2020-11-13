// @Description  TODO
// @Author  	 jiangyang  
// @Created  	 2020/10/30 5:30 下午
package config

import (
	"github.com/comeonjy/util/log"
	"github.com/comeonjy/util/mongodb"
	"github.com/comeonjy/util/mysql"
	"github.com/comeonjy/util/redis"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

var c Config

// 配置结构体
type Config struct {
	Mysql   mysql.Config   `mapstructure:"mysql"`
	Log     log.Config     `mapstructure:"log"`
	Mongodb mongodb.Config `mapstructure:"mongodb"`
	Redis   redis.Config   `mapstructure:"redis"`
}

func GetConfig() Config {
	return c
}

// 加载配置
func LoadConfig(cfgFile ...string) Config {
	if len(cfgFile) > 0 && cfgFile[0] != "" {
		viper.SetConfigFile(cfgFile[0])
	} else {
		viper.SetConfigName("config")
		viper.SetConfigType("yaml")
		viper.AddConfigPath(".")
		viper.AddConfigPath("./../")
	}
	if err := viper.ReadInConfig(); err != nil {
		logrus.Fatal(err)
	}

	logrus.Info("use config file:", viper.ConfigFileUsed())

	err := viper.Unmarshal(&c)
	if err != nil {
		logrus.Fatal(err)
	}
	return c
}
