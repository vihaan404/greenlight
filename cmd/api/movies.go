package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/vihaan404/greenlight/internal/data"
	"github.com/vihaan404/greenlight/internal/validator"
)

func (app *application) createMovieHandler(w http.ResponseWriter, r *http.Request) {
	var input struct {
		Title   string       `json:"title"`
		Year    int32        `json:"year"`
		Runtime data.Runtime `json:"runtime"`
		Genres  []string     `json:"genres"`
	}

	err := app.readJSON(w, r, &input)
	if err != nil {
		app.badRequestResponse(w, r, err)
		return
	}
	movie := &data.Movie{
		Title:   input.Title,
		Year:    input.Year,
		Runtime: input.Runtime,
		Genres:  input.Genres,
	}

	v := validator.New()
	data.ValidateMovie(v, movie)
	if !v.Valid() {
		app.failedValidationResponse(w, r, v.Errors)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	fmt.Fprintf(w, "Movie: %+v\n", movie)
}

func (app *application) showMovieHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	id, err := app.readIDParams(r)
	if err != nil {
		app.logError(r, err)
		app.notFoundResponse(w, r)
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
		app.serverErrorResponse(w, r, err)
	}
}
