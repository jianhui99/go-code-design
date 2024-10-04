package main

import (
	"fmt"
	"time"
)

type Server struct {
	Addr         string
	Port         int
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
	Timeout      time.Duration
}

func NewServer(addr string, port int, readTimeout, writeTimeout, timeout time.Duration) *Server {
	return &Server{
		Addr:         addr,
		Port:         port,
		ReadTimeout:  readTimeout,
		WriteTimeout: writeTimeout,
		Timeout:      timeout,
	}
}

func main() {
	server := NewServer("localhost", 8080, 2*time.Second, 2*time.Second, 4*time.Second)
	fmt.Println(server)
}
