package utils

import (
	"github.com/golang-jwt/jwt/v4"
	"go-go-zero/service/user/model"
	"time"
)

type JwtId string

const SecretKey = "xpdj"

type Claims struct {
	Id   int64  `json:"id"`
	Name string `json:"name"`
	jwt.RegisteredClaims
}

func GenToken(user *model.UserInfo) (string, error) {
	var now = time.Now().Local()
	var claim = &Claims{
		Id:   user.Id,
		Name: user.Name,
		RegisteredClaims: jwt.RegisteredClaims{
			IssuedAt:  jwt.NewNumericDate(now),
			ExpiresAt: jwt.NewNumericDate(now.Add(7 * 24 * time.Hour)),
		},
	}
	// 使用指定的签名方法和声明创建一个新token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)
	// 创建并返回一个完整的token（jwt）。令牌使用令牌中指定的签名 方法进行签名 。
	tokenString, err := token.SignedString([]byte(SecretKey))
	return tokenString, err
}

func ParseToken(token string) (*Claims, error) {
	tokenClaim, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(SecretKey), nil
	})
	if tokenClaim.Valid {
		if claim, ok := tokenClaim.Claims.(*Claims); ok {
			return claim, nil
		}
	}
	return nil, err
}
