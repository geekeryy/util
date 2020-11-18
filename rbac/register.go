// @Description  TODO
// @Author  	 jiangyang  
// @Created  	 2020/11/17 6:05 下午
package rbac

import (
	"github.com/comeonjy/util/ctx"
	"github.com/comeonjy/util/errno"
	"github.com/comeonjy/util/jwt"
	"github.com/comeonjy/util/middlewares"
	"github.com/gin-gonic/gin"
	"net/http"
)

const DefaultPrefix = "/rbac"

func getPrefix(prefixOptions ...string) string {
	prefix := DefaultPrefix
	if len(prefixOptions) > 0 {
		prefix = prefixOptions[0]
	}
	return prefix
}

func Register(r *gin.Engine, prefixOptions ...string) {

	prefix := getPrefix(prefixOptions...)

	prefixRouter := r.Group(prefix)
	{
		// 跳转到前端页面
		prefixRouter.GET("", func(ctx *gin.Context) {
			scheme := "http://"
			if ctx.Request.Proto == "HTTP/2" {
				scheme = "https://"
			}
			ctx.Redirect(http.StatusFound, cfg.Frontend+"?referer="+scheme+ctx.Request.Host+ctx.Request.URL.String())
		})
	}

	v1 := prefixRouter.Group("/api/v1")
	{
		// 公共访问组
		public := v1.Group("")
		{
			// 登录
			public.POST("/login", ctx.Handle(Login))
		}

		v1.Use(middlewares.JwtAuth())

		// 权限校验组
		auth := v1.Group("")
		{
			// 用户
			userGroup := auth.Group("/user")
			{
				userGroup.POST("", ctx.Handle(UserInsert))
			}

			// 角色
			roleGroup := auth.Group("/role")
			{
				roleGroup.POST("", ctx.Handle(UserInsert))
			}

		}
	}

}

func Login(ctx *ctx.Context) {
	admin := struct {
		User     string `json:"user"`
		Password string `json:"password"`
	}{}
	if err := ctx.ShouldBindJSON(&admin); err != nil {
		ctx.Fail(errno.ParamErr)
		return
	}

	if admin.User != cfg.User || admin.Password != cfg.Password {
		ctx.Fail(errno.UserPasswordErr)
		return
	}

	token, err := jwt.CreateToken(nil, 0)
	if err != nil {
		ctx.Fail(errno.SystemErr)
		return
	}

	ctx.Success(token)

}
