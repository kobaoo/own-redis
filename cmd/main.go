package main

import (
	"log"
	"own-redis/internal/config"
	"own-redis/internal/server"
)

func main() {
	err := config.ParseFlags()
	if err != nil {
		log.Fatal("Failed during flag parsing", err)
	}

	server, err := server.NewServer()
	if err != nil {
		log.Fatal("Error initializing new server")
	}

	
}
