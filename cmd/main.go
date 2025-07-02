package main

import (
	"log"
	"own-redis/internal/config"
)

func main() {
	err := config.ParseFlags()
	if err != nil {
		log.Fatal("Failed during flag parsing", err)
	}
}
