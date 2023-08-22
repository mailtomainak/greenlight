package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
)

type envelope map[string]interface{}

func (app *application) writeJSON(w http.ResponseWriter, status int, data interface{}, headers http.Header) error {

	js, err := json.Marshal(data)
	if err != nil {
		return err
	}
	for k, h := range headers {
		w.Header()[k] = h
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_, _ = w.Write(js)
	return nil
}

func (app *application) readJSON(_ http.ResponseWriter, r *http.Request, obj interface{}) error {
	err := json.NewDecoder(r.Body).Decode(obj)
	if err != nil {
		var syntaxErr *json.SyntaxError
		var unmarshalTypeErr *json.UnmarshalTypeError
		var invalidUnmarshalErr *json.InvalidUnmarshalError

		switch {
		case errors.As(err, &syntaxErr):
			return fmt.Errorf("body contains badly formed JSON (at character %d)", syntaxErr.Offset)
		case errors.Is(err, io.ErrUnexpectedEOF):
			return errors.New("body contains badly formed JSON")
		case errors.As(err, &unmarshalTypeErr):
			if unmarshalTypeErr.Field != "" {
				return fmt.Errorf("body contains incorrect JSON type for field %q", unmarshalTypeErr.Field)
			}
			return fmt.Errorf("body contains incorrect JSON type (at character %d)", unmarshalTypeErr.Offset)
		case errors.Is(err, io.EOF):
			return errors.New("body must not be empty")
		case errors.As(err, &invalidUnmarshalErr):
			panic(err)
		default:
			return err

		}
	}
	return nil
}
