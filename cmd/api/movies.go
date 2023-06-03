package main

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"strconv"
)

func (a *application) createMovieHandler(writer http.ResponseWriter, _ *http.Request) {
	_, _ = fmt.Fprintln(writer, "create a new movie")
}

func (a *application) showMovieHandler(writer http.ResponseWriter, request *http.Request) {
	params := httprouter.ParamsFromContext(request.Context())
	id, err := strconv.ParseInt(params.ByName("id"), 10, 64)
	if err != nil {
		http.NotFound(writer, request)
		return
	}
	if id < 1 {
		writer.WriteHeader(http.StatusBadRequest)
		return
	}
	_, _ = fmt.Fprintf(writer, "Show details of movie -> %d\n", id)
}
