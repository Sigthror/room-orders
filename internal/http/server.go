package http

import (
	"context"
	"fmt"
	"net"
	"net/http"
	"strconv"
	"time"
)

var (
	DefaultConfig = Config{
		Host:            "0.0.0.0",
		Port:            80,
		ShutdownTimeout: 5 * time.Second,
	}
)

type Config struct {
	Host            string
	Port            int
	ShutdownTimeout time.Duration
}

type Server struct {
	srv             *http.Server
	shutdownTimeout time.Duration
}

func NewServer(cfg Config, r http.Handler) *Server {
	srv := &http.Server{
		Addr:    net.JoinHostPort(cfg.Host, strconv.Itoa(cfg.Port)),
		Handler: r,
	}

	return &Server{
		srv:             srv,
		shutdownTimeout: cfg.ShutdownTimeout,
	}
}

func (s *Server) ListenAndServe(ctx context.Context) error {
	// TODO make logger for context
	sCtx, cancelCause := context.WithCancelCause(ctx)
	defer cancelCause(nil)

	go func() {
		fmt.Printf("starting http server at %s\n", s.srv.Addr)
		if err := s.srv.ListenAndServe(); err != nil {
			cancelCause(err)
		}
	}()

	<-sCtx.Done()
	if err := context.Cause(sCtx); err != nil && err != context.Canceled {
		return fmt.Errorf("failed to listen and serve %s: %w", s.srv.Addr, err)
	}

	sCtx, cancel := context.WithTimeout(ctx, s.shutdownTimeout)
	defer cancel()
	if err := s.srv.Shutdown(sCtx); err != nil && err != sCtx.Err() {
		return fmt.Errorf("failed gracefull shutdown the server: %w", err)
	}

	fmt.Println("gracefully shutdown the server")

	return nil
}
