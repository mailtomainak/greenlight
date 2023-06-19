package main

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"github.com/mailtomainak/greenlight/internal/data"
	"net/http"
	"strconv"
	"time"
)

func (app *application) createMovieHandler(writer http.ResponseWriter, _ *http.Request) {
	_, _ = fmt.Fprintln(writer, "create a new movie")
}

func (app *application) showMovieHandler(writer http.ResponseWriter, request *http.Request) {
	params := httprouter.ParamsFromContext(request.Context())
	id, err := strconv.ParseInt(params.ByName("id"), 10, 64)
	if err != nil {
		app.notFoundResponse(writer, request)
		return
	}
	if id < 1 {
		writer.WriteHeader(http.StatusBadRequest)
		return
	}

	movie := data.Movie{
		ID:        id,
		CreatedAt: time.Now(),
		Title:     "Casablanca",
		Runtime:   102,
		Genres:    []string{"drama", "romance", "war"},
		Version:   1,
	}

	err = app.writeJSON(writer, http.StatusOK, envelope{"movie": movie}, nil)
	if err != nil {
		app.serverErrorResponse(writer, request, err)
	}
}
