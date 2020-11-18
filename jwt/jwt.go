// @Description  jwt
// @Author  	 jiangyang
// @Created  	 2020/11/17 4:12 下午
package jwt

import (
	"github.com/pkg/errors"
	"time"

	jwtgo "github.com/dgrijalva/jwt-go"
)

const DefaultExpireDuration = time.Hour * 24 * 30

var (
	ErrTokenExpired     = errors.New("Token is expired")
	ErrTokenNotValidYet = errors.New("Token not active yet")
	ErrTokenMalformed   = errors.New("That's not even a token")
	ErrTokenInvalid     = errors.New("Couldn't handle this token")
	SignKey             = []byte("243223ffslsfsldfl412fdsfsdf")
)

type Business struct {
	UID  uint `json:"uid"`
	Role uint `json:"role"`
}

type CustomClaims struct {
	Business interface{}
	jwtgo.StandardClaims
}

type TokenResp struct {
	Token     string `json:"token"`
	ExpiredAt string `json:"expired_at"`
}

func Init(key string) {
	SignKey = []byte(key)
}

// 创建Token
func CreateToken(bus interface{}, expires time.Duration) (*TokenResp, error) {
	expiresAt := time.Now().Add(DefaultExpireDuration).Unix()
	if expires != 0 {
		expiresAt = time.Now().Add(expires).Unix()
	}
	claims := &CustomClaims{
		Business: bus,
		StandardClaims: jwtgo.StandardClaims{
			ExpiresAt: expiresAt,
		},
	}
	token := jwtgo.NewWithClaims(jwtgo.SigningMethodHS256, claims)
	tokenStr, err := token.SignedString(SignKey)
	if err != nil {
		return nil, err
	}
	return &TokenResp{
		Token:     tokenStr,
		ExpiredAt: time.Unix(expiresAt, 0).Format("2006-01-02 15:04:05"),
	}, nil
}

// 解析Token
func ParseToken(tokenString string) (interface{}, error) {
	customClaims := CustomClaims{}
	token, err := jwtgo.ParseWithClaims(tokenString, &customClaims, func(token *jwtgo.Token) (interface{}, error) {
		return SignKey, nil
	})

	if err != nil {
		if ve, ok := err.(*jwtgo.ValidationError); ok {
			if ve.Errors&jwtgo.ValidationErrorMalformed != 0 {
				return nil, ErrTokenMalformed
			} else if ve.Errors&jwtgo.ValidationErrorExpired != 0 {
				return nil, ErrTokenExpired
			} else if ve.Errors&jwtgo.ValidationErrorNotValidYet != 0 {
				return nil, ErrTokenNotValidYet
			} else {
				return nil, ErrTokenInvalid
			}
		}
	}

	if token == nil || !token.Valid {
		return nil, ErrTokenInvalid
	}
	return customClaims.Business, nil

}
