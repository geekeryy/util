// @Description  TODO
// @Author  	 jiangyang  
// @Created  	 2020/12/7 11:10 上午
package excle

import (
	"encoding/json"
	"github.com/360EntSecGroup-Skylar/excelize/v2"
	"github.com/pkg/errors"
	"os"
	"reflect"
	"strconv"
)

type Excel struct {
	FileName  string
	SheetName string
	file      *excelize.File
}

func New(fileName string, sheetName string) *Excel {
	return &Excel{
		FileName:  fileName,
		SheetName: sheetName,
	}
}

// 保存
func (e *Excel) Save(data interface{}) error {
	_, err := os.Stat(e.FileName)
	if err != nil {
		if os.IsNotExist(err) {
			return e.Create(data)
		}
	}
	return e.Insert(data)
}

// 编辑Excel
func (e *Excel) Insert(data interface{}) error {
	file, err := excelize.OpenFile(e.FileName)
	if err != nil {
		return err
	}
	e.file = file

	rows, err := e.file.GetRows(e.SheetName)
	if err != nil {
		return err
	}
	count := len(rows)

	v := reflect.ValueOf(data)
	t := reflect.TypeOf(data)
	if t.Kind() == reflect.Ptr {
		t = t.Elem()
		v = v.Elem()
	}

	if err := e.store(v.Interface(), 0, count-1); err != nil {
		return err
	}
	return e.file.Save()
}

// 创建Excel (覆盖原文件)
func (e *Excel) Create(data interface{}) error {
	e.file = excelize.NewFile()

	v := reflect.ValueOf(data)
	t := reflect.TypeOf(data)
	if t.Kind() == reflect.Ptr {
		t = t.Elem()
		v = v.Elem()
	}

	e.file.NewSheet(e.SheetName)

	if err := e.setTitle(t, 0, 0); err != nil {
		return err
	}

	if err := e.store(v.Interface(), 0, 0); err != nil {
		return err
	}

	return e.file.SaveAs(e.FileName)
}

// 读取Excel
// data: []struct
func (e *Excel) Read(data interface{}) error {
	file, err := excelize.OpenFile(e.FileName)
	if err != nil {
		return err
	}
	e.file = file
	rows, err := e.file.GetRows(e.SheetName)
	if err != nil {
		return err
	}

	v := reflect.ValueOf(data)
	t := reflect.TypeOf(data)

	if t.Kind() == reflect.Ptr {
		t = t.Elem()
		v = v.Elem()
	}

	if t.Kind() != reflect.Slice {
		return errors.New("can not analysis the type of " + t.Kind().String())
	}

	s := reflect.MakeSlice(t, len(rows)-1, len(rows)-1)

	for i, row := range rows[1:] {
		for k, cell := range row {
			switch s.Index(i).Field(k).Type().Kind() {
			case reflect.Int:
				in, err := strconv.Atoi(cell)
				if err == nil {
					s.Index(i).Field(k).Set(reflect.ValueOf(in))
				}
			case reflect.Uint:
				in, err := strconv.Atoi(cell)
				if err == nil {
					s.Index(i).Field(k).Set(reflect.ValueOf(uint(in)))
				}
			case reflect.Uint16:
				in, err := strconv.Atoi(cell)
				if err == nil {
					s.Index(i).Field(k).Set(reflect.ValueOf(uint16(in)))
				}
			case reflect.Uint32:
				in, err := strconv.Atoi(cell)
				if err == nil {
					s.Index(i).Field(k).Set(reflect.ValueOf(uint32(in)))
				}
			case reflect.Uint64:
				in, err := strconv.Atoi(cell)
				if err == nil {
					s.Index(i).Field(k).Set(reflect.ValueOf(uint64(in)))
				}
			case reflect.Int8:
				in, err := strconv.Atoi(cell)
				if err == nil {
					s.Index(i).Field(k).Set(reflect.ValueOf(int8(in)))
				}
			case reflect.Int16:
				in, err := strconv.Atoi(cell)
				if err == nil {
					s.Index(i).Field(k).Set(reflect.ValueOf(int16(in)))
				}
			case reflect.Int32:
				in, err := strconv.Atoi(cell)
				if err == nil {
					s.Index(i).Field(k).Set(reflect.ValueOf(int32(in)))
				}
			case reflect.Int64:
				in, err := strconv.Atoi(cell)
				if err == nil {
					s.Index(i).Field(k).Set(reflect.ValueOf(int64(in)))
				}
			case reflect.String:
				s.Index(i).Field(k).SetString(cell)
			case reflect.Bool:
				s.Index(i).Field(k).SetBool(cell == "true")
			case reflect.Float32, reflect.Float64:
				float, err := strconv.ParseFloat(cell, 64)
				if err == nil {
					s.Index(i).Field(k).SetFloat(float)
				}
			}

		}
	}

	marshal, err := json.Marshal(s.Interface())
	json.Unmarshal(marshal, data)

	return nil
}

// 保存到excel
// data: Struct,Slice,Ptr{Struct,Slice}
func (e *Excel) store(data interface{}, x, y int) error {
	v := reflect.ValueOf(data)
	t := reflect.TypeOf(data)
	switch t.Kind() {
	case reflect.Struct:
		if err := e.setStruct(v, x, y); err != nil {
			return err
		}
	case reflect.Slice:
		if err := e.setSlice(v, x, y); err != nil {
			return err
		}
	default:
		return errors.New("can not analysis the type of " + t.Kind().String())
	}

	return nil
}

// 设置默认表头
func (e *Excel) setTitle(t reflect.Type, x, y int) error {
	if t.Kind() == reflect.Slice || t.Kind() == reflect.Ptr {
		t = t.Elem()
	}
	for i := 0; i < t.NumField(); i++ {
		if err := e.file.SetCellStr(e.SheetName, Axis(i+1+x, 1+y), t.Field(i).Name); err != nil {
			return err
		}
	}
	return nil
}

// 设置切片类型值
func (e *Excel) setSlice(v reflect.Value, x, y int) error {
	for j := 0; j < v.Len(); j++ {
		for i := 0; i < v.Index(j).NumField(); i++ {
			if err := e.file.SetCellValue(e.SheetName, Axis(i+1+x, j+y+2), v.Index(j).Field(i).Interface()); err != nil {
				return err
			}
		}
	}
	return nil
}

// 保存结构体类型值
func (e *Excel) setStruct(v reflect.Value, x, y int) error {
	for i := 0; i < v.NumField(); i++ {
		if err := e.file.SetCellValue(e.SheetName, Axis(i+1+x, 2+y), v.Field(i).Interface()); err != nil {
			return err
		}
	}
	return nil
}

// excelize.CoordinatesToCellName
// excelize.CellNameToCoordinates

// 获取Excel坐标名
func Axis(i, j int) string {
	b := make([]byte, 0)
	to26(i, &b)
	return string(b) + strconv.Itoa(j)
}

// 索引转Excel行
func to26(i int, b *[]byte) {
	remainder := i % 26
	quotient := i / 26
	if remainder == 0 {
		quotient--
		remainder = 26
	}
	if quotient > 26 {
		to26(remainder, b)
	} else {
		if quotient > 0 {
			*b = append(*b, uint8(quotient)+64)
		}
	}
	*b = append(*b, uint8(remainder)+64)

}
