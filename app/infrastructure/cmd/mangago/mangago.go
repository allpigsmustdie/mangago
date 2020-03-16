package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"

	"github.com/allpigsmustdie/mangago/app/interfaces/repoitory"
	"github.com/allpigsmustdie/mangago/app/interfaces/rest"
	"github.com/allpigsmustdie/mangago/app/usecases"
)

func main() {

	const shutdownTimeout = 5

	mux := http.NewServeMux()

	db, err := gorm.Open("sqlite3", ":memory:")
	if err != nil {
		panic(err)
	}
	mangaRepo := repoitory.NewManga(db)
	mangaService := usecases.NewMangaService(mangaRepo)
	restHandler := rest.NewHandler(mangaService)

	const api = "/api"
	mux.Handle(api+"/", http.StripPrefix(api, restHandler))

	srv := http.Server{
		Handler: mux,
		Addr: ":8080",
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  15 * time.Second,
	}

	idleConnsClosed := make(chan struct{})

	go func() {
		sigint := make(chan os.Signal, 1)

		// interrupt signal sent from terminal
		signal.Notify(sigint, os.Interrupt)
		// sigterm signal sent from kubernetes
		signal.Notify(sigint, syscall.SIGTERM)

		<-sigint
		ctx, cancel := context.WithTimeout(context.Background(), shutdownTimeout * time.Second)
		defer cancel()

		if err := srv.Shutdown(ctx); err != nil {
			log.Printf("HTTP server Shutdown: %v\n", err)
		}
		close(idleConnsClosed)
	}()

	if err := srv.ListenAndServe(); err != http.ErrServerClosed {
		log.Printf("HTTP server ListenAndServe: %v\n", err)
	}

	<-idleConnsClosed
}
