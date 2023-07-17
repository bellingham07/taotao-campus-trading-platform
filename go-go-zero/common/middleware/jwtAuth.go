package middleware

import (
	"context"
	"fmt"
	"github.com/zeromicro/go-zero/rest/httpx"
	"go-go-zero/common/utils"
	"net/http"
	"strconv"
)

// JwtAuthenticate jwt校验中间件
func JwtAuthenticate(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("X-Middleware", "static-middleware")

		headerToken := r.Header.Get("Authorization")

		fmt.Println(headerToken)
		if headerToken == "" {
			httpx.WriteJson(w, http.StatusUnauthorized, "请先登录!😼")
			return
		}

		claim, err := utils.ParseToken(headerToken)
		if err != nil {
			httpx.WriteJson(w, http.StatusUnauthorized, "身份认证错误或过期，请重新登录!")
			return
		}

		id := claim.Id
		key := utils.UserLogin + strconv.FormatInt(id, 10)

		redisToken, err := utils.UserServiceRedis.Get(r.Context(), key).Result()
		if redisToken != headerToken {
			httpx.WriteJson(w, http.StatusUnauthorized, "身份认证过期，请重新登录!")
			return
		}

		ctx := context.WithValue(r.Context(), utils.JwtId("userId"), id)
		ctx = context.WithValue(ctx, utils.JwtName("name"), claim.Name)
		request := r.WithContext(ctx)
		next(w, request)
	}
}
