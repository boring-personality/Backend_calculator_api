package main

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"

	"github.com/go-playground/validator"
)

type jsonResponse struct {
	Message string `json:"message,omitempty"`
	Result  int    `json:"result,omitempty"`
}

type jsonInput struct {
	Input1 int `json:"input1" validate:"required"`
	Input2 int `json:"input2" validate:"required"`
}

type jsonArrayInput struct {
	Inputarray []int `json:"inputarray" validate:"required"`
}

func writeJSON(w http.ResponseWriter, status int, data any, headers ...http.Header) error {
	out, err := json.Marshal(data)
	if err != nil {
		return err
	}

	if len(headers) > 0 {
		for key, value := range headers[0] {
			w.Header()[key] = value
		}
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_, err = w.Write(out)
	if err != nil {
		return err
	}

	return nil
}

func readJSON[T any](w http.ResponseWriter, r *http.Request, data *T) error {
	maxBytes := 1048576 // 1 Mb

	// var data jsonInput

	r.Body = http.MaxBytesReader(w, r.Body, int64(maxBytes))

	dec := json.NewDecoder(r.Body)
	err := dec.Decode(&data)

	if err != nil {
		return err
	}

	err = dec.Decode(&struct{}{})

	if err != io.EOF {
		return errors.New("body must have only a single JSON value")
	}

	return err
}

func inputValidator(data jsonInput) bool {
	validate := validator.New()
	if err := validate.Struct(&data); err != nil {
		return false
	}
	return true
}

func inputValidatorArray(data jsonArrayInput) bool {
	validate := validator.New()
	if err := validate.Struct(&data); err != nil {
		return false
	}
	return true
}
