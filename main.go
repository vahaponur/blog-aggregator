package main

import (
	"context"
	"fmt"
	_ "github.com/lib/pq"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

var cfg *Config

func main() {

	cfg = createConfig()
	go fetchNext(5, 10)
	mainRouter := createRouter()
	http.Handle("/", mainRouter)

	server := &http.Server{Addr: cfg.env.PORT}
	fmt.Println("Starting the server...")
	go func() {
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			fmt.Printf("Error: %s\n", err)
		}
	}()
	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, os.Interrupt, syscall.SIGTERM)
	// Wait for signals
	<-signalChan
	fmt.Println("\nReceived interrupt signal. Shutting down gracefully...")

	// Shutdown the server gracefully
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		fmt.Printf("Error during server shutdown: %s\n", err)
	}

}
