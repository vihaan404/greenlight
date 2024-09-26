package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func (app *application) routes() *httprouter.Router {
	router := httprouter.New()

	router.NotFound = http.HandlerFunc(app.notFoundResponse)
	router.MethodNotAllowed = http.HandlerFunc(app.methodNotAllowedResponse)
	router.HandlerFunc(http.MethodGet, "/v1/health", app.healthCheckHandler)
	router.HandlerFunc(http.MethodPost, "/v1/movie", app.createMovieHandler)
	router.HandlerFunc(http.MethodGet, "/v1/movie/:id", app.showMovieHandler)
	router.HandlerFunc(http.MethodPost, "/v1/movie/:id", app.updateMovieHandler)

	router.HandlerFunc(http.MethodDelete, "/v1/movie/:id", app.deleteMovieHandler)
	return router
}
