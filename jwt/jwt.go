package jwt

import (
	"github.com/pkg/errors"

	jwtgo "github.com/dgrijalva/jwt-go"
)

// JWT struct.
type JWT struct {
	SigningKey []byte
}

var (
	// ErrTokenExpired .
	ErrTokenExpired = errors.New("Token is expired")
	// ErrTokenNotValidYet .
	ErrTokenNotValidYet = errors.New("Token not active yet")
	// ErrTokenMalformed .
	ErrTokenMalformed = errors.New("That's not even a token")
	// ErrTokenInvalid .
	ErrTokenInvalid = errors.New("Couldn't handle this token")
	// SignKey .
	SignKey = "class100"
)

// Customclaims struct.
type Customclaims struct {
	UID       uint  `json:"uid"`
	Role      uint8 `json:"role"`
	SchoolID  uint  `json:"school_id"`
	SubjectID uint  `json:"subject_id"`
	jwtgo.StandardClaims
}

// SetSignKey method
func SetSignKey(key string) string {
	SignKey = key
	return SignKey
}

// GetSignKey method
func GetSignKey() string {
	return SignKey
}

// NewJWT method return a JWT instance
func NewJWT() *JWT {
	return &JWT{
		[]byte(GetSignKey()),
	}
}

// CreateToken method
func (j *JWT) CreateToken(claims Customclaims) (string, error) {
	token := jwtgo.NewWithClaims(jwtgo.SigningMethodHS256, claims)
	return token.SignedString(j.SigningKey)
}

// ParseToken method
func (j *JWT) ParseToken(tokenString string) (*Customclaims, error) {
	token, err := jwtgo.ParseWithClaims(tokenString, &Customclaims{}, func(token *jwtgo.Token) (interface{}, error) {
		return j.SigningKey, nil
	})
	if err != nil {
		if ve, ok := err.(*jwtgo.ValidationError); ok {
			if ve.Errors&jwtgo.ValidationErrorMalformed != 0 {
				return nil, ErrTokenMalformed
			} else if ve.Errors&jwtgo.ValidationErrorExpired != 0 {
				// Token is expired
				return nil, ErrTokenExpired
			} else if ve.Errors&jwtgo.ValidationErrorNotValidYet != 0 {
				return nil, ErrTokenNotValidYet
			} else {
				return nil, ErrTokenInvalid
			}
		}
	}
	if claims, ok := token.Claims.(*Customclaims); ok && token.Valid {
		return claims, nil
	}
	return nil, ErrTokenInvalid
}
