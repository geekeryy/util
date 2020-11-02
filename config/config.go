// @Description  TODO
// @Author  	 jiangyang  
// @Created  	 2020/10/30 5:30 下午
package config

import (
	"github.com/comeonjy/util/mysql"
	"github.com/spf13/viper"
	"log"
)

var c Config

// 配置结构体
type Config struct {
	Mysql mysql.Config `mapstructure:"mysql"`
}

// 加载配置
func LoadConfig(filename string) Config {
	viper.SetConfigFile(filename)
	if err := viper.ReadInConfig();err!=nil{
		log.Fatal(err)
	}
	err := viper.Unmarshal(&c)
	if err != nil {
		log.Fatal(err)
	}
	return c
}
