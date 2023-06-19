package assist

import (
	"math/rand"
	"strconv"
)

// 不常用工具方法

// GenerateCode 生成验证码
func GenerateCode() string {
	code := ""
	for i := 0; i < 6; i++ {
		code += strconv.Itoa(rand.Intn(10))
	}
	return code
}
