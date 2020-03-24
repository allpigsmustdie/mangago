// +build wireinject

package main

import (
	"github.com/google/wire"

	"github.com/allpigsmustdie/mangago/app/infrastructure/db/sqlite"
	"github.com/allpigsmustdie/mangago/app/infrastructure/server"
	"github.com/allpigsmustdie/mangago/app/interfaces/http"
	"github.com/allpigsmustdie/mangago/app/interfaces/http/rest"
	"github.com/allpigsmustdie/mangago/app/interfaces/repoitory"
	"github.com/allpigsmustdie/mangago/app/usecases"
)

func InitApp() (App, error) {
	panic(wire.Build(
		NewApp,
		sqlite.InMemory, // TODO: postgreSQL
		repoitory.NewManga,
		usecases.NewMangaService,
		rest.NewHandler,
		http.NewMainHandler,
		server.NewServer,
	))
}
