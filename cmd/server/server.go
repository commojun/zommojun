package main

import (
	"log"

	"github.com/commojun/zommojun/server"
)

func main() {
	err := server.Start()
	if err != nil {
		log.Fatal(err)
	}

}
