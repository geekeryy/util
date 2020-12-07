// @Description  TODO
// @Author  	 jiangyang  
// @Created  	 2020/12/7 11:14 上午
package excle_test

import (
	"fmt"
	"github.com/comeonjy/util/excle"
	"testing"
)

type DemoModel struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
}

func TestSave(t *testing.T) {
	elx, err := excle.New("")
	if err != nil {
		t.Error(err)
	}
	demo := DemoModel{
		ID:   1,
		Name: "jy",
	}
	if err := elx.Save(demo); err != nil {
		t.Error(err)
	}
}

func TestAxis(t *testing.T) {
	fmt.Println(CompareString(excle.Axis(1, 1, 0), "A1"))
	fmt.Println(CompareString(excle.Axis(26, 1, 0), "Z1"))
	fmt.Println(CompareString(excle.Axis(26+1, 1, 0), "AA1"))
	fmt.Println(CompareString(excle.Axis(26+2, 1, 0), "AB1"))
	fmt.Println(CompareString(excle.Axis(26+26, 1, 0), "AZ1"))
	fmt.Println(CompareString(excle.Axis(26+26+1, 1, 0), "BA1"))
	fmt.Println(CompareString(excle.Axis(26+26+2, 1, 0), "BB1"))
	fmt.Println(CompareString(excle.Axis(26*26, 1, 0), "YZ1"))
}

func CompareString(a, b string) bool {
	return a == b
}
