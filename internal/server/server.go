package server

import (
	"context"
	"net/http"
	"time"
)

const (
	RWTimeout = 10 * time.Second

	port = "8000"
)

type Server struct {
	httpServer *http.Server
}

func NewServer(handler http.Handler) *Server {
	return &Server{
		httpServer: &http.Server{
			Addr:           ":" + port,
			Handler:        handler,
			ReadTimeout:    RWTimeout,
			WriteTimeout:   RWTimeout,
			MaxHeaderBytes: 1 << 20,
		},
	}
}

func (s *Server) Run() error {
	return s.httpServer.ListenAndServe()
}

func (s *Server) ShutDown(ctx context.Context) error {
	return s.httpServer.Shutdown(ctx)
}
