// @Description  日志
// @Author  	 jiangyang
// @Created  	 2020/11/17 4:12 下午

// Example Config:
// log:
//  format: json
//  hooks:
//   - elasticsearch
//   - email
//   - mobile
//  emails:
//   - 1********8@qq.com
//   - 151****1234
//  level: info
//  report_caller: true

package log

import (
	"github.com/comeonjy/util/elastic"
	"github.com/comeonjy/util/email"
	"github.com/sirupsen/logrus"
)

type Config struct {
	Format       string   `json:"format" yaml:"format"`
	Level        string   `json:"level" yaml:"level"`
	Hooks        []string `json:"hooks" yaml:"hooks"`
	Emails       []string `json:"emails" yaml:"emails"`
	Mobile       string   `json:"mobile" yaml:"mobile"`
	ReportCaller bool     `json:"report_caller" yaml:"report_caller" mapstructure:"report_caller"`
}

// 初始化日志
func Init(cfg Config) {

	level, err := logrus.ParseLevel(cfg.Level)
	if err != nil {
		logrus.Errorf("日志等级解析失败 need[panic,fatal,error,warn,warning,info,debug,trace] get:[%s]", cfg.Level)
	}

	logrus.SetLevel(level)

	logrus.SetReportCaller(cfg.ReportCaller)

	for _, hook := range cfg.Hooks {
		switch hook {
		case "elasticsearch":
			if elastic.Conn() != nil {
				logrus.AddHook(&EsHook{})
				logrus.Info("日志Hook添加成功：", hook)
			} else {
				logrus.Error("elasticsearch client 未初始化 hook未生效")
			}
		case "email":
			if email.Conn() != nil {
				logrus.AddHook(&EmailHook{MailTo: cfg.Emails})
				logrus.Info("日志Hook添加成功：", hook)
			} else {
				logrus.Error("email 未初始化 hook未生效")
			}
		case "mobile":
			// TODO
		}
	}

	switch cfg.Format {
	case "text":
		logrus.SetFormatter(&logrus.TextFormatter{})
	default:
		logrus.SetFormatter(&logrus.JSONFormatter{})
	}
}
