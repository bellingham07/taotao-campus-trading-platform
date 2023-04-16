package middleware

import (
	"com.xpwk/go-gin/cache"
	"com.xpwk/go-gin/model/response"
	"com.xpwk/go-gin/utils"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"time"
)

func JWTAuthenticate() func(c *gin.Context) {
	return func(c *gin.Context) {
		//获取到请求头中的token
		authHeader := c.Request.Header.Get("Authorization")
		log.Println(authHeader)
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{
				"code": response.FAIL,
				"msg":  "访问失败,请登录!",
			})
			c.Abort()
			return
		}
		claim, err := utils.ParseToken(authHeader)
		if err != nil {
			c.JSON(http.StatusOK, gin.H{
				"code": response.FAIL,
				"msg":  "身份认证错误，请重新登录!",
			})
			c.Abort()
			return
		}
		id := claim.Id
		key := cache.USERLOGIN + id
		err = cache.RedisClient.Expire(key, 7*24*time.Hour)
		if err != nil {
			log.Println(err.Error())
			c.JSON(http.StatusOK, gin.H{
				"code": response.FAIL,
				"msg":  "身份信息过期,请重新登录!",
			})
			c.Abort()
			return
		}

		// 将当前请求的userID信息保存到请求的上下文c上
		c.Set("userid", id)
		c.Next() // 后续的处理函数可以用过c.Get("username")来获取当前请求的用户信息
	}
}