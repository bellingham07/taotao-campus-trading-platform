package main

import (
	"com.xpwk/go-gin/initial"
	"com.xpwk/go-gin/server"
	"github.com/yitter/idgenerator-go/idgen"
	"log"
)

func main() {

	go initial.Initializer()
	server.ListenAndServe("12345")
	var newId = idgen.NextId()
	log.Printf("%v", newId)
}
