// Package etcd @Description  TODO
// @Author  	 jiangyang  
// @Created  	 2021/5/5 2:54 下午
package etcd_test

import (
	"context"
	"fmt"
	"github.com/comeonjy/util/config"
	"github.com/comeonjy/util/etcd"
	clientv3 "go.etcd.io/etcd/client/v3"
	"testing"
	"time"
)

func init()  {
	etcd.Init(config.GetConfig().Etcd)
}

func TestConn(t *testing.T) {
	putResp, err := etcd.Conn().Put(context.TODO(), "name1", "jy1",clientv3.WithPrevKV())
	if err != nil {
		t.Error(err)
	}
	fmt.Println(putResp.Header.Revision,putResp.PrevKv)
}

func TestWatch(t *testing.T)  {
	ctx, _ := context.WithTimeout(context.TODO(), time.Second*5)
	watch := etcd.Conn().Watch(ctx, "name*")
	//cancelFunc()
	for {
		select {
		case resp:=<-watch:
			for k,v:=range resp.Events {
				fmt.Println(k,v.Type,v.Kv)
			}
		case <-ctx.Done():
			fmt.Println("done")
			t.Error()
		}
	}
}
