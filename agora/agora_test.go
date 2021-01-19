// @Description  TODO
// @Author  	 jiangyang  
// @Created  	 2021/1/5 4:47 下午
package agora_test

import (
	"github.com/comeonjy/util/agora"
	"github.com/sirupsen/logrus"
	"testing"
)

func TestGenerateCredential(t *testing.T) {
	logrus.Info(agora.GenerateCredential())
}

func TestHttpGet(t *testing.T) {
	acquireReq:=agora.AcquireReq{
		Cname: "demo",
		UID:   "1",
	}
	if resp,err := agora.Acquire(&acquireReq);err!=nil{
		logrus.Fatal(err)
	}else{
		logrus.Info(resp)
	}
}