package server

import (
	"com.xpdj/go-gin/router"
	"log"
)

func ListenAndServe(port string) {

	e := router.Routers()

	err := e.Run(":" + port)
	if err != nil {
		log.Printf("服务启动错误！ error：%s", err.Error())
	}
}
