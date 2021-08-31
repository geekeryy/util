// @Description  TODO
// @Author  	 jiangyang  
// @Created  	 2020/11/16 3:09 下午
package jwt_test

import (
	"log"
	"testing"

	"github.com/comeonjy/util/jwt"
)

func TestCreateToken(t *testing.T) {
	token, err := jwt.CreateToken(jwt.Business{
		UID:  1,
		Role: 2,
	}, 0)
	log.Println(token, err)

	if bus, err := jwt.ParseToken(token.Token); err != nil {
		t.Errorf("%+v",err)
	} else {

		if b, ok := bus.(map[string]interface{}); ok {
			log.Println(b)
		}

		log.Println(bus)
	}

}