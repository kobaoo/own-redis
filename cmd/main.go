package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"own-redis/internal/config"
)

func main() {
	err := config.ParseFlags()
	if err != nil {
		log.Fatal("Failed during flag parsing", err)
	}

	p := make([]byte, 2048)
	conn, err := net.Dial("upd", fmt.Sprintf("127.0.0.1:%d", config.Port))
	if err != nil {
		log.Fatal("Failed to connect with udp")
	}
	fmt.Fprintf(conn, "HELLOO IT WORKS")
	_, err = bufio.NewReader(conn).Read(p)
	if err == nil {
		fmt.Printf("%s\n", p)
	} else {
		fmt.Printf("Some error %v\n", err)
	}
	conn.Close()
}
