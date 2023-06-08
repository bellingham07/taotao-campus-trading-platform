package utils

import (
	"github.com/golang-jwt/jwt/v4"
	"go-go-zero/service/user/model"
	"time"
)

const SECRET_KEY = "xpdj"

type Claims struct {
	Id     int64  `json:"id"`
	Name   string `json:"name"`
	Avatar string `json:"avatar"`
	jwt.RegisteredClaims
}

func GenToken(user *model.UserInfo) (string, error) {
	claim := &Claims{
		Id:     user.Id,
		Name:   user.Name,
		Avatar: user.Avatar,
		RegisteredClaims: jwt.RegisteredClaims{
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(7 * 24 * time.Hour)),
		},
	}
	// 使用指定的签名方法和声明创建一个新token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)
	// 创建并返回一个完整的token（jwt）。令牌使用令牌中指定的签名 方法进行签名 。
	tokenString, err := token.SignedString([]byte(SECRET_KEY))
	return tokenString, err
}

func ParseToken(token string) (*Claims, error) {
	tokenClaim, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(SECRET_KEY), nil
	})
	if tokenClaim != nil {
		if claim, ok := tokenClaim.Claims.(*Claims); ok && tokenClaim.Valid {
			return claim, nil
		}
	}
	return nil, err
}
