// @Description  TODO
// @Author  	 jiangyang  
// @Created  	 2020/11/9 5:54 下午
package main

import (
	"github.com/comeonjy/util/config"
	"github.com/comeonjy/util/ctx"
	"github.com/comeonjy/util/server"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"net/http"
	"time"
)

func init() {
	config.LoadConfig()
}

func main()  {
	r := gin.Default()
	r.GET("", func(ctx *gin.Context) {
		logrus.Info("sleep...start")
		time.Sleep(4*time.Second)
		logrus.Info("sleep...end")
		ctx.JSON(http.StatusOK, gin.H{"msg": "ok"})
	})

	r.GET("/ping", ctx.Handle(handle))

	server.Server(r, viper.GetInt("http_port"))

}

func handle(ctx *ctx.Context)  {
	ctx.Success(nil)
}