package main

import (
	"com.xpwk/go-gin/initial"
	"com.xpwk/go-gin/server"
	"github.com/yitter/idgenerator-go/idgen"
	"log"
)

func main() {

	go initial.Initializer()
	var newId = idgen.NextId()
	log.Printf("acsdfdsgadfgfdg%v", newId)
	server.ListenAndServe("12345")

}
