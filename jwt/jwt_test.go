// @Description  TODO
// @Author  	 jiangyang  
// @Created  	 2020/11/16 3:09 下午
package jwt_test

import (
	"github.com/comeonjy/util/jwt"
	"log"
	"testing"
)

func TestCreateToken(t *testing.T) {
	token, err := jwt.CreateToken(&jwt.Business{
		UID:       1,
		Role:      1,
	},0)
	log.Println(token, err)

	bus := &jwt.Business{}
	if err := jwt.ParseToken(token.Token, bus); err != nil {
		t.Error(err)
	}
	log.Println(err, bus)
}
