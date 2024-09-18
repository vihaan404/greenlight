package main

import (
	"net/http"
)

func (app *application) healthCheckHandler(w http.ResponseWriter, r *http.Request) {
	data := map[string]string{
		"status":      "available",
		"version":     version,
		"environment": app.config.env,
	}
	err := app.writeJson(w, r, http.StatusOK, envelope{"health": data}, nil)
	if err != nil {
		app.errorResponse(w, r, http.StatusInternalServerError, err)
	}
}
