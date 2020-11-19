// @Description  应用程序上下文
// @Author  	 jiangyang  
// @Created  	 2020/11/16 4:36 下午
package ctx

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"

	"github.com/comeonjy/util/errno"
)

type Context struct {
	*gin.Context
}

// 成功返回
func (c *Context) Success(data interface{}) {
	c.Context.JSON(http.StatusOK, gin.H{
		"code": 0,
		"msg":  "ok",
		"data": data,
	})
}

// 错误返回
func (c *Context) Fail(err error, statusArr ...int) {
	status := http.StatusBadRequest
	if len(statusArr) > 0 {
		status = statusArr[0]
	}
	logrus.Error(err)
	ret := gin.H{}
	if e, ok := err.(*errno.Errno); ok {
		ret["code"] = e.Code
		ret["msg"] = e.Msg
	} else {
		ret["code"] = -1
		ret["msg"] = "未知错误"
		ret["err"] = err
	}
	c.Context.AbortWithStatusJSON(status, ret)
}

// 包装上下文
func Handle(handle func(*Context)) gin.HandlerFunc {
	return func(c *gin.Context) {
		handle(&Context{c})
	}
}
