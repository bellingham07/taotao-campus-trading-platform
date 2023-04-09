package server

import (
	"com.xpwk/go-gin/internal/init"
	"log"
)

func ListenAndServe(port string) {

	e := init.Routers()

	err := e.Run(":" + port)
	if err != nil {
		log.Printf("服务启动错误！ error：%s", err.Error())
	}
}
