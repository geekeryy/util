// @Description  jwt中间件
// @Author  	 jiangyang  
// @Created  	 2020/11/16 5:12 下午
package middlewares

import (
	"net/http"

	"github.com/comeonjy/util/jwt"

	"github.com/gin-gonic/gin"
)

func JwtAuth() func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		token := ctx.GetHeader("Authorization")
		if len(token) == 0 {
			ctx.JSON(http.StatusForbidden, "Auth Forbidden")
			ctx.Abort()
			return
		}
		if bus, err := jwt.ParseToken(token); err != nil {
			ctx.JSON(http.StatusForbidden, err.Error())
			ctx.Abort()
			return
		} else {
			ctx.Set("business", bus)
		}

		ctx.Next()
	}
}
