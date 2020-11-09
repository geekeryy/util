// @Description  TODO
// @Author  	 jiangyang  
// @Created  	 2020/11/2 10:29 上午
package server

import (
	"context"
	"github.com/gin-contrib/pprof"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"syscall"
	"time"
)

func Server(router *gin.Engine, port int) {
	pprof.Register(router, "/pprof")
	srv := &http.Server{
		Addr:    ":" + strconv.Itoa(port),
		Handler: router,
	}

	logrus.Info("Server Starting...")

	go func() {
		if err := srv.ListenAndServe(); err != http.ErrServerClosed {
			logrus.Fatalf("Server start failed: %s", err)
		}
	}()

	logrus.Infof("Server started at http://127.0.0.1:%d", port)

	// 等待中断信号以优雅地关闭服务器（设置 5 秒的超时时间）
	quit := make(chan os.Signal)

	// kill (no param) default send syscall.SIGTERM
	// kill -2 is syscall.SIGINT
	// kill -9 is syscall.SIGKILL but can"t be catch, so don't need add it
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	logrus.Info("Shutdown Server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		logrus.Fatalf("Server Shutdown: %v", err)
	}

	logrus.Info("Server exiting")

}
