package main

import (
	"RainmanwareKYC/internal/service"
	"context"
	"log"
	"time"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	middleware := service.NewMiddlewareCoreDefault()

	if err := middleware.Start(ctx); err != nil {
		log.Fatalf("Service exited with error: %v", err)
	}

	time.Sleep(1 * time.Second)
	log.Println("Middleware core stopped gracefully")
}
