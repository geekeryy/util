// @Description  TODO
// @Author  	 jiangyang  
// @Created  	 2020/11/17 9:51 上午
package errno

var (
	SYSTEM_ERR = Errno{Code: 1,Msg: ""}
)

type Errno struct {
	Code int
	Msg  string
}

func (e *Errno) Error() string {
	return e.Msg
}
