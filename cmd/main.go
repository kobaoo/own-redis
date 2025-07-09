package main

import (
	"fmt"
	"log"
	"own-redis/internal/config"
	"own-redis/internal/server"
)

func main() {
	err := config.ParseFlags()
	if err != nil {
		log.Fatal("Failed during flag parsing", err)
	}

	srv, err := server.NewServer()
	if err != nil {
		log.Fatal("Error initializing new server")
	}
	fmt.Println("UDP server is running on port:", config.Port)

	srv.HandleRequest()
}
