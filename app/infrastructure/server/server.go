package server

import (
	"net/http"
	"time"
)

const (
	serverReadTimeout = 5
	serverWriteTimeout = 10
	serverIdleTimeout = 15
)

func NewServer(mux http.Handler) *http.Server {
	return &http.Server{
		Handler: mux,
		Addr:         ":8080", // TODO: from config
		ReadTimeout:  serverReadTimeout * time.Second,
		WriteTimeout: serverWriteTimeout * time.Second,
		IdleTimeout:  serverIdleTimeout * time.Second,
	}
}
