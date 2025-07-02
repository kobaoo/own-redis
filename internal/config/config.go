package config

import (
	"flag"
	"fmt"
)

var Port int

func ParseFlags() error {
	flag.IntVar(&Port, "port", 8080, "port to serve on")
	flag.Parse()
	flag.Usage = func() {
		printHelp()
	}

	if Port < 1024 || Port > 65565 {
		return fmt.Errorf("port number must be between 1024 and 65565")
	}
	return nil
}

func printHelp() {
	fmt.Println(`$ ./own-redis --help
Own Redis

Usage:
  own-redis [--port <N>]
  own-redis --help

Options:
  --help       Show this screen.
  --port N     Port number.`)
}
