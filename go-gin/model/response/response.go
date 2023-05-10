package response

import "github.com/gin-gonic/gin"

const (
	SUCCESS = "success"
	FAIL    = "error"
	OK      = 1
	ERROR   = -1
)

func genH(code int, msg string, data ...any) gin.H {
	return gin.H{"code": code, "msg": msg, "data": data}
}

func Ok() gin.H {
	return genH(OK, SUCCESS)
}

func Error() gin.H {
	return genH(ERROR, FAIL)
}

func OkData(data any) gin.H {
	return genH(OK, SUCCESS, data)
}

func ErrorData(data any) gin.H {
	return genH(OK, FAIL, data)
}

func OkMsg(msg string) gin.H {
	return genH(OK, msg)
}

func ErrorMsg(msg string) gin.H {
	return genH(ERROR, msg)
}

func OkMsgData(msg string, data any) gin.H {
	return genH(OK, msg, data)
}

func ErrorMsgData(msg string, data any) gin.H {
	return genH(ERROR, msg, data)
}
