package httpserver

import (
	"log"
	"net"
	"time"
)

// Option - represents http server option.
type Option func(*Server)

// Port - configures http server port.
func Port(port string) Option {
	if port == "" {
		log.Fatal("Port is empty")
	}
	return func(s *Server) {
		s.server.Addr = net.JoinHostPort("", port)
	}
}

// ReadTimeout - configures http server read timeout.
func ReadTimeout(timeout time.Duration) Option {
	return func(s *Server) {
		s.server.ReadTimeout = timeout
	}
}

// WriteTimeout - configures http server read timeout.
func WriteTimeout(timeout time.Duration) Option {
	return func(s *Server) {
		s.server.WriteTimeout = timeout
	}
}

// ShutdownTimeout - configures http server shutdown timeout.
func ShutdownTimeout(timeout time.Duration) Option {
	return func(s *Server) {
		s.shutdownTimeout = timeout
	}
}
