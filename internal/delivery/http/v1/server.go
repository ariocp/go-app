package v1

import (
	"context"
	"github.com/ariocp/go-app/config"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

type Server struct {
	httpServer *http.Server
}

func NewServer(cfg config.Config, handler http.Handler) *Server {
	httpServer := &http.Server{
		Addr:    cfg.Server.Host + ":" + cfg.Server.Port,
		Handler: handler,
	}
	return &Server{httpServer: httpServer}
}

func (s *Server) Start() error {
	go func() {
		if err := s.httpServer.ListenAndServe(); err != nil {
			panic(err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit

	if err := s.Stop(); err != nil {
		panic(err)
	}
	return nil
}

func (s *Server) Stop() error {
	return s.httpServer.Shutdown(context.Background())
}
