package response

import "github.com/gin-gonic/gin"

const (
	SUCCESS = "success"
	ERROR   = "error"
	OK      = 1
	FAIL    = -1
)

func GenH(code int, msg string, data ...any) gin.H {
	return gin.H{"code": code, "msg": msg, "data": data}
}
