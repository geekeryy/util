package jwt

import (
	"encoding/json"
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
	UID       uint          `json:"uid"`
	Role      uint8         `json:"role"`
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
func CreateToken(bus interface{},expires int64) (*TokenResp, error) {
	expiresAt := time.Now().Add(DefaultExpireDuration).Unix()
	if expires != 0 {
		expiresAt = time.Now().Add(time.Duration(expires)).Unix()
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
func ParseToken(tokenString string, bus interface{}) error {
	token, err := jwtgo.ParseWithClaims(tokenString, &CustomClaims{}, func(token *jwtgo.Token) (interface{}, error) {
		return SignKey, nil
	})

	if err != nil {
		if ve, ok := err.(*jwtgo.ValidationError); ok {
			if ve.Errors&jwtgo.ValidationErrorMalformed != 0 {
				return ErrTokenMalformed
			} else if ve.Errors&jwtgo.ValidationErrorExpired != 0 {
				return ErrTokenExpired
			} else if ve.Errors&jwtgo.ValidationErrorNotValidYet != 0 {
				return ErrTokenNotValidYet
			} else {
				return ErrTokenInvalid
			}
		}
	}

	if token != nil {
		if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid {
			marshal, err := json.Marshal(claims.Business)
			if err != nil {
				return err
			}
			if err := json.Unmarshal(marshal, bus); err != nil {
				return err
			}
			return nil
		}
	}

	return ErrTokenInvalid
}
