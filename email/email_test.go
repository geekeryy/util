package email

import (
	"github.com/comeonjy/util/config"
	"testing"
)

func init()  {
	config.LoadConfig()
}

func TestSendMail(t *testing.T) {
	err := SendMail([]string{"1126254578@qq.com"}, "subject", "你好")
	if err != nil {
		t.Error(err)
	}
}
