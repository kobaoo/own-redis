package server

import (
	"fmt"
	"net"
	"own-redis/internal/config"
	"time"
)

type Server struct {
	conn net.PacketConn
	storage map[string]Value
}

type Value struct {
	data string
	expiresAt time.Time
}

func NewServer() (*Server, error){
	conn, err := net.ListenPacket("udp", fmt.Sprintf(":%d", config.Port))
	if err != nil {
		return nil, err
	}
	defer conn.Close()

	return &Server{
		conn: conn,
		storage: make(map[string]Value),
	}, nil
}