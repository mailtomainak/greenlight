package main

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func (a *application) routes() *httprouter.Router {
	router := httprouter.New()
	router.HandlerFunc(http.MethodGet, "/v1/healthcheck", a.healthcheckHandler)
	router.HandlerFunc(http.MethodPost, "/v1/movies", a.createMovieHandler)
	router.HandlerFunc(http.MethodGet, "/v1/movies/:id", a.showMovieHandler)
	return router
}
