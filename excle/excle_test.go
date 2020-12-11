// @Description  TODO
// @Author  	 jiangyang  
// @Created  	 2020/12/7 11:14 上午
package excle_test

import (
	"fmt"
	"github.com/comeonjy/util/excle"
	"github.com/sirupsen/logrus"
	"testing"
)

type DemoModel struct {
	ID        uint    `json:"id"`
	Age       int64   `json:"age"`
	Name      string  `json:"name"`
	IsStudent bool    `json:"is_student"`
	Score     float32 `json:"score"`
	Point     float64 `json:"point"`
	Demo     *DemoModel `json:"demo"`
}

func TestSave(t *testing.T) {
	demo := &DemoModel{
		1, 10, "jy", true, 1.11, 1.234,&DemoModel{ID: 1},
	}
	if err := excle.New("./1.xlsx", "demo").Create(demo); err != nil {
		t.Error(err)
	}

	s := make([]DemoModel, 0)
	if err := excle.New("./1.xlsx", "demo").Read(&s); err != nil {
		t.Error(err)
	}
	logrus.Println(s)

}

func TestExcel_Save(t *testing.T) {

}

func TestAxis(t *testing.T) {
	fmt.Println(CompareString(excle.Axis(1, 1), "A1"))
	fmt.Println(CompareString(excle.Axis(26, 1), "Z1"))
	fmt.Println(CompareString(excle.Axis(26+1, 1), "AA1"))
	fmt.Println(CompareString(excle.Axis(26+2, 1), "AB1"))
	fmt.Println(CompareString(excle.Axis(26+26, 1), "AZ1"))
	fmt.Println(CompareString(excle.Axis(26+26+1, 1), "BA1"))
	fmt.Println(CompareString(excle.Axis(26+26+2, 1), "BB1"))
	fmt.Println(CompareString(excle.Axis(26*26, 1), "YZ1"))
}

func CompareString(a, b string) bool {
	return a == b
}
