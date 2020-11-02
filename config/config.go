// @Description  TODO
// @Author  	 jiangyang  
// @Created  	 2020/10/30 5:30 下午
package config

import (
	"github.com/comeonjy/util/mysql"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"log"
)

var c Config

// 配置结构体
type Config struct {
	Mysql mysql.Config `mapstructure:"mysql"`
}

func GetConfig() Config {
	return c
}

// 加载配置
func LoadConfig(cfgFile ...string) Config {
	if len(cfgFile)>0 && cfgFile[0] != "" {
		viper.SetConfigFile(cfgFile[0])
	} else {
		viper.SetConfigName("config")
		viper.SetConfigType("yaml")
		viper.AddConfigPath(".")
		viper.AddConfigPath("./../")
	}
	if err := viper.ReadInConfig(); err != nil {
		log.Fatal(err)
	}

	logrus.Info("use config file:", viper.ConfigFileUsed())

	err := viper.Unmarshal(&c)
	if err != nil {
		log.Fatal(err)
	}
	return c
}
