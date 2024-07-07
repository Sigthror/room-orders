package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"
	"top-selection-test/internal/http"
	"top-selection-test/internal/http/handlers"
)

func main() {
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM, syscall.SIGHUP)
	defer cancel()

	r := http.NewRouter(handlers.NewOrder(), nil)
	srv := http.NewServer(http.DefaultConfig, r)
	if err := srv.ListenAndServe(ctx); err != nil {
		panic(err)
	}
}
