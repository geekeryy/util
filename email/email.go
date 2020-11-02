package email

import (
	"github.com/spf13/viper"
	"gopkg.in/gomail.v2"
)

// 发送单封邮件
func SendMail(mailTo []string, subject string, body string) error {
	m := gomail.NewMessage()
	m.SetHeader("From", "<"+viper.GetString("email.user")+">")
	m.SetHeader("To", mailTo...)
	m.SetHeader("Subject", subject)
	m.SetBody("text/html", body)
	d := gomail.NewDialer(viper.GetString("email.host"), viper.GetInt("email.port"), viper.GetString("email.user"), viper.GetString("email.pass"))
	err := d.DialAndSend(m)
	return err
}
