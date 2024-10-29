package utils

import (
	"errors"
	"project/global"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

var (
	ErrInvalidToken = errors.New("invalid token")
	ErrTokenExpired = errors.New("token is expired")
)

// Custom claims structure
type CustomClaims struct {
	BaseClaims
	BufferTime int64
	jwt.RegisteredClaims
}

type BaseClaims struct {
	ID     uint
	RoleId uint
}

func GetToken(c *gin.Context) string {
	token := c.Request.Header.Get("x-token")
	return token
}

func SetToken(token string, c *gin.Context) {
	c.Header("x-token", token)
}

func SetRefreshToken(refreshToken string, c *gin.Context) {
	c.Header("x-refresh-token", refreshToken)
}

func ClearToken(c *gin.Context) {
	c.Header("x-token", "")
	c.Header("x-refresh-token", "")
}

// 创建一个token
func NewToken(baseClaims BaseClaims) (string, error) {
	// bf, _ := ParseDuration(global.Config.JWT.BufferTime)
	ep, _ := ParseDuration(global.Config.JWT.ExpiresTime)
	claims := CustomClaims{
		BaseClaims: baseClaims,
		// BufferTime: int64(bf / time.Second), //
		RegisteredClaims: jwt.RegisteredClaims{
			NotBefore: jwt.NewNumericDate(time.Now().Add(-1000)), // 签名生效时间
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(ep)),    // 过期时间
			Issuer:    global.Config.JWT.Issuer,                  // 签名的发行者
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(global.Config.JWT.SigningKey))
}

// 解析 token
func ParseToken(tokenString string) (*CustomClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &CustomClaims{}, func(token *jwt.Token) (i interface{}, e error) {
		return []byte(global.Config.JWT.SigningKey), nil
	})

	if err != nil {
		if errors.Is(err, jwt.ErrTokenExpired) {
			return nil, ErrTokenExpired
		}
		return nil, ErrInvalidToken
	}
	if token.Valid {
		claims, ok := token.Claims.(*CustomClaims)
		if ok {
			return claims, nil
		}
	}

	return nil, ErrInvalidToken
}

func RefreshToken(baseClaims BaseClaims) (string, error) {
	return NewToken(baseClaims)
}

func AutoRefreshToken(claims *CustomClaims, c *gin.Context) {

}

func Auth(ctx *gin.Context) BaseClaims {
	auth, ok := ctx.Get("auth")
	if !ok {
		return BaseClaims{}
	}
	return auth.(BaseClaims)
}
