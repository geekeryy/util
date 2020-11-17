// @Description  日志
// @Author  	 jiangyang
// @Created  	 2020/11/17 4:12 下午
package log

import (
	"github.com/sirupsen/logrus"
)

type Config struct {
	Format string `json:"format" yaml:"format"`
}

// 设置日志格式
func Init(cfg Config)  {
	logrus.SetReportCaller(true)
	switch cfg.Format {
	case "json":
		logrus.SetFormatter(&logrus.JSONFormatter{})
	case "text":
		logrus.SetFormatter(&logrus.TextFormatter{})
	}
}
