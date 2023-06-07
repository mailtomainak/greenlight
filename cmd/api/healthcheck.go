package main

import (
	"net/http"
)

func (a *application) healthcheckHandler(writer http.ResponseWriter, _ *http.Request) {
	data := envelope{
		"status":      "available",
		"environment": a.config.env,
		"version":     version,
	}
	err := a.writeJSON(writer, http.StatusOK, data, nil)

	if err != nil {
		a.logger.Println(err)
		http.Error(writer, "The server encountered an error", http.StatusInternalServerError)
	}

}
