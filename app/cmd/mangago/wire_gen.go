// Code generated by Wire. DO NOT EDIT.

//go:generate wire
//+build !wireinject

package main

import (
	"github.com/allpigsmustdie/mangago/app/infrastructure/db/sqlite"
	"github.com/allpigsmustdie/mangago/app/infrastructure/server"
	"github.com/allpigsmustdie/mangago/app/interfaces/http"
	"github.com/allpigsmustdie/mangago/app/interfaces/http/rest"
	"github.com/allpigsmustdie/mangago/app/interfaces/repoitory"
	"github.com/allpigsmustdie/mangago/app/usecases"
)

// Injectors from wire.go:

func InitApp() (App, error) {
	db, err := sqlite.InMemory()
	if err != nil {
		return App{}, err
	}
	manga, err := repoitory.NewManga(db)
	if err != nil {
		return App{}, err
	}
	serviceManga := usecases.NewMangaService(manga)
	restHandler := rest.NewHandler(serviceManga)
	handler := http.NewMainHandler(restHandler)
	httpServer := server.NewServer(handler)
	app := NewApp(httpServer)
	return app, nil
}
