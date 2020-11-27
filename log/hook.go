// @Description  logrus hook
// @Author  	 jiangyang  
// @Created  	 2020/11/27 11:27 上午
package log

import (
	"fmt"
	"github.com/comeonjy/util/elastic"
	"github.com/comeonjy/util/email"
	"github.com/sirupsen/logrus"
	"time"
)

// EsHook
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
	if entry.Caller != nil {
		doc["caller"] = fmt.Sprintf("%s:%d", entry.Caller.File, entry.Caller.Line)
	}

	if err := elastic.Index("demo", doc); err != nil {
		return err
	}
	return nil
}


// EmailHook
type EmailHook struct{
	MailTo []string
}

func (hook *EmailHook) Levels() []logrus.Level {
	return []logrus.Level{
		logrus.ErrorLevel,
	}
}
func (hook *EmailHook) Fire(entry *logrus.Entry) error {
	body,err:=entry.String()
	if err != nil {
		return err
	}
	return email.SendMail(hook.MailTo, "服务报错：", body)
}
