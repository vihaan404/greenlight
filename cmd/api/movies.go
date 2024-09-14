package main

import (
	"net/http"
	"time"

	"github.com/vihaan404/greenlight/internal/data"
)

func (app *application) createMovieHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(`{"message": "movie created"}`))
}

func (app *application) showMovieHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	id, err := app.readIDParams(r)
	if err != nil {
		http.NotFound(w, r)
		return
	}

	movie := data.Movie{
		ID:       id,
		Title:    "Inception",
		CratedAt: time.Now(),
		Year:     2010,
		Runtime:  148,
		Genres:   []string{"Action", "Adventure", "Sci-Fi"},
		Version:  1,
	}

	err = app.writeJson(w, r, http.StatusOK, envelope{"movie": movie}, nil)
	if err != nil {
		app.logger.Println(err)
		http.Error(w, "the server encountered a problem and could not process your request", http.StatusInternalServerError)
	}
}
