// Package excel
// @Description  TODO 不够好
// @Author  	 jiangyang  
// @Created  	 2020/12/7 11:10 上午
package excel

import (
	"encoding/json"
	"fmt"
	"github.com/360EntSecGroup-Skylar/excelize/v2"
	"github.com/pkg/errors"
	"os"
	"reflect"
	"strconv"
)

const (
	StructTag = "excel"
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

// Save 保存
func (e *Excel) Save(data interface{}) error {
	_, err := os.Stat(e.FileName)
	if err != nil {
		if os.IsNotExist(err) {
			return e.Create(data)
		}
	}
	return e.Insert(data)
}

func (e *Excel) Check() error {

	return nil
}

// Insert 编辑Excel
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

// Create 创建Excel (覆盖原文件)
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

// 读取Excel 第一行为表头
// data: []struct
func (e *Excel) Read(data interface{}) error {
	var err error

	v := reflect.ValueOf(data)
	t := reflect.TypeOf(data)

	if t.Kind() == reflect.Ptr {
		t = t.Elem()
		v = v.Elem()
	}

	if e.file, err = excelize.OpenFile(e.FileName); err != nil {
		return err
	}

	rows, err := e.file.GetRows(e.SheetName)
	if err != nil {
		return err
	}

	if t.Elem().NumField() < len(rows[0]) {
		return errors.New("invalid data can not match field num")
	}

	x := 0
	for i, rowTag := range rows[0] {
		var tag string
		for ; i+x < t.Elem().NumField(); x++ {
			tag = t.Elem().Field(i + x).Tag.Get(StructTag)
			if tag != "-" {
				if tag != rowTag {
					return errors.New("field " + tag + " not find")
				} else {
					break
				}
			}

		}
	}

	if t.Kind() != reflect.Slice {
		return errors.New("can not analysis the type of " + t.Kind().String())
	}

	s := reflect.MakeSlice(t, len(rows)-1, len(rows)-1)

	for i, row := range rows[1:] {
		index := 0
		for k, cell := range row {
			if t.Kind() == reflect.Slice {
				t = t.Elem()
			}

			for ; k+index < t.NumField(); index++ {
				if t.Field(k+index).Tag.Get(StructTag) != "-" {
					break
				}
			}
			if k+index >= t.NumField() {
				break
			}

			field := s.Index(i).Field(k + index)
			switch field.Type().Kind() {
			case reflect.Bool:
				field.SetBool(cell == "true" || cell == "1" || cell == "是")
			case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
				in, err := strconv.ParseInt(cell, 10, 64)
				if err == nil {
					field.SetInt(in)
				}
			case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
				in, err := strconv.ParseUint(cell, 10, 64)
				if err == nil {
					field.SetUint(in)
				}
			case reflect.String:
				field.SetString(cell)
			case reflect.Float32, reflect.Float64:
				float, err := strconv.ParseFloat(cell, 64)
				if err == nil {
					field.SetFloat(float)
				}
			default:
				return errors.New(fmt.Sprintf("not supported kind of %s", field.Type().Kind()))
			}

		}
	}

	marshal, _ := json.Marshal(s.Interface())
	_ = json.Unmarshal(marshal, data)

	return nil
}

// 保存到excel
// data: Struct,Slice,Ptr{Struct,Slice}
// x,y: 偏移量
func (e *Excel) store(data interface{}, x, y int) error {
	t := reflect.TypeOf(data)
	switch t.Kind() {
	case reflect.Struct:
		if err := e.setStruct(data, x, y); err != nil {
			return err
		}
	case reflect.Slice:
		if err := e.setSlice(data, x, y); err != nil {
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
		tag := t.Field(i).Tag.Get(StructTag)
		if tag == "-" {
			x--
		} else {
			if err := e.file.SetCellStr(e.SheetName, Axis(i+1+x, 1+y), tag); err != nil {
				return err
			}
		}
	}
	return nil
}

// 设置切片类型值
func (e *Excel) setSlice(data interface{}, x, y int) error {
	v := reflect.ValueOf(data)
	t := reflect.TypeOf(data)
	for j := 0; j < v.Len(); j++ {
		index := 0
		for i := 0; i < v.Index(j).NumField(); i++ {
			if t.Elem().Field(i).Tag.Get(StructTag) == "-" {
				index--
			} else {
				if err := e.file.SetCellValue(e.SheetName, Axis(i+1+x+index, j+y+2), v.Index(j).Field(i).Interface()); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

// 保存结构体类型值
func (e *Excel) setStruct(data interface{}, x, y int) error {
	v := reflect.ValueOf(data)
	t := reflect.TypeOf(data)
	for i := 0; i < v.NumField(); i++ {
		if t.Field(i).Tag.Get(StructTag) == "-" {
			x--
		} else {
			if err := e.file.SetCellValue(e.SheetName, Axis(i+1+x, 2+y), v.Field(i).Interface()); err != nil {
				return err
			}
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
