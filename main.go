package main

import (
	"log"

	"github.com/phatwasin01/guessing-game-api/api"
)

func main() {
	server, err := api.NewServer()
	if err != nil {
		log.Fatal("Cannot Setup Server:", err)
	}
	err = server.Start()
	if err != nil {
		log.Fatal("Cannot Start Server:", err)
	}

}
