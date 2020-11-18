// @Description  TODO
// @Author  	 jiangyang  
// @Created  	 2020/11/17 11:26 上午
package rbac

import "github.com/jinzhu/gorm"

type User struct {
	gorm.Model
	Name string `json:"name"`
}

type Role struct {
	gorm.Model
	Name string `json:"name"`
}

type Power struct {
	gorm.Model
	Title string `json:"title"`
	Url   string `json:"url"`
}

type UserRole struct {
	gorm.Model
	UID uint `json:"uid"`
	RID uint `json:"rid"`
}

type RolePower struct {
	gorm.Model
	RID     uint `json:"rid"`
	PowerID uint `json:"power_id"`
}
