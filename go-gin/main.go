package main

import (
	"com.xpdj/go-gin/initial"
	"com.xpdj/go-gin/server"
)

func main() {

	go initial.Initializer()
	server.ListenAndServe("12345")

}
