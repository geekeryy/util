// @Description  发送邮件
// @Author  	 jiangyang  
// @Created  	 2020/11/17 4:12 下午
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

