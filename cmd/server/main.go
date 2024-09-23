package main

import (
	"context"
	"log"
	"os/signal"
	"syscall"

	"github.com/denislyubo/word-of-wisdom/internal/config"
	"github.com/denislyubo/word-of-wisdom/internal/usecase/server"
)

func main() {
	ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGSTOP, syscall.SIGTERM)
	go func() {
		<-ctx.Done()
		log.Println("Server: ", "Shutting down...")
		cancel()
	}()

	var cfg config.ServerConfig
	c, err := config.Load(cfg)
	if err != nil {
		log.Fatal("Server: ", "failed to load config: ", err.Error())
	}

	if err := server.New(c).ListenAndServe(ctx); err != nil {
		log.Println("Server: ", err)
	}
}
