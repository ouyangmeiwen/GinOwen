package auth

import (
	"time"

	"github.com/dgrijalva/jwt-go"
)

var jwtKey = []byte("jwtowenjwt")

type Claims struct {
	UserID uint `json:"user_id"`
	RoleID uint `json:"role_id"`
	jwt.StandardClaims
}

// 生成JWT令牌
func GenerateJWT(userID uint, roleID uint, expireTime ...int) (string, error) {
	// 设置默认过期时间为24小时
	defaultExpireTime := 24
	if len(expireTime) > 0 {
		defaultExpireTime = expireTime[0]
	}

	expirationTime := time.Now().Add(time.Duration(defaultExpireTime) * time.Hour)
	claims := &Claims{
		UserID: userID,
		RoleID: roleID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtKey)
}

// 解析并验证JWT令牌
func ParseJWT(tokenStr string) (*Claims, error) {
	claims := &Claims{}
	token, err := jwt.ParseWithClaims(tokenStr, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})
	if err != nil || !token.Valid {
		return nil, err
	}
	return claims, nil
}
