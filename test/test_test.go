// @Description  TODO
// @Author  	 jiangyang  
// @Created  	 2021/3/17 10:33 上午
package test_test

import (
	"fmt"
	"regexp"
	"testing"
)

func foo() (a int) {
	p:=&a
	defer func() {
		a++
		*p++
	}()
	a++
	return a+1
}

func TestDemo(t *testing.T)  {
	re, err:= regexp.Compile(".*网申.*")
	if err!=nil{
		t.Error(err)
	}
	fmt.Println(re.MatchString("qwe网1申wqe"))

}
