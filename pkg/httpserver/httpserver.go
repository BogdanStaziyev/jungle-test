package httpserver

import (
	"context"
	"log"
	"net/http"
	"time"
)

const (
	_defaultReadTimeout     = 5 * time.Second
	_defaultWriteTimeout    = 5 * time.Second
	_defaultShutdownTimeout = 3 * time.Second
)

// Server - represents http server.
type Server struct {
	server          *http.Server
	notify          chan error
	shutdownTimeout time.Duration
}

// New - creates instance of new http server.
func New(handler http.Handler, opts ...Option) *Server {
	httpServer := &http.Server{
		Handler:      handler,
		ReadTimeout:  _defaultReadTimeout,
		WriteTimeout: _defaultWriteTimeout,
	}

	s := &Server{
		server:          httpServer,
		notify:          make(chan error, 1),
		shutdownTimeout: _defaultShutdownTimeout,
	}

	// add custom options
	for _, opt := range opts {
		opt(s)
	}

	s.start()

	return s
}

// Start - bootstraps http server.
func (s *Server) start() {
	log.Printf("Starting HTTP server on port %s", s.server.Addr)
	go func() {
		s.notify <- s.server.ListenAndServe()
		close(s.notify)
	}()
}

// Notify - returns error notification channel.
func (s *Server) Notify() <-chan error {
	return s.notify
}

// Shutdown - shuts down http server gracefully.
func (s *Server) Shutdown() error {
	ctx, cancel := context.WithTimeout(context.Background(), s.shutdownTimeout)
	defer cancel()

	return s.server.Shutdown(ctx)
}
