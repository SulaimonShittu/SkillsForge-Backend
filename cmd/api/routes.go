package main

import (
	"net/http"
)

func (app *application) route() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/v1/healthcheck", app.healthcheck)
	mux.HandleFunc("/v1/login", login)
	mux.HandleFunc("/v1/logout", logout)
	mux.HandleFunc("/v1/comments", app.getCreateCommentsHandler)
	mux.HandleFunc("/v1/comments/", app.getDeleteCommentsHandler)
	return mux
}
