package main

import "net/http"

type App struct {
	Server *http.Server
}

func NewApp(server *http.Server) App {
	return App{
		server,
	}
}