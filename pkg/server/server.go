package server

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/ariocp/go-app/config"
	"github.com/sirupsen/logrus"
)

type Server struct {
	*http.Server
}

func NewServer(cfg config.Config, handler http.Handler) *Server {
	httpServer := &http.Server{
		Addr:    cfg.Server.Host + ":" + cfg.Server.Port,
		Handler: handler,
	}
	return &Server{httpServer}
}

func (s *Server) Run() error {
	go func() {
		if err := s.ListenAndServe(); err != nil {
			return
		}
	}()

	logrus.Info()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit

	logrus.Info()

	if err := s.Shutdown(context.Background()); err != nil {
		return err
	}
	return nil
}
