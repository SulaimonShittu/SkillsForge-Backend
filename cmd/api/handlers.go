package main

import (
	"fmt"
	"net/http"
)

func (app *application) healthcheck(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}
	fmt.Fprintln(w, "status: available")
	fmt.Fprintf(w, "environment: %s\n", app.config.env)
	fmt.Fprintf(w, "version: %s\n", version)
}

func (app *application) getCreateBooksHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		fmt.Fprintln(w, `<h1>Hello World</h1>`)
		return
	}
	if r.Method == http.MethodPost {
		fmt.Fprintln(w, `<h1>Came as hard</h1>`)
		return
	}
}

func (app *application) getBooks(w http.ResponseWriter, r *http.Request) {

}
