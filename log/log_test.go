// @Description  TODO
// @Author  	 jiangyang  
// @Created  	 2020/11/10 5:59 下午
package log_test

import (
	"github.com/comeonjy/util/config"
	"github.com/comeonjy/util/log"
	"github.com/sirupsen/logrus"
	"testing"
)

func init()  {
	config.LoadConfig()
}

func TestInit(t *testing.T) {
	log.Init(config.GetConfig().Log)
	logrus.Info("xixi")
}
