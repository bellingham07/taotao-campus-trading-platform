package main

import (
	"com.xpwk/go-gin/internal/initial"
	"com.xpwk/go-gin/internal/server"
)

func main() {

	go initial.Initializer()
	server.ListenAndServe("12345")

}
