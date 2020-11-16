// @Description  TODO
// @Author  	 jiangyang  
// @Created  	 2020/11/16 4:36 下午
package ctx

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
)

type Context struct {
	*gin.Context
}

type Errno struct {
	Code int
	Msg  string
}

func (e *Errno) Error() string {
	return e.Msg
}

func (c *Context) Success(data interface{}) {
	c.Context.JSON(http.StatusOK, gin.H{
		"code": 0,
		"msg":  "ok",
		"data": data,
	})
}

func (c *Context) Fail(err error) {
	logrus.Error(err)
	ret := gin.H{
		"code": -1,
		"msg":  "未知错误",
	}
	if errno, ok := err.(*Errno); ok {
		ret["code"] = errno.Code
		ret["msg"] = errno.Msg
	}
	c.Context.AbortWithStatusJSON(http.StatusBadRequest, ret)
}

func Handle(handle func(*Context)) gin.HandlerFunc {
	return func(c *gin.Context) {
		handle(&Context{c})
	}
}
