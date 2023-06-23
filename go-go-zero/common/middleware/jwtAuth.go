package middleware

import (
	"context"
	"fmt"
	"github.com/zeromicro/go-zero/rest/httpx"
	"go-go-zero/common/utils"
	"log"
	"net/http"
)

// JWTAuthenticate jwt校验中间件
func JWTAuthenticate(next http.HandlerFunc) http.HandlerFunc {
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
		fmt.Println(id)
		//key := cache.UserLogin + strconv.FormatInt(id, 10)
		//err = cache.RedisUtil.EXPIRE(key, 7*24*time.Hour)
		if err != nil {
			httpx.WriteJson(w, http.StatusUnauthorized, "身份认证过期，请重新登录!")
			return
		}
		k := JwtId("userId")
		ctx := context.WithValue(r.Context(), k, id)
		next(w, r.WithContext(ctx))
	}
}
