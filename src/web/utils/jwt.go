package utils

import (
	"family-web-server/src/web/models/eneity/login"
	"github.com/dgrijalva/jwt-go"
	"time"
)

// Claims JWT 结构体
type Claims struct {
	UserId      int
	Username    string
	Role        *login.Role
	Permissions []*login.Permission
	jwt.StandardClaims
}

func GenerateToken(
	UserId int,
	username string,
	role *login.Role,
	permissions []*login.Permission,
	serviceName string,
	expireTime int,
	secretKey string,
) (string, error) {
	// 创建一个声明对象
	claims := Claims{
		UserId:      UserId,
		Username:    username,
		Role:        role,
		Permissions: permissions,
		StandardClaims: jwt.StandardClaims{
			Audience:  serviceName,                                                    // 可以设定为服务名
			ExpiresAt: time.Now().Add(time.Duration(expireTime) * time.Minute).Unix(), // 设置有效期为 24 小时
			IssuedAt:  time.Now().Unix(),                                              // 当前时间作为签发时间
			Issuer:    serviceName,                                                    // 设置签发人
			NotBefore: time.Now().Unix(),                                              // 设置生效时间为当前时间
			Subject:   username,                                                       // 这里可以设置为用户名，或者是其他的用户标识
		},
	}
	// 使用 HMAC 签名算法创建 JWT
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	// 生成 Token 字符串
	return token.SignedString([]byte(secretKey))
}

// ParseToken 解析 JWT Token
func ParseToken(tokenString string, secretKey string) (*Claims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (any, error) {
		return []byte(secretKey), nil
	})

	if err != nil || !token.Valid {
		return nil, err
	}

	claims, ok := token.Claims.(*Claims)
	if !ok {
		return nil, err
	}
	return claims, nil
}
