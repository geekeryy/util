// @Description  发送邮件
// @Author  	 jiangyang  
// @Created  	 2020/11/17 4:12 下午

// Example Config:
// email:
//   user:
//   pass:
//   host: smtp.qq.com
//   port: 465

package email

import (
	"github.com/sirupsen/logrus"
	"gopkg.in/gomail.v2"
)

var cfg *Config

type Config struct {
	User string `json:"user" yaml:"user"`
	Pass string `json:"pass" yaml:"pass"`
	Host string `json:"host" yaml:"host"`
	Port int    `json:"port" yaml:"port"`
}

func Init(c Config) {
	cfg = &c
	logrus.Info("email init successfully")
}

func Conn() *Config {
	return cfg
}

// 发送单封邮件
func SendMail(mailTo []string, subject string, body string) error {
	m := gomail.NewMessage()
	m.SetHeader("From", "<"+cfg.User+">")
	m.SetHeader("To", mailTo...)
	m.SetHeader("Subject", subject)
	m.SetBody("text/html", body)
	d := gomail.NewDialer(cfg.Host, cfg.Port, cfg.User, cfg.Pass)
	err := d.DialAndSend(m)
	return err
}
