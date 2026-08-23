package server

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/OkciD/whos_on_call/internal/shared/pkg/logger"
)

type Server struct {
	name   string
	config Config
	logger logger.Logger

	server http.Server
}

func New(name string, config Config, logger logger.Logger, mux http.Handler) *Server {
	return &Server{
		name:   name,
		config: config,
		logger: logger.WithField("serverName", name),

		server: http.Server{
			Addr:        config.ListenAddr,
			Handler:     mux,
			ReadTimeout: config.ReadTimeout.Duration,
		},
	}
}

func (s *Server) Start() error {
	s.logger.WithField("addr", s.server.Addr).Info("server starting")

	if err := s.server.ListenAndServe(); !errors.Is(err, http.ErrServerClosed) {
		s.logger.WithError(err).Error("listenAndServe error")

		return fmt.Errorf("server %s start error: %w", s.name, err)
	}

	s.logger.Info("stopped serving new connections")

	return nil
}

func (s *Server) Stop() error {
	shutdownCtx, shutdownRelease := context.WithTimeout(context.Background(), s.config.ShutdownTimeout.Duration)
	defer shutdownRelease()

	if err := s.server.Shutdown(shutdownCtx); err != nil {
		s.logger.WithError(err).Warn("server shutdown error, trying force close")

		if err = s.server.Close(); err != nil {
			s.logger.WithError(err).Error("server close error")

			return fmt.Errorf("failed to stop server %s: %w", s.name, err)
		}
	}

	return nil
}
