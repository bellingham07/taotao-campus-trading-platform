package middleware

import (
	"com.xpdj/go-gin/model/response"
	"com.xpdj/go-gin/utils"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
	"time"
)

func JWTAuthenticate(c *gin.Context) {
	//获取到请求头中的token
	authHeader := c.Request.Header.Get("Authorization")
	log.Println(authHeader)
	if authHeader == "" {
		c.JSON(http.StatusUnauthorized, response.GenH(response.FAIL, "访问失败,请登录!"))
		c.Abort()
		return
	}
	claim, err := utils.ParseToken(authHeader)
	if err != nil {
		c.JSON(http.StatusOK, response.GenH(response.FAIL, "身份认证错误或过期，请重新登录!"))
		c.Abort()
		return
	}
	id := claim.Id
	key := utils.USERLOGIN + id
	err = utils.RedisUtil.EXPIRE(key, 7*24*time.Hour)
	if err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusOK, response.GenH(response.FAIL, "身份信息过期,请重新登录!"))
		c.Abort()
		return
	}
	// 将当前请求的userID信息保存到请求的上下文c上
	c.Set("userid", id)
	c.Next() // 后续的处理函数可以用过c.GET("username")来获取当前请求的用户信息
}

func GetUserIdStr(c *gin.Context) string {
	userIdAny, _ := c.Get("userid")
	userIdStr := userIdAny.(string)
	return userIdStr
}

func GetUserId(c *gin.Context) int64 {
	userIdAny, _ := c.Get("userid")
	userIdStr := userIdAny.(string)
	userId, _ := strconv.ParseInt(userIdStr, 10, 64)
	return userId
}
