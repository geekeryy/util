// @Description  TODO
// @Author  	 jiangyang  
// @Created  	 2021/2/5 2:39 下午
package tencent_cos_test

import (
	"github.com/comeonjy/util/config"
	"github.com/comeonjy/util/tencent/tencent_cos"
	"testing"
)

func TestInsert(t *testing.T) {
	c := tencent_cos.NewClient(config.GetConfig().TencentCos)
	c.Insert()
}

func TestDelete(t *testing.T) {
	c := tencent_cos.NewClient(config.GetConfig().TencentCos)
	c.Delete()
}
