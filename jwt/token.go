package jwt

import (
	"time"

	jwtgo "github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

// DefaultExpireDuration .
const DefaultExpireDuration = time.Hour * 24 * 30

// TokenReq struct.
type TokenReq struct {
	UID       uint
	Role      uint8
	SchoolID  uint
	SubjectID uint
	Expire    time.Duration
}

// TokenResp struct.
type TokenResp struct {
	Token     string `json:"token"`
	ExpiredAt string `json:"expired_at"`
}

// GenerateToken 生成Token.
func GenerateToken(c *gin.Context, token TokenReq) (*TokenResp, error) {
	j := &JWT{
		[]byte(SignKey),
	}

	if token.Expire == 0 {
		token.Expire = DefaultExpireDuration
	}

	expireSec := int64(token.Expire.Seconds())
	nowTS := int64(time.Now().Unix())
	expireTS := int64(nowTS + expireSec)
	claims := Customclaims{
		token.UID,
		token.Role,
		token.SchoolID,
		token.SubjectID,
		jwtgo.StandardClaims{
			NotBefore: nowTS,    //签名生效时间
			ExpiresAt: expireTS, //签名过期时间
			Issuer:    SignKey,  //签名发行者
		},
	}

	tokenStr, err := j.CreateToken(claims)
	if err != nil {
		return nil, err
	}
	return &TokenResp{
		tokenStr,
		time.Unix(expireTS, 0).Format("2006-01-02 15:04:05"),
	}, nil
}

// 获取token信息
func GetTokenInfo(ctx *gin.Context) (*Customclaims, error) {
	claims, exists := ctx.Get("token")
	if !exists {
		return nil, ErrTokenInvalid
	}
	if tokenInfo, ok := claims.(*Customclaims); ok {
		return tokenInfo, nil
	}
	return nil, ErrTokenInvalid
}
