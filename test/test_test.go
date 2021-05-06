// @Description  TODO
// @Author  	 jiangyang  
// @Created  	 2021/3/17 10:33 上午
package test_test

import (
	"fmt"
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
	fmt.Println(foo())
}
