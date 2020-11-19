// @Description  错误处理
// @Author  	 jiangyang  
// @Created  	 2020/11/17 9:51 上午
package errno

var (
	SystemErr = &Errno{Code: 1, Msg: "系统错误"}

	BusNotFound     = &Errno{101, "Token中业务数据不存在"}
	ParamErr        = &Errno{102, "参数错误"}
	UserPasswordErr = &Errno{103, "账号或密码错误"}

	UserNotFound   = &Errno{1001, "用户不存在"}
	UserRoleErr    = &Errno{1002, "用户角色错误"}
	UserNoPowerErr = &Errno{1003, "用户权限不足"}

	// RBAC系统错误

)

type Errno struct {
	Code int
	Msg  string
}

func (e *Errno) Error() string {
	return e.Msg
}
