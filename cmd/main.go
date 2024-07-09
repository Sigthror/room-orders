package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"
	"top-selection-test/internal/http"
	v1 "top-selection-test/internal/http/handlers/v1"
	"top-selection-test/internal/repository"
)

func main() {
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM, syscall.SIGHUP)
	defer cancel()

	orderRepository := repository.NewOrders()
	orderHandler := v1.NewOrders(orderRepository)
	r := http.NewRouter(orderHandler)
	srv := http.NewServer(http.DefaultConfig, r)
	if err := srv.ListenAndServe(ctx); err != nil {
		panic(err)
	}
}
