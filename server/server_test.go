// @Description  TODO
// @Author  	 jiangyang  
// @Created  	 2020/11/2 10:30 上午
package server_test

import (
	"github.com/comeonjy/util/config"
	"github.com/comeonjy/util/server"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"testing"
)

func init()  {
	config.LoadConfig()
}

func TestServer(t *testing.T) {
	server.Server(gin.Default(),viper.GetInt("http_port"))
}
