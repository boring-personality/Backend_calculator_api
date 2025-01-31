package main

import (
	"encoding/json"
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

func inputValidator(data jsonInput) bool {
	validate := validator.New()
	if err := validate.Struct(&data); err != nil {
		return false
	}
	return true
}
