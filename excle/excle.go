// @Description  TODO
// @Author  	 jiangyang  
// @Created  	 2020/12/7 11:10 上午
package excle

import (
	"fmt"
	"github.com/360EntSecGroup-Skylar/excelize/v2"
	"reflect"
	"strconv"
)

type Excel struct {
	FileName string
	file     *excelize.File
}

func New(fileName string) (*Excel, error) {
	e := &Excel{
		FileName: fileName,
	}
	if fileName != "" {
		file, err := excelize.OpenFile(fileName)
		if err != nil {
			return nil, err
		}
		e.file = file

	} else {
		e.file = excelize.NewFile()
	}
	return e, nil
}

func (e *Excel) Save(data interface{}) error {
	//row := make([][]string, 0)
	t := reflect.TypeOf(data)
	for {
		if t.Kind() == reflect.Ptr {
			t = t.Elem()
		} else {
			break
		}
	}

	SheetName := t.Name()

	v := reflect.ValueOf(data)

	for i := 0; i < t.NumField(); i++ {
		name := t.Field(i).Name
		e.file.SetCellStr(SheetName, "", name)
	}

	fmt.Println(SheetName, v)

	return nil
}

func Axis(i, j int, offset int) string {
	b := make([]byte, 0)
	To24(i+offset, &b)
	fmt.Println(string(b) + strconv.Itoa(j))
	return string(b) + strconv.Itoa(j)
}

// 十进制转26进制
func To24(i int, b *[]byte) {
	remainder := i % 26
	quotient := i / 26
	if remainder == 0 {
		quotient--
		remainder=26
	}
	if quotient > 26 {
		To24(remainder, b)
	} else {
		if quotient > 0 {
			*b = append(*b, uint8(quotient)+64)
		}
	}
	*b = append(*b, uint8(remainder)+64)

}
