// @Description  日志
// @Author  	 jiangyang
// @Created  	 2020/11/17 4:12 下午
package log

import (
	"fmt"
	"github.com/comeonjy/util/elastic"
	"github.com/sirupsen/logrus"
	"time"
)

type Config struct {
	Format string `json:"format" yaml:"format"`
	Hook   string `json:"hook" yaml:"hook"`
}

// 设置日志格式
func Init(cfg Config) {
	logrus.SetReportCaller(true)
	switch cfg.Hook {
	case "elasticsearch":
		logrus.AddHook(&EsHook{})
	}
	switch cfg.Format {
	case "json":
		logrus.SetFormatter(&logrus.JSONFormatter{})
	case "text":
		logrus.SetFormatter(&logrus.TextFormatter{})
	}
	logrus.Info("log ", cfg.Hook, cfg.Hook)
}

type EsHook struct{}

func (hook *EsHook) Levels() []logrus.Level {
	return logrus.AllLevels
}
func (hook *EsHook) Fire(entry *logrus.Entry) error {
	doc := make(map[string]interface{})
	for k, v := range entry.Data {
		doc[k] = v
	}
	doc["timestamp"] = time.Now().Local()
	doc["level"] = entry.Level
	doc["message"] = entry.Message
	doc["caller"] = fmt.Sprintf("%s:%d", entry.Caller.File, entry.Caller.Line)

	if err := elastic.Index("demo", doc); err != nil {
		return err
	}
	return nil
}
