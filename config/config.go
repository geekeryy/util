// @Description  加载配置
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

var c  *Config

// 配置结构体
type Config struct {
	Mysql   mysql.Config   `mapstructure:"mysql"`
	Log     log.Config     `mapstructure:"log"`
	Mongodb mongodb.Config `mapstructure:"mongodb"`
	Redis   redis.Config   `mapstructure:"redis"`
}

// 获取配置信息
// 全局只加载一次配置
func GetConfig(cfgFile ...string) *Config {
	if c != nil {
		return c
	}
	return LoadConfig(cfgFile...)
}

// 加载配置
// 可多次加载不同配置
func LoadConfig(cfgFile ...string) *Config {
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

	c = &Config{}
	if err := viper.Unmarshal(c); err != nil {
		logrus.Fatal(err)
	}

	return c
}
