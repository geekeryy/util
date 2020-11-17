// @Description  rbac中间件
// @Author  	 jiangyang  
// @Created  	 2020/11/17 11:32 上午
package middlewares

import (
	"errors"
	"github.com/comeonjy/util/jwt"
	"github.com/comeonjy/util/tool"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Rbac(checkFunc func(interface{}) error) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		bus, exists := ctx.Get("business")
		if !exists {
			ctx.JSON(http.StatusBadRequest, errors.New("business not found"))
			ctx.Abort()
			return
		}

		if checkFunc != nil {
			if err := checkFunc(bus); err != nil {
				ctx.JSON(http.StatusForbidden, err.Error())
				ctx.Abort()
				return
			}
		}

		ctx.Next()
	}
}

// Example:
// 权限校验例子
// bus: ctx中存储的interface类型的业务相关信息
func checkFunc(bus interface{}) error {
	b := jwt.Business{}
	if err := tool.InterfaceToPointer(&b, bus); err != nil {
		return err
	}

	// TODO 权限校验 开箱即用


	return nil
}