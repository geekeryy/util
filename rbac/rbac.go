// @Description  基于角色的权限控制
// @Author  	 jiangyang  
// @Created  	 2020/11/17 11:26 上午

// Example Config:
// rbac:
//   user: admin
//   password: 123456

package rbac

import (
	"github.com/comeonjy/util/errno"
	"github.com/comeonjy/util/jwt"
	"github.com/comeonjy/util/mysql"
	"github.com/comeonjy/util/tool"
	"github.com/jinzhu/gorm"
)

type Config struct {
	User     string `json:"user" yaml:"user"`
	Password string `json:"password" yaml:"password"`
	Frontend string `json:"frontend" yaml:"frontend"`
}


var cfg Config

func Init(c Config) {
	mysql.Conn().AutoMigrate(User{}, Role{}, Power{}, UserRole{}, RolePower{})
	cfg = c
}

func Check(bus interface{}, url string) error {
	b := jwt.Business{}
	if err := tool.InterfaceToPointer(&b, bus); err != nil {
		return err
	}

	return check(b.UID, url)
}

func check(uid uint, url string) error {
	db := mysql.Conn()

	user := User{}
	if err := db.Model(User{}).First(&user, uid).Error; err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return errno.UserNotFound
		}
		return err
	}

	userRoleIDs := make([]uint, 0)
	if err := db.Model(UserRole{}).Where("uid = ?", uid).Pluck("id", &userRoleIDs).Error; err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return errno.UserRoleErr
		}
		return err
	}

	rolePowerIDs := make([]uint, 0)
	if err := db.Model(RolePower{}).Where("rid in (?)", userRoleIDs).Pluck("id", &rolePowerIDs).Error; err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return errno.UserNoPowerErr
		}
		return err
	}

	powers := make([]Power, 0)
	if err := db.Model(Power{}).Where("id in (?)", rolePowerIDs).Find(&powers).Error; err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return errno.UserNoPowerErr
		}
		return err
	}

	for _, v := range powers {
		if v.Url == url {
			return nil
		}
	}

	return errno.UserNoPowerErr
}
