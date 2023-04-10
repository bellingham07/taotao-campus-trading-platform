package server

import (
	"com.xpwk/go-gin/internal/initial"
	"log"
)

func ListenAndServe(port string) {

	e := initial.Routers()

	err := e.Run(":" + port)
	if err != nil {
		log.Printf("服务启动错误！ error：%s", err.Error())
	}
}
