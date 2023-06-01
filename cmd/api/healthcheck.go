package main

import (
	"fmt"
	"net/http"
)

func (a application) healthcheckHandler(writer http.ResponseWriter, _ *http.Request) {
	_, _ = fmt.Fprintln(writer, "status available")
	_, _ = fmt.Fprintf(writer, "environment %s\n", a.config.env)
	_, _ = fmt.Fprintf(writer, "version %s\n", version)
}
