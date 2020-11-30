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
	Level string
}

func (hook *EmailHook) Levels() []logrus.Level {
	arr:=make([]logrus.Level,0)
	level,_:=logrus.ParseLevel(hook.Level)
	for _,v:=range logrus.AllLevels{
		if v<=level {
			arr=append(arr, v)
		}
	}
	return arr
}
func (hook *EmailHook) Fire(entry *logrus.Entry) error {
	if entry.Caller != nil {
		entry.WithField("caller",fmt.Sprintf("%s:%d", entry.Caller.File, entry.Caller.Line))
	}
	body,err:=entry.String()
	if err != nil {
		return err
	}
	return email.SendMail(hook.MailTo, "服务报错：", body)
}
