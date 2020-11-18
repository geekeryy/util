package email_test

import (
	"github.com/comeonjy/util/config"
	"github.com/comeonjy/util/email"
	"testing"
)

func init()  {
	email.Init(config.GetConfig().Email)
}

func TestSendMail(t *testing.T) {
	err := email.SendMail([]string{"1126254578@qq.com"}, "subject", "你好")
	if err != nil {
		t.Error(err)
	}
}
