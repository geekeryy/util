// @Description  TODO
// @Author  	 jiangyang  
// @Created  	 2020/10/30 5:38 下午
package config_test

import (
	"github.com/comeonjy/util/config"
	"log"
	"testing"
)

func TestInitConfig(t *testing.T) {
	log.Println(config.LoadConfig())
}
