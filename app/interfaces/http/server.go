package http

import (
	"net/http"
	"time"
)

const (
 serverReadTimeout = 5
 serverWriteTimeout = 10
 serverIdleTimeout = 15
)

const (
	rest = "/rest"
)

func NewServer(restHandler RESTHandler) *http.Server {
	mux := http.NewServeMux()
	mux.Handle(rest+"/", http.StripPrefix(rest, restHandler))

	return &http.Server{
		Handler:      mux,
		Addr:         ":8080",
		ReadTimeout:  serverReadTimeout * time.Second,
		WriteTimeout: serverWriteTimeout * time.Second,
		IdleTimeout:  serverIdleTimeout * time.Second,
	}
}