// @Description  rbac中间件
// @Author  	 jiangyang  
// @Created  	 2020/11/17 11:32 上午
package middlewares

import (
	"net/http"

	"github.com/gin-gonic/gin"

	core "github.com/comeonjy/util/ctx"
	"github.com/comeonjy/util/errno"
	"github.com/comeonjy/util/jwt"
	"github.com/comeonjy/util/tool"
)

func Rbac(checkFunc func(interface{}, string) error) func(context *gin.Context) {
	return func(context *gin.Context) {
		ctx := core.Context{
			Context: context,
		}
		bus, exists := ctx.Get("business")
		if !exists {
			ctx.Fail(errno.BusNotFound)
			return
		}

		if checkFunc != nil {
			if err := checkFunc(bus, ctx.Request.URL.String()); err != nil {
				ctx.Fail(err, http.StatusForbidden)
				return
			}
		}

		ctx.Next()
	}
}

// Example:
// 权限校验例子
// bus: ctx中存储的interface类型的业务相关信息
func checkFunc(bus interface{}, url string) error {
	b := jwt.Business{}
	if err := tool.InterfaceToPointer(&b, bus); err != nil {
		return err
	}

	// TODO 权限校验 开箱即用

	return nil
}
