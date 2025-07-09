package server

import (
	"fmt"
	"net"
	"own-redis/internal/config"
	"strconv"
	"strings"
	"sync"
	"time"
)

type Server struct {
	Conn net.PacketConn
	Storage map[string]Value
	mtx sync.Mutex
}

type Value struct {
	Data string
	ExpiresAt time.Time
}

func NewServer() (*Server, error){
	conn, err := net.ListenPacket("udp", fmt.Sprintf(":%d", config.Port))
	if err != nil {
		return nil, err
	}

	return &Server{
		Conn: conn,
		Storage: make(map[string]Value),
	}, nil
}


func (s *Server) HandleRequest() {
	defer s.Conn.Close()
	for {
		buf := make([]byte, 1024)
		n, addr, err := s.Conn.ReadFrom(buf)
		if err != nil {
			fmt.Println("Error reading request:", err)
			continue
		}
		request := strings.TrimSpace(string(buf[:n]))
		fmt.Println("Received from", addr, ":", request)

		response := s.processCommand(request)
		
		_, err 	= s.Conn.WriteTo([]byte(response), addr)
		if err != nil {
			fmt.Println("Error sending response:", err)
		}
	}
}

func (s *Server) processCommand(req string) string {
	args := strings.Split(req, " ")
	if len(args) == 0 {
		return "(error) empty command\n"
	}

	switch strings.ToLower(args[0]) {
	case "ping":
		return "PONG\n"
	case "set":
		return s.handleSet(args)
	case "get":
		return s.handleGet(args)
	}

	return "(error) something went wrong\n"
}

func (s *Server) handleSet(args []string) string{
	if len(args) < 3 {
		return "(error) ERR wrong number of arguments for 'SET' command\n"
	}
	var val = Value{ExpiresAt: time.Time{}}
	lenArgs := len(args)
	if strings.ToLower(args[lenArgs-2]) == "px" {
		ms, err := strconv.Atoi(args[lenArgs-1])
		if err != nil {
			return fmt.Sprintf("(error) ERR %s\n", err)
		}
		lenArgs = lenArgs-2
		val.ExpiresAt = time.Now().Add(time.Duration(ms)*time.Millisecond)
	}
	val.Data = strings.Join(args[2:lenArgs], " ")
	s.mtx.Lock()	
	s.Storage[args[1]] = val
	s.mtx.Unlock()
	return "OK\n"
}

func (s *Server) handleGet(args []string) string{
	if len(args) != 2 {
		return "(error) ERR wrong number of arguments for 'GET' command\n"
	}
	s.mtx.Lock()
	val, ok := s.Storage[args[1]]
	if !ok || (!val.ExpiresAt.IsZero() && time.Now().After(val.ExpiresAt)) {
		delete(s.Storage, args[1])
		s.mtx.Unlock()
		return "(nil)\n"
	}
	s.mtx.Unlock()
	return val.Data+"\n"
}