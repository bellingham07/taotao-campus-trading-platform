package utils

import "github.com/golang-jwt/jwt/v4"

func NewWithClaims(method jwt.SigningMethod, claims jwt.Claims) *jwt.Token {

	return &jwt.Token{
		Header: map[string]interface{}{
			"typ": "JWT",
			"alg": method.Alg(),
		},
		Claims: claims,
		Method: method,
	}
}
