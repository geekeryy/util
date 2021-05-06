// @Description  TODO
// @Author  	 jiangyang  
// @Created  	 2021/2/5 2:39 下午
package tencent_cos_test

import (
	"context"
	"fmt"
	"github.com/comeonjy/util/config"
	"github.com/comeonjy/util/tencent/tencent_cos"
	"testing"
)

var c *tencent_cos.TencentCos

func init() {
	c = tencent_cos.NewClient(config.GetConfig().TencentCos)
}

func TestDemo(t *testing.T)  {
	fmt.Println("123"[2:])
}

func TestTencentCos_Get(t *testing.T) {
	fmt.Println(c.Get("examcloud/prod/deep3/2773"))
}

func TestInsertLifecycle(t *testing.T) {
	c.InsertLifecycle()
}

func TestDeleteLifecycle(t *testing.T) {
	c.DeleteLifecycle()
}

func TestTencentCos_DownLoadFile(t *testing.T) {
	_, err := c.Object.GetToFile(context.Background(), "q/demo1.jpg", "./xixi.jpg", nil)
	if err != nil {
		t.Error(err)
	}
}

func TestTencentCos_UploadFile(t *testing.T) {
	_, err := c.Object.PutFromFile(context.Background(), "", "./xixi.jpg", nil)
	if err != nil {
		t.Error(err)
	}
}
