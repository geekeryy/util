// @Description  TODO
// @Author  	 jiangyang  
// @Created  	 2020/11/16 4:36 下午
package ctx

import (
	"github.com/comeonjy/util/errno"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
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
func (c *Context) Fail(err error) {
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
	c.Context.AbortWithStatusJSON(http.StatusBadRequest, ret)
}

// 包装上下文
func Handle(handle func(*Context)) gin.HandlerFunc {
	return func(c *gin.Context) {
		handle(&Context{c})
	}
}

