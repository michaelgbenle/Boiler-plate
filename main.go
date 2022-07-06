package main

import (
	"github.com/michaelgbenle/Boiler-plate/server"
	"log"
)

func main() {
	err := server.Start()
	if err != nil {
		log.Fatal(err)
	}
}
