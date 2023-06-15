package middleware

import (
	"go-go-zero/common/utils"
	"net/http"
)

type JwtAuthMiddleware struct {
}

func NewJwtAuthMiddleware() *JwtAuthMiddleware {
	return &JwtAuthMiddleware{}
}

func (m *JwtAuthMiddleware) Handle(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("X-Middleware", "static-middleware")
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			http.Error(w, "访问失败,请登录!😼", http.StatusUnauthorized)
			return
		}
		_, err := utils.ParseToken(authHeader)
		if err != nil {
			http.Error(w, "身份认证错误或过期，请重新登录!", http.StatusUnauthorized)
			return
		}
		if err != nil {
			http.Error(w, "身份认证过期，请重新登录!", http.StatusUnauthorized)
			return
		}
		next(w, r)
	}
}
