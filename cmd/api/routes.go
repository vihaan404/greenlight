package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func (app *application) routes() *httprouter.Router {
	router := httprouter.New()
	router.HandlerFunc(http.MethodGet, "/v1/health", app.healthCheckHandler)
	router.HandlerFunc(http.MethodPost, "/v1/movie/:id", app.createMovieHandler)
	router.HandlerFunc(http.MethodGet, "/v1/movie/:id", app.showMovieHandler)
	return router
}
