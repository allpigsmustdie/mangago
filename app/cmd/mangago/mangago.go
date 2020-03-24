package main

import (
	"context"
	"log"
	net "net/http"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

const shutdownTimeout = 5

func main() {
	app, err := InitApp()
	if err != nil {
		panic(err)
	}

	mainCtx, cancelMainCtx := context.WithCancel(context.Background())
	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		defer wg.Done()
		if err := app.Server.ListenAndServe(); err != net.ErrServerClosed {
			log.Printf("HTTP server ListenAndServe: %v\n", err)
		}
	}()

	signals := make(chan os.Signal, 1)
	// interrupt signal sent from terminal
	signal.Notify(signals, os.Interrupt)
	// sigterm signal sent from kubernetes
	signal.Notify(signals, syscall.SIGTERM)

	log.Printf("Got signal %v\n", <-signals)

	ctx, _ := context.WithTimeout(mainCtx, shutdownTimeout * time.Second)

	wg.Add(1)
	go func() {
		defer wg.Done()
		if err := app.Server.Shutdown(ctx); err != nil {
			log.Printf("HTTP server Shutdown: %v\n", err)
		}
	}()

	go func() {
		wg.Wait()
		cancelMainCtx()
	}()

	<-ctx.Done()
}
