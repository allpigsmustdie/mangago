package http

import (
	"net/http"
)

const (
	restPath = "/rest"
)

func NewMainHandler(restHandler RESTHandler) http.Handler {
	mux := http.NewServeMux()
	mux.Handle(restPath+"/", http.StripPrefix(restPath, restHandler))
	return mux
}