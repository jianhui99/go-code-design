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

type Option func(*Server)

func WithAddr(addr string) Option {
	return func(s *Server) {
		s.Addr = addr
	}
}

func WithPort(port int) Option {
	return func(s *Server) {
		s.Port = port
	}
}

func WithReadTimeout(readTimeout time.Duration) Option {
	return func(s *Server) {
		s.ReadTimeout = readTimeout
	}
}

func WithWriteTimeout(writeTimeout time.Duration) Option {
	return func(s *Server) {
		s.WriteTimeout = writeTimeout
	}
}

func WithTimeout(timeout time.Duration) Option {
	return func(s *Server) {
		s.Timeout = timeout
	}
}

func NewServer(options ...Option) *Server {
	server := &Server{
		Addr:         "localhost",
		Port:         8080,
		ReadTimeout:  2 * time.Second,
		WriteTimeout: 2 * time.Second,
		Timeout:      4 * time.Second,
	}

	for _, option := range options {
		option(server)
	}

	return server
}

func main() {
	server1 := NewServer()
	fmt.Println(server1)

	add := WithAddr("127.0.0.1")
	timeout := WithTimeout(5 * time.Second)

	server2 := NewServer(add, timeout)
	fmt.Println(server2)
}
