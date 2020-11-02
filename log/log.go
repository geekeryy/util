package log

import (
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

// 设置日志格式
func Init()  {
	logrus.SetReportCaller(true)
	switch viper.GetString("log.format") {
	case "json":
		logrus.SetFormatter(&logrus.JSONFormatter{})
	case "text":
		logrus.SetFormatter(&logrus.TextFormatter{})
	}
}
