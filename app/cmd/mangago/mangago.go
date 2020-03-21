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

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"

	"github.com/allpigsmustdie/mangago/app/interfaces/http"
	"github.com/allpigsmustdie/mangago/app/interfaces/http/rest"
	"github.com/allpigsmustdie/mangago/app/interfaces/repoitory"
	"github.com/allpigsmustdie/mangago/app/usecases"
)

const shutdownTimeout = 5

func main() {
	db, err := gorm.Open("sqlite3", ":memory:")
	if err != nil {
		panic(err)
	}
	mangaRepo := repoitory.NewManga(db)
	mangaService := usecases.NewMangaService(mangaRepo)
	restHandler := rest.NewHandler(mangaService)
	server := http.NewServer(restHandler)

	mainCtx, cancel := context.WithCancel(context.Background())
	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		defer wg.Done()
		if err := server.ListenAndServe(); err != net.ErrServerClosed {
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
		if err := server.Shutdown(ctx); err != nil {
			log.Printf("HTTP server Shutdown: %v\n", err)
		}
	}()

	go func() {
		wg.Wait()
		cancel()
	}()

	<-ctx.Done()
}
