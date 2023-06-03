package main

import (
	"encoding/json"
	"net/http"
)

func (a *application) healthcheckHandler(writer http.ResponseWriter, _ *http.Request) {
	data := map[string]string{
		"status":      "available",
		"environment": a.config.env,
		"version":     version,
	}

	js, err := json.Marshal(data)
	if err != nil {
		a.logger.Println(err)
		http.Error(writer, "The server encountered an error", http.StatusInternalServerError)
		return
	}
	writer.Header().Set("Content-Type", "application/json")
	_, _ = writer.Write(js)
}
