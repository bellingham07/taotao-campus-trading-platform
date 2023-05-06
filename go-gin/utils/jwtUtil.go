package utils

import (
	"com.xpdj/go-gin/model"
	"github.com/golang-jwt/jwt/v4"
	"strconv"
	"time"
)

const SECRETKEY = "xpdj"

type Claims struct {
	Id   string `json:"id"`
	Name string `json:"name"`
	//Avatar string `json:"avatar"`
	jwt.RegisteredClaims
}

func GenerateToken(user *model.UserInfo) (string, error) {
	claim := &Claims{
		Id:   strconv.FormatInt(user.Id, 10),
		Name: user.Name,
		//Avatar: user.Avatar,
		RegisteredClaims: jwt.RegisteredClaims{
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(7 * 24 * time.Hour)),
		},
	}
	// 使用指定的签名方法和声明创建一个新token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)
	// 创建并返回一个完整的token（jwt）。令牌使用令牌中指定的签名 方法进行签名 。
	tokenString, err := token.SignedString([]byte(SECRETKEY))
	return tokenString, err
}

func ParseToken(token string) (*Claims, error) {
	tokenClaim, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(SECRETKEY), nil
	})
	if tokenClaim != nil {
		if claim, ok := tokenClaim.Claims.(*Claims); ok && tokenClaim.Valid {
			return claim, nil
		}
	}
	return nil, err
}
