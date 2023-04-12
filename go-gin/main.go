package main

import (
	"com.xpwk/go-gin/initial"
	"com.xpwk/go-gin/server"
)

func main() {

	go initial.Initializer()
	server.ListenAndServe("12345")

}
