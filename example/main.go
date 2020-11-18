// @Description  TODO
// @Author  	 jiangyang  
// @Created  	 2020/11/9 5:54 下午
package main

import (
	"encoding/json"
	"github.com/comeonjy/util/config"
	"github.com/comeonjy/util/ctx"
	"github.com/comeonjy/util/errno"
	"github.com/comeonjy/util/jwt"
	"github.com/comeonjy/util/middlewares"
	"github.com/comeonjy/util/mysql"
	"github.com/comeonjy/util/rbac"
	"github.com/comeonjy/util/server"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"net/http"
	"time"
)

func init() {
	config.LoadConfig()
	mysql.Init(config.GetConfig().Mysql)
	rbac.Init(config.GetConfig().Rbac)
}

func main() {
	r := gin.Default()
	r.GET("", func(ctx *gin.Context) {
		logrus.Info("sleep...start")
		//time.Sleep(4 * time.Second)
		logrus.Info("sleep...end")
		ctx.JSON(http.StatusOK, gin.H{"msg": "ok"})
	})

	r.GET("token", ctx.Handle(token))

	auth := r.Group("")
	auth.Use(middlewares.JwtAuth())
	auth.GET("/ping", ctx.Handle(ping))
	auth.Use(middlewares.Rbac(rbac.Check)).GET("/auth", ctx.Handle(ping))
	server.Server(r, viper.GetInt("http_port"))

}

func token(ctx *ctx.Context) {
	bus := jwt.Business{
		UID:  1,
		Role: 2,
	}
	tokenResp, err := jwt.CreateToken(bus, 24*time.Hour)
	if err != nil {
		ctx.Fail(err)
		return
	}
	ctx.Success(tokenResp)
}

func ping(ctx *ctx.Context) {
	bus, exists := ctx.Get("business")
	if !exists {
		ctx.Fail(errno.BusNotFound)
		return
	}
	b := jwt.Business{}
	marshal, err := json.Marshal(bus)
	if err != nil {
		ctx.Fail(err)
		return
	}
	if err := json.Unmarshal(marshal, &b); err != nil {
		ctx.Fail(err)
		return
	}
	ctx.Success(b)
}
