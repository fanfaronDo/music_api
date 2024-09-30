package server

import (
	"context"
	"github.com/fanfaronDo/music_api/pkg/config"
	"net/http"
)

type Server struct {
	httpServer *http.Server
}

func (s *Server) Run(host string, httpHandler http.Handler, cfg config.HttpServer) error {
	s.httpServer = &http.Server{
		Addr:           host,
		Handler:        httpHandler,
		MaxHeaderBytes: 1 << 20,
		ReadTimeout:    cfg.Timeout,
		WriteTimeout:   cfg.Timeout,
		IdleTimeout:    cfg.IdleTimeout,
	}
	return s.httpServer.ListenAndServe()
}

func (s *Server) Shutdown(ctx context.Context) error {
	return s.httpServer.Shutdown(ctx)
}
