package middleware

import (
	"context"
	"fmt"
	"github.com/zeromicro/go-zero/rest/httpx"
	"go-go-zero/common/utils"
	"log"
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
		log.Println(authHeader)
		if authHeader == "" {
			httpx.WriteJson(w, http.StatusUnauthorized, "访问失败,请登录!😼")
			return
		}
		claim, err := utils.ParseToken(authHeader)
		if err != nil {
			httpx.WriteJson(w, http.StatusUnauthorized, "身份认证错误或过期，请重新登录!")
			return
		}
		id := claim.Id
		fmt.Println("asd", id)
		//key := cache.UserLogin + strconv.FormatInt(id, 10)
		//err = cache.RedisUtil.EXPIRE(key, 7*24*time.Hour)
		if err != nil {
			httpx.WriteJson(w, http.StatusUnauthorized, "身份认证过期，请重新登录!")
			return
		}
		k := utils.JwtId("userId")
		var asd int64 = 123
		ctx := context.WithValue(r.Context(), k, asd)
		request := r.WithContext(ctx)
		fmt.Println(request.Context().Value(utils.JwtId("userId")))
		next(w, request)
	}
}
