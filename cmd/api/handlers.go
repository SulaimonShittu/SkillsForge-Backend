package main

import (
	data "SkillsForge-Backend/cmd/internal"
	"encoding/json"
	"fmt"
	"net/http"
	"net/mail"
	"strconv"
	"time"
)

func (app *application) healthcheck(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}
	data := map[string]string{
		"status":      "available",
		"environment": app.config.env,
		"version":     version,
	}
	js, err := json.Marshal(data)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	js = append(js, '\n')
	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}

func (app *application) getCreateCommentsHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		comments := []data.Comment{
			{
				ID:         1,
				TimePosted: time.Now(),
				SenderName: "Paul Silas",
				Message:    "I want to make enquiry about the project.",
				Email: mail.Address{
					Address: "sulele04@gmail.com",
				},
			},
			{
				ID:         2,
				TimePosted: time.Now(),
				SenderName: "Peter thomas",
				Message:    "I want to make enquiry about the project.",
				Email: mail.Address{
					Address: "sulaimonshittu2004@gmail.com",
				},
			},
		}
		js, err := json.Marshal(comments)
		if err != nil {
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}
		js = append(js, '\n')
		w.Header().Set("Content-Type", "application/json")
		w.Write(js)
		return
	}
	if r.Method == http.MethodPost {
		fmt.Fprintln(w, `<h1>Came as hard</h1>`)
		return
	}
}

func (app *application) getDeleteCommentsHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		app.getComment(w, r)
	case http.MethodDelete:
		app.deleteComment(w, r)
	default:
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
	}
}

func (app *application) getComment(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Path[len("/v1/comments/"):]
	idInt, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}
	comment := data.Comment{
		ID:         idInt,
		TimePosted: time.Now(),
		SenderName: "Paul Silas",
		Message:    "I want to make enquiry about the project.",
		Email: mail.Address{
			Address: "sulele04@gmail.com",
		},
	}
	js, err := json.Marshal(comment)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	js = append(js, '\n')
	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}

func (app *application) deleteComment(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Path[len("/v1/comments/"):]
	idInt, err := strconv.Atoi(id)
	if err != nil {
		http.Error(w, "Bad Request", http.StatusBadRequest)
	}
	fmt.Println(idInt)
}
