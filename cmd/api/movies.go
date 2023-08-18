package main

import (
	"encoding/json"
	"fmt"
	"github.com/julienschmidt/httprouter"
	"github.com/mailtomainak/greenlight/internal/data"
	"net/http"
	"strconv"
	"time"
)

func (app *application) createMovieHandler(writer http.ResponseWriter, request *http.Request) {
	var input struct {
		Title   string   `json:"title"`
		Year    int32    `json:"year"`
		Runtime int32    `json:"runtime"`
		Genres  []string `json:"genres"`
	}
	err := json.NewDecoder(request.Body).Decode(&input)
	if err != nil {
		app.errorResponse(writer, request, http.StatusBadRequest, err.Error())
		return
	}
	_, _ = fmt.Fprintf(writer, "%+v\n", input)
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
