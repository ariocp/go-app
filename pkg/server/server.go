package server

import (
	"context"
	"github.com/ariocp/go-app/config"
	"net/http"
	"os"
	"os/signal"
	"time"
)

type Server struct {
	httpServer *http.Server
}

func NewServer(cfg config.Config, handler http.Handler) *Server {
	return &Server{httpServer: &http.Server{
		Addr:              cfg.Server.Host + ":" + cfg.Server.Port,
		Handler:           handler,
		ReadHeaderTimeout: 10 * time.Second,
		WriteTimeout:      10 * time.Second,
		MaxHeaderBytes:    1 << 20,
	}}
}

func (s *Server) Start() error {
	go func() {
		if err := s.httpServer.ListenAndServe(); err != nil {
			panic(err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit

	return s.Shutdown()
}

func (s *Server) Shutdown() error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	return s.httpServer.Shutdown(ctx)
}
