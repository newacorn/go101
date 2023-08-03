package main

import (
	"log"
	"runtime"

	"helloworld"
)

func main() {
	log.Println(runtime.GOOS)
	helloworld.SayHello()
}
