// @Description  TODO
// @Author  	 jiangyang  
// @Created  	 2020/12/7 11:14 上午
package excel_test

import (
	"fmt"
	"testing"

	"github.com/comeonjy/util/excel"
)

type DemoModel struct {
	ID        uint32                 `json:"id" excel:"序号"`
	Age       int16                  `json:"age" excel:"年龄"`
	Name      string                 `json:"name" excel:"姓名"`
	IsStudent bool                   `json:"is_student" excel:"是否学生"`
	Score     float32                `json:"score" excel:"得分"`
	Point     float64                `json:"point" excel:"点"`
	No        interface{}            `json:"no" excel:"-"`
}

func TestSave(t *testing.T) {
	demo := &DemoModel{
		1, 10, "jy", true, 1.11, 1.234,2,
	}
	demos := make([]DemoModel, 0)
	demos = append(demos, *demo, *demo)
	t.Run("create", func(t *testing.T) {
		if err := excel.New(excel.FileNameOption("./1.xlsx"), excel.SheetNameOption("Sheet1")).Create(demos); err != nil {
			t.Error(err)
		}
	})

	t.Run("insert struct", func(t *testing.T) {
		if err := excel.New(excel.FileNameOption("./1.xlsx"), excel.SheetNameOption("Sheet1")).Insert(demo); err != nil {
			t.Error(err)
		}
	})

	t.Run("insert slice", func(t *testing.T) {
		if err := excel.New(excel.FileNameOption("./1.xlsx"), excel.SheetNameOption("Sheet1")).Insert(demos); err != nil {
			t.Error(err)
		}
	})

	t.Run("read", func(t *testing.T) {
		s := make([]DemoModel, 0)
		if err := excel.New(excel.FileNameOption("./1.xlsx"), excel.SheetNameOption("Sheet1")).Read(&s); err != nil {
			t.Error(err)
		} else {
			for _,v:=range s{
				fmt.Println(v)
			}
		}
	})

}

func TestExcel_Save(t *testing.T) {

}

func TestAxis(t *testing.T) {
	fmt.Println(CompareString(excel.Axis(1, 1), "A1"))
	fmt.Println(CompareString(excel.Axis(26, 1), "Z1"))
	fmt.Println(CompareString(excel.Axis(26+1, 1), "AA1"))
	fmt.Println(CompareString(excel.Axis(26+2, 1), "AB1"))
	fmt.Println(CompareString(excel.Axis(26+26, 1), "AZ1"))
	fmt.Println(CompareString(excel.Axis(26+26+1, 1), "BA1"))
	fmt.Println(CompareString(excel.Axis(26+26+2, 1), "BB1"))
	fmt.Println(CompareString(excel.Axis(26*26, 1), "YZ1"))
}

func CompareString(a, b string) bool {
	return a == b
}
