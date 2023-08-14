package main

import (
	"com.xpdj/go-gin/server"
)

func main() {

	server.Initializer()
	server.ListenAndServe("12345")

}
