// @Description  TODO
// @Author  	 jiangyang  
// @Created  	 2020/11/17 11:26 上午
package rbac

type User struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
}

type Role struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
}

type Power struct {
	ID    uint   `json:"id"`
	Title string `json:"title"`
}

type UserRole struct {
	ID     uint `json:"id"`
	UID    uint `json:"uid"`
	RoleID uint `json:"role_id"`
}

type RolePower struct {
	ID      uint `json:"id"`
	RID     uint `json:"rid"`
	PowerID uint `json:"power_id"`
}
