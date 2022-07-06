package main

import "log"

func main() {
	err := server.Start()
	if err != nil {
		log.Fatal(err)
	}
}
